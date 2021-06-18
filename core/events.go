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

package core

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/rlp"
	"github.com/MOACChain/MoacLib/types"
	"github.com/MOACChain/xchain/dkg"
)

// TxPreEvent is posted when a transaction enters the transaction pool.
type TxPreEvent struct{ Tx *types.Transaction }

// TxPreEvent is posted when a transaction enters the transaction pool.
type ShardingTxEvent struct{ Tx *types.Transaction }

// PendingLogsEvent is posted pre mining and notifies of pending logs.
type PendingLogsEvent struct {
	Logs []*types.Log
}

// PendingStateEvent is posted pre mining and notifies of pending state changes.
type PendingStateEvent struct{}

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// NewSigShareEvent is
type NewSigShareEvent struct{ SigShare *dkg.SigShareMessage }

// RemovedTransactionEvent is posted when a reorg happens
type RemovedTransactionEvent struct{ Txs types.Transactions }

// RemovedLogsEvent is posted when a reorg happens
type RemovedLogsEvent struct{ Logs []*types.Log }

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  []*types.Log
}

type ChainSideEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }

type TxShardJoinEvent struct{ Tx *types.Transaction }

type SigShares []*dkg.SigShareMessage

type VaultEvents []*VaultEvent

type VaultEventWithSigs []*VaultEventWithSig

type VaultEvent struct {
	Vault         common.Address `json:"vault"          gencodec:"required"`
	SourceChainid *big.Int       `json:"sourcechainid"  gencodec:"required"`
	SourceToken   common.Address `json:"sourcetoken"    gencodec:"required"`
	MappedChainid *big.Int       `json:"mappedchainid"  gencodec:"required"`
	MappedToken   common.Address `json:"mappedtoken"    gencodec:"required"`
	To            common.Address `json:"to"             gencodec:"required"`
	Amount        *big.Int       `json:"amount"         gencodec:"required"`
	Nonce         *big.Int       `json:"nonce"          gencodec:"required"`
	BlockNumber   *big.Int       `json:"blockNumber"    gencodec:"required"`
}

func (vaultEvent *VaultEvent) Hash() common.Hash {
	return common.RlpHash(vaultEvent)
}

func (vaultEvent *VaultEvent) TokenMapping() string {
	return fmt.Sprintf(
		"%d,%x,%d,%x",
		vaultEvent.SourceChainid,
		vaultEvent.SourceToken.Bytes(),
		vaultEvent.MappedChainid,
		vaultEvent.MappedToken.Bytes(),
	)
}

func (vaultEvent *VaultEvent) TokenMappingSha256() [32]byte {
	h := sha256.New()
	h.Write([]byte(vaultEvent.TokenMapping()))
	sha := h.Sum(nil)
	var sha32 [32]byte
	copy(sha32[:], sha)
	return sha32
}

func (vaultEvent *VaultEvent) TokenMappingWithVault() string {
	return fmt.Sprintf(
		"%x,%d,%x,%d,%x",
		vaultEvent.Vault.Bytes(),
		vaultEvent.SourceChainid,
		vaultEvent.SourceToken.Bytes(),
		vaultEvent.MappedChainid,
		vaultEvent.MappedToken.Bytes(),
	)
}

func (vaultEvent *VaultEvent) Bytes() []byte {
	data, err := rlp.EncodeToBytes(vaultEvent)
	if err != nil {
		log.Errorf("vault event rlp encode, err:", err)
		return []byte{}
	}

	return data
}

type VaultEventWithSig struct {
	Event  VaultEvent `json:"event"  gencodec:"required"`
	Blssig []byte     `json:"blssig"  gencodec:"required"`
}

func (vaultEventWithSig *VaultEventWithSig) Hash() common.Hash {
	return common.RlpHash(vaultEventWithSig)
}

func (vaultEvent *VaultEvent) String() string {
	return fmt.Sprintf(
		"vault[%x], source[%s]: %x, mapped[%s]: %x, account: %x, amount: %s",
		vaultEvent.Vault.Bytes()[:8],
		vaultEvent.SourceChainid,
		vaultEvent.SourceToken.Bytes()[:8],
		vaultEvent.MappedToken,
		vaultEvent.MappedToken.Bytes()[:8],
		vaultEvent.To.Bytes()[:8],
		vaultEvent.Amount,
	)
}
