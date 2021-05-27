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
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
)

type Sentinel struct {
	vaultEventFeed event.Feed
	scope          event.SubscriptionScope
	chainHeadCh    chan core.ChainHeadEvent
	chainHeadSub   event.Subscription
}

func New(bc *core.BlockChain) *Sentinel {
	sentinel := &Sentinel{
		chainHeadCh: make(chan core.ChainHeadEvent, 10),
	}
	//sentinel.chainHeadSub = bc.SubscribeChainHeadEvent(sentinel.chainHeadCh)
	go sentinel.start()

	return sentinel
}

func (sentinel *Sentinel) SubscribeVaultEvent(ch chan<- core.VaultEvent) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventFeed.Subscribe(ch))
}

func (sentinel *Sentinel) start() {
	lastBlock := uint64(0)
	for {
		time.Sleep(10 * time.Second)
		log.Infof("This is sentinel: [[      ***  %d  ***      ]]", lastBlock)

		// init rpc client
		client, err := mcclient.Dial("http://172.21.0.11:8545")
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
		itr, err := vaultx.FilterTokenDeposit(
			filterOpts,
			[]common.Address{},
			[]common.Address{},
			[]*big.Int{},
		)
		for itr.Next() {
			event := itr.Event
			log.Infof("vault event 1 %x, %x, %x, %s, %d",
				event.SourceToken.Bytes(),
				event.MappedToken.Bytes(),
				event.From.Bytes(),
				event.Amount,
				event.DepositNonce.Uint64(),
			)

			vaultEvent := core.VaultEvent{
				event.SourceToken,
				event.MappedToken,
				event.From,
				event.Amount.Uint64(),
				event.DepositNonce.Uint64(),
				[]byte{},
			}
			log.Infof("vault event 2 %v", vaultEvent)

			go sentinel.vaultEventFeed.Send(vaultEvent)
		}

		lastBlock = currentBlock
	}
}
