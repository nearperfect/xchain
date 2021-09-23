package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
//
// To run: ./xchainRescue http://127.0.0.1:18545 x2y 0xa7eb59c3074fe608419796ae2eb73bae3f576079 600000
// parameters:
//     1. rpc address
//     2. which vault: x2y for vault x, y2x for vault y
//     3. specify vault address
//     4. reset height

func main() {
	// read config file
	xchainRPC := os.Args[1]
	direction := os.Args[2]

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
	var xeventsAddr common.Address
	if direction == "x2y" {
		xeventsAddr = common.HexToAddress("0x0000000000000000000000000000000000010001")
	} else if direction == "y2x" {
		xeventsAddr = common.HexToAddress("0x0000000000000000000000000000000000010002")
	} else {
		log.Printf("Direction \"%s\" not known. Use either \"x2y\" or \"y2x\"", direction)
	}

	xeventsContract, err := xevents.NewXEvents(
		xeventsAddr,
		client,
	)
	if err != nil {
		log.Printf("Initialize xevents contract failed: %v", err)
		return
	}
	transactor, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	vault := common.HexToAddress(os.Args[3])
	blockNumber, _ := strconv.Atoi(os.Args[4])
	// call initilaize()
	tx, err := xeventsContract.RescueVault(transactor, vault, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Printf("Call xconfig Initialize() failed: %v", err)
	} else {
		log.Printf("Initialize() Tx: %v", tx)
	}
}
