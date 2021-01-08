// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ICrossFee

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

// ICrossFeeABI is the input ABI used to generate the binding from.
const ICrossFeeABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"calCrossFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"crossFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ICrossFee is an auto generated Go binding around an Ethereum contract.
type ICrossFee struct {
	ICrossFeeCaller     // Read-only binding to the contract
	ICrossFeeTransactor // Write-only binding to the contract
	ICrossFeeFilterer   // Log filterer for contract events
}

// ICrossFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossFeeSession struct {
	Contract     *ICrossFee        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICrossFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossFeeCallerSession struct {
	Contract *ICrossFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ICrossFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossFeeTransactorSession struct {
	Contract     *ICrossFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICrossFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossFeeRaw struct {
	Contract *ICrossFee // Generic contract binding to access the raw methods on
}

// ICrossFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossFeeCallerRaw struct {
	Contract *ICrossFeeCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossFeeTransactorRaw struct {
	Contract *ICrossFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossFee creates a new instance of ICrossFee, bound to a specific deployed contract.
func NewICrossFee(address common.Address, backend bind.ContractBackend) (*ICrossFee, error) {
	contract, err := bindICrossFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossFee{ICrossFeeCaller: ICrossFeeCaller{contract: contract}, ICrossFeeTransactor: ICrossFeeTransactor{contract: contract}, ICrossFeeFilterer: ICrossFeeFilterer{contract: contract}}, nil
}

// NewICrossFeeCaller creates a new read-only instance of ICrossFee, bound to a specific deployed contract.
func NewICrossFeeCaller(address common.Address, caller bind.ContractCaller) (*ICrossFeeCaller, error) {
	contract, err := bindICrossFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossFeeCaller{contract: contract}, nil
}

// NewICrossFeeTransactor creates a new write-only instance of ICrossFee, bound to a specific deployed contract.
func NewICrossFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossFeeTransactor, error) {
	contract, err := bindICrossFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossFeeTransactor{contract: contract}, nil
}

// NewICrossFeeFilterer creates a new log filterer instance of ICrossFee, bound to a specific deployed contract.
func NewICrossFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossFeeFilterer, error) {
	contract, err := bindICrossFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossFeeFilterer{contract: contract}, nil
}

// bindICrossFee binds a generic wrapper to an already deployed contract.
func bindICrossFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICrossFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossFee *ICrossFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossFee.Contract.ICrossFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossFee *ICrossFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossFee.Contract.ICrossFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossFee *ICrossFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossFee.Contract.ICrossFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossFee *ICrossFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossFee *ICrossFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossFee *ICrossFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossFee.Contract.contract.Transact(opts, method, params...)
}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 crossFee, uint256 gasFee)
func (_ICrossFee *ICrossFeeCaller) CalCrossFee(opts *bind.CallOpts, resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	CrossFee *big.Int
	GasFee   *big.Int
}, error) {
	var out []interface{}
	err := _ICrossFee.contract.Call(opts, &out, "calCrossFee", resourceId, inputAmount, gasPrice)

	outstruct := new(struct {
		CrossFee *big.Int
		GasFee   *big.Int
	})

	outstruct.CrossFee = out[0].(*big.Int)
	outstruct.GasFee = out[1].(*big.Int)

	return *outstruct, err

}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 crossFee, uint256 gasFee)
func (_ICrossFee *ICrossFeeSession) CalCrossFee(resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	CrossFee *big.Int
	GasFee   *big.Int
}, error) {
	return _ICrossFee.Contract.CalCrossFee(&_ICrossFee.CallOpts, resourceId, inputAmount, gasPrice)
}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 crossFee, uint256 gasFee)
func (_ICrossFee *ICrossFeeCallerSession) CalCrossFee(resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	CrossFee *big.Int
	GasFee   *big.Int
}, error) {
	return _ICrossFee.Contract.CalCrossFee(&_ICrossFee.CallOpts, resourceId, inputAmount, gasPrice)
}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_ICrossFee *ICrossFeeCaller) EstimateFee(opts *bind.CallOpts, resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICrossFee.contract.Call(opts, &out, "estimateFee", resourceId, inputAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_ICrossFee *ICrossFeeSession) EstimateFee(resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	return _ICrossFee.Contract.EstimateFee(&_ICrossFee.CallOpts, resourceId, inputAmount)
}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_ICrossFee *ICrossFeeCallerSession) EstimateFee(resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	return _ICrossFee.Contract.EstimateFee(&_ICrossFee.CallOpts, resourceId, inputAmount)
}
