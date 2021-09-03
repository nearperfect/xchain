package main

import (
	"context"
	"log"

	"github.com/MOACChain/xchain/mcclient"
)

func main() {
	client, err := mcclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545")
	client.SetFuncPrefix("eth")
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	currentBlock, _ := client.BlockNumber(context.Background())
	log.Printf("Chain ID: %d, Block Number: %d\n", chainID, currentBlock)
}
