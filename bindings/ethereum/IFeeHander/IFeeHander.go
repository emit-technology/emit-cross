// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IFeeHander

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IFeeHanderABI is the input ABI used to generate the binding from.
const IFeeHanderABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reosurceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setResourceFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"gasFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"relayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IFeeHander is an auto generated Go binding around an Ethereum contract.
type IFeeHander struct {
	IFeeHanderCaller     // Read-only binding to the contract
	IFeeHanderTransactor // Write-only binding to the contract
	IFeeHanderFilterer   // Log filterer for contract events
}

// IFeeHanderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeHanderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeHanderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeHanderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeHanderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeHanderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeHanderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeHanderSession struct {
	Contract     *IFeeHander       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeHanderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeHanderCallerSession struct {
	Contract *IFeeHanderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IFeeHanderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeHanderTransactorSession struct {
	Contract     *IFeeHanderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IFeeHanderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeHanderRaw struct {
	Contract *IFeeHander // Generic contract binding to access the raw methods on
}

// IFeeHanderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeHanderCallerRaw struct {
	Contract *IFeeHanderCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeHanderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeHanderTransactorRaw struct {
	Contract *IFeeHanderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeHander creates a new instance of IFeeHander, bound to a specific deployed contract.
func NewIFeeHander(address common.Address, backend bind.ContractBackend) (*IFeeHander, error) {
	contract, err := bindIFeeHander(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeHander{IFeeHanderCaller: IFeeHanderCaller{contract: contract}, IFeeHanderTransactor: IFeeHanderTransactor{contract: contract}, IFeeHanderFilterer: IFeeHanderFilterer{contract: contract}}, nil
}

// NewIFeeHanderCaller creates a new read-only instance of IFeeHander, bound to a specific deployed contract.
func NewIFeeHanderCaller(address common.Address, caller bind.ContractCaller) (*IFeeHanderCaller, error) {
	contract, err := bindIFeeHander(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeHanderCaller{contract: contract}, nil
}

// NewIFeeHanderTransactor creates a new write-only instance of IFeeHander, bound to a specific deployed contract.
func NewIFeeHanderTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeHanderTransactor, error) {
	contract, err := bindIFeeHander(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeHanderTransactor{contract: contract}, nil
}

// NewIFeeHanderFilterer creates a new log filterer instance of IFeeHander, bound to a specific deployed contract.
func NewIFeeHanderFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeHanderFilterer, error) {
	contract, err := bindIFeeHander(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeHanderFilterer{contract: contract}, nil
}

// bindIFeeHander binds a generic wrapper to an already deployed contract.
func bindIFeeHander(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFeeHanderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeHander *IFeeHanderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeHander.Contract.IFeeHanderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeHander *IFeeHanderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeHander.Contract.IFeeHanderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeHander *IFeeHanderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeHander.Contract.IFeeHanderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeHander *IFeeHanderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeHander.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeHander *IFeeHanderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeHander.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeHander *IFeeHanderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeHander.Contract.contract.Transact(opts, method, params...)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_IFeeHander *IFeeHanderTransactor) SetResourceFeeContract(opts *bind.TransactOpts, reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IFeeHander.contract.Transact(opts, "setResourceFeeContract", reosurceId, tokenAddress)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_IFeeHander *IFeeHanderSession) SetResourceFeeContract(reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IFeeHander.Contract.SetResourceFeeContract(&_IFeeHander.TransactOpts, reosurceId, tokenAddress)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_IFeeHander *IFeeHanderTransactorSession) SetResourceFeeContract(reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _IFeeHander.Contract.SetResourceFeeContract(&_IFeeHander.TransactOpts, reosurceId, tokenAddress)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_IFeeHander *IFeeHanderTransactor) TransferFee(opts *bind.TransactOpts, resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFeeHander.contract.Transact(opts, "transferFee", resourceID, gasFeeRecipient, relayers, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_IFeeHander *IFeeHanderSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFeeHander.Contract.TransferFee(&_IFeeHander.TransactOpts, resourceID, gasFeeRecipient, relayers, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_IFeeHander *IFeeHanderTransactorSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFeeHander.Contract.TransferFee(&_IFeeHander.TransactOpts, resourceID, gasFeeRecipient, relayers, amount)
}
