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
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/mcdb"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/dkg"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
	"github.com/MOACChain/xchain/vnode/config"
)

const (
	VaultCheckInterval = 7
	VssCheckInterval   = 5
)

type Sentinel struct {
	vaultEventFeed event.Feed
	scope          event.SubscriptionScope
	chainHeadCh    chan core.ChainHeadEvent
	chainHeadSub   event.Subscription
	config         *config.Configuration
	vaultsConfig   *VaultPairListConfig
	db             mcdb.Database
	dkg            *dkg.DKG

	// vault events
	VaultEventsReceived  map[common.Hash]int
	VaultEvents          map[common.Hash]*core.VaultEvent
	VaultEventsNonce     map[TokenMapping]map[*big.Int]*core.VaultEvent
	VaultEventsWatermark map[TokenMapping]*big.Int

	// channel between go routines for the same token mapping pair
	VaultEventsChanXY map[int](chan *core.VaultEvent)
	VaultEventsChanYX map[int](chan *core.VaultEvent)
}

type TokenMapping struct {
	SourceAddress common.Address
	SourceChainid *big.Int
	MappedAddress common.Address
	MappedChainid *big.Int
}

func New(bc *core.BlockChain, vaultsConfig *VaultPairListConfig, db mcdb.Database, dkg *dkg.DKG) *Sentinel {
	sentinel := &Sentinel{
		db:                   db,
		chainHeadCh:          make(chan core.ChainHeadEvent, 10),
		VaultEvents:          make(map[common.Hash]*core.VaultEvent),
		VaultEventsReceived:  make(map[common.Hash]int),
		VaultEventsNonce:     make(map[TokenMapping]map[*big.Int]*core.VaultEvent),
		VaultEventsWatermark: make(map[TokenMapping]*big.Int),
		config:               bc.VnodeConfig(),
		vaultsConfig:         vaultsConfig,
		VaultEventsChanXY:    make(map[int](chan *core.VaultEvent)),
		VaultEventsChanYX:    make(map[int](chan *core.VaultEvent)),
		dkg:                  dkg,
	}
	sentinel.scope.Open()
	sentinel.chainHeadSub = bc.SubscribeChainHeadEvent(sentinel.chainHeadCh)
	log.Infof("sentinel start with config: %v", sentinel.vaultsConfig)
	go sentinel.watchNewBlock()
	//go sentinel.start()

	return sentinel
}

func (sentinel *Sentinel) SubscribeVaultEvent(ch chan<- core.VaultEvent) event.Subscription {
	if sentinel == nil {
		panic(0)
	}

	return sentinel.scope.Track(sentinel.vaultEventFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	for hash, received := range sentinel.VaultEventsReceived {
		log.Infof("Vault events received:")
		log.Infof("\t%x, %d", hash.Bytes(), received)
	}
}

func (sentinel *Sentinel) MarkVaultEvent(vaultContract common.Address, vaultEvent *core.VaultEvent) {
	// keep count of the event
	sentinel.VaultEventsReceived[vaultEvent.Hash()] += 1
	sentinel.VaultEvents[vaultEvent.Hash()] = vaultEvent

	// update VaultEventsNonce
	nonce := vaultEvent.Nonce
	tokenMapping := TokenMapping{
		vaultEvent.SourceToken,
		vaultEvent.SourceChainid,
		vaultEvent.MappedToken,
		vaultEvent.MappedChainid,
	}
	if nonceMapping, found := sentinel.VaultEventsNonce[tokenMapping]; found {
		nonceMapping[nonce] = vaultEvent
	} else {
		sentinel.VaultEventsNonce[tokenMapping] = make(map[*big.Int]*core.VaultEvent)
		sentinel.VaultEventsNonce[tokenMapping][nonce] = vaultEvent
	}

	// update water mark
	sentinel.VaultEventsWatermark[tokenMapping] = nonce
}

func (sentinel *Sentinel) watchDeposit(
	chainId uint64,
	chainFuncPrefix string,
	chainRPC string,
	vaultContract common.Address,
	queue chan *core.VaultEvent,
) {
	lastBlock := uint64(0)
	for {
		time.Sleep(VaultCheckInterval * time.Second)
		log.Infof(
			"This is sentinel for chain[%d], vault X [0x%x]: [[      ***  %d  ***      ]]",
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
			log.Errorf("Chain ID does not match, have: %d, want: %d", chainId_, chainId)
			return
		}

		currentBlock, _ := client.BlockNumber(context.Background())
		vaultx, _ := xdefi.NewVaultX(
			vaultContract,
			client,
		)
		// get events from vaultx
		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &currentBlock,
		}
		log.Debugf("sentinel: start %d, end: %d", lastBlock, currentBlock)
		itr, err := vaultx.FilterTokenDeposit(
			filterOpts,
			[]common.Address{},
			[]common.Address{},
			[]*big.Int{},
		)
		if err != nil {
			log.Errorf("Unable to get token deposit iterator: %v", err)
			continue
		}
		for itr.Next() {
			event := itr.Event
			vaultEvent := core.VaultEvent{
				event.SourceChainid,
				event.SourceToken,
				event.MappedChainid,
				event.MappedToken,
				event.From,
				event.Amount,
				event.DepositNonce,
				[]byte{},
			}

			// record state in sentinel, include this node itself.
			sentinel.MarkVaultEvent(vaultContract, &vaultEvent)

			// broad cast to other nodes
			sentinel.vaultEventFeed.Send(vaultEvent)
		}
		lastBlock = currentBlock
	}
}

func (sentinel *Sentinel) callMint(queue chan *core.VaultEvent) {
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
			log.Errorf("Chain ID does not match, have: %d, want: %d", chainId_, chainId)
			return
		}

		currentBlock, _ := client.BlockNumber(context.Background())
		log.Debugf("sentinel: start %d, end: %d", lastBlock, currentBlock)
	}
}

func (sentinel *Sentinel) callWithdraw(queue chan *core.VaultEvent) {
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
	}
	log.Infof("In sentinel start(): vss is ready, t=%d", sentinel.dkg.Bls.Threshold)
	// create all go routines for watching vault contracts on various blockchains
	for pairIndex, vaultPairConfig := range sentinel.vaultsConfig.Vaults {
		pairId := vaultPairConfig.Id()
		GetLastBlock(sentinel.db, pairId)
		WriteLastBlock(sentinel.db, pairId, 0)

		sentinel.VaultEventsChanXY[pairIndex] = make(chan *core.VaultEvent)
		// watch vaultx
		vaultx := vaultPairConfig.VaultX
		go sentinel.watchDeposit(
			vaultx.ChainId,
			vaultx.ChainFuncPrefix,
			vaultx.ChainRPC,
			common.HexToAddress(vaultx.VaultAddress),
			sentinel.VaultEventsChanXY[pairIndex],
		)
		go sentinel.callMint(
			sentinel.VaultEventsChanXY[pairIndex],
		)

		sentinel.VaultEventsChanYX[pairIndex] = make(chan *core.VaultEvent)
		// watch vaulty
		vaulty := vaultPairConfig.VaultY
		go sentinel.watchBurn(
			vaulty.ChainId,
			vaulty.ChainFuncPrefix,
			vaulty.ChainRPC,
			common.HexToAddress(vaulty.VaultAddress),
			sentinel.VaultEventsChanYX[pairIndex],
		)
		go sentinel.callWithdraw(
			sentinel.VaultEventsChanYX[pairIndex],
		)
	}
}
