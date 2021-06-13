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
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"time"

	blsmlib "github.com/innowells/bls_lib/v2"
	gocache "github.com/patrickmn/go-cache"
	kyber "go.dedis.ch/kyber/v3"
	ed25519 "go.dedis.ch/kyber/v3/group/edwards25519"
	bn256 "go.dedis.ch/kyber/v3/pairing/bn256"
	share "go.dedis.ch/kyber/v3/share"
	blslib "go.dedis.ch/kyber/v3/sign/bls"
	eddsa "go.dedis.ch/kyber/v3/sign/eddsa"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/xchain/accounts/keystore"
	"github.com/MOACChain/xchain/mcclient/xdefi"
	"github.com/MOACChain/xchain/params"
)

type RevealedShare struct {
	PubShare      []byte
	PubSig        []byte
	PriShare      []byte
	PriSig        []byte
	Revealed      []byte
	Violator      common.Address
	Whistleblower common.Address
}

const (
	VssConfigVote    = 0
	VssConfigReveal  = 1
	VssConfigRecheck = 2
	VssConfigOmit    = 3

	MaxConfigUpdateTimeoutInBlocks  = 10
	MaxConfigUploadTimeoutInSeconds = 120 * time.Second
	GetVSSSharesTimeout             = 60 * time.Second

	SlashReveal        = 0
	CounterSlashReveal = 1
	NoSlashAction      = 2
)

var (
	SlowNodeChan        = make(chan uint64, 10)
	RevealedShareChan   = make(chan uint64, 10)
	UploadVSSConfigChan = make(chan uint64, 10)
	VssStateChan        = make(chan uint64, 10)
)

func (pos *Pos) IsVSSReady() bool {
	ret := pos.vssEnabled && pos.Bls != nil && pos.Bls.GroupPrivateShare != nil
	log.Infof("--------------------------- IsVSSReady: %t", ret)
	return ret
}

func (pos *Pos) ShouldSkipVSS(blockNumber int64) bool {
	return !pos.vssEnabled || blockNumber < int64(params.MinBLSBlockNumber)
}

func (pos *Pos) GetVssThreshold() uint64 {
	n, _ := pos.vssbase.VssThreshold(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetVssBaseContractAddress() common.Address {
	return common.HexToAddress("0x0000")
}

func (pos *Pos) GetVSSNodeIndex() *big.Int {
	n, _ := pos.vssbase.GetVSSNodeIndex(pos.callOpts, pos.Vssid)
	return n
}

func (pos *Pos) isVSSConfigReady(configVersion int) (bool, error) {
	b, err := pos.vssbase.IsConfigReady(pos.callOpts, big.NewInt(int64(configVersion)))
	return b, err
}

func (pos *Pos) GetVSSNodesPubkey(
	activeNodeList []common.Address,
) (map[common.Address][]byte, error) {
	result := make(map[common.Address][]byte)
	emptyPubkey := 0
	publickeys, _ := pos.vssbase.GetVSSNodesPubkey(pos.callOpts, activeNodeList)
	for i, addr := range activeNodeList {
		result[addr] = publickeys[i][:]
		log.Debugf("pub key for addr: %x is %x", addr.Bytes(), publickeys[i])
		if len(publickeys[i]) == 32 {
			emptyPubkey++
		}
	}

	if emptyPubkey > 0 {
		return result, fmt.Errorf("empty pub key in for active node.")
	} else {
		return result, nil
	}
}

func (pos *Pos) IsActiveVss() bool {
	// noreg = 0, active = 1, inactive = 2
	ret, _ := pos.vssbase.VssNodeMemberships(pos.callOpts, pos.Vssid)
	return ret.Int64() == 1
}

func (pos *Pos) GetVSSNodesIndexs(activeMemberList []common.Address) map[common.Address]int {
	result := make(map[common.Address]int)
	indexs, _ := pos.vssbase.GetVSSNodesIndexs(pos.callOpts, activeMemberList)
	for i, addr := range activeMemberList {
		result[addr] = int(indexs[i].Int64())
	}
	return result
}

func (pos *Pos) GetActiveVSSMemberList() ([]common.Address, error) {
	activeMemberList, err := pos.vssbase.GetActiveVSSMemberList(pos.callOpts)
	return activeMemberList, err
}

func (pos *Pos) GetLastSlashVoted() uint64 {
	n, _ := pos.vssbase.GetLastSlashVoted(pos.callOpts, pos.Vssid)
	return n.Uint64()
}

func (pos *Pos) GetLastConfigUpload(address common.Address) uint64 {
	n, _ := pos.vssbase.LastConfigUpload(pos.callOpts, address)
	return n.Uint64()
}

func (pos *Pos) GetVssNodeCount() uint64 {
	n, _ := pos.vssbase.VssNodeCount(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetVssConfigVersion() uint64 {
	n, _ := pos.vssbase.VssConfigVersion(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetRevealIndex() uint64 {
	n, _ := pos.vssbase.RevealIndex(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetLastNodeChangeConfigVersion() uint64 {
	n, _ := pos.vssbase.LastNodeChangeConfigVersion(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetLastNodeChangeBlock() uint64 {
	n, _ := pos.vssbase.LastNodeChangeBlock(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetSlowNodeThreshold() uint64 {
	n, _ := pos.vssbase.SlowNodeThreshold(pos.callOpts)
	return n.Uint64()
}

func (pos *Pos) GetVSSShares(
	address common.Address,
	endpoint string,
	activeNodeList []common.Address,
) (map[common.Address][]byte, error) {
	receivedVSSShares := make(map[common.Address][]byte)

	// wait up to n seconds
	timeout := time.Now().Add(GetVSSSharesTimeout)

	// query vss shares for addresses in active node list in a loop
	// sleep for 1 second each round and timout after n seconds
	for {
		for _, addr := range activeNodeList {
			if _, found := receivedVSSShares[addr]; found {
				continue
			} else {
				var _vssShares []byte
				var err error
				if endpoint == "getPublicShares" {
					_vssShares, err = pos.vssbase.GetPublicShares(pos.callOpts, addr)
				} else {
					_vssShares, err = pos.vssbase.GetPrivateShares(pos.callOpts, addr)
				}

				if len(_vssShares) > 1 {
					// ungzip the bytes
					bufzip := bytes.NewBuffer(_vssShares)
					bufunzip := new(bytes.Buffer)
					zr, _ := gzip.NewReader(bufzip)
					io.Copy(bufunzip, zr)
					vssShares := bufunzip.Bytes()
					receivedVSSShares[addr] = vssShares
					log.Debugf(
						"Received (%v/%v) %v shares: %v, err: %v",
						len(receivedVSSShares),
						len(activeNodeList),
						endpoint,
						string(vssShares[:]),
						err,
					)
				}
			}
		}
		if len(receivedVSSShares) == len(activeNodeList) {
			// if we received all the active vss shares, break the outer loop
			break
		}

		if time.Now().After(timeout) {
			missingShares := make([]string, 0)
			for _, addr := range activeNodeList {
				if _, found := receivedVSSShares[addr]; !found {
					missingShares = append(missingShares, addr.String())
				}
			}
			return receivedVSSShares, fmt.Errorf(
				"get vss shares timed out, missing shares from addresses: %v",
				missingShares,
			)
		} else {
			// otherwise sleep for 1 second and retry
			time.Sleep(1 * time.Second)
		}
	}

	return receivedVSSShares, nil
}

func (pos *Pos) GetBlsSig(sigShareKey string) ([]byte, error) {
	timeoutInterval := params.BlockInterval
	log.Debugf(
		"vss call GetBlsSig() with key: %v, allsigs count: %d",
		sigShareKey,
		pos.AllSigs.ItemCount(),
	)
	// check every 1s and timeout after 30s
	for i := uint64(0); i < timeoutInterval; i++ {
		if value, ok := pos.AllSigs.Get(sigShareKey); ok {
			// find the sig sha
			sig := value.([]byte)
			return sig, nil
		}
		time.Sleep(1 * time.Second)
	}

	log.Debugf("Wait for vss result timeout")
	return []byte{}, fmt.Errorf("wait for vss result timeout")
}

func (pos *Pos) ValidateProposerSig(
	proposerSig []byte, proposer common.Address, blsSig []byte, fixBlock bool, blockNumber int64,
) bool {
	if fixBlock {
		return true
	}

	// for the first n blocks of the chain, we don't validate bls
	if pos.ShouldSkipVSS(blockNumber) {
		return true
	}

	if !pos.IsVSSReady() {
		log.Debugf("validate proposer sig failed, bls is not ready")
		return false
	}

	pubkeyBytes, found := pos.nodesPubkey[proposer]
	if found {
		suite25519 := ed25519.NewBlakeSHA256Ed25519()
		pubkey := suite25519.Point()
		pubkey.UnmarshalFrom(bytes.NewReader(pubkeyBytes))
		validated, err := VerifyED25519(pubkey, blsSig, proposerSig)
		if validated {
			log.Debugf(
				"validate proposer sig finished successfully, proposer sig: %x, bls sig: %x, pubkey: %x",
				proposerSig,
				blsSig,
				pubkeyBytes,
			)
			return true
		} else {
			log.Debugf(
				"validate proposer sig failed, err: %v, proposer sig: %x, bls sig: %x, pubkey: %x",
				err,
				proposerSig,
				blsSig,
				pubkeyBytes,
			)
		}
	}

	return false
}

func (pos *Pos) ValidateBLSSig(
	toSign, blsSig []byte, fixBlock bool, blockNumber int64,
) bool {
	// 1. monitor don't participate in bls sig
	// 2. when sync, don't validate bls sig since it's likely
	// we don't have the past bls curve for legacy blocks
	if fixBlock {
		return true
	}

	// for the first n blocks of the chain, we don't validate bls
	if pos.ShouldSkipVSS(blockNumber) {
		return true
	}

	if !pos.IsVSSReady() {
		log.Debugf("validate bls sig failed, bls is not ready")
		return false
	}

	suite := bn256.NewSuite()
	errVerify := blslib.Verify(
		suite,
		pos.Bls.GroupPublicPoly.Commit(),
		toSign,
		blsSig,
	)
	if errVerify != nil {
		log.Debugf(
			"vss validate bls sig failed, sig not verify with current bls, h = %x, s = %x",
			toSign,
			blsSig,
		)
	} else {
		log.Debugf("vss validate bls sig ok")
		return true
	}

	// loop through previous bls to see if the sig verify
	bls := pos.prevbls.Front()
	for {
		if bls == nil {
			break
		} else {
			_bls := bls.Value.(*BLS)
			errVerify := blslib.Verify(
				suite, _bls.GroupPublicPoly.Commit(),
				toSign,
				blsSig,
			)
			if errVerify != nil {
				log.Debugf(
					"vss validate bls sig failed, sig not verify with prev bls, h = %x, s = %x",
					toSign,
					blsSig,
				)
			} else {
				log.Debugf(
					"vss validate bls sig ok with prev bls, h = %x, s = %x",
					toSign,
					blsSig,
				)
				return true
			}
			bls = bls.Next()
		}
	}

	return false
}

func (pos *Pos) HandleSigShares() {
	// blockhash -> nodexindex -> sigshare
	var allSigShares = make(map[string]map[int][]byte)

	for {
		sigShare := <-pos.Vss.SigShareChan
		sigShareKey := sigShare.Key()

		if _, ok := allSigShares[sigShareKey]; !ok {
			allSigShares[sigShareKey] = make(map[int][]byte)
		}
		allSigShares[sigShareKey][sigShare.FromIndex] = sigShare.Sig

		/*
			// update each node's height
			from := pos.IndexToAddress[sigShare.FromIndex]
			// block number in sigshare is next block number
			sci.RecordNodeBlockHeight(
				from,
				int(sigShare.BlockNumber)-1,
				sigShare.BlockHash,
				int(sigShare.Epoch),
			)*/

		// Try generate randome number if we have enough sig collected
		if pos.IsVSSReady() {
			sigs := allSigShares[sigShareKey]

			// skip if there is not enough signature
			if len(sigs) < pos.Bls.Threshold {
				continue
			}

			vssResult, err := pos.Bls.GenerateRandomNumber(
				[]byte(sigShareKey),
				sigs,
			)
			if err == nil {
				log.Debugf(
					"vss Sigs verified: %s, %x",
					string(vssResult.ToSign),
					vssResult.Sig,
				)
				pos.AllSigs.Set(
					sigShareKey,
					vssResult.Sig,
					gocache.DefaultExpiration,
				)

				delete(allSigShares, sigShareKey)
			} else {
				log.Errorf("vss sigs with error: %v", err)
			}
		}
	}
}

func (pos *Pos) LoadVSSKey() {
	vsskeyBytes := keystore.GetVSSKey(common.HexToAddress(pos.vnodeconfig.VssBaseAddr))
	vsskey := &VSSKey{}
	if len(vsskeyBytes) == 0 {
		vsskey = GenerateVSSKey()

		// store the vss key in db
		pVSSKey := ToPersistVSSKey(vsskey)
		data, _ := json.Marshal(pVSSKey)
		log.Debugf("vss can not find vss key & generated new one: %s", string(data))
		if err := keystore.PutVSSKey(
			common.HexToAddress(pos.vnodeconfig.VssBaseAddr), data,
		); err != nil {
			log.Errorf("vss key can not be stored: %v", err)
		}

	} else {
		var pVSSKey PersistVSSKey
		json.Unmarshal(vsskeyBytes, &pVSSKey)
		suite25519 := ed25519.NewBlakeSHA256Ed25519()
		publickey := suite25519.Point()
		publickey.UnmarshalFrom(bytes.NewReader(pVSSKey.Public))
		privatekey := suite25519.Scalar()
		privatekey.UnmarshalFrom(bytes.NewReader(pVSSKey.Private))
		vsskey.Public = publickey
		vsskey.Private = privatekey
		log.Debugf("loaded vss key from db: %s", string(vsskeyBytes))
	}

	pos.VssKey = vsskey
	log.Debugf("load vsskey as: %v", vsskey)
}

func (pos *Pos) NewVnodeBlockLoop() {
	// this is 5 seconds if block interval is 10 seconds
	interval := params.BlockInterval / 2
	if interval < params.MinVSSRunInterval {
		interval = params.MinVSSRunInterval
	}
	log.Debugf("New vnode block loop, run every %d seconds", interval)
	t := time.NewTicker(time.Duration(int(interval)) * time.Second)
	defer t.Stop()
	defer log.Infof("new vnode block loop exits")
	lastBlock := uint64(0)
	for {
		<-t.C
		currentVnodeBlockNumber, _ := pos.client.BlockNumber(context.Background())
		log.Debugf("new vnode block loop block number = %d", currentVnodeBlockNumber)
		if currentVnodeBlockNumber > lastBlock {
			SlowNodeChan <- currentVnodeBlockNumber
			RevealedShareChan <- currentVnodeBlockNumber
			UploadVSSConfigChan <- currentVnodeBlockNumber
			VssStateChan <- currentVnodeBlockNumber

			// update last block number
			lastBlock = currentVnodeBlockNumber
		}
	}
}

// runs indefinitely and update vss config and update vss config if needed
func (pos *Pos) VssStateLoop() {
	log.Debugf("vss state loop, runs on new block")
	defer log.Infof("vss state loop exits")
	for {
		<-VssStateChan
		// download secret shares
		pos.RunVssStateMachine()
	}
}

// runs indefinitely and upload vss config if needed.
func (pos *Pos) VssUploadConfigLoop() {
	log.Debugf("vss upload config loop, runs on new block")
	defer log.Infof("vss upload config loop exit")
	for {
		<-UploadVSSConfigChan
		// upload secret shares
		log.Debugf("call UploadVSSConfig from VssStateLoop()")
		pos.UploadVSSConfig(false)
	}
}

// runs indefinitely and participate slashing
func (pos *Pos) VssSlashingLoop() {
	// there is time discrepancy between block write and read
	// so we need to keep track of slash status in a local map too
	// mapping: slash index -> bool
	slashed := make(map[uint64]bool)
	log.Debugf("vss slashing loop, runs on new block")
	defer log.Infof("vss slashing loop exit")

	for {
		select {
		case <-SlowNodeChan:
			activeNodeList, _ := pos.GetActiveVSSMemberList()
			for _, addr := range activeNodeList {
				lastConfigUpload := pos.GetLastConfigUpload(addr)
				lastNodeChangeConfigVersion := pos.GetLastNodeChangeConfigVersion()
				lastNodeChangeBlock := pos.GetLastNodeChangeBlock()
				threshold := pos.GetSlowNodeThreshold()
				currentVnodeBlockNumber, _ := pos.client.BlockNumber(context.Background())
				// if node is behind node change config version and should upload new vss config
				log.Debugf(
					"vss report slow node[%x]: last config upload: %d, last node change: %d, block: %d, last node change block: %d, threshold: %d",
					addr, lastConfigUpload, lastNodeChangeConfigVersion, currentVnodeBlockNumber, lastNodeChangeBlock, threshold,
				)

				if lastConfigUpload < lastNodeChangeConfigVersion && currentVnodeBlockNumber-lastNodeChangeBlock > threshold {
					log.Debugf("vss report slow node exceed threshold: addr %x", addr)
					pos.vssReportSlowNode(addr)
				}
			}
		case <-RevealedShareChan:
			lastSlashVoted := pos.GetLastSlashVoted()
			contractRevealIndex := pos.GetRevealIndex()
			log.Infof(
				"-------------------------------  vss slashing loop, lastSlashVoted: %d, contractRevealIndex: %d",
				lastSlashVoted,
				contractRevealIndex,
			)

			// if there are revealed shares we haven't slashed yet
			lastRevealedIndex := contractRevealIndex - 1
			if lastSlashVoted < lastRevealedIndex {
				for index := lastSlashVoted; index <= lastRevealedIndex; index++ {
					// don't slash twice
					if slashed[index] {
						continue
					}
					log.Debugf(
						"vss slashing loop, index: %d, lastSlashVoted: %d, lastRevealedIndex: %d, contractRevealIndex: %d",
						index,
						lastSlashVoted,
						lastRevealedIndex,
						contractRevealIndex,
					)
					if revealedShare, err := pos.vssbase.GetRevealedShare(
						pos.callOpts, big.NewInt(int64(index))); err == nil {
						log.Debugf(
							"vss get revealed share pubshare: %s, pubsig: %x, prishare: %s, prisig: %x, revealed: %x, violator: %x, Whistleblower: %x",
							string(revealedShare.PubShare[:]),
							revealedShare.PubSig,
							string(revealedShare.PriShare[:]),
							revealedShare.PriSig,
							revealedShare.Revealed,
							revealedShare.Violator,
							revealedShare.Whistleblower,
						)

						// determine if we should slash
						violator, slash, err := pos.verifyRevealedShare(revealedShare)
						log.Debugf(
							"vss verifyRevealedShare violator: %x, slash: %d, err: %v",
							violator, slash, err,
						)

						// either slash or counter slash
						if slash == SlashReveal {
							//pos.vssSlashing(index, true)
							slashed[index] = true
						} else if slash == CounterSlashReveal {
							//pos.vssSlashing(index, false)
							slashed[index] = true
						}
					} else {
						log.Debugf("vss get revealed share err: %v", err)
					}
				}
			}
		}
	}
}

func (pos *Pos) verifyRevealedShare(revealedShare xdefi.VssBaseRevealedShare) (
	common.Address, int, error) {
	var priShares []PersistPrivateShare
	var pubShares []PersistPublicShare

	if pos.Bls == nil {
		return revealedShare.Violator, NoSlashAction, fmt.Errorf("vss verifyRevealedShare pos.Bls is nil")
	}

	// get violator's pubkey
	pubkeyBytes, found := pos.nodesPubkey[revealedShare.Violator]
	if !found {
		log.Debugf("vss verifyRevealedShare can not find pubkey for %x", revealedShare.Violator)
		return revealedShare.Violator, NoSlashAction, fmt.Errorf("vss verifyRevealedShare can not find pubkey for %x", revealedShare.Violator)
	}

	// use the pubkey to check sig so we are sure these shares are sent by the violator
	suite25519 := ed25519.NewBlakeSHA256Ed25519()
	pubkey := suite25519.Point()
	pubkey.UnmarshalFrom(bytes.NewReader(pubkeyBytes))
	verifiedPub, errPub := VerifyED25519(pubkey, revealedShare.PubShare, revealedShare.PubSig)
	verifiedPri, errPri := VerifyED25519(pubkey, revealedShare.PriShare, revealedShare.PriSig)
	if !verifiedPub || !verifiedPri || errPub != nil || errPri != nil {
		return revealedShare.Violator, CounterSlashReveal, fmt.Errorf("vss verifyRevealedShare sig does not verify for %x", revealedShare.Violator)
	}

	json.Unmarshal(revealedShare.PubShare, &pubShares)
	json.Unmarshal(revealedShare.PriShare, &priShares)

	activeNodeList, _ := pos.GetActiveVSSMemberList()
	activeNodeAddressAndIndex := pos.GetVSSNodesIndexs(
		activeNodeList,
	)
	index := activeNodeAddressAndIndex[revealedShare.Whistleblower]
	suite := bn256.NewSuite()
	for _, pubshare := range pubShares {
		for _, prishare := range priShares {
			if pubshare.I == index && prishare.I == index {
				pub := suite.G2().Point()
				pub.UnmarshalFrom(bytes.NewReader([]byte(pubshare.BV)))
				pri := suite.G1().Scalar()
				err := pri.UnmarshalBinary(revealedShare.Revealed)
				if err != nil {
					log.Errorf("verifyRevealedShare UnmarshalBinary() error %v", err)
				}
				verifyPub := suite.G2().Point().Mul(pri, nil)
				if verifyPub.Equal(pub) {
					log.Debugf(
						"vss verifyRevealedShare share verified for sender: %x",
						revealedShare.Violator,
					)
					break
				} else {
					return revealedShare.Violator, SlashReveal, nil
				}
			}
		}
	}

	return revealedShare.Violator, NoSlashAction, nil
}

func (pos *Pos) RunVssStateMachine() {
	// always check the latest config version so that eventually everyone will
	// converge to the same state given that config version change is rare (add new node
	// or delete old node).
	vssConfigVersion := pos.GetVssConfigVersion()
	if seen := pos.vssSeenConfigs[int(vssConfigVersion)]; !seen {
		// if not found, the it is a new state
		// we should call updateVssConfig() to get the latest vss setting
		updateResult := pos.UpdateVSSConfig()
		log.Debugf(
			"vss first seen config version: %d, call updateVssConfig() = %d",
			vssConfigVersion,
			updateResult,
		)
		// if we don't need to recheck, mark this as seen
		if updateResult != VssConfigRecheck {
			pos.vssSeenConfigs[int(vssConfigVersion)] = true
		}
	} else {
		// we have seen this vss config before, so we either voted or revealed
		if ret, _ := pos.isVSSConfigReady(int(vssConfigVersion)); ret {
			log.Debugf("vss switch to new config version %d", vssConfigVersion)
			// switch to new config
			pos.setVssConfig(int(vssConfigVersion))
			return
		}
	}
}

// read the current vss settings from subchainbase contract, return 3 status:
// 1: vote, means all configs are correct
// 2: slash, means at least one of the configs are incorrect
// 3: recheck, means it needs to recheck this config again
func (pos *Pos) UpdateVSSConfig() int {
	//singleton of UpdateVSSConfig
	pos.vssIsrunningMutex.Lock()
	defer pos.vssIsrunningMutex.Unlock()

	log.Debugf("vss updateVssConfig")
	vssConfigVersion := pos.GetVssConfigVersion()
	activeNodeList, listErr := pos.GetActiveVSSMemberList()
	addressToIndex := pos.GetVSSNodesIndexs(activeNodeList)
	pos.AddressToIndex = make(map[string]int)
	pos.IndexToAddress = make(map[int]string)
	for address, index := range addressToIndex {
		pos.AddressToIndex[address.String()] = index
		pos.IndexToAddress[index] = address.String()
	}

	if listErr != nil {
		log.Debugf("skip reload vss, return error: %v", listErr)
		return VssConfigRecheck
	}
	// if no change in config, we don't need to reload vss
	if !pos.shouldUpdateVSSConfig(int(vssConfigVersion), activeNodeList) {
		log.Debugf("skip reload vss, return early: vssConfigVersion %d", vssConfigVersion)
		return VssConfigOmit
	}
	// get vss configs from subchainbase
	threshold := pos.GetVssThreshold()
	vssNodeIndex := pos.GetVSSNodeIndex()
	vssNodeCount := pos.GetVssNodeCount()

	// check if this node is active in the vssbase contract
	if !pos.IsNodeActive(activeNodeList) {
		return VssConfigRecheck
	}

	pos.nodesPubkey, _ = pos.GetVSSNodesPubkey(activeNodeList)
	// if vss in subchainbase is not yet finished setting up.
	if len(activeNodeList) < int(threshold) {
		log.Errorf("vss not enough active nodes")
		return VssConfigRecheck
	}

	log.Debugf(
		"vss config: threshold = %d, config_version: %d, total node count: %d, active ones %d, local node index: %d",
		threshold,
		vssConfigVersion,
		vssNodeCount,
		len(activeNodeList),
		vssNodeIndex,
	)

	// initialize bls and vss data
	sciBls := NewBLS(int(threshold), activeNodeList, int(vssNodeIndex.Int64()), int(vssNodeCount))
	suite25519 := ed25519.NewBlakeSHA256Ed25519()

	// compare this to lastConfigUploadAfter to determine if
	// config changes in between
	lastConfigUploadBefore := make(map[common.Address]uint64)
	for _, addr := range activeNodeList {
		lastConfigUploadBefore[addr] = pos.GetLastConfigUpload(addr)
	}

	// get all public shares, including non-active ones
	allPublicShares := make([]*[]*share.PubShare, 0)
	verifyPublicShares := make(map[common.Address]*[]*share.PubShare)
	signedPublicShares := make(map[common.Address]PersistPublicShares)
	pubshares, err := pos.GetVSSShares(
		common.HexToAddress(pos.vnodeconfig.VssBaseAddr),
		"getPublicShares",
		activeNodeList,
	)
	if err != nil {
		log.Errorf("vss updateVssConfig() failed with %v", err)
		return VssConfigRecheck
	}
	pubShareCount := 1
	for addr, publicShares := range pubshares {
		// need the pubkey of the sender to verify the received share
		pubkeyBytes, found := pos.nodesPubkey[addr]
		if found {
			pubkey := suite25519.Point()
			pubkey.UnmarshalFrom(bytes.NewReader(pubkeyBytes))
			shares, signedShares := UnmarshalPublicShares(&publicShares, pubkey)
			allPublicShares = append(allPublicShares, shares)
			verifyPublicShares[addr] = shares
			signedPublicShares[addr] = signedShares
			log.Debugf(
				"Decoded (%d/%d) public shares: %v",
				pubShareCount,
				len(pubshares),
				shares,
			)
		}
		pubShareCount++
	}

	// get all private shares, including non-active ones
	allPrivateShares := make([]*[]*share.PriShare, 0)
	verifyPrivateShares := make(map[common.Address]*[]*share.PriShare)
	signedPrivateShares := make(map[common.Address]PersistPrivateShares)
	prishares, err := pos.GetVSSShares(
		common.HexToAddress(pos.vnodeconfig.VssBaseAddr),
		"getPrivateShares",
		activeNodeList,
	)
	if err != nil {
		log.Debugf("vss updateVssConfig() failed with %v", err)
		return VssConfigRecheck
	}

	// compare this to lastConfigUploadBefore to determine if
	// config changes in between. If so, return recheck
	var lastConfigUploadAfter uint64
	for _, addr := range activeNodeList {
		lastConfigUploadAfter = pos.GetLastConfigUpload(addr)
		if lastConfigUploadBefore[addr] != lastConfigUploadAfter {
			return VssConfigRecheck
		}
	}

	priShareCount := 1
	for addr, activePrivateShare := range prishares {
		// get the key to decrypt
		var buf bytes.Buffer
		pos.VssKey.Private.MarshalTo(&buf)
		stream := blsmlib.ConstantStream(buf.Bytes())
		edDSA := eddsa.NewEdDSA(stream)

		// need the pubkey of the sender to verify the received share
		pubkeyBytes, found := pos.nodesPubkey[addr]
		if found {
			pubkey := suite25519.Point()
			pubkey.UnmarshalFrom(bytes.NewReader(pubkeyBytes))
			shares, signedShares := UnmarshalPrivateShares(
				&activePrivateShare,
				edDSA.Secret,
				sciBls.NodeIndex,
				pubkey,
			)
			allPrivateShares = append(allPrivateShares, shares)
			verifyPrivateShares[addr] = shares
			signedPrivateShares[addr] = signedShares
			log.Debugf(
				"Decoded (%d/%d) private shares: %v",
				priShareCount,
				len(prishares),
				shares,
			)

		}
		priShareCount++
	}

	// create a new curve suite
	suite := bn256.NewSuite()

	// verify pri/pub shares before we recontruct poly
	// if verify fails, report vssbase contract
	notVerify := pos.verifyShares(sciBls, suite, verifyPublicShares, verifyPrivateShares)
	if len(notVerify) > 0 {
		// return a mapping from address to index
		activeNodeAddressAndIndex := pos.GetVSSNodesIndexs(activeNodeList)
		for addr, myPriShare := range notVerify {
			index := activeNodeAddressAndIndex[addr]
			log.Debugf("vss reveal: share not verified for index %d: %x", index, addr)
			signedPubShare := signedPublicShares[addr]
			signedPriShare := signedPrivateShares[addr]
			BV, errBV := myPriShare.V.MarshalBinary()
			if errBV != nil {
				log.Debugf("vss reveal BV: %v, errBV: %v, myPriShare: %v", BV, errBV, myPriShare)
			}
			pos.vssReveal(
				addr,
				[]byte(signedPubShare.PersistPublicShares),
				[]byte(signedPubShare.Sig),
				[]byte(signedPriShare.PersistPrivateShares),
				[]byte(signedPriShare.Sig),
				BV,
			)
		}
		log.Debugf("vss updateVssConfig() failed, unable to verify shares")
		return VssConfigReveal
	}

	// reconstruct group public poly
	// set initial poly to be the first one
	ps := *(allPublicShares[0])
	// set the empty one to be nil pointer so that RecoverPubPoly() will ignore it
	for i := 0; i < len(ps); i++ {
		if ps[i].I < 0 {
			ps[i] = nil
		}
	}
	log.Debugf("vss recover pub shares: %d/%d/%d, pubshares(%d) %v",
		sciBls.Threshold, sciBls.N, sciBls.TotalN, len(ps), ps,
	)
	pubpoly, e0 := share.RecoverPubPoly(
		suite.G2(),
		ps,
		sciBls.Threshold,
		sciBls.TotalN,
	)
	if e0 != nil || pubpoly == nil {
		log.Errorf("vss reconstruct group public poly error %v", e0)
		return VssConfigRecheck
	}

	// combine all pubpoly to get group pubpoly
	sciBls.GroupPublicPoly = pubpoly
	nIndex := 1
	for _, publicShares := range allPublicShares[1:] {
		_ps := *publicShares
		// set the empty one to be nil pointer so that RecoverPubPoly() will ignore it
		for i := 0; i < len(_ps); i++ {
			if _ps[i].I < 0 {
				_ps[i] = nil
			}
		}
		pubpoly, e := share.RecoverPubPoly(
			suite.G2(),
			_ps,
			sciBls.Threshold,
			sciBls.TotalN,
		)

		if e != nil || pubpoly == nil {
			log.Errorf("vss reconstruct group public poly in loop error %v", e)
			return VssConfigRecheck
		}
		var _e error
		sciBls.GroupPublicPoly, _e = sciBls.GroupPublicPoly.Add(pubpoly)
		if _e != nil {
			log.Errorf("vss reconstruct group public poly error %v, %v", e, _e)
			return VssConfigRecheck
		}
		nIndex++
	}

	// reconstruct secret share for this node from group private poly
	log.Debugf("vss allPrivateShares(%d): %v", len(allPrivateShares), allPrivateShares)
	groupPrivateShare := suite.G2().Scalar().Zero()
	for _, privateShares := range allPrivateShares {
		for _, prishare := range *privateShares {
			if prishare.I == sciBls.NodeIndex {
				groupPrivateShare = suite.G2().Scalar().Add(
					groupPrivateShare,
					prishare.V,
				)
			}
		}
	}
	sciBls.GroupPrivateShare = &share.PriShare{
		sciBls.NodeIndex,
		groupPrivateShare,
	}
	verifyPubShareV := suite.G2().Point().Mul(sciBls.GroupPrivateShare.V, nil)
	if sciBls.GroupPublicPoly.Eval(sciBls.NodeIndex).V.Equal(verifyPubShareV) {
		log.Debugf("vss Pub/Pri shares verify")
		// vote the config to be correct
		pos.vssVote(int(vssConfigVersion))

		// record the bls setting by config version number
		// we will later check vote result to determine
		// if we will swap the existing setting with new one
		pos.vssSettings[int(vssConfigVersion)] = sciBls
		return VssConfigVote
	} else {
		log.Errorf("vss Pub/Pri shares do NOT verify")
		return VssConfigRecheck
	}
}

func (pos *Pos) setVssConfig(vssConfigVersion int) {
	// set only if config is ready locally and the current one out-dated
	if bls, found := pos.vssSettings[vssConfigVersion]; found {
		if pos.currentConfigVersion < vssConfigVersion {
			log.Debugf(
				"vss new config version %d is ready locally, will swap with current config",
				vssConfigVersion,
			)
			// update pos.Bls list
			if pos.Bls != nil {
				pos.prevbls.PushFront(pos.Bls)
			}
			if pos.prevbls.Len() > int(params.MaximumBLSHistorySize) {
				e := pos.prevbls.Back()
				if e != nil {
					pos.prevbls.Remove(e)
				}
			}

			// swap to be current
			pos.Bls = bls
			pos.currentConfigVersion = vssConfigVersion
		}
	} else {
		log.Debugf("vss config version %d is not ready locally", vssConfigVersion)
	}
}

func (pos *Pos) IsVSSEnabled() bool {
	var NullAddress common.Address
	res := bytes.Compare(
		common.HexToAddress(pos.vnodeconfig.VssBaseAddr).Bytes()[:common.AddressLength],
		NullAddress[:common.AddressLength],
	)
	return res != 0
}

func (pos *Pos) shouldUpdateVSSConfig(
	vssConfigVersion int,
	activeNodeList []common.Address,
) bool {
	// if bls is configed and version is lower
	if pos.currentConfigVersion < vssConfigVersion {
		// if some nodes are behind due to a node change, wait up to n seconds
		// for them and don't update config immediately
		timeout := time.Now().Add(GetVSSSharesTimeout)
		for {
			allNodesUptoDate := true
			lastNodeChangeConfigVersion := pos.GetLastNodeChangeConfigVersion()
			threshold := pos.GetVssThreshold()
			nodeUptoDateCount := 0
			for _, addr := range activeNodeList {
				lastConfigUpload := pos.GetLastConfigUpload(addr)
				log.Debugf(
					"in shouldUpdateVSSConfig() node: %s, last config update: %d, last node change: %d",
					addr.String(),
					lastConfigUpload,
					lastNodeChangeConfigVersion,
				)
				// in ShouldUploadVSSConfig, if lastConfigUpload <= lastNodeChangeConfigVersion
				// we will always upload new config, so for every node,
				// its lastConfigUpload should be always > lastNodeChangeConfigVersion
				// if it's up to date
				if lastConfigUpload < lastNodeChangeConfigVersion {
					allNodesUptoDate = false
				} else {
					nodeUptoDateCount++
				}
			}
			log.Debugf(
				"in shouldUpdateVSSConfig: all nodes up to date = %t, nodes up to date: %d",
				allNodesUptoDate,
				nodeUptoDateCount,
			)
			if allNodesUptoDate {
				return true
			}

			log.Debugf(
				"in shouldUpdateVSSConfig: threshold: %d, node up to date: %d, now: %v, timeout: %v",
				threshold,
				nodeUptoDateCount,
				time.Now(),
				timeout,
			)
			if time.Now().After(timeout) {
				if nodeUptoDateCount >= int(threshold) {
					return true
				} else {
					// break to return false
					// todo: slash slow nodes
					break
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}

	return false
}

// GetVSSNodesPubkey() returns a mapping of address => public key
// for each node

// check if this node is active or not
func (pos *Pos) IsNodeActive(activeNodeList []common.Address) bool {
	isNodeActive := false
	for _, nodeAddr := range activeNodeList {
		log.Debugf("vss activeNodeList: %s, self: %s", nodeAddr.String(), pos.Vssid.String())
		if pos.Vssid == nodeAddr {
			isNodeActive = true
		}
	}

	if !isNodeActive {
		log.Debugf("vss node not active yet")
	}

	return isNodeActive
}

func (pos *Pos) ShouldUploadVSSConfig(
	activeNodeList []common.Address,
	threshold int,
	forceUpdate bool,
) (bool, error) {
	// #1 check if this node is active or not
	if !pos.IsNodeActive(activeNodeList) {
		log.Debugf("ShouldUploadVSSConfig: false, reason: vss node not active yet")
		return false, fmt.Errorf("vss node not active yet")
	}

	// #2 allow force upload config
	if forceUpdate {
		log.Debugf("ShouldUploadVSSConfig: true, reason: force update")
		return true, nil
	}

	// #3 don't upload repeatly
	if time.Now().Before(pos.uploadedConfigTime.Add(MaxConfigUploadTimeoutInSeconds)) {
		log.Debugf("ShouldUploadVSSConfig: false, reason: no repeat update")
		return false, fmt.Errorf(
			"within %v of last config upload, last: %v, now: %v",
			MaxConfigUploadTimeoutInSeconds,
			pos.uploadedConfigTime,
			time.Now(),
		)
	}

	// #4 if we never upload before
	lastConfigUpload := pos.GetLastConfigUpload(pos.Vssid)
	lastNodeChangeConfigVersion := pos.GetLastNodeChangeConfigVersion()
	if lastConfigUpload == 0 {
		log.Debugf("ShouldUploadVSSConfig: true, reason: first vss config update")
		return true, nil
	} else {
		// #5 if we uploaded config before but are behind node change
		if lastConfigUpload <= lastNodeChangeConfigVersion {
			log.Debugf("ShouldUploadVSSConfig: true, reason: behind node change")
			return true, nil
		}
	}

	// #6 false for all other cases
	log.Debugf("ShouldUploadVSSConfig: false, reason: no change")
	return false, fmt.Errorf(
		"active node list [%d] matches",
		len(activeNodeList),
	)
}

// UploadVSSConfig generates pub and pri shares by reading configs from the subchainbase contract
// and then upload them to the contract
func (pos *Pos) UploadVSSConfig(forceUpdate bool) error {
	// send request to contract at contractAddress.
	// return a list of active vss nodes addresses (address[])
	activeNodeList, listErr := pos.GetActiveVSSMemberList()
	if listErr != nil {
		log.Debugf("in UploadVSSConfig() return with error: %v", listErr)
		return listErr
	}
	threshold := pos.GetVssThreshold()

	// if we don't need to upload, just return
	if ret, _err := pos.ShouldUploadVSSConfig(activeNodeList, int(threshold), forceUpdate); !ret {
		log.Debugf("in UploadVSSConfig() no need to upload vss config, forceUpdate: %t, reason: %v",
			forceUpdate,
			_err,
		)
		return nil
	}

	// send request to contract at contractAddress,
	// internally, it will send our own scs id, then return its node index
	vssNodeIndex := pos.GetVSSNodeIndex()

	// return a mapping from address to index
	activeNodeAddressAndIndex := pos.GetVSSNodesIndexs(activeNodeList)

	// total number of vss nodes recorded before in subchainbase
	vssNodeCount := pos.GetVssNodeCount()

	log.Debugf(
		"vss configs activeNodeList: %v, vssNodeIndex: %d, activeNodeAddressAndIndex: %v, vssNodeCount: %d",
		activeNodeList,
		vssNodeIndex,
		activeNodeAddressAndIndex,
		vssNodeCount,
	)

	// get public keys for all active nodes
	vssAddressPublicKeys, _ := pos.GetVSSNodesPubkey(activeNodeList)
	// save it in pos. we need it later to verify proposer signature
	pos.nodesPubkey = vssAddressPublicKeys

	counter := 1
	for addr, pubkey := range vssAddressPublicKeys {
		log.Debugf("vss public key [%d/%d]: %s, %s", counter, len(vssAddressPublicKeys), common.Bytes2Hex(addr.Bytes()), common.Bytes2Hex(pubkey))
		counter++
	}
	vssIndexPublicKeys := make(map[int]kyber.Point)
	suite25519 := ed25519.NewBlakeSHA256Ed25519()
	for address, publickeyBytes := range vssAddressPublicKeys {
		if index, ok := activeNodeAddressAndIndex[address]; ok {
			publickey := suite25519.Point()
			_, err := publickey.UnmarshalFrom(bytes.NewReader(publickeyBytes))
			if err != nil {
				log.Errorf(
					"vss publickey can not be unmarshalfrom for address and pubkey bytes %v, %v",
					common.Bytes2Hex(address.Bytes()),
					common.Bytes2Hex(publickeyBytes),
				)
			}
			vssIndexPublicKeys[index] = publickey
		} else {
			log.Errorf("vss can not find index for address: %s", common.Bytes2Hex(address.Bytes()))
		}
	}

	// initialize new bls, should use len(activeNodeList) as n here instead of total node count
	sciBls := NewBLS(int(threshold), activeNodeList, int(vssNodeIndex.Int64()), int(vssNodeCount))
	privateShares := sciBls.GetPrivateShares(sciBls.TotalN)
	publicShares := sciBls.GetPublicShares(sciBls.TotalN)
	suite := bn256.NewSuite()

	// generate public shares
	allPublicShares := make([]*share.PubShare, 0)
	for _, nodeIndex := range activeNodeAddressAndIndex {
		for _, ps := range publicShares {
			if ps.I == nodeIndex {
				allPublicShares = append(allPublicShares, ps)
			}
		}
	}
	// set non active node's public share to be empty
	// so that the total size of allPublicShares always equals to total node count
	for i := sciBls.N; i < sciBls.TotalN; i++ {
		emptyShare := share.PubShare{
			I: -1,
			V: suite.G2().Point(),
		}
		allPublicShares = append(allPublicShares, &emptyShare)
	}
	// generate private shares
	allPrivateShares := make([]*share.PriShare, 0)
	for _, nodeIndex := range activeNodeAddressAndIndex {
		for _, ps := range privateShares {
			if ps.I == nodeIndex {
				allPrivateShares = append(allPrivateShares, ps)
			}
		}
	}
	// set non active node's private share to be empty
	// so that the total size of allPrivateShares always equals to total node count
	for i := sciBls.N; i < sciBls.TotalN; i++ {
		emptyShare := share.PriShare{
			I: -1,
			V: suite.G2().Scalar().Zero(),
		}
		allPrivateShares = append(allPrivateShares, &emptyShare)
	}

	// marshal the data. Private data need to encypted by the public key
	_publicShares := MarshalPublicShares(allPublicShares, pos.VssKey.Private)
	_privateShares := MarshalPrivateShares(allPrivateShares, vssIndexPublicKeys, pos.VssKey.Private)

	// zip the data for publib shares
	var bufPublicShares bytes.Buffer
	zw, _ := gzip.NewWriterLevel(&bufPublicShares, gzip.BestCompression)
	zw.Write(_publicShares)
	zw.Close()
	log.Debugf("upload vss config gzip: before (%d), after (%d)", len(_publicShares), len(bufPublicShares.Bytes()))
	// zip the data for private shares
	var bufPrivateShares bytes.Buffer
	zw, _ = gzip.NewWriterLevel(&bufPrivateShares, gzip.BestCompression)
	zw.Write(_privateShares)
	zw.Close()
	log.Debugf("upload vss config gzip: before (%d), after (%d)", len(_privateShares), len(bufPrivateShares.Bytes()))
	log.Debugf(
		"\nupload vss config with vss public shares: %v\nupload vss config with vss private shares: %v",
		string(_publicShares[:]),
		string(_privateShares[:]),
	)

	// set flush to true so we can use all 9000000 gaslimit. This tx is big
	// and consume more than register gaslimit settings. see sendtxtovnode()
	// for details
	pos.SendUploadVSSConfig(bufPublicShares.Bytes(), bufPrivateShares.Bytes())
	pos.uploadedConfigTime = time.Now()

	return nil
}

func (pos *Pos) SendUploadVSSConfig(publicShares []byte, privateShares []byte) {
	pos.vssbase.UploadVSSConfig(pos.transactOpts, publicShares, privateShares)
}

func (pos *Pos) vssVote(configVersion int) {
	_, err := pos.vssbase.Vote(pos.transactOpts, big.NewInt(int64(configVersion)))
	log.Debugf("vss vote config version: %d, err: %v", configVersion, err)
}

func (pos *Pos) vssReveal(
	violator common.Address,
	pubShare []byte,
	pubSig []byte,
	priShare []byte,
	priSig []byte,
	revealed []byte,
) {
	// set last parameter to true to use all the gas in case revealedSecrets is a long bytes array
	_, err := pos.vssbase.Reveal(pos.transactOpts, violator, pubShare, pubSig, priShare, priSig, revealed)
	log.Debugf(
		"vss reveal violator: %s, pubShare: %s, priShare: %s, "+
			"pubSig: %x, priSig: %x, revealed: %x, err: %v",
		violator.String(),
		string(pubShare[:]),
		string(priShare[:]),
		pubSig,
		priSig,
		revealed,
		err,
	)
}

func (pos *Pos) vssSlashing(index int, slash bool) {
	_, err := pos.vssbase.Slashing(pos.transactOpts, big.NewInt(int64(index)), slash)
	log.Debugf("vss slashing: %d, %t, err: %v", index, slash, err)
}

func (pos *Pos) vssReportSlowNode(slowNode common.Address) {
	_, err := pos.vssbase.ReportSlowNode(pos.transactOpts, slowNode)
	log.Debugf("vss report slow node: %s, %v", slowNode, err)
}

func (pos *Pos) registerVss() {
	var pubkey32 [32]byte
	vsskey := ToPersistVSSKey(pos.VssKey)
	copy(pubkey32[:], vsskey.Public[:32])

	_, err := pos.vssbase.RegisterVSS(pos.transactOpts, pos.Vssid, pubkey32)
	log.Debugf("vss register vss: %v", err)
}

func (pos *Pos) activateVss() {
	_, err := pos.vssbase.ActivateVSS(pos.transactOpts, pos.Vssid)
	log.Debugf("vss activate vss: %v", err)
}

func (pos *Pos) FindAddressInNodelist(myAddr common.Address) (int, int) {
	nodelist, _ := pos.GetActiveVSSMemberList()
	for i := 0; i < len(nodelist); i++ {
		if nodelist[i] == myAddr {
			return i, len(nodelist)
		}
	}
	return -1, len(nodelist)
}

func (pos *Pos) verifyShares(
	sciBls *BLS,
	suite *bn256.Suite,
	verifyPublicShares map[common.Address]*[]*share.PubShare,
	verifyPrivateShares map[common.Address]*[]*share.PriShare,
) map[common.Address]*share.PriShare {
	// verify received pub and pri shares
	notVerify := make(map[common.Address]*share.PriShare)
	for sender, pubshares := range verifyPublicShares {
		for _, pubshare := range *pubshares {
			// locate the pubshare for this node
			if pubshare.I == sciBls.NodeIndex {
				if prishares, found := verifyPrivateShares[sender]; found {
					for _, prishare := range *prishares {
						// locate the prishare for this node
						if prishare.I == sciBls.NodeIndex {
							// verify if they match
							verifyPub := suite.G2().Point().Mul(prishare.V, nil)
							if verifyPub.Equal(pubshare.V) {
								log.Debugf("share verified for sender: %x", sender)
							} else {
								notVerify[sender] = prishare
							}
						}
					}
				}
			}
		}
	}

	return notVerify
}
