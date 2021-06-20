// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xevents

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

// XEventsABI is the input ABI used to generate the binding from.
const XEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"describe\",\"type\":\"string\"}],\"internalType\":\"structRoleAccess.Role[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"mintWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storeCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"vaultEventDone\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"vaultEventWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultEvents\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"eventData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultStoreCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultWatermark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMapping\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"eventData\",\"type\":\"bytes\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMapping\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"doMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"updateVaultWatermark\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMapping\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"done\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n_\",\"type\":\"uint256\"}],\"name\":\"setStoreCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storeCounter_\",\"type\":\"uint256\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMapping\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"rescueVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// XEvents is an auto generated Go binding around an Ethereum contract.
type XEvents struct {
	XEventsCaller     // Read-only binding to the contract
	XEventsTransactor // Write-only binding to the contract
	XEventsFilterer   // Log filterer for contract events
}

// XEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type XEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XEventsSession struct {
	Contract     *XEvents          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XEventsCallerSession struct {
	Contract *XEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// XEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XEventsTransactorSession struct {
	Contract     *XEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// XEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type XEventsRaw struct {
	Contract *XEvents // Generic contract binding to access the raw methods on
}

// XEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XEventsCallerRaw struct {
	Contract *XEventsCaller // Generic read-only contract binding to access the raw methods on
}

// XEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XEventsTransactorRaw struct {
	Contract *XEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXEvents creates a new instance of XEvents, bound to a specific deployed contract.
func NewXEvents(address common.Address, backend bind.ContractBackend) (*XEvents, error) {
	contract, err := bindXEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XEvents{XEventsCaller: XEventsCaller{contract: contract}, XEventsTransactor: XEventsTransactor{contract: contract}, XEventsFilterer: XEventsFilterer{contract: contract}}, nil
}

// NewXEventsCaller creates a new read-only instance of XEvents, bound to a specific deployed contract.
func NewXEventsCaller(address common.Address, caller bind.ContractCaller) (*XEventsCaller, error) {
	contract, err := bindXEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XEventsCaller{contract: contract}, nil
}

// NewXEventsTransactor creates a new write-only instance of XEvents, bound to a specific deployed contract.
func NewXEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*XEventsTransactor, error) {
	contract, err := bindXEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XEventsTransactor{contract: contract}, nil
}

// NewXEventsFilterer creates a new log filterer instance of XEvents, bound to a specific deployed contract.
func NewXEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*XEventsFilterer, error) {
	contract, err := bindXEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XEventsFilterer{contract: contract}, nil
}

// bindXEvents binds a generic wrapper to an already deployed contract.
func bindXEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XEvents *XEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XEvents.Contract.XEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XEvents *XEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XEvents.Contract.XEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XEvents *XEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XEvents.Contract.XEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XEvents *XEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XEvents *XEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XEvents *XEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XEvents.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XEvents *XEventsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XEvents *XEventsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _XEvents.Contract.DEFAULTADMINROLE(&_XEvents.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XEvents *XEventsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _XEvents.Contract.DEFAULTADMINROLE(&_XEvents.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XEvents *XEventsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XEvents *XEventsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _XEvents.Contract.GetRoleAdmin(&_XEvents.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XEvents *XEventsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _XEvents.Contract.GetRoleAdmin(&_XEvents.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XEvents *XEventsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XEvents *XEventsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _XEvents.Contract.GetRoleMember(&_XEvents.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XEvents *XEventsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _XEvents.Contract.GetRoleMember(&_XEvents.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XEvents *XEventsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XEvents *XEventsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _XEvents.Contract.GetRoleMemberCount(&_XEvents.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XEvents *XEventsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _XEvents.Contract.GetRoleMemberCount(&_XEvents.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XEvents *XEventsCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XEvents *XEventsSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _XEvents.Contract.GetRoleMembers(&_XEvents.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XEvents *XEventsCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _XEvents.Contract.GetRoleMembers(&_XEvents.CallOpts, role)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XEvents *XEventsCaller) GetRoles(opts *bind.CallOpts) ([]RoleAccessRole, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "getRoles")

	if err != nil {
		return *new([]RoleAccessRole), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleAccessRole)).(*[]RoleAccessRole)

	return out0, err

}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XEvents *XEventsSession) GetRoles() ([]RoleAccessRole, error) {
	return _XEvents.Contract.GetRoles(&_XEvents.CallOpts)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XEvents *XEventsCallerSession) GetRoles() ([]RoleAccessRole, error) {
	return _XEvents.Contract.GetRoles(&_XEvents.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XEvents *XEventsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XEvents *XEventsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _XEvents.Contract.HasRole(&_XEvents.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XEvents *XEventsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _XEvents.Contract.HasRole(&_XEvents.CallOpts, role, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_XEvents *XEventsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_XEvents *XEventsSession) Initialized() (bool, error) {
	return _XEvents.Contract.Initialized(&_XEvents.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_XEvents *XEventsCallerSession) Initialized() (bool, error) {
	return _XEvents.Contract.Initialized(&_XEvents.CallOpts)
}

// MintWatermark is a free data retrieval call binding the contract method 0xbb4f797f.
//
// Solidity: function mintWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCaller) MintWatermark(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "mintWatermark", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintWatermark is a free data retrieval call binding the contract method 0xbb4f797f.
//
// Solidity: function mintWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsSession) MintWatermark(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.MintWatermark(&_XEvents.CallOpts, arg0, arg1)
}

// MintWatermark is a free data retrieval call binding the contract method 0xbb4f797f.
//
// Solidity: function mintWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCallerSession) MintWatermark(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.MintWatermark(&_XEvents.CallOpts, arg0, arg1)
}

// StoreCounter is a free data retrieval call binding the contract method 0xb9ca868d.
//
// Solidity: function storeCounter() view returns(uint256)
func (_XEvents *XEventsCaller) StoreCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "storeCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoreCounter is a free data retrieval call binding the contract method 0xb9ca868d.
//
// Solidity: function storeCounter() view returns(uint256)
func (_XEvents *XEventsSession) StoreCounter() (*big.Int, error) {
	return _XEvents.Contract.StoreCounter(&_XEvents.CallOpts)
}

// StoreCounter is a free data retrieval call binding the contract method 0xb9ca868d.
//
// Solidity: function storeCounter() view returns(uint256)
func (_XEvents *XEventsCallerSession) StoreCounter() (*big.Int, error) {
	return _XEvents.Contract.StoreCounter(&_XEvents.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XEvents *XEventsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XEvents *XEventsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _XEvents.Contract.SupportsInterface(&_XEvents.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XEvents *XEventsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _XEvents.Contract.SupportsInterface(&_XEvents.CallOpts, interfaceId)
}

// VaultEventDone is a free data retrieval call binding the contract method 0xad9ef616.
//
// Solidity: function vaultEventDone(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCaller) VaultEventDone(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "vaultEventDone", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultEventDone is a free data retrieval call binding the contract method 0xad9ef616.
//
// Solidity: function vaultEventDone(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsSession) VaultEventDone(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.VaultEventDone(&_XEvents.CallOpts, arg0, arg1)
}

// VaultEventDone is a free data retrieval call binding the contract method 0xad9ef616.
//
// Solidity: function vaultEventDone(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCallerSession) VaultEventDone(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.VaultEventDone(&_XEvents.CallOpts, arg0, arg1)
}

// VaultEventWatermark is a free data retrieval call binding the contract method 0xa5067868.
//
// Solidity: function vaultEventWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCaller) VaultEventWatermark(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "vaultEventWatermark", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultEventWatermark is a free data retrieval call binding the contract method 0xa5067868.
//
// Solidity: function vaultEventWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsSession) VaultEventWatermark(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.VaultEventWatermark(&_XEvents.CallOpts, arg0, arg1)
}

// VaultEventWatermark is a free data retrieval call binding the contract method 0xa5067868.
//
// Solidity: function vaultEventWatermark(address , bytes32 ) view returns(uint256)
func (_XEvents *XEventsCallerSession) VaultEventWatermark(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _XEvents.Contract.VaultEventWatermark(&_XEvents.CallOpts, arg0, arg1)
}

// VaultEvents is a free data retrieval call binding the contract method 0x68741d5e.
//
// Solidity: function vaultEvents(address , bytes32 , uint256 ) view returns(bytes eventData, bytes sig, uint256 blockNumber)
func (_XEvents *XEventsCaller) VaultEvents(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "vaultEvents", arg0, arg1, arg2)

	outstruct := new(struct {
		EventData   []byte
		Sig         []byte
		BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EventData = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Sig = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VaultEvents is a free data retrieval call binding the contract method 0x68741d5e.
//
// Solidity: function vaultEvents(address , bytes32 , uint256 ) view returns(bytes eventData, bytes sig, uint256 blockNumber)
func (_XEvents *XEventsSession) VaultEvents(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}, error) {
	return _XEvents.Contract.VaultEvents(&_XEvents.CallOpts, arg0, arg1, arg2)
}

// VaultEvents is a free data retrieval call binding the contract method 0x68741d5e.
//
// Solidity: function vaultEvents(address , bytes32 , uint256 ) view returns(bytes eventData, bytes sig, uint256 blockNumber)
func (_XEvents *XEventsCallerSession) VaultEvents(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}, error) {
	return _XEvents.Contract.VaultEvents(&_XEvents.CallOpts, arg0, arg1, arg2)
}

// VaultStoreCounter is a free data retrieval call binding the contract method 0x1e2cb95f.
//
// Solidity: function vaultStoreCounter(address , uint256 ) view returns(uint256)
func (_XEvents *XEventsCaller) VaultStoreCounter(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "vaultStoreCounter", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultStoreCounter is a free data retrieval call binding the contract method 0x1e2cb95f.
//
// Solidity: function vaultStoreCounter(address , uint256 ) view returns(uint256)
func (_XEvents *XEventsSession) VaultStoreCounter(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _XEvents.Contract.VaultStoreCounter(&_XEvents.CallOpts, arg0, arg1)
}

// VaultStoreCounter is a free data retrieval call binding the contract method 0x1e2cb95f.
//
// Solidity: function vaultStoreCounter(address , uint256 ) view returns(uint256)
func (_XEvents *XEventsCallerSession) VaultStoreCounter(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _XEvents.Contract.VaultStoreCounter(&_XEvents.CallOpts, arg0, arg1)
}

// VaultWatermark is a free data retrieval call binding the contract method 0x5b65b990.
//
// Solidity: function vaultWatermark(address ) view returns(uint256)
func (_XEvents *XEventsCaller) VaultWatermark(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _XEvents.contract.Call(opts, &out, "vaultWatermark", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultWatermark is a free data retrieval call binding the contract method 0x5b65b990.
//
// Solidity: function vaultWatermark(address ) view returns(uint256)
func (_XEvents *XEventsSession) VaultWatermark(arg0 common.Address) (*big.Int, error) {
	return _XEvents.Contract.VaultWatermark(&_XEvents.CallOpts, arg0)
}

// VaultWatermark is a free data retrieval call binding the contract method 0x5b65b990.
//
// Solidity: function vaultWatermark(address ) view returns(uint256)
func (_XEvents *XEventsCallerSession) VaultWatermark(arg0 common.Address) (*big.Int, error) {
	return _XEvents.Contract.VaultWatermark(&_XEvents.CallOpts, arg0)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsTransactor) AddRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "addRoleMember", role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.AddRoleMember(&_XEvents.TransactOpts, role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsTransactorSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.AddRoleMember(&_XEvents.TransactOpts, role, member)
}

// DoMint is a paid mutator transaction binding the contract method 0xb5562cbd.
//
// Solidity: function doMint(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactor) DoMint(opts *bind.TransactOpts, vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "doMint", vault, tokenMapping, nonce)
}

// DoMint is a paid mutator transaction binding the contract method 0xb5562cbd.
//
// Solidity: function doMint(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsSession) DoMint(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.DoMint(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// DoMint is a paid mutator transaction binding the contract method 0xb5562cbd.
//
// Solidity: function doMint(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactorSession) DoMint(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.DoMint(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// Done is a paid mutator transaction binding the contract method 0x0cdec404.
//
// Solidity: function done(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactor) Done(opts *bind.TransactOpts, vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "done", vault, tokenMapping, nonce)
}

// Done is a paid mutator transaction binding the contract method 0x0cdec404.
//
// Solidity: function done(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsSession) Done(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.Done(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// Done is a paid mutator transaction binding the contract method 0x0cdec404.
//
// Solidity: function done(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactorSession) Done(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.Done(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XEvents *XEventsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.GrantRole(&_XEvents.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.GrantRole(&_XEvents.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XEvents *XEventsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XEvents *XEventsSession) Initialize() (*types.Transaction, error) {
	return _XEvents.Contract.Initialize(&_XEvents.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XEvents *XEventsTransactorSession) Initialize() (*types.Transaction, error) {
	return _XEvents.Contract.Initialize(&_XEvents.TransactOpts)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsTransactor) RemoveRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "removeRoleMember", role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RemoveRoleMember(&_XEvents.TransactOpts, role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XEvents *XEventsTransactorSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RemoveRoleMember(&_XEvents.TransactOpts, role, member)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XEvents *XEventsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RenounceRole(&_XEvents.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RenounceRole(&_XEvents.TransactOpts, role, account)
}

// Rescue is a paid mutator transaction binding the contract method 0x18c023ae.
//
// Solidity: function rescue(address vault, uint256 blockNumber, uint256 storeCounter_) returns()
func (_XEvents *XEventsTransactor) Rescue(opts *bind.TransactOpts, vault common.Address, blockNumber *big.Int, storeCounter_ *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "rescue", vault, blockNumber, storeCounter_)
}

// Rescue is a paid mutator transaction binding the contract method 0x18c023ae.
//
// Solidity: function rescue(address vault, uint256 blockNumber, uint256 storeCounter_) returns()
func (_XEvents *XEventsSession) Rescue(vault common.Address, blockNumber *big.Int, storeCounter_ *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.Rescue(&_XEvents.TransactOpts, vault, blockNumber, storeCounter_)
}

// Rescue is a paid mutator transaction binding the contract method 0x18c023ae.
//
// Solidity: function rescue(address vault, uint256 blockNumber, uint256 storeCounter_) returns()
func (_XEvents *XEventsTransactorSession) Rescue(vault common.Address, blockNumber *big.Int, storeCounter_ *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.Rescue(&_XEvents.TransactOpts, vault, blockNumber, storeCounter_)
}

// RescueVault is a paid mutator transaction binding the contract method 0x63e19513.
//
// Solidity: function rescueVault(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactor) RescueVault(opts *bind.TransactOpts, vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "rescueVault", vault, tokenMapping, nonce)
}

// RescueVault is a paid mutator transaction binding the contract method 0x63e19513.
//
// Solidity: function rescueVault(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsSession) RescueVault(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.RescueVault(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// RescueVault is a paid mutator transaction binding the contract method 0x63e19513.
//
// Solidity: function rescueVault(address vault, bytes32 tokenMapping, uint256 nonce) returns()
func (_XEvents *XEventsTransactorSession) RescueVault(vault common.Address, tokenMapping [32]byte, nonce *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.RescueVault(&_XEvents.TransactOpts, vault, tokenMapping, nonce)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XEvents *XEventsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RevokeRole(&_XEvents.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XEvents *XEventsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XEvents.Contract.RevokeRole(&_XEvents.TransactOpts, role, account)
}

// SetStoreCounter is a paid mutator transaction binding the contract method 0x284a26a0.
//
// Solidity: function setStoreCounter(uint256 n_) returns()
func (_XEvents *XEventsTransactor) SetStoreCounter(opts *bind.TransactOpts, n_ *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "setStoreCounter", n_)
}

// SetStoreCounter is a paid mutator transaction binding the contract method 0x284a26a0.
//
// Solidity: function setStoreCounter(uint256 n_) returns()
func (_XEvents *XEventsSession) SetStoreCounter(n_ *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.SetStoreCounter(&_XEvents.TransactOpts, n_)
}

// SetStoreCounter is a paid mutator transaction binding the contract method 0x284a26a0.
//
// Solidity: function setStoreCounter(uint256 n_) returns()
func (_XEvents *XEventsTransactorSession) SetStoreCounter(n_ *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.SetStoreCounter(&_XEvents.TransactOpts, n_)
}

// Store is a paid mutator transaction binding the contract method 0xfebb7ca1.
//
// Solidity: function store(bytes sig, address vault, uint256 nonce, bytes32 tokenMapping, uint256 blockNumber, bytes eventData) returns()
func (_XEvents *XEventsTransactor) Store(opts *bind.TransactOpts, sig []byte, vault common.Address, nonce *big.Int, tokenMapping [32]byte, blockNumber *big.Int, eventData []byte) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "store", sig, vault, nonce, tokenMapping, blockNumber, eventData)
}

// Store is a paid mutator transaction binding the contract method 0xfebb7ca1.
//
// Solidity: function store(bytes sig, address vault, uint256 nonce, bytes32 tokenMapping, uint256 blockNumber, bytes eventData) returns()
func (_XEvents *XEventsSession) Store(sig []byte, vault common.Address, nonce *big.Int, tokenMapping [32]byte, blockNumber *big.Int, eventData []byte) (*types.Transaction, error) {
	return _XEvents.Contract.Store(&_XEvents.TransactOpts, sig, vault, nonce, tokenMapping, blockNumber, eventData)
}

// Store is a paid mutator transaction binding the contract method 0xfebb7ca1.
//
// Solidity: function store(bytes sig, address vault, uint256 nonce, bytes32 tokenMapping, uint256 blockNumber, bytes eventData) returns()
func (_XEvents *XEventsTransactorSession) Store(sig []byte, vault common.Address, nonce *big.Int, tokenMapping [32]byte, blockNumber *big.Int, eventData []byte) (*types.Transaction, error) {
	return _XEvents.Contract.Store(&_XEvents.TransactOpts, sig, vault, nonce, tokenMapping, blockNumber, eventData)
}

// UpdateVaultWatermark is a paid mutator transaction binding the contract method 0x63bea0d9.
//
// Solidity: function updateVaultWatermark(address vault, uint256 blockNumber) returns()
func (_XEvents *XEventsTransactor) UpdateVaultWatermark(opts *bind.TransactOpts, vault common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _XEvents.contract.Transact(opts, "updateVaultWatermark", vault, blockNumber)
}

// UpdateVaultWatermark is a paid mutator transaction binding the contract method 0x63bea0d9.
//
// Solidity: function updateVaultWatermark(address vault, uint256 blockNumber) returns()
func (_XEvents *XEventsSession) UpdateVaultWatermark(vault common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.UpdateVaultWatermark(&_XEvents.TransactOpts, vault, blockNumber)
}

// UpdateVaultWatermark is a paid mutator transaction binding the contract method 0x63bea0d9.
//
// Solidity: function updateVaultWatermark(address vault, uint256 blockNumber) returns()
func (_XEvents *XEventsTransactorSession) UpdateVaultWatermark(vault common.Address, blockNumber *big.Int) (*types.Transaction, error) {
	return _XEvents.Contract.UpdateVaultWatermark(&_XEvents.TransactOpts, vault, blockNumber)
}

// XEventsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the XEvents contract.
type XEventsRoleAdminChangedIterator struct {
	Event *XEventsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *XEventsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XEventsRoleAdminChanged)
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
		it.Event = new(XEventsRoleAdminChanged)
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
func (it *XEventsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XEventsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XEventsRoleAdminChanged represents a RoleAdminChanged event raised by the XEvents contract.
type XEventsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_XEvents *XEventsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*XEventsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _XEvents.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &XEventsRoleAdminChangedIterator{contract: _XEvents.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_XEvents *XEventsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *XEventsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _XEvents.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XEventsRoleAdminChanged)
				if err := _XEvents.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_XEvents *XEventsFilterer) ParseRoleAdminChanged(log types.Log) (*XEventsRoleAdminChanged, error) {
	event := new(XEventsRoleAdminChanged)
	if err := _XEvents.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XEventsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the XEvents contract.
type XEventsRoleGrantedIterator struct {
	Event *XEventsRoleGranted // Event containing the contract specifics and raw log

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
func (it *XEventsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XEventsRoleGranted)
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
		it.Event = new(XEventsRoleGranted)
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
func (it *XEventsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XEventsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XEventsRoleGranted represents a RoleGranted event raised by the XEvents contract.
type XEventsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_XEvents *XEventsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XEventsRoleGrantedIterator, error) {

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

	logs, sub, err := _XEvents.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XEventsRoleGrantedIterator{contract: _XEvents.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_XEvents *XEventsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *XEventsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _XEvents.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XEventsRoleGranted)
				if err := _XEvents.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_XEvents *XEventsFilterer) ParseRoleGranted(log types.Log) (*XEventsRoleGranted, error) {
	event := new(XEventsRoleGranted)
	if err := _XEvents.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XEventsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the XEvents contract.
type XEventsRoleRevokedIterator struct {
	Event *XEventsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *XEventsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XEventsRoleRevoked)
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
		it.Event = new(XEventsRoleRevoked)
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
func (it *XEventsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XEventsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XEventsRoleRevoked represents a RoleRevoked event raised by the XEvents contract.
type XEventsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_XEvents *XEventsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XEventsRoleRevokedIterator, error) {

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

	logs, sub, err := _XEvents.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XEventsRoleRevokedIterator{contract: _XEvents.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_XEvents *XEventsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *XEventsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _XEvents.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XEventsRoleRevoked)
				if err := _XEvents.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_XEvents *XEventsFilterer) ParseRoleRevoked(log types.Log) (*XEventsRoleRevoked, error) {
	event := new(XEventsRoleRevoked)
	if err := _XEvents.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
