package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

func waitBlocks(blocks uint64, client *mcclient.Client) {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		panic("Can not get block number")
	}

	lastBlockNumber := blockNumber
	for {
		if blockNumber, err := client.BlockNumber(context.Background()); err != nil {
			panic("Can not get block number")
		} else {
			fmt.Printf("[wait] block number: %d\n", blockNumber)
			if blockNumber-lastBlockNumber > 3 {
				break
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func IsInitialized(xe *xevents.XEvents, expectedAdmin common.Address) bool {
	initialized := false
	callOpts := &bind.CallOpts{}
	roles, _ := xe.GetRoles(callOpts)
	for _, role := range roles {
		fmt.Printf("[role: %s] %x\n", role.Describe, role.Role)
		if role.Describe == "admin" {
			if admins, err := xe.GetRoleMembers(callOpts, role.Role); err != nil {
				panic("can not get admins")
			} else {
				for _, admin := range admins {
					if admin == expectedAdmin {
						initialized = true
						fmt.Printf("-------- Initialized with admin: 0x%x --------\n", expectedAdmin)
					}
				}
			}
		}
	}

	return initialized
}

func main() {
	// read config file
	xchainRPC := os.Args[1]
	validator := common.HexToAddress(os.Args[2])

	client, err := mcclient.Dial(xchainRPC)
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

	xeventsXYAddr := common.HexToAddress("0x0000000000000000000000000000000000010001")
	xeventsXY, _ := xevents.NewXEvents(
		xeventsXYAddr,
		client,
	)
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	// call initialize()
	initialized := IsInitialized(xeventsXY, fromAddress)
	if !initialized {
		if tx, err := xeventsXY.Initialize(transactOpts); err == nil {
			log.Printf("Initialize(): tx: %v", tx)
		} else {
			log.Printf("Initialize(): err: %v", err)
		}
	}
	waitBlocks(3, client)
	// call grantValidator
	if tx, err := xeventsXY.GrantValidator(transactOpts, validator); err == nil {
		log.Printf("grantValidator(): tx: %v", tx)
	} else {
		log.Printf("grantValidator(): err: %v", err)
	}

	xeventsYXAddr := common.HexToAddress("0x0000000000000000000000000000000000010002")
	xeventsYX, _ := xevents.NewXEvents(
		xeventsYXAddr,
		client,
	)
	waitBlocks(3, client)
	// call initialize()
	initialized = IsInitialized(xeventsYX, fromAddress)
	if !initialized {
		if tx, err := xeventsYX.Initialize(transactOpts); err == nil {
			log.Printf("Initialize(): tx: %v", tx)
		} else {
			log.Printf("Initialize(): err: %v", err)
		}
	}
	waitBlocks(3, client)
	// call grantValidator
	if tx, err := xeventsYX.GrantValidator(transactOpts, validator); err == nil {
		log.Printf("grantValidator(): tx: %v", tx)
	} else {
		log.Printf("grantValidator(): err: %v", err)
	}
}
