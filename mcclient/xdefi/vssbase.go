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

// VssBaseRevealedShare is an auto generated low-level Go binding around an user-defined struct.
type VssBaseRevealedShare struct {
	PubShare      []byte
	PubSig        []byte
	PriShare      []byte
	PriSig        []byte
	Revealed      []byte
	Violator      common.Address
	Whistleblower common.Address
}

// VssBaseABI is the input ABI used to generate the binding from.
const VssBaseABI = "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"threshold\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"caller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"generalAttributes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastConfigUpload\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastConfigUploadByBlock\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastNodeChangeBlock\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastNodeChangeConfigVersion\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastSlashingVoted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"revealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"revealedReporter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"revealedSigMapping\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"revealedViolator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"reveals\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"PubShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PubSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PriShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PriSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Revealed\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Whistleblower\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"slashed\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"slashingRejects\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashingVoters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"slashingVotes\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slowNodeThreshold\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slowNodeVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"slowNodeVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vssConfigVersion\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vssNodeCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vssNodeMemberships\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vssThreshold\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"namespace\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"key\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setGeneralAttributes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"newThreshold\",\"type\":\"int256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"callerAddr\",\"type\":\"address\"}],\"name\":\"setCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"newThreshold\",\"type\":\"int256\"}],\"name\":\"setSlowNodeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publickey\",\"type\":\"bytes32\"}],\"name\":\"registerVSS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"unregisterVSS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"deactivateVSS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"activateVSS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicShares\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"privateShares\",\"type\":\"bytes\"}],\"name\":\"uploadVSSConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"configVersion\",\"type\":\"int256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"}],\"name\":\"reportSlowNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"priShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"priSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"revealedPri\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"getRevealedShare\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"PubShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PubSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PriShare\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"PriSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Revealed\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"Violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Whistleblower\",\"type\":\"address\"}],\"internalType\":\"structVssBase.RevealedShare\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"slash\",\"type\":\"bool\"}],\"name\":\"slashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLastSlashVoted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"configVersion\",\"type\":\"int256\"}],\"name\":\"isConfigReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"configVersion\",\"type\":\"int256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetLastSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getPublicShares\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"getPrivateShares\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"}],\"name\":\"getVSSNodesPubkey\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveVSSMemberCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveVSSMemberList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"}],\"name\":\"getVSSNodesIndexs\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scs\",\"type\":\"address\"}],\"name\":\"getVSSNodeIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetVSSGroup\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VssBase is an auto generated Go binding around an Ethereum contract.
type VssBase struct {
	VssBaseCaller     // Read-only binding to the contract
	VssBaseTransactor // Write-only binding to the contract
	VssBaseFilterer   // Log filterer for contract events
}

// VssBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type VssBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VssBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VssBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VssBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VssBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VssBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VssBaseSession struct {
	Contract     *VssBase          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VssBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VssBaseCallerSession struct {
	Contract *VssBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VssBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VssBaseTransactorSession struct {
	Contract     *VssBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VssBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type VssBaseRaw struct {
	Contract *VssBase // Generic contract binding to access the raw methods on
}

// VssBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VssBaseCallerRaw struct {
	Contract *VssBaseCaller // Generic read-only contract binding to access the raw methods on
}

// VssBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VssBaseTransactorRaw struct {
	Contract *VssBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVssBase creates a new instance of VssBase, bound to a specific deployed contract.
func NewVssBase(address common.Address, backend bind.ContractBackend) (*VssBase, error) {
	contract, err := bindVssBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VssBase{VssBaseCaller: VssBaseCaller{contract: contract}, VssBaseTransactor: VssBaseTransactor{contract: contract}, VssBaseFilterer: VssBaseFilterer{contract: contract}}, nil
}

// NewVssBaseCaller creates a new read-only instance of VssBase, bound to a specific deployed contract.
func NewVssBaseCaller(address common.Address, caller bind.ContractCaller) (*VssBaseCaller, error) {
	contract, err := bindVssBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VssBaseCaller{contract: contract}, nil
}

// NewVssBaseTransactor creates a new write-only instance of VssBase, bound to a specific deployed contract.
func NewVssBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*VssBaseTransactor, error) {
	contract, err := bindVssBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VssBaseTransactor{contract: contract}, nil
}

// NewVssBaseFilterer creates a new log filterer instance of VssBase, bound to a specific deployed contract.
func NewVssBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*VssBaseFilterer, error) {
	contract, err := bindVssBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VssBaseFilterer{contract: contract}, nil
}

// bindVssBase binds a generic wrapper to an already deployed contract.
func bindVssBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VssBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VssBase *VssBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VssBase.Contract.VssBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VssBase *VssBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VssBase.Contract.VssBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VssBase *VssBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VssBase.Contract.VssBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VssBase *VssBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VssBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VssBase *VssBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VssBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VssBase *VssBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VssBase.Contract.contract.Transact(opts, method, params...)
}

// GetCaller is a free data retrieval call binding the contract method 0xd89da49c.
//
// Solidity: function GetCaller() view returns(address)
func (_VssBase *VssBaseCaller) GetCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "GetCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCaller is a free data retrieval call binding the contract method 0xd89da49c.
//
// Solidity: function GetCaller() view returns(address)
func (_VssBase *VssBaseSession) GetCaller() (common.Address, error) {
	return _VssBase.Contract.GetCaller(&_VssBase.CallOpts)
}

// GetCaller is a free data retrieval call binding the contract method 0xd89da49c.
//
// Solidity: function GetCaller() view returns(address)
func (_VssBase *VssBaseCallerSession) GetCaller() (common.Address, error) {
	return _VssBase.Contract.GetCaller(&_VssBase.CallOpts)
}

// GetLastSender is a free data retrieval call binding the contract method 0x9f1e5f0e.
//
// Solidity: function GetLastSender() view returns(address)
func (_VssBase *VssBaseCaller) GetLastSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "GetLastSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLastSender is a free data retrieval call binding the contract method 0x9f1e5f0e.
//
// Solidity: function GetLastSender() view returns(address)
func (_VssBase *VssBaseSession) GetLastSender() (common.Address, error) {
	return _VssBase.Contract.GetLastSender(&_VssBase.CallOpts)
}

// GetLastSender is a free data retrieval call binding the contract method 0x9f1e5f0e.
//
// Solidity: function GetLastSender() view returns(address)
func (_VssBase *VssBaseCallerSession) GetLastSender() (common.Address, error) {
	return _VssBase.Contract.GetLastSender(&_VssBase.CallOpts)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_VssBase *VssBaseCaller) Caller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "caller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_VssBase *VssBaseSession) Caller() (common.Address, error) {
	return _VssBase.Contract.Caller(&_VssBase.CallOpts)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_VssBase *VssBaseCallerSession) Caller() (common.Address, error) {
	return _VssBase.Contract.Caller(&_VssBase.CallOpts)
}

// GeneralAttributes is a free data retrieval call binding the contract method 0x50743a45.
//
// Solidity: function generalAttributes(bytes32 , int256 ) view returns(bytes)
func (_VssBase *VssBaseCaller) GeneralAttributes(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "generalAttributes", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GeneralAttributes is a free data retrieval call binding the contract method 0x50743a45.
//
// Solidity: function generalAttributes(bytes32 , int256 ) view returns(bytes)
func (_VssBase *VssBaseSession) GeneralAttributes(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _VssBase.Contract.GeneralAttributes(&_VssBase.CallOpts, arg0, arg1)
}

// GeneralAttributes is a free data retrieval call binding the contract method 0x50743a45.
//
// Solidity: function generalAttributes(bytes32 , int256 ) view returns(bytes)
func (_VssBase *VssBaseCallerSession) GeneralAttributes(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _VssBase.Contract.GeneralAttributes(&_VssBase.CallOpts, arg0, arg1)
}

// GetActiveVSSMemberCount is a free data retrieval call binding the contract method 0x029471e0.
//
// Solidity: function getActiveVSSMemberCount() view returns(int256)
func (_VssBase *VssBaseCaller) GetActiveVSSMemberCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getActiveVSSMemberCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVSSMemberCount is a free data retrieval call binding the contract method 0x029471e0.
//
// Solidity: function getActiveVSSMemberCount() view returns(int256)
func (_VssBase *VssBaseSession) GetActiveVSSMemberCount() (*big.Int, error) {
	return _VssBase.Contract.GetActiveVSSMemberCount(&_VssBase.CallOpts)
}

// GetActiveVSSMemberCount is a free data retrieval call binding the contract method 0x029471e0.
//
// Solidity: function getActiveVSSMemberCount() view returns(int256)
func (_VssBase *VssBaseCallerSession) GetActiveVSSMemberCount() (*big.Int, error) {
	return _VssBase.Contract.GetActiveVSSMemberCount(&_VssBase.CallOpts)
}

// GetActiveVSSMemberList is a free data retrieval call binding the contract method 0xed0c88ef.
//
// Solidity: function getActiveVSSMemberList() view returns(address[])
func (_VssBase *VssBaseCaller) GetActiveVSSMemberList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getActiveVSSMemberList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveVSSMemberList is a free data retrieval call binding the contract method 0xed0c88ef.
//
// Solidity: function getActiveVSSMemberList() view returns(address[])
func (_VssBase *VssBaseSession) GetActiveVSSMemberList() ([]common.Address, error) {
	return _VssBase.Contract.GetActiveVSSMemberList(&_VssBase.CallOpts)
}

// GetActiveVSSMemberList is a free data retrieval call binding the contract method 0xed0c88ef.
//
// Solidity: function getActiveVSSMemberList() view returns(address[])
func (_VssBase *VssBaseCallerSession) GetActiveVSSMemberList() ([]common.Address, error) {
	return _VssBase.Contract.GetActiveVSSMemberList(&_VssBase.CallOpts)
}

// GetLastSlashVoted is a free data retrieval call binding the contract method 0x7e2a1ed3.
//
// Solidity: function getLastSlashVoted(address addr) view returns(int256)
func (_VssBase *VssBaseCaller) GetLastSlashVoted(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getLastSlashVoted", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastSlashVoted is a free data retrieval call binding the contract method 0x7e2a1ed3.
//
// Solidity: function getLastSlashVoted(address addr) view returns(int256)
func (_VssBase *VssBaseSession) GetLastSlashVoted(addr common.Address) (*big.Int, error) {
	return _VssBase.Contract.GetLastSlashVoted(&_VssBase.CallOpts, addr)
}

// GetLastSlashVoted is a free data retrieval call binding the contract method 0x7e2a1ed3.
//
// Solidity: function getLastSlashVoted(address addr) view returns(int256)
func (_VssBase *VssBaseCallerSession) GetLastSlashVoted(addr common.Address) (*big.Int, error) {
	return _VssBase.Contract.GetLastSlashVoted(&_VssBase.CallOpts, addr)
}

// GetPrivateShares is a free data retrieval call binding the contract method 0x9496db22.
//
// Solidity: function getPrivateShares(address node) view returns(bytes)
func (_VssBase *VssBaseCaller) GetPrivateShares(opts *bind.CallOpts, node common.Address) ([]byte, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getPrivateShares", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPrivateShares is a free data retrieval call binding the contract method 0x9496db22.
//
// Solidity: function getPrivateShares(address node) view returns(bytes)
func (_VssBase *VssBaseSession) GetPrivateShares(node common.Address) ([]byte, error) {
	return _VssBase.Contract.GetPrivateShares(&_VssBase.CallOpts, node)
}

// GetPrivateShares is a free data retrieval call binding the contract method 0x9496db22.
//
// Solidity: function getPrivateShares(address node) view returns(bytes)
func (_VssBase *VssBaseCallerSession) GetPrivateShares(node common.Address) ([]byte, error) {
	return _VssBase.Contract.GetPrivateShares(&_VssBase.CallOpts, node)
}

// GetPublicShares is a free data retrieval call binding the contract method 0x80c3fe62.
//
// Solidity: function getPublicShares(address node) view returns(bytes)
func (_VssBase *VssBaseCaller) GetPublicShares(opts *bind.CallOpts, node common.Address) ([]byte, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getPublicShares", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPublicShares is a free data retrieval call binding the contract method 0x80c3fe62.
//
// Solidity: function getPublicShares(address node) view returns(bytes)
func (_VssBase *VssBaseSession) GetPublicShares(node common.Address) ([]byte, error) {
	return _VssBase.Contract.GetPublicShares(&_VssBase.CallOpts, node)
}

// GetPublicShares is a free data retrieval call binding the contract method 0x80c3fe62.
//
// Solidity: function getPublicShares(address node) view returns(bytes)
func (_VssBase *VssBaseCallerSession) GetPublicShares(node common.Address) ([]byte, error) {
	return _VssBase.Contract.GetPublicShares(&_VssBase.CallOpts, node)
}

// GetRevealedShare is a free data retrieval call binding the contract method 0x8b051563.
//
// Solidity: function getRevealedShare(int256 index) view returns((bytes,bytes,bytes,bytes,bytes,address,address))
func (_VssBase *VssBaseCaller) GetRevealedShare(opts *bind.CallOpts, index *big.Int) (VssBaseRevealedShare, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getRevealedShare", index)

	if err != nil {
		return *new(VssBaseRevealedShare), err
	}

	out0 := *abi.ConvertType(out[0], new(VssBaseRevealedShare)).(*VssBaseRevealedShare)

	return out0, err

}

// GetRevealedShare is a free data retrieval call binding the contract method 0x8b051563.
//
// Solidity: function getRevealedShare(int256 index) view returns((bytes,bytes,bytes,bytes,bytes,address,address))
func (_VssBase *VssBaseSession) GetRevealedShare(index *big.Int) (VssBaseRevealedShare, error) {
	return _VssBase.Contract.GetRevealedShare(&_VssBase.CallOpts, index)
}

// GetRevealedShare is a free data retrieval call binding the contract method 0x8b051563.
//
// Solidity: function getRevealedShare(int256 index) view returns((bytes,bytes,bytes,bytes,bytes,address,address))
func (_VssBase *VssBaseCallerSession) GetRevealedShare(index *big.Int) (VssBaseRevealedShare, error) {
	return _VssBase.Contract.GetRevealedShare(&_VssBase.CallOpts, index)
}

// GetVSSNodeIndex is a free data retrieval call binding the contract method 0xfc054aa8.
//
// Solidity: function getVSSNodeIndex(address scs) view returns(int256)
func (_VssBase *VssBaseCaller) GetVSSNodeIndex(opts *bind.CallOpts, scs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getVSSNodeIndex", scs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVSSNodeIndex is a free data retrieval call binding the contract method 0xfc054aa8.
//
// Solidity: function getVSSNodeIndex(address scs) view returns(int256)
func (_VssBase *VssBaseSession) GetVSSNodeIndex(scs common.Address) (*big.Int, error) {
	return _VssBase.Contract.GetVSSNodeIndex(&_VssBase.CallOpts, scs)
}

// GetVSSNodeIndex is a free data retrieval call binding the contract method 0xfc054aa8.
//
// Solidity: function getVSSNodeIndex(address scs) view returns(int256)
func (_VssBase *VssBaseCallerSession) GetVSSNodeIndex(scs common.Address) (*big.Int, error) {
	return _VssBase.Contract.GetVSSNodeIndex(&_VssBase.CallOpts, scs)
}

// GetVSSNodesIndexs is a free data retrieval call binding the contract method 0x2a292dc7.
//
// Solidity: function getVSSNodesIndexs(address[] nodes) view returns(int256[])
func (_VssBase *VssBaseCaller) GetVSSNodesIndexs(opts *bind.CallOpts, nodes []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getVSSNodesIndexs", nodes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVSSNodesIndexs is a free data retrieval call binding the contract method 0x2a292dc7.
//
// Solidity: function getVSSNodesIndexs(address[] nodes) view returns(int256[])
func (_VssBase *VssBaseSession) GetVSSNodesIndexs(nodes []common.Address) ([]*big.Int, error) {
	return _VssBase.Contract.GetVSSNodesIndexs(&_VssBase.CallOpts, nodes)
}

// GetVSSNodesIndexs is a free data retrieval call binding the contract method 0x2a292dc7.
//
// Solidity: function getVSSNodesIndexs(address[] nodes) view returns(int256[])
func (_VssBase *VssBaseCallerSession) GetVSSNodesIndexs(nodes []common.Address) ([]*big.Int, error) {
	return _VssBase.Contract.GetVSSNodesIndexs(&_VssBase.CallOpts, nodes)
}

// GetVSSNodesPubkey is a free data retrieval call binding the contract method 0x2512b947.
//
// Solidity: function getVSSNodesPubkey(address[] nodes) view returns(bytes32[])
func (_VssBase *VssBaseCaller) GetVSSNodesPubkey(opts *bind.CallOpts, nodes []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "getVSSNodesPubkey", nodes)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetVSSNodesPubkey is a free data retrieval call binding the contract method 0x2512b947.
//
// Solidity: function getVSSNodesPubkey(address[] nodes) view returns(bytes32[])
func (_VssBase *VssBaseSession) GetVSSNodesPubkey(nodes []common.Address) ([][32]byte, error) {
	return _VssBase.Contract.GetVSSNodesPubkey(&_VssBase.CallOpts, nodes)
}

// GetVSSNodesPubkey is a free data retrieval call binding the contract method 0x2512b947.
//
// Solidity: function getVSSNodesPubkey(address[] nodes) view returns(bytes32[])
func (_VssBase *VssBaseCallerSession) GetVSSNodesPubkey(nodes []common.Address) ([][32]byte, error) {
	return _VssBase.Contract.GetVSSNodesPubkey(&_VssBase.CallOpts, nodes)
}

// IsConfigReady is a free data retrieval call binding the contract method 0x91e4d7f7.
//
// Solidity: function isConfigReady(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseCaller) IsConfigReady(opts *bind.CallOpts, configVersion *big.Int) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "isConfigReady", configVersion)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfigReady is a free data retrieval call binding the contract method 0x91e4d7f7.
//
// Solidity: function isConfigReady(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseSession) IsConfigReady(configVersion *big.Int) (bool, error) {
	return _VssBase.Contract.IsConfigReady(&_VssBase.CallOpts, configVersion)
}

// IsConfigReady is a free data retrieval call binding the contract method 0x91e4d7f7.
//
// Solidity: function isConfigReady(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseCallerSession) IsConfigReady(configVersion *big.Int) (bool, error) {
	return _VssBase.Contract.IsConfigReady(&_VssBase.CallOpts, configVersion)
}

// IsSlashed is a free data retrieval call binding the contract method 0x9d6e8c6c.
//
// Solidity: function isSlashed(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseCaller) IsSlashed(opts *bind.CallOpts, configVersion *big.Int) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "isSlashed", configVersion)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0x9d6e8c6c.
//
// Solidity: function isSlashed(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseSession) IsSlashed(configVersion *big.Int) (bool, error) {
	return _VssBase.Contract.IsSlashed(&_VssBase.CallOpts, configVersion)
}

// IsSlashed is a free data retrieval call binding the contract method 0x9d6e8c6c.
//
// Solidity: function isSlashed(int256 configVersion) view returns(bool)
func (_VssBase *VssBaseCallerSession) IsSlashed(configVersion *big.Int) (bool, error) {
	return _VssBase.Contract.IsSlashed(&_VssBase.CallOpts, configVersion)
}

// LastConfigUpload is a free data retrieval call binding the contract method 0x2a56ce43.
//
// Solidity: function lastConfigUpload(address ) view returns(int256)
func (_VssBase *VssBaseCaller) LastConfigUpload(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastConfigUpload", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastConfigUpload is a free data retrieval call binding the contract method 0x2a56ce43.
//
// Solidity: function lastConfigUpload(address ) view returns(int256)
func (_VssBase *VssBaseSession) LastConfigUpload(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastConfigUpload(&_VssBase.CallOpts, arg0)
}

// LastConfigUpload is a free data retrieval call binding the contract method 0x2a56ce43.
//
// Solidity: function lastConfigUpload(address ) view returns(int256)
func (_VssBase *VssBaseCallerSession) LastConfigUpload(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastConfigUpload(&_VssBase.CallOpts, arg0)
}

// LastConfigUploadByBlock is a free data retrieval call binding the contract method 0xb5ef046e.
//
// Solidity: function lastConfigUploadByBlock(address ) view returns(int256)
func (_VssBase *VssBaseCaller) LastConfigUploadByBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastConfigUploadByBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastConfigUploadByBlock is a free data retrieval call binding the contract method 0xb5ef046e.
//
// Solidity: function lastConfigUploadByBlock(address ) view returns(int256)
func (_VssBase *VssBaseSession) LastConfigUploadByBlock(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastConfigUploadByBlock(&_VssBase.CallOpts, arg0)
}

// LastConfigUploadByBlock is a free data retrieval call binding the contract method 0xb5ef046e.
//
// Solidity: function lastConfigUploadByBlock(address ) view returns(int256)
func (_VssBase *VssBaseCallerSession) LastConfigUploadByBlock(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastConfigUploadByBlock(&_VssBase.CallOpts, arg0)
}

// LastNodeChangeBlock is a free data retrieval call binding the contract method 0x7d0f6c77.
//
// Solidity: function lastNodeChangeBlock() view returns(int256)
func (_VssBase *VssBaseCaller) LastNodeChangeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastNodeChangeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastNodeChangeBlock is a free data retrieval call binding the contract method 0x7d0f6c77.
//
// Solidity: function lastNodeChangeBlock() view returns(int256)
func (_VssBase *VssBaseSession) LastNodeChangeBlock() (*big.Int, error) {
	return _VssBase.Contract.LastNodeChangeBlock(&_VssBase.CallOpts)
}

// LastNodeChangeBlock is a free data retrieval call binding the contract method 0x7d0f6c77.
//
// Solidity: function lastNodeChangeBlock() view returns(int256)
func (_VssBase *VssBaseCallerSession) LastNodeChangeBlock() (*big.Int, error) {
	return _VssBase.Contract.LastNodeChangeBlock(&_VssBase.CallOpts)
}

// LastNodeChangeConfigVersion is a free data retrieval call binding the contract method 0x514cac35.
//
// Solidity: function lastNodeChangeConfigVersion() view returns(int256)
func (_VssBase *VssBaseCaller) LastNodeChangeConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastNodeChangeConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastNodeChangeConfigVersion is a free data retrieval call binding the contract method 0x514cac35.
//
// Solidity: function lastNodeChangeConfigVersion() view returns(int256)
func (_VssBase *VssBaseSession) LastNodeChangeConfigVersion() (*big.Int, error) {
	return _VssBase.Contract.LastNodeChangeConfigVersion(&_VssBase.CallOpts)
}

// LastNodeChangeConfigVersion is a free data retrieval call binding the contract method 0x514cac35.
//
// Solidity: function lastNodeChangeConfigVersion() view returns(int256)
func (_VssBase *VssBaseCallerSession) LastNodeChangeConfigVersion() (*big.Int, error) {
	return _VssBase.Contract.LastNodeChangeConfigVersion(&_VssBase.CallOpts)
}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(address)
func (_VssBase *VssBaseCaller) LastSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(address)
func (_VssBase *VssBaseSession) LastSender() (common.Address, error) {
	return _VssBase.Contract.LastSender(&_VssBase.CallOpts)
}

// LastSender is a free data retrieval call binding the contract method 0x256fec88.
//
// Solidity: function lastSender() view returns(address)
func (_VssBase *VssBaseCallerSession) LastSender() (common.Address, error) {
	return _VssBase.Contract.LastSender(&_VssBase.CallOpts)
}

// LastSlashingVoted is a free data retrieval call binding the contract method 0x7457b355.
//
// Solidity: function lastSlashingVoted(address ) view returns(int256)
func (_VssBase *VssBaseCaller) LastSlashingVoted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "lastSlashingVoted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSlashingVoted is a free data retrieval call binding the contract method 0x7457b355.
//
// Solidity: function lastSlashingVoted(address ) view returns(int256)
func (_VssBase *VssBaseSession) LastSlashingVoted(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastSlashingVoted(&_VssBase.CallOpts, arg0)
}

// LastSlashingVoted is a free data retrieval call binding the contract method 0x7457b355.
//
// Solidity: function lastSlashingVoted(address ) view returns(int256)
func (_VssBase *VssBaseCallerSession) LastSlashingVoted(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.LastSlashingVoted(&_VssBase.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VssBase *VssBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VssBase *VssBaseSession) Owner() (common.Address, error) {
	return _VssBase.Contract.Owner(&_VssBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VssBase *VssBaseCallerSession) Owner() (common.Address, error) {
	return _VssBase.Contract.Owner(&_VssBase.CallOpts)
}

// ResetVSSGroup is a free data retrieval call binding the contract method 0x1a1efdaf.
//
// Solidity: function resetVSSGroup() view returns()
func (_VssBase *VssBaseCaller) ResetVSSGroup(opts *bind.CallOpts) error {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "resetVSSGroup")

	if err != nil {
		return err
	}

	return err

}

// ResetVSSGroup is a free data retrieval call binding the contract method 0x1a1efdaf.
//
// Solidity: function resetVSSGroup() view returns()
func (_VssBase *VssBaseSession) ResetVSSGroup() error {
	return _VssBase.Contract.ResetVSSGroup(&_VssBase.CallOpts)
}

// ResetVSSGroup is a free data retrieval call binding the contract method 0x1a1efdaf.
//
// Solidity: function resetVSSGroup() view returns()
func (_VssBase *VssBaseCallerSession) ResetVSSGroup() error {
	return _VssBase.Contract.ResetVSSGroup(&_VssBase.CallOpts)
}

// RevealIndex is a free data retrieval call binding the contract method 0x99ba5a92.
//
// Solidity: function revealIndex() view returns(int256)
func (_VssBase *VssBaseCaller) RevealIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "revealIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealIndex is a free data retrieval call binding the contract method 0x99ba5a92.
//
// Solidity: function revealIndex() view returns(int256)
func (_VssBase *VssBaseSession) RevealIndex() (*big.Int, error) {
	return _VssBase.Contract.RevealIndex(&_VssBase.CallOpts)
}

// RevealIndex is a free data retrieval call binding the contract method 0x99ba5a92.
//
// Solidity: function revealIndex() view returns(int256)
func (_VssBase *VssBaseCallerSession) RevealIndex() (*big.Int, error) {
	return _VssBase.Contract.RevealIndex(&_VssBase.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x0b927b32.
//
// Solidity: function revealed(bytes32 ) view returns(bool)
func (_VssBase *VssBaseCaller) Revealed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "revealed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revealed is a free data retrieval call binding the contract method 0x0b927b32.
//
// Solidity: function revealed(bytes32 ) view returns(bool)
func (_VssBase *VssBaseSession) Revealed(arg0 [32]byte) (bool, error) {
	return _VssBase.Contract.Revealed(&_VssBase.CallOpts, arg0)
}

// Revealed is a free data retrieval call binding the contract method 0x0b927b32.
//
// Solidity: function revealed(bytes32 ) view returns(bool)
func (_VssBase *VssBaseCallerSession) Revealed(arg0 [32]byte) (bool, error) {
	return _VssBase.Contract.Revealed(&_VssBase.CallOpts, arg0)
}

// RevealedReporter is a free data retrieval call binding the contract method 0x307f201a.
//
// Solidity: function revealedReporter(int256 ) view returns(address)
func (_VssBase *VssBaseCaller) RevealedReporter(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "revealedReporter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevealedReporter is a free data retrieval call binding the contract method 0x307f201a.
//
// Solidity: function revealedReporter(int256 ) view returns(address)
func (_VssBase *VssBaseSession) RevealedReporter(arg0 *big.Int) (common.Address, error) {
	return _VssBase.Contract.RevealedReporter(&_VssBase.CallOpts, arg0)
}

// RevealedReporter is a free data retrieval call binding the contract method 0x307f201a.
//
// Solidity: function revealedReporter(int256 ) view returns(address)
func (_VssBase *VssBaseCallerSession) RevealedReporter(arg0 *big.Int) (common.Address, error) {
	return _VssBase.Contract.RevealedReporter(&_VssBase.CallOpts, arg0)
}

// RevealedSigMapping is a free data retrieval call binding the contract method 0x59495523.
//
// Solidity: function revealedSigMapping(bytes32 ) view returns(int256)
func (_VssBase *VssBaseCaller) RevealedSigMapping(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "revealedSigMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealedSigMapping is a free data retrieval call binding the contract method 0x59495523.
//
// Solidity: function revealedSigMapping(bytes32 ) view returns(int256)
func (_VssBase *VssBaseSession) RevealedSigMapping(arg0 [32]byte) (*big.Int, error) {
	return _VssBase.Contract.RevealedSigMapping(&_VssBase.CallOpts, arg0)
}

// RevealedSigMapping is a free data retrieval call binding the contract method 0x59495523.
//
// Solidity: function revealedSigMapping(bytes32 ) view returns(int256)
func (_VssBase *VssBaseCallerSession) RevealedSigMapping(arg0 [32]byte) (*big.Int, error) {
	return _VssBase.Contract.RevealedSigMapping(&_VssBase.CallOpts, arg0)
}

// RevealedViolator is a free data retrieval call binding the contract method 0x51ed07e4.
//
// Solidity: function revealedViolator(int256 ) view returns(address)
func (_VssBase *VssBaseCaller) RevealedViolator(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "revealedViolator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevealedViolator is a free data retrieval call binding the contract method 0x51ed07e4.
//
// Solidity: function revealedViolator(int256 ) view returns(address)
func (_VssBase *VssBaseSession) RevealedViolator(arg0 *big.Int) (common.Address, error) {
	return _VssBase.Contract.RevealedViolator(&_VssBase.CallOpts, arg0)
}

// RevealedViolator is a free data retrieval call binding the contract method 0x51ed07e4.
//
// Solidity: function revealedViolator(int256 ) view returns(address)
func (_VssBase *VssBaseCallerSession) RevealedViolator(arg0 *big.Int) (common.Address, error) {
	return _VssBase.Contract.RevealedViolator(&_VssBase.CallOpts, arg0)
}

// Reveals is a free data retrieval call binding the contract method 0x778ada10.
//
// Solidity: function reveals(int256 ) view returns(bytes PubShare, bytes PubSig, bytes PriShare, bytes PriSig, bytes Revealed, address Violator, address Whistleblower)
func (_VssBase *VssBaseCaller) Reveals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PubShare      []byte
	PubSig        []byte
	PriShare      []byte
	PriSig        []byte
	Revealed      []byte
	Violator      common.Address
	Whistleblower common.Address
}, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "reveals", arg0)

	outstruct := new(struct {
		PubShare      []byte
		PubSig        []byte
		PriShare      []byte
		PriSig        []byte
		Revealed      []byte
		Violator      common.Address
		Whistleblower common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PubShare = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.PubSig = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.PriShare = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.PriSig = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Revealed = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Violator = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Whistleblower = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Reveals is a free data retrieval call binding the contract method 0x778ada10.
//
// Solidity: function reveals(int256 ) view returns(bytes PubShare, bytes PubSig, bytes PriShare, bytes PriSig, bytes Revealed, address Violator, address Whistleblower)
func (_VssBase *VssBaseSession) Reveals(arg0 *big.Int) (struct {
	PubShare      []byte
	PubSig        []byte
	PriShare      []byte
	PriSig        []byte
	Revealed      []byte
	Violator      common.Address
	Whistleblower common.Address
}, error) {
	return _VssBase.Contract.Reveals(&_VssBase.CallOpts, arg0)
}

// Reveals is a free data retrieval call binding the contract method 0x778ada10.
//
// Solidity: function reveals(int256 ) view returns(bytes PubShare, bytes PubSig, bytes PriShare, bytes PriSig, bytes Revealed, address Violator, address Whistleblower)
func (_VssBase *VssBaseCallerSession) Reveals(arg0 *big.Int) (struct {
	PubShare      []byte
	PubSig        []byte
	PriShare      []byte
	PriSig        []byte
	Revealed      []byte
	Violator      common.Address
	Whistleblower common.Address
}, error) {
	return _VssBase.Contract.Reveals(&_VssBase.CallOpts, arg0)
}

// Slashed is a free data retrieval call binding the contract method 0x1c2c6b4d.
//
// Solidity: function slashed(int256 ) view returns(int256)
func (_VssBase *VssBaseCaller) Slashed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slashed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slashed is a free data retrieval call binding the contract method 0x1c2c6b4d.
//
// Solidity: function slashed(int256 ) view returns(int256)
func (_VssBase *VssBaseSession) Slashed(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.Slashed(&_VssBase.CallOpts, arg0)
}

// Slashed is a free data retrieval call binding the contract method 0x1c2c6b4d.
//
// Solidity: function slashed(int256 ) view returns(int256)
func (_VssBase *VssBaseCallerSession) Slashed(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.Slashed(&_VssBase.CallOpts, arg0)
}

// SlashingRejects is a free data retrieval call binding the contract method 0x5e296853.
//
// Solidity: function slashingRejects(int256 ) view returns(int256)
func (_VssBase *VssBaseCaller) SlashingRejects(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slashingRejects", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingRejects is a free data retrieval call binding the contract method 0x5e296853.
//
// Solidity: function slashingRejects(int256 ) view returns(int256)
func (_VssBase *VssBaseSession) SlashingRejects(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlashingRejects(&_VssBase.CallOpts, arg0)
}

// SlashingRejects is a free data retrieval call binding the contract method 0x5e296853.
//
// Solidity: function slashingRejects(int256 ) view returns(int256)
func (_VssBase *VssBaseCallerSession) SlashingRejects(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlashingRejects(&_VssBase.CallOpts, arg0)
}

// SlashingVoters is a free data retrieval call binding the contract method 0x3430cadd.
//
// Solidity: function slashingVoters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseCaller) SlashingVoters(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slashingVoters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashingVoters is a free data retrieval call binding the contract method 0x3430cadd.
//
// Solidity: function slashingVoters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseSession) SlashingVoters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VssBase.Contract.SlashingVoters(&_VssBase.CallOpts, arg0, arg1)
}

// SlashingVoters is a free data retrieval call binding the contract method 0x3430cadd.
//
// Solidity: function slashingVoters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseCallerSession) SlashingVoters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VssBase.Contract.SlashingVoters(&_VssBase.CallOpts, arg0, arg1)
}

// SlashingVotes is a free data retrieval call binding the contract method 0x94a629be.
//
// Solidity: function slashingVotes(int256 ) view returns(int256)
func (_VssBase *VssBaseCaller) SlashingVotes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slashingVotes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingVotes is a free data retrieval call binding the contract method 0x94a629be.
//
// Solidity: function slashingVotes(int256 ) view returns(int256)
func (_VssBase *VssBaseSession) SlashingVotes(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlashingVotes(&_VssBase.CallOpts, arg0)
}

// SlashingVotes is a free data retrieval call binding the contract method 0x94a629be.
//
// Solidity: function slashingVotes(int256 ) view returns(int256)
func (_VssBase *VssBaseCallerSession) SlashingVotes(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlashingVotes(&_VssBase.CallOpts, arg0)
}

// SlowNodeThreshold is a free data retrieval call binding the contract method 0x3ed4f116.
//
// Solidity: function slowNodeThreshold() view returns(int256)
func (_VssBase *VssBaseCaller) SlowNodeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slowNodeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlowNodeThreshold is a free data retrieval call binding the contract method 0x3ed4f116.
//
// Solidity: function slowNodeThreshold() view returns(int256)
func (_VssBase *VssBaseSession) SlowNodeThreshold() (*big.Int, error) {
	return _VssBase.Contract.SlowNodeThreshold(&_VssBase.CallOpts)
}

// SlowNodeThreshold is a free data retrieval call binding the contract method 0x3ed4f116.
//
// Solidity: function slowNodeThreshold() view returns(int256)
func (_VssBase *VssBaseCallerSession) SlowNodeThreshold() (*big.Int, error) {
	return _VssBase.Contract.SlowNodeThreshold(&_VssBase.CallOpts)
}

// SlowNodeVoted is a free data retrieval call binding the contract method 0xdbccad54.
//
// Solidity: function slowNodeVoted(address , int256 , address ) view returns(bool)
func (_VssBase *VssBaseCaller) SlowNodeVoted(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slowNodeVoted", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlowNodeVoted is a free data retrieval call binding the contract method 0xdbccad54.
//
// Solidity: function slowNodeVoted(address , int256 , address ) view returns(bool)
func (_VssBase *VssBaseSession) SlowNodeVoted(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _VssBase.Contract.SlowNodeVoted(&_VssBase.CallOpts, arg0, arg1, arg2)
}

// SlowNodeVoted is a free data retrieval call binding the contract method 0xdbccad54.
//
// Solidity: function slowNodeVoted(address , int256 , address ) view returns(bool)
func (_VssBase *VssBaseCallerSession) SlowNodeVoted(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (bool, error) {
	return _VssBase.Contract.SlowNodeVoted(&_VssBase.CallOpts, arg0, arg1, arg2)
}

// SlowNodeVotes is a free data retrieval call binding the contract method 0x6911af65.
//
// Solidity: function slowNodeVotes(address , int256 ) view returns(uint256)
func (_VssBase *VssBaseCaller) SlowNodeVotes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "slowNodeVotes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlowNodeVotes is a free data retrieval call binding the contract method 0x6911af65.
//
// Solidity: function slowNodeVotes(address , int256 ) view returns(uint256)
func (_VssBase *VssBaseSession) SlowNodeVotes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlowNodeVotes(&_VssBase.CallOpts, arg0, arg1)
}

// SlowNodeVotes is a free data retrieval call binding the contract method 0x6911af65.
//
// Solidity: function slowNodeVotes(address , int256 ) view returns(uint256)
func (_VssBase *VssBaseCallerSession) SlowNodeVotes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.SlowNodeVotes(&_VssBase.CallOpts, arg0, arg1)
}

// Voters is a free data retrieval call binding the contract method 0x9c081852.
//
// Solidity: function voters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseCaller) Voters(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "voters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0x9c081852.
//
// Solidity: function voters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseSession) Voters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VssBase.Contract.Voters(&_VssBase.CallOpts, arg0, arg1)
}

// Voters is a free data retrieval call binding the contract method 0x9c081852.
//
// Solidity: function voters(int256 , address ) view returns(bool)
func (_VssBase *VssBaseCallerSession) Voters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _VssBase.Contract.Voters(&_VssBase.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xd0bef4ae.
//
// Solidity: function votes(int256 ) view returns(int256)
func (_VssBase *VssBaseCaller) Votes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd0bef4ae.
//
// Solidity: function votes(int256 ) view returns(int256)
func (_VssBase *VssBaseSession) Votes(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.Votes(&_VssBase.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd0bef4ae.
//
// Solidity: function votes(int256 ) view returns(int256)
func (_VssBase *VssBaseCallerSession) Votes(arg0 *big.Int) (*big.Int, error) {
	return _VssBase.Contract.Votes(&_VssBase.CallOpts, arg0)
}

// VssConfigVersion is a free data retrieval call binding the contract method 0xbbc3d237.
//
// Solidity: function vssConfigVersion() view returns(int256)
func (_VssBase *VssBaseCaller) VssConfigVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "vssConfigVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VssConfigVersion is a free data retrieval call binding the contract method 0xbbc3d237.
//
// Solidity: function vssConfigVersion() view returns(int256)
func (_VssBase *VssBaseSession) VssConfigVersion() (*big.Int, error) {
	return _VssBase.Contract.VssConfigVersion(&_VssBase.CallOpts)
}

// VssConfigVersion is a free data retrieval call binding the contract method 0xbbc3d237.
//
// Solidity: function vssConfigVersion() view returns(int256)
func (_VssBase *VssBaseCallerSession) VssConfigVersion() (*big.Int, error) {
	return _VssBase.Contract.VssConfigVersion(&_VssBase.CallOpts)
}

// VssNodeCount is a free data retrieval call binding the contract method 0xcc30a60f.
//
// Solidity: function vssNodeCount() view returns(int256)
func (_VssBase *VssBaseCaller) VssNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "vssNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VssNodeCount is a free data retrieval call binding the contract method 0xcc30a60f.
//
// Solidity: function vssNodeCount() view returns(int256)
func (_VssBase *VssBaseSession) VssNodeCount() (*big.Int, error) {
	return _VssBase.Contract.VssNodeCount(&_VssBase.CallOpts)
}

// VssNodeCount is a free data retrieval call binding the contract method 0xcc30a60f.
//
// Solidity: function vssNodeCount() view returns(int256)
func (_VssBase *VssBaseCallerSession) VssNodeCount() (*big.Int, error) {
	return _VssBase.Contract.VssNodeCount(&_VssBase.CallOpts)
}

// VssNodeMemberships is a free data retrieval call binding the contract method 0x98b72f98.
//
// Solidity: function vssNodeMemberships(address ) view returns(uint256)
func (_VssBase *VssBaseCaller) VssNodeMemberships(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "vssNodeMemberships", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VssNodeMemberships is a free data retrieval call binding the contract method 0x98b72f98.
//
// Solidity: function vssNodeMemberships(address ) view returns(uint256)
func (_VssBase *VssBaseSession) VssNodeMemberships(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.VssNodeMemberships(&_VssBase.CallOpts, arg0)
}

// VssNodeMemberships is a free data retrieval call binding the contract method 0x98b72f98.
//
// Solidity: function vssNodeMemberships(address ) view returns(uint256)
func (_VssBase *VssBaseCallerSession) VssNodeMemberships(arg0 common.Address) (*big.Int, error) {
	return _VssBase.Contract.VssNodeMemberships(&_VssBase.CallOpts, arg0)
}

// VssThreshold is a free data retrieval call binding the contract method 0x984624ba.
//
// Solidity: function vssThreshold() view returns(int256)
func (_VssBase *VssBaseCaller) VssThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VssBase.contract.Call(opts, &out, "vssThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VssThreshold is a free data retrieval call binding the contract method 0x984624ba.
//
// Solidity: function vssThreshold() view returns(int256)
func (_VssBase *VssBaseSession) VssThreshold() (*big.Int, error) {
	return _VssBase.Contract.VssThreshold(&_VssBase.CallOpts)
}

// VssThreshold is a free data retrieval call binding the contract method 0x984624ba.
//
// Solidity: function vssThreshold() view returns(int256)
func (_VssBase *VssBaseCallerSession) VssThreshold() (*big.Int, error) {
	return _VssBase.Contract.VssThreshold(&_VssBase.CallOpts)
}

// ActivateVSS is a paid mutator transaction binding the contract method 0xe4c1de98.
//
// Solidity: function activateVSS(address sender) returns()
func (_VssBase *VssBaseTransactor) ActivateVSS(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "activateVSS", sender)
}

// ActivateVSS is a paid mutator transaction binding the contract method 0xe4c1de98.
//
// Solidity: function activateVSS(address sender) returns()
func (_VssBase *VssBaseSession) ActivateVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.ActivateVSS(&_VssBase.TransactOpts, sender)
}

// ActivateVSS is a paid mutator transaction binding the contract method 0xe4c1de98.
//
// Solidity: function activateVSS(address sender) returns()
func (_VssBase *VssBaseTransactorSession) ActivateVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.ActivateVSS(&_VssBase.TransactOpts, sender)
}

// DeactivateVSS is a paid mutator transaction binding the contract method 0xde79b856.
//
// Solidity: function deactivateVSS(address sender) returns()
func (_VssBase *VssBaseTransactor) DeactivateVSS(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "deactivateVSS", sender)
}

// DeactivateVSS is a paid mutator transaction binding the contract method 0xde79b856.
//
// Solidity: function deactivateVSS(address sender) returns()
func (_VssBase *VssBaseSession) DeactivateVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.DeactivateVSS(&_VssBase.TransactOpts, sender)
}

// DeactivateVSS is a paid mutator transaction binding the contract method 0xde79b856.
//
// Solidity: function deactivateVSS(address sender) returns()
func (_VssBase *VssBaseTransactorSession) DeactivateVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.DeactivateVSS(&_VssBase.TransactOpts, sender)
}

// RegisterVSS is a paid mutator transaction binding the contract method 0xeefb4227.
//
// Solidity: function registerVSS(address sender, bytes32 publickey) returns()
func (_VssBase *VssBaseTransactor) RegisterVSS(opts *bind.TransactOpts, sender common.Address, publickey [32]byte) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "registerVSS", sender, publickey)
}

// RegisterVSS is a paid mutator transaction binding the contract method 0xeefb4227.
//
// Solidity: function registerVSS(address sender, bytes32 publickey) returns()
func (_VssBase *VssBaseSession) RegisterVSS(sender common.Address, publickey [32]byte) (*types.Transaction, error) {
	return _VssBase.Contract.RegisterVSS(&_VssBase.TransactOpts, sender, publickey)
}

// RegisterVSS is a paid mutator transaction binding the contract method 0xeefb4227.
//
// Solidity: function registerVSS(address sender, bytes32 publickey) returns()
func (_VssBase *VssBaseTransactorSession) RegisterVSS(sender common.Address, publickey [32]byte) (*types.Transaction, error) {
	return _VssBase.Contract.RegisterVSS(&_VssBase.TransactOpts, sender, publickey)
}

// ReportSlowNode is a paid mutator transaction binding the contract method 0x7ce54cc8.
//
// Solidity: function reportSlowNode(address violator) returns()
func (_VssBase *VssBaseTransactor) ReportSlowNode(opts *bind.TransactOpts, violator common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "reportSlowNode", violator)
}

// ReportSlowNode is a paid mutator transaction binding the contract method 0x7ce54cc8.
//
// Solidity: function reportSlowNode(address violator) returns()
func (_VssBase *VssBaseSession) ReportSlowNode(violator common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.ReportSlowNode(&_VssBase.TransactOpts, violator)
}

// ReportSlowNode is a paid mutator transaction binding the contract method 0x7ce54cc8.
//
// Solidity: function reportSlowNode(address violator) returns()
func (_VssBase *VssBaseTransactorSession) ReportSlowNode(violator common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.ReportSlowNode(&_VssBase.TransactOpts, violator)
}

// Reveal is a paid mutator transaction binding the contract method 0x282c0b8f.
//
// Solidity: function reveal(address violator, bytes pubShare, bytes pubSig, bytes priShare, bytes priSig, bytes revealedPri) returns()
func (_VssBase *VssBaseTransactor) Reveal(opts *bind.TransactOpts, violator common.Address, pubShare []byte, pubSig []byte, priShare []byte, priSig []byte, revealedPri []byte) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "reveal", violator, pubShare, pubSig, priShare, priSig, revealedPri)
}

// Reveal is a paid mutator transaction binding the contract method 0x282c0b8f.
//
// Solidity: function reveal(address violator, bytes pubShare, bytes pubSig, bytes priShare, bytes priSig, bytes revealedPri) returns()
func (_VssBase *VssBaseSession) Reveal(violator common.Address, pubShare []byte, pubSig []byte, priShare []byte, priSig []byte, revealedPri []byte) (*types.Transaction, error) {
	return _VssBase.Contract.Reveal(&_VssBase.TransactOpts, violator, pubShare, pubSig, priShare, priSig, revealedPri)
}

// Reveal is a paid mutator transaction binding the contract method 0x282c0b8f.
//
// Solidity: function reveal(address violator, bytes pubShare, bytes pubSig, bytes priShare, bytes priSig, bytes revealedPri) returns()
func (_VssBase *VssBaseTransactorSession) Reveal(violator common.Address, pubShare []byte, pubSig []byte, priShare []byte, priSig []byte, revealedPri []byte) (*types.Transaction, error) {
	return _VssBase.Contract.Reveal(&_VssBase.TransactOpts, violator, pubShare, pubSig, priShare, priSig, revealedPri)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address callerAddr) returns()
func (_VssBase *VssBaseTransactor) SetCaller(opts *bind.TransactOpts, callerAddr common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "setCaller", callerAddr)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address callerAddr) returns()
func (_VssBase *VssBaseSession) SetCaller(callerAddr common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.SetCaller(&_VssBase.TransactOpts, callerAddr)
}

// SetCaller is a paid mutator transaction binding the contract method 0xbeb92f55.
//
// Solidity: function setCaller(address callerAddr) returns()
func (_VssBase *VssBaseTransactorSession) SetCaller(callerAddr common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.SetCaller(&_VssBase.TransactOpts, callerAddr)
}

// SetGeneralAttributes is a paid mutator transaction binding the contract method 0xcba3f588.
//
// Solidity: function setGeneralAttributes(bytes32 namespace, int256 key, bytes value) returns()
func (_VssBase *VssBaseTransactor) SetGeneralAttributes(opts *bind.TransactOpts, namespace [32]byte, key *big.Int, value []byte) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "setGeneralAttributes", namespace, key, value)
}

// SetGeneralAttributes is a paid mutator transaction binding the contract method 0xcba3f588.
//
// Solidity: function setGeneralAttributes(bytes32 namespace, int256 key, bytes value) returns()
func (_VssBase *VssBaseSession) SetGeneralAttributes(namespace [32]byte, key *big.Int, value []byte) (*types.Transaction, error) {
	return _VssBase.Contract.SetGeneralAttributes(&_VssBase.TransactOpts, namespace, key, value)
}

// SetGeneralAttributes is a paid mutator transaction binding the contract method 0xcba3f588.
//
// Solidity: function setGeneralAttributes(bytes32 namespace, int256 key, bytes value) returns()
func (_VssBase *VssBaseTransactorSession) SetGeneralAttributes(namespace [32]byte, key *big.Int, value []byte) (*types.Transaction, error) {
	return _VssBase.Contract.SetGeneralAttributes(&_VssBase.TransactOpts, namespace, key, value)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_VssBase *VssBaseTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_VssBase *VssBaseSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.SetOwner(&_VssBase.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_VssBase *VssBaseTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.SetOwner(&_VssBase.TransactOpts, newOwner)
}

// SetSlowNodeThreshold is a paid mutator transaction binding the contract method 0xcd199420.
//
// Solidity: function setSlowNodeThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseTransactor) SetSlowNodeThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "setSlowNodeThreshold", newThreshold)
}

// SetSlowNodeThreshold is a paid mutator transaction binding the contract method 0xcd199420.
//
// Solidity: function setSlowNodeThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseSession) SetSlowNodeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.SetSlowNodeThreshold(&_VssBase.TransactOpts, newThreshold)
}

// SetSlowNodeThreshold is a paid mutator transaction binding the contract method 0xcd199420.
//
// Solidity: function setSlowNodeThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseTransactorSession) SetSlowNodeThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.SetSlowNodeThreshold(&_VssBase.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.SetThreshold(&_VssBase.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xef0fb558.
//
// Solidity: function setThreshold(int256 newThreshold) returns()
func (_VssBase *VssBaseTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.SetThreshold(&_VssBase.TransactOpts, newThreshold)
}

// Slashing is a paid mutator transaction binding the contract method 0x56a09939.
//
// Solidity: function slashing(int256 index, bool slash) returns()
func (_VssBase *VssBaseTransactor) Slashing(opts *bind.TransactOpts, index *big.Int, slash bool) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "slashing", index, slash)
}

// Slashing is a paid mutator transaction binding the contract method 0x56a09939.
//
// Solidity: function slashing(int256 index, bool slash) returns()
func (_VssBase *VssBaseSession) Slashing(index *big.Int, slash bool) (*types.Transaction, error) {
	return _VssBase.Contract.Slashing(&_VssBase.TransactOpts, index, slash)
}

// Slashing is a paid mutator transaction binding the contract method 0x56a09939.
//
// Solidity: function slashing(int256 index, bool slash) returns()
func (_VssBase *VssBaseTransactorSession) Slashing(index *big.Int, slash bool) (*types.Transaction, error) {
	return _VssBase.Contract.Slashing(&_VssBase.TransactOpts, index, slash)
}

// UnregisterVSS is a paid mutator transaction binding the contract method 0xf06dc92d.
//
// Solidity: function unregisterVSS(address sender) returns()
func (_VssBase *VssBaseTransactor) UnregisterVSS(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "unregisterVSS", sender)
}

// UnregisterVSS is a paid mutator transaction binding the contract method 0xf06dc92d.
//
// Solidity: function unregisterVSS(address sender) returns()
func (_VssBase *VssBaseSession) UnregisterVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.UnregisterVSS(&_VssBase.TransactOpts, sender)
}

// UnregisterVSS is a paid mutator transaction binding the contract method 0xf06dc92d.
//
// Solidity: function unregisterVSS(address sender) returns()
func (_VssBase *VssBaseTransactorSession) UnregisterVSS(sender common.Address) (*types.Transaction, error) {
	return _VssBase.Contract.UnregisterVSS(&_VssBase.TransactOpts, sender)
}

// UploadVSSConfig is a paid mutator transaction binding the contract method 0xca6c8a31.
//
// Solidity: function uploadVSSConfig(bytes publicShares, bytes privateShares) returns()
func (_VssBase *VssBaseTransactor) UploadVSSConfig(opts *bind.TransactOpts, publicShares []byte, privateShares []byte) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "uploadVSSConfig", publicShares, privateShares)
}

// UploadVSSConfig is a paid mutator transaction binding the contract method 0xca6c8a31.
//
// Solidity: function uploadVSSConfig(bytes publicShares, bytes privateShares) returns()
func (_VssBase *VssBaseSession) UploadVSSConfig(publicShares []byte, privateShares []byte) (*types.Transaction, error) {
	return _VssBase.Contract.UploadVSSConfig(&_VssBase.TransactOpts, publicShares, privateShares)
}

// UploadVSSConfig is a paid mutator transaction binding the contract method 0xca6c8a31.
//
// Solidity: function uploadVSSConfig(bytes publicShares, bytes privateShares) returns()
func (_VssBase *VssBaseTransactorSession) UploadVSSConfig(publicShares []byte, privateShares []byte) (*types.Transaction, error) {
	return _VssBase.Contract.UploadVSSConfig(&_VssBase.TransactOpts, publicShares, privateShares)
}

// Vote is a paid mutator transaction binding the contract method 0x3891c320.
//
// Solidity: function vote(int256 configVersion) returns()
func (_VssBase *VssBaseTransactor) Vote(opts *bind.TransactOpts, configVersion *big.Int) (*types.Transaction, error) {
	return _VssBase.contract.Transact(opts, "vote", configVersion)
}

// Vote is a paid mutator transaction binding the contract method 0x3891c320.
//
// Solidity: function vote(int256 configVersion) returns()
func (_VssBase *VssBaseSession) Vote(configVersion *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.Vote(&_VssBase.TransactOpts, configVersion)
}

// Vote is a paid mutator transaction binding the contract method 0x3891c320.
//
// Solidity: function vote(int256 configVersion) returns()
func (_VssBase *VssBaseTransactorSession) Vote(configVersion *big.Int) (*types.Transaction, error) {
	return _VssBase.Contract.Vote(&_VssBase.TransactOpts, configVersion)
}
