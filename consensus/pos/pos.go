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
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
)

// Pos is a consensus engine
type Pos struct {
	// vss
	Vssid                  common.Address
	VssBaseContractAddress common.Address
	Bls                    *BLS
	prevbls                *list.List
	Vss                    *VSS
	AllSigs                *gocache.Cache // blockhash -> sig sha
	VssKey                 *VSSKey
	vssIsrunningMutex      sync.Mutex
	vssEnabled             bool
	nodesPubkey            map[common.Address][]byte
	Position               int
	NodeList               []common.Address
	vssSeenConfigs         map[uint64]bool
	vssSettings            map[int]*BLS // config version ==> bls obj
	currentConfigVersion   int          // current vss config version
	uploadedConfigTime     time.Time    // time when the last config was uploaded
	AddressToIndex         map[string]int
	IndexToAddress         map[int]string

	// vss stats
	ReceivedBlocks  map[string]bool
	muSubchainStats sync.RWMutex

	// vss mc client
	client       *mcclient.Client
	callOpts     *bind.CallOpts
	transactOpts *bind.TransactOpts
	vssbase      *xdefi.VssBase
}

// New creates a full sized pos PoW scheme.
func New() *Pos {
	client, _ := mcclient.Dial("http://172.21.0.11:8545")
	vssbase, _ := xdefi.NewVssBase(
		common.HexToAddress("0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0"),
		client,
	)
	pos := &Pos{
		vssEnabled: true,
		client:     client,
		callOpts:   &bind.CallOpts{},
		vssbase:    vssbase,
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
