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
	ScanStep                             = uint64(5)
	MaxBlockNumber                       = uint64(1000000000000)
)

var (
	XeventsAddr = common.HexToAddress("0x0000000000000000000000000000000000010000")
)

type Sentinel struct {
	vaultEventWithSigFeed event.Feed
	scope                 event.SubscriptionScope
	chainHeadCh           chan core.ChainHeadEvent
	chainHeadSub          event.Subscription
	config                *config.Configuration
	rpc                   string
	vaultsConfig          *VaultPairListConfig
	db                    mcdb.Database
	dkg                   *dkg.DKG
	key                   *keystore.Key

	// vault events
	VaultEventsReceived map[common.Hash]*set.Set
	VaultEvents         map[common.Hash]*core.VaultEvent
	//[batchNumber][event hash] => event count
	VaultEventsBatch    map[uint64]map[common.Hash]int
	VaultEventProcessMu sync.RWMutex

	// channel between go routines for the same token mapping pair
	VaultEventsChanXY map[int](chan *core.VaultEvent)
	VaultEventsChanYX map[int](chan *core.VaultEvent)

	// persist
	PersistSeenVaultEventWithSigChan chan PersistSeenVaultEventWithSig
}

type PersistSeenVaultEventWithSig struct {
	batchNumber       uint64
	vaultEventWithSig *core.VaultEventWithSig
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
		db:                  db,
		chainHeadCh:         make(chan core.ChainHeadEvent, 10),
		VaultEvents:         make(map[common.Hash]*core.VaultEvent),
		VaultEventsReceived: make(map[common.Hash]*set.Set),
		VaultEventsBatch:    make(map[uint64]map[common.Hash]int),
		config:              bc.VnodeConfig(),
		vaultsConfig:        vaultsConfig,
		VaultEventsChanXY:   make(map[int](chan *core.VaultEvent)),
		VaultEventsChanYX:   make(map[int](chan *core.VaultEvent)),
		dkg:                 dkg,
		key:                 key,
		rpc:                 "http://" + rpc,
		PersistSeenVaultEventWithSigChan: make(
			chan PersistSeenVaultEventWithSig, PersistSeenVaultEventWithSigChanSize),
	}
	sentinel.scope.Open()
	sentinel.chainHeadSub = bc.SubscribeChainHeadEvent(sentinel.chainHeadCh)
	log.Infof("sentinel start with config: %v", sentinel.vaultsConfig)
	go sentinel.watchNewBlock()
	go sentinel.start()
	go sentinel.xeventsLoop()

	return sentinel
}

func (sentinel *Sentinel) SubscribeVaultEventWithSig(ch chan<- core.VaultEventWithSig) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventWithSigFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	log.Infof("Vault events received:")
	for hash, received := range sentinel.VaultEventsReceived {
		log.Infof("\t%x, %d", hash.Bytes()[:8], received.Size())
	}
}

func (sentinel *Sentinel) xeventsLoop() {
	defer log.Infof("Loop exit: xevents loop")
	client, err := mcclient.Dial(sentinel.rpc)
	if err != nil {
		log.Errorf("Unable to connect to network:%v, rpc: %v\n", err, sentinel.rpc)
		return
	}
	client.SetFuncPrefix("mc")
	xeventsContract, err := xevents.NewXEvents(
		XeventsAddr,
		client,
	)
	log.Infof("%v", xeventsContract)
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		big.NewInt(0),
	)

	for {
		select {
		case persistSeenVaultEventWithSig := <-sentinel.PersistSeenVaultEventWithSigChan:
			vaultEventWithSig := persistSeenVaultEventWithSig.vaultEventWithSig
			log.Infof("%v", vaultEventWithSig)
			vaultEvent := vaultEventWithSig.Event
			_, err := xeventsContract.Store(
				transactor,
				vaultEventWithSig.Blssig,
				vaultEvent.Vault,
				vaultEvent.Nonce,
				vaultEvent.TokenMappingSha256(),
				vaultEvent.BlockNumber,
				vaultEvent.Bytes(),
			)
			if err != nil {
				log.Errorf("----------------------sentinel xevents store err: %v ------------------", err)
			}
		}
	}
}

func (sentinel *Sentinel) IsVaultEventsBatchDone(batchNumber uint64) bool {
	return false
	//sentinel.VaultEventsBatch[batchNumber]
}

func (sentinel *Sentinel) ProcessVaultEventWithSig(
	vaultEventWithSig *core.VaultEventWithSig,
	batchNumber uint64,
) {
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
	// update vault event batch
	count := sentinel.VaultEventsReceived[vaultEventHash].Size()
	if batchNumber != 0 {
		if batchmap, found := sentinel.VaultEventsBatch[batchNumber]; found {
			batchmap[vaultEventHash] = count
		} else {
			sentinel.VaultEventsBatch[batchNumber] = make(map[common.Hash]int)
			sentinel.VaultEventsBatch[batchNumber][vaultEventHash] = count
		}
	}
	// keep the event
	sentinel.VaultEvents[vaultEventHash] = &vaultEvent

	// persist the event in xevent contract if more than
	// threshold number of nodes have seen it
	if sentinel.dkg.Bls != nil {
		log.Infof(
			"@@@@@@@@@@@@@@@@@@@@@  size: %d, threshold: %d",
			sentinel.VaultEventsReceived[vaultEventHash].Size(),
			sentinel.dkg.Bls.Threshold,
		)
		if sentinel.VaultEventsReceived[vaultEventHash].Size() >= sentinel.dkg.Bls.Threshold {
			persistSeenVaultEventWithSig := PersistSeenVaultEventWithSig{
				batchNumber,
				vaultEventWithSig,
			}
			sentinel.PersistSeenVaultEventWithSigChan <- persistSeenVaultEventWithSig
		}
	}
}

func (sentinel *Sentinel) shouldIUpdate(lastBlock uint64) bool {
	// use node index to determine if this node should update xevents this round
	nodeIndex, nodeCount := sentinel.dkg.FindAddressInNodelist(sentinel.dkg.Vssid)
	shouldIUpdate := false
	if nodeIndex == -1 {
		log.Errorf(
			"------------shouldIUpdate err: index: %d, count: %d, last block: %d-----------------",
			nodeIndex,
			nodeCount,
			lastBlock,
		)
		return false
	} else {
		shouldIUpdate = lastBlock%uint64(nodeCount) == uint64(nodeIndex)
	}

	return shouldIUpdate
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
			lastBlock = createdAt.Uint64()
		}
	} else {
		// we've scanned watermark block, proceed to next one
		// if store() does not store all vault events
		// (e.g. partially complete for all events in one block),
		// the next call to it will fail because of nonce check.
		lastBlock = vaultWatermark.Uint64() + 1
	}

	defer sentinel.scanStatus(lastBlock, &lastBlock, &eventCount)

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
			sentinel.ProcessVaultEventWithSig(&vaultEventWithSig, lastBlock)
			// broad cast to other nodes
			sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
		}
		lastBlock += ScanStep
		log.Infof(
			"@@@@@@@@@@@@@@@@@@@@@@s vault: %x,  watermark block: %d, current: %d, last: %d, end: %d, count: %d",
			vaultContract.Bytes()[:8], vaultWatermark, currentBlock, lastBlock, endBlock, eventCount,
		)
	}
}

func (sentinel *Sentinel) scanStatus(begin uint64, end *uint64, eventCount *uint64) {
	log.Infof(
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
	defer log.Infof("*********************END WATCH DEPOSIT***********************")
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
	defer log.Infof("*********************END CALL MINT***********************")
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
	defer log.Infof("*********************END WATCH BURN***********************")
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
	defer log.Infof("*********************END CALL WITHDRAW***********************")
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

func (sentinel *Sentinel) watchNewBlock() {
	for {
		select {
		case event := <-sentinel.chainHeadCh:
			log.Infof(
				"[-----------sentinel watch chain head: %d, vss ready: %t -------------]",
				event.Block.Number(), sentinel.dkg.IsVSSReady(),
			)
		}
	}
}

func (sentinel *Sentinel) start() {
	for {
		time.Sleep(VssCheckInterval * time.Second)
		if sentinel.dkg.IsVSSReady() {
			break
		}
		log.Infof("In sentinel start(): wait for vss ready")
	}
	log.Infof("In sentinel start(): vss is ready, t=%d", sentinel.dkg.Bls.Threshold)
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
