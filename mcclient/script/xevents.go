package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
)

func main() {
	client, err := mcclient.Dial("http://0.0.0.0:18545")
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
	log.Printf("%x, %d, %d\n", fromAddress, chainID, nonce)

	xevents, _ := xdefi.NewXEvents(
		common.HexToAddress("0x0000000000000000000000000000000000010000"),
		client,
	)

	callOpts := &bind.CallOpts{}
	input, err := xevents.Input(callOpts)
	log.Printf("input = %d, err: %v", input, err)

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	_, err = xevents.Test(transactOpts, big.NewInt(100))
	log.Printf("Test(): err: %v", err)

	inputAfter, err := xevents.Input(callOpts)
	log.Printf("input after = %d, err: %v ", inputAfter, err)
}
