package miner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"sort"
	"sync/atomic"
	//"strconv"
	"sync"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/params"
	"github.com/MOACChain/MoacLib/state"
	"github.com/MOACChain/MoacLib/types"
	"github.com/MOACChain/MoacLib/vm"
	"github.com/MOACChain/xchain/consensus"
	"github.com/MOACChain/xchain/consensus/pos"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/event"
	xparams "github.com/MOACChain/xchain/params"
)

const (
	ChainHeadChanSize = 10
)

// Agent can register themself with the worker
type Agent interface {
	Work() chan<- *Work
	SetReturnCh(chan<- *Result)
	Stop()
	Start()
	GetHashRate() int64
}

// Work is the workers current environment and holds
// all of the current state information
type Work struct {
	config    *params.ChainConfig
	signer    types.Signer
	state     *state.StateDB // apply state changes here
	txsCount  int            // tx count in cycle
	Block     *types.Block   // the new block
	header    *types.Header
	txs       []*types.Transaction
	receipts  []*types.Receipt
	createdAt time.Time
}

// worker is the main object which takes care of applying messages to the new state
type worker struct {
	config           *params.ChainConfig
	chain            *core.BlockChain
	pos              pos.Pos
	mu               sync.Mutex
	mux              *event.TypeMux // update loop
	blockGenTimerCh  chan core.BlockGenTimerEvent
	blockGenTimerSub event.Subscription
	currentMu        sync.Mutex
	blockGenTicker   *core.Ticker
	blockGenEpochBLS chan int64
	blockGenRunning  bool
	mc               Backend
	coinbase         common.Address
	mining           int32
	atWork           int32
	agents           map[Agent]struct{}
	current          *Work
	wg               sync.WaitGroup
	recv             chan *Result
	extra            []byte
}

type Result struct {
	Work  *Work
	Block *types.Block
}

func NewWorker(
	config *params.ChainConfig,
	engine consensus.Engine,
	coinbase common.Address,
	mc Backend,
	mux *event.TypeMux,
) *worker {
	log.Debugf("create new chain worker")
	worker := &worker{
		config:           config,
		mux:              mux,
		blockGenTimerCh:  make(chan core.BlockGenTimerEvent, ChainHeadChanSize),
		chain:            mc.BlockChain(),
		pos:              engine.(pos.Pos),
		blockGenEpochBLS: make(chan int64),
		mc:               mc,
	}

	worker.blockGenTimerSub = mc.BlockChain().SubscribeBlockGenTimerEvent(worker.blockGenTimerCh)
	go worker.update()

	return worker
}

func newWorkerAsMonitor(config *params.ChainConfig, bc *core.BlockChain) *worker {
	worker := &worker{
		config:          config,
		mux:             new(event.TypeMux),
		blockGenTimerCh: make(chan core.BlockGenTimerEvent, ChainHeadChanSize),
		chain:           bc,
	}

	worker.blockGenTimerSub = bc.SubscribeBlockGenTimerEvent(worker.blockGenTimerCh)
	go worker.update()

	return worker
}

func (worker *worker) setMoacbase(addr common.Address) {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.coinbase = addr
}

func (worker *worker) setExtra(extra []byte) {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.extra = extra
}

func (worker *worker) pending() (*types.Block, *state.StateDB) {
	worker.currentMu.Lock()
	defer worker.currentMu.Unlock()

	if atomic.LoadInt32(&worker.mining) == 0 {
		return types.NewBlock(
			worker.current.header,
			worker.current.txs,
			nil,
			worker.current.receipts,
		), worker.current.state.Copy()
	}
	return worker.current.Block, worker.current.state.Copy()
}

func (worker *worker) pendingBlock() *types.Block {
	worker.currentMu.Lock()
	defer worker.currentMu.Unlock()

	if atomic.LoadInt32(&worker.mining) == 0 {
		return types.NewBlock(
			worker.current.header,
			worker.current.txs,
			nil,
			worker.current.receipts,
		)
	}
	return worker.current.Block
}

func (worker *worker) start() {
	worker.mu.Lock()
	defer worker.mu.Unlock()

	// bls loop
	log.Infof("run vss worker")
	go worker.blockGenBLSBeacon()
	go worker.blockGenBLSLoop()

	atomic.StoreInt32(&worker.mining, 1)
}

func (worker *worker) stop() {
	worker.wg.Wait()
	worker.mu.Lock()
	defer worker.mu.Unlock()
	atomic.StoreInt32(&worker.mining, 0)
	atomic.StoreInt32(&worker.atWork, 0)
}

func (worker *worker) register(agent Agent) {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.agents[agent] = struct{}{}
	agent.SetReturnCh(worker.recv)
}

func (worker *worker) unregister(agent Agent) {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	delete(worker.agents, agent)
	agent.Stop()
}

func (worker *worker) UpdateTimer(length int) {
	if worker.blockGenTicker != nil && length > 0 {
		//remove current ticker
		worker.blockGenTicker.Period = time.Duration(length) * time.Second
		worker.blockGenTicker.ResetTicker()
	}
}

func (worker *worker) StopTimer() {
	// blockGenTicker of monitor scs is nil
	if worker.blockGenTicker != nil {
		worker.blockGenTicker.Stop()
	}
}

func (worker *worker) blockGenBLSBeacon() {
	ticker := time.NewTicker(50 * time.Millisecond)
	timeMillisecond := time.Now().UnixNano() / int64(time.Millisecond)
	epochCurrent := timeMillisecond / 10000 * 10000
	epochPrev := epochCurrent

	go func() {
		for {
			<-ticker.C
			timeMillisecond := time.Now().UnixNano() / int64(time.Millisecond)
			epochCurrent = timeMillisecond / 10000 * 10000
			if epochCurrent != epochPrev {
				epoch := epochCurrent / 10000
				worker.blockGenEpochBLS <- epoch
				epochPrev = epochCurrent
			}
		}
	}()
}

// Use the bls signature in current head block to determine the next proposer
func (worker *worker) getProposer(block *types.Block, epoch int64) common.Address {
	// get the random number based on bls signature from head block and epoch
	extraData, _ := block.Header().ExtraData()
	blssig := (*extraData)[params.BLSSigField]
	rounds := 100
	seed256 := sha256.Sum256([]byte(fmt.Sprintf("%x,%d", blssig, epoch)))
	for i := 0; i < rounds; i++ {
		seed256 = sha256.Sum256(seed256[:32])
	}
	seed64 := common.BytesToInt64(seed256[:8])
	r := rand.New(rand.NewSource(seed64))

	// sort the nodelist
	nodelist := worker.pos.NodeList
	nodelistSorted := []string{}
	for _, node := range nodelist {
		nodelistSorted = append(nodelistSorted, fmt.Sprintf("%s", node.Bytes()))
	}
	sort.Strings(nodelistSorted)

	// determine the proposer
	proposer := nodelistSorted[r.Intn(len(nodelistSorted))]

	return common.BytesToAddress([]byte(proposer))
}

func (worker *worker) genSigShare(epoch int64) pos.SigShareMessage {
	var sigShareMessage pos.SigShareMessage
	if worker.pos.IsVSSReady() {
		// prepare fields to sign
		currentblock := worker.chain.CurrentBlock()
		blockHash := currentblock.Hash().Bytes()
		nextBlockNumber := new(big.Int).Add(currentblock.Number(), big.NewInt(1)).Int64()
		proposer := worker.getProposer(currentblock, epoch)

		// format the toSign string and sign with bls
		toSign := fmt.Sprintf(
			"%x,%x,%d,%d",
			blockHash,
			proposer.Bytes(),
			nextBlockNumber,
			epoch,
		)
		sig := worker.pos.Bls.SignBytes([]byte(toSign))
		log.Infof(
			"gen sig share hash: %x, epoch: %d, proposer: %x, sig: %x",
			blockHash[:8], epoch, proposer.Bytes(), sig,
		)

		// generate the sig share message
		sigShareMessage = pos.SigShareMessage{
			FromIndex:   worker.pos.Bls.NodeIndex,
			Sig:         sig,
			ForProposer: proposer.Bytes(),
			BlockNumber: nextBlockNumber,
			BlockHash:   blockHash,
			Epoch:       epoch,
		}
	}

	return sigShareMessage
}

func (worker *worker) sendSigShare(sigShare *pos.SigShareMessage) error {
	// broadcast to other peers
	worker.mux.Post(core.NewSigShareEvent{SigShare: sigShare})
	// send one to self by directly sending to the channel
	worker.pos.Vss.SigShareChan <- sigShare
	return nil
}

func (worker *worker) blockGenBLS(sigShareMessage pos.SigShareMessage) {
	if worker.pos.IsVSSReady() {
		//commitNewWorkTime := int(time.Now().Unix())
		log.Debugf("Generate new block vss, current head: %d", worker.chain.CurrentBlock())
		work := worker.commitNewWork(&sigShareMessage)
		log.Debugf("after commitNewWork() vss: work = %v", work)
		var block *types.Block
		if work != nil {
			if block = worker.GenBlock(work); block != nil {
				// update block proposer
				//extraData, _ := block.Header().ExtraData()
				//epoch := (*extraData)[params.EpochField]
				//epochInt, _ := strconv.Atoi(string(epoch))
				/*
					worker.pos.RecordBlockInfo(
						block.Header().Coinbase.String(),
						int(block.Header().Number.Int64()),
						block.Header().Hash().Bytes(),
						block.Header().ParentHash.Bytes(),
						epochInt,
					)*/
			}
		}
	} else {
		log.Infof("block gen bls epoch: BLS Not Ready")
	}
}

func (worker *worker) tryBlockGenBLS(sigShareMessage pos.SigShareMessage) {
	// make sure at anytime, there is just one instance
	// of this gorouting
	// todo: use trylock once golang supports it in mutex
	worker.blockGenRunning = true
	defer func() {
		log.Debugf("in tryblockgenbls(), end block generation")
		worker.blockGenRunning = false
	}()

	log.Debugf("in tryblockgenbls(), begin block generation")
	if !worker.pos.IsConsensus() {
		log.Debugf("vss scs not in consensus group, will continue")
		return
	} else {
		log.Debugf("vss scs in consensus group, will participate")
	}

	// if I'm proposer, I should collect all sigs and generate the block
	// otherwise no-op
	proposer := common.BytesToAddress(sigShareMessage.ForProposer)
	if proposer == worker.pos.Vssid {
		log.Debugf(
			"I am the proposer vss %s, self: %x, proposer: %x",
			sigShareMessage.Key(),
			worker.pos.Vssid.Bytes(),
			sigShareMessage.ForProposer,
		)
		worker.blockGenBLS(sigShareMessage)
	} else {
		log.Debugf(
			"I am NOT the proposer vss %s, self: %x, proposer: %x",
			sigShareMessage.Key(),
			worker.pos.Vssid.Bytes(),
			sigShareMessage.ForProposer,
		)
	}
}

func (worker *worker) blockGenBLSLoop() {
	for {
		select {
		case epoch := <-worker.blockGenEpochBLS:
			if worker.pos.IsVSSReady() {
				sigShareMessage := worker.genSigShare(epoch)
				// broadcast sig shares at the begining of every epoch
				go worker.sendSigShare(&sigShareMessage)

				// try generating block but avoid concurrent runs
				if !worker.blockGenRunning {
					go worker.tryBlockGenBLS(sigShareMessage)
				}
			} else {
				log.Errorf("sci bls is not ready, can not generate block")
			}
		}
	}
}

func (worker *worker) update() {
	defer worker.blockGenTimerSub.Unsubscribe()
	for {
		// A real event arrived, process interesting content
		select {
		case blk := <-worker.blockGenTimerCh:
			// Update backupscs to nodelist if need
			if worker.pos.IsConsensus() {
				//no need to create new work here.
				//get block generator index in vnode info
				loc, nodecnt := worker.pos.FindAddressInNodelist(blk.Block.Header().Coinbase)
				if loc < 0 {
					continue
				}

				newwait := 0
				if loc >= worker.pos.Position {
					newwait = (worker.pos.Position - loc + nodecnt) * int(xparams.BlockInterval)
				} else {
					newwait = (worker.pos.Position - loc) * int(xparams.BlockInterval)
				}
				// now reset block gen timer
				log.Debugf(
					"blockGenTimerCh UpdateTimer:Position: %v, newwait: %v",
					worker.pos.Position,
					newwait,
				)
				worker.UpdateTimer(newwait)
			}

		// System stopped
		case err := <-worker.blockGenTimerSub.Err():
			log.Error("blockGenTimerSub", "err", err)
			return
		}
	}
}

func (worker *worker) GenBlock(work *Work) *types.Block {
	// seal block
	block, _ := worker.chain.Engine().Seal(
		worker.chain,
		work.Block,
		make(chan struct{}),
	)
	if block == nil {
		log.Error("Failed to seal block")
		return nil
	}

	// Update the block hash in all logs since it is now available and not when the
	// receipt/log of individual transactions were created.
	for _, receipt := range work.receipts {
		for _, log := range receipt.Logs {
			log.BlockHash = block.Hash()
		}
	}
	for _, log := range work.state.Logs() {
		log.BlockHash = block.Hash()
	}

	stat, err := worker.chain.WriteBlockAndState(block, work.receipts, work.state)
	if err != nil {
		log.Errorf("Failed writing block to chain: number %d, hash: %x, error: %v", block.Number(), block.Hash().Bytes()[:8], err)
		return nil
	}

	// broadcast the block
	worker.mux.Post(core.NewMinedBlockEvent{Block: block})
	var (
		events []interface{}
		logs   = work.state.Logs()
	)
	events = append(events, core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
	if stat == core.CanonStatTy {
		events = append(events, core.ChainHeadEvent{Block: block})
		events = append(events, core.BlockGenTimerEvent{Block: block})
	}

	worker.chain.PostChainEvents(events, logs)
	return block
}

// makeNewWork creates a new environment for the current cycle.
func (worker *worker) makeNewWork(parent *types.Block, header *types.Header) (*Work, error) {
	state, err := worker.chain.StateAt(parent.Root(), parent.Number())
	if err != nil {
		log.Error("worker StateAt", "err", err)
		return nil, err
	}
	work := &Work{
		config:    worker.config,
		signer:    types.NewPanguSigner(worker.config.ChainId),
		state:     state,
		header:    header,
		createdAt: time.Now(),
	}

	// Keep track of transactions which return errors so they can be removed
	work.txsCount = 0
	return work, nil
}

func (worker *worker) commitNewWork(sigShareMessage *pos.SigShareMessage) *Work {
	worker.mu.Lock()
	defer worker.mu.Unlock()
	worker.currentMu.Lock()
	defer worker.currentMu.Unlock()

	tstart := time.Now()
	parent := worker.chain.CurrentBlock()
	if parent == nil {
		log.Error("Failed to get current block.")
		return nil
	}

	tstamp := tstart.Unix()
	if parent.Header().Time.Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Header().Time.Int64() + 1
	}
	// this will ensure we're not going off too far in the future
	if now := time.Now().Unix(); tstamp > now+1 {
		wait := time.Duration(tstamp-now) * time.Second
		log.Debug("Mining too far in the future", "wait", common.PrettyDuration(wait))
		time.Sleep(wait)
	}

	num := parent.Number()
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     num.Add(num, common.Big1),
		Time:       big.NewInt(tstamp),
	}
	header.Coinbase = worker.pos.Vssid
	log.Debugf("commitNewWork parent num: %d, parent hash %x", num, header.ParentHash.Bytes())
	log.Debugf("engine type: %v", reflect.TypeOf(worker.chain.Engine()))

	// make new work
	work, err := worker.makeNewWork(parent, header)
	if err != nil {
		log.Error("Failed to create mining context", "err", err)
		return nil
	}

	pending, err := worker.mc.TxPool().Pending()
	if err != nil {
		log.Error("Failed to fetch pending transactions", "err", err)
		return nil
	}
	txs := types.NewTransactionsByPriceAndNonce(work.signer, pending)
	work.commitTransactions(worker.mux, txs, worker.chain)

	extraData := make(map[int][]byte)
	var blsSig []byte
	if !worker.pos.ShouldSkipVSS(header.Number.Int64()) {
		var err error
		blsSig, err = worker.pos.GetBlsSig(sigShareMessage.Key())
		log.Debugf("vss bls sig: %x, %v", blsSig, err)
		if err != nil {
			log.Debugf("Failed to get vss bls sig")
			return nil
		}

		// proposer will sign the bls signature with their vss private key
		// the vss public key is registered on the subchainbase which
		// can be used by other partys to verify the signature
		proposerSig, _ := pos.SignED25519(worker.pos.VssKey.Private, blsSig)

		epoch := fmt.Sprintf("%d", sigShareMessage.Epoch)
		extraData[params.BLSSigField] = blsSig
		extraData[params.EpochField] = []byte(epoch)
		extraData[params.ParentHashField] = header.ParentHash.Bytes()
		extraData[params.ProposerSig] = proposerSig
	}
	header.SetExtraData(&extraData)

	// Create the new block to seal with the consensus engine
	uncles := []*types.Header{}
	liveFlag := false
	if work.Block, err = worker.chain.Engine().Finalize(
		worker.chain, header, work.state, work.txs, uncles, work.receipts, liveFlag); err != nil {
		log.Error("Failed to finalize block for sealing", "err", err)
		return nil
	}
	// We only care about logging if we're actually mining.
	log.Infof(
		"Commit new mining work number %d, hash: %x, txs: %d, elapsed %v",
		work.Block.Number(),
		work.Block.Hash().Bytes()[:8],
		work.txsCount,
		common.PrettyDuration(time.Since(tstart)),
	)

	return work
}

func (worker *Work) commitTransactions(
	mux *event.TypeMux, txs *types.TransactionsByPriceAndNonce, bc *core.BlockChain,
) {
	gp := new(core.GasPool).AddGas(params.TargetGasLimit)
	for {
		if remain := gp.Value(); remain < params.TxGas || remain > params.TargetGasLimit.Uint64() {
			log.Trace("Not enough gas for further transactions", "have", remain, "want", params.TxGas)
			break
		}
		// Retrieve the next transaction and abort if all done
		tx := txs.Peek()
		if tx == nil {
			break
		}

		from, err := types.Sender(worker.signer, tx)
		if err != nil {
			continue
		}

		// Start executing the transaction
		worker.state.Prepare(tx.Hash(), common.Hash{}, worker.txsCount)
		err, _ = worker.commitTransaction(tx, bc, gp)
		switch err {
		case core.ErrGasLimitReached:
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Errorf(
				"GasRemaining limit exceeded for current block, sender: %x, tx hash: %x",
				from.Bytes()[:8], tx.Hash().Bytes()[:8],
			)
			txs.Pop()

		case core.ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Errorf(
				"Skipping transaction with low nonce, sender: %x, nonce: %d",
				from.Bytes()[:8], tx.Nonce(),
			)
			txs.Shift()

		case core.ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Errorf(
				"Skipping account with hight nonce, sender: %x, nonce: %d",
				from.Bytes()[:8], tx.Nonce(),
			)
			txs.Pop()

		case nil:
			worker.txsCount++
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Error("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}
	}
}

func (worker *Work) commitResetTransactions(
	mux *event.TypeMux, txs types.Transactions, bc *core.BlockChain,
) {
	gp := new(core.GasPool).AddGas(params.TargetGasLimit)
	for i := 0; i < len(txs); i++ {
		if remain := gp.Value(); remain < params.TxGas || remain > params.TargetGasLimit.Uint64() {
			log.Trace("Not enough gas for further transactions", "have", remain, "want", params.TxGas)
			break
		}
		// Retrieve the next transaction and abort if all done
		tx := txs[i]
		if tx == nil {
			break
		}

		from, err := types.Sender(worker.signer, tx)
		if err != nil {
			continue
		}

		// Start executing the transaction
		worker.state.Prepare(tx.Hash(), common.Hash{}, worker.txsCount)
		err, _ = worker.commitTransaction(tx, bc, gp)
		switch err {
		case core.ErrGasLimitReached:
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Errorf(
				"GasRemaining limit exceeded for current block, sender: %x",
				from.Bytes()[:8],
			)
		case core.ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Errorf(
				"Skipping transaction with low nonce, sender: %x, nonce: %d",
				from.Bytes()[:8], tx.Nonce(),
			)
		case core.ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Errorf(
				"Skipping account with hight nonce, sender: %x, nonce: %d",
				from.Bytes()[:8], tx.Nonce(),
			)
		case nil:
			worker.txsCount++
		default:
			log.Errorf(
				"Transaction failed, account skipped, hash: %x, err: %s",
				tx.Hash().Bytes()[:8], err,
			)
		}
	}
}

func (worker *Work) commitTransaction(
	tx *types.Transaction, bc *core.BlockChain, gp *core.GasPool,
) (error, *types.Receipt) {
	snap := worker.state.Snapshot()
	totalUsedGas := big.NewInt(0)
	receipt, _, err := core.ApplyTransaction(
		worker.config,
		bc,
		nil,
		gp,
		worker.state,
		worker.header,
		tx,
		totalUsedGas,
		vm.Config{},
		nil,
	)
	if err != nil {
		worker.state.RevertToSnapshot(snap)
		return err, nil
	}
	worker.txs = append(worker.txs, tx)
	if receipt != nil {
		worker.receipts = append(worker.receipts, receipt)
	}

	return nil, receipt
}
