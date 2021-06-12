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

package pos

import (
	"errors"
	"fmt"
	"math/big"
	"runtime"
	//	"strconv"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/params"
	"github.com/MOACChain/MoacLib/state"
	"github.com/MOACChain/MoacLib/types"
	"github.com/MOACChain/xchain/consensus"
	"github.com/MOACChain/xchain/rpc"
)

// Pos proof-of-work protocol constants.

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	errLargeBlockTime = errors.New("timestamp too big")
	errZeroBlockTime  = errors.New("timestamp equals parent's")
)

// Author implements consensus.Engine, returning the header's coinbase as the
// proof-of-work verified author of the block.
func (pos Pos) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (pos Pos) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications.
func (pos Pos) VerifyHeaders(
	chain consensus.ChainReader,
	headers []*types.Header,
	seals []bool,
) (chan<- struct{}, <-chan error) {
	fixBlock := false
	fromBroadcast := true

	// Spawn as many workers as allowed threads
	workers := runtime.GOMAXPROCS(0)
	if len(headers) < workers {
		workers = len(headers)
	}

	// Create a task channel and spawn the verifiers
	var (
		inputs = make(chan int)
		done   = make(chan int, workers)
		errors = make([]error, len(headers))
		abort  = make(chan struct{})
	)
	for i := 0; i < workers; i++ {
		go func() {
			for index := range inputs {
				errors[index] = pos.verifyHeaderWorker(
					chain,
					headers,
					seals,
					index,
					fixBlock,
					fromBroadcast,
				)
				done <- index
			}
		}()
	}

	errorsOut := make(chan error, len(headers))
	go func() {
		defer close(inputs)
		var (
			in, out = 0, 0
			checked = make([]bool, len(headers))
			inputs  = inputs
		)
		for {
			select {
			case inputs <- in:
				if in++; in == len(headers) {
					// Reached end of headers. Stop sending to workers.
					inputs = nil
				}
			case index := <-done:
				for checked[index] = true; checked[out]; out++ {
					errorsOut <- errors[out]
					if out == len(headers)-1 {
						return
					}
				}
			case <-abort:
				return
			}
		}
	}()
	return abort, errorsOut
}

func (pos Pos) verifyHeaderWorker(
	chain consensus.ChainReader,
	headers []*types.Header,
	seals []bool,
	index int,
	fixBlock bool,
	fromBroadcast bool,
) (err error) {
	// use defer to record received blocks before return
	defer func() {
		// if there is no header error or just unknown ancestor err
		if err == nil || err == consensus.ErrUnknownAncestor {
			//header := headers[index]
			//extraData, _ := header.ExtraData()
			//epoch := (*extraData)[params.EpochField]
			//epochInt, _ := strconv.Atoi(string(epoch))

			// only record block from broadcast or mining (not from sync)
			if fromBroadcast {
				/*
					sci.RecordBlockInfo(
						header.Coinbase.String(),
						int(header.Number.Int64()),
						header.Hash().Bytes(),
						header.ParentHash.Bytes(),
						epochInt,
					)*/
			}
		}
	}()

	var parent *types.Header
	if index == 0 {
		parent = chain.GetHeader(headers[index].ParentHash, headers[index].Number.Uint64()-1)
	} else if headers[index-1].Hash() == headers[index].ParentHash {
		parent = headers[index-1]
	}
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// known block
	if chain.GetHeader(headers[index].Hash(), headers[index].Number.Uint64()) != nil {
		return nil
	}

	// if we find the parent, this is the normal case
	return pos.verifyHeader(chain, headers[index], parent, seals[index], fixBlock)
}

func (pos Pos) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return nil
}

// verifyHeader checks whether a header conforms to the consensus rules of the
// stock MoacNode pos engine.
func (pos Pos) verifyHeader(
	chain consensus.ChainReader,
	header, parent *types.Header,
	seal bool,
	fixBlock bool,
) error {
	// Ensure that the header's extra-data section is of a reasonable size
	if uint64(len(header.Extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra-data too long: %d > %d", len(header.Extra), params.MaximumExtraDataSize)
	}

	// Verify that the block number is parent + 1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}

	if header.Time.Cmp(parent.Time) <= 0 {
		return errZeroBlockTime
	}

	// Verify the engine specific seal securing the block
	if seal {
		if err := pos.VerifySeal(chain, header); err != nil {
			return err
		}
	}

	// verify bls related fields
	if pos.IsVSSEnabled() {
		// get extra data field
		extraData, err := header.ExtraData()
		if err != nil {
			return fmt.Errorf("header extra data error: %v", err)
		}
		log.Debugf("extra data received %v", extraData)

		// verify bls signature in extra data
		blsSig := (*extraData)[params.BLSSigField]
		epoch := (*extraData)[params.EpochField]
		parentHash := (*extraData)[params.ParentHashField]
		// coinbase is proposer address
		toSign := fmt.Sprintf(
			"%x,%x,%d,%s",
			parentHash, header.Coinbase.Bytes(), header.Number.Int64(), epoch,
		)
		validated := pos.ValidateBLSSig([]byte(toSign), blsSig, fixBlock, header.Number.Int64())
		log.Debugf(
			"validate bls sig block number = %d, result = %t, fixblock = %t, bissig = %v, toSign = %s",
			header.Number.Int64(),
			validated,
			fixBlock,
			common.Bytes2Hex(blsSig),
			toSign,
		)
		if !validated {
			return fmt.Errorf("bls sig not verify")
		}

		// verify proposer signature in extra data
		proposerSig := (*extraData)[params.ProposerSig]
		validatedProposerSig := pos.ValidateProposerSig(
			proposerSig, header.Coinbase, blsSig, fixBlock, header.Number.Int64(),
		)
		if !validatedProposerSig {
			return fmt.Errorf("proposer signature not verify")
		}
	}

	return nil
}

// VerifySeal implements consensus.Engine. Empty for now.
func (pos Pos) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// Finalize implements consensus.Engine, accumulating the block and uncle rewards,
// setting the final state and assembling the block.
func (pos Pos) Finalize(
	chain consensus.ChainReader,
	header *types.Header,
	state *state.StateDB,
	txs []*types.Transaction,
	uncles []*types.Header,
	receipts []*types.Receipt,
	liveFlag bool,
) (*types.Block, error) {
	header.Root = state.IntermediateRoot(true)
	log.Debugf(
		"Finalize/IntermediateRoot(true) num: %d, root: 0x%x, live: %t",
		header.Number, header.Root, liveFlag,
	)
	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, uncles, receipts), nil
}

// APIs implements consensus.Engine, returning the user facing RPC APIs. Currently
// that is empty.
func (pos Pos) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}

// currently empty
func (pos Pos) Prepare(chain consensus.ChainReader, header *types.Header) error {
	return nil
}

// return true, assuming no monitor nodes
func (pos Pos) IsConsensus() bool {
	return true
}
