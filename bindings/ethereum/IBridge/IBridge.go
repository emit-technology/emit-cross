// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBridge

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

// IBridgeABI is the input ABI used to generate the binding from.
const IBridgeABI = "[{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"}],\"name\":\"minCrossAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"}],\"name\":\"resourceIDToLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IBridge is an auto generated Go binding around an Ethereum contract.
type IBridge struct {
	IBridgeCaller     // Read-only binding to the contract
	IBridgeTransactor // Write-only binding to the contract
	IBridgeFilterer   // Log filterer for contract events
}

// IBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeSession struct {
	Contract     *IBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCallerSession struct {
	Contract *IBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeTransactorSession struct {
	Contract     *IBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeRaw struct {
	Contract *IBridge // Generic contract binding to access the raw methods on
}

// IBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCallerRaw struct {
	Contract *IBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeTransactorRaw struct {
	Contract *IBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridge creates a new instance of IBridge, bound to a specific deployed contract.
func NewIBridge(address common.Address, backend bind.ContractBackend) (*IBridge, error) {
	contract, err := bindIBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridge{IBridgeCaller: IBridgeCaller{contract: contract}, IBridgeTransactor: IBridgeTransactor{contract: contract}, IBridgeFilterer: IBridgeFilterer{contract: contract}}, nil
}

// NewIBridgeCaller creates a new read-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCaller, error) {
	contract, err := bindIBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCaller{contract: contract}, nil
}

// NewIBridgeTransactor creates a new write-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeTransactor, error) {
	contract, err := bindIBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeTransactor{contract: contract}, nil
}

// NewIBridgeFilterer creates a new log filterer instance of IBridge, bound to a specific deployed contract.
func NewIBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFilterer, error) {
	contract, err := bindIBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFilterer{contract: contract}, nil
}

// bindIBridge binds a generic wrapper to an already deployed contract.
func bindIBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.IBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transact(opts, method, params...)
}

// MinCrossAmount is a free data retrieval call binding the contract method 0xfca08900.
//
// Solidity: function minCrossAmount(bytes32 resourceId) view returns(uint256)
func (_IBridge *IBridgeCaller) MinCrossAmount(opts *bind.CallOpts, resourceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "minCrossAmount", resourceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCrossAmount is a free data retrieval call binding the contract method 0xfca08900.
//
// Solidity: function minCrossAmount(bytes32 resourceId) view returns(uint256)
func (_IBridge *IBridgeSession) MinCrossAmount(resourceId [32]byte) (*big.Int, error) {
	return _IBridge.Contract.MinCrossAmount(&_IBridge.CallOpts, resourceId)
}

// MinCrossAmount is a free data retrieval call binding the contract method 0xfca08900.
//
// Solidity: function minCrossAmount(bytes32 resourceId) view returns(uint256)
func (_IBridge *IBridgeCallerSession) MinCrossAmount(resourceId [32]byte) (*big.Int, error) {
	return _IBridge.Contract.MinCrossAmount(&_IBridge.CallOpts, resourceId)
}

// ResourceIDToLimit is a free data retrieval call binding the contract method 0xfa6bbe47.
//
// Solidity: function resourceIDToLimit(bytes32 resourceId) view returns(uint256 minAmount, uint256 maxAmount)
func (_IBridge *IBridgeCaller) ResourceIDToLimit(opts *bind.CallOpts, resourceId [32]byte) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
}, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "resourceIDToLimit", resourceId)

	outstruct := new(struct {
		MinAmount *big.Int
		MaxAmount *big.Int
	})

	outstruct.MinAmount = out[0].(*big.Int)
	outstruct.MaxAmount = out[1].(*big.Int)

	return *outstruct, err

}

// ResourceIDToLimit is a free data retrieval call binding the contract method 0xfa6bbe47.
//
// Solidity: function resourceIDToLimit(bytes32 resourceId) view returns(uint256 minAmount, uint256 maxAmount)
func (_IBridge *IBridgeSession) ResourceIDToLimit(resourceId [32]byte) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
}, error) {
	return _IBridge.Contract.ResourceIDToLimit(&_IBridge.CallOpts, resourceId)
}

// ResourceIDToLimit is a free data retrieval call binding the contract method 0xfa6bbe47.
//
// Solidity: function resourceIDToLimit(bytes32 resourceId) view returns(uint256 minAmount, uint256 maxAmount)
func (_IBridge *IBridgeCallerSession) ResourceIDToLimit(resourceId [32]byte) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
}, error) {
	return _IBridge.Contract.ResourceIDToLimit(&_IBridge.CallOpts, resourceId)
}

// ChainID is a paid mutator transaction binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() returns(uint8)
func (_IBridge *IBridgeTransactor) ChainID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "_chainID")
}

// ChainID is a paid mutator transaction binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() returns(uint8)
func (_IBridge *IBridgeSession) ChainID() (*types.Transaction, error) {
	return _IBridge.Contract.ChainID(&_IBridge.TransactOpts)
}

// ChainID is a paid mutator transaction binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() returns(uint8)
func (_IBridge *IBridgeTransactorSession) ChainID() (*types.Transaction, error) {
	return _IBridge.Contract.ChainID(&_IBridge.TransactOpts)
}

// DepositFT is a paid mutator transaction binding the contract method 0x1ca2388d.
//
// Solidity: function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes recipient, uint256 amount) payable returns()
func (_IBridge *IBridgeTransactor) DepositFT(opts *bind.TransactOpts, destinationChainID uint8, resourceID [32]byte, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "depositFT", destinationChainID, resourceID, recipient, amount)
}

// DepositFT is a paid mutator transaction binding the contract method 0x1ca2388d.
//
// Solidity: function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes recipient, uint256 amount) payable returns()
func (_IBridge *IBridgeSession) DepositFT(destinationChainID uint8, resourceID [32]byte, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.DepositFT(&_IBridge.TransactOpts, destinationChainID, resourceID, recipient, amount)
}

// DepositFT is a paid mutator transaction binding the contract method 0x1ca2388d.
//
// Solidity: function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes recipient, uint256 amount) payable returns()
func (_IBridge *IBridgeTransactorSession) DepositFT(destinationChainID uint8, resourceID [32]byte, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.DepositFT(&_IBridge.TransactOpts, destinationChainID, resourceID, recipient, amount)
}
