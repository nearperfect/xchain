// Copyright 2015 The MOAC-core Authors
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
	"github.com/MOACChain/MoacLib/mcdb"
)

var (
	PairPrefix = []byte("TokenMappingLastBlock")
)

// DatabaseReader wraps the Get method of a backing data store.
type DatabaseReader interface {
	Get(key []byte) (value []byte, err error)
}

func LastBlockKey(pair string) []byte {
	return append(PairPrefix, []byte(pair)...)
}

func GetLastBlock(db DatabaseReader, pair string) uint64 {
	key := LastBlockKey(pair)
	_, _ = db.Get(key)
	return 0
}

func WriteLastBlock(db mcdb.Putter, pair string, value uint64) error {
	return nil
}
