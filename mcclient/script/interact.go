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
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/mcclient/xdefi"
)

func main() {
	client, err := mcclient.Dial("http://172.21.0.11:8545")
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

	vaultx, _ := xdefi.NewVaultX(
		common.HexToAddress("0xABE1A1A941C9666ac221B041aC1cFE6167e1F1D0"),
		client,
	)
	log.Printf("%v\n", vaultx)
	callOpts := &bind.CallOpts{}
	validators, _ := vaultx.GetValidators(callOpts)
	for _, validator := range validators {
		log.Printf("%x\n", validator.Bytes())
	}

	allValidators := []string{
		"a35add395b804c3faacf7c7829638e42ffa1d051",
		"da8ad06b2a20c6f92641d185c22f0479b00a90f3",
		"f34c3a04099a76dda80517373b21409391540b82",
		"903ee4f9753b3717aa6a295b02095aa0c94036d0",
		"f084d898a6329d0d9159ddccca0380d651ee1c17",
		"3563e38cc436bd6835da191228115fe7869a382c",
		"a221d547d2e3821f24924d7bd89e443045d81f6e",
		"7ac799d9fb930fafc3d50937b10ea30f0c1c30ce",
		"b544fd6b593807b864998836db91ab0d81626745",
		"b94b69cc0fb38a6b1be2cd5466f0676b7d5be7f8",
		"0fcdc2ec292878c15449d02d4e4928694f9a5baf",
		"9526b1366d5328d8a3bb0eca0489deff7282a5fb",
		"6586f8ad114271c2c84287a1b2e2b4794aee3868",
		"1be6959a8498c91fe2599cf54d084fdd347e3929",
		"cececba5e000d0b2c63de7e2f862be5c9d1e6123",
		"0ae5b3913922eca2d781a9bc56d517987bdb9176",
		"ad2139c3c35e61e11fbc880780accffb14c367e2",
		"a6eca2a8b0109aefe1c4f6e9041c641cf79a76b3",
		"b639f28531e832e7c90362f883f4a82bc5bab5ee",
		"ec6234cbaae7ee4fd85d8c288054f89f3be29c81",
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	for _, validator := range allValidators {
		//tx, err := vaultx.AddValidator(transactOpts, common.HexToAddress(validator))
		//log.Printf("%s, %v", tx, err)
		vaultx.AddValidator(transactOpts, common.HexToAddress(validator))
	}

	// get events
	filterOpts := &bind.FilterOpts{
		Context: context.Background(),
		Start:   1,
		End:     nil,
	}
	itr, err := vaultx.FilterRoleGranted(
		filterOpts,
		[][32]byte{},
		[]common.Address{},
		[]common.Address{},
	)
	for itr.Next() {
		event := itr.Event
		log.Printf("%x, %x, %x\n", event.Role, event.Account.Bytes(), event.Sender.Bytes())
	}

	// get events from vaultx
	filterOpts2 := &bind.FilterOpts{
		Context: context.Background(),
		Start:   1,
		End:     nil,
	}
	itr2, err := vaultx.FilterTokenDeposit(
		filterOpts2,
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	for itr2.Next() {
		event := itr2.Event
		log.Printf("vault event 1 %x, %x, %x, %s, %d",
			event.SourceToken.Bytes(),
			event.MappedToken.Bytes(),
			event.From.Bytes(),
			event.Amount,
			event.DepositNonce.Uint64(),
		)

		vaultEvent := core.VaultEvent{
			event.SourceChainid,
			event.SourceToken,
			event.MappedChainid,
			event.MappedToken,
			event.From,
			event.Amount,
			event.DepositNonce,
			[]byte{},
		}
		log.Printf("vault event 2 %v", vaultEvent)
	}
}
