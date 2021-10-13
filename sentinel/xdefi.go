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
	"sort"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/rlp"
	"github.com/MOACChain/MoacLib/types"
	moaccore "github.com/MOACChain/xchain"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/xdefi/vaultx"
	"github.com/MOACChain/xchain/xdefi/vaulty"
	"github.com/MOACChain/xchain/xdefi/xconfig"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

const (
	VaultCheckInterval                = 5 // in seconds
	VaultEventsBatchInterval          = 5 // in seconds
	BatchCommitThreshold              = 3 // in blocks
	SigCheckDelay                     = 3 // in seconds
	ScanRangeMulti                    = uint64(20)
	ScanRangeDouble                   = uint64(2)
	MinScanBlockNumber                = uint64(20)
	CatchupWithinBlocks               = uint64(100)
	VaultEventsBatchQueueSize         = 100
	VaultEventsBatchPendingSize       = 100
	CommittedVaultEventsBatchChanSize = 100
	VaultEventScanInterval            = 2
	MaxBlockNumber                    = uint64(1000000000000)
	Gwei                              = int64(1000000000)
	DefaultGasPrice                   = int64(20) * Gwei
	DefaultGasLimit                   = int64(6000000)
	BlockDelay                        = 12
	MaxEmptyBatchBlocks               = 60
	MintBatchSize                     = 30
	MintIntervalBlocks                = 50
	MyTurnSeed                        = 10000

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

type XdefiVaultEvent struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

type XdefiVaultEventIterator struct {
	Event    *XdefiVaultEvent      // Event containing the contract specifics and raw log
	contract *bind.BoundContract   // Generic contract to use for unpacking event data
	event    string                // Event name to use for unpacking event data
	logs     chan types.Log        // Log channel receiving the found contract events
	sub      moaccore.Subscription // Subscription for errors, completion and termination
	done     bool                  // Whether the subscription completed delivering logs
	fail     error                 // Occurred error to stop iteration
}

type XdefiVault interface {
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
	EventType                     uint64
	Config                        *XdefiContextConfig
	ConfigVersion                 uint64
	Sentinel                      *Sentinel
	ClientX                       *mcclient.Client
	ClientY                       *mcclient.Client
	ClientXevents                 *mcclient.Client
	VaultX                        *vaultx.VaultX
	VaultY                        *vaulty.VaultY
	XeventsXY                     *xevents.XEvents
	XeventsYX                     *xevents.XEvents
	ClientXws                     *mcclient.Client
	ClientYws                     *mcclient.Client
	VaultXws                      *vaultx.VaultX
	VaultYws                      *vaulty.VaultY
	Xconfig                       *xconfig.XConfig
	VaultScanBitmap               map[common.Address]map[uint64]bool
	QueueVaultEventsBatchChan     chan *VaultEventsBatch
	PendingVaultEventsBatchChan   chan *VaultEventsBatch
	CommittedVaultEventsBatchFeed event.Feed
}

func NewXdefiContext(
	eventType uint64,
	xdefiContextConfig *XdefiContextConfig,
	sentinel *Sentinel,
	configVersion uint64,
) *XdefiContext {
	xdefiContext := &XdefiContext{}
	xdefiContext.EventType = eventType
	xdefiContext.Config = xdefiContextConfig
	xdefiContext.Sentinel = sentinel
	xdefiContext.ConfigVersion = configVersion
	xdefiContext.QueueVaultEventsBatchChan = make(chan *VaultEventsBatch, VaultEventsBatchQueueSize)
	xdefiContext.PendingVaultEventsBatchChan = make(chan *VaultEventsBatch, VaultEventsBatchPendingSize)
	xdefiContext.VaultScanBitmap = make(map[common.Address]map[uint64]bool)

	return xdefiContext
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

func (xdefiContext *XdefiContext) LogInfo() func(format string, args ...interface{}) {
	return xdefiContext.logFunc(false)
}

func (xdefiContext *XdefiContext) LogErr() func(format string, args ...interface{}) {
	return xdefiContext.logFunc(true)
}

func (xdefiContext *XdefiContext) logFunc(err bool) func(format string, args ...interface{}) {
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

func (xdefiContext *XdefiContext) RecordVaultScan(vault common.Address, blockStart, blockEnd uint64) {
	if _, found := xdefiContext.VaultScanBitmap[vault]; !found {
		xdefiContext.VaultScanBitmap[vault] = make(map[uint64]bool)
	}

	for i := blockStart; i <= blockEnd; i++ {
		xdefiContext.VaultScanBitmap[vault][i] = true
	}
}

// remove records that are under the watermark
func (xdefiContext *XdefiContext) ClearVaultScan(vault common.Address, watermark uint64) {
	if vaultScan, found := xdefiContext.VaultScanBitmap[vault]; !found {
		for block, _ := range vaultScan {
			if block < watermark {
				delete(vaultScan, block)
			}
		}
	}
}

// see if we have already scanned the block range for the vault
func (xdefiContext *XdefiContext) ScanExists(vault common.Address, blockStart, blockEnd uint64) bool {
	if vaultScan, found := xdefiContext.VaultScanBitmap[vault]; found {
		for i := blockStart; i <= blockEnd; i++ {
			if _, exist := vaultScan[i]; !exist {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func (xdefiContext *XdefiContext) InitClients() {
	var xclient *mcclient.Client
	var yclient *mcclient.Client
	var err error

	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
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
		LogErr("[InitClients] Unable to connect to network http rpc:%v\n", err)
		return
	}
	xclient.SetFuncPrefix(xChainFuncPrefix)
	// sanity check chain id
	xchainId_, err := xclient.ChainID(context.Background())
	if err != nil {
		LogErr("[InitClients]\t client chain id err: %v", err)
		return
	}
	if xChainId != xchainId_.Uint64() {
		err := fmt.Errorf(
			"[InitClients]Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			xchainId_, xChainId,
		)
		panic(err)
	}
	xclientWS, err := mcclient.Dial(xChainWS)
	if err != nil {
		LogErr("[InitClients] Unable to connect to network ws:%v\n", err)
	} else {
		xclientWS.SetFuncPrefix(xChainFuncPrefix)
	}

	///////////////////////////////////////////////////////////

	yclient, err = mcclient.Dial(yChainRPC)
	if err != nil {
		LogInfo("[InitClients]Unable to connect to network:%v\n", err)
		return
	}
	yclient.SetFuncPrefix(yChainFuncPrefix)
	// sanity check chain id
	ychainId_, err := yclient.ChainID(context.Background())
	if err != nil {
		LogErr("[InitClients]\t client chain id err: %v", err)
		return
	}
	if yChainId != ychainId_.Uint64() {
		err := fmt.Errorf(
			"[InitClients]Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			ychainId_, yChainId,
		)
		panic(err)
	}
	yclientWS, err := mcclient.Dial(yChainWS)
	if err != nil {
		LogErr("[InitClients] Unable to connect to network ws:%v\n", err)
	} else {
		yclientWS.SetFuncPrefix(yChainFuncPrefix)
	}

	//////////////////////////////////////////////////////////////////////////

	clientXevents, err := mcclient.Dial(xdefiContext.Sentinel.Rpc)
	if err != nil {
		LogErr("[InitClients]\t client xevents err: %v %s", err, xdefiContext.Sentinel.Rpc)
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
		err := fmt.Errorf("[InitContracts]\t http client new vault X contract err: %v", err)
		panic(err)
	}

	vaultxContractWS, err := vaultx.NewVaultX(
		vaultXAddr,
		clientXws,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t ws client new vault X contract err: %v", err)
		panic(err)
	}

	vaultyContract, err := vaulty.NewVaultY(
		vaultYAddr,
		clientY,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t http client new vault Y contract err: %v", err)
		panic(err)
	}

	vaultyContractWS, err := vaulty.NewVaultY(
		vaultYAddr,
		clientYws,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t ws client new vault Y contract err: %v", err)
		panic(err)
	}

	// xevents contract
	xeventsXYContract, err := xevents.NewXEvents(
		XeventsXYAddr,
		clientXevents,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t client new xevents XY contract err: %v", err)
		panic(err)
	}

	xeventsYXContract, err := xevents.NewXEvents(
		XeventsYXAddr,
		clientXevents,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t client new xevents YX contract err: %v", err)
		panic(err)
	}

	// xconfig contract
	xconfigContract, err := xconfig.NewXConfig(
		XConfigAddr,
		clientXevents,
	)
	if err != nil {
		err := fmt.Errorf("[InitContracts]\t client new xconfig contract err: %v", err)
		panic(err)
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
			"[IsReady]\t Prepare common context: one of clients/contracts "+
				"is nil: %v, # %v, # %v, # %v, # %v, # %v # %v       ",
			xdefiContext.ClientX, xdefiContext.ClientY, xdefiContext.ClientXevents,
			xdefiContext.VaultX, xdefiContext.VaultY,
			xdefiContext.XeventsXY, xdefiContext.XeventsYX,
		)
	}

	return !oneIsNil
}

func (xdefiContext *XdefiContext) scanVaultEvents(
	sentinel *Sentinel, newConfigChan chan uint64,
) *VaultEventsBatch {
	lastBlock := xdefiContext.getLastBlock(sentinel)
	eventCount := uint64(0)
	vaultX := xdefiContext.VaultX
	vaultY := xdefiContext.VaultY
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	clientFrom := xdefiContext.ClientFrom()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	startBlock := lastBlock
	lastCurrentBlock := uint64(0)
	defer LogInfo("[scanVaultEvents]\t batch = %d scan ends [SCAN END]", startBlock)

	// the outer for loop is to skip void blocks with no events
	for {
		// throttling for sending query
		time.Sleep(VaultEventScanInterval * time.Second)
		// once the last filter call returns events, we're done for this round.
		if eventCount > 0 {
			return nil
		}

		// if there is new config, return
		if len(newConfigChan) > 0 {
			newConfigVersion := <-newConfigChan
			LogInfo("[scanVaultEvents]\t new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				// put it back to chan so parent func will check and exit
				newConfigChan <- newConfigVersion
				return nil
			}
		}

		currentBlock, err := clientFrom.BlockNumber(context.Background())
		if err != nil {
			LogErr(
				"[scanVaultEvents]\t  sentinel: unable to get current block number from chain: %v", err)
		}
		if currentBlock == lastCurrentBlock && lastCurrentBlock-lastBlock < CatchupWithinBlocks {
			LogInfo("[scanVaultEvents]\t scan omitted: block head does NOT change on target chain")
			continue
		}
		lastCurrentBlock = currentBlock

		// set scan horizon
		scanHorizon := currentBlock - BlockDelay
		// if blockchain just started producing block
		if scanHorizon < MinScanBlockNumber {
			return nil
		}

		// filter events from vaultx, [start, end] are inclusive
		var endBlock uint64
		if lastBlock+ScanRangeMulti-1 <= scanHorizon {
			endBlock = lastBlock + ScanRangeMulti - 1
		} else {
			endBlock = lastBlock + ScanRangeDouble - 1
		}

		// sanity check lastblock and endblock
		if endBlock > scanHorizon {
			endBlock = scanHorizon
		}
		if lastBlock > endBlock {
			lastBlock = endBlock
		}

		rawBatch := make(map[common.Hash]bool)
		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &endBlock,
		}
		LogInfo(
			"[scanVaultEvents]\t filter opts:[batch=%d], start: %d, end: %d, cur: %d, horz: %d, step: %d",
			startBlock, lastBlock, endBlock, currentBlock, scanHorizon, endBlock-lastBlock,
		)
		if xdefiContext.IsDeposit() {
			itrX, err := vaultX.FilterTokenDeposit(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)
			sentinel.LogRpcStat("vault", "FilterTokenDeposit")

			if err != nil {
				LogErr(
					"[scanVaultEvents]\t sentinel: unable to get token deposit iterator: %v", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				LogErr("[scanVaultEvents]\t sentinel: unable to get bls signer: %v", err)
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
					event.TotalNonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				LogInfo(
					"[scanVaultEvents]\t Scanned vault event *****: %x, nonce: %d, "+
						"block: %d, source: %x, mapped: %x",
					vaultAddrFrom.Bytes()[:4],
					event.Nonce,
					event.BlockNumber,
					event.SourceToken.Bytes()[:4],
					event.MappedToken.Bytes()[:4],
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.ProcessVaultEventWithSig(&vaultEventWithSig)
				rawBatch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			xdefiContext.RecordVaultScan(vaultAddrFrom, lastBlock, endBlock)
			LogInfo(
				"[scanVaultEvents]\t filter result[batch=%d, start=%d, end=%d]: raw batch size = %d",
				startBlock, lastBlock, endBlock, len(rawBatch),
			)
		} else {
			itrY, err := vaultY.FilterTokenBurn(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)
			sentinel.LogRpcStat("vault", "FilterTokenBurn")

			if err != nil {
				LogErr(
					"[scanVaultEvents]\t sentinel: unable to get token burn iterator: %v", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				LogErr("[scanVaultEvents]\t sentinel: unable to get bls signer: %v", err)
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
					event.TotalNonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				LogInfo(
					"[scanVaultEvents]\t Scanned vault event: %x, nonce: %d, "+
						"block: %d, source: %x, mapped: %x",
					vaultAddrFrom.Bytes()[:4],
					event.Nonce,
					event.BlockNumber,
					event.SourceToken.Bytes()[:4],
					event.MappedToken.Bytes()[:4],
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.ProcessVaultEventWithSig(&vaultEventWithSig)
				rawBatch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			xdefiContext.RecordVaultScan(vaultAddrFrom, lastBlock, endBlock)
			LogInfo(
				"[scanVaultEvents]\t filter result[batch=%d, start=%d, end=%d]: raw batch size = %d",
				startBlock, lastBlock, endBlock, len(rawBatch),
			)
		}

		// 2 return scenarios: 1) empty or 2)with vault events
		if len(rawBatch) == 0 {
			// corner case: generate an empty batch so that we can move
			// the watermark higher for a gap between events. We reach here
			// because all previous filters returns empty results and
			// this filter returns empty result as well
			if endBlock-startBlock >= MaxEmptyBatchBlocks {
				LogInfo("[scanVaultEvents]\t max empty batch blocks reached, startBlock: %d, endBlock %d",
					startBlock, endBlock,
				)
				return &VaultEventsBatch{
					rawBatch,
					vaultAddrFrom,
					startBlock,
					endBlock,
				}
			}
		} else {
			// once we scanned any vault events,
			// this round is done, just return the batch
			if sentinel.batchNumber != 0 {
				LogInfo(
					"[scanVaultEvents]\t New batch mined (commit = %t), "+
						"number: %d, start: %d, end: %d, events: %d",
					sentinel.myTurn(MyTurnSeed),
					sentinel.batchNumber,
					startBlock,
					endBlock,
					len(rawBatch),
				)
				return &VaultEventsBatch{
					rawBatch,
					vaultAddrFrom,
					startBlock,
					endBlock,
				}
			}
		}

		// step forward
		lastBlock = endBlock + 1
	}
}

func (xdefiContext *XdefiContext) ScanVaultEvents(sentinel *Sentinel) {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	worker := "scanVaultEvents"
	LogInfo("*********************enter scan vault events loop***********************")
	sentinel.addWorker(worker)
	defer LogInfo("*********************end scan vault events loop***********************")
	defer sentinel.reduceWorker(worker)

	ticker := time.NewTicker(VaultCheckInterval * time.Second)

	// setup feeds
	newConfigChan := make(chan uint64, 10)
	subNewConfig := sentinel.configVersionFeed.Subscribe(newConfigChan)
	committedVaultEventsBatchChan := make(chan *VaultEventsBatch, CommittedVaultEventsBatchChanSize)
	subCommittedBatch := xdefiContext.CommittedVaultEventsBatchFeed.Subscribe(
		committedVaultEventsBatchChan,
	)
	defer subNewConfig.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case err := <-subCommittedBatch.Err():
			LogErr("[scanVaultEvents]\t committed batch sub err: %v", err)
		case committedBatch := <-committedVaultEventsBatchChan:
			LogInfo("[scanVaultEvents]\t committed batch: %v", committedBatch)
		case err := <-subNewConfig.Err():
			LogErr("[scanVaultEvents]\t new config sub err: %v", err)
			break
		case newConfigVersion := <-newConfigChan:
			LogInfo("[scanVaultEvents]\t new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			// if there is a new config, stop all current watchers
			// since new ones will be created with updated config.
			// The if statement prevents accidentally kill new watchers.
			if newConfigVersion > xdefiContext.ConfigVersion {
				return
			}
		case <-ticker.C:
			if len(newConfigChan) > 0 {
				newConfigVersion := <-newConfigChan
				LogInfo("[scanVaultEvents]\t new config version received: %d -> %d",
					xdefiContext.ConfigVersion, newConfigVersion,
				)
				if newConfigVersion > xdefiContext.ConfigVersion {
					return
				}
			}
			batch := xdefiContext.scanVaultEvents(sentinel, newConfigChan)
			if batch == nil {
				continue
			}

			// send to vault events batch queue
			xdefiContext.QueueVaultEventsBatchChan <- batch
		}
	}
}

func (xdefiContext *XdefiContext) QueueVaultEventsBatch(sentinel *Sentinel) {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	worker := "QueueVaultEventsBatch"
	LogInfo("*********************enter queue vault events loop***********************")
	sentinel.addWorker(worker)
	defer LogInfo("QueueVaultEventsBatch loop exit")
	defer sentinel.reduceWorker(worker)

	// feed setup
	newConfigChan := make(chan uint64, 10)
	subNewConfig := sentinel.configVersionFeed.Subscribe(newConfigChan)
	committedVaultEventsBatchChan := make(chan *VaultEventsBatch, CommittedVaultEventsBatchChanSize)
	subCommittedBatch := xdefiContext.CommittedVaultEventsBatchFeed.Subscribe(
		committedVaultEventsBatchChan,
	)
	defer subNewConfig.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case err := <-subCommittedBatch.Err():
			LogErr("[QueueVaultEventsBatch] committed batch sub err: %v", err)
		case committedBatch := <-committedVaultEventsBatchChan:
			LogInfo("[QueueVaultEventsBatch] committed batch: %v", committedBatch)
		case newConfigVersion := <-newConfigChan:
			LogInfo("[QueueVaultEventsBatch] new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				return
			}
		case err := <-subNewConfig.Err():
			LogErr("[QueueVaultEventsBatch] new config sub err: %v", err)
		case batch := <-xdefiContext.QueueVaultEventsBatchChan:
			// it is empty, no need to wait for sigs
			if len(batch.Batch) == 0 {
				xdefiContext.PendingVaultEventsBatchChan <- batch
			} else {
				// wait for sigs
				time.Sleep(SigCheckDelay * time.Second)

				// if we received all sigs
				if readyCount, received := xdefiContext.ReceivedAllSigs(batch, sentinel); received {
					LogInfo("[QueueVaultEventsBatch] All sigs received for batch [%d], %d/%d",
						batch.Start, readyCount, len(batch.Batch),
					)
					xdefiContext.PendingVaultEventsBatchChan <- batch
				} else {
					LogInfo("[QueueVaultEventsBatch] Not enough sigs for batch [%d], %d/%d",
						batch.Start, readyCount, len(batch.Batch),
					)
				}
			}
		}
	}
}

func (xdefiContext *XdefiContext) ReceivedAllSigs(batch *VaultEventsBatch, sentinel *Sentinel) (int, bool) {
	readyCount := 0
	for vaultEventHash, _ := range batch.Batch {
		// if sigs do not exist or not enough sigs
		if sigs, found := sentinel.VaultEventsSigs[vaultEventHash]; found {
			if sigs.Size() >= sentinel.Threshold() {
				readyCount += 1
			}
		}
	}

	return readyCount, readyCount == len(batch.Batch)
}

func (xdefiContext *XdefiContext) PendingVaultEventsBatch(sentinel *Sentinel) {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	worker := "PendingVaultEventsBatch"
	LogInfo("*********************enter pending vault events loop***********************")
	sentinel.addWorker(worker)
	defer LogInfo("exit pending vault events loop")
	defer sentinel.reduceWorker(worker)

	newConfigChan := make(chan uint64, 10)
	subNewConfig := sentinel.configVersionFeed.Subscribe(newConfigChan)
	defer subNewConfig.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case err := <-subNewConfig.Err():
			LogErr("[PendingVaultEvents] new config sub err: %v", err)
		case newConfigVersion := <-newConfigChan:
			LogInfo("[PendingVaultEvents] new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				return
			}
		case batch := <-xdefiContext.PendingVaultEventsBatchChan:
			if sentinel.myTurn(MyTurnSeed) {
				//  commit the batch
				succeeded, committed, errors, transactor := xdefiContext.commitVaultEvents(
					sentinel,
					batch,
				)
				LogInfo(
					"[PendingVaultEvents]\t Before move vault watermark, "+
						"batch size: %d, succeeded: %v, commit %v, errors: %v",
					len(batch.Batch),
					succeeded,
					committed,
					errors,
				)
				// something's seriously wrong, we don't even get transactor
				if transactor == nil {
					continue
				}

				// wait for confirmation
				if err := xdefiContext.confirmVaultEvents(
					sentinel, newConfigChan, batch, succeeded, transactor); err == nil {
					// notify all feed subscribers
					xdefiContext.CommittedVaultEventsBatchFeed.Send(batch)
				}
			}
		}
	}
}

func (xdefiContext *XdefiContext) confirmVaultEvents(
	sentinel *Sentinel,
	newConfigChan chan uint64,
	batch *VaultEventsBatch,
	succeeded []uint64,
	transactor *bind.TransactOpts,
) error {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()

	// wait for the confirmation
	blockNumberBefore, err := xdefiContext.ClientXevents.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	for {
		// if we have new config event, don't wait for confirmation
		if len(newConfigChan) > 0 {
			newConfigVersion := <-newConfigChan
			LogInfo("[PendingVaultEvents] new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				// put it back to chan so parent func will check and exit
				newConfigChan <- newConfigVersion
				return nil
			}
		}

		time.Sleep(3 * time.Second)
		blockNumberAfter, err := xdefiContext.ClientXevents.BlockNumber(context.Background())
		if err != nil {
			return err
		}

		// wait till one block is mined and all txs in the batch confirmed
		if blockNumberAfter > blockNumberBefore && len(succeeded) == len(batch.Batch) {
			xevents := xdefiContext.Xevents()
			_, err := xevents.UpdateVaultWatermark(
				transactor,
				batch.Vault,
				big.NewInt(int64(batch.End)),
			)
			LogInfo("[PendingVaultEvents]\t MOVE vault watermark HIGHER to %d, err: %v",
				batch.End, err,
			)
			return nil
		}

		// wait time out
		if blockNumberAfter-blockNumberBefore > BatchCommitThreshold {
			LogErr(
				"[PendingVaultEvents]\t batch commit wait threshold time out: xchain block [%d -> %d]",
				blockNumberBefore, blockNumberAfter,
			)
			break
		}
	}

	return fmt.Errorf("time out waiting for vault watermark commit")
}

func (xdefiContext *XdefiContext) forwardVaultEvents(
	sentinel *Sentinel,
	tokenMapping TokenMapping,
	newConfigChan chan uint64,
) {
	clientTo := xdefiContext.ClientTo()
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()

	////////////////////////////////////////////////////////////////
	/////////// 1. get current watermark
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
			LogErr("[ForwardVaultEvents]\t Vault TO: watermark err: %v", err)
			return
		}
	} else {
		waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		sentinel.LogRpcStat("vault", "TokenMappingWatermark")
		if err != nil {
			LogErr("[ForwardVaultEvents]\t Vault TO: watermark err: %v", err)
			return
		}
	}

	LogInfo("[ForwardVaultEvents]\t Vault TO: watermark: %d, %s, %s",
		waterMark,
		tokenMapping.SourceToken,
		tokenMapping.MappedToken,
	)

	////////////////////////////////////////////////////////////////
	/////////// 2. commit new vault events to vault Y

	// each node will mint for n blocks, then hand it over to next node
	if !sentinel.myTurn(MyTurnSeed) {
		LogInfo("[ForwardVaultEvents]\t Vault TO: Not My Turn to Mint")
		return
	} else {
		LogInfo("[ForwardVaultEvents]\t Vault TO: My Turn to Mint")
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
		LogErr(
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
			LogErr("[ForwardVaultEvents]\t  Xevents: retrieve vault events err: %v", err)
		}
		rlp.DecodeBytes(vaultEventData.EventData, &vaultEvent)
		LogInfo(
			"[ForwardVaultEvents]\t Vault [TO] MINT/WITHDRAW tx: nonce: %d "+
				"amount: %d, tip: %d, to: %x, mappedToken: %x, txtor.GasLimit: %d, txtor.Nonce: %d",
			big.NewInt(waterMark.Int64()+i),
			vaultEvent.Amount,
			vaultEvent.Tip,
			vaultEvent.To.Bytes()[:8],
			vaultEvent.MappedToken.Bytes()[:8],
			transactor.GasLimit,
			transactor.Nonce,
		)
		// if no more vault event
		if vaultEvent.Amount == nil || vaultEvent.Amount.Int64() == 0 {
			break
		} else {
			var tx *types.Transaction
			var err error
			if xdefiContext.IsDeposit() {
				tx, err = xdefiContext.VaultY.Mint(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				sentinel.LogRpcStat("vault", "Mint")

			} else {
				tx, err = xdefiContext.VaultX.Withdraw(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				sentinel.LogRpcStat("vault", "Withdraw")
			}

			if err != nil {
				errors += 1
				log.Errorf("[ForwardVaultEvents]\t sentinel vault TO 'mint' tx err: %v", err)
			} else {
				LogInfo(
					"[ForwardVaultEvents]\t sentinel vault [TO] mint() tx: %x, "+
						"from: %x, nonce: %d, %d, %d       ",
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

	//////////////////////////////////////////////////////////////
	///////// 3. wait for confirmation from vault Y
	waterMarkBefore := waterMark.Uint64()
	// wait for watermark to update
	wait := 0
	for {
		if committed == 0 {
			LogInfo("[ForwardVaultEvents]\t No vault TO mint/withdraw tx called")
			break
		}
		var waterMark *big.Int
		if xdefiContext.IsDeposit() {
			waterMark, err = xdefiContext.VaultY.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
		} else {
			waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
		}
		sentinel.LogRpcStat("vault", "TokenMappingWatermark")
		LogInfo(
			"[ForwardVaultEvents]\t Vault TO: watermark: %d, before: %d, committed: %d",
			waterMark, waterMarkBefore, committed,
		)
		if err != nil {
			log.Errorf("[ForwardVaultEvents]\t Err with calling vault Y watermark: %v", err)
			continue
		}

		// 3 break conditions
		if waterMark.Uint64() == waterMarkBefore+committed {
			break
		}
		time.Sleep(5 * time.Second)
		wait += 1
		if wait > 20 {
			LogInfo(
				"[ForwardVaultEvents]\t Vault Y watermark wait time out, "+
					"watermark: %d, before: %d, committed: %d",
				waterMark, waterMarkBefore, committed,
			)
			break
		}
		if len(newConfigChan) > 0 {
			newConfigVersion := <-newConfigChan
			LogInfo("[ForwardVaultEvents] new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				// put it back to chan so parent func will check and exit
				newConfigChan <- newConfigVersion
				break
			}
		}
	}
}

func (xdefiContext *XdefiContext) ForwardVaultEvents(
	sentinel *Sentinel, tokenMapping TokenMapping,
) {
	LogInfo := xdefiContext.LogInfo()
	LogInfo("------------------enter forward vault events ------------------")
	worker := "ForwardVaultEvents"
	sentinel.addWorker(worker)
	defer LogInfo("--------------------end foward Vault Events loop --------------------")
	defer sentinel.reduceWorker(worker)

	ticker := time.NewTicker(VaultCheckInterval * time.Second)
	newConfigChan := make(chan uint64, 10)
	sub := sentinel.configVersionFeed.Subscribe(newConfigChan)
	defer sub.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case <-sub.Err():
			break
		case newConfigVersion := <-newConfigChan:
			LogInfo("[forwardVault] new config version received: %d -> %d",
				xdefiContext.ConfigVersion, newConfigVersion,
			)
			if newConfigVersion > xdefiContext.ConfigVersion {
				return
			}
		case <-ticker.C:
			xdefiContext.forwardVaultEvents(
				sentinel,
				tokenMapping,
				newConfigChan,
			)
		}
	}
}

func (xdefiContext *XdefiContext) getLastBlock(sentinel *Sentinel) uint64 {
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xevents.VaultWatermark(callOpts, vaultAddrFrom)
	sentinel.LogRpcStat("xevents", "VaultWatermark")
	if err != nil {
		log.Errorf("[getLastBlock]\t client watermark block err: %v -------", err)
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
				"[getLastBlock]\t client CreateAt err: %v, set lastBlock to %d",
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
			log.Errorf("[getLastBlock]\t client vault store counter error %v      ", err)
			return 0
		}
	}

	// share with other go routine
	sentinel.batchNumber = lastBlock

	return lastBlock
}

func (xdefiContext *XdefiContext) commitVaultEvents(
	sentinel *Sentinel,
	batch *VaultEventsBatch,
) ([]uint64, []uint64, []uint64, *bind.TransactOpts) {
	clientXevents := xdefiContext.ClientXevents
	xeventsCnt := xdefiContext.Xevents()
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()

	LogInfo(
		"[commitVaultEvents]\t Vault FROM commit batch BEGIN from %d to %d ]",
		batch.Start, batch.End,
	)
	defer LogInfo(
		"[commitVaultEvents]\t Vault FROM commit batch END from %d to %d",
		batch.Start, batch.End,
	)

	balance, err := clientXevents.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	nonceAt, err := clientXevents.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	LogInfo(
		"[commitVaultEvents]\t %x, balance: %d, account nonce: %d",
		sentinel.key.Address.Bytes(), balance, nonceAt,
	)
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		return []uint64{}, []uint64{}, []uint64{}, nil
	}

	vaultWatermark, err := xeventsCnt.VaultWatermark(
		&bind.CallOpts{}, XeventsXYAddr,
	)
	if err != nil {
		LogErr("[commitVaultEvents]\t client watermark block err: %v", err)
		return []uint64{}, []uint64{}, []uint64{}, nil
	}
	sentinel.LogRpcStat("xevents", "VaultWatermark")

	// prevent re-entry
	if vaultWatermark.Uint64() >= batch.Start {
		LogErr(
			"[commitVaultEvents]\t Vault FROM skip batch commit: vault: %d, batch number: %d  ",
			vaultWatermark.Uint64(), batch.Start,
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
		LogInfo(
			"[commitVaultEvents]\t Vault event: vault: %x, tokenMapping: %s, sha256: %x, nonce: %d",
			vaultEvent.Vault,
			vaultEvent.TokenMapping(),
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
	}
	// nonces is a string list of |vault|tokenmapping sha|nonce|
	sort.Strings(nonces)

	var vaultEventHash common.Hash
	LogInfo("[commitVaultEvents]\t Vault FROM nonces: %v", nonces)

	committed := []uint64{}
	errors := []uint64{}
	succeeded := []uint64{}
	xeventsVaultEventExts := []xevents.XEventsVaultEventExt{}
	// commit all events in order
	for _, nonce := range nonces {
		vaultEventHash = noncesMap[nonce]
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event

		// if tokenmapping nonce is low, mark as succeeded
		tokenMappingWatermark, err := xeventsCnt.TokenMappingWatermark(
			&bind.CallOpts{}, vaultEvent.Vault, vaultEvent.TokenMappingSha256(),
		)
		sentinel.LogRpcStat("xevents", "VaultEventWatermark")
		if err != nil {
			errors = append(errors, vaultEvent.Nonce.Uint64())
			LogErr("[commitVaultEvents]\t Vault FROM watermark err: %v", err)
			continue
		} else {
			LogInfo(
				"[commitVaultEvents]\t xevents tokenmapping watermark %d, this tokenmapping nonce: %d",
				tokenMappingWatermark, vaultEvent.Nonce,
			)
		}
		if tokenMappingWatermark.Uint64() > vaultEvent.Nonce.Uint64() {
			succeeded = append(succeeded, vaultEvent.Nonce.Uint64())
			LogInfo(
				"[commitVaultEvents]\t xevents tokenmapping nonce low, nonce mark succeeded: %d",
				vaultEvent.Nonce.Uint64(),
			)
			continue
		}

		xeventsVaultEventExts = append(xeventsVaultEventExts, xevents.XEventsVaultEventExt{
			vaultEventWithSig.Blssig,
			vaultEvent.Vault,
			vaultEvent.Nonce,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Bytes(),
		})
	}

	// if we actually need to commit vault events
	if len(xeventsVaultEventExts) > 0 {
		allNonces := []uint64{}
		for _, event := range xeventsVaultEventExts {
			allNonces = append(allNonces, event.Nonce.Uint64())
		}
		LogInfo("[commitVaultEvents]\t record vault event batch: %v", allNonces)
		tx, err := xeventsCnt.RecordVaultEventBatch(
			transactor, xeventsVaultEventExts,
		)
		sentinel.LogRpcStat("xevents", "Store Batch")
		//transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)
		if err != nil {
			LogErr("[commitVaultEvents]\t Vault FROM sentinel xevents store err: "+
				"%v, nonce: %d, gaslimit: %d, gasPrice: %d",
				err, transactor.Nonce, transactor.GasLimit, transactor.GasPrice)
			for _, event := range xeventsVaultEventExts {
				errors = append(errors, event.Nonce.Uint64())
			}
		} else {
			LogInfo(
				"[commitVaultEvents]\t Vault FROM sentinel xevents store tx: "+
					"%x, from: %x, nonce: %d",
				tx.Hash(), sentinel.key.Address.Bytes(), tx.Nonce(),
			)
			for _, event := range xeventsVaultEventExts {
				committed = append(committed, event.Nonce.Uint64())
			}
		}
	}

	return succeeded, committed, errors, transactor
}

///////////////////////////////////////////////////////////////////////
/////////////////////  web socket watcher        //////////////////////
///////////////////////////////////////////////////////////////////////

func (xdefiContext *XdefiContext) WatchVaultXLogs(sentinel *Sentinel) {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	defer LogInfo("Watch vault X logs exit")

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
		LogErr("Can not watch vault X logs: %v", err)
		return
	}

	// new config version
	newConfigChan := make(chan uint64, 10)
	subConfigVersion := sentinel.configVersionFeed.Subscribe(newConfigChan)
	defer subConfigVersion.Unsubscribe()
	defer close(newConfigChan)

	for {
		select {
		case <-subConfigVersion.Err():
			break
		case newConfigVersion := <-newConfigChan:
			if newConfigVersion > xdefiContext.ConfigVersion {
				break
			}
		case event := <-depositChan:
			LogInfo("notified with NEW DEPOSIT: vault: %x, nonce: %d, block: %d, source: %x, mapped: %x",
				event.Vault.Bytes()[:4],
				event.Nonce,
				event.BlockNumber,
				event.SourceToken.Bytes()[:4],
				event.MappedToken.Bytes()[:4],
			)
		case err := <-subX.Err():
			LogErr("Can not watch vault X logs, sub err: %v", err)
			break
		}
	}
}

func (xdefiContext *XdefiContext) WatchVaultYLogs(sentinel *Sentinel) {
	LogInfo := xdefiContext.LogInfo()
	LogErr := xdefiContext.LogErr()
	defer LogInfo("Watch vault Y logs exit")

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
	defer subConfigVersion.Unsubscribe()
	defer close(newConfigChan)

	if err != nil {
		LogErr("can not watch vault Y logs: %v", err)
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
			LogInfo("notified with NEW BURN: vault: %x, nonce: %d, block: %d, source: %x, mapped: %x",
				event.Vault.Bytes()[:4],
				event.Nonce,
				event.BlockNumber,
				event.SourceToken.Bytes()[:4],
				event.MappedToken.Bytes()[:4],
			)
		case err := <-subY.Err():
			LogErr("can not watch vault Y logs, sub err: %v", err)
			break
		}
	}
}
