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
	"github.com/MOACChain/xchain/xdefi/xconfig"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

const (
	VaultCheckInterval                   = 3
	VssCheckInterval                     = 3
	BatchCommitThreshold                 = 3
	PersistSeenVaultEventWithSigChanSize = 8192
	VaultEventsBatchChanSize             = 0
	ScanRangeBlock                       = uint64(20)
	ScanRangeSingle                      = uint64(2)
	MinScanBlockNumber                   = uint64(20)
	VaultEventScanInterval               = 2
	MaxBlockNumber                       = uint64(1000000000000)
	Gwei                                 = int64(1000000000)
	DefaultGasPrice                      = int64(5) * Gwei
	DefaultGasLimit                      = int64(300000)
	BlockDelay                           = 12
	MaxEmptyBatchBlocks                  = 200
	MintBatchSize                        = 30
	MintIntervalBlocks                   = 50
	MyTurnSeed                           = 10000

	// event type
	DEPOSIT = 10
	BURN    = 11

	// event name
	DEPOSITNAME = "DEPOSIT"
	BURNNAME    = "BURN"
)

var (
	XConfigAddr   = common.HexToAddress("0x0000000000000000000000000000000000010000")
	XeventsXYAddr = common.HexToAddress("0x0000000000000000000000000000000000010001")
	XeventsYXAddr = common.HexToAddress("0x0000000000000000000000000000000000010002")
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

	// config settings
	configVersion     uint64
	xconfig           *xconfig.XConfig
	configVersionFeed event.Feed

	// vault events
	VaultEventsReceived map[common.Hash]*set.Set
	VaultEvents         map[common.Hash]*VaultEventWithSig
	VaultEventProcessMu sync.RWMutex

	// persist
	PersistSeenVaultEventWithSigChan chan PersistSeenVaultEventWithSig
	VaultEventsBatchChan             chan VaultEventsBatch

	// Rpc stats
	RpcCallPerClient map[string]map[string]int
	RpcCallStatMu    sync.RWMutex

	// chan for vault events
	depositChan map[common.Address]chan *vaultx.VaultXTokenDeposit
	burnChan    map[common.Address]chan *vaulty.VaultYTokenBurn
}

type EventData struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}

type PersistSeenVaultEventWithSig struct {
	batchNumber       uint64
	vaultEventWithSig *VaultEventWithSig
}

type VaultEventsBatch struct {
	Batch            map[common.Hash]bool
	Vault            common.Address
	StartBlockNumber uint64
	EndBlockNumber   uint64
}

type XdefiContextConfig struct {
	XChainId         uint64
	XChainFuncPrefix string
	XChainRPC        string
	XChainWS         string
	YChainId         uint64
	YChainFuncPrefix string
	YChainRPC        string
	YChainWS         string
	VaultXAddr       common.Address
	VaultYAddr       common.Address
	ConfigVersion    uint64
}

func (config *XdefiContextConfig) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %s, %d, %s, %s, %s, %x, %x",
		config.XChainId,
		config.XChainFuncPrefix,
		config.XChainRPC,
		config.XChainWS,
		config.YChainId,
		config.YChainFuncPrefix,
		config.YChainRPC,
		config.YChainWS,
		config.VaultXAddr,
		config.VaultYAddr,
	)
}

/////////////////////////////////////////////////////////////
///////////////    XdefiContext       ///////////////////////
/////////////////////////////////////////////////////////////

type XdefiContext struct {
	EventType     uint64
	Config        *XdefiContextConfig
	ConfigVersion uint64
	Sentinel      *Sentinel
	ClientX       *mcclient.Client
	ClientY       *mcclient.Client
	ClientXevents *mcclient.Client
	VaultX        *vaultx.VaultX
	VaultY        *vaulty.VaultY
	XeventsXY     *xevents.XEvents
	XeventsYX     *xevents.XEvents
	ClientXws     *mcclient.Client
	ClientYws     *mcclient.Client
	VaultXws      *vaultx.VaultX
	VaultYws      *vaulty.VaultY
	Xconfig       *xconfig.XConfig
}

func (xdefiContext *XdefiContext) IsDeposit() bool {
	return xdefiContext.EventType == DEPOSIT
}

func (xdefiContext *XdefiContext) Type() string {
	if xdefiContext.EventType == DEPOSIT {
		return DEPOSITNAME
	} else {
		return BURNNAME
	}
}

func logX2Y(err bool, name string, from, to common.Address, format string, args ...interface{}) {
	all := []interface{}{name, from.Bytes()[:3], to.Bytes()[:3]}
	all = append(all, args...)
	if err {
		log.Errorf(" %s %x -> %x "+format, all...)
	} else {
		log.Infof(" %s %x -> %x "+format, all...)
	}
}

func (xdefiContext *XdefiContext) LogFunc(err bool) func(format string, args ...interface{}) {
	var fromAddr common.Address
	var toAddr common.Address
	var name string
	if xdefiContext.IsDeposit() {
		name = fmt.Sprintf("[ver:%d] X -> Y", xdefiContext.ConfigVersion)
		fromAddr = xdefiContext.Config.VaultXAddr
		toAddr = xdefiContext.Config.VaultYAddr

	} else {
		fromAddr = xdefiContext.Config.VaultYAddr
		toAddr = xdefiContext.Config.VaultXAddr
		name = fmt.Sprintf("[ver:%d] Y -> X", xdefiContext.ConfigVersion)
	}

	ret := func(format string, args ...interface{}) {
		logX2Y(err, name, fromAddr, toAddr, format, args...)
	}
	return ret
}

func (xdefiContext *XdefiContext) Xevents() *xevents.XEvents {
	if xdefiContext.IsDeposit() {
		return xdefiContext.XeventsXY
	} else {
		return xdefiContext.XeventsYX
	}
}

func (xdefiContext *XdefiContext) VaultAddrFrom() common.Address {
	if xdefiContext.IsDeposit() {
		return xdefiContext.Config.VaultXAddr
	} else {
		return xdefiContext.Config.VaultYAddr
	}
}

func (xdefiContext *XdefiContext) VaultAddrTo() common.Address {
	if xdefiContext.IsDeposit() {
		return xdefiContext.Config.VaultYAddr
	} else {
		return xdefiContext.Config.VaultXAddr
	}
}

func (xdefiContext *XdefiContext) ClientFrom() *mcclient.Client {
	if xdefiContext.IsDeposit() {
		return xdefiContext.ClientX
	} else {
		return xdefiContext.ClientY
	}
}

func (xdefiContext *XdefiContext) ClientTo() *mcclient.Client {
	if xdefiContext.IsDeposit() {
		return xdefiContext.ClientY
	} else {
		return xdefiContext.ClientX
	}
}

func (xdefiContext *XdefiContext) InitClients() {
	var xclient *mcclient.Client
	var yclient *mcclient.Client
	var err error

	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)
	xChainId := xdefiContext.Config.XChainId
	xChainFuncPrefix := xdefiContext.Config.XChainFuncPrefix
	xChainRPC := xdefiContext.Config.XChainRPC
	xChainWS := xdefiContext.Config.XChainWS
	yChainId := xdefiContext.Config.YChainId
	yChainFuncPrefix := xdefiContext.Config.YChainFuncPrefix
	yChainRPC := xdefiContext.Config.YChainRPC
	yChainWS := xdefiContext.Config.YChainWS

	xclient, err = mcclient.Dial(xChainRPC)
	if err != nil {
		logFuncErr("[InitClients] Unable to connect to network http rpc:%v\n", err)
		return
	}
	xclient.SetFuncPrefix(xChainFuncPrefix)
	// sanity check chain id
	xchainId_, err := xclient.ChainID(context.Background())
	if err != nil {
		log.Errorf("[InitClients]\t client chain id err: %v", err)
		return
	}
	if xChainId != xchainId_.Uint64() {
		log.Errorf(
			"[InitClients]Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			xchainId_, xChainId,
		)
		return
	}
	xclientWS, err := mcclient.Dial(xChainWS)
	if err != nil {
		logFunc("[InitClients] Unable to connect to network ws:%v\n", err)
		return
	}
	xclientWS.SetFuncPrefix(xChainFuncPrefix)

	///////////////////////////////////////////////////////////

	yclient, err = mcclient.Dial(yChainRPC)
	if err != nil {
		logFunc("[InitClients]Unable to connect to network:%v\n", err)
		return
	}
	yclient.SetFuncPrefix(yChainFuncPrefix)
	// sanity check chain id
	ychainId_, err := yclient.ChainID(context.Background())
	if err != nil {
		log.Errorf("[InitClients]\t client chain id err: %v", err)
		return
	}
	if yChainId != ychainId_.Uint64() {
		log.Errorf(
			"[InitClients]Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			ychainId_, yChainId,
		)
		return
	}
	yclientWS, err := mcclient.Dial(yChainWS)
	if err != nil {
		logFunc("[InitClients] Unable to connect to network ws:%v\n", err)
		return
	}
	yclientWS.SetFuncPrefix(yChainFuncPrefix)

	//////////////////////////////////////////////////////////////////////////

	clientXevents, err := mcclient.Dial(xdefiContext.Sentinel.Rpc)
	if err != nil {
		log.Errorf("[InitClients]\t client x err: %v %s", err, xdefiContext.Sentinel.Rpc)
		return
	}

	xdefiContext.ClientX = xclient
	xdefiContext.ClientY = yclient
	xdefiContext.ClientXws = xclientWS
	xdefiContext.ClientYws = yclientWS
	xdefiContext.ClientXevents = clientXevents

	return
}

func (xdefiContext *XdefiContext) InitContracts() {
	vaultXAddr := xdefiContext.Config.VaultXAddr
	vaultYAddr := xdefiContext.Config.VaultYAddr
	clientX := xdefiContext.ClientX
	clientY := xdefiContext.ClientY
	clientXws := xdefiContext.ClientXws
	clientYws := xdefiContext.ClientYws
	clientXevents := xdefiContext.ClientXevents

	vaultxContract, err := vaultx.NewVaultX(
		vaultXAddr,
		clientX,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t http client new vault X contract err: %v", err)
		return
	}

	vaultxContractWS, err := vaultx.NewVaultX(
		vaultXAddr,
		clientXws,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t ws client new vault X contract err: %v", err)
		return
	}

	vaultyContract, err := vaulty.NewVaultY(
		vaultYAddr,
		clientY,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t http client new vault Y contract err: %v", err)
		return
	}

	vaultyContractWS, err := vaulty.NewVaultY(
		vaultYAddr,
		clientYws,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t ws client new vault Y contract err: %v", err)
		return
	}

	// xevents contract
	xeventsXYContract, err := xevents.NewXEvents(
		XeventsXYAddr,
		clientXevents,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t client new xevents XY contract err: %v", err)
		return
	}

	xeventsYXContract, err := xevents.NewXEvents(
		XeventsYXAddr,
		clientXevents,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t client new xevents YX contract err: %v", err)
		return
	}

	// xconfig contract
	xconfigContract, err := xconfig.NewXConfig(
		XConfigAddr,
		clientXevents,
	)
	if err != nil {
		log.Errorf("[InitContracts]\t client new xconfig contract err: %v", err)
		return
	}

	xdefiContext.VaultX = vaultxContract
	xdefiContext.VaultY = vaultyContract
	xdefiContext.VaultXws = vaultxContractWS
	xdefiContext.VaultYws = vaultyContractWS
	xdefiContext.XeventsXY = xeventsXYContract
	xdefiContext.XeventsYX = xeventsYXContract
	xdefiContext.Xconfig = xconfigContract
	return
}

func (xdefiContext *XdefiContext) IsReady() bool {
	oneIsNil := xdefiContext.ClientX == nil
	oneIsNil = oneIsNil || xdefiContext.ClientY == nil
	oneIsNil = oneIsNil || xdefiContext.ClientXevents == nil
	oneIsNil = oneIsNil || xdefiContext.VaultX == nil
	oneIsNil = oneIsNil || xdefiContext.VaultY == nil
	oneIsNil = oneIsNil || xdefiContext.XeventsXY == nil
	oneIsNil = oneIsNil || xdefiContext.XeventsYX == nil

	if oneIsNil {
		log.Errorf(
			"[IsReady]\t Prepare common context: one of clients/contracts is nil: %v, # %v, # %v, # %v, # %v, # %v # %v       ",
			xdefiContext.ClientX, xdefiContext.ClientY, xdefiContext.ClientXevents,
			xdefiContext.VaultX, xdefiContext.VaultY,
			xdefiContext.XeventsXY, xdefiContext.XeventsYX,
		)
	}

	return !oneIsNil
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
		db:                   db,
		VaultEvents:          make(map[common.Hash]*VaultEventWithSig),
		VaultEventsReceived:  make(map[common.Hash]*set.Set),
		config:               bc.VnodeConfig(),
		vaultsConfig:         vaultsConfig,
		dkg:                  dkg,
		key:                  key,
		Rpc:                  "http://" + rpc,
		RpcCallPerClient:     make(map[string]map[string]int),
		VaultEventsBatchChan: make(chan VaultEventsBatch, VaultEventsBatchChanSize),
		PersistSeenVaultEventWithSigChan: make(
			chan PersistSeenVaultEventWithSig, PersistSeenVaultEventWithSigChanSize),
		depositChan:   make(map[common.Address]chan *vaultx.VaultXTokenDeposit),
		burnChan:      make(map[common.Address]chan *vaulty.VaultYTokenBurn),
		configVersion: uint64(0),
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
	ticker := time.NewTicker(ConfigCheckInterval * time.Second)
	for {
		select {
		case <-ticker.C:
			callOpts := &bind.CallOpts{}
			configVersion, err := sentinel.xconfig.ConfigVersion(callOpts)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}

			// update if sentinel's version is lower
			var newConfig VaultPairListConfig
			newConfigBytes, err := sentinel.xconfig.Configs(callOpts, configVersion)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}
			log.Infof("New Xconfig: %s", string(newConfigBytes))
			err = json.Unmarshal(newConfigBytes, &newConfig)
			if err != nil {
				log.Errorf("New xconfig update failed: %v", err)
				continue
			}
			sentinel.vaultsConfig = &newConfig
			sentinel.configVersion = configVersion.Uint64()

			// send out the config version regardless of its value
			// so that new feed subscriber will receive it
			sentinel.configVersionFeed.Send(configVersion.Uint64())
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

func (sentinel *Sentinel) SubscribeVaultEventWithSig(ch chan<- VaultEventWithSig) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventWithSigFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	log.Debugf("Vault events received:")
	for hash, received := range sentinel.VaultEventsReceived {
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

	nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	sentinel.LogRpcStat("vault", "NonceAt")
	if err != nil {
		return nil
	}
	transactor.Nonce = big.NewInt(int64(nonceAt))

	return transactor
}

func (sentinel *Sentinel) IsVaultEventsBatchAllReceived(batch VaultEventsBatch, threshold int) bool {
	for vaultEventHash, _ := range batch.Batch {
		if sentinel.VaultEventsReceived[vaultEventHash].Size() < threshold {
			return false
		}
	}
	return true
}

func (sentinel *Sentinel) RecordVaultEventWithSig(
	vaultEventWithSig *VaultEventWithSig,
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
	xdefiContext := &XdefiContext{}
	xdefiContext.EventType = eventType
	xdefiContext.Config = xdefiContextConfig
	xdefiContext.Sentinel = sentinel

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

func (sentinel *Sentinel) initVaultParams(xdefiContext *XdefiContext) uint64 {
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xevents.VaultWatermark(callOpts, vaultAddrFrom)
	sentinel.LogRpcStat("xevents", "VaultWatermark")
	if err != nil {
		log.Errorf("[initVaultParams]\t client watermark block err: %v       -----", err)
		return 0
	}

	lastBlock := uint64(0)
	var createdAt *big.Int
	// fallback to use createAt
	if vaultWatermark.Uint64() == 0 {
		if xdefiContext.IsDeposit() {
			createdAt, err = xdefiContext.VaultX.CreatedAt(callOpts)
			sentinel.LogRpcStat("vault", "CreateAt")
		} else {
			createdAt, err = xdefiContext.VaultY.CreatedAt(callOpts)
			sentinel.LogRpcStat("vault", "CreateAt")
		}
		if err != nil {
			log.Errorf(
				"[initVaultParams]\t client CreateAt err: %v, set lastBlock to %d",
				err, lastBlock,
			)
			return 0
		} else {
			// first time
			lastBlock = createdAt.Uint64()
		}
	} else {
		// normal case
		// we've scanned watermark block, proceed to next one
		lastBlock = vaultWatermark.Uint64() + 1
		if err != nil {
			log.Errorf("[initVaultParams]\t client vault store counter error %v      ", err)
			return 0
		}
	}

	// share with other go routine
	sentinel.batchNumber = lastBlock

	return lastBlock
}

func (sentinel *Sentinel) scanVaultEvents(
	xdefiContext *XdefiContext,
	lastBlock uint64,
) *VaultEventsBatch {
	eventCount := uint64(0)
	vaultX := xdefiContext.VaultX
	vaultY := xdefiContext.VaultY
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)
	clientFrom := xdefiContext.ClientFrom()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	startBlock := lastBlock
	lastCurrentBlock := uint64(0)

	// the outer for loop is to skip void blocks with no events
	for {
		// throttling for sending query
		time.Sleep(VaultEventScanInterval * time.Second)
		// once the filter return some events, we're done for this round.
		if eventCount > 0 {
			return nil
		}
		currentBlock, err := clientFrom.BlockNumber(context.Background())
		if currentBlock == lastCurrentBlock {
			logFunc("scan omitted []: block head does NOT change on target chain")
			continue
		} else {
			// new block from client FROM
			lastCurrentBlock = currentBlock
		}
		if err != nil {
			logFuncErr(
				"[scanVaultEvents]\t  sentinel: unable to get current block number from chain: %v", err)
		}
		currentBlock = currentBlock - BlockDelay
		if currentBlock < MinScanBlockNumber {
			return nil
		}

		// filter events from vaultx, [start, end] are inclusive
		ScanRange := ScanRangeBlock
		if lastBlock+ScanRangeBlock-1 > currentBlock {
			ScanRange = ScanRangeSingle
		}
		endBlock := lastBlock + ScanRange - 1

		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &endBlock,
		}

		logFunc("[scanVaultEvents]\t filter opts:[batch=%d], start: %d, end: %d",
			startBlock, lastBlock, endBlock,
		)

		//var itrY *vaulty.VaultYTokenDepositIterator
		batch := make(map[common.Hash]bool)
		if xdefiContext.IsDeposit() {
			itrX, err := vaultX.FilterTokenDeposit(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)
			sentinel.LogRpcStat("vault", "FilterTokenDeposit")

			if err != nil {
				logFuncErr(
					"[scanVaultEvents]\t sentinel: unable to get token deposit iterator: %v", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				logFuncErr("[scanVaultEvents]\t sentinel: unable to get bls signer: %v", err)
				return nil
			}

			for itrX.Next() {
				eventCount += 1
				event := itrX.Event
				vaultEvent := VaultEvent{
					vaultAddrFrom,
					event.SourceChainid,
					event.SourceToken,
					event.MappedChainid,
					event.MappedToken,
					event.From,
					event.Amount,
					event.Nonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				logFunc(
					"[scanVaultEvents]\t Scanned vault event: %x, nonce: %d, block: %d, source: %x, mapped: %x",
					vaultAddrFrom,
					event.Nonce,
					event.BlockNumber,
					event.SourceToken,
					event.MappedToken,
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.RecordVaultEventWithSig(&vaultEventWithSig)
				batch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			logFunc("[scanVaultEvents]\t raw filter result: batch size = %d", len(batch))
		} else {
			itrY, err := vaultY.FilterTokenBurn(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)
			sentinel.LogRpcStat("vault", "FilterTokenBurn")

			if err != nil {
				logFuncErr(
					"[scanVaultEvents]\t sentinel: unable to get token burn iterator: %v", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				logFuncErr("[scanVaultEvents]\t sentinel: unable to get bls signer: %v", err)
				return nil
			}

			for itrY.Next() {
				eventCount += 1
				event := itrY.Event
				vaultEvent := VaultEvent{
					vaultAddrFrom,
					event.SourceChainid,
					event.SourceToken,
					event.MappedChainid,
					event.MappedToken,
					event.From,
					event.Amount,
					event.Nonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				logFunc(
					"[scanVaultEvents]\t Scanned vault event: %x, nonce: %d, block: %d, source: %x, mapped: %x",
					vaultAddrFrom,
					event.Nonce,
					event.BlockNumber,
					event.SourceToken,
					event.MappedToken,
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.RecordVaultEventWithSig(&vaultEventWithSig)
				batch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			logFunc("[scanVaultEvents]\t raw filter result: batch size = %d", len(batch))
		}

		// once we receive the batch with events,
		// scan for this round is done, just return
		if sentinel.batchNumber != 0 && len(batch) > 0 {
			logFunc(
				"[scanVaultEvents]\t New batch mined (commit = %t), number: %d, start: %d, end: %d, events: %d",
				sentinel.myTurn(MyTurnSeed),
				sentinel.batchNumber,
				startBlock,
				endBlock,
				len(batch),
			)
			return &VaultEventsBatch{
				batch,
				vaultAddrFrom,
				startBlock,
				endBlock,
			}
		}

		// corner case: generate an empty batch so that we can move
		// the watermark higher for a gap between events
		if endBlock-startBlock >= MaxEmptyBatchBlocks && len(batch) == 0 {
			return &VaultEventsBatch{
				batch,
				vaultAddrFrom,
				startBlock,
				endBlock,
			}
		}

		lastBlock = endBlock + 1
	}
}

func (sentinel *Sentinel) commitBatch(
	xdefiContext *XdefiContext,
	batch *VaultEventsBatch,
) ([]uint64, []uint64, []uint64, *bind.TransactOpts) {
	clientXevents := xdefiContext.ClientXevents
	xevents := xdefiContext.Xevents()
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)

	logFunc(
		"[  commitBatch  ]\t Vault FROM commit batch BEGIN from %d to %d ]",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)
	defer logFunc(
		"[  commitBatch  ]\t Vault FROM commit batch END from %d to %d",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)

	balance, err := clientXevents.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	nonceAt, err := clientXevents.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	logFunc(
		"[  commitBatch  ]\t   %x, balance: %d, nonce: %d",
		sentinel.key.Address.Bytes(), balance, nonceAt,
	)
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}

	vaultWatermark, err := xevents.VaultWatermark(
		&bind.CallOpts{}, XeventsXYAddr,
	)
	if err != nil {
		logFuncErr("[  commitBatch  ]\t client watermark block err: %v", err)
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	sentinel.LogRpcStat("xevents", "VaultWatermark")

	// prevent re-entry
	if vaultWatermark.Uint64() >= batch.StartBlockNumber {
		logFuncErr(
			"[  commitBatch  ]\t Vault FROM skip batch commit: vault: %d, batch number: %d       -",
			vaultWatermark.Uint64(),
			batch.StartBlockNumber,
		)
		return []uint64{}, []uint64{}, []uint64{}, nil
	}

	// setup transactor
	transactor := sentinel.getTransactor(
		clientXevents,
		DefaultGasPrice,
		DefaultGasLimit,
	)

	// order the commit sequence by nonce
	nonces := make([]string, 0)
	noncesMap := make(map[string]common.Hash)
	for vaultEventHash, _ := range batch.Batch {
		vaultEvent := sentinel.VaultEvents[vaultEventHash].Event
		key := fmt.Sprintf(
			"%x,%x,%12d",
			vaultEvent.Vault,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
		nonces = append(nonces, key)
		noncesMap[key] = vaultEventHash
		logFunc(
			"[  commitBatch  ]\t Vault event: vault: %x, tokenMapping: %s, sha256: %x, nonce: %d",
			vaultEvent.Vault,
			vaultEvent.TokenMapping(),
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
	}
	// nonces is a string list of |vault|tokenmapping sha|nonce|
	sort.Strings(nonces)

	var vaultEventHash common.Hash
	logFunc("[  commitBatch  ]\t Vault FROM nonces: %v", nonces)

	committed := []uint64{}
	errors := []uint64{}
	succeeded := []uint64{}
	// commit all events in order
	for _, nonce := range nonces {
		vaultEventHash = noncesMap[nonce]
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event

		// if tokenmapping nonce is low, mark as succeeded
		tokenMappingWatermark, err := xevents.TokenMappingWatermark(
			&bind.CallOpts{}, vaultEvent.Vault, vaultEvent.TokenMappingSha256(),
		)
		sentinel.LogRpcStat("xevents", "VaultEventWatermark")
		if err != nil {
			errors = append(errors, vaultEvent.Nonce.Uint64())
			logFuncErr("[  commitBatch  ]\t Vault FROM watermark err: %v", err)
			continue
		} else {
			logFunc(
				"[  commitBatch  ]\t xevents tokenmapping watermark %d, this tokenmapping nonce: %d",
				tokenMappingWatermark, vaultEvent.Nonce,
			)
		}
		if tokenMappingWatermark.Uint64() > vaultEvent.Nonce.Uint64() {
			succeeded = append(succeeded, vaultEvent.Nonce.Uint64())
			logFunc(
				"[  commitBatch  ]\t xevents tokenmapping nonce low, nonce mark succeeded: %d",
				vaultEvent.Nonce.Uint64(),
			)
			continue
		}

		logFunc(
			"[  commitBatch  ]\t  xevents Store tx: vault: %x, tokenmapping: %x, nonce: %d, block: %d, %d, %d",
			vaultEvent.Vault.Bytes()[:8],
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
			vaultEvent.BlockNumber,
			transactor.GasPrice,
			transactor.GasLimit,
		)
		tx, err := xevents.RecordVaultEvent(
			transactor,
			vaultEventWithSig.Blssig,
			vaultEvent.Vault,
			vaultEvent.Nonce,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.BlockNumber,
			vaultEvent.Bytes(),
		)
		sentinel.LogRpcStat("xevents", "Store")
		transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)
		if err != nil {
			errors = append(errors, vaultEvent.Nonce.Uint64())
			logFuncErr("[  commitBatch  ]\t Vault FROM sentinel xevents store err: %v", err)
			continue
		} else {
			logFunc(
				"[  commitBatch  ]\t Vault FROM sentinel xevents store tx: %x, from: %x, nonce: %d       ",
				tx.Hash(), sentinel.key.Address.Bytes(), tx.Nonce(),
			)
			committed = append(committed, vaultEvent.Nonce.Uint64())
		}
	}

	return succeeded, committed, errors, transactor
}

func (sentinel *Sentinel) watchVault(xdefiContext *XdefiContext) {
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)
	defer logFunc("*********************END WATCHER LOOP***********************")

	ticker := time.NewTicker(VaultCheckInterval * time.Second)
	newConfigChan := make(chan uint64, 10)
	sub := sentinel.configVersionFeed.Subscribe(newConfigChan)

	for {
		select {
		case err := <-sub.Err():
			logFuncErr("new config sub err: %v", err)
			break
		case newConfigVersion := <-newConfigChan:
			// if there is a new config, stop all current watchers
			// since new ones will be created with updated config.
			// The if statement prevents accidentally kill new watchers.
			if newConfigVersion > xdefiContext.ConfigVersion {
				break
			}
		case <-ticker.C:
			sentinel.PrintRpcStat()
			// # 0
			xevents := xdefiContext.Xevents()
			lastBlock := sentinel.initVaultParams(xdefiContext)

			// # 1
			batch := sentinel.scanVaultEvents(
				xdefiContext,
				lastBlock,
			)

			if batch == nil {
				continue
			}

			if sentinel.myTurn(MyTurnSeed) {
				blockNumberBefore, _ := xdefiContext.ClientXevents.BlockNumber(context.Background())

				//  commit the batch
				succeeded, committed, errors, transactor := sentinel.commitBatch(
					xdefiContext,
					batch,
				)
				logFunc(
					"[  In WatchVault  ]\t Before MOVE vault watermark, batch size: %d, succeeded: %v, commit %v, errors: %v",
					len(batch.Batch),
					succeeded,
					committed,
					errors,
				)
				// something's seriously wrong, we don't even get transactor
				if transactor == nil {
					continue
				}

				// wait for the confirmation
				for {
					time.Sleep(3 * time.Second)
					blockNumberAfter, _ := xdefiContext.ClientXevents.BlockNumber(context.Background())
					// wait till one block is mined and all txs in the batch confirmed
					if blockNumberAfter > blockNumberBefore && len(succeeded) == len(batch.Batch) {
						xevents.UpdateVaultWatermark(
							transactor,
							batch.Vault,
							big.NewInt(int64(batch.EndBlockNumber)),
						)
						logFunc("[  In WatchVault  ]\t MOVE vault watermark HIGHER to %d",
							batch.EndBlockNumber,
						)
						break
					}
					if blockNumberAfter-blockNumberBefore > BatchCommitThreshold {
						logFuncErr(
							"[ In WatchVault ]\t batch commit wait threshold time out: block [%d -> %d]",
							blockNumberBefore, blockNumberAfter,
						)
						break
					}
				}
			}
		}
	}
}

func (sentinel *Sentinel) ForwardVaultEvents(
	xdefiContext *XdefiContext,
	tokenMapping TokenMapping,
) {
	clientTo := xdefiContext.ClientTo()
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)

	callOpts := &bind.CallOpts{}
	var waterMark *big.Int
	var err error
	if xdefiContext.IsDeposit() {
		waterMark, err = xdefiContext.VaultY.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		sentinel.LogRpcStat("vault", "TokenMappingWatermark")
		if err != nil {
			logFuncErr("[ForwardVaultEvents]\t Vault TO: watermark err: %v", err)
			return
		}
		logFunc("[ForwardVaultEvents]\t Vault TO: watermark: %d, %s, %s",
			waterMark,
			tokenMapping.SourceToken,
			tokenMapping.MappedToken,
		)
	} else {
		waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		sentinel.LogRpcStat("vault", "TokenMappingWatermark")
		if err != nil {
			logFuncErr("[ForwardVaultEvents]\t Vault TO: watermark err: %v", err)
			return
		}
		logFunc("[ForwardVaultEvents]\t Vault TO: watermark: %d, %s, %s",
			waterMark,
			tokenMapping.SourceToken,
			tokenMapping.MappedToken,
		)

	}

	// each node will mint for n blocks, then hand it over to next node
	if !sentinel.myTurn(MyTurnSeed) {
		logFunc("[ForwardVaultEvents]\t Vault TO: Not My Turn to Mint")
		return
	} else {
		logFunc("[ForwardVaultEvents]\t Vault TO: My Turn to Mint")
	}

	// setup transactor
	transactor := sentinel.getTransactor(
		clientTo,
		DefaultGasPrice,
		DefaultGasLimit,
	)
	if transactor == nil {
		return
	}

	// check balance
	balance, err := clientTo.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return
	}
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		logFuncErr(
			"[ForwardVaultEvents]\t No balance for account on Vault TO chain: %x",
			sentinel.key.Address,
		)
		return
	}

	committed := uint64(0)
	errors := uint64(0)
	for i := int64(0); i < int64(MintBatchSize); i++ {
		var vaultEvent VaultEvent
		vaultEventData, err := xevents.VaultEvents(
			callOpts, vaultAddrFrom,
			tokenMapping.Sha256(),
			big.NewInt(waterMark.Int64()+i),
		)
		sentinel.LogRpcStat("xevents", "VaultEvents")
		if err != nil {
			logFuncErr("[ForwardVaultEvents]\t  Xevents: retrieve vault events err: %v", err)
		}
		rlp.DecodeBytes(vaultEventData.EventData, &vaultEvent)
		logFunc(
			"[ForwardVaultEvents]\t Vault [TO] TO MINT: nonce: %d amount: %d, tip: %d, to: %x, mappedToken: %x",
			big.NewInt(waterMark.Int64()+i),
			vaultEvent.Amount,
			vaultEvent.Tip,
			vaultEvent.To.Bytes,
			vaultEvent.MappedToken.Bytes,
		)
		// if no more vault event
		if vaultEvent.Amount == nil || vaultEvent.Amount.Int64() == 0 {
			break
		} else {
			if xdefiContext.IsDeposit() {
				tx, err := xdefiContext.VaultY.Mint(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				sentinel.LogRpcStat("vault", "Mint")
				transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

				if err != nil {
					errors += 1
					log.Errorf("[ForwardVaultEvents]\t sentinel vault TO 'mint' tx err: %v", err)
				} else {
					logFunc(
						"[ForwardVaultEvents]\t sentinel vault [TO] mint() tx: %x, from: %x, nonce: %d, %d, %d       ",
						tx.Hash(),
						sentinel.key.Address.Bytes(),
						tx.Nonce(),
						transactor.GasPrice,
						transactor.GasLimit,
					)
					committed += 1
				}
			} else {
				tx, err := xdefiContext.VaultX.Withdraw(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				sentinel.LogRpcStat("vault", "Withdraw")
				transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

				if err != nil {
					errors += 1
					log.Errorf("[ForwardVaultEvents]\t sentinel vault TO 'withdraw' tx err: %v", err)
				} else {
					logFunc(
						"[ForwardVaultEvents]\t sentinel vault TO withdraw() tx: %x, from: %x, nonce: %d, %d, %d       ",
						tx.Hash(),
						sentinel.key.Address.Bytes(),
						tx.Nonce(),
						transactor.GasPrice,
						transactor.GasLimit,
					)
					committed += 1
				}
			}
		}
	}

	waterMarkBefore := waterMark.Uint64()
	// wait for watermark to update
	wait := 0
	for {
		if committed == 0 {
			logFunc("[ForwardVaultEvents]\t No vault TO mint/withdraw tx called")
			break
		}
		var waterMark *big.Int
		if xdefiContext.IsDeposit() {
			waterMark, err = xdefiContext.VaultY.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
			sentinel.LogRpcStat("vault", "TokenMappingWatermark")
			logFunc(
				"[ForwardVaultEvents]\t Vault TO: watermark: %d, before: %d, committed: %d",
				waterMark, waterMarkBefore, committed,
			)
		} else {
			waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
			sentinel.LogRpcStat("vault", "TokenMappingWatermark")
			logFunc(
				"[ForwardVaultEvents]\t Vault TO: watermark: %d, before: %d, committed: %d       ",
				waterMark, waterMarkBefore, committed,
			)
		}
		if err != nil {
			log.Errorf("[ForwardVaultEvents]\t Err with calling vault Y watermark: %v", err)
			continue
		}
		if waterMark.Uint64() == waterMarkBefore+committed {
			break
		}
		time.Sleep(10 * time.Second)
		wait += 1
		if wait > 12 {
			logFunc(
				"[ForwardVaultEvents]\t Vault Y watermark wait time out, watermark: %d, before: %d, committed: %d",
				waterMark, waterMarkBefore, committed,
			)
			break
		}
	}
}

func (sentinel *Sentinel) doForward(
	xdefiContext *XdefiContext, tokenMapping TokenMapping,
) {
	logFunc := xdefiContext.LogFunc(false)
	defer logFunc("*********************END LOOP Foward Vault Events***********************")

	ticker := time.NewTicker(VaultCheckInterval * time.Second)
	newConfigChan := make(chan uint64, 10)
	sub := sentinel.configVersionFeed.Subscribe(newConfigChan)

	for {
		select {
		case <-sub.Err():
			break
		case newConfigVersion := <-newConfigChan:
			if newConfigVersion > xdefiContext.ConfigVersion {
				break
			}
		case <-ticker.C:
			sentinel.ForwardVaultEvents(
				xdefiContext,
				tokenMapping,
			)
		}
	}
}

func (sentinel *Sentinel) startWatchersAndForwarders() {
	defer log.Errorf("*********************END WATCHER BOOTER LOOP***********************")

	newConfigChan := make(chan uint64, 10)
	sub := sentinel.configVersionFeed.Subscribe(newConfigChan)

	for {
		select {
		case <-sub.Err():
			break
		case newConfigVersion := <-newConfigChan:
			// exhaust the channel
			for i := 0; i < len(newConfigChan); i++ {
				newConfigVersion = <-newConfigChan
			}
			log.Infof("sentinel: new config version received: %d", newConfigVersion)

			if newConfigVersion != sentinel.configVersion {
				log.Errorf(
					"mismatch between chan config version: %d and sentinel config version: %d",
					newConfigVersion, sentinel.configVersion,
				)
				break
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
				go sentinel.watchVault(xdefiContextXY)
				for _, tokenMapping := range vaultPairConfig.TokenMappings {
					go sentinel.doForward(xdefiContextXY, tokenMapping)
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
				go sentinel.watchVault(xdefiContextYX)
				for _, tokenMapping := range vaultPairConfig.TokenMappings {
					go sentinel.doForward(xdefiContextYX, tokenMapping)
				}

				// sanity check
				if !(xdefiContextXY.EventType == DEPOSIT && xdefiContextYX.EventType == BURN) {
					panic("Missing Deposit or Burn go routine")
				}

				go sentinel.WatchVaultXLogs(xdefiContextXY)
				go sentinel.WatchVaultYLogs(xdefiContextYX)
			}
		}
	}
}

func (sentinel *Sentinel) WatchVaultXLogs(xdefiContext *XdefiContext) {
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)
	defer logFuncErr("Watch vault X logs exit")

	start := uint64(1)
	optsX := new(bind.WatchOpts)
	optsX.Start = &start
	optsX.Context = context.Background()

	// vault event chan
	if _, found := sentinel.depositChan[xdefiContext.Config.VaultXAddr]; !found {
		sentinel.depositChan[xdefiContext.Config.VaultXAddr] = make(chan *vaultx.VaultXTokenDeposit)
	}
	depositChan := sentinel.depositChan[xdefiContext.Config.VaultXAddr]
	subX, err := xdefiContext.VaultXws.WatchTokenDeposit(
		optsX,
		depositChan,
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		logFuncErr("Can not watch vault X logs: %v", err)
		return
	}

	// new config version
	newConfigChan := make(chan uint64, 10)
	subConfigVersion := sentinel.configVersionFeed.Subscribe(newConfigChan)

	for {
		select {
		case <-subConfigVersion.Err():
			break
		case newConfigVersion := <-newConfigChan:
			if newConfigVersion > xdefiContext.ConfigVersion {
				break
			}
		case event := <-depositChan:
			logFunc("notified with NEW DEPOSIT: vault: %x, nonce: %d, block: %d, source: %x, mapped: %x",
				event.Vault.Bytes()[:4],
				event.Nonce,
				event.BlockNumber,
				event.SourceToken.Bytes()[:4],
				event.MappedToken.Bytes()[:4],
			)
		case err := <-subX.Err():
			logFuncErr("Can not watch vault X logs, sub err: %v", err)
			break
		}
	}
}

func (sentinel *Sentinel) WatchVaultYLogs(xdefiContext *XdefiContext) {
	logFunc := xdefiContext.LogFunc(false)
	logFuncErr := xdefiContext.LogFunc(true)
	defer logFuncErr("Watch vault Y logs exit")

	start := uint64(1)
	optsY := new(bind.WatchOpts)
	optsY.Start = &start
	optsY.Context = context.Background()

	// vault event chan
	if _, found := sentinel.burnChan[xdefiContext.Config.VaultYAddr]; !found {
		sentinel.burnChan[xdefiContext.Config.VaultYAddr] = make(chan *vaulty.VaultYTokenBurn)
	}
	burnChan := sentinel.burnChan[xdefiContext.Config.VaultYAddr]
	subY, err := xdefiContext.VaultYws.WatchTokenBurn(
		optsY,
		burnChan,
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)

	// new config version
	newConfigChan := make(chan uint64, 10)
	subConfigVersion := sentinel.configVersionFeed.Subscribe(newConfigChan)

	if err != nil {
		log.Errorf("can not watch vault Y logs: %v", err)
		return
	}

	for {
		select {
		case <-subConfigVersion.Err():
			break
		case newConfigVersion := <-newConfigChan:
			if newConfigVersion > xdefiContext.ConfigVersion {
				break
			}
		case event := <-burnChan:
			logFunc("notified with NEW BURN: vault: %x, nonce: %d, block: %d, source: %x, mapped: %x",
				event.Vault.Bytes()[:4],
				event.Nonce,
				event.BlockNumber,
				event.SourceToken.Bytes()[:4],
				event.MappedToken.Bytes()[:4],
			)
		case err := <-subY.Err():
			log.Errorf("can not watch vault Y logs, sub err: %v", err)
			break
		}
	}
}
