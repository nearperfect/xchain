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
	//"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/types"
	"github.com/MOACChain/xchain/consensus"
)

// Seal implements consensus.Engine. Set the signer in the header.
/*
func (pos *Pos) Seal(block *types.Block, signer common.Address) *types.Block {
	result := block
	result.SetHeader = types.CopyHeader(block.Header)
	result.Header.Coinbase = signer
	return result
}*/

func (pos Pos) Seal(reader consensus.ChainReader, block *types.Block, c <-chan struct{}) (*types.Block, error) {
	return block, nil
}
