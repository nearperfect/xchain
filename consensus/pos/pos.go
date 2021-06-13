// Copyright 2017 The MOAC-core Authors
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

// Package pos implements the pos proof-of-work consensus engine.
package pos

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
	"github.com/MOACChain/xchain/vnode/config"
)

// Pos is a consensus engine
type Pos struct {
	// vss
	Vssid                common.Address
	Bls                  *BLS
	prevbls              *list.List
	Vss                  *VSS
	AllSigs              *gocache.Cache // blockhash -> sig sha
	VssKey               *VSSKey
	vssIsrunningMutex    sync.Mutex
	vssEnabled           bool
	nodesPubkey          map[common.Address][]byte
	Position             int
	NodeList             []common.Address
	vssSeenConfigs       map[uint64]bool
	vssSettings          map[int]*BLS // config version ==> bls obj
	currentConfigVersion int          // current vss config version
	uploadedConfigTime   time.Time    // time when the last config was uploaded
	AddressToIndex       map[string]int
	IndexToAddress       map[int]string

	// vss stats
	ReceivedBlocks  map[string]bool
	muSubchainStats sync.RWMutex

	// vss mc client
	client       *mcclient.Client
	callOpts     *bind.CallOpts
	transactOpts *bind.TransactOpts
	vssbase      *xdefi.VssBase
	vnodeconfig  *config.Configuration
}

// New creates a full sized pos PoW scheme.
func New(cfg *config.Configuration) *Pos {
	url := fmt.Sprintf("http://%s:%s", cfg.VnodeIP, cfg.VnodePort)
	client, _ := mcclient.Dial(url)
	vssbase, _ := xdefi.NewVssBase(
		common.HexToAddress(cfg.VssBaseAddr),
		client,
	)
	pos := &Pos{
		vssEnabled:  true,
		client:      client,
		callOpts:    &bind.CallOpts{},
		vssbase:     vssbase,
		vnodeconfig: cfg,
	}

	// vss config loop
	if pos.vssEnabled {
		log.Debugf("vss enabled, ready to run vss loop")
		// init vsskey
		pos.LoadVSSKey()
		go pos.VssStateLoop()        // for updating config
		go pos.VssSlashingLoop()     // for checking slash
		go pos.VssUploadConfigLoop() // for uploading config
		go pos.NewVnodeBlockLoop()   // for checking new block in vnode
	}

	return pos
}
