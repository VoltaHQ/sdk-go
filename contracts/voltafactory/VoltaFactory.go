// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voltafactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VoltafactoryMetaData contains all meta data concerning the Voltafactory contract.
var VoltafactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKeyValidatorImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executorImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minNumSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractVoltaAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKeyValidatorImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executorImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minNumSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"name\":\"setEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoltafactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VoltafactoryMetaData.ABI instead.
var VoltafactoryABI = VoltafactoryMetaData.ABI

// Voltafactory is an auto generated Go binding around an Ethereum contract.
type Voltafactory struct {
	VoltafactoryCaller     // Read-only binding to the contract
	VoltafactoryTransactor // Write-only binding to the contract
	VoltafactoryFilterer   // Log filterer for contract events
}

// VoltafactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoltafactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltafactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoltafactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltafactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoltafactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltafactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoltafactorySession struct {
	Contract     *Voltafactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoltafactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoltafactoryCallerSession struct {
	Contract *VoltafactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoltafactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoltafactoryTransactorSession struct {
	Contract     *VoltafactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoltafactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoltafactoryRaw struct {
	Contract *Voltafactory // Generic contract binding to access the raw methods on
}

// VoltafactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoltafactoryCallerRaw struct {
	Contract *VoltafactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VoltafactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoltafactoryTransactorRaw struct {
	Contract *VoltafactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoltafactory creates a new instance of Voltafactory, bound to a specific deployed contract.
func NewVoltafactory(address common.Address, backend bind.ContractBackend) (*Voltafactory, error) {
	contract, err := bindVoltafactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voltafactory{VoltafactoryCaller: VoltafactoryCaller{contract: contract}, VoltafactoryTransactor: VoltafactoryTransactor{contract: contract}, VoltafactoryFilterer: VoltafactoryFilterer{contract: contract}}, nil
}

// NewVoltafactoryCaller creates a new read-only instance of Voltafactory, bound to a specific deployed contract.
func NewVoltafactoryCaller(address common.Address, caller bind.ContractCaller) (*VoltafactoryCaller, error) {
	contract, err := bindVoltafactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryCaller{contract: contract}, nil
}

// NewVoltafactoryTransactor creates a new write-only instance of Voltafactory, bound to a specific deployed contract.
func NewVoltafactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VoltafactoryTransactor, error) {
	contract, err := bindVoltafactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryTransactor{contract: contract}, nil
}

// NewVoltafactoryFilterer creates a new log filterer instance of Voltafactory, bound to a specific deployed contract.
func NewVoltafactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VoltafactoryFilterer, error) {
	contract, err := bindVoltafactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryFilterer{contract: contract}, nil
}

// bindVoltafactory binds a generic wrapper to an already deployed contract.
func bindVoltafactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoltafactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voltafactory *VoltafactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voltafactory.Contract.VoltafactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voltafactory *VoltafactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.Contract.VoltafactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voltafactory *VoltafactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voltafactory.Contract.VoltafactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voltafactory *VoltafactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voltafactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voltafactory *VoltafactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voltafactory *VoltafactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voltafactory.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voltafactory *VoltafactoryCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voltafactory.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voltafactory *VoltafactorySession) EntryPoint() (common.Address, error) {
	return _Voltafactory.Contract.EntryPoint(&_Voltafactory.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Voltafactory *VoltafactoryCallerSession) EntryPoint() (common.Address, error) {
	return _Voltafactory.Contract.EntryPoint(&_Voltafactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Voltafactory *VoltafactoryCaller) GetAddress(opts *bind.CallOpts, accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Voltafactory.contract.Call(opts, &out, "getAddress", accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Voltafactory *VoltafactorySession) GetAddress(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	return _Voltafactory.Contract.GetAddress(&_Voltafactory.CallOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Voltafactory *VoltafactoryCallerSession) GetAddress(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	return _Voltafactory.Contract.GetAddress(&_Voltafactory.CallOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Voltafactory *VoltafactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voltafactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Voltafactory *VoltafactorySession) Owner() (common.Address, error) {
	return _Voltafactory.Contract.Owner(&_Voltafactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Voltafactory *VoltafactoryCallerSession) Owner() (common.Address, error) {
	return _Voltafactory.Contract.Owner(&_Voltafactory.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Voltafactory *VoltafactoryCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Voltafactory.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Voltafactory *VoltafactorySession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Voltafactory.Contract.OwnershipHandoverExpiresAt(&_Voltafactory.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Voltafactory *VoltafactoryCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Voltafactory.Contract.OwnershipHandoverExpiresAt(&_Voltafactory.CallOpts, pendingOwner)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Voltafactory *VoltafactoryTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Voltafactory *VoltafactorySession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Voltafactory.Contract.AddStake(&_Voltafactory.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _Voltafactory.Contract.AddStake(&_Voltafactory.TransactOpts, unstakeDelaySec)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactoryTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactorySession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Voltafactory.Contract.CancelOwnershipHandover(&_Voltafactory.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Voltafactory.Contract.CancelOwnershipHandover(&_Voltafactory.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Voltafactory *VoltafactoryTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Voltafactory *VoltafactorySession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.CompleteOwnershipHandover(&_Voltafactory.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.CompleteOwnershipHandover(&_Voltafactory.TransactOpts, pendingOwner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Voltafactory *VoltafactoryTransactor) CreateAccount(opts *bind.TransactOpts, accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "createAccount", accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Voltafactory *VoltafactorySession) CreateAccount(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Voltafactory.Contract.CreateAccount(&_Voltafactory.TransactOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Voltafactory *VoltafactoryTransactorSession) CreateAccount(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Voltafactory.Contract.CreateAccount(&_Voltafactory.TransactOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Voltafactory *VoltafactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Voltafactory *VoltafactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Voltafactory.Contract.RenounceOwnership(&_Voltafactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Voltafactory.Contract.RenounceOwnership(&_Voltafactory.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactoryTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactorySession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Voltafactory.Contract.RequestOwnershipHandover(&_Voltafactory.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Voltafactory.Contract.RequestOwnershipHandover(&_Voltafactory.TransactOpts)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_Voltafactory *VoltafactoryTransactor) SetEntryPoint(opts *bind.TransactOpts, _entryPoint common.Address) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "setEntryPoint", _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_Voltafactory *VoltafactorySession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.SetEntryPoint(&_Voltafactory.TransactOpts, _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_Voltafactory *VoltafactoryTransactorSession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.SetEntryPoint(&_Voltafactory.TransactOpts, _entryPoint)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Voltafactory *VoltafactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Voltafactory *VoltafactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.TransferOwnership(&_Voltafactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Voltafactory *VoltafactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.TransferOwnership(&_Voltafactory.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Voltafactory *VoltafactoryTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Voltafactory *VoltafactorySession) UnlockStake() (*types.Transaction, error) {
	return _Voltafactory.Contract.UnlockStake(&_Voltafactory.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_Voltafactory *VoltafactoryTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _Voltafactory.Contract.UnlockStake(&_Voltafactory.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Voltafactory *VoltafactoryTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _Voltafactory.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Voltafactory *VoltafactorySession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.WithdrawStake(&_Voltafactory.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_Voltafactory *VoltafactoryTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _Voltafactory.Contract.WithdrawStake(&_Voltafactory.TransactOpts, withdrawAddress)
}

// VoltafactoryOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Voltafactory contract.
type VoltafactoryOwnershipHandoverCanceledIterator struct {
	Event *VoltafactoryOwnershipHandoverCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoltafactoryOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltafactoryOwnershipHandoverCanceled)
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
		it.Event = new(VoltafactoryOwnershipHandoverCanceled)
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
func (it *VoltafactoryOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltafactoryOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltafactoryOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Voltafactory contract.
type VoltafactoryOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*VoltafactoryOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryOwnershipHandoverCanceledIterator{contract: _Voltafactory.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *VoltafactoryOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltafactoryOwnershipHandoverCanceled)
				if err := _Voltafactory.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*VoltafactoryOwnershipHandoverCanceled, error) {
	event := new(VoltafactoryOwnershipHandoverCanceled)
	if err := _Voltafactory.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltafactoryOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Voltafactory contract.
type VoltafactoryOwnershipHandoverRequestedIterator struct {
	Event *VoltafactoryOwnershipHandoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoltafactoryOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltafactoryOwnershipHandoverRequested)
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
		it.Event = new(VoltafactoryOwnershipHandoverRequested)
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
func (it *VoltafactoryOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltafactoryOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltafactoryOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Voltafactory contract.
type VoltafactoryOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*VoltafactoryOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryOwnershipHandoverRequestedIterator{contract: _Voltafactory.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *VoltafactoryOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltafactoryOwnershipHandoverRequested)
				if err := _Voltafactory.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Voltafactory *VoltafactoryFilterer) ParseOwnershipHandoverRequested(log types.Log) (*VoltafactoryOwnershipHandoverRequested, error) {
	event := new(VoltafactoryOwnershipHandoverRequested)
	if err := _Voltafactory.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltafactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Voltafactory contract.
type VoltafactoryOwnershipTransferredIterator struct {
	Event *VoltafactoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VoltafactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltafactoryOwnershipTransferred)
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
		it.Event = new(VoltafactoryOwnershipTransferred)
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
func (it *VoltafactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltafactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltafactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Voltafactory contract.
type VoltafactoryOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Voltafactory *VoltafactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*VoltafactoryOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoltafactoryOwnershipTransferredIterator{contract: _Voltafactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Voltafactory *VoltafactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VoltafactoryOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Voltafactory.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltafactoryOwnershipTransferred)
				if err := _Voltafactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Voltafactory *VoltafactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VoltafactoryOwnershipTransferred, error) {
	event := new(VoltafactoryOwnershipTransferred)
	if err := _Voltafactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
