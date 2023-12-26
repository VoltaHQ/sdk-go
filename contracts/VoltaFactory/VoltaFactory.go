// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKeyValidatorImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executorImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minNumSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKeyValidatorImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executorImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"minNumSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractVoltaAccount\",\"name\":\"ret\",\"type\":\"address\"}]}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Factory *FactoryCaller) GetAddress(opts *bind.CallOpts, accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getAddress", accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Factory *FactorySession) GetAddress(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	return _Factory.Contract.GetAddress(&_Factory.CallOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8b2f057a.
//
// Solidity: function getAddress(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) view returns(address)
func (_Factory *FactoryCallerSession) GetAddress(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (common.Address, error) {
	return _Factory.Contract.GetAddress(&_Factory.CallOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Factory *FactoryTransactor) CreateAccount(opts *bind.TransactOpts, accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createAccount", accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Factory *FactorySession) CreateAccount(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreateAccount(&_Factory.TransactOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x2181f423.
//
// Solidity: function createAccount(address accountImplementation, address sessionKeyValidatorImplementation, address executorImplementation, address[] owners, uint256 minNumSigners, uint256 salt) returns(address ret)
func (_Factory *FactoryTransactorSession) CreateAccount(accountImplementation common.Address, sessionKeyValidatorImplementation common.Address, executorImplementation common.Address, owners []common.Address, minNumSigners *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreateAccount(&_Factory.TransactOpts, accountImplementation, sessionKeyValidatorImplementation, executorImplementation, owners, minNumSigners, salt)
}
