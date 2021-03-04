// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDepositNFTExecute

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

// IDepositNFTExecuteABI is the input ABI used to generate the binding from.
const IDepositNFTExecuteABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDepositNFTExecute is an auto generated Go binding around an Ethereum contract.
type IDepositNFTExecute struct {
	IDepositNFTExecuteCaller     // Read-only binding to the contract
	IDepositNFTExecuteTransactor // Write-only binding to the contract
	IDepositNFTExecuteFilterer   // Log filterer for contract events
}

// IDepositNFTExecuteCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDepositNFTExecuteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositNFTExecuteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDepositNFTExecuteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositNFTExecuteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDepositNFTExecuteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDepositNFTExecuteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDepositNFTExecuteSession struct {
	Contract     *IDepositNFTExecute // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDepositNFTExecuteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDepositNFTExecuteCallerSession struct {
	Contract *IDepositNFTExecuteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IDepositNFTExecuteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDepositNFTExecuteTransactorSession struct {
	Contract     *IDepositNFTExecuteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IDepositNFTExecuteRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDepositNFTExecuteRaw struct {
	Contract *IDepositNFTExecute // Generic contract binding to access the raw methods on
}

// IDepositNFTExecuteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDepositNFTExecuteCallerRaw struct {
	Contract *IDepositNFTExecuteCaller // Generic read-only contract binding to access the raw methods on
}

// IDepositNFTExecuteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDepositNFTExecuteTransactorRaw struct {
	Contract *IDepositNFTExecuteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDepositNFTExecute creates a new instance of IDepositNFTExecute, bound to a specific deployed contract.
func NewIDepositNFTExecute(address common.Address, backend bind.ContractBackend) (*IDepositNFTExecute, error) {
	contract, err := bindIDepositNFTExecute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDepositNFTExecute{IDepositNFTExecuteCaller: IDepositNFTExecuteCaller{contract: contract}, IDepositNFTExecuteTransactor: IDepositNFTExecuteTransactor{contract: contract}, IDepositNFTExecuteFilterer: IDepositNFTExecuteFilterer{contract: contract}}, nil
}

// NewIDepositNFTExecuteCaller creates a new read-only instance of IDepositNFTExecute, bound to a specific deployed contract.
func NewIDepositNFTExecuteCaller(address common.Address, caller bind.ContractCaller) (*IDepositNFTExecuteCaller, error) {
	contract, err := bindIDepositNFTExecute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositNFTExecuteCaller{contract: contract}, nil
}

// NewIDepositNFTExecuteTransactor creates a new write-only instance of IDepositNFTExecute, bound to a specific deployed contract.
func NewIDepositNFTExecuteTransactor(address common.Address, transactor bind.ContractTransactor) (*IDepositNFTExecuteTransactor, error) {
	contract, err := bindIDepositNFTExecute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDepositNFTExecuteTransactor{contract: contract}, nil
}

// NewIDepositNFTExecuteFilterer creates a new log filterer instance of IDepositNFTExecute, bound to a specific deployed contract.
func NewIDepositNFTExecuteFilterer(address common.Address, filterer bind.ContractFilterer) (*IDepositNFTExecuteFilterer, error) {
	contract, err := bindIDepositNFTExecute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDepositNFTExecuteFilterer{contract: contract}, nil
}

// bindIDepositNFTExecute binds a generic wrapper to an already deployed contract.
func bindIDepositNFTExecute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDepositNFTExecuteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositNFTExecute *IDepositNFTExecuteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositNFTExecute.Contract.IDepositNFTExecuteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositNFTExecute *IDepositNFTExecuteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.IDepositNFTExecuteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositNFTExecute *IDepositNFTExecuteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.IDepositNFTExecuteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDepositNFTExecute *IDepositNFTExecuteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDepositNFTExecute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDepositNFTExecute *IDepositNFTExecuteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDepositNFTExecute *IDepositNFTExecuteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 tokenId) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, tokenId *big.Int) (*types.Transaction, error) {
	return _IDepositNFTExecute.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipient, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 tokenId) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, tokenId *big.Int) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.Deposit(&_IDepositNFTExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, tokenId)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 tokenId) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, tokenId *big.Int) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.Deposit(&_IDepositNFTExecute.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, tokenId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 tokenId, bytes metadata) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, recipient common.Address, tokenId *big.Int, metadata []byte) (*types.Transaction, error) {
	return _IDepositNFTExecute.contract.Transact(opts, "executeProposal", resourceID, recipient, tokenId, metadata)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 tokenId, bytes metadata) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteSession) ExecuteProposal(resourceID [32]byte, recipient common.Address, tokenId *big.Int, metadata []byte) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.ExecuteProposal(&_IDepositNFTExecute.TransactOpts, resourceID, recipient, tokenId, metadata)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipient, uint256 tokenId, bytes metadata) returns()
func (_IDepositNFTExecute *IDepositNFTExecuteTransactorSession) ExecuteProposal(resourceID [32]byte, recipient common.Address, tokenId *big.Int, metadata []byte) (*types.Transaction, error) {
	return _IDepositNFTExecute.Contract.ExecuteProposal(&_IDepositNFTExecute.TransactOpts, resourceID, recipient, tokenId, metadata)
}
