// Copyright 2016 The MOAC-core Authors
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

package params

import (
	"math/big"
)

const (
	MinVSSRunInterval     uint64 = 2   // Min interval in seconds between two UpdateVSSConfig()
	MaximumBLSHistorySize uint64 = 100 // Maximum size of bls history stored
	MinBLSBlockNumber     uint64 = 1   // Minimum block number that bls sig in block header is required
	BlockInterval         uint64 = 10  // New block generate time interval
)

var (
	PosDifficultyPerBlock = big.NewInt(1000)
	PosBaseDifficulty     = big.NewInt(100000000000)
)
