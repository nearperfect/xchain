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
package pos

import (
	"github.com/MOACChain/MoacLib/common"
	blslib "github.com/innowells/bls_lib/v2"
	kyber "go.dedis.ch/kyber/v3"
	share "go.dedis.ch/kyber/v3/share"
)

type BLS blslib.BLS

type VSS struct {
	VSSNodeChan  chan *VSSNode
	SigShareChan chan *SigShareMessage
}

type SigShareMessage blslib.SigShareMessage

type VSSNode blslib.VSSNode

type PersistPrivateShare blslib.PersistPrivateShare

type PersistPrivateShares blslib.PersistPrivateShares

type PersistPublicShare blslib.PersistPublicShare

type PersistPublicShares blslib.PersistPublicShares

type VSSKey blslib.VSSKey

type PersistVSSKey blslib.PersistVSSKey

type VSSResult blslib.VSSResult

func (s SigShareMessage) String() string {
	return (blslib.SigShareMessage)(s).String()
}

func (s SigShareMessage) Key() string {
	return (blslib.SigShareMessage)(s).Key()
}

func (s SigShareMessage) Hash() common.Hash {
	return (blslib.SigShareMessage)(s).Hash()
}

func (bls *BLS) SignBlockHash(blockHash common.Hash) []byte {
	return (*blslib.BLS)(bls).SignBlockHash(blockHash)
}

func (bls *BLS) SignBytes(bytes []byte) []byte {
	return (*blslib.BLS)(bls).SignBytes(bytes)
}

func (bls *BLS) GetPrivateShares(n int) []*share.PriShare {
	return (*blslib.BLS)(bls).GetPrivateShares(n)
}

func (bls *BLS) GetPublicShares(n int) []*share.PubShare {
	return (*blslib.BLS)(bls).GetPublicShares(n)
}

func (bls *BLS) GenerateRandomNumber(toSign []byte, sigs map[int][]byte) (VSSResult, error) {
	r, err := (*blslib.BLS)(bls).GenerateRandomNumber(toSign, sigs)
	return VSSResult(r), err
}

func ToPersistVSSKey(vsskey *VSSKey) *PersistVSSKey {
	r := blslib.ToPersistVSSKey((*blslib.VSSKey)(vsskey))
	return (*PersistVSSKey)(r)
}

func GenerateVSSKey() *VSSKey {
	r := blslib.GenerateVSSKey()
	return (*VSSKey)(r)
}

func MarshalPrivateShares(activePrishares []*share.PriShare, publickeys map[int]kyber.Point, key kyber.Scalar) []byte {
	return blslib.MarshalPrivateShares(activePrishares, publickeys, key)
}

func UnmarshalPrivateShares(input *[]byte, private kyber.Scalar, localNodeIndex int, pubkey kyber.Point) (*[]*share.PriShare, PersistPrivateShares) {
	r, v := blslib.UnmarshalPrivateShares(input, private, localNodeIndex, pubkey)
	return r, PersistPrivateShares(v)
}

func MarshalPublicShares(pubshares []*share.PubShare, key kyber.Scalar) []byte {
	return blslib.MarshalPublicShares(pubshares, key)
}

func UnmarshalPublicShares(input *[]byte, pubkey kyber.Point) (*[]*share.PubShare, PersistPublicShares) {
	r, v := blslib.UnmarshalPublicShares(input, pubkey)
	return r, PersistPublicShares(v)
}

// sign message using ed25519 private key
func SignED25519(private kyber.Scalar, message []byte) ([]byte, error) {
	return blslib.SignED25519(private, message)
}

// verify message using ed25519 public key
func VerifyED25519(public kyber.Point, msg, sig []byte) (bool, error) {
	return blslib.VerifyED25519(public, msg, sig)
}

func NewBLS(threshold int, activeNodeList []common.Address, localNodeIndex int, totaln int) *BLS {
	r := blslib.NewBLS(threshold, activeNodeList, localNodeIndex, totaln)
	return (*BLS)(r)
}
