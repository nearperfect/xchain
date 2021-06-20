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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/MOACChain/xchain/sentinel"
)

func main() {
	vaultX := sentinel.VaultConfig{
		95125,
		"http://172.21.0.11:8545",
		"mc",
		"0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0",
	}
	vaultY := sentinel.VaultConfig{
		95125,
		"http://172.21.0.11:8545",
		"mc",
		"0xcCa8BAA2d1E83A38bdbcF52a9e5BbB530f50493A",
	}

	tokenMapping := sentinel.TokenMapping{
		95125,
		"0x350e47237eb2515b3b30c2f232268b998e392409",
		95125,
		"0x8553ce822a9072b5ff0992da9a61d5ce54a1f5df",
	}
	vaultPairs := []sentinel.VaultPairConfig{
		sentinel.VaultPairConfig{
			vaultX,
			vaultY,
			[]sentinel.TokenMapping{tokenMapping},
		},
	}
	vaults := sentinel.VaultPairListConfig{vaultPairs}

	result, _ := json.Marshal(vaults)
	fmt.Println(string(result))

	var vaults_ sentinel.VaultPairListConfig
	err := json.Unmarshal(result, &vaults_)
	fmt.Printf("%v, err: %v", vaults_, err)
}
