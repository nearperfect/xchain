// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultx

import (
	"math/big"
	"strings"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/types"
	moaccore "github.com/MOACChain/xchain"
	"github.com/MOACChain/xchain/accounts/abi"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = moaccore.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RoleAccessRole is an auto generated low-level Go binding around an user-defined struct.
type RoleAccessRole struct {
	Role     [32]byte
	Describe string
}

// VaultXtokenPairInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultXtokenPairInfo struct {
	SourceToken        common.Address
	SourceTokenChainid *big.Int
	SourceTokenSymbol  string
	MappedToken        common.Address
	MappedTokenChainid *big.Int
	MappedTokenSymbol  string
	Paused             bool
}

// VaultXABI is the input ABI used to generate the binding from.
const VaultXABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeAccount\",\"type\":\"address\"}],\"name\":\"FeeAccountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newFiatCurrency\",\"type\":\"string\"}],\"name\":\"FiatCurrencyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFiatFeeAmount\",\"type\":\"uint256\"}],\"name\":\"FiatFeeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"IgnoreNonces\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"}],\"name\":\"SkipNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTipAccount\",\"type\":\"address\"}],\"name\":\"TipAccountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTipRate\",\"type\":\"uint256\"}],\"name\":\"TipRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mappedChainid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mappedChainid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"CreatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getRoles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"describe\",\"type\":\"string\"}],\"internalType\":\"structRoleAccess.Role[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"omitNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"pauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeAccount\",\"type\":\"address\"}],\"name\":\"setFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newFiatCurrency\",\"type\":\"string\"}],\"name\":\"setFiatCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setFiatFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceoracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTipAccount\",\"type\":\"address\"}],\"name\":\"setTipAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTipRate\",\"type\":\"uint256\"}],\"name\":\"setTipRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingDepositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"unpauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mappedChainid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceTokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mappedTokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tipRate\",\"type\":\"uint256\"}],\"name\":\"setupTokenMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenChainid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sourceTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mappedTokenChainid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"mappedTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"internalType\":\"structVaultX.tokenPairInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"depositNative\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"batchWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tipX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tipY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tipCashout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tipBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"addNoncesToOmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nonces\",\"type\":\"uint256[]\"}],\"name\":\"removeNoncesToOmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"skip\",\"type\":\"uint256\"}],\"name\":\"skipWithdrawWatermark\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultX is an auto generated Go binding around an Ethereum contract.
type VaultX struct {
	VaultXCaller     // Read-only binding to the contract
	VaultXTransactor // Write-only binding to the contract
	VaultXFilterer   // Log filterer for contract events
}

// VaultXCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultXSession struct {
	Contract     *VaultX           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultXCallerSession struct {
	Contract *VaultXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultXTransactorSession struct {
	Contract     *VaultXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultXRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultXRaw struct {
	Contract *VaultX // Generic contract binding to access the raw methods on
}

// VaultXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultXCallerRaw struct {
	Contract *VaultXCaller // Generic read-only contract binding to access the raw methods on
}

// VaultXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultXTransactorRaw struct {
	Contract *VaultXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultX creates a new instance of VaultX, bound to a specific deployed contract.
func NewVaultX(address common.Address, backend bind.ContractBackend) (*VaultX, error) {
	contract, err := bindVaultX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultX{VaultXCaller: VaultXCaller{contract: contract}, VaultXTransactor: VaultXTransactor{contract: contract}, VaultXFilterer: VaultXFilterer{contract: contract}}, nil
}

// NewVaultXCaller creates a new read-only instance of VaultX, bound to a specific deployed contract.
func NewVaultXCaller(address common.Address, caller bind.ContractCaller) (*VaultXCaller, error) {
	contract, err := bindVaultX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultXCaller{contract: contract}, nil
}

// NewVaultXTransactor creates a new write-only instance of VaultX, bound to a specific deployed contract.
func NewVaultXTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultXTransactor, error) {
	contract, err := bindVaultX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultXTransactor{contract: contract}, nil
}

// NewVaultXFilterer creates a new log filterer instance of VaultX, bound to a specific deployed contract.
func NewVaultXFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultXFilterer, error) {
	contract, err := bindVaultX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultXFilterer{contract: contract}, nil
}

// bindVaultX binds a generic wrapper to an already deployed contract.
func bindVaultX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultX *VaultXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultX.Contract.VaultXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultX *VaultXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.Contract.VaultXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultX *VaultXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultX.Contract.VaultXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultX *VaultXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultX *VaultXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultX *VaultXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultX.Contract.contract.Transact(opts, method, params...)
}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultX *VaultXCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "CreatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultX *VaultXSession) CreatedAt() (*big.Int, error) {
	return _VaultX.Contract.CreatedAt(&_VaultX.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0x748140a3.
//
// Solidity: function CreatedAt() view returns(uint256)
func (_VaultX *VaultXCallerSession) CreatedAt() (*big.Int, error) {
	return _VaultX.Contract.CreatedAt(&_VaultX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultX.Contract.DEFAULTADMINROLE(&_VaultX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultX.Contract.DEFAULTADMINROLE(&_VaultX.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultX *VaultXCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultX *VaultXSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultX.Contract.GetRoleAdmin(&_VaultX.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultX *VaultXCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultX.Contract.GetRoleAdmin(&_VaultX.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultX *VaultXCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultX *VaultXSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultX.Contract.GetRoleMember(&_VaultX.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultX *VaultXCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultX.Contract.GetRoleMember(&_VaultX.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultX *VaultXCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultX *VaultXSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultX.Contract.GetRoleMemberCount(&_VaultX.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultX *VaultXCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultX.Contract.GetRoleMemberCount(&_VaultX.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultX *VaultXCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultX *VaultXSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _VaultX.Contract.GetRoleMembers(&_VaultX.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_VaultX *VaultXCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _VaultX.Contract.GetRoleMembers(&_VaultX.CallOpts, role)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultX *VaultXCaller) GetRoles(opts *bind.CallOpts) ([]RoleAccessRole, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getRoles")

	if err != nil {
		return *new([]RoleAccessRole), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleAccessRole)).(*[]RoleAccessRole)

	return out0, err

}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultX *VaultXSession) GetRoles() ([]RoleAccessRole, error) {
	return _VaultX.Contract.GetRoles(&_VaultX.CallOpts)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_VaultX *VaultXCallerSession) GetRoles() ([]RoleAccessRole, error) {
	return _VaultX.Contract.GetRoles(&_VaultX.CallOpts)
}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultX *VaultXCaller) GetTip(opts *bind.CallOpts, sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getTip", sourceToken, mappedToken, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultX *VaultXSession) GetTip(sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	return _VaultX.Contract.GetTip(&_VaultX.CallOpts, sourceToken, mappedToken, amount)
}

// GetTip is a free data retrieval call binding the contract method 0xb58f745b.
//
// Solidity: function getTip(address sourceToken, address mappedToken, uint256 amount) view returns(uint256)
func (_VaultX *VaultXCallerSession) GetTip(sourceToken common.Address, mappedToken common.Address, amount *big.Int) (*big.Int, error) {
	return _VaultX.Contract.GetTip(&_VaultX.CallOpts, sourceToken, mappedToken, amount)
}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultX *VaultXCaller) GetTokenPairs(opts *bind.CallOpts) ([]VaultXtokenPairInfo, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getTokenPairs")

	if err != nil {
		return *new([]VaultXtokenPairInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]VaultXtokenPairInfo)).(*[]VaultXtokenPairInfo)

	return out0, err

}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultX *VaultXSession) GetTokenPairs() ([]VaultXtokenPairInfo, error) {
	return _VaultX.Contract.GetTokenPairs(&_VaultX.CallOpts)
}

// GetTokenPairs is a free data retrieval call binding the contract method 0xe24e4fdb.
//
// Solidity: function getTokenPairs() view returns((address,uint256,string,address,uint256,string,bool)[])
func (_VaultX *VaultXCallerSession) GetTokenPairs() ([]VaultXtokenPairInfo, error) {
	return _VaultX.Contract.GetTokenPairs(&_VaultX.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultX *VaultXCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultX *VaultXSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultX.Contract.HasRole(&_VaultX.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultX *VaultXCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultX.Contract.HasRole(&_VaultX.CallOpts, role, account)
}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultX *VaultXCaller) OmitNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "omitNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultX *VaultXSession) OmitNonces(arg0 *big.Int) (bool, error) {
	return _VaultX.Contract.OmitNonces(&_VaultX.CallOpts, arg0)
}

// OmitNonces is a free data retrieval call binding the contract method 0x1bcd5a4e.
//
// Solidity: function omitNonces(uint256 ) view returns(bool)
func (_VaultX *VaultXCallerSession) OmitNonces(arg0 *big.Int) (bool, error) {
	return _VaultX.Contract.OmitNonces(&_VaultX.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultX *VaultXCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultX *VaultXSession) Paused() (bool, error) {
	return _VaultX.Contract.Paused(&_VaultX.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultX *VaultXCallerSession) Paused() (bool, error) {
	return _VaultX.Contract.Paused(&_VaultX.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultX *VaultXCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultX *VaultXSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultX.Contract.SupportsInterface(&_VaultX.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultX *VaultXCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultX.Contract.SupportsInterface(&_VaultX.CallOpts, interfaceId)
}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultX *VaultXCaller) TipBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tipBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultX *VaultXSession) TipBalance(token common.Address) (*big.Int, error) {
	return _VaultX.Contract.TipBalance(&_VaultX.CallOpts, token)
}

// TipBalance is a free data retrieval call binding the contract method 0xefc03425.
//
// Solidity: function tipBalance(address token) view returns(uint256)
func (_VaultX *VaultXCallerSession) TipBalance(token common.Address) (*big.Int, error) {
	return _VaultX.Contract.TipBalance(&_VaultX.CallOpts, token)
}

// TokenMappingDepositNonce is a free data retrieval call binding the contract method 0x356caeef.
//
// Solidity: function tokenMappingDepositNonce(address , address ) view returns(uint256)
func (_VaultX *VaultXCaller) TokenMappingDepositNonce(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMappingDepositNonce", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMappingDepositNonce is a free data retrieval call binding the contract method 0x356caeef.
//
// Solidity: function tokenMappingDepositNonce(address , address ) view returns(uint256)
func (_VaultX *VaultXSession) TokenMappingDepositNonce(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultX.Contract.TokenMappingDepositNonce(&_VaultX.CallOpts, arg0, arg1)
}

// TokenMappingDepositNonce is a free data retrieval call binding the contract method 0x356caeef.
//
// Solidity: function tokenMappingDepositNonce(address , address ) view returns(uint256)
func (_VaultX *VaultXCallerSession) TokenMappingDepositNonce(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultX.Contract.TokenMappingDepositNonce(&_VaultX.CallOpts, arg0, arg1)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultX *VaultXCaller) TokenMappingWatermark(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMappingWatermark", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultX *VaultXSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultX.Contract.TokenMappingWatermark(&_VaultX.CallOpts, arg0, arg1)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultX *VaultXCallerSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultX.Contract.TokenMappingWatermark(&_VaultX.CallOpts, arg0, arg1)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXTransactor) AddNoncesToOmit(opts *bind.TransactOpts, nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "addNoncesToOmit", nonces)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXSession) AddNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.AddNoncesToOmit(&_VaultX.TransactOpts, nonces)
}

// AddNoncesToOmit is a paid mutator transaction binding the contract method 0x379e271e.
//
// Solidity: function addNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXTransactorSession) AddNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.AddNoncesToOmit(&_VaultX.TransactOpts, nonces)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXTransactor) AddRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "addRoleMember", role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.AddRoleMember(&_VaultX.TransactOpts, role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXTransactorSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.AddRoleMember(&_VaultX.TransactOpts, role, member)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0xe1554b91.
//
// Solidity: function batchWithdraw(bytes signature, bytes input) returns()
func (_VaultX *VaultXTransactor) BatchWithdraw(opts *bind.TransactOpts, signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "batchWithdraw", signature, input)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0xe1554b91.
//
// Solidity: function batchWithdraw(bytes signature, bytes input) returns()
func (_VaultX *VaultXSession) BatchWithdraw(signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultX.Contract.BatchWithdraw(&_VaultX.TransactOpts, signature, input)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0xe1554b91.
//
// Solidity: function batchWithdraw(bytes signature, bytes input) returns()
func (_VaultX *VaultXTransactorSession) BatchWithdraw(signature []byte, input []byte) (*types.Transaction, error) {
	return _VaultX.Contract.BatchWithdraw(&_VaultX.TransactOpts, signature, input)
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns(bool)
func (_VaultX *VaultXTransactor) DepositNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "depositNative")
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns(bool)
func (_VaultX *VaultXSession) DepositNative() (*types.Transaction, error) {
	return _VaultX.Contract.DepositNative(&_VaultX.TransactOpts)
}

// DepositNative is a paid mutator transaction binding the contract method 0xdb6b5246.
//
// Solidity: function depositNative() payable returns(bool)
func (_VaultX *VaultXTransactorSession) DepositNative() (*types.Transaction, error) {
	return _VaultX.Contract.DepositNative(&_VaultX.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address sourceToken, uint256 amount) returns(bool)
func (_VaultX *VaultXTransactor) DepositToken(opts *bind.TransactOpts, sourceToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "depositToken", sourceToken, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address sourceToken, uint256 amount) returns(bool)
func (_VaultX *VaultXSession) DepositToken(sourceToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.DepositToken(&_VaultX.TransactOpts, sourceToken, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(address sourceToken, uint256 amount) returns(bool)
func (_VaultX *VaultXTransactorSession) DepositToken(sourceToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.DepositToken(&_VaultX.TransactOpts, sourceToken, amount)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultX *VaultXTransactor) GrantMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "grantMinter", minter)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultX *VaultXSession) GrantMinter(minter common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.GrantMinter(&_VaultX.TransactOpts, minter)
}

// GrantMinter is a paid mutator transaction binding the contract method 0x261707fa.
//
// Solidity: function grantMinter(address minter) returns(bool)
func (_VaultX *VaultXTransactorSession) GrantMinter(minter common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.GrantMinter(&_VaultX.TransactOpts, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultX *VaultXSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.GrantRole(&_VaultX.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.GrantRole(&_VaultX.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultX *VaultXTransactor) Mint(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "mint", sourceToken, mappedToken, to, amount, tipX, nonce)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultX *VaultXSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Mint(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, tipX, nonce)
}

// Mint is a paid mutator transaction binding the contract method 0x3e401bf8.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipX, uint256 nonce) returns()
func (_VaultX *VaultXTransactorSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipX *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Mint(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, tipX, nonce)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultX *VaultXTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultX *VaultXSession) PauseAll() (*types.Transaction, error) {
	return _VaultX.Contract.PauseAll(&_VaultX.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultX *VaultXTransactorSession) PauseAll() (*types.Transaction, error) {
	return _VaultX.Contract.PauseAll(&_VaultX.TransactOpts)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXTransactor) PauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "pauseTokenMapping", sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.PauseTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXTransactorSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.PauseTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXTransactor) Refund(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "refund", token, to, amount)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXSession) Refund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Refund(&_VaultX.TransactOpts, token, to, amount)
}

// Refund is a paid mutator transaction binding the contract method 0x82ad6f35.
//
// Solidity: function refund(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXTransactorSession) Refund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Refund(&_VaultX.TransactOpts, token, to, amount)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXTransactor) RemoveNoncesToOmit(opts *bind.TransactOpts, nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "removeNoncesToOmit", nonces)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXSession) RemoveNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveNoncesToOmit(&_VaultX.TransactOpts, nonces)
}

// RemoveNoncesToOmit is a paid mutator transaction binding the contract method 0xcb79b053.
//
// Solidity: function removeNoncesToOmit(uint256[] nonces) returns()
func (_VaultX *VaultXTransactorSession) RemoveNoncesToOmit(nonces []*big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveNoncesToOmit(&_VaultX.TransactOpts, nonces)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXTransactor) RemoveRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "removeRoleMember", role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveRoleMember(&_VaultX.TransactOpts, role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_VaultX *VaultXTransactorSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveRoleMember(&_VaultX.TransactOpts, role, member)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultX *VaultXSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RenounceRole(&_VaultX.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RenounceRole(&_VaultX.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultX *VaultXSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RevokeRole(&_VaultX.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultX *VaultXTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RevokeRole(&_VaultX.TransactOpts, role, account)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultX *VaultXTransactor) SetFeeAccount(opts *bind.TransactOpts, newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setFeeAccount", newFeeAccount)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultX *VaultXSession) SetFeeAccount(newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetFeeAccount(&_VaultX.TransactOpts, newFeeAccount)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address newFeeAccount) returns()
func (_VaultX *VaultXTransactorSession) SetFeeAccount(newFeeAccount common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetFeeAccount(&_VaultX.TransactOpts, newFeeAccount)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultX *VaultXTransactor) SetFiatCurrency(opts *bind.TransactOpts, newFiatCurrency string) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setFiatCurrency", newFiatCurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultX *VaultXSession) SetFiatCurrency(newFiatCurrency string) (*types.Transaction, error) {
	return _VaultX.Contract.SetFiatCurrency(&_VaultX.TransactOpts, newFiatCurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string newFiatCurrency) returns()
func (_VaultX *VaultXTransactorSession) SetFiatCurrency(newFiatCurrency string) (*types.Transaction, error) {
	return _VaultX.Contract.SetFiatCurrency(&_VaultX.TransactOpts, newFiatCurrency)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultX *VaultXTransactor) SetFiatFeeAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setFiatFeeAmount", amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultX *VaultXSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetFiatFeeAmount(&_VaultX.TransactOpts, amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultX *VaultXTransactorSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetFiatFeeAmount(&_VaultX.TransactOpts, amount)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultX *VaultXTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setPriceOracle", newPriceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultX *VaultXSession) SetPriceOracle(newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetPriceOracle(&_VaultX.TransactOpts, newPriceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceoracle) returns()
func (_VaultX *VaultXTransactorSession) SetPriceOracle(newPriceoracle common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetPriceOracle(&_VaultX.TransactOpts, newPriceoracle)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultX *VaultXTransactor) SetTipAccount(opts *bind.TransactOpts, newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setTipAccount", newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultX *VaultXSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetTipAccount(&_VaultX.TransactOpts, newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultX *VaultXTransactorSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetTipAccount(&_VaultX.TransactOpts, newTipAccount)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultX *VaultXTransactor) SetTipRate(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setTipRate", sourceToken, mappedToken, newTipRate)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultX *VaultXSession) SetTipRate(sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetTipRate(&_VaultX.TransactOpts, sourceToken, mappedToken, newTipRate)
}

// SetTipRate is a paid mutator transaction binding the contract method 0xda78defd.
//
// Solidity: function setTipRate(address sourceToken, address mappedToken, uint256 newTipRate) returns()
func (_VaultX *VaultXTransactorSession) SetTipRate(sourceToken common.Address, mappedToken common.Address, newTipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetTipRate(&_VaultX.TransactOpts, sourceToken, mappedToken, newTipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 mappedChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultX *VaultXTransactor) SetupTokenMapping(opts *bind.TransactOpts, mappedChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setupTokenMapping", mappedChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 mappedChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultX *VaultXSession) SetupTokenMapping(mappedChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetupTokenMapping(&_VaultX.TransactOpts, mappedChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9c0d9230.
//
// Solidity: function setupTokenMapping(uint256 mappedChainid, address sourceToken, address mappedToken, string sourceTokenSymbol_, string mappedTokenSymbol_, uint256 tipRate) returns(bool)
func (_VaultX *VaultXTransactorSession) SetupTokenMapping(mappedChainid *big.Int, sourceToken common.Address, mappedToken common.Address, sourceTokenSymbol_ string, mappedTokenSymbol_ string, tipRate *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SetupTokenMapping(&_VaultX.TransactOpts, mappedChainid, sourceToken, mappedToken, sourceTokenSymbol_, mappedTokenSymbol_, tipRate)
}

// SkipWithdrawWatermark is a paid mutator transaction binding the contract method 0x12ffc05d.
//
// Solidity: function skipWithdrawWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultX *VaultXTransactor) SkipWithdrawWatermark(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "skipWithdrawWatermark", sourceToken, mappedToken, skip)
}

// SkipWithdrawWatermark is a paid mutator transaction binding the contract method 0x12ffc05d.
//
// Solidity: function skipWithdrawWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultX *VaultXSession) SkipWithdrawWatermark(sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SkipWithdrawWatermark(&_VaultX.TransactOpts, sourceToken, mappedToken, skip)
}

// SkipWithdrawWatermark is a paid mutator transaction binding the contract method 0x12ffc05d.
//
// Solidity: function skipWithdrawWatermark(address sourceToken, address mappedToken, uint256 skip) returns()
func (_VaultX *VaultXTransactorSession) SkipWithdrawWatermark(sourceToken common.Address, mappedToken common.Address, skip *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SkipWithdrawWatermark(&_VaultX.TransactOpts, sourceToken, mappedToken, skip)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXTransactor) TipCashout(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "tipCashout", token, to, amount)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXSession) TipCashout(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.TipCashout(&_VaultX.TransactOpts, token, to, amount)
}

// TipCashout is a paid mutator transaction binding the contract method 0x1a46b010.
//
// Solidity: function tipCashout(address token, address to, uint256 amount) returns()
func (_VaultX *VaultXTransactorSession) TipCashout(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.TipCashout(&_VaultX.TransactOpts, token, to, amount)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultX *VaultXTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultX *VaultXSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultX.Contract.UnpauseAll(&_VaultX.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultX *VaultXTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultX.Contract.UnpauseAll(&_VaultX.TransactOpts)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXTransactor) UnpauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "unpauseTokenMapping", sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.UnpauseTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultX *VaultXTransactorSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.UnpauseTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultX *VaultXTransactor) Withdraw(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "withdraw", sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultX *VaultXSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Withdraw(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0xac6f7e3b.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 tipY, uint256 nonce) returns()
func (_VaultX *VaultXTransactorSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, tipY *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Withdraw(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, tipY, nonce)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultX *VaultXTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VaultX.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultX *VaultXSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultX.Contract.Fallback(&_VaultX.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultX *VaultXTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultX.Contract.Fallback(&_VaultX.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultX *VaultXTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultX.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultX *VaultXSession) Receive() (*types.Transaction, error) {
	return _VaultX.Contract.Receive(&_VaultX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultX *VaultXTransactorSession) Receive() (*types.Transaction, error) {
	return _VaultX.Contract.Receive(&_VaultX.TransactOpts)
}

// VaultXFeeAccountChangedIterator is returned from FilterFeeAccountChanged and is used to iterate over the raw logs and unpacked data for FeeAccountChanged events raised by the VaultX contract.
type VaultXFeeAccountChangedIterator struct {
	Event *VaultXFeeAccountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXFeeAccountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXFeeAccountChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXFeeAccountChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXFeeAccountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXFeeAccountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXFeeAccountChanged represents a FeeAccountChanged event raised by the VaultX contract.
type VaultXFeeAccountChanged struct {
	Sender        common.Address
	NewFeeAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeAccountChanged is a free log retrieval operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultX *VaultXFilterer) FilterFeeAccountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXFeeAccountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "FeeAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXFeeAccountChangedIterator{contract: _VaultX.contract, event: "FeeAccountChanged", logs: logs, sub: sub}, nil
}

// WatchFeeAccountChanged is a free log subscription operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultX *VaultXFilterer) WatchFeeAccountChanged(opts *bind.WatchOpts, sink chan<- *VaultXFeeAccountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "FeeAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXFeeAccountChanged)
				if err := _VaultX.contract.UnpackLog(event, "FeeAccountChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeAccountChanged is a log parse operation binding the contract event 0x508dcd584678b460e90a73b67f0c71c3c9eb2506cc1100a01ec04f6a19fac648.
//
// Solidity: event FeeAccountChanged(address indexed sender, address newFeeAccount)
func (_VaultX *VaultXFilterer) ParseFeeAccountChanged(log types.Log) (*VaultXFeeAccountChanged, error) {
	event := new(VaultXFeeAccountChanged)
	if err := _VaultX.contract.UnpackLog(event, "FeeAccountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXFiatCurrencyChangedIterator is returned from FilterFiatCurrencyChanged and is used to iterate over the raw logs and unpacked data for FiatCurrencyChanged events raised by the VaultX contract.
type VaultXFiatCurrencyChangedIterator struct {
	Event *VaultXFiatCurrencyChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXFiatCurrencyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXFiatCurrencyChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXFiatCurrencyChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXFiatCurrencyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXFiatCurrencyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXFiatCurrencyChanged represents a FiatCurrencyChanged event raised by the VaultX contract.
type VaultXFiatCurrencyChanged struct {
	Sender          common.Address
	NewFiatCurrency string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFiatCurrencyChanged is a free log retrieval operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultX *VaultXFilterer) FilterFiatCurrencyChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXFiatCurrencyChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "FiatCurrencyChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXFiatCurrencyChangedIterator{contract: _VaultX.contract, event: "FiatCurrencyChanged", logs: logs, sub: sub}, nil
}

// WatchFiatCurrencyChanged is a free log subscription operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultX *VaultXFilterer) WatchFiatCurrencyChanged(opts *bind.WatchOpts, sink chan<- *VaultXFiatCurrencyChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "FiatCurrencyChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXFiatCurrencyChanged)
				if err := _VaultX.contract.UnpackLog(event, "FiatCurrencyChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFiatCurrencyChanged is a log parse operation binding the contract event 0x7ed7441c7c9c83f1fd4152cec0a268d64e2c8784a2a05edf810befb09eb2a328.
//
// Solidity: event FiatCurrencyChanged(address indexed sender, string newFiatCurrency)
func (_VaultX *VaultXFilterer) ParseFiatCurrencyChanged(log types.Log) (*VaultXFiatCurrencyChanged, error) {
	event := new(VaultXFiatCurrencyChanged)
	if err := _VaultX.contract.UnpackLog(event, "FiatCurrencyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXFiatFeeAmountChangedIterator is returned from FilterFiatFeeAmountChanged and is used to iterate over the raw logs and unpacked data for FiatFeeAmountChanged events raised by the VaultX contract.
type VaultXFiatFeeAmountChangedIterator struct {
	Event *VaultXFiatFeeAmountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXFiatFeeAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXFiatFeeAmountChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXFiatFeeAmountChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXFiatFeeAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXFiatFeeAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXFiatFeeAmountChanged represents a FiatFeeAmountChanged event raised by the VaultX contract.
type VaultXFiatFeeAmountChanged struct {
	Sender           common.Address
	NewFiatFeeAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFiatFeeAmountChanged is a free log retrieval operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultX *VaultXFilterer) FilterFiatFeeAmountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXFiatFeeAmountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "FiatFeeAmountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXFiatFeeAmountChangedIterator{contract: _VaultX.contract, event: "FiatFeeAmountChanged", logs: logs, sub: sub}, nil
}

// WatchFiatFeeAmountChanged is a free log subscription operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultX *VaultXFilterer) WatchFiatFeeAmountChanged(opts *bind.WatchOpts, sink chan<- *VaultXFiatFeeAmountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "FiatFeeAmountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXFiatFeeAmountChanged)
				if err := _VaultX.contract.UnpackLog(event, "FiatFeeAmountChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFiatFeeAmountChanged is a log parse operation binding the contract event 0x2460d22a2a151edef78f3e2e7bf595d43fabf127aea7b47450f81fafef9d515b.
//
// Solidity: event FiatFeeAmountChanged(address indexed sender, uint256 newFiatFeeAmount)
func (_VaultX *VaultXFilterer) ParseFiatFeeAmountChanged(log types.Log) (*VaultXFiatFeeAmountChanged, error) {
	event := new(VaultXFiatFeeAmountChanged)
	if err := _VaultX.contract.UnpackLog(event, "FiatFeeAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXIgnoreNoncesIterator is returned from FilterIgnoreNonces and is used to iterate over the raw logs and unpacked data for IgnoreNonces events raised by the VaultX contract.
type VaultXIgnoreNoncesIterator struct {
	Event *VaultXIgnoreNonces // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXIgnoreNoncesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXIgnoreNonces)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXIgnoreNonces)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXIgnoreNoncesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXIgnoreNoncesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXIgnoreNonces represents a IgnoreNonces event raised by the VaultX contract.
type VaultXIgnoreNonces struct {
	Sender common.Address
	Nonces []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIgnoreNonces is a free log retrieval operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultX *VaultXFilterer) FilterIgnoreNonces(opts *bind.FilterOpts) (*VaultXIgnoreNoncesIterator, error) {

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "IgnoreNonces")
	if err != nil {
		return nil, err
	}
	return &VaultXIgnoreNoncesIterator{contract: _VaultX.contract, event: "IgnoreNonces", logs: logs, sub: sub}, nil
}

// WatchIgnoreNonces is a free log subscription operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultX *VaultXFilterer) WatchIgnoreNonces(opts *bind.WatchOpts, sink chan<- *VaultXIgnoreNonces) (event.Subscription, error) {

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "IgnoreNonces")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXIgnoreNonces)
				if err := _VaultX.contract.UnpackLog(event, "IgnoreNonces", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIgnoreNonces is a log parse operation binding the contract event 0x33fe1cf6879c44bf1b70173d9ba045fa144fdcaafc5135a64285cd6e8b18989c.
//
// Solidity: event IgnoreNonces(address sender, uint256[] nonces)
func (_VaultX *VaultXFilterer) ParseIgnoreNonces(log types.Log) (*VaultXIgnoreNonces, error) {
	event := new(VaultXIgnoreNonces)
	if err := _VaultX.contract.UnpackLog(event, "IgnoreNonces", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VaultX contract.
type VaultXPausedIterator struct {
	Event *VaultXPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXPaused represents a Paused event raised by the VaultX contract.
type VaultXPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultX *VaultXFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultXPausedIterator, error) {

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultXPausedIterator{contract: _VaultX.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultX *VaultXFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultXPaused) (event.Subscription, error) {

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXPaused)
				if err := _VaultX.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultX *VaultXFilterer) ParsePaused(log types.Log) (*VaultXPaused, error) {
	event := new(VaultXPaused)
	if err := _VaultX.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXPriceOracleChangedIterator is returned from FilterPriceOracleChanged and is used to iterate over the raw logs and unpacked data for PriceOracleChanged events raised by the VaultX contract.
type VaultXPriceOracleChangedIterator struct {
	Event *VaultXPriceOracleChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXPriceOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXPriceOracleChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXPriceOracleChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXPriceOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXPriceOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXPriceOracleChanged represents a PriceOracleChanged event raised by the VaultX contract.
type VaultXPriceOracleChanged struct {
	Sender         common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleChanged is a free log retrieval operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultX *VaultXFilterer) FilterPriceOracleChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXPriceOracleChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "PriceOracleChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXPriceOracleChangedIterator{contract: _VaultX.contract, event: "PriceOracleChanged", logs: logs, sub: sub}, nil
}

// WatchPriceOracleChanged is a free log subscription operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultX *VaultXFilterer) WatchPriceOracleChanged(opts *bind.WatchOpts, sink chan<- *VaultXPriceOracleChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "PriceOracleChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXPriceOracleChanged)
				if err := _VaultX.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceOracleChanged is a log parse operation binding the contract event 0x40bddd72ea96b80dae14e3d13e8ce2c4ecd2500d88c6c0004d24a00deab28f9c.
//
// Solidity: event PriceOracleChanged(address indexed sender, address newPriceOracle)
func (_VaultX *VaultXFilterer) ParsePriceOracleChanged(log types.Log) (*VaultXPriceOracleChanged, error) {
	event := new(VaultXPriceOracleChanged)
	if err := _VaultX.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the VaultX contract.
type VaultXRefundIterator struct {
	Event *VaultXRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXRefund represents a Refund event raised by the VaultX contract.
type VaultXRefund struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultX *VaultXFilterer) FilterRefund(opts *bind.FilterOpts) (*VaultXRefundIterator, error) {

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &VaultXRefundIterator{contract: _VaultX.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultX *VaultXFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *VaultXRefund) (event.Subscription, error) {

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXRefund)
				if err := _VaultX.contract.UnpackLog(event, "Refund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefund is a log parse operation binding the contract event 0xf40cc8c1a1d17359049ba500cfc894596a692cffc9d03943cd92ec2e159cf6ae.
//
// Solidity: event Refund(address token, address to, uint256 amount)
func (_VaultX *VaultXFilterer) ParseRefund(log types.Log) (*VaultXRefund, error) {
	event := new(VaultXRefund)
	if err := _VaultX.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VaultX contract.
type VaultXRoleAdminChangedIterator struct {
	Event *VaultXRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXRoleAdminChanged represents a RoleAdminChanged event raised by the VaultX contract.
type VaultXRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultX *VaultXFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultXRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultXRoleAdminChangedIterator{contract: _VaultX.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultX *VaultXFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultXRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXRoleAdminChanged)
				if err := _VaultX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultX *VaultXFilterer) ParseRoleAdminChanged(log types.Log) (*VaultXRoleAdminChanged, error) {
	event := new(VaultXRoleAdminChanged)
	if err := _VaultX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultX contract.
type VaultXRoleGrantedIterator struct {
	Event *VaultXRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXRoleGranted represents a RoleGranted event raised by the VaultX contract.
type VaultXRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultXRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXRoleGrantedIterator{contract: _VaultX.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultXRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXRoleGranted)
				if err := _VaultX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) ParseRoleGranted(log types.Log) (*VaultXRoleGranted, error) {
	event := new(VaultXRoleGranted)
	if err := _VaultX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultX contract.
type VaultXRoleRevokedIterator struct {
	Event *VaultXRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXRoleRevoked represents a RoleRevoked event raised by the VaultX contract.
type VaultXRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultXRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXRoleRevokedIterator{contract: _VaultX.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultXRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXRoleRevoked)
				if err := _VaultX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultX *VaultXFilterer) ParseRoleRevoked(log types.Log) (*VaultXRoleRevoked, error) {
	event := new(VaultXRoleRevoked)
	if err := _VaultX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXSkipNonceIterator is returned from FilterSkipNonce and is used to iterate over the raw logs and unpacked data for SkipNonce events raised by the VaultX contract.
type VaultXSkipNonceIterator struct {
	Event *VaultXSkipNonce // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXSkipNonceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXSkipNonce)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXSkipNonce)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXSkipNonceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXSkipNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXSkipNonce represents a SkipNonce event raised by the VaultX contract.
type VaultXSkipNonce struct {
	Start *big.Int
	Step  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSkipNonce is a free log retrieval operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultX *VaultXFilterer) FilterSkipNonce(opts *bind.FilterOpts) (*VaultXSkipNonceIterator, error) {

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "SkipNonce")
	if err != nil {
		return nil, err
	}
	return &VaultXSkipNonceIterator{contract: _VaultX.contract, event: "SkipNonce", logs: logs, sub: sub}, nil
}

// WatchSkipNonce is a free log subscription operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultX *VaultXFilterer) WatchSkipNonce(opts *bind.WatchOpts, sink chan<- *VaultXSkipNonce) (event.Subscription, error) {

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "SkipNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXSkipNonce)
				if err := _VaultX.contract.UnpackLog(event, "SkipNonce", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSkipNonce is a log parse operation binding the contract event 0x503d8507b9fedcafb1c6a87c33912dcfc181e9f19b17345e8b53c26b3e9fc29e.
//
// Solidity: event SkipNonce(uint256 start, uint256 step)
func (_VaultX *VaultXFilterer) ParseSkipNonce(log types.Log) (*VaultXSkipNonce, error) {
	event := new(VaultXSkipNonce)
	if err := _VaultX.contract.UnpackLog(event, "SkipNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the VaultX contract.
type VaultXStakeAddedIterator struct {
	Event *VaultXStakeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXStakeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXStakeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXStakeAdded represents a StakeAdded event raised by the VaultX contract.
type VaultXStakeAdded struct {
	From   common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) FilterStakeAdded(opts *bind.FilterOpts, from []common.Address, epoch []*big.Int) (*VaultXStakeAddedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "StakeAdded", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &VaultXStakeAddedIterator{contract: _VaultX.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *VaultXStakeAdded, from []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "StakeAdded", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXStakeAdded)
				if err := _VaultX.contract.UnpackLog(event, "StakeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeAdded is a log parse operation binding the contract event 0x270d6dd254edd1d985c81cf7861b8f28fb06b6d719df04d90464034d43412440.
//
// Solidity: event StakeAdded(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) ParseStakeAdded(log types.Log) (*VaultXStakeAdded, error) {
	event := new(VaultXStakeAdded)
	if err := _VaultX.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the VaultX contract.
type VaultXStakeRemovedIterator struct {
	Event *VaultXStakeRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXStakeRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXStakeRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXStakeRemoved represents a StakeRemoved event raised by the VaultX contract.
type VaultXStakeRemoved struct {
	From   common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) FilterStakeRemoved(opts *bind.FilterOpts, from []common.Address, epoch []*big.Int) (*VaultXStakeRemovedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "StakeRemoved", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &VaultXStakeRemovedIterator{contract: _VaultX.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *VaultXStakeRemoved, from []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "StakeRemoved", fromRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXStakeRemoved)
				if err := _VaultX.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeRemoved is a log parse operation binding the contract event 0xdcca95406ac9554449be02d88dcdf9c877f96e4c02bdad4bd5cadefc98a20e3d.
//
// Solidity: event StakeRemoved(address indexed from, uint256 indexed epoch, uint256 amount)
func (_VaultX *VaultXFilterer) ParseStakeRemoved(log types.Log) (*VaultXStakeRemoved, error) {
	event := new(VaultXStakeRemoved)
	if err := _VaultX.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXTipAccountChangedIterator is returned from FilterTipAccountChanged and is used to iterate over the raw logs and unpacked data for TipAccountChanged events raised by the VaultX contract.
type VaultXTipAccountChangedIterator struct {
	Event *VaultXTipAccountChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXTipAccountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXTipAccountChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXTipAccountChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXTipAccountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXTipAccountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXTipAccountChanged represents a TipAccountChanged event raised by the VaultX contract.
type VaultXTipAccountChanged struct {
	Sender        common.Address
	NewTipAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTipAccountChanged is a free log retrieval operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultX *VaultXFilterer) FilterTipAccountChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXTipAccountChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TipAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTipAccountChangedIterator{contract: _VaultX.contract, event: "TipAccountChanged", logs: logs, sub: sub}, nil
}

// WatchTipAccountChanged is a free log subscription operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultX *VaultXFilterer) WatchTipAccountChanged(opts *bind.WatchOpts, sink chan<- *VaultXTipAccountChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TipAccountChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXTipAccountChanged)
				if err := _VaultX.contract.UnpackLog(event, "TipAccountChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipAccountChanged is a log parse operation binding the contract event 0x028dd631374b48803cfa88ca1e6693001c280e4ffa9e573bf1002723cb03ae15.
//
// Solidity: event TipAccountChanged(address indexed sender, address newTipAccount)
func (_VaultX *VaultXFilterer) ParseTipAccountChanged(log types.Log) (*VaultXTipAccountChanged, error) {
	event := new(VaultXTipAccountChanged)
	if err := _VaultX.contract.UnpackLog(event, "TipAccountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXTipRateChangedIterator is returned from FilterTipRateChanged and is used to iterate over the raw logs and unpacked data for TipRateChanged events raised by the VaultX contract.
type VaultXTipRateChangedIterator struct {
	Event *VaultXTipRateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXTipRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXTipRateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXTipRateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXTipRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXTipRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXTipRateChanged represents a TipRateChanged event raised by the VaultX contract.
type VaultXTipRateChanged struct {
	Sender      common.Address
	SourceToken common.Address
	MappedToken common.Address
	NewTipRate  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipRateChanged is a free log retrieval operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultX *VaultXFilterer) FilterTipRateChanged(opts *bind.FilterOpts, sender []common.Address) (*VaultXTipRateChangedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TipRateChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTipRateChangedIterator{contract: _VaultX.contract, event: "TipRateChanged", logs: logs, sub: sub}, nil
}

// WatchTipRateChanged is a free log subscription operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultX *VaultXFilterer) WatchTipRateChanged(opts *bind.WatchOpts, sink chan<- *VaultXTipRateChanged, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TipRateChanged", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXTipRateChanged)
				if err := _VaultX.contract.UnpackLog(event, "TipRateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipRateChanged is a log parse operation binding the contract event 0x7551456b2f0d91a3f37c1ac5ab2c19b7ed477502a8236b6e280061e05bd144a8.
//
// Solidity: event TipRateChanged(address indexed sender, address sourceToken, address mappedToken, uint256 newTipRate)
func (_VaultX *VaultXFilterer) ParseTipRateChanged(log types.Log) (*VaultXTipRateChanged, error) {
	event := new(VaultXTipRateChanged)
	if err := _VaultX.contract.UnpackLog(event, "TipRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXTokenBurnIterator is returned from FilterTokenBurn and is used to iterate over the raw logs and unpacked data for TokenBurn events raised by the VaultX contract.
type VaultXTokenBurnIterator struct {
	Event *VaultXTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXTokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXTokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXTokenBurn represents a TokenBurn event raised by the VaultX contract.
type VaultXTokenBurn struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBurn is a free log retrieval operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) FilterTokenBurn(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (*VaultXTokenBurnIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTokenBurnIterator{contract: _VaultX.contract, event: "TokenBurn", logs: logs, sub: sub}, nil
}

// WatchTokenBurn is a free log subscription operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) WatchTokenBurn(opts *bind.WatchOpts, sink chan<- *VaultXTokenBurn, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXTokenBurn)
				if err := _VaultX.contract.UnpackLog(event, "TokenBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenBurn is a log parse operation binding the contract event 0xe6064164b9ee4a11aab4f75dcadd3198373528b269d446c73ca48565eaf1c28f.
//
// Solidity: event TokenBurn(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) ParseTokenBurn(log types.Log) (*VaultXTokenBurn, error) {
	event := new(VaultXTokenBurn)
	if err := _VaultX.contract.UnpackLog(event, "TokenBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the VaultX contract.
type VaultXTokenDepositIterator struct {
	Event *VaultXTokenDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXTokenDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXTokenDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXTokenDeposit represents a TokenDeposit event raised by the VaultX contract.
type VaultXTokenDeposit struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) FilterTokenDeposit(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (*VaultXTokenDepositIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTokenDepositIterator{contract: _VaultX.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *VaultXTokenDeposit, sourceToken []common.Address, mappedToken []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXTokenDeposit)
				if err := _VaultX.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenDeposit is a log parse operation binding the contract event 0x41f138048b690e2a81e12b36d21c06d77c6a6ee2a5b4ed203690f54a4d5ca4bf.
//
// Solidity: event TokenDeposit(address vault, uint256 sourceChainid, uint256 mappedChainid, address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 tip, uint256 indexed nonce, uint256 blockNumber)
func (_VaultX *VaultXFilterer) ParseTokenDeposit(log types.Log) (*VaultXTokenDeposit, error) {
	event := new(VaultXTokenDeposit)
	if err := _VaultX.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VaultX contract.
type VaultXUnpausedIterator struct {
	Event *VaultXUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  moaccore.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultXUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultXUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultXUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXUnpaused represents a Unpaused event raised by the VaultX contract.
type VaultXUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultX *VaultXFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultXUnpausedIterator, error) {

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultXUnpausedIterator{contract: _VaultX.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultX *VaultXFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultXUnpaused) (event.Subscription, error) {

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXUnpaused)
				if err := _VaultX.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultX *VaultXFilterer) ParseUnpaused(log types.Log) (*VaultXUnpaused, error) {
	event := new(VaultXUnpaused)
	if err := _VaultX.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
