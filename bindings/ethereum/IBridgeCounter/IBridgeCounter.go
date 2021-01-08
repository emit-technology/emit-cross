// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBridgeCounter

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

// IBridgeCounterABI is the input ABI used to generate the binding from.
const IBridgeCounterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"}],\"name\":\"incr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBridgeCounter is an auto generated Go binding around an Ethereum contract.
type IBridgeCounter struct {
	IBridgeCounterCaller     // Read-only binding to the contract
	IBridgeCounterTransactor // Write-only binding to the contract
	IBridgeCounterFilterer   // Log filterer for contract events
}

// IBridgeCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeCounterSession struct {
	Contract     *IBridgeCounter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCounterCallerSession struct {
	Contract *IBridgeCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IBridgeCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeCounterTransactorSession struct {
	Contract     *IBridgeCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IBridgeCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeCounterRaw struct {
	Contract *IBridgeCounter // Generic contract binding to access the raw methods on
}

// IBridgeCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCounterCallerRaw struct {
	Contract *IBridgeCounterCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeCounterTransactorRaw struct {
	Contract *IBridgeCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeCounter creates a new instance of IBridgeCounter, bound to a specific deployed contract.
func NewIBridgeCounter(address common.Address, backend bind.ContractBackend) (*IBridgeCounter, error) {
	contract, err := bindIBridgeCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeCounter{IBridgeCounterCaller: IBridgeCounterCaller{contract: contract}, IBridgeCounterTransactor: IBridgeCounterTransactor{contract: contract}, IBridgeCounterFilterer: IBridgeCounterFilterer{contract: contract}}, nil
}

// NewIBridgeCounterCaller creates a new read-only instance of IBridgeCounter, bound to a specific deployed contract.
func NewIBridgeCounterCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCounterCaller, error) {
	contract, err := bindIBridgeCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCounterCaller{contract: contract}, nil
}

// NewIBridgeCounterTransactor creates a new write-only instance of IBridgeCounter, bound to a specific deployed contract.
func NewIBridgeCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeCounterTransactor, error) {
	contract, err := bindIBridgeCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCounterTransactor{contract: contract}, nil
}

// NewIBridgeCounterFilterer creates a new log filterer instance of IBridgeCounter, bound to a specific deployed contract.
func NewIBridgeCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeCounterFilterer, error) {
	contract, err := bindIBridgeCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeCounterFilterer{contract: contract}, nil
}

// bindIBridgeCounter binds a generic wrapper to an already deployed contract.
func bindIBridgeCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeCounterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeCounter *IBridgeCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeCounter.Contract.IBridgeCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeCounter *IBridgeCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.IBridgeCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeCounter *IBridgeCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.IBridgeCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeCounter *IBridgeCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeCounter *IBridgeCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeCounter *IBridgeCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.contract.Transact(opts, method, params...)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_IBridgeCounter *IBridgeCounterTransactor) Incr(opts *bind.TransactOpts, destinationChainID uint8) (*types.Transaction, error) {
	return _IBridgeCounter.contract.Transact(opts, "incr", destinationChainID)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_IBridgeCounter *IBridgeCounterSession) Incr(destinationChainID uint8) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.Incr(&_IBridgeCounter.TransactOpts, destinationChainID)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_IBridgeCounter *IBridgeCounterTransactorSession) Incr(destinationChainID uint8) (*types.Transaction, error) {
	return _IBridgeCounter.Contract.Incr(&_IBridgeCounter.TransactOpts, destinationChainID)
}
