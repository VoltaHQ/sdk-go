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

// VoltaAccountMetaData contains all meta data concerning the VoltaAccount contract.
var VoltaAccountMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sessionKeyValidatorImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executorImplementation\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_minQuorum\",\"type\":\"uint256\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"__entryPoint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newExecutor\",\"type\":\"address\"}],\"name\":\"upgradeExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSessionKeyValidator\",\"type\":\"address\"}],\"name\":\"upgradeSessionKeyValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}],",
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

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VoltaAccount *VoltaAccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VoltaAccount *VoltaAccountSession) EntryPoint() (common.Address, error) {
	return _VoltaAccount.Contract.EntryPoint(&_VoltaAccount.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VoltaAccount *VoltaAccountCallerSession) EntryPoint() (common.Address, error) {
	return _VoltaAccount.Contract.EntryPoint(&_VoltaAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VoltaAccount *VoltaAccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VoltaAccount *VoltaAccountSession) GetDeposit() (*big.Int, error) {
	return _VoltaAccount.Contract.GetDeposit(&_VoltaAccount.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VoltaAccount *VoltaAccountCallerSession) GetDeposit() (*big.Int, error) {
	return _VoltaAccount.Contract.GetDeposit(&_VoltaAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_VoltaAccount *VoltaAccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_VoltaAccount *VoltaAccountSession) GetNonce() (*big.Int, error) {
	return _VoltaAccount.Contract.GetNonce(&_VoltaAccount.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_VoltaAccount *VoltaAccountCallerSession) GetNonce() (*big.Int, error) {
	return _VoltaAccount.Contract.GetNonce(&_VoltaAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC1155BatchReceived(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC1155BatchReceived(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC1155Received(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC1155Received(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC721Received(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_VoltaAccount *VoltaAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _VoltaAccount.Contract.OnERC721Received(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_VoltaAccount *VoltaAccountCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_VoltaAccount *VoltaAccountSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _VoltaAccount.Contract.Owners(&_VoltaAccount.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_VoltaAccount *VoltaAccountCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _VoltaAccount.Contract.Owners(&_VoltaAccount.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VoltaAccount *VoltaAccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VoltaAccount *VoltaAccountSession) ProxiableUUID() ([32]byte, error) {
	return _VoltaAccount.Contract.ProxiableUUID(&_VoltaAccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VoltaAccount *VoltaAccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VoltaAccount.Contract.ProxiableUUID(&_VoltaAccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoltaAccount *VoltaAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoltaAccount *VoltaAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoltaAccount.Contract.SupportsInterface(&_VoltaAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoltaAccount *VoltaAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoltaAccount.Contract.SupportsInterface(&_VoltaAccount.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_VoltaAccount *VoltaAccountCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _VoltaAccount.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_VoltaAccount *VoltaAccountSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _VoltaAccount.Contract.TokensReceived(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_VoltaAccount *VoltaAccountCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _VoltaAccount.Contract.TokensReceived(&_VoltaAccount.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_VoltaAccount *VoltaAccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_VoltaAccount *VoltaAccountSession) AddDeposit() (*types.Transaction, error) {
	return _VoltaAccount.Contract.AddDeposit(&_VoltaAccount.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_VoltaAccount *VoltaAccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _VoltaAccount.Contract.AddDeposit(&_VoltaAccount.TransactOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountTransactor) Disable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "disable", _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountSession) Disable(_data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Disable(&_VoltaAccount.TransactOpts, _data)
}

// Disable is a paid mutator transaction binding the contract method 0x8fc925aa.
//
// Solidity: function disable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) Disable(_data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Disable(&_VoltaAccount.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountTransactor) Enable(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "enable", _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountSession) Enable(_data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Enable(&_VoltaAccount.TransactOpts, _data)
}

// Enable is a paid mutator transaction binding the contract method 0x0c959556.
//
// Solidity: function enable(bytes _data) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) Enable(_data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Enable(&_VoltaAccount.TransactOpts, _data)
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

// Initialize is a paid mutator transaction binding the contract method 0x264a83b8.
//
// Solidity: function initialize(address _sessionKeyValidatorImplementation, address _executorImplementation, address[] _owners, uint256 _minQuorum, address __entryPoint) returns()
func (_VoltaAccount *VoltaAccountTransactor) Initialize(opts *bind.TransactOpts, _sessionKeyValidatorImplementation common.Address, _executorImplementation common.Address, _owners []common.Address, _minQuorum *big.Int, __entryPoint common.Address) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "initialize", _sessionKeyValidatorImplementation, _executorImplementation, _owners, _minQuorum, __entryPoint)
}

// Initialize is a paid mutator transaction binding the contract method 0x264a83b8.
//
// Solidity: function initialize(address _sessionKeyValidatorImplementation, address _executorImplementation, address[] _owners, uint256 _minQuorum, address __entryPoint) returns()
func (_VoltaAccount *VoltaAccountSession) Initialize(_sessionKeyValidatorImplementation common.Address, _executorImplementation common.Address, _owners []common.Address, _minQuorum *big.Int, __entryPoint common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Initialize(&_VoltaAccount.TransactOpts, _sessionKeyValidatorImplementation, _executorImplementation, _owners, _minQuorum, __entryPoint)
}

// Initialize is a paid mutator transaction binding the contract method 0x264a83b8.
//
// Solidity: function initialize(address _sessionKeyValidatorImplementation, address _executorImplementation, address[] _owners, uint256 _minQuorum, address __entryPoint) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) Initialize(_sessionKeyValidatorImplementation common.Address, _executorImplementation common.Address, _owners []common.Address, _minQuorum *big.Int, __entryPoint common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.Initialize(&_VoltaAccount.TransactOpts, _sessionKeyValidatorImplementation, _executorImplementation, _owners, _minQuorum, __entryPoint)
}

// UpgradeExecutor is a paid mutator transaction binding the contract method 0x6563b5a6.
//
// Solidity: function upgradeExecutor(address newExecutor) returns()
func (_VoltaAccount *VoltaAccountTransactor) UpgradeExecutor(opts *bind.TransactOpts, newExecutor common.Address) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "upgradeExecutor", newExecutor)
}

// UpgradeExecutor is a paid mutator transaction binding the contract method 0x6563b5a6.
//
// Solidity: function upgradeExecutor(address newExecutor) returns()
func (_VoltaAccount *VoltaAccountSession) UpgradeExecutor(newExecutor common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeExecutor(&_VoltaAccount.TransactOpts, newExecutor)
}

// UpgradeExecutor is a paid mutator transaction binding the contract method 0x6563b5a6.
//
// Solidity: function upgradeExecutor(address newExecutor) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) UpgradeExecutor(newExecutor common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeExecutor(&_VoltaAccount.TransactOpts, newExecutor)
}

// UpgradeSessionKeyValidator is a paid mutator transaction binding the contract method 0x4bd191a4.
//
// Solidity: function upgradeSessionKeyValidator(address newSessionKeyValidator) returns()
func (_VoltaAccount *VoltaAccountTransactor) UpgradeSessionKeyValidator(opts *bind.TransactOpts, newSessionKeyValidator common.Address) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "upgradeSessionKeyValidator", newSessionKeyValidator)
}

// UpgradeSessionKeyValidator is a paid mutator transaction binding the contract method 0x4bd191a4.
//
// Solidity: function upgradeSessionKeyValidator(address newSessionKeyValidator) returns()
func (_VoltaAccount *VoltaAccountSession) UpgradeSessionKeyValidator(newSessionKeyValidator common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeSessionKeyValidator(&_VoltaAccount.TransactOpts, newSessionKeyValidator)
}

// UpgradeSessionKeyValidator is a paid mutator transaction binding the contract method 0x4bd191a4.
//
// Solidity: function upgradeSessionKeyValidator(address newSessionKeyValidator) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) UpgradeSessionKeyValidator(newSessionKeyValidator common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeSessionKeyValidator(&_VoltaAccount.TransactOpts, newSessionKeyValidator)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VoltaAccount *VoltaAccountTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VoltaAccount *VoltaAccountSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeTo(&_VoltaAccount.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeTo(&_VoltaAccount.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VoltaAccount *VoltaAccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VoltaAccount *VoltaAccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeToAndCall(&_VoltaAccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VoltaAccount *VoltaAccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VoltaAccount.Contract.UpgradeToAndCall(&_VoltaAccount.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_VoltaAccount *VoltaAccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_VoltaAccount *VoltaAccountSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.Contract.ValidateUserOp(&_VoltaAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_VoltaAccount *VoltaAccountTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.Contract.ValidateUserOp(&_VoltaAccount.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_VoltaAccount *VoltaAccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_VoltaAccount *VoltaAccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.Contract.WithdrawDepositTo(&_VoltaAccount.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_VoltaAccount *VoltaAccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VoltaAccount.Contract.WithdrawDepositTo(&_VoltaAccount.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VoltaAccount *VoltaAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoltaAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VoltaAccount *VoltaAccountSession) Receive() (*types.Transaction, error) {
	return _VoltaAccount.Contract.Receive(&_VoltaAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VoltaAccount *VoltaAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _VoltaAccount.Contract.Receive(&_VoltaAccount.TransactOpts)
}

// VoltaAccountAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VoltaAccount contract.
type VoltaAccountAdminChangedIterator struct {
	Event *VoltaAccountAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoltaAccountAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaAccountAdminChanged)
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
		it.Event = new(VoltaAccountAdminChanged)
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
func (it *VoltaAccountAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaAccountAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaAccountAdminChanged represents a AdminChanged event raised by the VoltaAccount contract.
type VoltaAccountAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VoltaAccount *VoltaAccountFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VoltaAccountAdminChangedIterator, error) {

	logs, sub, err := _VoltaAccount.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VoltaAccountAdminChangedIterator{contract: _VoltaAccount.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VoltaAccount *VoltaAccountFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VoltaAccountAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VoltaAccount.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaAccountAdminChanged)
				if err := _VoltaAccount.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_VoltaAccount *VoltaAccountFilterer) ParseAdminChanged(log types.Log) (*VoltaAccountAdminChanged, error) {
	event := new(VoltaAccountAdminChanged)
	if err := _VoltaAccount.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaAccountBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VoltaAccount contract.
type VoltaAccountBeaconUpgradedIterator struct {
	Event *VoltaAccountBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VoltaAccountBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaAccountBeaconUpgraded)
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
		it.Event = new(VoltaAccountBeaconUpgraded)
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
func (it *VoltaAccountBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaAccountBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaAccountBeaconUpgraded represents a BeaconUpgraded event raised by the VoltaAccount contract.
type VoltaAccountBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VoltaAccount *VoltaAccountFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VoltaAccountBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VoltaAccount.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VoltaAccountBeaconUpgradedIterator{contract: _VoltaAccount.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VoltaAccount *VoltaAccountFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VoltaAccountBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VoltaAccount.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaAccountBeaconUpgraded)
				if err := _VoltaAccount.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_VoltaAccount *VoltaAccountFilterer) ParseBeaconUpgraded(log types.Log) (*VoltaAccountBeaconUpgraded, error) {
	event := new(VoltaAccountBeaconUpgraded)
	if err := _VoltaAccount.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VoltaAccount contract.
type VoltaAccountInitializedIterator struct {
	Event *VoltaAccountInitialized // Event containing the contract specifics and raw log

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
func (it *VoltaAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaAccountInitialized)
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
		it.Event = new(VoltaAccountInitialized)
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
func (it *VoltaAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaAccountInitialized represents a Initialized event raised by the VoltaAccount contract.
type VoltaAccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VoltaAccount *VoltaAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*VoltaAccountInitializedIterator, error) {

	logs, sub, err := _VoltaAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VoltaAccountInitializedIterator{contract: _VoltaAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VoltaAccount *VoltaAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VoltaAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _VoltaAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaAccountInitialized)
				if err := _VoltaAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VoltaAccount *VoltaAccountFilterer) ParseInitialized(log types.Log) (*VoltaAccountInitialized, error) {
	event := new(VoltaAccountInitialized)
	if err := _VoltaAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoltaAccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VoltaAccount contract.
type VoltaAccountUpgradedIterator struct {
	Event *VoltaAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *VoltaAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoltaAccountUpgraded)
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
		it.Event = new(VoltaAccountUpgraded)
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
func (it *VoltaAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoltaAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoltaAccountUpgraded represents a Upgraded event raised by the VoltaAccount contract.
type VoltaAccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VoltaAccount *VoltaAccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VoltaAccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VoltaAccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VoltaAccountUpgradedIterator{contract: _VoltaAccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VoltaAccount *VoltaAccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VoltaAccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VoltaAccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoltaAccountUpgraded)
				if err := _VoltaAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VoltaAccount *VoltaAccountFilterer) ParseUpgraded(log types.Log) (*VoltaAccountUpgraded, error) {
	event := new(VoltaAccountUpgraded)
	if err := _VoltaAccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
