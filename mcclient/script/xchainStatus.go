package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/sentinel"
	"github.com/MOACChain/xchain/xdefi/vaultx"
	"github.com/MOACChain/xchain/xdefi/vaulty"
	"github.com/MOACChain/xchain/xdefi/xconfig"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

/////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////
//
// To run: ./xchainStatus http://127.0.0.1:18545
//

func main() {
	// read config file
	xchainRPC := os.Args[1]
	callOpts := &bind.CallOpts{}

	client, err := mcclient.Dial(xchainRPC)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
		return
	}

	// initialize xconfig
	xconfigAddr := common.HexToAddress("0x0000000000000000000000000000000000010000")
	xconfigContract, err := xconfig.NewXConfig(
		xconfigAddr,
		client,
	)

	configVersion, err := xconfigContract.VaultConfigVersion(callOpts)
	data, err := xconfigContract.VaultConfigs(callOpts, configVersion)

	// validate the content of the config file
	var vaultsConfig sentinel.VaultPairListConfig
	err = json.Unmarshal(data, &vaultsConfig)
	if err == nil {
		//fmt.Printf("Config: \n%s", string(data))
	} else {
		log.Fatalf("Config err: %v", err)
		return
	}

	// initialize xevents vault contract
	xeventsXYAddr := common.HexToAddress("0x0000000000000000000000000000000000010001")
	xeventsXY, err := xevents.NewXEvents(
		xeventsXYAddr,
		client,
	)
	if err != nil {
		log.Printf("Initialize xevents XY contract failed: %v", err)
		return
	}
	xeventsYXAddr := common.HexToAddress("0x0000000000000000000000000000000000010002")
	xeventsYX, err := xevents.NewXEvents(
		xeventsYXAddr,
		client,
	)
	if err != nil {
		log.Printf("Initialize xevents YX contract failed: %v", err)
		return
	}

	fmt.Printf("###############################\n")
	fmt.Printf("#######   Vault Scan  #########\n")
	fmt.Printf("###############################\n")
	for index, vaultConfig := range vaultsConfig.Vaults {
		vaultXAddr := common.HexToAddress(vaultConfig.VaultX.VaultAddress)
		vaultYAddr := common.HexToAddress(vaultConfig.VaultY.VaultAddress)

		clientX, err := mcclient.Dial(vaultConfig.VaultX.ChainRPC)
		vaultXContract, _ := vaultx.NewVaultX(
			vaultXAddr,
			clientX,
		)
		clientY, err := mcclient.Dial(vaultConfig.VaultY.ChainRPC)
		clientY.SetFuncPrefix("eth")
		vaultYContract, _ := vaulty.NewVaultY(
			vaultYAddr,
			clientY,
		)

		currentBlockX, errx := clientX.BlockNumber(context.Background())
		currentBlockY, erry := clientY.BlockNumber(context.Background())
		if errx != nil {
			fmt.Printf("errx %v", errx)
		}
		if erry != nil {
			fmt.Printf("erry %v", erry)
		}

		fmt.Printf(
			"\nVault Pair X <-> Y (vault %x -> vault %x) # %d:\n",
			vaultXAddr.Bytes()[:3], vaultYAddr.Bytes()[:3],
			index,
		)

		for _, tokenMapping := range vaultConfig.TokenMappings {
			sourceToken := common.HexToAddress(tokenMapping.SourceToken)
			mappedToken := common.HexToAddress(tokenMapping.MappedToken)

			sha256 := sentinel.TokenMappingSha256(sentinel.TokenMappingString(
				big.NewInt(int64(tokenMapping.SourceChainId)),
				sourceToken.Bytes(),
				big.NewInt(int64(tokenMapping.MappedChainId)),
				mappedToken.Bytes(),
			))

			// deposit nonce
			depositNonce, _ := vaultXContract.TokenMappingDepositNonce(
				callOpts, sourceToken, mappedToken,
			)
			fmt.Printf("\n")
			fmt.Printf("\tToken mapping: %x <-> %x\n\n", sourceToken.Bytes()[:3], mappedToken.Bytes()[:3])
			fmt.Printf("\t\tX deposit nonce: %d\n", depositNonce)

			// burn nonce
			burnNonce, _ := vaultYContract.TokenMappingBurnNonce(
				callOpts, sourceToken, mappedToken,
			)
			fmt.Printf("\t\tY burn nonce: %d\n", burnNonce)
			fmt.Printf("\n\t\t-------------\n\n")

			res, _ := xeventsXY.TokenMappingWatermark(callOpts, vaultXAddr, sha256)
			fmt.Printf("\t\tX -> Y nonce: %d\n", res)
			res, _ = xeventsYX.TokenMappingWatermark(callOpts, vaultYAddr, sha256)
			fmt.Printf("\t\tY -> X nonce: %d\n", res)
		}
		fmt.Printf("\n")
		watermarkX, err := xeventsXY.VaultWatermark(callOpts, vaultXAddr)

		if err != nil {
			fmt.Printf("\tvault X scan err: %v\n", err)
		} else {
			fmt.Printf(
				"\tvault X scan block: %d (chain id:%d, #%d)\n",
				watermarkX, vaultConfig.VaultX.ChainId, currentBlockX,
			)
		}

		watermarkY, err := xeventsYX.VaultWatermark(callOpts, vaultYAddr)

		if err != nil {
			fmt.Printf("\tvault Y scan err: %v\n", err)
		} else {
			fmt.Printf(
				"\tvault Y scan block: %d (chain id:%d, #%d)\n",
				watermarkY, vaultConfig.VaultY.ChainId, currentBlockY,
			)
		}
	}

	fmt.Printf("\n###############################\n")
	fmt.Printf("#######   Vault Config  #######\n")
	fmt.Printf("###############################\n\n")
	fmt.Printf("Config version: %d\n\n", configVersion)
	fmt.Printf("%s\n", data)
}
