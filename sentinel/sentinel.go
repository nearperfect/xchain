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
	"github.com/MOACChain/MoacLib/rlp"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/accounts/keystore"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/dkg"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/vnode/config"
	"github.com/MOACChain/xchain/xdefi/vaultx"
	"github.com/MOACChain/xchain/xdefi/vaulty"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

const (
	VaultCheckInterval                   = 3
	VssCheckInterval                     = 3
	PersistSeenVaultEventWithSigChanSize = 8192
	VaultEventsBatchChanSize             = 0
	ScanStep                             = uint64(20)
	MaxBlockNumber                       = uint64(1000000000000)
	Gwei                                 = int64(1000000000)
	BlockDelay                           = 12
	MaxEmptyBatchBlocks                  = 200
	MintBatchSize                        = 5
	MintIntervalBlocks                   = 50
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

func (sentinel *Sentinel) prepareClients(
	xChainId uint64,
	xChainFuncPrefix string,
	xChainRPC string,
	yChainId uint64,
	yChainFuncPrefix string,
	yChainRPC string,
) (*mcclient.Client, *mcclient.Client, *mcclient.Client) {
	var clientX *mcclient.Client
	var err error
	if xChainId != 0 {
		clientX, err = mcclient.Dial(xChainRPC)
		if err != nil {
			log.Errorf("Unable to connect to network:%v\n", err)
			return nil, nil, nil
		}
		clientX.SetFuncPrefix(xChainFuncPrefix)
		// sanity check chain id
		chainId_, err := clientX.ChainID(context.Background())
		if err != nil {
			log.Errorf("------------client chain id err: %v -----------------", err)
			return nil, nil, nil
		}
		if xChainId != chainId_.Uint64() {
			log.Errorf(
				"Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
				chainId_, xChainId,
			)
			return nil, nil, nil
		}
	}
	var clientY *mcclient.Client
	if yChainId != 0 {
		clientY, err = mcclient.Dial(yChainRPC)
		if err != nil {
			log.Errorf("Unable to connect to network:%v\n", err)
			return nil, nil, nil
		}
		clientY.SetFuncPrefix(yChainFuncPrefix)
		// sanity check chain id
		chainId_, err := clientY.ChainID(context.Background())
		if err != nil {
			log.Errorf("------------client chain id err: %v -----------------", err)
			return nil, nil, nil
		}
		if yChainId != chainId_.Uint64() {
			log.Errorf(
				"Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
				chainId_, yChainId,
			)
			return nil, nil, nil
		}
	}
	clientXevents, err := mcclient.Dial(sentinel.rpc)
	if err != nil {
		log.Errorf("------------client x err: %v %s-----------------", err, sentinel.rpc)
		return nil, nil, nil
	}
	return clientX, clientY, clientXevents
}

func (sentinel *Sentinel) prepareContracts(
	vaultXAddr common.Address,
	vaultYAddr common.Address,
	needXevents bool,
	clientX *mcclient.Client,
	clientY *mcclient.Client,
	clientXevents *mcclient.Client,
) (*vaultx.VaultX, *vaulty.VaultY, *xevents.XEvents) {
	var vaultxContract *vaultx.VaultX
	var err error
	if (vaultXAddr != common.Address{}) {
		vaultxContract, err = vaultx.NewVaultX(
			vaultXAddr,
			clientX,
		)
		if err != nil {
			log.Errorf("------------client new vault X contract err: %v -----------------", err)
			return nil, nil, nil
		}
	} else {
		vaultxContract = nil
	}
	var vaultyContract *vaulty.VaultY
	if (vaultYAddr != common.Address{}) {
		vaultyContract, err = vaulty.NewVaultY(
			vaultYAddr,
			clientY,
		)
		if err != nil {
			log.Errorf("------------client new vault Y contract err: %v -----------------", err)
			return nil, nil, nil
		}
	} else {
		vaultyContract = nil
	}
	var xeventsContract *xevents.XEvents
	if needXevents {
		xeventsContract, err = xevents.NewXEvents(
			XeventsAddr,
			clientXevents,
		)
		if err != nil {
			log.Errorf("------------client new xevents contract err: %v -----------------", err)
			return nil, nil, nil
		}
	}
	return vaultxContract, vaultyContract, xeventsContract
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

func (sentinel *Sentinel) myTurn(number uint64) bool {
	// use node index to determine if this node should update xevents this round
	// batchNumber's value should be consistent between validators
	nodeIndex, nodeCount := sentinel.dkg.FindAddressInNodelist(sentinel.dkg.Vssid)
	myTurn := false
	if nodeIndex == -1 {
		log.Errorf(
			"----------- My turn err: index: %d, count: %d, last block: %d-----------------",
			nodeIndex,
			nodeCount,
			number,
		)
		return false
	} else {
		myTurn = number%uint64(nodeCount) == uint64(nodeIndex)
	}

	return myTurn
}

func (sentinel *Sentinel) scanDeposit(
	client *mcclient.Client,
	vaultx *vaultx.VaultX,
	xevents *xevents.XEvents,
	notMyTurn bool,
	lastBlock uint64,
	vaultContract common.Address,
	storeCounter uint64,
) *VaultEventsBatch {
	eventCount := uint64(0)
	defer sentinel.scanStatus(lastBlock, &lastBlock, &eventCount)

	startBlock := lastBlock
	// the outer for loop is to skip void blocks with no events
	for {
		// throttling for sending query
		time.Sleep(800 * time.Millisecond)
		// once the filter return some events, we're done for this round.
		if eventCount > 0 {
			return nil
		}
		currentBlock, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Errorf(
				"----------------- sentinel: unable to get current block number from chain: %v ---------------", err)
		}
		if lastBlock > currentBlock {
			return nil
		}

		// filter events from vaultx, [start, end] are inclusive
		endBlock := lastBlock + ScanStep - 1
		if endBlock > currentBlock-BlockDelay {
			endBlock = currentBlock - BlockDelay
		}
		if endBlock < lastBlock {
			endBlock = lastBlock + 1
		}

		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &endBlock,
		}
		log.Errorf("--------------- filter opts:[batch=%d], start: %d, end: %d-----------", startBlock, lastBlock, endBlock)
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
				event.Tip,
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
		log.Errorf("---------------- raw filter result: batch size = %d--------------", len(batch))
		// once we receive the batch with events,
		// scan for this round is done, just return
		if sentinel.batchNumber != 0 && len(batch) > 0 {
			log.Errorf(
				"----------- New batch mined (commit = %t), number: %d, start: %d, end: %d, events: %d -------------",
				!notMyTurn,
				sentinel.batchNumber,
				startBlock,
				endBlock,
				len(batch),
			)
			return &VaultEventsBatch{
				batch,
				vaultContract,
				startBlock,
				endBlock,
				storeCounter,
			}
		}

		// corner case where we scan many blocks but getting no events
		if endBlock-startBlock >= MaxEmptyBatchBlocks && len(batch) == 0 {
			return &VaultEventsBatch{
				batch,
				vaultContract,
				startBlock,
				endBlock,
				storeCounter,
			}
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

func (sentinel *Sentinel) prepareWatchDeposit(
	xChainId uint64,
	xChainFuncPrefix string,
	xChainRPC string,
	xVaultAddr common.Address,
) (*mcclient.Client, *mcclient.Client, *vaultx.VaultX,
	*xevents.XEvents, bool, uint64, uint64) {
	clientX, _, clientXevents := sentinel.prepareClients(
		xChainId,
		xChainFuncPrefix,
		xChainRPC,
		0,
		"",
		"",
	)

	needXevents := true
	vaultxContract, _, xeventsContract := sentinel.prepareContracts(
		xVaultAddr,
		common.Address{},
		needXevents,
		clientX,
		nil,
		clientXevents,
	)

	if clientX == nil || clientXevents == nil || vaultxContract == nil || xeventsContract == nil {
		log.Errorf(
			"---------watch deposit: one of clients/contracts is nil: %v, # %v, # %v, # %v ------------",
			clientX, clientXevents, vaultxContract, xeventsContract,
		)
		return nil, nil, nil, nil, false, 0, 0
	}

	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xeventsContract.VaultWatermark(
		callOpts, xVaultAddr,
	)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return nil, nil, nil, nil, false, 0, 0
	}

	batchNumber := uint64(0)
	storeCounter := uint64(0)
	lastBlock := uint64(0)
	// fallback to use createAt
	if vaultWatermark.Uint64() == 0 {
		createdAt, err := vaultxContract.CreatedAt(callOpts)
		if err != nil {
			log.Errorf(
				"------------client CreateAt err: %v, set lastBlock to %d-----------------",
				err, lastBlock,
			)
			return nil, nil, nil, nil, false, 0, 0
		} else {
			// first time
			lastBlock = createdAt.Uint64()
			storeCounter = uint64(0)
		}
	} else {
		// normal case
		// we've scanned watermark block, proceed to next one
		batchNumber = vaultWatermark.Uint64()
		lastBlock = vaultWatermark.Uint64() + 1
		counter, err := xeventsContract.VaultStoreCounter(
			callOpts, xVaultAddr, big.NewInt(int64(batchNumber)),
		)
		if err != nil {
			log.Errorf("---------------client vault store counter error %v------------", err)
			return nil, nil, nil, nil, false, 0, 0
		}
		storeCounter = counter.Uint64()
	}

	notMyTurn := !sentinel.myTurn(lastBlock)
	if notMyTurn {
		log.Errorf("------------Not my turn, scan only %d --------------", lastBlock)
	} else {
		log.Errorf("-------- My turn, scan & commit batch----------")
	}

	// share with other go routine
	sentinel.batchNumber = lastBlock

	return clientX, clientXevents, vaultxContract,
		xeventsContract, notMyTurn, lastBlock, storeCounter
}

func (sentinel *Sentinel) watchDeposit(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
) {
	defer log.Errorf("*********************END WATCH DEPOSIT***********************")
	for {
		// sleep for interval
		time.Sleep(VaultCheckInterval * time.Second)
		// # 0
		clientX, clientXevents, vaultx, xevents, notMyTurn, lastBlock, storeCounter :=
			sentinel.prepareWatchDeposit(
				chainId,
				chainFuncPrefix,
				chainRPC,
				vaultContract,
			)

		// # 1
		batch := sentinel.scanDeposit(
			clientX,
			vaultx,
			xevents,
			notMyTurn,
			lastBlock,
			vaultContract,
			storeCounter,
		)

		if batch == nil {
			continue
		}

		if !notMyTurn {
			// # 2
			omitted, committed, errors, transactor := sentinel.commitDepositBatch(
				clientXevents,
				xevents,
				batch,
			)

			if transactor == nil {
				continue
			}

			// # 3
			wait := 0
			for {
				time.Sleep(3 * time.Second)
				callOpts := &bind.CallOpts{}
				newStoreCounter, err := xevents.StoreCounter(callOpts)
				if err == nil {
					log.Errorf(
						"------- before vault watermark commit, store counter %d, prev %d, batch size: %d, omitted: %d, commit %d, errors: %d -----------",
						newStoreCounter,
						batch.StoreCounter,
						len(batch.Batch),
						omitted,
						committed,
						errors,
					)

					// wait for the new store counter to increase and then update vault water mark
					oldBatchPass := newStoreCounter.Uint64() == batch.StoreCounter && uint64(len(batch.Batch)) == omitted
					newBatchPass := newStoreCounter.Uint64() == batch.StoreCounter+uint64(len(batch.Batch))
					if oldBatchPass || newBatchPass {
						log.Errorf(
							"-------------------MOVE vault watermark HIGHER to %d---------------",
							batch.EndBlockNumber,
						)
						xevents.UpdateVaultWatermark(
							transactor,
							batch.Vault,
							big.NewInt(int64(batch.EndBlockNumber)),
						)
						break
					}
				}
				wait += 1
				if wait > 20 {
					log.Errorf("--------------wait time out for storeCounter------------")
					break
				}
			}
		}
	}
}

type EventData struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}

func (sentinel *Sentinel) scanAndCallMint(
	clientXevents *mcclient.Client,
	clientY *mcclient.Client,
	xevents *xevents.XEvents,
	vaultY *vaulty.VaultY,
	xVaultAddr common.Address,
	tokenMapping TokenMapping,
) {
	callOpts := &bind.CallOpts{}
	waterMark, err := vaultY.TokenMappingWatermark(
		callOpts,
		common.HexToAddress(tokenMapping.SourceToken),
		common.HexToAddress(tokenMapping.MappedToken),
	)
	if err != nil {
		log.Errorf("-------------- mint watermark err: %v--------------", err)
		return
	}
	log.Errorf("-----------Mint watermark: %d, %s, %s------------------------------",
		waterMark,
		tokenMapping.SourceToken,
		tokenMapping.MappedToken,
	)
	xchainBlockNumber, err := clientXevents.BlockNumber(context.Background())
	if err != nil {
		log.Errorf("---------------Can not get xchain block number, NO MINT ----------------")
	}

	// each node will mint for n blocks, then hand it over to next node
	if !sentinel.myTurn(xchainBlockNumber / MintIntervalBlocks) {
		log.Errorf("---------Not My Turn to Mint: %d ------------------",
			sentinel.batchNumber,
		)
		return
	} else {
		log.Errorf("---------My Turn to Mint !!!!: %d ------------------",
			sentinel.batchNumber,
		)
	}

	// setup transactor
	gasPrice := int64(5) // 5gwei
	gasLimit := int64(300000)
	transactor, nonceAt := sentinel.getTransactor(clientY, gasPrice, gasLimit)
	if transactor == nil {
		return
	}
	transactor.Nonce = big.NewInt(int64(nonceAt))

	// check balance
	balance, err := clientY.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return
	}
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		log.Errorf(
			"---------------No balance for account on Vault Y chain: %x--------------------",
			sentinel.key.Address,
		)
		return
	}

	committed := uint64(0)
	errors := uint64(0)
	for i := int64(0); i < int64(MintBatchSize); i++ {
		var vaultEvent core.VaultEvent
		vaultEventData, err := xevents.VaultEvents(
			callOpts, xVaultAddr,
			tokenMapping.Sha256(),
			big.NewInt(waterMark.Int64()+i),
		)
		if err != nil {
			log.Errorf("------------ retrieve vault events err: %v---------------", err)
		}
		rlp.DecodeBytes(vaultEventData.EventData, &vaultEvent)
		log.Errorf(
			"-----------TO MINT: nonce: %d amount: %d, tip: %d, to: %x, mappedToken: %x-------------",
			big.NewInt(waterMark.Int64()+i),
			vaultEvent.Amount,
			vaultEvent.Tip,
			vaultEvent.To,
			vaultEvent.MappedToken,
		)
		// if no more vault event
		if vaultEvent.Amount == nil || vaultEvent.Amount.Int64() == 0 {
			break
		} else {
			tx, err := vaultY.Mint(
				transactor,
				vaultEvent.SourceToken,
				vaultEvent.MappedToken,
				vaultEvent.To,
				vaultEvent.Amount,
				vaultEvent.Tip,
				vaultEvent.Nonce,
			)
			transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

			if err != nil {
				errors += 1
				log.Errorf("---------------sentinel vault Y mint tx err: %v ------------------", err)
			} else {
				log.Errorf(
					"---------sentinel vault Y mint() tx: %x, from: %x, nonce: %d ---------",
					tx.Hash(), sentinel.key.Address.Bytes(), tx.Nonce(),
				)
				committed += 1
			}
		}
	}

	waterMarkBefore := waterMark.Int64()
	// wait for watermark to update
	for {
		if committed == 0 {
			log.Errorf("--------------No vault Y mint tx called-------------------")
			break
		}
		waterMark, _ := vaultY.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		log.Errorf(
			"---------Vault Y watermark: %d, before: %d, committed: %d ---------",
			waterMark, waterMarkBefore, committed,
		)
		time.Sleep(10 * time.Second)
	}
}

func (sentinel *Sentinel) getTransactor(
	client *mcclient.Client,
	gasPrice int64,
	gasLimit int64,
) (*bind.TransactOpts, uint64) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, 0
	}

	// build transactor
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		chainId,
	)
	transactor.GasPrice = big.NewInt(gasPrice * Gwei)
	transactor.GasLimit = uint64(gasLimit)

	nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return nil, 0
	}

	return transactor, nonceAt
}

func (sentinel *Sentinel) doMint(
	yChainId uint64,
	yChainFuncPrefix string,
	yChainRPC string,
	xVaultAddr common.Address,
	yVaultAddr common.Address,
	tokenMapping TokenMapping,
) {
	defer log.Errorf("*********************END CALL MINT***********************")
	for {
		_, clientY, clientXevents := sentinel.prepareClients(
			0,
			"",
			"",
			yChainId,
			yChainFuncPrefix,
			yChainRPC,
		)

		needXevents := true
		_, vaultyContract, xeventsContract := sentinel.prepareContracts(
			common.Address{},
			yVaultAddr,
			needXevents,
			nil,
			clientY,
			clientXevents,
		)

		if clientY == nil || clientXevents == nil || vaultyContract == nil || xeventsContract == nil {
			log.Errorf(
				"---------do mint: one of clients/contracts is nil: %v, # %v, # %v, # %v ----------------",
				clientY, clientXevents, vaultyContract, xeventsContract,
			)
			return
		}

		sentinel.scanAndCallMint(
			clientXevents,
			clientY,
			xeventsContract,
			vaultyContract,
			xVaultAddr,
			tokenMapping,
		)
		time.Sleep(VaultCheckInterval * time.Second)
	}
}

func (sentinel *Sentinel) commitDepositBatch(
	client *mcclient.Client,
	xevents *xevents.XEvents,
	batch *VaultEventsBatch,
) (uint64, uint64, uint64, *bind.TransactOpts) {
	log.Errorf(
		"-----------commit batch BEGIN from %d to %d ]-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)
	defer log.Errorf(
		"-----------commit batch END from %d to %d-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)

	// sanity check chain id
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client chain id err: %v -----------------", err)
		return 0, 0, 0, nil
	}

	balance, err := client.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, 0, 0, nil
	}
	nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, 0, 0, nil
	}
	log.Errorf(
		"$$$$$$$$  %x, balance: %d, nonce: %d, chainId: %d",
		sentinel.key.Address.Bytes(), balance, nonceAt, chainId,
	)
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		return 0, 0, 0, nil
	}
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xevents.VaultWatermark(
		callOpts, XeventsAddr,
	)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return 0, 0, 0, nil
	}
	// prevent re-entry
	if vaultWatermark.Uint64() >= batch.StartBlockNumber {
		log.Errorf(
			"---------------skip batch commit: vault: %d, batch number: %d ----------",
			vaultWatermark.Uint64(),
			batch.StartBlockNumber,
		)
		return 0, 0, 0, nil
	}

	// build transactor
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		chainId,
	)
	transactor.GasPrice = big.NewInt(5 * Gwei)
	transactor.GasLimit = uint64(300000)
	transactor.Nonce = big.NewInt(int64(nonceAt))
	log.Debugf("%v, %v", xevents, transactor)

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
			"$$$$$$$$$$ @@@@@@@ %x, %s, %x, %d ",
			vaultEvent.Vault,
			vaultEvent.TokenMapping(),
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
	}
	sort.Ints(nonces)

	var vaultEventHash common.Hash
	log.Errorf("-------------- nonces: %v -------------", nonces)

	committed := uint64(0)
	errors := uint64(0)
	omitted := uint64(0)
	// commit all events
	for index, nonce := range nonces {
		vaultEventHash = noncesMap[nonce]
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event

		callOpts := &bind.CallOpts{}
		watermark, err := xevents.VaultEventWatermark(
			callOpts, vaultEvent.Vault, vaultEvent.TokenMappingSha256(),
		)
		if err != nil {
			log.Errorf("-----------------@@@@@@@@@@@ err: %v", err)
		} else {
			log.Errorf(
				"$$$$$$$  tokenmapping watermark %d, this tokenmappin nonce: %d",
				watermark, vaultEvent.Nonce,
			)
		}

		// if first nonce is different, omit the whole batch
		if watermark.Int64() != vaultEvent.Nonce.Int64() && index == 0 {
			omitted = uint64(len(batch.Batch))
			log.Errorf(
				"----------------tokenmappin nonce low, omit: %v-----------------------------",
				nonces,
			)
			break
		}

		log.Errorf(
			"-------------- Store tx: vault: %x, tokenmapping: %x, nonce: %d------------",
			vaultEvent.Vault.Bytes()[:8],
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
		tx, err := xevents.Store(
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

	return omitted, committed, errors, transactor
}

func (sentinel *Sentinel) watchBurn(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
	queue chan *core.VaultEvent,
) {
	defer log.Errorf("*********************END WATCH BURN***********************")
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
	defer log.Errorf("*********************END CALL WITHDRAW***********************")
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
		vaulty := vaultPairConfig.VaultY
		go sentinel.watchDeposit(
			vaultx.ChainId,
			vaultx.ChainFuncPrefix,
			vaultx.ChainRPC,
			common.HexToAddress(vaultx.VaultAddress),
		)
		// mint
		for _, tokenMapping := range vaultPairConfig.TokenMappings {
			go sentinel.doMint(
				vaulty.ChainId,
				vaulty.ChainFuncPrefix,
				vaulty.ChainRPC,
				common.HexToAddress(vaultx.VaultAddress),
				common.HexToAddress(vaulty.VaultAddress),
				tokenMapping,
			)
		}

		/*


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
