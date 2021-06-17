// Copyright 2015 The MOAC-core Authors
// This file is part of the MOAC-core library.
//
// The MOAC-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The MOAC-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the MOAC-core library. If not, see <http://www.gnu.org/licenses/>.

// +build bls

package miner

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
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
)

const (
	// txChanSize is the size of channel listening to TxPreEvent.
	// The number is referenced from the size of tx pool.
	txChanSize = 4096
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
	// chainSideChanSize is the size of channel listening to ChainSideEvent.
	chainSideChanSize = 10
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
	tcount    int
}

// worker is the main object which takes care of applying messages to the new state
type worker struct {
	config           *params.ChainConfig
	chain            *core.BlockChain
	pos              *pos.Pos
	mu               sync.Mutex
	mux              *event.TypeMux // update loop
	currentMu        sync.Mutex
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

	// subscribe to channels
	txCh         chan core.TxPreEvent
	txSub        event.Subscription
	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription
	chainSideCh  chan core.ChainSideEvent
	chainSideSub event.Subscription
}

type Result struct {
	Work  *Work
	Block *types.Block
}

func newWorker(
	config *params.ChainConfig,
	engine consensus.Engine,
	coinbase common.Address,
	mc Backend,
	mux *event.TypeMux,
) *worker {
	log.Infof("create new chain worker")
	worker := &worker{
		config:           config,
		mux:              mux,
		txCh:             make(chan core.TxPreEvent, txChanSize),
		chainHeadCh:      make(chan core.ChainHeadEvent, chainHeadChanSize),
		chainSideCh:      make(chan core.ChainSideEvent, chainSideChanSize),
		chain:            mc.BlockChain(),
		pos:              engine.(*pos.Pos),
		blockGenEpochBLS: make(chan int64),
		mc:               mc,
		agents:           make(map[Agent]struct{}),
		coinbase:         coinbase,
	}

	worker.txSub = mc.TxPool().SubscribeTxPreEvent(worker.txCh)
	worker.chainHeadSub = mc.BlockChain().SubscribeChainHeadEvent(worker.chainHeadCh)
	worker.chainSideSub = mc.BlockChain().SubscribeChainSideEvent(worker.chainSideCh)

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
	var blssig []byte
	// if it's the genesis block
	if block.Number().Int64() == 0 {
		blssig = make([]byte, 32, 32)
	} else {
		blssig = (*extraData)[params.BLSSigField]
	}
	rounds := 100
	seed256 := sha256.Sum256([]byte(fmt.Sprintf("%x,%d", blssig, epoch)))
	for i := 0; i < rounds; i++ {
		seed256 = sha256.Sum256(seed256[:32])
	}
	seed64 := common.BytesToInt64(seed256[:8])
	r := rand.New(rand.NewSource(seed64))

	// sort the nodelist
	nodelist := worker.pos.NodeList
	log.Debugf("worker.pos.NodeList: %d", len(worker.pos.NodeList))
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
		currentBlock := worker.chain.CurrentBlock()
		blockHash := currentBlock.Hash().Bytes()
		nextBlockNumber := new(big.Int).Add(currentBlock.Number(), big.NewInt(1)).Int64()
		proposer := worker.getProposer(currentBlock, epoch)

		// format the toSign string and sign with bls
		toSign := fmt.Sprintf(
			"%x,%x,%d,%d",
			blockHash,
			proposer.Bytes(),
			nextBlockNumber,
			epoch,
		)
		sig := worker.pos.Bls.SignBytes([]byte(toSign))
		log.Debugf(
			"gen sig share,block hash: %x, epoch: %d, proposer: %x, sig: %x",
			blockHash[:8], epoch, proposer.Bytes(), sig,
		)

		// generate the sig share message
		sigShareMessage = pos.SigShareMessage{
			FromIndex:   big.NewInt(int64(worker.pos.Bls.NodeIndex)),
			Sig:         sig,
			ForProposer: proposer.Bytes(),
			BlockNumber: big.NewInt(nextBlockNumber),
			BlockHash:   blockHash,
			Epoch:       big.NewInt(epoch),
			Difficulty:  currentBlock.Difficulty(),
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
		log.Infof("To generate new block vss, current head: %d", worker.chain.CurrentBlock())
		work := worker.commitNewWork(&sigShareMessage)
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
		log.Infof("In tryBlockGenBLS(), end block generation")
		worker.blockGenRunning = false
	}()

	log.Infof("In tryBlockGenBLS(), begin block generation")
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
		log.Infof(
			"I am the proposer vss %s, self: %x, proposer: %x",
			sigShareMessage.Key(),
			worker.pos.Vssid.Bytes(),
			sigShareMessage.ForProposer,
		)
		worker.blockGenBLS(sigShareMessage)
	} else {
		log.Infof(
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
				log.Infof(
					"blockGenBLSLoop: current block: %d, %x",
					worker.chain.CurrentBlock().Number(),
					worker.chain.CurrentBlock().Hash().Bytes()[:8],
				)
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
	defer worker.txSub.Unsubscribe()
	defer worker.chainHeadSub.Unsubscribe()
	defer worker.chainSideSub.Unsubscribe()

	for {
		// A real event arrived, process interesting content
		select {
		// Handle ChainHeadEvent
		case ev := <-worker.chainHeadCh:
			log.Debugf("%v", ev)
		// Handle ChainSideEvent
		case ev := <-worker.chainSideCh:
			log.Debugf("%v", ev)
		// Handle TxPreEvent
		case ev := <-worker.txCh:
			log.Debugf("%v", ev)
		// System stopped
		case <-worker.txSub.Err():
			return
		case <-worker.chainHeadSub.Err():
			return
		case <-worker.chainSideSub.Err():
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
	log.Infof("Before worker.mux.post @@@@@@@@@@@@@@@@@@@@@@@@@@@ block: %d, hash: %x", block.Number(), block.Hash().Bytes()[:4])
	// broadcast the block
	worker.mux.Post(core.NewMinedBlockEvent{Block: block})
	log.Infof("After worker.mux.Post @@@@@@@@@@@@@@@@@@@@@@@@@@@ block: %d, hash: %x", block.Number(), block.Hash().Bytes()[:4])
	var (
		events []interface{}
		logs   = work.state.Logs()
	)
	events = append(events, core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
	if stat == core.CanonStatTy {
		events = append(events, core.ChainHeadEvent{Block: block})
	}

	//worker.chain.PostChainEvents(events, logs)
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
		log.Info("Mining too far in the future", "wait", common.PrettyDuration(wait))
		time.Sleep(wait)
	}

	num := parent.Number()
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     num.Add(num, common.Big1),
		Time:       big.NewInt(tstamp),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Extra:      worker.extra,
		Difficulty: new(big.Int),
	}
	header.Coinbase = worker.pos.Vssid
	if err := worker.pos.Prepare(worker.chain, header); err != nil {
		log.Error("Failed to prepare header for mining", "err", err)
		return nil
	}

	log.Debugf("commitNewWork parent num: %d, parent hash %x", num, header.ParentHash.Bytes())

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
	txsset := []*types.TransactionsByPriceAndNonce{txs}
	work.commitTransactions(worker.mux, txsset, worker.chain, worker.coinbase, worker.mc)

	extraData := make(map[int][]byte)
	var blsSig []byte
	if !worker.pos.ShouldSkipVSS(header.Number.Int64()) {
		var err error
		blsSig, err = worker.pos.GetBlsSig(sigShareMessage.Key())
		if err != nil {
			log.Errorf("Worker FAILED to get BLS sig: %v", err)
			return nil
		} else {
			log.Infof(
				"^^^^^^^^^^^^^^^^^^^^^   Worker get BLS sig = [%x], toSign: %s, err: %v",
				blsSig[:8], sigShareMessage.Key(), err,
			)
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

func (env *Work) commitTransactions(mux *event.TypeMux, txs []*types.TransactionsByPriceAndNonce, bc *core.BlockChain, coinbase common.Address, mc Backend) {
	gp := new(core.GasPool).AddGas(env.header.GasLimit)
	var coalescedLogs []*types.Log
	var queryQueue []*types.QueryContract

	{
		systx := core.CreateSysTx(env.state)

		err, receipt := env.commitTransaction(systx, bc, coinbase, gp, mc)
		if err == nil {
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, receipt.Logs...)
			env.tcount++
		} else {
			log.Debugf("systx failed ")
		}
	}

	length := len(txs)
	for i := 0; i < length; i++ {
		for {
			// Retrieve the next transaction and abort if all done
			tx := txs[i].Peek()
			if tx == nil {
				break
			}

			// Error may be ignored here. The error has already been checked
			// during transaction acceptance is the transaction pool.
			//
			// We use the eip155 signer regardless of the current hf.
			// MOAC, use EIP155 signer
			from, _ := types.Sender(env.signer, tx)

			// Start executing the transaction
			env.state.Prepare(tx.Hash(), common.Hash{}, env.tcount)

			err, receipt := env.commitTransaction(tx, bc, coinbase, gp, mc)
			switch err {
			case core.ErrGasLimitReached:
				// Pop the current out-of-gas transaction without shifting in the next from the account
				log.Trace("GasRemaining limit exceeded for current block", "sender", from)
				txs[i].Pop()

			case core.ErrNonceTooLow:
				// New head notification data race between the transaction pool and miner, shift
				log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
				txs[i].Shift()

			case core.ErrNonceTooHigh:
				// Reorg notification data race between the transaction pool and miner, skip account =
				log.Trace("Skipping account with hight nonce", "sender", from, "nonce", tx.Nonce())
				txs[i].Pop()

			//fix 1
			case vm.ErrEmptyCode:
				// Reorg notification data race between the transaction pool and miner, skip account =
				log.Trace("Skipping account with empty code", "sender", from, "nonce", tx.Nonce())
				txs[i].Pop()

			case vm.ErrInvalidCode:
				// Reorg notification data race between the transaction pool and miner, skip account =
				log.Trace("Skipping account with invalid code, check compiler version", "sender", from, "nonce", tx.Nonce())
				txs[i].Pop()

			case nil:
				if receipt.QueryInBlock > 0 {
					//add to queue for system contract call
					queryQueue = append(queryQueue, &types.QueryContract{Block: receipt.QueryInBlock, ContractAddress: receipt.ContractAddress})
				}
				// Everything ok, collect the logs and shift in the next transaction from the same account
				coalescedLogs = append(coalescedLogs, receipt.Logs...)
				env.tcount++
				txs[i].Shift()

			default:
				// Strange error, discard the transaction and get the next in line (note, the
				// nonce-too-high clause will prevent us from executing in vain).
				log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
				txs[i].Shift()
			}
		}
	}

	if len(coalescedLogs) > 0 || env.tcount > 0 {
		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make([]*types.Log, len(coalescedLogs))
		for i, l := range coalescedLogs {
			cpy[i] = new(types.Log)
			*cpy[i] = *l
		}
		go func(logs []*types.Log, tcount int) {
			if len(logs) > 0 {
				mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				mux.Post(core.PendingStateEvent{})
			}
		}(cpy, env.tcount)
	}
}

func (env *Work) commitTransaction(tx *types.Transaction, bc *core.BlockChain, coinbase common.Address, gp *core.GasPool, mc Backend) (error, *types.Receipt) {
	snap := env.state.Snapshot()

	//here NetworkRelay becomes NetworkRelayInterface item.
	receipt, _, err := core.ApplyTransaction(
		env.config,
		bc,
		&coinbase,
		gp,
		env.state,
		env.header,
		tx,
		env.header.GasUsed,
		vm.Config{},
		mc.TxPool(),
	)
	if err != nil {
		env.state.RevertToSnapshot(snap)
		return err, nil
	}
	env.txs = append(env.txs, tx)
	env.receipts = append(env.receipts, receipt)

	return nil, receipt
}

/*
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

		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		//
		// We use the eip155 signer regardless of the current hf.
		// MOAC, use EIP155 signer
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
*/
