// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDepositFTExecute

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

// IDepositFTExecuteABI is the input ABI used to generate the binding from.
const IDepositFTExecuteABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"relayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDepositFTExecute is an auto generated Go binding around an Ethereum contract.
type IDepositFTExecute struct {
	IDepositFTExecuteCaller     // Read-only binding to the contract
	IDepositFTExecuteTransactor // Write-only binding to the contract
	IDepositFTExecuteFilterer   // Log filterer for contract events
}

// IDepositFTExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositFTExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositFTExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositFTExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositFTExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositFTExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositFTExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositFTExecuteSession struct {
	Contract     *IDepositFTExecute // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IDepositFTExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositFTExecuteCallerSession struct {
	Contract *IDepositFTExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IDepositFTExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositFTExecuteTransactorSession struct {
	Contract     *IDepositFTExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IDepositFTExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositFTExecuteRaw struct {
	Contract *IDepositFTExecute // Generic contract binding to access the raw methods on
}

// IDepositFTExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositFTExecuteCallerRaw struct {
	Contract *IDepositFTExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositFTExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositFTExecuteTransactorRaw struct {
	Contract *IDepositFTExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDepositFTExecute creates a new instance of IDepositFTExecute, bound to a specific deployed contract.
func NewIDepositFTExecute(address common.Address, backend bind.ContractBackend) (*IDepositFTExecute, error) {
	contract, err := bindIDepositFTExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDepositFTExecute{IDepositFTExecuteCaller: IDepositFTExecuteCaller{contract: contract}, IDepositFTExecuteTransactor: IDepositFTExecuteTransactor{contract: contract}, IDepositFTExecuteFilterer: IDepositFTExecuteFilterer{contract: contract}}, nil
}

// NewIDepositFTExecuteCaller creates a new read-only instance of IDepositFTExecute, bound to a specific deployed contract.
func NewIDepositFTExecuteCaller(address common.Address, caller bind.ContractCaller) (*IDepositFTExecuteCaller, error) {
	contract, err := bindIDepositFTExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositFTExecuteCaller{contract: contract}, nil
}

// NewIDepositFTExecuteTransactor creates a new write-only instance of IDepositFTExecute, bound to a specific deployed contract.
func NewIDepositFTExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositFTExecuteTransactor, error) {
	contract, err := bindIDepositFTExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositFTExecuteTransactor{contract: contract}, nil
}

// NewIDepositFTExecuteFilterer creates a new log filterer instance of IDepositFTExecute, bound to a specific deployed contract.
func NewIDepositFTExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositFTExecuteFilterer, error) {
	contract, err := bindIDepositFTExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositFTExecuteFilterer{contract: contract}, nil
}

// bindIDepositFTExecute binds a generic wrapper to an already deployed contract.
func bindIDepositFTExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDepositFTExecuteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositFTExecute *IDepositFTExecuteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositFTExecute.Contract.IDepositFTExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositFTExecute *IDepositFTExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.IDepositFTExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositFTExecute *IDepositFTExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.IDepositFTExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositFTExecute *IDepositFTExecuteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositFTExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositFTExecute *IDepositFTExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositFTExecute *IDepositFTExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.Deposit(&_IDepositFTExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.Deposit(&_IDepositFTExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.contract.Transact(opts, "executeProposal", resourceID, recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteSession) ExecuteProposal(resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.ExecuteProposal(&_IDepositFTExecute.TransactOpts, resourceID, recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 amount) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactorSession) ExecuteProposal(resourceID [32]byte, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.ExecuteProposal(&_IDepositFTExecute.TransactOpts, resourceID, recipient, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address recipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactor) TransferFee(opts *bind.TransactOpts, resourceID [32]byte, recipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.contract.Transact(opts, "transferFee", resourceID, recipient, gasFee, relayers, relayFee)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address recipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_IDepositFTExecute *IDepositFTExecuteSession) TransferFee(resourceID [32]byte, recipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.TransferFee(&_IDepositFTExecute.TransactOpts, resourceID, recipient, gasFee, relayers, relayFee)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address recipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_IDepositFTExecute *IDepositFTExecuteTransactorSession) TransferFee(resourceID [32]byte, recipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _IDepositFTExecute.Contract.TransferFee(&_IDepositFTExecute.TransactOpts, resourceID, recipient, gasFee, relayers, relayFee)
}
