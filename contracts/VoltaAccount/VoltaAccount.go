// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

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

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
}

// VoltaAccountMetaData contains all meta data concerning the VoltaAccount contract.
var VoltaAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoltaAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use VoltaAccountMetaData.ABI instead.
var VoltaAccountABI = VoltaAccountMetaData.ABI

// VoltaAccount is an auto generated Go binding around an Ethereum contract.
type VoltaAccount struct {
	VoltaAccountCaller     // Read-only binding to the contract
	VoltaAccountTransactor // Write-only binding to the contract
	VoltaAccountFilterer   // Log filterer for contract events
}

// VoltaAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoltaAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoltaAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoltaAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoltaAccountSession struct {
	Contract     *VoltaAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoltaAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoltaAccountCallerSession struct {
	Contract *VoltaAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoltaAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoltaAccountTransactorSession struct {
	Contract     *VoltaAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoltaAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoltaAccountRaw struct {
	Contract *VoltaAccount // Generic contract binding to access the raw methods on
}

// VoltaAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoltaAccountCallerRaw struct {
	Contract *VoltaAccountCaller // Generic read-only contract binding to access the raw methods on
}

// VoltaAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoltaAccountTransactorRaw struct {
	Contract *VoltaAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoltaAccount creates a new instance of VoltaAccount, bound to a specific deployed contract.
func NewVoltaAccount(address common.Address, backend bind.ContractBackend) (*VoltaAccount, error) {
	contract, err := bindVoltaAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoltaAccount{VoltaAccountCaller: VoltaAccountCaller{contract: contract}, VoltaAccountTransactor: VoltaAccountTransactor{contract: contract}, VoltaAccountFilterer: VoltaAccountFilterer{contract: contract}}, nil
}

// NewVoltaAccountCaller creates a new read-only instance of VoltaAccount, bound to a specific deployed contract.
func NewVoltaAccountCaller(address common.Address, caller bind.ContractCaller) (*VoltaAccountCaller, error) {
	contract, err := bindVoltaAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoltaAccountCaller{contract: contract}, nil
}

// NewVoltaAccountTransactor creates a new write-only instance of VoltaAccount, bound to a specific deployed contract.
func NewVoltaAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*VoltaAccountTransactor, error) {
	contract, err := bindVoltaAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoltaAccountTransactor{contract: contract}, nil
}

// NewVoltaAccountFilterer creates a new log filterer instance of VoltaAccount, bound to a specific deployed contract.
func NewVoltaAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*VoltaAccountFilterer, error) {
	contract, err := bindVoltaAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoltaAccountFilterer{contract: contract}, nil
}

// bindVoltaAccount binds a generic wrapper to an already deployed contract.
func bindVoltaAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoltaAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoltaAccount *VoltaAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoltaAccount.Contract.VoltaAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoltaAccount *VoltaAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoltaAccount.Contract.VoltaAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoltaAccount *VoltaAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoltaAccount.Contract.VoltaAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoltaAccount *VoltaAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoltaAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoltaAccount *VoltaAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoltaAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoltaAccount *VoltaAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoltaAccount.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes data) returns()
func (_VoltaAccount *VoltaAccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "execute", dest, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes data) returns()
func (_VoltaAccount *VoltaAccountSession) Execute(dest common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Execute(&_VoltaAccount.TransactOpts, dest, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes data) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) Execute(dest common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Execute(&_VoltaAccount.TransactOpts, dest, value, data)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_VoltaAccount *VoltaAccountTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []Call) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_VoltaAccount *VoltaAccountSession) ExecuteBatch(calls []Call) (*types.Transaction, error) {
	return _VoltaAccount.Contract.ExecuteBatch(&_VoltaAccount.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) ExecuteBatch(calls []Call) (*types.Transaction, error) {
	return _VoltaAccount.Contract.ExecuteBatch(&_VoltaAccount.TransactOpts, calls)
}
