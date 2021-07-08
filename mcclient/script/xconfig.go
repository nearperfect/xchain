package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/sentinel"
	"github.com/MOACChain/xchain/xdefi/xconfig"
)

/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
//
// To run: go run mcclient/script/xconfig.go http://192.168.0.156:18545 ./vaults.json
//

func main() {
	// read config file
	xchainRPC := os.Args[1]
	configFile := os.Args[2]
	data, err := ioutil.ReadFile(configFile)

	// validate the content of the config file
	var vaultsConfig sentinel.VaultPairListConfig
	err = json.Unmarshal(data, &vaultsConfig)
	if err == nil {
		log.Printf("Config: %s", string(data))
	} else {
		log.Fatalf("Config err: %v", err)
		return
	}

	// initialize client, key, chainid and nonce
	client, err := mcclient.Dial(xchainRPC)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return
	}
	privateKey, err := crypto.HexToECDSA("393873d6bbc61b9d83ba923e08375b7bf8210a12bed4ea2016d96021e9378cc9")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("0x%x, %d, %d\n", fromAddress, chainID, nonce)

	// initialize xconfig vault contract
	xconfigAddr := common.HexToAddress("0x0000000000000000000000000000000000010000")
	xconfig, err := xconfig.NewXConfig(
		xconfigAddr,
		client,
	)
	if err != nil {
		log.Printf("Initialize xconfig contract failed: %v", err)
		return
	}

	blockNumberBefore, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("fail get block number: %v", err)
		return
	}
	transactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	// call initilaize()
	tx, err := xconfig.Initialize(transactor)
	if err != nil {
		log.Printf("Call xconfig Initialize() failed: %v", err)
	} else {
		log.Printf("Initialize() Tx: %v", tx)
	}

	// wait for initilaize() tx to be confirmed
	for {
		time.Sleep(5 * time.Second)
		blockNumberAfter, err := client.BlockNumber(context.Background())
		if err != nil {
			continue
		}
		if blockNumberAfter > blockNumberBefore {
			break
		}
	}

	// call updateConfig()
	tx, err = xconfig.UpdateVaultConfig(transactor, data)
	if err != nil {
		log.Printf("call xconfig update config failed: %v, %v", err, transactor.GasLimit)
	} else {
		log.Printf("UpdateConfig() Tx: %v", tx)
	}
}
