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
	"fmt"
	"math/big"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
	"github.com/MOACChain/xchain/vnode/config"
)

type Sentinel struct {
	vaultEventFeed event.Feed
	scope          event.SubscriptionScope
	chainHeadCh    chan core.ChainHeadEvent
	chainHeadSub   event.Subscription
	config         *config.Configuration

	// vault events
	VaultEventsReceived  map[common.Hash]int
	VaultEvents          map[common.Hash]*core.VaultEvent
	VaultEventsNonce     map[TokenMapping]map[*big.Int]*core.VaultEvent
	VaultEventsWatermark map[TokenMapping]*big.Int
}

type TokenMapping struct {
	SourceAddress common.Address
	SourceChainid *big.Int
	MappedAddress common.Address
	MappedChainid *big.Int
}

func New(bc *core.BlockChain) *Sentinel {
	sentinel := &Sentinel{
		chainHeadCh:          make(chan core.ChainHeadEvent, 10),
		VaultEvents:          make(map[common.Hash]*core.VaultEvent),
		VaultEventsReceived:  make(map[common.Hash]int),
		VaultEventsNonce:     make(map[TokenMapping]map[*big.Int]*core.VaultEvent),
		VaultEventsWatermark: make(map[TokenMapping]*big.Int),
		config:               bc.VnodeConfig(),
	}
	sentinel.scope.Open()
	sentinel.chainHeadSub = bc.SubscribeChainHeadEvent(sentinel.chainHeadCh)
	go sentinel.start()

	return sentinel
}

func (sentinel *Sentinel) SubscribeVaultEvent(ch chan<- core.VaultEvent) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	for hash, received := range sentinel.VaultEventsReceived {
		log.Infof("Vault events received:")
		log.Infof("\t%x, %d", hash.Bytes(), received)
	}
}

func (sentinel *Sentinel) MarkVaultEvent(vaultEvent *core.VaultEvent) {
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

func (sentinel *Sentinel) start() {
	lastBlock := uint64(0)
	for {
		time.Sleep(10 * time.Second)
		log.Debugf("This is sentinel: [[      ***  %d  ***      ]]", lastBlock)
		sentinel.PrintVaultEventsReceived()

		// init rpc client
		url := fmt.Sprintf(
			"http://%s:%s",
			sentinel.config.VnodeIP,
			sentinel.config.VnodePort,
		)
		client, err := mcclient.Dial(url)
		if err != nil {
			log.Errorf("Unable to connect to network:%v\n", err)
			continue
		}

		currentBlock, _ := client.BlockNumber(context.Background())
		vaultx, _ := xdefi.NewVaultX(
			common.HexToAddress("0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0"),
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
			sentinel.MarkVaultEvent(&vaultEvent)

			// broad cast to other nodes
			sentinel.vaultEventFeed.Send(vaultEvent)
		}
		lastBlock = currentBlock
	}
}
