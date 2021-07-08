// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xconfig

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

// XConfigABI is the input ABI used to generate the binding from.
const XConfigABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VaultConfigVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"VaultConfigs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VssConfigVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"VssConfigs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeVssMemberList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"addRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"describe\",\"type\":\"string\"}],\"internalType\":\"structRoleAccess.Role[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"removeRoleMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateVssConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"name\":\"UpdateVaultConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// XConfig is an auto generated Go binding around an Ethereum contract.
type XConfig struct {
	XConfigCaller     // Read-only binding to the contract
	XConfigTransactor // Write-only binding to the contract
	XConfigFilterer   // Log filterer for contract events
}

// XConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type XConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XConfigSession struct {
	Contract     *XConfig          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XConfigCallerSession struct {
	Contract *XConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// XConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XConfigTransactorSession struct {
	Contract     *XConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// XConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type XConfigRaw struct {
	Contract *XConfig // Generic contract binding to access the raw methods on
}

// XConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XConfigCallerRaw struct {
	Contract *XConfigCaller // Generic read-only contract binding to access the raw methods on
}

// XConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XConfigTransactorRaw struct {
	Contract *XConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXConfig creates a new instance of XConfig, bound to a specific deployed contract.
func NewXConfig(address common.Address, backend bind.ContractBackend) (*XConfig, error) {
	contract, err := bindXConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XConfig{XConfigCaller: XConfigCaller{contract: contract}, XConfigTransactor: XConfigTransactor{contract: contract}, XConfigFilterer: XConfigFilterer{contract: contract}}, nil
}

// NewXConfigCaller creates a new read-only instance of XConfig, bound to a specific deployed contract.
func NewXConfigCaller(address common.Address, caller bind.ContractCaller) (*XConfigCaller, error) {
	contract, err := bindXConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XConfigCaller{contract: contract}, nil
}

// NewXConfigTransactor creates a new write-only instance of XConfig, bound to a specific deployed contract.
func NewXConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*XConfigTransactor, error) {
	contract, err := bindXConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XConfigTransactor{contract: contract}, nil
}

// NewXConfigFilterer creates a new log filterer instance of XConfig, bound to a specific deployed contract.
func NewXConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*XConfigFilterer, error) {
	contract, err := bindXConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XConfigFilterer{contract: contract}, nil
}

// bindXConfig binds a generic wrapper to an already deployed contract.
func bindXConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XConfig *XConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XConfig.Contract.XConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XConfig *XConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XConfig.Contract.XConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XConfig *XConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XConfig.Contract.XConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XConfig *XConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XConfig *XConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XConfig *XConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XConfig.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XConfig *XConfigCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XConfig *XConfigSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _XConfig.Contract.DEFAULTADMINROLE(&_XConfig.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_XConfig *XConfigCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _XConfig.Contract.DEFAULTADMINROLE(&_XConfig.CallOpts)
}

// VaultConfigVersion is a free data retrieval call binding the contract method 0x11321ab8.
//
// Solidity: function VaultConfigVersion() view returns(uint256)
func (_XConfig *XConfigCaller) VaultConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "VaultConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultConfigVersion is a free data retrieval call binding the contract method 0x11321ab8.
//
// Solidity: function VaultConfigVersion() view returns(uint256)
func (_XConfig *XConfigSession) VaultConfigVersion() (*big.Int, error) {
	return _XConfig.Contract.VaultConfigVersion(&_XConfig.CallOpts)
}

// VaultConfigVersion is a free data retrieval call binding the contract method 0x11321ab8.
//
// Solidity: function VaultConfigVersion() view returns(uint256)
func (_XConfig *XConfigCallerSession) VaultConfigVersion() (*big.Int, error) {
	return _XConfig.Contract.VaultConfigVersion(&_XConfig.CallOpts)
}

// VaultConfigs is a free data retrieval call binding the contract method 0xc70bdcf7.
//
// Solidity: function VaultConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigCaller) VaultConfigs(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "VaultConfigs", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// VaultConfigs is a free data retrieval call binding the contract method 0xc70bdcf7.
//
// Solidity: function VaultConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigSession) VaultConfigs(arg0 *big.Int) ([]byte, error) {
	return _XConfig.Contract.VaultConfigs(&_XConfig.CallOpts, arg0)
}

// VaultConfigs is a free data retrieval call binding the contract method 0xc70bdcf7.
//
// Solidity: function VaultConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigCallerSession) VaultConfigs(arg0 *big.Int) ([]byte, error) {
	return _XConfig.Contract.VaultConfigs(&_XConfig.CallOpts, arg0)
}

// VssConfigVersion is a free data retrieval call binding the contract method 0x68b75c6b.
//
// Solidity: function VssConfigVersion() view returns(uint256)
func (_XConfig *XConfigCaller) VssConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "VssConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VssConfigVersion is a free data retrieval call binding the contract method 0x68b75c6b.
//
// Solidity: function VssConfigVersion() view returns(uint256)
func (_XConfig *XConfigSession) VssConfigVersion() (*big.Int, error) {
	return _XConfig.Contract.VssConfigVersion(&_XConfig.CallOpts)
}

// VssConfigVersion is a free data retrieval call binding the contract method 0x68b75c6b.
//
// Solidity: function VssConfigVersion() view returns(uint256)
func (_XConfig *XConfigCallerSession) VssConfigVersion() (*big.Int, error) {
	return _XConfig.Contract.VssConfigVersion(&_XConfig.CallOpts)
}

// VssConfigs is a free data retrieval call binding the contract method 0xdb187255.
//
// Solidity: function VssConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigCaller) VssConfigs(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "VssConfigs", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// VssConfigs is a free data retrieval call binding the contract method 0xdb187255.
//
// Solidity: function VssConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigSession) VssConfigs(arg0 *big.Int) ([]byte, error) {
	return _XConfig.Contract.VssConfigs(&_XConfig.CallOpts, arg0)
}

// VssConfigs is a free data retrieval call binding the contract method 0xdb187255.
//
// Solidity: function VssConfigs(uint256 ) view returns(bytes)
func (_XConfig *XConfigCallerSession) VssConfigs(arg0 *big.Int) ([]byte, error) {
	return _XConfig.Contract.VssConfigs(&_XConfig.CallOpts, arg0)
}

// ActiveVssMemberList is a free data retrieval call binding the contract method 0x06faf263.
//
// Solidity: function activeVssMemberList(uint256 ) view returns(address)
func (_XConfig *XConfigCaller) ActiveVssMemberList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "activeVssMemberList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveVssMemberList is a free data retrieval call binding the contract method 0x06faf263.
//
// Solidity: function activeVssMemberList(uint256 ) view returns(address)
func (_XConfig *XConfigSession) ActiveVssMemberList(arg0 *big.Int) (common.Address, error) {
	return _XConfig.Contract.ActiveVssMemberList(&_XConfig.CallOpts, arg0)
}

// ActiveVssMemberList is a free data retrieval call binding the contract method 0x06faf263.
//
// Solidity: function activeVssMemberList(uint256 ) view returns(address)
func (_XConfig *XConfigCallerSession) ActiveVssMemberList(arg0 *big.Int) (common.Address, error) {
	return _XConfig.Contract.ActiveVssMemberList(&_XConfig.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XConfig *XConfigCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XConfig *XConfigSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _XConfig.Contract.GetRoleAdmin(&_XConfig.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_XConfig *XConfigCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _XConfig.Contract.GetRoleAdmin(&_XConfig.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XConfig *XConfigCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XConfig *XConfigSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _XConfig.Contract.GetRoleMember(&_XConfig.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_XConfig *XConfigCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _XConfig.Contract.GetRoleMember(&_XConfig.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XConfig *XConfigCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XConfig *XConfigSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _XConfig.Contract.GetRoleMemberCount(&_XConfig.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_XConfig *XConfigCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _XConfig.Contract.GetRoleMemberCount(&_XConfig.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XConfig *XConfigCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XConfig *XConfigSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _XConfig.Contract.GetRoleMembers(&_XConfig.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_XConfig *XConfigCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _XConfig.Contract.GetRoleMembers(&_XConfig.CallOpts, role)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XConfig *XConfigCaller) GetRoles(opts *bind.CallOpts) ([]RoleAccessRole, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "getRoles")

	if err != nil {
		return *new([]RoleAccessRole), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoleAccessRole)).(*[]RoleAccessRole)

	return out0, err

}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XConfig *XConfigSession) GetRoles() ([]RoleAccessRole, error) {
	return _XConfig.Contract.GetRoles(&_XConfig.CallOpts)
}

// GetRoles is a free data retrieval call binding the contract method 0x71061398.
//
// Solidity: function getRoles() pure returns((bytes32,string)[])
func (_XConfig *XConfigCallerSession) GetRoles() ([]RoleAccessRole, error) {
	return _XConfig.Contract.GetRoles(&_XConfig.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XConfig *XConfigCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XConfig *XConfigSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _XConfig.Contract.HasRole(&_XConfig.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_XConfig *XConfigCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _XConfig.Contract.HasRole(&_XConfig.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XConfig *XConfigCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _XConfig.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XConfig *XConfigSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _XConfig.Contract.SupportsInterface(&_XConfig.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_XConfig *XConfigCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _XConfig.Contract.SupportsInterface(&_XConfig.CallOpts, interfaceId)
}

// UpdateVaultConfig is a paid mutator transaction binding the contract method 0x971eecc4.
//
// Solidity: function UpdateVaultConfig(bytes config) returns()
func (_XConfig *XConfigTransactor) UpdateVaultConfig(opts *bind.TransactOpts, config []byte) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "UpdateVaultConfig", config)
}

// UpdateVaultConfig is a paid mutator transaction binding the contract method 0x971eecc4.
//
// Solidity: function UpdateVaultConfig(bytes config) returns()
func (_XConfig *XConfigSession) UpdateVaultConfig(config []byte) (*types.Transaction, error) {
	return _XConfig.Contract.UpdateVaultConfig(&_XConfig.TransactOpts, config)
}

// UpdateVaultConfig is a paid mutator transaction binding the contract method 0x971eecc4.
//
// Solidity: function UpdateVaultConfig(bytes config) returns()
func (_XConfig *XConfigTransactorSession) UpdateVaultConfig(config []byte) (*types.Transaction, error) {
	return _XConfig.Contract.UpdateVaultConfig(&_XConfig.TransactOpts, config)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigTransactor) AddRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "addRoleMember", role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.AddRoleMember(&_XConfig.TransactOpts, role, member)
}

// AddRoleMember is a paid mutator transaction binding the contract method 0x1b65471f.
//
// Solidity: function addRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigTransactorSession) AddRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.AddRoleMember(&_XConfig.TransactOpts, role, member)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XConfig *XConfigSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.GrantRole(&_XConfig.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.GrantRole(&_XConfig.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XConfig *XConfigTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XConfig *XConfigSession) Initialize() (*types.Transaction, error) {
	return _XConfig.Contract.Initialize(&_XConfig.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_XConfig *XConfigTransactorSession) Initialize() (*types.Transaction, error) {
	return _XConfig.Contract.Initialize(&_XConfig.TransactOpts)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigTransactor) RemoveRoleMember(opts *bind.TransactOpts, role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "removeRoleMember", role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RemoveRoleMember(&_XConfig.TransactOpts, role, member)
}

// RemoveRoleMember is a paid mutator transaction binding the contract method 0x4dd8fac8.
//
// Solidity: function removeRoleMember(bytes32 role, address member) returns(bool)
func (_XConfig *XConfigTransactorSession) RemoveRoleMember(role [32]byte, member common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RemoveRoleMember(&_XConfig.TransactOpts, role, member)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XConfig *XConfigSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RenounceRole(&_XConfig.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RenounceRole(&_XConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XConfig *XConfigSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RevokeRole(&_XConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_XConfig *XConfigTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _XConfig.Contract.RevokeRole(&_XConfig.TransactOpts, role, account)
}

// UpdateVssConfig is a paid mutator transaction binding the contract method 0xc5093e3d.
//
// Solidity: function updateVssConfig() returns()
func (_XConfig *XConfigTransactor) UpdateVssConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XConfig.contract.Transact(opts, "updateVssConfig")
}

// UpdateVssConfig is a paid mutator transaction binding the contract method 0xc5093e3d.
//
// Solidity: function updateVssConfig() returns()
func (_XConfig *XConfigSession) UpdateVssConfig() (*types.Transaction, error) {
	return _XConfig.Contract.UpdateVssConfig(&_XConfig.TransactOpts)
}

// UpdateVssConfig is a paid mutator transaction binding the contract method 0xc5093e3d.
//
// Solidity: function updateVssConfig() returns()
func (_XConfig *XConfigTransactorSession) UpdateVssConfig() (*types.Transaction, error) {
	return _XConfig.Contract.UpdateVssConfig(&_XConfig.TransactOpts)
}

// XConfigRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the XConfig contract.
type XConfigRoleAdminChangedIterator struct {
	Event *XConfigRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *XConfigRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XConfigRoleAdminChanged)
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
		it.Event = new(XConfigRoleAdminChanged)
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
func (it *XConfigRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XConfigRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XConfigRoleAdminChanged represents a RoleAdminChanged event raised by the XConfig contract.
type XConfigRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_XConfig *XConfigFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*XConfigRoleAdminChangedIterator, error) {

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

	logs, sub, err := _XConfig.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &XConfigRoleAdminChangedIterator{contract: _XConfig.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_XConfig *XConfigFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *XConfigRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _XConfig.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XConfigRoleAdminChanged)
				if err := _XConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_XConfig *XConfigFilterer) ParseRoleAdminChanged(log types.Log) (*XConfigRoleAdminChanged, error) {
	event := new(XConfigRoleAdminChanged)
	if err := _XConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XConfigRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the XConfig contract.
type XConfigRoleGrantedIterator struct {
	Event *XConfigRoleGranted // Event containing the contract specifics and raw log

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
func (it *XConfigRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XConfigRoleGranted)
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
		it.Event = new(XConfigRoleGranted)
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
func (it *XConfigRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XConfigRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XConfigRoleGranted represents a RoleGranted event raised by the XConfig contract.
type XConfigRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_XConfig *XConfigFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XConfigRoleGrantedIterator, error) {

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

	logs, sub, err := _XConfig.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XConfigRoleGrantedIterator{contract: _XConfig.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_XConfig *XConfigFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *XConfigRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _XConfig.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XConfigRoleGranted)
				if err := _XConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_XConfig *XConfigFilterer) ParseRoleGranted(log types.Log) (*XConfigRoleGranted, error) {
	event := new(XConfigRoleGranted)
	if err := _XConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XConfigRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the XConfig contract.
type XConfigRoleRevokedIterator struct {
	Event *XConfigRoleRevoked // Event containing the contract specifics and raw log

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
func (it *XConfigRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XConfigRoleRevoked)
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
		it.Event = new(XConfigRoleRevoked)
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
func (it *XConfigRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XConfigRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XConfigRoleRevoked represents a RoleRevoked event raised by the XConfig contract.
type XConfigRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_XConfig *XConfigFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*XConfigRoleRevokedIterator, error) {

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

	logs, sub, err := _XConfig.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &XConfigRoleRevokedIterator{contract: _XConfig.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_XConfig *XConfigFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *XConfigRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _XConfig.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XConfigRoleRevoked)
				if err := _XConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_XConfig *XConfigFilterer) ParseRoleRevoked(log types.Log) (*XConfigRoleRevoked, error) {
	event := new(XConfigRoleRevoked)
	if err := _XConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
