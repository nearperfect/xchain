// Copyright 2014 The MOAC-core Authors
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

package sentinel

import (
	"context"
	"math/big"
	"sort"
	"sync"
	"time"

	"gopkg.in/fatih/set.v0"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/mcdb"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/accounts/keystore"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/dkg"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/vnode/config"
	"github.com/MOACChain/xchain/xdefi/vaultx"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

const (
	VaultCheckInterval                   = 7
	VssCheckInterval                     = 5
	PersistSeenVaultEventWithSigChanSize = 8192
	VaultEventsBatchChanSize             = 128
	ScanStep                             = uint64(10)
	MaxBlockNumber                       = uint64(1000000000000)
	Gwei                                 = int64(1000000000)
)

var (
	XeventsAddr = common.HexToAddress("0x0000000000000000000000000000000000010000")
)

type Sentinel struct {
	vaultEventWithSigFeed event.Feed
	scope                 event.SubscriptionScope
	chainHeadSub          event.Subscription
	config                *config.Configuration
	rpc                   string
	vaultsConfig          *VaultPairListConfig
	db                    mcdb.Database
	dkg                   *dkg.DKG
	key                   *keystore.Key
	batchNumber           uint64
	batchEndNumber        uint64

	// vault events
	VaultEventsReceived map[common.Hash]*set.Set
	VaultEvents         map[common.Hash]*core.VaultEventWithSig
	VaultEventsBatchMap map[uint64]VaultEventsBatch
	VaultEventProcessMu sync.RWMutex

	// channel between go routines for the same token mapping pair
	VaultEventsChanXY map[int](chan *core.VaultEvent)
	VaultEventsChanYX map[int](chan *core.VaultEvent)

	// persist
	PersistSeenVaultEventWithSigChan chan PersistSeenVaultEventWithSig
	VaultEventsBatchChan             chan VaultEventsBatch
}

type PersistSeenVaultEventWithSig struct {
	batchNumber       uint64
	vaultEventWithSig *core.VaultEventWithSig
}

type VaultEventsBatch struct {
	Batch            map[common.Hash]bool
	Vault            common.Address
	StartBlockNumber uint64
	EndBlockNumber   uint64
	StoreCounter     uint64
}

func New(
	bc *core.BlockChain,
	vaultsConfig *VaultPairListConfig,
	db mcdb.Database,
	dkg *dkg.DKG,
	rpc string,
	key *keystore.Key,
) *Sentinel {
	sentinel := &Sentinel{
		db:                   db,
		VaultEvents:          make(map[common.Hash]*core.VaultEventWithSig),
		VaultEventsReceived:  make(map[common.Hash]*set.Set),
		VaultEventsBatchMap:  make(map[uint64]VaultEventsBatch),
		config:               bc.VnodeConfig(),
		vaultsConfig:         vaultsConfig,
		VaultEventsChanXY:    make(map[int](chan *core.VaultEvent)),
		VaultEventsChanYX:    make(map[int](chan *core.VaultEvent)),
		dkg:                  dkg,
		key:                  key,
		rpc:                  "http://" + rpc,
		VaultEventsBatchChan: make(chan VaultEventsBatch, VaultEventsBatchChanSize),
		PersistSeenVaultEventWithSigChan: make(
			chan PersistSeenVaultEventWithSig, PersistSeenVaultEventWithSigChanSize),
	}
	sentinel.scope.Open()
	log.Debugf("sentinel start with config: %v", sentinel.vaultsConfig)
	go sentinel.start()
	go sentinel.xeventsLoop()

	return sentinel
}

func (sentinel *Sentinel) SubscribeVaultEventWithSig(ch chan<- core.VaultEventWithSig) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventWithSigFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	log.Debugf("Vault events received:")
	for hash, received := range sentinel.VaultEventsReceived {
		log.Debugf("\t%x, %d", hash.Bytes()[:8], received.Size())
	}
}

func (sentinel *Sentinel) xeventsLoop() {
	defer log.Debugf("Loop exit: xevents loop")

	for {
		select {
		case batch := <-sentinel.VaultEventsBatchChan:
			log.Errorf("-------------receive new batch -------------")
			if sentinel.dkg.Bls == nil {
				log.Debugf("------------xeventsLoop: bls not ready-------------")
				continue
			}

			threshold := sentinel.dkg.Bls.Threshold
			for {
				time.Sleep(5 * time.Second)
				if sentinel.IsVaultEventsBatchAllReceived(batch, threshold) {
					break
				}
				log.Errorf("--------------Not all received-----------------------")
			}
			// commit the batch
			sentinel.commitBatch(batch)
			sentinel.PrintVaultEventsReceived()
		}
	}
}

func (sentinel *Sentinel) commitBatch(batch VaultEventsBatch) (uint64, error) {
	log.Errorf(
		"-----------commit batch BEGIN from %d to %d ]-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)
	defer log.Debugf(
		"-----------commit batch END from %d to %d-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)
	client, err := mcclient.Dial(sentinel.rpc)
	if err != nil {
		log.Errorf("Unable to connect to network:%v, rpc: %v\n", err, sentinel.rpc)
		return 0, nil
	}
	client.SetFuncPrefix("mc")
	xeventsContract, err := xevents.NewXEvents(
		XeventsAddr,
		client,
	)
	// sanity check chain id
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client chain id err: %v -----------------", err)
		return 0, nil
	}
	log.Debugf("111111111111111 connect xevents with rpc: %s, chain id: %d", sentinel.rpc, chainId)
	balance, err := client.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, nil
	}
	nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, nil
	}
	log.Errorf("$$$$$$$$  %x, balance: %d, nonce: %d", sentinel.key.Address.Bytes(), balance, nonceAt)
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		return 0, nil
	}
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xeventsContract.VaultWatermark(
		callOpts, XeventsAddr,
	)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return 0, nil
	}
	// prevent re-entry
	if vaultWatermark.Uint64() >= batch.StartBlockNumber {
		log.Errorf(
			"---------------skip batch commit: vault: %d, batch number: %d ----------",
			vaultWatermark.Uint64(),
			batch.StartBlockNumber,
		)
		return 0, nil
	}

	// build transactor
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		chainId,
	)
	transactor.GasPrice = big.NewInt(5 * Gwei)
	transactor.GasLimit = uint64(300000)
	transactor.Nonce = big.NewInt(int64(nonceAt))
	log.Debugf("%v, %v", xeventsContract, transactor)

	if err != nil {
		log.Errorf("Unable to connect to network:%v, rpc: %v\n", err, sentinel.rpc)
		return 0, nil
	}

	// order the commit sequence by nonce
	nonce := 0
	nonces := make([]int, 0)
	noncesMap := make(map[int]common.Hash)
	for vaultEventHash, _ := range batch.Batch {
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event
		nonce = int(vaultEvent.Nonce.Int64())
		nonces = append(nonces, int(vaultEvent.Nonce.Int64()))
		noncesMap[nonce] = vaultEventHash
		log.Errorf(
			"$$$$$$$$$$ %x, %x, %d ",
			vaultEvent.Vault,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
	}
	sort.Ints(nonces)

	var vaultEventHash common.Hash
	omitted := uint64(0)
	committed := uint64(0)
	total := uint64(0)
	errors := uint64(0)
	for _, nonce := range nonces {
		total += 1
		vaultEventHash = noncesMap[nonce]
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event

		callOpts := &bind.CallOpts{}
		watermark, err := xeventsContract.VaultEventWatermark(
			callOpts, vaultEvent.Vault, vaultEvent.TokenMappingSha256(),
		)
		if err != nil {
			log.Errorf("-----------------@@@@@@@@@@@ err: %v", err)
		} else {
			log.Errorf(
				"$$$$$$$$  tokenmapping watermark %d, this tokenmappin nonce: %d",
				watermark, vaultEvent.Nonce,
			)
		}
		if watermark.Int64() > vaultEvent.Nonce.Int64() {
			omitted += 1
			continue
		}

		log.Debugf(
			"-------------- Store tx: vault: %x, tokenmapping: %x, nonce: %d------------",
			vaultEvent.Vault.Bytes()[:8],
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
		tx, err := xeventsContract.Store(
			transactor,
			vaultEventWithSig.Blssig,
			vaultEvent.Vault,
			vaultEvent.Nonce,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.BlockNumber,
			vaultEvent.Bytes(),
		)
		transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

		if err != nil {
			errors += 1
			log.Errorf("---------------sentinel xevents store err: %v ------------------", err)
		} else {
			log.Debugf(
				"---------sentinel xevents store tx: %x, from: %x, nonce: %d ---------",
				tx.Hash(), sentinel.key.Address.Bytes(), tx.Nonce(),
			)
			committed += 1
		}
	}

	wait := 0
	for {
		callOpts := &bind.CallOpts{}
		storeCounter, err := xeventsContract.StoreCounter(callOpts)
		if err != nil {
			log.Errorf("-------------- batch commit, store counter err: %v -----------", err)
		} else {
			log.Errorf(
				"---batch commit, store counter %d, prev %d, total %d, commit %d, omit: %d, errors: %d ---",
				storeCounter,
				batch.StoreCounter,
				total,
				committed,
				omitted,
				errors,
			)
		}
		// if storeCounter gets updated, we are done with this batch
		if storeCounter.Uint64() == batch.StoreCounter+committed {
			xeventsContract.UpdateVaultWatermark(
				transactor,
				batch.Vault,
				big.NewInt(int64(batch.EndBlockNumber)),
			)
			break
		} else {
			time.Sleep(5 * time.Second)
		}

		// break without update vault watermark, waited more than 120 seconds
		if wait > 24 {
			break
		}
		wait += 1
	}

	return committed, nil
}

func (sentinel *Sentinel) IsVaultEventsBatchAllReceived(batch VaultEventsBatch, threshold int) bool {
	for vaultEventHash, _ := range batch.Batch {
		if sentinel.VaultEventsReceived[vaultEventHash].Size() < threshold {
			return false
		}
	}
	return true
}

func (sentinel *Sentinel) ProcessVaultEventWithSig(
	vaultEventWithSig *core.VaultEventWithSig,
) common.Hash {
	// both mc/handler and this process call this function
	sentinel.VaultEventProcessMu.Lock()
	defer sentinel.VaultEventProcessMu.Unlock()

	vaultEvent := vaultEventWithSig.Event
	vaultEventHash := vaultEvent.Hash()
	// keep count of the event
	if received, found := sentinel.VaultEventsReceived[vaultEventHash]; found {
		received.Add(string(vaultEventWithSig.Blssig))
	} else {
		sentinel.VaultEventsReceived[vaultEventHash] = set.New()
		sentinel.VaultEventsReceived[vaultEventHash].Add(string(vaultEventWithSig.Blssig))
	}
	// keep the event
	sentinel.VaultEvents[vaultEventHash] = vaultEventWithSig

	return vaultEventHash
}

func (sentinel *Sentinel) shouldIStore(batchNumber uint64) bool {
	// use node index to determine if this node should update xevents this round
	// batchNumber's value should be consistent between validators
	nodeIndex, nodeCount := sentinel.dkg.FindAddressInNodelist(sentinel.dkg.Vssid)
	shouldIStore := false
	if nodeIndex == -1 {
		log.Errorf(
			"------------shouldIStore err: index: %d, count: %d, last block: %d-----------------",
			nodeIndex,
			nodeCount,
			batchNumber,
		)
		return false
	} else {
		shouldIStore = batchNumber%uint64(nodeCount) == uint64(nodeIndex)
	}

	return shouldIStore
}

func (sentinel *Sentinel) scanOneRound(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
	queue chan *core.VaultEvent,
) {
	eventCount := uint64(0)
	lastBlock := uint64(0)
	// prepare the client and vaultx instance
	client, err := mcclient.Dial(chainRPC)
	if err != nil {
		log.Errorf("Unable to connect to network:%v\n", err)
		return
	}
	client.SetFuncPrefix(chainFuncPrefix)
	// sanity check chain id
	chainId_, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client chain id err: %v -----------------", err)
		return
	}
	if chainId != chainId_.Uint64() {
		log.Errorf("Chain ID does not match, have: %d, want: %d, check vaultx.json and restart", chainId_, chainId)
		return
	}
	vaultx, err := vaultx.NewVaultX(
		vaultContract,
		client,
	)
	if err != nil {
		log.Errorf("------------client new vault contract err: %v -----------------", err)
		return
	}
	clientx, err := mcclient.Dial(sentinel.rpc)
	if err != nil {
		log.Errorf("------------client x err: %v %s-----------------", err, sentinel.rpc)
	}
	xeventsContract, err := xevents.NewXEvents(
		XeventsAddr,
		clientx,
	)
	if err != nil {
		log.Errorf("------------client new xevents contract err: %v -----------------", err)
		return
	}
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xeventsContract.VaultWatermark(
		callOpts, vaultContract,
	)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return
	}
	// fallback to use createAt
	if vaultWatermark.Uint64() == 0 {
		createdAt, err := vaultx.CreatedAt(callOpts)
		if err != nil {
			log.Errorf(
				"------------client CreateAt err: %v, set lastBlock to %d-----------------",
				err, lastBlock,
			)
			return
		} else {
			// first time
			lastBlock = createdAt.Uint64()
		}
	} else {
		// normal case
		// we've scanned watermark block, proceed to next one
		lastBlock = vaultWatermark.Uint64() + 1
	}

	notMyTurn := !sentinel.shouldIStore(lastBlock)
	if notMyTurn {
		log.Errorf("------------Not my turn, scan only %d --------------", lastBlock)
	} else {
		log.Errorf("-------- My turn, scan & commit batch----------")
	}

	// share with other go routine
	sentinel.batchNumber = lastBlock
	// get store counter
	storeCounter, err := xeventsContract.StoreCounter(callOpts)
	if err != nil {
		log.Errorf("-------------client store counter err: %v ---------------", err)
	}

	defer sentinel.scanStatus(lastBlock, &lastBlock, &eventCount)

	startBlock := lastBlock
	// the outer for loop is to skip void blocks with no events
	for {
		// throttling for sending query
		time.Sleep(1 * time.Second)
		// once the filter return some events, we're done for this round.
		if eventCount > 0 {
			return
		}
		currentBlock, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Errorf(
				"----------------- sentinel: unable to get current block number from chain: %v ---------------", err)
		}
		if lastBlock > currentBlock {
			return
		}

		// filter events from vaultx, [start, end] are inclusive
		endBlock := lastBlock + ScanStep - 1
		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &endBlock,
		}
		itr, err := vaultx.FilterTokenDeposit(
			filterOpts,
			[]common.Address{},
			[]common.Address{},
			[]*big.Int{},
		)
		if err != nil {
			log.Errorf("------------------ sentinel: unable to get token deposit iterator: %v ---------------", err)
			continue
		}

		if !sentinel.dkg.IsVSSReady() {
			log.Errorf("------------------ sentinel: unable to get bls signer: %v ------------------", err)
			continue
		}

		batch := make(map[common.Hash]bool)
		for itr.Next() {
			eventCount += 1
			event := itr.Event
			vaultEvent := core.VaultEvent{
				vaultContract,
				event.SourceChainid,
				event.SourceToken,
				event.MappedChainid,
				event.MappedToken,
				event.From,
				event.Amount,
				event.DepositNonce,
				event.BlockNumber,
			}
			blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
			vaultEventWithSig := core.VaultEventWithSig{
				vaultEvent,
				blssig,
			}

			// process the event in this sentinel, include this node itself.
			vaultEventHash := sentinel.ProcessVaultEventWithSig(&vaultEventWithSig)
			batch[vaultEventHash] = true

			// broad cast to other nodes
			sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
		}

		// once we receive the batch with events,
		// scan for this round is done, just return
		if sentinel.batchNumber != 0 && len(batch) > 0 {
			// if scanned multiple times and end up with the same result
			if batch_, found := sentinel.VaultEventsBatchMap[sentinel.batchNumber]; found {
				if len(batch) == len(batch_.Batch) && endBlock == batch_.EndBlockNumber {
					break
				}
			}
			log.Errorf(
				"----------- New batch mined, number: %d, start: %d, end: %d, events: %d -------------",
				sentinel.batchNumber,
				startBlock,
				endBlock,
				len(batch),
			)
			newBatch := VaultEventsBatch{
				batch,
				vaultContract,
				startBlock,
				endBlock,
				storeCounter.Uint64(),
			}
			sentinel.VaultEventsBatchMap[sentinel.batchNumber] = newBatch
			if !notMyTurn {
				sentinel.VaultEventsBatchChan <- newBatch
			}
			break
		}

		lastBlock += ScanStep
	}
}

func (sentinel *Sentinel) scanStatus(begin uint64, end *uint64, eventCount *uint64) {
	log.Debugf(
		"-------------- Sentinel Scan One Round: start %d, end: %d, event count: %d ----------------",
		begin, *end, *eventCount,
	)
}

func (sentinel *Sentinel) watchDeposit(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
	queue chan *core.VaultEvent,
) {
	defer log.Debugf("*********************END WATCH DEPOSIT***********************")
	for {
		// sleep for interval
		time.Sleep(VaultCheckInterval * time.Second)
		/*

			if lastBlock == 0 {

			}*/

		sentinel.scanOneRound(
			chainId,
			chainFuncPrefix,
			chainRPC,
			vaultContract,
			queue,
		)
	}
}

func (sentinel *Sentinel) callMint(queue chan *core.VaultEvent) {
	defer log.Debugf("*********************END CALL MINT***********************")
	for {
		time.Sleep(VaultCheckInterval * time.Second)
		select {
		case <-queue:
			//process event
		}
	}
}

func (sentinel *Sentinel) watchBurn(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
	queue chan *core.VaultEvent,
) {
	defer log.Debugf("*********************END WATCH BURN***********************")
	lastBlock := uint64(0)
	for {
		time.Sleep(VaultCheckInterval * time.Second)
		log.Debugf(
			"This is sentinel for chain[%d], vault Y [0x%x]: [[      ***  %d  ***      ]]",
			chainId, vaultContract.Bytes()[:8], lastBlock,
		)
		sentinel.PrintVaultEventsReceived()

		client, err := mcclient.Dial(chainRPC)
		if err != nil {
			log.Errorf("Unable to connect to network:%v\n", err)
			continue
		}
		client.SetFuncPrefix(chainFuncPrefix)
		// sanity check chain id
		chainId_, err := client.ChainID(context.Background())
		if chainId != chainId_.Uint64() {
			log.Errorf("Chain ID does not match, have: %d, want: %d, check vaultx.json and restart", chainId_, chainId)
			return
		}

		currentBlock, _ := client.BlockNumber(context.Background())
		log.Debugf("sentinel: start %d, end: %d", lastBlock, currentBlock)
	}
}

func (sentinel *Sentinel) callWithdraw(queue chan *core.VaultEvent) {
	defer log.Debugf("*********************END CALL WITHDRAW***********************")
	for {
		time.Sleep(VaultCheckInterval * time.Second)
		select {
		case <-queue:
			//process event
		}
	}
}

func (sentinel *Sentinel) threshold() int {
	if sentinel.dkg.IsVSSReady() {
		return sentinel.dkg.Bls.Threshold
	} else {
		return 0
	}
}

func (sentinel *Sentinel) start() {
	for {
		time.Sleep(VssCheckInterval * time.Second)
		if sentinel.dkg.IsVSSReady() {
			break
		}
		log.Debugf("In sentinel start(): wait for vss ready")
	}
	log.Debugf("In sentinel start(): vss is ready, t=%d", sentinel.dkg.Bls.Threshold)
	// create all go routines for watching vault contracts on various blockchains
	for pairIndex, vaultPairConfig := range sentinel.vaultsConfig.Vaults {
		sentinel.VaultEventsChanXY[pairIndex] = make(chan *core.VaultEvent)
		sentinel.VaultEventsChanYX[pairIndex] = make(chan *core.VaultEvent)

		// deposit
		vaultx := vaultPairConfig.VaultX
		go sentinel.watchDeposit(
			vaultx.ChainId,
			vaultx.ChainFuncPrefix,
			vaultx.ChainRPC,
			common.HexToAddress(vaultx.VaultAddress),
			sentinel.VaultEventsChanXY[pairIndex],
		)

		/*
				// mint
				go sentinel.callMint(
					sentinel.VaultEventsChanXY[pairIndex],
				)


			// burn
			vaulty := vaultPairConfig.VaultY
			go sentinel.watchBurn(
				vaulty.ChainId,
				vaulty.ChainFuncPrefix,
				vaulty.ChainRPC,
				common.HexToAddress(vaulty.VaultAddress),
				sentinel.VaultEventsChanYX[pairIndex],
			)




				// withdraw
				go sentinel.callWithdraw(
					sentinel.VaultEventsChanYX[pairIndex],
				)*/
	}
}
