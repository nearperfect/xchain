// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xdefi

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

// VaultYABI is the input ABI used to generate the binding from.
const VaultYABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ExitNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenBalanceAfter\",\"type\":\"uint256\"}],\"name\":\"TokenBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenBalanceAfter\",\"type\":\"uint256\"}],\"name\":\"TokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fiatCurrency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fiatFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stagingAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingExitNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMappingMintdone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingReversed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenStagingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWithdrawFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTipAccount\",\"type\":\"address\"}],\"name\":\"setTipAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fiatcurrency\",\"type\":\"string\"}],\"name\":\"setFiatCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setFiatFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nativetoken\",\"type\":\"string\"}],\"name\":\"setNativeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceoracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stagingaccount\",\"type\":\"address\"}],\"name\":\"setStagingAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"pauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"unpauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"setupTokenMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintNonce\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintNonce\",\"type\":\"uint256\"}],\"name\":\"skipMintdone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultY is an auto generated Go binding around an Ethereum contract.
type VaultY struct {
	VaultYCaller     // Read-only binding to the contract
	VaultYTransactor // Write-only binding to the contract
	VaultYFilterer   // Log filterer for contract events
}

// VaultYCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultYSession struct {
	Contract     *VaultY           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultYCallerSession struct {
	Contract *VaultYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultYTransactorSession struct {
	Contract     *VaultYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultYRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultYRaw struct {
	Contract *VaultY // Generic contract binding to access the raw methods on
}

// VaultYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultYCallerRaw struct {
	Contract *VaultYCaller // Generic read-only contract binding to access the raw methods on
}

// VaultYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultYTransactorRaw struct {
	Contract *VaultYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultY creates a new instance of VaultY, bound to a specific deployed contract.
func NewVaultY(address common.Address, backend bind.ContractBackend) (*VaultY, error) {
	contract, err := bindVaultY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultY{VaultYCaller: VaultYCaller{contract: contract}, VaultYTransactor: VaultYTransactor{contract: contract}, VaultYFilterer: VaultYFilterer{contract: contract}}, nil
}

// NewVaultYCaller creates a new read-only instance of VaultY, bound to a specific deployed contract.
func NewVaultYCaller(address common.Address, caller bind.ContractCaller) (*VaultYCaller, error) {
	contract, err := bindVaultY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultYCaller{contract: contract}, nil
}

// NewVaultYTransactor creates a new write-only instance of VaultY, bound to a specific deployed contract.
func NewVaultYTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultYTransactor, error) {
	contract, err := bindVaultY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultYTransactor{contract: contract}, nil
}

// NewVaultYFilterer creates a new log filterer instance of VaultY, bound to a specific deployed contract.
func NewVaultYFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultYFilterer, error) {
	contract, err := bindVaultY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultYFilterer{contract: contract}, nil
}

// bindVaultY binds a generic wrapper to an already deployed contract.
func bindVaultY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultY *VaultYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultY.Contract.VaultYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultY *VaultYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.Contract.VaultYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultY *VaultYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultY.Contract.VaultYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultY *VaultYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultY *VaultYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultY *VaultYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultY.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYSession) ADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.ADMINROLE(&_VaultY.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCallerSession) ADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.ADMINROLE(&_VaultY.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.DEFAULTADMINROLE(&_VaultY.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultY *VaultYCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultY.Contract.DEFAULTADMINROLE(&_VaultY.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultY *VaultYCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultY *VaultYSession) VALIDATORROLE() ([32]byte, error) {
	return _VaultY.Contract.VALIDATORROLE(&_VaultY.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultY *VaultYCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _VaultY.Contract.VALIDATORROLE(&_VaultY.CallOpts)
}

// FiatCurrency is a free data retrieval call binding the contract method 0x0d7b69e2.
//
// Solidity: function fiatCurrency() view returns(string)
func (_VaultY *VaultYCaller) FiatCurrency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "fiatCurrency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FiatCurrency is a free data retrieval call binding the contract method 0x0d7b69e2.
//
// Solidity: function fiatCurrency() view returns(string)
func (_VaultY *VaultYSession) FiatCurrency() (string, error) {
	return _VaultY.Contract.FiatCurrency(&_VaultY.CallOpts)
}

// FiatCurrency is a free data retrieval call binding the contract method 0x0d7b69e2.
//
// Solidity: function fiatCurrency() view returns(string)
func (_VaultY *VaultYCallerSession) FiatCurrency() (string, error) {
	return _VaultY.Contract.FiatCurrency(&_VaultY.CallOpts)
}

// FiatFeeAmount is a free data retrieval call binding the contract method 0x7b74400a.
//
// Solidity: function fiatFeeAmount() view returns(uint256)
func (_VaultY *VaultYCaller) FiatFeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "fiatFeeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FiatFeeAmount is a free data retrieval call binding the contract method 0x7b74400a.
//
// Solidity: function fiatFeeAmount() view returns(uint256)
func (_VaultY *VaultYSession) FiatFeeAmount() (*big.Int, error) {
	return _VaultY.Contract.FiatFeeAmount(&_VaultY.CallOpts)
}

// FiatFeeAmount is a free data retrieval call binding the contract method 0x7b74400a.
//
// Solidity: function fiatFeeAmount() view returns(uint256)
func (_VaultY *VaultYCallerSession) FiatFeeAmount() (*big.Int, error) {
	return _VaultY.Contract.FiatFeeAmount(&_VaultY.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultY.Contract.GetRoleAdmin(&_VaultY.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultY *VaultYCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultY.Contract.GetRoleAdmin(&_VaultY.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultY.Contract.GetRoleMember(&_VaultY.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_VaultY *VaultYCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _VaultY.Contract.GetRoleMember(&_VaultY.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultY.Contract.GetRoleMemberCount(&_VaultY.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_VaultY *VaultYCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _VaultY.Contract.GetRoleMemberCount(&_VaultY.CallOpts, role)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultY *VaultYCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultY *VaultYSession) GetValidators() ([]common.Address, error) {
	return _VaultY.Contract.GetValidators(&_VaultY.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultY *VaultYCallerSession) GetValidators() ([]common.Address, error) {
	return _VaultY.Contract.GetValidators(&_VaultY.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultY.Contract.HasRole(&_VaultY.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultY *VaultYCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultY.Contract.HasRole(&_VaultY.CallOpts, role, account)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(string)
func (_VaultY *VaultYCaller) NativeToken(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(string)
func (_VaultY *VaultYSession) NativeToken() (string, error) {
	return _VaultY.Contract.NativeToken(&_VaultY.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(string)
func (_VaultY *VaultYCallerSession) NativeToken() (string, error) {
	return _VaultY.Contract.NativeToken(&_VaultY.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYSession) Paused() (bool, error) {
	return _VaultY.Contract.Paused(&_VaultY.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VaultY *VaultYCallerSession) Paused() (bool, error) {
	return _VaultY.Contract.Paused(&_VaultY.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_VaultY *VaultYCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_VaultY *VaultYSession) PriceOracle() (common.Address, error) {
	return _VaultY.Contract.PriceOracle(&_VaultY.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_VaultY *VaultYCallerSession) PriceOracle() (common.Address, error) {
	return _VaultY.Contract.PriceOracle(&_VaultY.CallOpts)
}

// StagingAccount is a free data retrieval call binding the contract method 0xabd375f8.
//
// Solidity: function stagingAccount() view returns(address)
func (_VaultY *VaultYCaller) StagingAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "stagingAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StagingAccount is a free data retrieval call binding the contract method 0xabd375f8.
//
// Solidity: function stagingAccount() view returns(address)
func (_VaultY *VaultYSession) StagingAccount() (common.Address, error) {
	return _VaultY.Contract.StagingAccount(&_VaultY.CallOpts)
}

// StagingAccount is a free data retrieval call binding the contract method 0xabd375f8.
//
// Solidity: function stagingAccount() view returns(address)
func (_VaultY *VaultYCallerSession) StagingAccount() (common.Address, error) {
	return _VaultY.Contract.StagingAccount(&_VaultY.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultY.Contract.SupportsInterface(&_VaultY.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultY *VaultYCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultY.Contract.SupportsInterface(&_VaultY.CallOpts, interfaceId)
}

// TipAccount is a free data retrieval call binding the contract method 0x980be60f.
//
// Solidity: function tipAccount() view returns(address)
func (_VaultY *VaultYCaller) TipAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tipAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TipAccount is a free data retrieval call binding the contract method 0x980be60f.
//
// Solidity: function tipAccount() view returns(address)
func (_VaultY *VaultYSession) TipAccount() (common.Address, error) {
	return _VaultY.Contract.TipAccount(&_VaultY.CallOpts)
}

// TipAccount is a free data retrieval call binding the contract method 0x980be60f.
//
// Solidity: function tipAccount() view returns(address)
func (_VaultY *VaultYCallerSession) TipAccount() (common.Address, error) {
	return _VaultY.Contract.TipAccount(&_VaultY.CallOpts)
}

// TipRate is a free data retrieval call binding the contract method 0x498e76a0.
//
// Solidity: function tipRate() view returns(uint256)
func (_VaultY *VaultYCaller) TipRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tipRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TipRate is a free data retrieval call binding the contract method 0x498e76a0.
//
// Solidity: function tipRate() view returns(uint256)
func (_VaultY *VaultYSession) TipRate() (*big.Int, error) {
	return _VaultY.Contract.TipRate(&_VaultY.CallOpts)
}

// TipRate is a free data retrieval call binding the contract method 0x498e76a0.
//
// Solidity: function tipRate() view returns(uint256)
func (_VaultY *VaultYCallerSession) TipRate() (*big.Int, error) {
	return _VaultY.Contract.TipRate(&_VaultY.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultY *VaultYCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultY *VaultYSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _VaultY.Contract.TokenMapping(&_VaultY.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultY *VaultYCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _VaultY.Contract.TokenMapping(&_VaultY.CallOpts, arg0)
}

// TokenMappingExitNonce is a free data retrieval call binding the contract method 0xa7eb8d8c.
//
// Solidity: function tokenMappingExitNonce(address , address ) view returns(uint256)
func (_VaultY *VaultYCaller) TokenMappingExitNonce(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingExitNonce", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMappingExitNonce is a free data retrieval call binding the contract method 0xa7eb8d8c.
//
// Solidity: function tokenMappingExitNonce(address , address ) view returns(uint256)
func (_VaultY *VaultYSession) TokenMappingExitNonce(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingExitNonce(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingExitNonce is a free data retrieval call binding the contract method 0xa7eb8d8c.
//
// Solidity: function tokenMappingExitNonce(address , address ) view returns(uint256)
func (_VaultY *VaultYCallerSession) TokenMappingExitNonce(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingExitNonce(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingMintdone is a free data retrieval call binding the contract method 0x35028ff0.
//
// Solidity: function tokenMappingMintdone(address , address , uint256 ) view returns(bool)
func (_VaultY *VaultYCaller) TokenMappingMintdone(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingMintdone", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMappingMintdone is a free data retrieval call binding the contract method 0x35028ff0.
//
// Solidity: function tokenMappingMintdone(address , address , uint256 ) view returns(bool)
func (_VaultY *VaultYSession) TokenMappingMintdone(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _VaultY.Contract.TokenMappingMintdone(&_VaultY.CallOpts, arg0, arg1, arg2)
}

// TokenMappingMintdone is a free data retrieval call binding the contract method 0x35028ff0.
//
// Solidity: function tokenMappingMintdone(address , address , uint256 ) view returns(bool)
func (_VaultY *VaultYCallerSession) TokenMappingMintdone(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _VaultY.Contract.TokenMappingMintdone(&_VaultY.CallOpts, arg0, arg1, arg2)
}

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultY *VaultYCaller) TokenMappingPaused(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingPaused", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultY *VaultYSession) TokenMappingPaused(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultY.Contract.TokenMappingPaused(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultY *VaultYCallerSession) TokenMappingPaused(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultY.Contract.TokenMappingPaused(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultY *VaultYCaller) TokenMappingReversed(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingReversed", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultY *VaultYSession) TokenMappingReversed(arg0 common.Address) (common.Address, error) {
	return _VaultY.Contract.TokenMappingReversed(&_VaultY.CallOpts, arg0)
}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultY *VaultYCallerSession) TokenMappingReversed(arg0 common.Address) (common.Address, error) {
	return _VaultY.Contract.TokenMappingReversed(&_VaultY.CallOpts, arg0)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYCaller) TokenMappingWatermark(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenMappingWatermark", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingWatermark(&_VaultY.CallOpts, arg0, arg1)
}

// TokenMappingWatermark is a free data retrieval call binding the contract method 0x55275c6f.
//
// Solidity: function tokenMappingWatermark(address , address ) view returns(uint256)
func (_VaultY *VaultYCallerSession) TokenMappingWatermark(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenMappingWatermark(&_VaultY.CallOpts, arg0, arg1)
}

// TokenStagingBalances is a free data retrieval call binding the contract method 0x335701e4.
//
// Solidity: function tokenStagingBalances(address , address ) view returns(uint256)
func (_VaultY *VaultYCaller) TokenStagingBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenStagingBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenStagingBalances is a free data retrieval call binding the contract method 0x335701e4.
//
// Solidity: function tokenStagingBalances(address , address ) view returns(uint256)
func (_VaultY *VaultYSession) TokenStagingBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenStagingBalances(&_VaultY.CallOpts, arg0, arg1)
}

// TokenStagingBalances is a free data retrieval call binding the contract method 0x335701e4.
//
// Solidity: function tokenStagingBalances(address , address ) view returns(uint256)
func (_VaultY *VaultYCallerSession) TokenStagingBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenStagingBalances(&_VaultY.CallOpts, arg0, arg1)
}

// TokenWithdrawFees is a free data retrieval call binding the contract method 0xadc072aa.
//
// Solidity: function tokenWithdrawFees(address , address ) view returns(uint256)
func (_VaultY *VaultYCaller) TokenWithdrawFees(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultY.contract.Call(opts, &out, "tokenWithdrawFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWithdrawFees is a free data retrieval call binding the contract method 0xadc072aa.
//
// Solidity: function tokenWithdrawFees(address , address ) view returns(uint256)
func (_VaultY *VaultYSession) TokenWithdrawFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenWithdrawFees(&_VaultY.CallOpts, arg0, arg1)
}

// TokenWithdrawFees is a free data retrieval call binding the contract method 0xadc072aa.
//
// Solidity: function tokenWithdrawFees(address , address ) view returns(uint256)
func (_VaultY *VaultYCallerSession) TokenWithdrawFees(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VaultY.Contract.TokenWithdrawFees(&_VaultY.CallOpts, arg0, arg1)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultY *VaultYTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "addValidator", validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultY *VaultYSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.AddValidator(&_VaultY.TransactOpts, validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultY *VaultYTransactorSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.AddValidator(&_VaultY.TransactOpts, validator)
}

// CollectFee is a paid mutator transaction binding the contract method 0x2ec0ff6c.
//
// Solidity: function collectFee(address to, uint256 amount) returns()
func (_VaultY *VaultYTransactor) CollectFee(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "collectFee", to, amount)
}

// CollectFee is a paid mutator transaction binding the contract method 0x2ec0ff6c.
//
// Solidity: function collectFee(address to, uint256 amount) returns()
func (_VaultY *VaultYSession) CollectFee(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.CollectFee(&_VaultY.TransactOpts, to, amount)
}

// CollectFee is a paid mutator transaction binding the contract method 0x2ec0ff6c.
//
// Solidity: function collectFee(address to, uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) CollectFee(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.CollectFee(&_VaultY.TransactOpts, to, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address mappedToken, address from, uint256 amount) returns()
func (_VaultY *VaultYTransactor) Exit(opts *bind.TransactOpts, mappedToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "exit", mappedToken, from, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address mappedToken, address from, uint256 amount) returns()
func (_VaultY *VaultYSession) Exit(mappedToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Exit(&_VaultY.TransactOpts, mappedToken, from, amount)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address mappedToken, address from, uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) Exit(mappedToken common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Exit(&_VaultY.TransactOpts, mappedToken, from, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantRole(&_VaultY.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.GrantRole(&_VaultY.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xe9bca719.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 mintNonce) returns()
func (_VaultY *VaultYTransactor) Mint(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "mint", sourceToken, mappedToken, to, amount, mintNonce)
}

// Mint is a paid mutator transaction binding the contract method 0xe9bca719.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 mintNonce) returns()
func (_VaultY *VaultYSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Mint(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, mintNonce)
}

// Mint is a paid mutator transaction binding the contract method 0xe9bca719.
//
// Solidity: function mint(address sourceToken, address mappedToken, address to, uint256 amount, uint256 mintNonce) returns()
func (_VaultY *VaultYTransactorSession) Mint(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.Mint(&_VaultY.TransactOpts, sourceToken, mappedToken, to, amount, mintNonce)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYSession) PauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.PauseAll(&_VaultY.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_VaultY *VaultYTransactorSession) PauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.PauseAll(&_VaultY.TransactOpts)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactor) PauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "pauseTokenMapping", sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.PauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// PauseTokenMapping is a paid mutator transaction binding the contract method 0x0020d792.
//
// Solidity: function pauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactorSession) PauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.PauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultY *VaultYTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultY *VaultYSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveValidator(&_VaultY.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultY *VaultYTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RemoveValidator(&_VaultY.TransactOpts, validator)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RenounceRole(&_VaultY.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RenounceRole(&_VaultY.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RevokeRole(&_VaultY.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultY *VaultYTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.RevokeRole(&_VaultY.TransactOpts, role, account)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string fiatcurrency) returns()
func (_VaultY *VaultYTransactor) SetFiatCurrency(opts *bind.TransactOpts, fiatcurrency string) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setFiatCurrency", fiatcurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string fiatcurrency) returns()
func (_VaultY *VaultYSession) SetFiatCurrency(fiatcurrency string) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatCurrency(&_VaultY.TransactOpts, fiatcurrency)
}

// SetFiatCurrency is a paid mutator transaction binding the contract method 0x10804064.
//
// Solidity: function setFiatCurrency(string fiatcurrency) returns()
func (_VaultY *VaultYTransactorSession) SetFiatCurrency(fiatcurrency string) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatCurrency(&_VaultY.TransactOpts, fiatcurrency)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYTransactor) SetFiatFeeAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setFiatFeeAmount", amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatFeeAmount(&_VaultY.TransactOpts, amount)
}

// SetFiatFeeAmount is a paid mutator transaction binding the contract method 0xb7fac573.
//
// Solidity: function setFiatFeeAmount(uint256 amount) returns()
func (_VaultY *VaultYTransactorSession) SetFiatFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SetFiatFeeAmount(&_VaultY.TransactOpts, amount)
}

// SetNativeToken is a paid mutator transaction binding the contract method 0xb6600568.
//
// Solidity: function setNativeToken(string nativetoken) returns()
func (_VaultY *VaultYTransactor) SetNativeToken(opts *bind.TransactOpts, nativetoken string) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setNativeToken", nativetoken)
}

// SetNativeToken is a paid mutator transaction binding the contract method 0xb6600568.
//
// Solidity: function setNativeToken(string nativetoken) returns()
func (_VaultY *VaultYSession) SetNativeToken(nativetoken string) (*types.Transaction, error) {
	return _VaultY.Contract.SetNativeToken(&_VaultY.TransactOpts, nativetoken)
}

// SetNativeToken is a paid mutator transaction binding the contract method 0xb6600568.
//
// Solidity: function setNativeToken(string nativetoken) returns()
func (_VaultY *VaultYTransactorSession) SetNativeToken(nativetoken string) (*types.Transaction, error) {
	return _VaultY.Contract.SetNativeToken(&_VaultY.TransactOpts, nativetoken)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceoracle) returns()
func (_VaultY *VaultYTransactor) SetPriceOracle(opts *bind.TransactOpts, priceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setPriceOracle", priceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceoracle) returns()
func (_VaultY *VaultYSession) SetPriceOracle(priceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetPriceOracle(&_VaultY.TransactOpts, priceoracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceoracle) returns()
func (_VaultY *VaultYTransactorSession) SetPriceOracle(priceoracle common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetPriceOracle(&_VaultY.TransactOpts, priceoracle)
}

// SetStagingAccount is a paid mutator transaction binding the contract method 0xc3323798.
//
// Solidity: function setStagingAccount(address stagingaccount) returns()
func (_VaultY *VaultYTransactor) SetStagingAccount(opts *bind.TransactOpts, stagingaccount common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setStagingAccount", stagingaccount)
}

// SetStagingAccount is a paid mutator transaction binding the contract method 0xc3323798.
//
// Solidity: function setStagingAccount(address stagingaccount) returns()
func (_VaultY *VaultYSession) SetStagingAccount(stagingaccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetStagingAccount(&_VaultY.TransactOpts, stagingaccount)
}

// SetStagingAccount is a paid mutator transaction binding the contract method 0xc3323798.
//
// Solidity: function setStagingAccount(address stagingaccount) returns()
func (_VaultY *VaultYTransactorSession) SetStagingAccount(stagingaccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetStagingAccount(&_VaultY.TransactOpts, stagingaccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYTransactor) SetTipAccount(opts *bind.TransactOpts, newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setTipAccount", newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipAccount(&_VaultY.TransactOpts, newTipAccount)
}

// SetTipAccount is a paid mutator transaction binding the contract method 0x0cddaa22.
//
// Solidity: function setTipAccount(address newTipAccount) returns()
func (_VaultY *VaultYTransactorSession) SetTipAccount(newTipAccount common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetTipAccount(&_VaultY.TransactOpts, newTipAccount)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultY *VaultYTransactor) SetupTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "setupTokenMapping", sourceToken, mappedToken)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultY *VaultYSession) SetupTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetupTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultY *VaultYTransactorSession) SetupTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.SetupTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// SkipMintdone is a paid mutator transaction binding the contract method 0x6ed201c4.
//
// Solidity: function skipMintdone(address sourceToken, address mappedToken, uint256 mintNonce) returns()
func (_VaultY *VaultYTransactor) SkipMintdone(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "skipMintdone", sourceToken, mappedToken, mintNonce)
}

// SkipMintdone is a paid mutator transaction binding the contract method 0x6ed201c4.
//
// Solidity: function skipMintdone(address sourceToken, address mappedToken, uint256 mintNonce) returns()
func (_VaultY *VaultYSession) SkipMintdone(sourceToken common.Address, mappedToken common.Address, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SkipMintdone(&_VaultY.TransactOpts, sourceToken, mappedToken, mintNonce)
}

// SkipMintdone is a paid mutator transaction binding the contract method 0x6ed201c4.
//
// Solidity: function skipMintdone(address sourceToken, address mappedToken, uint256 mintNonce) returns()
func (_VaultY *VaultYTransactorSession) SkipMintdone(sourceToken common.Address, mappedToken common.Address, mintNonce *big.Int) (*types.Transaction, error) {
	return _VaultY.Contract.SkipMintdone(&_VaultY.TransactOpts, sourceToken, mappedToken, mintNonce)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseAll(&_VaultY.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_VaultY *VaultYTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseAll(&_VaultY.TransactOpts)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactor) UnpauseTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "unpauseTokenMapping", sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// UnpauseTokenMapping is a paid mutator transaction binding the contract method 0x7b32a34d.
//
// Solidity: function unpauseTokenMapping(address sourceToken, address mappedToken) returns()
func (_VaultY *VaultYTransactorSession) UnpauseTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.UnpauseTokenMapping(&_VaultY.TransactOpts, sourceToken, mappedToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address mappedToken, address to) payable returns()
func (_VaultY *VaultYTransactor) Withdraw(opts *bind.TransactOpts, mappedToken common.Address, to common.Address) (*types.Transaction, error) {
	return _VaultY.contract.Transact(opts, "withdraw", mappedToken, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address mappedToken, address to) payable returns()
func (_VaultY *VaultYSession) Withdraw(mappedToken common.Address, to common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.Withdraw(&_VaultY.TransactOpts, mappedToken, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address mappedToken, address to) payable returns()
func (_VaultY *VaultYTransactorSession) Withdraw(mappedToken common.Address, to common.Address) (*types.Transaction, error) {
	return _VaultY.Contract.Withdraw(&_VaultY.TransactOpts, mappedToken, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VaultY.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultY.Contract.Fallback(&_VaultY.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_VaultY *VaultYTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultY.Contract.Fallback(&_VaultY.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultY.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYSession) Receive() (*types.Transaction, error) {
	return _VaultY.Contract.Receive(&_VaultY.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VaultY *VaultYTransactorSession) Receive() (*types.Transaction, error) {
	return _VaultY.Contract.Receive(&_VaultY.TransactOpts)
}

// VaultYPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VaultY contract.
type VaultYPausedIterator struct {
	Event *VaultYPaused // Event containing the contract specifics and raw log

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
func (it *VaultYPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYPaused)
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
		it.Event = new(VaultYPaused)
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
func (it *VaultYPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYPaused represents a Paused event raised by the VaultY contract.
type VaultYPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultY *VaultYFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultYPausedIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultYPausedIterator{contract: _VaultY.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VaultY *VaultYFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultYPaused) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYPaused)
				if err := _VaultY.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VaultY *VaultYFilterer) ParsePaused(log types.Log) (*VaultYPaused, error) {
	event := new(VaultYPaused)
	if err := _VaultY.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VaultY contract.
type VaultYRoleAdminChangedIterator struct {
	Event *VaultYRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VaultYRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleAdminChanged)
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
		it.Event = new(VaultYRoleAdminChanged)
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
func (it *VaultYRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleAdminChanged represents a RoleAdminChanged event raised by the VaultY contract.
type VaultYRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultY *VaultYFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultYRoleAdminChangedIterator, error) {

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

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleAdminChangedIterator{contract: _VaultY.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultY *VaultYFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultYRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleAdminChanged)
				if err := _VaultY.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_VaultY *VaultYFilterer) ParseRoleAdminChanged(log types.Log) (*VaultYRoleAdminChanged, error) {
	event := new(VaultYRoleAdminChanged)
	if err := _VaultY.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultY contract.
type VaultYRoleGrantedIterator struct {
	Event *VaultYRoleGranted // Event containing the contract specifics and raw log

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
func (it *VaultYRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleGranted)
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
		it.Event = new(VaultYRoleGranted)
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
func (it *VaultYRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleGranted represents a RoleGranted event raised by the VaultY contract.
type VaultYRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultYRoleGrantedIterator, error) {

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

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleGrantedIterator{contract: _VaultY.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultYRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleGranted)
				if err := _VaultY.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VaultY *VaultYFilterer) ParseRoleGranted(log types.Log) (*VaultYRoleGranted, error) {
	event := new(VaultYRoleGranted)
	if err := _VaultY.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultY contract.
type VaultYRoleRevokedIterator struct {
	Event *VaultYRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VaultYRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYRoleRevoked)
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
		it.Event = new(VaultYRoleRevoked)
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
func (it *VaultYRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYRoleRevoked represents a RoleRevoked event raised by the VaultY contract.
type VaultYRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultYRoleRevokedIterator, error) {

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

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultYRoleRevokedIterator{contract: _VaultY.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultY *VaultYFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultYRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYRoleRevoked)
				if err := _VaultY.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VaultY *VaultYFilterer) ParseRoleRevoked(log types.Log) (*VaultYRoleRevoked, error) {
	event := new(VaultYRoleRevoked)
	if err := _VaultY.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTokenBurnIterator is returned from FilterTokenBurn and is used to iterate over the raw logs and unpacked data for TokenBurn events raised by the VaultY contract.
type VaultYTokenBurnIterator struct {
	Event *VaultYTokenBurn // Event containing the contract specifics and raw log

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
func (it *VaultYTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTokenBurn)
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
		it.Event = new(VaultYTokenBurn)
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
func (it *VaultYTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTokenBurn represents a TokenBurn event raised by the VaultY contract.
type VaultYTokenBurn struct {
	SourceToken       common.Address
	MappedToken       common.Address
	Account           common.Address
	Amount            *big.Int
	ExitNonce         *big.Int
	TokenBalanceAfter *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokenBurn is a free log retrieval operation binding the contract event 0x190f1ba8a56661f2810c1874c30137729454d745e2a4b66490906a074a4035fc.
//
// Solidity: event TokenBurn(address indexed sourceToken, address indexed mappedToken, address account, uint256 amount, uint256 indexed ExitNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) FilterTokenBurn(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, ExitNonce []*big.Int) (*VaultYTokenBurnIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var ExitNonceRule []interface{}
	for _, ExitNonceItem := range ExitNonce {
		ExitNonceRule = append(ExitNonceRule, ExitNonceItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, ExitNonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTokenBurnIterator{contract: _VaultY.contract, event: "TokenBurn", logs: logs, sub: sub}, nil
}

// WatchTokenBurn is a free log subscription operation binding the contract event 0x190f1ba8a56661f2810c1874c30137729454d745e2a4b66490906a074a4035fc.
//
// Solidity: event TokenBurn(address indexed sourceToken, address indexed mappedToken, address account, uint256 amount, uint256 indexed ExitNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) WatchTokenBurn(opts *bind.WatchOpts, sink chan<- *VaultYTokenBurn, sourceToken []common.Address, mappedToken []common.Address, ExitNonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var ExitNonceRule []interface{}
	for _, ExitNonceItem := range ExitNonce {
		ExitNonceRule = append(ExitNonceRule, ExitNonceItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TokenBurn", sourceTokenRule, mappedTokenRule, ExitNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTokenBurn)
				if err := _VaultY.contract.UnpackLog(event, "TokenBurn", log); err != nil {
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

// ParseTokenBurn is a log parse operation binding the contract event 0x190f1ba8a56661f2810c1874c30137729454d745e2a4b66490906a074a4035fc.
//
// Solidity: event TokenBurn(address indexed sourceToken, address indexed mappedToken, address account, uint256 amount, uint256 indexed ExitNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) ParseTokenBurn(log types.Log) (*VaultYTokenBurn, error) {
	event := new(VaultYTokenBurn)
	if err := _VaultY.contract.UnpackLog(event, "TokenBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYTokenMintIterator is returned from FilterTokenMint and is used to iterate over the raw logs and unpacked data for TokenMint events raised by the VaultY contract.
type VaultYTokenMintIterator struct {
	Event *VaultYTokenMint // Event containing the contract specifics and raw log

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
func (it *VaultYTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYTokenMint)
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
		it.Event = new(VaultYTokenMint)
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
func (it *VaultYTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYTokenMint represents a TokenMint event raised by the VaultY contract.
type VaultYTokenMint struct {
	SourceToken       common.Address
	MappedToken       common.Address
	To                common.Address
	Amount            *big.Int
	Tip               *big.Int
	DepositNonce      *big.Int
	TokenBalanceAfter *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokenMint is a free log retrieval operation binding the contract event 0xcb4365d1a29bcf52dd2a22528308a1958a0726bab3b9ace84581df181f3857b8.
//
// Solidity: event TokenMint(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 tip, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) FilterTokenMint(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, depositNonce []*big.Int) (*VaultYTokenMintIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "TokenMint", sourceTokenRule, mappedTokenRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultYTokenMintIterator{contract: _VaultY.contract, event: "TokenMint", logs: logs, sub: sub}, nil
}

// WatchTokenMint is a free log subscription operation binding the contract event 0xcb4365d1a29bcf52dd2a22528308a1958a0726bab3b9ace84581df181f3857b8.
//
// Solidity: event TokenMint(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 tip, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) WatchTokenMint(opts *bind.WatchOpts, sink chan<- *VaultYTokenMint, sourceToken []common.Address, mappedToken []common.Address, depositNonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "TokenMint", sourceTokenRule, mappedTokenRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYTokenMint)
				if err := _VaultY.contract.UnpackLog(event, "TokenMint", log); err != nil {
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

// ParseTokenMint is a log parse operation binding the contract event 0xcb4365d1a29bcf52dd2a22528308a1958a0726bab3b9ace84581df181f3857b8.
//
// Solidity: event TokenMint(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 tip, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultY *VaultYFilterer) ParseTokenMint(log types.Log) (*VaultYTokenMint, error) {
	event := new(VaultYTokenMint)
	if err := _VaultY.contract.UnpackLog(event, "TokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultYUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VaultY contract.
type VaultYUnpausedIterator struct {
	Event *VaultYUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultYUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultYUnpaused)
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
		it.Event = new(VaultYUnpaused)
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
func (it *VaultYUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultYUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultYUnpaused represents a Unpaused event raised by the VaultY contract.
type VaultYUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultY *VaultYFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultYUnpausedIterator, error) {

	logs, sub, err := _VaultY.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultYUnpausedIterator{contract: _VaultY.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VaultY *VaultYFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultYUnpaused) (event.Subscription, error) {

	logs, sub, err := _VaultY.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultYUnpaused)
				if err := _VaultY.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VaultY *VaultYFilterer) ParseUnpaused(log types.Log) (*VaultYUnpaused, error) {
	event := new(VaultYUnpaused)
	if err := _VaultY.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
