// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package volta

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

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// VoltaMetaData contains all meta data concerning the Volta contract.
var VoltaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"DisabledAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"EnabledAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"owner\",\"type\":\"address[]\"}],\"name\":\"VoltaAccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"test\",\"type\":\"string\"}],\"name\":\"VoltaAccountLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"test\",\"type\":\"uint256\"}],\"name\":\"VoltaAccountLogNum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_minQuorum\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessionKeyContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"sessionKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint32\",\"name\":\"addressOffset\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VoltaABI is the input ABI used to generate the binding from.
// Deprecated: Use VoltaMetaData.ABI instead.
var VoltaABI = VoltaMetaData.ABI

// Volta is an auto generated Go binding around an Ethereum contract.
type Volta struct {
	VoltaCaller     // Read-only binding to the contract
	VoltaTransactor // Write-only binding to the contract
	VoltaFilterer   // Log filterer for contract events
}

// VoltaCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoltaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoltaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoltaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoltaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoltaSession struct {
	Contract     *Volta            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoltaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoltaCallerSession struct {
	Contract *VoltaCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoltaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoltaTransactorSession struct {
	Contract     *VoltaTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoltaRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoltaRaw struct {
	Contract *Volta // Generic contract binding to access the raw methods on
}

// VoltaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoltaCallerRaw struct {
	Contract *VoltaCaller // Generic read-only contract binding to access the raw methods on
}

// VoltaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoltaTransactorRaw struct {
	Contract *VoltaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolta creates a new instance of Volta, bound to a specific deployed contract.
func NewVolta(address common.Address, backend bind.ContractBackend) (*Volta, error) {
	contract, err := bindVolta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Volta{VoltaCaller: VoltaCaller{contract: contract}, VoltaTransactor: VoltaTransactor{contract: contract}, VoltaFilterer: VoltaFilterer{contract: contract}}, nil
}

// NewVoltaCaller creates a new read-only instance of Volta, bound to a specific deployed contract.
func NewVoltaCaller(address common.Address, caller bind.ContractCaller) (*VoltaCaller, error) {
	contract, err := bindVolta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoltaCaller{contract: contract}, nil
}

// NewVoltaTransactor creates a new write-only instance of Volta, bound to a specific deployed contract.
func NewVoltaTransactor(address common.Address, transactor bind.ContractTransactor) (*VoltaTransactor, error) {
	contract, err := bindVolta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoltaTransactor{contract: contract}, nil
}

// NewVoltaFilterer creates a new log filterer instance of Volta, bound to a specific deployed contract.
func NewVoltaFilterer(address common.Address, filterer bind.ContractFilterer) (*VoltaFilterer, error) {
	contract, err := bindVolta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoltaFilterer{contract: contract}, nil
}

// bindVolta binds a generic wrapper to an already deployed contract.
func bindVolta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoltaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Volta *VoltaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Volta.Contract.VoltaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Volta *VoltaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volta.Contract.VoltaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Volta *VoltaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Volta.Contract.VoltaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Volta *VoltaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Volta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Volta *VoltaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Volta *VoltaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Volta.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Volta *VoltaCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Volta *VoltaSession) EntryPoint() (common.Address, error) {
	return _Volta.Contract.EntryPoint(&_Volta.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Volta *VoltaCallerSession) EntryPoint() (common.Address, error) {
	return _Volta.Contract.EntryPoint(&_Volta.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Volta *VoltaCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Volta *VoltaSession) GetDeposit() (*big.Int, error) {
	return _Volta.Contract.GetDeposit(&_Volta.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Volta *VoltaCallerSession) GetDeposit() (*big.Int, error) {
	return _Volta.Contract.GetDeposit(&_Volta.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Volta *VoltaCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Volta *VoltaSession) GetNonce() (*big.Int, error) {
	return _Volta.Contract.GetNonce(&_Volta.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Volta *VoltaCallerSession) GetNonce() (*big.Int, error) {
	return _Volta.Contract.GetNonce(&_Volta.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Volta *VoltaCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Volta *VoltaSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC1155BatchReceived(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Volta *VoltaCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC1155BatchReceived(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC1155Received(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC1155Received(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC721Received(&_Volta.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Volta *VoltaCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Volta.Contract.OnERC721Received(&_Volta.CallOpts, arg0, arg1, arg2, arg3)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Volta *VoltaCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Volta *VoltaSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Volta.Contract.Owners(&_Volta.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Volta *VoltaCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Volta.Contract.Owners(&_Volta.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Volta *VoltaCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Volta *VoltaSession) ProxiableUUID() ([32]byte, error) {
	return _Volta.Contract.ProxiableUUID(&_Volta.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Volta *VoltaCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Volta.Contract.ProxiableUUID(&_Volta.CallOpts)
}

// SessionKeyContracts is a free data retrieval call binding the contract method 0x47c6e063.
//
// Solidity: function sessionKeyContracts(address , uint256 ) view returns(address)
func (_Volta *VoltaCaller) SessionKeyContracts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "sessionKeyContracts", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SessionKeyContracts is a free data retrieval call binding the contract method 0x47c6e063.
//
// Solidity: function sessionKeyContracts(address , uint256 ) view returns(address)
func (_Volta *VoltaSession) SessionKeyContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Volta.Contract.SessionKeyContracts(&_Volta.CallOpts, arg0, arg1)
}

// SessionKeyContracts is a free data retrieval call binding the contract method 0x47c6e063.
//
// Solidity: function sessionKeyContracts(address , uint256 ) view returns(address)
func (_Volta *VoltaCallerSession) SessionKeyContracts(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Volta.Contract.SessionKeyContracts(&_Volta.CallOpts, arg0, arg1)
}

// SessionKeys is a free data retrieval call binding the contract method 0x96ade1f9.
//
// Solidity: function sessionKeys(address sessionKey, address contractAddress) view returns(bool enabled, bytes4 selector, uint32 addressOffset)
func (_Volta *VoltaCaller) SessionKeys(opts *bind.CallOpts, sessionKey common.Address, contractAddress common.Address) (struct {
	Enabled       bool
	Selector      [4]byte
	AddressOffset uint32
}, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "sessionKeys", sessionKey, contractAddress)

	outstruct := new(struct {
		Enabled       bool
		Selector      [4]byte
		AddressOffset uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Selector = *abi.ConvertType(out[1], new([4]byte)).(*[4]byte)
	outstruct.AddressOffset = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// SessionKeys is a free data retrieval call binding the contract method 0x96ade1f9.
//
// Solidity: function sessionKeys(address sessionKey, address contractAddress) view returns(bool enabled, bytes4 selector, uint32 addressOffset)
func (_Volta *VoltaSession) SessionKeys(sessionKey common.Address, contractAddress common.Address) (struct {
	Enabled       bool
	Selector      [4]byte
	AddressOffset uint32
}, error) {
	return _Volta.Contract.SessionKeys(&_Volta.CallOpts, sessionKey, contractAddress)
}

// SessionKeys is a free data retrieval call binding the contract method 0x96ade1f9.
//
// Solidity: function sessionKeys(address sessionKey, address contractAddress) view returns(bool enabled, bytes4 selector, uint32 addressOffset)
func (_Volta *VoltaCallerSession) SessionKeys(sessionKey common.Address, contractAddress common.Address) (struct {
	Enabled       bool
	Selector      [4]byte
	AddressOffset uint32
}, error) {
	return _Volta.Contract.SessionKeys(&_Volta.CallOpts, sessionKey, contractAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Volta *VoltaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Volta *VoltaSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Volta.Contract.SupportsInterface(&_Volta.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Volta *VoltaCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Volta.Contract.SupportsInterface(&_Volta.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Volta *VoltaCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Volta.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Volta *VoltaSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Volta.Contract.TokensReceived(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Volta *VoltaCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Volta.Contract.TokensReceived(&_Volta.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Volta *VoltaTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Volta *VoltaSession) AddDeposit() (*types.Transaction, error) {
	return _Volta.Contract.AddDeposit(&_Volta.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Volta *VoltaTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Volta.Contract.AddDeposit(&_Volta.TransactOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_Volta *VoltaTransactor) Disable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "disable", _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_Volta *VoltaSession) Disable(_data []byte) (*types.Transaction, error) {
	return _Volta.Contract.Disable(&_Volta.TransactOpts, _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_Volta *VoltaTransactorSession) Disable(_data []byte) (*types.Transaction, error) {
	return _Volta.Contract.Disable(&_Volta.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_Volta *VoltaTransactor) Enable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "enable", _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_Volta *VoltaSession) Enable(_data []byte) (*types.Transaction, error) {
	return _Volta.Contract.Enable(&_Volta.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_Volta *VoltaTransactorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _Volta.Contract.Enable(&_Volta.TransactOpts, _data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Volta *VoltaTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Volta *VoltaSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Volta.Contract.Execute(&_Volta.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Volta *VoltaTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Volta.Contract.Execute(&_Volta.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Volta *VoltaTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "executeBatch", dest, values, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Volta *VoltaSession) ExecuteBatch(dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Volta.Contract.ExecuteBatch(&_Volta.TransactOpts, dest, values, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Volta *VoltaTransactorSession) ExecuteBatch(dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Volta.Contract.ExecuteBatch(&_Volta.TransactOpts, dest, values, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0x60b5bb3f.
//
// Solidity: function initialize(address[] _owners, uint256 _minQuorum) returns()
func (_Volta *VoltaTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address, _minQuorum *big.Int) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "initialize", _owners, _minQuorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x60b5bb3f.
//
// Solidity: function initialize(address[] _owners, uint256 _minQuorum) returns()
func (_Volta *VoltaSession) Initialize(_owners []common.Address, _minQuorum *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.Initialize(&_Volta.TransactOpts, _owners, _minQuorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x60b5bb3f.
//
// Solidity: function initialize(address[] _owners, uint256 _minQuorum) returns()
func (_Volta *VoltaTransactorSession) Initialize(_owners []common.Address, _minQuorum *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.Initialize(&_Volta.TransactOpts, _owners, _minQuorum)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Volta *VoltaTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Volta *VoltaSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Volta.Contract.UpgradeTo(&_Volta.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Volta *VoltaTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Volta.Contract.UpgradeTo(&_Volta.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Volta *VoltaTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Volta *VoltaSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Volta.Contract.UpgradeToAndCall(&_Volta.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Volta *VoltaTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Volta.Contract.UpgradeToAndCall(&_Volta.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Volta *VoltaTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Volta *VoltaSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.ValidateUserOp(&_Volta.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Volta *VoltaTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.ValidateUserOp(&_Volta.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Volta *VoltaTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Volta.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Volta *VoltaSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.WithdrawDepositTo(&_Volta.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Volta *VoltaTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Volta.Contract.WithdrawDepositTo(&_Volta.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Volta *VoltaTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volta.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Volta *VoltaSession) Receive() (*types.Transaction, error) {
	return _Volta.Contract.Receive(&_Volta.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Volta *VoltaTransactorSession) Receive() (*types.Transaction, error) {
	return _Volta.Contract.Receive(&_Volta.TransactOpts)
}

// VoltaAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Volta contract.
type VoltaAdminChangedIterator struct {
	Event *VoltaAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoltaAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaAdminChanged)
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
		it.Event = new(VoltaAdminChanged)
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
func (it *VoltaAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaAdminChanged represents a AdminChanged event raised by the Volta contract.
type VoltaAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Volta *VoltaFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VoltaAdminChangedIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VoltaAdminChangedIterator{contract: _Volta.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Volta *VoltaFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VoltaAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaAdminChanged)
				if err := _Volta.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Volta *VoltaFilterer) ParseAdminChanged(log types.Log) (*VoltaAdminChanged, error) {
	event := new(VoltaAdminChanged)
	if err := _Volta.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Volta contract.
type VoltaBeaconUpgradedIterator struct {
	Event *VoltaBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VoltaBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaBeaconUpgraded)
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
		it.Event = new(VoltaBeaconUpgraded)
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
func (it *VoltaBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaBeaconUpgraded represents a BeaconUpgraded event raised by the Volta contract.
type VoltaBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Volta *VoltaFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VoltaBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Volta.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VoltaBeaconUpgradedIterator{contract: _Volta.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Volta *VoltaFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VoltaBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Volta.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaBeaconUpgraded)
				if err := _Volta.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Volta *VoltaFilterer) ParseBeaconUpgraded(log types.Log) (*VoltaBeaconUpgraded, error) {
	event := new(VoltaBeaconUpgraded)
	if err := _Volta.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaDisabledAccountIterator is returned from FilterDisabledAccount and is used to iterate over the raw logs and unpacked data for DisabledAccount events raised by the Volta contract.
type VoltaDisabledAccountIterator struct {
	Event *VoltaDisabledAccount // Event containing the contract specifics and raw log

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
func (it *VoltaDisabledAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaDisabledAccount)
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
		it.Event = new(VoltaDisabledAccount)
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
func (it *VoltaDisabledAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaDisabledAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaDisabledAccount represents a DisabledAccount event raised by the Volta contract.
type VoltaDisabledAccount struct {
	User            common.Address
	ContractAddress common.Address
	Selector        [4]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDisabledAccount is a free log retrieval operation binding the contract event 0x43bc2e4b4ba88224d5b9e5ff1183a4f1d2d0a9b7762c9a46cb5920ac2b9eff22.
//
// Solidity: event DisabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) FilterDisabledAccount(opts *bind.FilterOpts) (*VoltaDisabledAccountIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "DisabledAccount")
	if err != nil {
		return nil, err
	}
	return &VoltaDisabledAccountIterator{contract: _Volta.contract, event: "DisabledAccount", logs: logs, sub: sub}, nil
}

// WatchDisabledAccount is a free log subscription operation binding the contract event 0x43bc2e4b4ba88224d5b9e5ff1183a4f1d2d0a9b7762c9a46cb5920ac2b9eff22.
//
// Solidity: event DisabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) WatchDisabledAccount(opts *bind.WatchOpts, sink chan<- *VoltaDisabledAccount) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "DisabledAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaDisabledAccount)
				if err := _Volta.contract.UnpackLog(event, "DisabledAccount", log); err != nil {
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

// ParseDisabledAccount is a log parse operation binding the contract event 0x43bc2e4b4ba88224d5b9e5ff1183a4f1d2d0a9b7762c9a46cb5920ac2b9eff22.
//
// Solidity: event DisabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) ParseDisabledAccount(log types.Log) (*VoltaDisabledAccount, error) {
	event := new(VoltaDisabledAccount)
	if err := _Volta.contract.UnpackLog(event, "DisabledAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaEnabledAccountIterator is returned from FilterEnabledAccount and is used to iterate over the raw logs and unpacked data for EnabledAccount events raised by the Volta contract.
type VoltaEnabledAccountIterator struct {
	Event *VoltaEnabledAccount // Event containing the contract specifics and raw log

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
func (it *VoltaEnabledAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaEnabledAccount)
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
		it.Event = new(VoltaEnabledAccount)
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
func (it *VoltaEnabledAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaEnabledAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaEnabledAccount represents a EnabledAccount event raised by the Volta contract.
type VoltaEnabledAccount struct {
	User            common.Address
	ContractAddress common.Address
	Selector        [4]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnabledAccount is a free log retrieval operation binding the contract event 0x1c048ce30809390ea9a9a3805b85146032ada9a0cbe3a9eeffa81e5651de41c7.
//
// Solidity: event EnabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) FilterEnabledAccount(opts *bind.FilterOpts) (*VoltaEnabledAccountIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "EnabledAccount")
	if err != nil {
		return nil, err
	}
	return &VoltaEnabledAccountIterator{contract: _Volta.contract, event: "EnabledAccount", logs: logs, sub: sub}, nil
}

// WatchEnabledAccount is a free log subscription operation binding the contract event 0x1c048ce30809390ea9a9a3805b85146032ada9a0cbe3a9eeffa81e5651de41c7.
//
// Solidity: event EnabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) WatchEnabledAccount(opts *bind.WatchOpts, sink chan<- *VoltaEnabledAccount) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "EnabledAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaEnabledAccount)
				if err := _Volta.contract.UnpackLog(event, "EnabledAccount", log); err != nil {
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

// ParseEnabledAccount is a log parse operation binding the contract event 0x1c048ce30809390ea9a9a3805b85146032ada9a0cbe3a9eeffa81e5651de41c7.
//
// Solidity: event EnabledAccount(address user, address contractAddress, bytes4 selector)
func (_Volta *VoltaFilterer) ParseEnabledAccount(log types.Log) (*VoltaEnabledAccount, error) {
	event := new(VoltaEnabledAccount)
	if err := _Volta.contract.UnpackLog(event, "EnabledAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Volta contract.
type VoltaInitializedIterator struct {
	Event *VoltaInitialized // Event containing the contract specifics and raw log

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
func (it *VoltaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaInitialized)
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
		it.Event = new(VoltaInitialized)
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
func (it *VoltaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaInitialized represents a Initialized event raised by the Volta contract.
type VoltaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Volta *VoltaFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoltaInitializedIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoltaInitializedIterator{contract: _Volta.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Volta *VoltaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoltaInitialized) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaInitialized)
				if err := _Volta.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Volta *VoltaFilterer) ParseInitialized(log types.Log) (*VoltaInitialized, error) {
	event := new(VoltaInitialized)
	if err := _Volta.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Volta contract.
type VoltaUpgradedIterator struct {
	Event *VoltaUpgraded // Event containing the contract specifics and raw log

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
func (it *VoltaUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaUpgraded)
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
		it.Event = new(VoltaUpgraded)
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
func (it *VoltaUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaUpgraded represents a Upgraded event raised by the Volta contract.
type VoltaUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Volta *VoltaFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VoltaUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Volta.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VoltaUpgradedIterator{contract: _Volta.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Volta *VoltaFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VoltaUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Volta.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaUpgraded)
				if err := _Volta.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Volta *VoltaFilterer) ParseUpgraded(log types.Log) (*VoltaUpgraded, error) {
	event := new(VoltaUpgraded)
	if err := _Volta.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaVoltaAccountInitializedIterator is returned from FilterVoltaAccountInitialized and is used to iterate over the raw logs and unpacked data for VoltaAccountInitialized events raised by the Volta contract.
type VoltaVoltaAccountInitializedIterator struct {
	Event *VoltaVoltaAccountInitialized // Event containing the contract specifics and raw log

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
func (it *VoltaVoltaAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaVoltaAccountInitialized)
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
		it.Event = new(VoltaVoltaAccountInitialized)
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
func (it *VoltaVoltaAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaVoltaAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaVoltaAccountInitialized represents a VoltaAccountInitialized event raised by the Volta contract.
type VoltaVoltaAccountInitialized struct {
	EntryPoint common.Address
	Owner      []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoltaAccountInitialized is a free log retrieval operation binding the contract event 0xa63508622b5e1b8a9cff6b677fca72435b171a2f29a219fc90bbad3a52c5e049.
//
// Solidity: event VoltaAccountInitialized(address indexed entryPoint, address[] indexed owner)
func (_Volta *VoltaFilterer) FilterVoltaAccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address, owner [][]common.Address) (*VoltaVoltaAccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Volta.contract.FilterLogs(opts, "VoltaAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VoltaVoltaAccountInitializedIterator{contract: _Volta.contract, event: "VoltaAccountInitialized", logs: logs, sub: sub}, nil
}

// WatchVoltaAccountInitialized is a free log subscription operation binding the contract event 0xa63508622b5e1b8a9cff6b677fca72435b171a2f29a219fc90bbad3a52c5e049.
//
// Solidity: event VoltaAccountInitialized(address indexed entryPoint, address[] indexed owner)
func (_Volta *VoltaFilterer) WatchVoltaAccountInitialized(opts *bind.WatchOpts, sink chan<- *VoltaVoltaAccountInitialized, entryPoint []common.Address, owner [][]common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Volta.contract.WatchLogs(opts, "VoltaAccountInitialized", entryPointRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaVoltaAccountInitialized)
				if err := _Volta.contract.UnpackLog(event, "VoltaAccountInitialized", log); err != nil {
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

// ParseVoltaAccountInitialized is a log parse operation binding the contract event 0xa63508622b5e1b8a9cff6b677fca72435b171a2f29a219fc90bbad3a52c5e049.
//
// Solidity: event VoltaAccountInitialized(address indexed entryPoint, address[] indexed owner)
func (_Volta *VoltaFilterer) ParseVoltaAccountInitialized(log types.Log) (*VoltaVoltaAccountInitialized, error) {
	event := new(VoltaVoltaAccountInitialized)
	if err := _Volta.contract.UnpackLog(event, "VoltaAccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaVoltaAccountLogIterator is returned from FilterVoltaAccountLog and is used to iterate over the raw logs and unpacked data for VoltaAccountLog events raised by the Volta contract.
type VoltaVoltaAccountLogIterator struct {
	Event *VoltaVoltaAccountLog // Event containing the contract specifics and raw log

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
func (it *VoltaVoltaAccountLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaVoltaAccountLog)
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
		it.Event = new(VoltaVoltaAccountLog)
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
func (it *VoltaVoltaAccountLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaVoltaAccountLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaVoltaAccountLog represents a VoltaAccountLog event raised by the Volta contract.
type VoltaVoltaAccountLog struct {
	Test string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVoltaAccountLog is a free log retrieval operation binding the contract event 0xb72bbc6e6149f87a9153f9b2dbcb8ce07b4ca9f31c50b1941be2072227ee1a7e.
//
// Solidity: event VoltaAccountLog(string test)
func (_Volta *VoltaFilterer) FilterVoltaAccountLog(opts *bind.FilterOpts) (*VoltaVoltaAccountLogIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "VoltaAccountLog")
	if err != nil {
		return nil, err
	}
	return &VoltaVoltaAccountLogIterator{contract: _Volta.contract, event: "VoltaAccountLog", logs: logs, sub: sub}, nil
}

// WatchVoltaAccountLog is a free log subscription operation binding the contract event 0xb72bbc6e6149f87a9153f9b2dbcb8ce07b4ca9f31c50b1941be2072227ee1a7e.
//
// Solidity: event VoltaAccountLog(string test)
func (_Volta *VoltaFilterer) WatchVoltaAccountLog(opts *bind.WatchOpts, sink chan<- *VoltaVoltaAccountLog) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "VoltaAccountLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaVoltaAccountLog)
				if err := _Volta.contract.UnpackLog(event, "VoltaAccountLog", log); err != nil {
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

// ParseVoltaAccountLog is a log parse operation binding the contract event 0xb72bbc6e6149f87a9153f9b2dbcb8ce07b4ca9f31c50b1941be2072227ee1a7e.
//
// Solidity: event VoltaAccountLog(string test)
func (_Volta *VoltaFilterer) ParseVoltaAccountLog(log types.Log) (*VoltaVoltaAccountLog, error) {
	event := new(VoltaVoltaAccountLog)
	if err := _Volta.contract.UnpackLog(event, "VoltaAccountLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaVoltaAccountLogNumIterator is returned from FilterVoltaAccountLogNum and is used to iterate over the raw logs and unpacked data for VoltaAccountLogNum events raised by the Volta contract.
type VoltaVoltaAccountLogNumIterator struct {
	Event *VoltaVoltaAccountLogNum // Event containing the contract specifics and raw log

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
func (it *VoltaVoltaAccountLogNumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaVoltaAccountLogNum)
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
		it.Event = new(VoltaVoltaAccountLogNum)
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
func (it *VoltaVoltaAccountLogNumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaVoltaAccountLogNumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaVoltaAccountLogNum represents a VoltaAccountLogNum event raised by the Volta contract.
type VoltaVoltaAccountLogNum struct {
	Test *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVoltaAccountLogNum is a free log retrieval operation binding the contract event 0xb0f064b5932cd9f0065447ab73a5fea198c3a50aae5567e1d55b8ad72cc0b755.
//
// Solidity: event VoltaAccountLogNum(uint256 test)
func (_Volta *VoltaFilterer) FilterVoltaAccountLogNum(opts *bind.FilterOpts) (*VoltaVoltaAccountLogNumIterator, error) {

	logs, sub, err := _Volta.contract.FilterLogs(opts, "VoltaAccountLogNum")
	if err != nil {
		return nil, err
	}
	return &VoltaVoltaAccountLogNumIterator{contract: _Volta.contract, event: "VoltaAccountLogNum", logs: logs, sub: sub}, nil
}

// WatchVoltaAccountLogNum is a free log subscription operation binding the contract event 0xb0f064b5932cd9f0065447ab73a5fea198c3a50aae5567e1d55b8ad72cc0b755.
//
// Solidity: event VoltaAccountLogNum(uint256 test)
func (_Volta *VoltaFilterer) WatchVoltaAccountLogNum(opts *bind.WatchOpts, sink chan<- *VoltaVoltaAccountLogNum) (event.Subscription, error) {

	logs, sub, err := _Volta.contract.WatchLogs(opts, "VoltaAccountLogNum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaVoltaAccountLogNum)
				if err := _Volta.contract.UnpackLog(event, "VoltaAccountLogNum", log); err != nil {
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

// ParseVoltaAccountLogNum is a log parse operation binding the contract event 0xb0f064b5932cd9f0065447ab73a5fea198c3a50aae5567e1d55b8ad72cc0b755.
//
// Solidity: event VoltaAccountLogNum(uint256 test)
func (_Volta *VoltaFilterer) ParseVoltaAccountLogNum(log types.Log) (*VoltaVoltaAccountLogNum, error) {
	event := new(VoltaVoltaAccountLogNum)
	if err := _Volta.contract.UnpackLog(event, "VoltaAccountLogNum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
