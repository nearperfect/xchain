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

// VaultXABI is the input ABI used to generate the binding from.
const VaultXABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenBalanceAfter\",\"type\":\"uint256\"}],\"name\":\"TokenDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"withdrawNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenBalanceAfter\",\"type\":\"uint256\"}],\"name\":\"TokenWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingDepositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingReversed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMappingWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMappingWithdrawdone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"pauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"unpauseTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"}],\"name\":\"setupTokenMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"depositNative\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawNonce\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawNonce\",\"type\":\"uint256\"}],\"name\":\"SkipWithdrawdone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXSession) ADMINROLE() ([32]byte, error) {
	return _VaultX.Contract.ADMINROLE(&_VaultX.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_VaultX *VaultXCallerSession) ADMINROLE() ([32]byte, error) {
	return _VaultX.Contract.ADMINROLE(&_VaultX.CallOpts)
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

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultX *VaultXCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultX *VaultXSession) VALIDATORROLE() ([32]byte, error) {
	return _VaultX.Contract.VALIDATORROLE(&_VaultX.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_VaultX *VaultXCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _VaultX.Contract.VALIDATORROLE(&_VaultX.CallOpts)
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

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultX *VaultXCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultX *VaultXSession) GetValidators() ([]common.Address, error) {
	return _VaultX.Contract.GetValidators(&_VaultX.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_VaultX *VaultXCallerSession) GetValidators() ([]common.Address, error) {
	return _VaultX.Contract.GetValidators(&_VaultX.CallOpts)
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

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultX *VaultXCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultX *VaultXSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _VaultX.Contract.TokenMapping(&_VaultX.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_VaultX *VaultXCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _VaultX.Contract.TokenMapping(&_VaultX.CallOpts, arg0)
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

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultX *VaultXCaller) TokenMappingPaused(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMappingPaused", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultX *VaultXSession) TokenMappingPaused(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultX.Contract.TokenMappingPaused(&_VaultX.CallOpts, arg0, arg1)
}

// TokenMappingPaused is a free data retrieval call binding the contract method 0x19955009.
//
// Solidity: function tokenMappingPaused(address , address ) view returns(bool)
func (_VaultX *VaultXCallerSession) TokenMappingPaused(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultX.Contract.TokenMappingPaused(&_VaultX.CallOpts, arg0, arg1)
}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultX *VaultXCaller) TokenMappingReversed(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMappingReversed", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultX *VaultXSession) TokenMappingReversed(arg0 common.Address) (common.Address, error) {
	return _VaultX.Contract.TokenMappingReversed(&_VaultX.CallOpts, arg0)
}

// TokenMappingReversed is a free data retrieval call binding the contract method 0x0804568e.
//
// Solidity: function tokenMappingReversed(address ) view returns(address)
func (_VaultX *VaultXCallerSession) TokenMappingReversed(arg0 common.Address) (common.Address, error) {
	return _VaultX.Contract.TokenMappingReversed(&_VaultX.CallOpts, arg0)
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

// TokenMappingWithdrawdone is a free data retrieval call binding the contract method 0x57a0076a.
//
// Solidity: function tokenMappingWithdrawdone(address , address , uint256 ) view returns(bool)
func (_VaultX *VaultXCaller) TokenMappingWithdrawdone(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "tokenMappingWithdrawdone", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMappingWithdrawdone is a free data retrieval call binding the contract method 0x57a0076a.
//
// Solidity: function tokenMappingWithdrawdone(address , address , uint256 ) view returns(bool)
func (_VaultX *VaultXSession) TokenMappingWithdrawdone(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _VaultX.Contract.TokenMappingWithdrawdone(&_VaultX.CallOpts, arg0, arg1, arg2)
}

// TokenMappingWithdrawdone is a free data retrieval call binding the contract method 0x57a0076a.
//
// Solidity: function tokenMappingWithdrawdone(address , address , uint256 ) view returns(bool)
func (_VaultX *VaultXCallerSession) TokenMappingWithdrawdone(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _VaultX.Contract.TokenMappingWithdrawdone(&_VaultX.CallOpts, arg0, arg1, arg2)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VaultX *VaultXCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultX.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VaultX *VaultXSession) WrappedNativeToken() (common.Address, error) {
	return _VaultX.Contract.WrappedNativeToken(&_VaultX.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VaultX *VaultXCallerSession) WrappedNativeToken() (common.Address, error) {
	return _VaultX.Contract.WrappedNativeToken(&_VaultX.CallOpts)
}

// SkipWithdrawdone is a paid mutator transaction binding the contract method 0x67f64877.
//
// Solidity: function SkipWithdrawdone(address sourceToken, address mappedToken, uint256 withdrawNonce) returns()
func (_VaultX *VaultXTransactor) SkipWithdrawdone(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "SkipWithdrawdone", sourceToken, mappedToken, withdrawNonce)
}

// SkipWithdrawdone is a paid mutator transaction binding the contract method 0x67f64877.
//
// Solidity: function SkipWithdrawdone(address sourceToken, address mappedToken, uint256 withdrawNonce) returns()
func (_VaultX *VaultXSession) SkipWithdrawdone(sourceToken common.Address, mappedToken common.Address, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SkipWithdrawdone(&_VaultX.TransactOpts, sourceToken, mappedToken, withdrawNonce)
}

// SkipWithdrawdone is a paid mutator transaction binding the contract method 0x67f64877.
//
// Solidity: function SkipWithdrawdone(address sourceToken, address mappedToken, uint256 withdrawNonce) returns()
func (_VaultX *VaultXTransactorSession) SkipWithdrawdone(sourceToken common.Address, mappedToken common.Address, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.SkipWithdrawdone(&_VaultX.TransactOpts, sourceToken, mappedToken, withdrawNonce)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultX *VaultXTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "addValidator", validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultX *VaultXSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.AddValidator(&_VaultX.TransactOpts, validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns(bool)
func (_VaultX *VaultXTransactorSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.AddValidator(&_VaultX.TransactOpts, validator)
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

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultX *VaultXTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultX *VaultXSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveValidator(&_VaultX.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns(bool)
func (_VaultX *VaultXTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.RemoveValidator(&_VaultX.TransactOpts, validator)
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

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultX *VaultXTransactor) SetupTokenMapping(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "setupTokenMapping", sourceToken, mappedToken)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultX *VaultXSession) SetupTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetupTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
}

// SetupTokenMapping is a paid mutator transaction binding the contract method 0x9812588b.
//
// Solidity: function setupTokenMapping(address sourceToken, address mappedToken) returns(bool)
func (_VaultX *VaultXTransactorSession) SetupTokenMapping(sourceToken common.Address, mappedToken common.Address) (*types.Transaction, error) {
	return _VaultX.Contract.SetupTokenMapping(&_VaultX.TransactOpts, sourceToken, mappedToken)
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

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 withdrawNonce) returns()
func (_VaultX *VaultXTransactor) Withdraw(opts *bind.TransactOpts, sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.contract.Transact(opts, "withdraw", sourceToken, mappedToken, to, amount, withdrawNonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 withdrawNonce) returns()
func (_VaultX *VaultXSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Withdraw(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, withdrawNonce)
}

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address sourceToken, address mappedToken, address to, uint256 amount, uint256 withdrawNonce) returns()
func (_VaultX *VaultXTransactorSession) Withdraw(sourceToken common.Address, mappedToken common.Address, to common.Address, amount *big.Int, withdrawNonce *big.Int) (*types.Transaction, error) {
	return _VaultX.Contract.Withdraw(&_VaultX.TransactOpts, sourceToken, mappedToken, to, amount, withdrawNonce)
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
	SourceToken       common.Address
	MappedToken       common.Address
	From              common.Address
	Amount            *big.Int
	DepositNonce      *big.Int
	TokenBalanceAfter *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0x8a70fb791a985ce0e8a56c0d1b95ad3a8cae91c7e17811afa33a2b1a8359d9a6.
//
// Solidity: event TokenDeposit(address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) FilterTokenDeposit(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, depositNonce []*big.Int) (*VaultXTokenDepositIterator, error) {

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

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTokenDepositIterator{contract: _VaultX.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0x8a70fb791a985ce0e8a56c0d1b95ad3a8cae91c7e17811afa33a2b1a8359d9a6.
//
// Solidity: event TokenDeposit(address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *VaultXTokenDeposit, sourceToken []common.Address, mappedToken []common.Address, depositNonce []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TokenDeposit", sourceTokenRule, mappedTokenRule, depositNonceRule)
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

// ParseTokenDeposit is a log parse operation binding the contract event 0x8a70fb791a985ce0e8a56c0d1b95ad3a8cae91c7e17811afa33a2b1a8359d9a6.
//
// Solidity: event TokenDeposit(address indexed sourceToken, address indexed mappedToken, address from, uint256 amount, uint256 indexed depositNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) ParseTokenDeposit(log types.Log) (*VaultXTokenDeposit, error) {
	event := new(VaultXTokenDeposit)
	if err := _VaultX.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultXTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the VaultX contract.
type VaultXTokenWithdrawIterator struct {
	Event *VaultXTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultXTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultXTokenWithdraw)
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
		it.Event = new(VaultXTokenWithdraw)
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
func (it *VaultXTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultXTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultXTokenWithdraw represents a TokenWithdraw event raised by the VaultX contract.
type VaultXTokenWithdraw struct {
	SourceToken       common.Address
	MappedToken       common.Address
	To                common.Address
	Amount            *big.Int
	WithdrawNonce     *big.Int
	TokenBalanceAfter *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x019749a15e6ce057e78b5de5437a98547a807f60cbba0574b5274f89cf2770ff.
//
// Solidity: event TokenWithdraw(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 indexed withdrawNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) FilterTokenWithdraw(opts *bind.FilterOpts, sourceToken []common.Address, mappedToken []common.Address, withdrawNonce []*big.Int) (*VaultXTokenWithdrawIterator, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var withdrawNonceRule []interface{}
	for _, withdrawNonceItem := range withdrawNonce {
		withdrawNonceRule = append(withdrawNonceRule, withdrawNonceItem)
	}

	logs, sub, err := _VaultX.contract.FilterLogs(opts, "TokenWithdraw", sourceTokenRule, mappedTokenRule, withdrawNonceRule)
	if err != nil {
		return nil, err
	}
	return &VaultXTokenWithdrawIterator{contract: _VaultX.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x019749a15e6ce057e78b5de5437a98547a807f60cbba0574b5274f89cf2770ff.
//
// Solidity: event TokenWithdraw(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 indexed withdrawNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *VaultXTokenWithdraw, sourceToken []common.Address, mappedToken []common.Address, withdrawNonce []*big.Int) (event.Subscription, error) {

	var sourceTokenRule []interface{}
	for _, sourceTokenItem := range sourceToken {
		sourceTokenRule = append(sourceTokenRule, sourceTokenItem)
	}
	var mappedTokenRule []interface{}
	for _, mappedTokenItem := range mappedToken {
		mappedTokenRule = append(mappedTokenRule, mappedTokenItem)
	}

	var withdrawNonceRule []interface{}
	for _, withdrawNonceItem := range withdrawNonce {
		withdrawNonceRule = append(withdrawNonceRule, withdrawNonceItem)
	}

	logs, sub, err := _VaultX.contract.WatchLogs(opts, "TokenWithdraw", sourceTokenRule, mappedTokenRule, withdrawNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultXTokenWithdraw)
				if err := _VaultX.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x019749a15e6ce057e78b5de5437a98547a807f60cbba0574b5274f89cf2770ff.
//
// Solidity: event TokenWithdraw(address indexed sourceToken, address indexed mappedToken, address to, uint256 amount, uint256 indexed withdrawNonce, uint256 tokenBalanceAfter)
func (_VaultX *VaultXFilterer) ParseTokenWithdraw(log types.Log) (*VaultXTokenWithdraw, error) {
	event := new(VaultXTokenWithdraw)
	if err := _VaultX.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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
