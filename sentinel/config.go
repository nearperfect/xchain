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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/MOACChain/MoacLib/log"
)

/*
(chain x id, chain x rpc, chain prefix, vaultx contract address) <- watch deposit
(chain y id, chain y rpc, chain prefix, vaulty contract address) <- watch burn
*/
type VaultConfig struct {
	ChainId         uint64 `json:"id"  gencodec:"required"`
	ChainRPC        string `json:"rpc"  gencodec:"required"`
	ChainFuncPrefix string `json:"prefix"  gencodec:"required"`
	VaultAddress    string `json:"address"  gencodec:"required"`
}

type VaultPairConfig struct {
	VaultX VaultConfig `json:"vaultx"  gencodec:"required"`
	VaultY VaultConfig `json:"vaulty"  gencodec:"required"`
}

type VaultPairListConfig struct {
	Vaults []VaultPairConfig `json:"vaults"  gencodec:"required"`
}

func (pair *VaultPairConfig) Id() string {
	// x chainid, x vault addr, y chainid, y vault addr
	return fmt.Sprintf(
		"%d,%s,%d,%s",
		pair.VaultX.ChainId,
		pair.VaultX.VaultAddress,
		pair.VaultY.ChainId,
		pair.VaultY.VaultAddress,
	)
}

func GetConfiguration(configFilePath string) (*VaultPairListConfig, error) {
	//check the file path
	filePath, _ := filepath.Abs(configFilePath)
	log.Debugf("load vnode config file from %v", filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		//If no config file exists, return nil
		log.Infof("%v not exists, use default settings", configFilePath)
		return nil, nil
	}

	if _, err := os.Stat(configFilePath); err != nil {

		log.Errorf("Open %v error: \n%v\n", configFilePath, err)

		return nil, err
	}

	return ParseConfig(configFilePath)
}

func ParseConfig(filepath string) (*VaultPairListConfig, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize the Confgiruation array
	var conf VaultPairListConfig

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &conf)

	return &conf, nil
}
