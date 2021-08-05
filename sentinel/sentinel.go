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
	"encoding/json"
	"fmt"
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
	"github.com/MOACChain/xchain/xdefi/vaulty"
	"github.com/MOACChain/xchain/xdefi/xconfig"
)

type Sentinel struct {
	db                    mcdb.Database
	dkg                   *dkg.DKG
	key                   *keystore.Key
	scope                 event.SubscriptionScope
	chainHeadSub          event.Subscription
	config                *config.Configuration
	Rpc                   string
	vaultsConfig          *VaultPairListConfig
	batchNumber           uint64
	batchEndNumber        uint64
	vaultEventWithSigFeed event.Feed
	workerStatusMu        sync.RWMutex

	// config settings
	configVersion     uint64
	xconfig           *xconfig.XConfig
	configVersionFeed event.Feed

	// vault events
	VaultEventsSigs     map[common.Hash]*set.Set
	VaultEvents         map[common.Hash]*VaultEventWithSig
	VaultEventProcessMu sync.RWMutex

	// Rpc stats
	RpcCallPerClient map[string]map[string]int
	RpcCallStatMu    sync.RWMutex

	// chan for vault events
	depositChan map[common.Address]chan *vaultx.VaultXTokenDeposit
	burnChan    map[common.Address]chan *vaulty.VaultYTokenBurn

	// worker status
	workerStatus map[string]int
}

type VaultEventsBatch struct {
	Batch map[common.Hash]bool
	Vault common.Address
	Start uint64
	End   uint64
}

/////////////////////////////////////////////////////////////
///////////////       Sentinel        ///////////////////////
/////////////////////////////////////////////////////////////

func New(
	bc *core.BlockChain,
	vaultsConfig *VaultPairListConfig,
	db mcdb.Database,
	dkg *dkg.DKG,
	rpc string,
	key *keystore.Key,
) *Sentinel {
	sentinel := &Sentinel{
		db:               db,
		VaultEvents:      make(map[common.Hash]*VaultEventWithSig),
		VaultEventsSigs:  make(map[common.Hash]*set.Set),
		config:           bc.VnodeConfig(),
		vaultsConfig:     vaultsConfig,
		dkg:              dkg,
		key:              key,
		Rpc:              "http://" + rpc,
		RpcCallPerClient: make(map[string]map[string]int),
		depositChan:      make(map[common.Address]chan *vaultx.VaultXTokenDeposit),
		burnChan:         make(map[common.Address]chan *vaulty.VaultYTokenBurn),
		workerStatus:     make(map[string]int),
		configVersion:    uint64(0),
	}

	err := sentinel.InitXconfig()
	if err != nil {
		panic("sentinel init failed")
	}

	sentinel.scope.Open()
	log.Infof("sentinel start with config: %v", sentinel.vaultsConfig)

	go sentinel.startWatchersAndForwarders()

	// make sure startWatchersAndForwarders subscribe to config feed
	// before ReloadConfigLoop send out the first feed
	time.Sleep(time.Second)

	// periodically check for new configs
	go sentinel.ReloadConfigLoop()

	// periodically check for worker goroutines
	go sentinel.workerStatusLoop()

	return sentinel
}

func (sentinel *Sentinel) InitXconfig() error {
	xconfigClient, err := mcclient.Dial(sentinel.Rpc)
	if err != nil {
		return fmt.Errorf("can not init xconfig client, err: %v", err)
	}

	sentinel.xconfig, err = xconfig.NewXConfig(XConfigAddr, xconfigClient)
	if err != nil {
		return fmt.Errorf("can not init xconfig client, err: %v", err)
	}
	return nil
}

func (sentinel *Sentinel) ReloadConfigLoop() {
	log.Infof("------------enter reload config loop--------------")
	defer log.Errorf("-------exit reload config loop ---------------")
	ticker := time.NewTicker(ConfigCheckInterval * time.Second)
	for {
		select {
		case <-ticker.C:
			callOpts := &bind.CallOpts{}

			// get config version
			configVersion, err := sentinel.xconfig.VaultConfigVersion(callOpts)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}

			// get config data
			var newConfig VaultPairListConfig
			newConfigBytes, err := sentinel.xconfig.VaultConfigs(callOpts, configVersion)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}
			log.Infof("New Xconfig [%d]: %s", configVersion.Uint64(), string(newConfigBytes))
			err = json.Unmarshal(newConfigBytes, &newConfig)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}

			// update config with sentinel if it is new
			if configVersion.Uint64() > sentinel.configVersion {
				sentinel.vaultsConfig = &newConfig
				sentinel.configVersion = configVersion.Uint64()
			}

			// send out the config version regardless of its value
			// so that new feed subscriber will receive it
			sentinel.configVersionFeed.Send(configVersion.Uint64())
		}
	}
}

func (sentinel *Sentinel) Threshold() int {
	if sentinel.dkg.IsVSSReady() {
		return sentinel.dkg.Bls.Threshold
	} else {
		return 0
	}
}

func (sentinel *Sentinel) SubscribeVaultEventWithSig(ch chan<- VaultEventWithSig) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventWithSigFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsSigs() {
	log.Debugf("Vault events received:")
	for hash, received := range sentinel.VaultEventsSigs {
		log.Debugf("\t%x, %d", hash.Bytes()[:8], received.Size())
	}
}

func (sentinel *Sentinel) LogRpcStat(client, method string) {
	sentinel.RpcCallStatMu.Lock()
	defer sentinel.RpcCallStatMu.Unlock()
	if clientStat, found := sentinel.RpcCallPerClient[client]; found {
		clientStat[method] += 1
	} else {
		clientStat := make(map[string]int)
		clientStat[method] += 1
		sentinel.RpcCallPerClient[client] = clientStat
	}
}

func (sentinel *Sentinel) PrintRpcStat() {
	sentinel.RpcCallStatMu.Lock()
	defer sentinel.RpcCallStatMu.Unlock()
	for client, clientStat := range sentinel.RpcCallPerClient {
		allCount := 0
		for method, count := range clientStat {
			log.Infof("[RPC] \t%s => [%s]:\t %d", client, method, count)
			allCount += count
		}
		log.Infof("[RPC] \t%s => ALL:\t\t %d", client, allCount)
	}
}

func (sentinel *Sentinel) getTransactor(
	client *mcclient.Client,
	gasPrice int64,
	gasLimit int64,
) *bind.TransactOpts {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil
	}
	sentinel.LogRpcStat("vault", "ChainID")

	// build transactor
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		chainId,
	)
	transactor.GasPrice = big.NewInt(gasPrice)
	transactor.GasLimit = uint64(gasLimit)

	//nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	//sentinel.LogRpcStat("vault", "NonceAt")
	//if err != nil {
	//	return nil
	//}
	//transactor.Nonce = big.NewInt(int64(nonceAt))
	return transactor
}

func (sentinel *Sentinel) ProcessVaultEventWithSig(
	vaultEventWithSig *VaultEventWithSig,
) common.Hash {
	// both mc/handler and this process call this function
	sentinel.VaultEventProcessMu.Lock()
	defer sentinel.VaultEventProcessMu.Unlock()

	vaultEvent := vaultEventWithSig.Event
	vaultEventHash := vaultEvent.Hash()
	// keep count of the event
	if received, found := sentinel.VaultEventsSigs[vaultEventHash]; found {
		received.Add(string(vaultEventWithSig.Blssig))
	} else {
		sentinel.VaultEventsSigs[vaultEventHash] = set.New()
		sentinel.VaultEventsSigs[vaultEventHash].Add(fmt.Sprintf("%x", vaultEventWithSig.Blssig))
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
			"[myTurn]\t My turn err: index: %d, count: %d, last block: %d",
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

func (sentinel *Sentinel) prepareCommonContext(
	xChainId uint64,
	xChainFuncPrefix string,
	xChainRPC string,
	xChainWS string,
	xVaultAddr common.Address,
	yChainId uint64,
	yChainFuncPrefix string,
	yChainRPC string,
	yChainWS string,
	yVaultAddr common.Address,
	configVersion uint64,
	eventType uint64,
) *XdefiContext {
	xdefiContextConfig := &XdefiContextConfig{
		xChainId,
		xChainFuncPrefix,
		xChainRPC,
		xChainWS,
		yChainId,
		yChainFuncPrefix,
		yChainRPC,
		yChainWS,
		xVaultAddr,
		yVaultAddr,
		configVersion,
	}
	xdefiContext := NewXdefiContext(eventType, xdefiContextConfig, sentinel, configVersion)
	log.Infof("Xdefi context config: %s", xdefiContext.Config)

	// clients are contracts are set inline with context
	xdefiContext.InitClients()
	xdefiContext.InitContracts()

	if xdefiContext.IsReady() {
		return xdefiContext
	} else {
		return nil
	}
}

func (sentinel *Sentinel) addWorker(worker string) {
	sentinel.workerStatusMu.Lock()
	defer sentinel.workerStatusMu.Unlock()
	sentinel.workerStatus[worker] += 1
}

func (sentinel *Sentinel) reduceWorker(worker string) {
	sentinel.workerStatusMu.Lock()
	defer sentinel.workerStatusMu.Unlock()
	sentinel.workerStatus[worker] -= 1
}

func (sentinel *Sentinel) workerStatusLoop() {
	ticker := time.NewTicker(ConfigCheckInterval * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Infof("[Worker status] ----------------")
			for worker, count := range sentinel.workerStatus {
				log.Infof("[Worker status] %s: %d", worker, count)
			}
		}
	}
}

func (sentinel *Sentinel) startWatchersAndForwarders() {
	defer log.Errorf("*********************end watcher and fowarder loop***********************")

	currentConfigVersion := uint64(0)
	newConfigChan := make(chan uint64, 10)
	sub := sentinel.configVersionFeed.Subscribe(newConfigChan)
	defer sub.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case <-sub.Err():
			break
		case newConfigVersion := <-newConfigChan:
			// exhaust the channel
			for i := 0; i < len(newConfigChan); i++ {
				newConfigVersion = <-newConfigChan
			}
			log.Infof("[startWatchersAndForwarders] new config version received: %d -> %d",
				sentinel.configVersion, newConfigVersion,
			)

			// no need to create go routine if incoming config is lower
			if sentinel.configVersion <= currentConfigVersion {
				continue
			}
			currentConfigVersion = sentinel.configVersion

			// sanity check
			if newConfigVersion != sentinel.configVersion {
				log.Errorf(
					"mismatch between chan config version: %d and sentinel config version: %d",
					newConfigVersion, sentinel.configVersion,
				)
				continue
			}

			// create all go routines for watching vault contracts on various blockchains
			for _, vaultPairConfig := range sentinel.vaultsConfig.Vaults {
				vaultx := vaultPairConfig.VaultX
				vaulty := vaultPairConfig.VaultY

				log.Infof("new config [%d]: %v, %v", sentinel.configVersion, vaultx, vaulty)
				// deposit & mint
				xdefiContextXY := sentinel.prepareCommonContext(
					vaultx.ChainId,
					vaultx.ChainFuncPrefix,
					vaultx.ChainRPC,
					vaultx.ChainWS,
					common.HexToAddress(vaultx.VaultAddress),
					vaulty.ChainId,
					vaulty.ChainFuncPrefix,
					vaulty.ChainRPC,
					vaulty.ChainWS,
					common.HexToAddress(vaulty.VaultAddress),
					sentinel.configVersion,
					DEPOSIT,
				)
				// 3 steps to record vault events in xchain
				// 1) scan vault in source chain
				// 2) queue: wait for enough sigs
				// 3) pending: commit to xchain
				go xdefiContextXY.ScanVaultEvents(sentinel)
				go xdefiContextXY.QueueVaultEventsBatch(sentinel)
				go xdefiContextXY.PendingVaultEventsBatch(sentinel)

				// forward vault events from xchain to target chain
				for _, tokenMapping := range vaultPairConfig.TokenMappings {
					go xdefiContextXY.ForwardVaultEvents(sentinel, tokenMapping)
				}

				// burn & withdarw
				xdefiContextYX := sentinel.prepareCommonContext(
					vaultx.ChainId,
					vaultx.ChainFuncPrefix,
					vaultx.ChainRPC,
					vaultx.ChainWS,
					common.HexToAddress(vaultx.VaultAddress),
					vaulty.ChainId,
					vaulty.ChainFuncPrefix,
					vaulty.ChainRPC,
					vaulty.ChainWS,
					common.HexToAddress(vaulty.VaultAddress),
					sentinel.configVersion,
					BURN,
				)

				// 3steps to record vault events in xchain
				// 1) scan vault in source chain
				// 2) queue: wait for enough sigs
				// 3) pending: commit to xchain
				go xdefiContextYX.ScanVaultEvents(sentinel)
				go xdefiContextYX.QueueVaultEventsBatch(sentinel)
				go xdefiContextYX.PendingVaultEventsBatch(sentinel)

				// forward vault events from xchain to target chain
				for _, tokenMapping := range vaultPairConfig.TokenMappings {
					go xdefiContextYX.ForwardVaultEvents(sentinel, tokenMapping)
				}

				//////////////////////////////////////////////////////////////////////
				//////////////////////////////////////////////////////////////////////
				// sanity check
				if !(xdefiContextXY.EventType == DEPOSIT && xdefiContextYX.EventType == BURN) {
					panic("Missing Deposit or Burn go routine")
				}

				// if ws is enabled
				//go sentinel.WatchVaultXLogs(xdefiContextXY)
				//go sentinel.WatchVaultYLogs(xdefiContextYX)
			}
		}
	}
}
