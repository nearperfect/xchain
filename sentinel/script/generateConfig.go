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
		56,
		"https://data-seed-prebsc-2-s3.binance.org:8545",
		"eth",
		"0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0",
	}
	vaultY := sentinel.VaultConfig{
		95125,
		"http://172.20.0.11:8545",
		"mc",
		"0x1B041aC1cFE6167e1F1D0ABE1A1A941C9666ac22",
	}

	vaultPairs := []sentinel.VaultPairConfig{sentinel.VaultPairConfig{vaultX, vaultY}}
	vaults := sentinel.VaultPairListConfig{vaultPairs}

	result, _ := json.Marshal(vaults)
	fmt.Println(string(result))

	var vaults_ sentinel.VaultPairListConfig
	err := json.Unmarshal(result, &vaults_)
	fmt.Printf("%v, err: %v", vaults_, err)
}
