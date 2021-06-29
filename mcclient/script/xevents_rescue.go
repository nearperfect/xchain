package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

func main() {
	rpc := os.Args[1]
	client, err := mcclient.Dial(fmt.Sprintf("http://%s", rpc))
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return
	} else {
		fmt.Println("%s", client)
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

	xeventsAddr := common.HexToAddress("0x0000000000000000000000000000000000010000")
	xevents, _ := xevents.NewXEvents(
		xeventsAddr,
		client,
	)

	vault := common.HexToAddress(os.Args[2])
	blockNumber, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Printf("failed getting blockNumber: %v", err)
		return
	}
	storeCounter, err := xevents.StoreCounter(&bind.CallOpts{})
	if err != nil {
		log.Printf("failed getting store counter: %v", err)
		return
	}
	log.Printf("Reset vault: %x to height %d, storeCounter: %d",
		vault, big.NewInt(int64(blockNumber)), storeCounter,
	)

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Printf("failed getting transactor: %v", err)
		return
	}

	MaxRepeats := 18
	for i := 0; i < MaxRepeats; i++ {
		time.Sleep(10 * time.Second)
		tx, err := xevents.Rescue(transactor, vault, big.NewInt(int64(blockNumber)), storeCounter)
		if err != nil {
			log.Printf("xevents rescue failed: %v", err)
		} else {
			log.Printf("xeevnts rescue Tx: %v", tx)
		}
	}
}
