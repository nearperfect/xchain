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
	"fmt"
	"math"
)

func Max(input []uint64) (uint64, error) {
	max := uint64(0)
	if len(input) == 0 {
		return max, fmt.Errorf("Empty input array")
	}
	for _, n := range input {
		if n > max {
			max = n
		}
	}
	return max, nil
}

func Min(input []uint64) (uint64, error) {
	min := uint64(math.MaxUint64)
	if len(input) == 0 {
		return min, fmt.Errorf("Empty input array")
	}
	for _, n := range input {
		if n < min {
			min = n
		}
	}
	return min, nil
}
