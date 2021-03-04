// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721FeeHandler

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

// ERC721FeeHandlerABI is the input ABI used to generate the binding from.
const ERC721FeeHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reosurceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setResourceFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"gasFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"relayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721FeeHandlerBin is the compiled bytecode used for deploying new contracts.
var ERC721FeeHandlerBin = "0x608060405234801561001057600080fd5b50604051610a75380380610a758339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b600081519050610087816100e8565b92915050565b60006020828403121561009f57600080fd5b60006100ad84828501610078565b91505092915050565b60006100c1826100c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6100f1816100b6565b81146100fc57600080fd5b50565b6109678061010e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063318c136e1461005157806388617c2e1461006f57806395601f091461008b578063b581b552146100a7575b600080fd5b6100596100c3565b604051610066919061074e565b60405180910390f35b610089600480360381019061008491906105a8565b6100e8565b005b6100a560048036038101906100a091906104f4565b61013b565b005b6100c160048036038101906100bc919061056c565b610152565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100f06101b0565b60006001600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610133818684610241565b505050505050565b600083905061014c818430856102b9565b50505050565b61015a6101b0565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461023f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610236906107e9565b60405180910390fd5b565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016102819291906107a0565b600060405180830381600087803b15801561029b57600080fd5b505af11580156102af573d6000803e3d6000fd5b5050505050505050565b61033c846323b872dd60e01b8585856040516024016102da93929190610769565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610342565b50505050565b600060608373ffffffffffffffffffffffffffffffffffffffff168360405161036b9190610737565b6000604051808303816000865af19150503d80600081146103a8576040519150601f19603f3d011682016040523d82523d6000602084013e6103ad565b606091505b5091509150816103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e990610809565b60405180910390fd5b60008151111561045057808060200190518101906104109190610543565b61044f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610446906107c9565b60405180910390fd5b5b50505050565b600081359050610465816108d5565b92915050565b60008083601f84011261047d57600080fd5b8235905067ffffffffffffffff81111561049657600080fd5b6020830191508360208202830111156104ae57600080fd5b9250929050565b6000815190506104c4816108ec565b92915050565b6000813590506104d981610903565b92915050565b6000813590506104ee8161091a565b92915050565b60008060006060848603121561050957600080fd5b600061051786828701610456565b935050602061052886828701610456565b9250506040610539868287016104df565b9150509250925092565b60006020828403121561055557600080fd5b6000610563848285016104b5565b91505092915050565b6000806040838503121561057f57600080fd5b600061058d858286016104ca565b925050602061059e85828601610456565b9150509250929050565b6000806000806000608086880312156105c057600080fd5b60006105ce888289016104ca565b95505060206105df88828901610456565b945050604086013567ffffffffffffffff8111156105fc57600080fd5b6106088882890161046b565b9350935050606061061b888289016104df565b9150509295509295909350565b61063181610850565b82525050565b600061064282610829565b61064c8185610834565b935061065c8185602086016108a2565b80840191505092915050565b600061067560208361083f565b91507f45524332303a206f7065726174696f6e20646964206e6f7420737563636565646000830152602082019050919050565b60006106b5601e8361083f565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006106f560128361083f565b91507f45524332303a2063616c6c206661696c656400000000000000000000000000006000830152602082019050919050565b61073181610898565b82525050565b60006107438284610637565b915081905092915050565b60006020820190506107636000830184610628565b92915050565b600060608201905061077e6000830186610628565b61078b6020830185610628565b6107986040830184610728565b949350505050565b60006040820190506107b56000830185610628565b6107c26020830184610728565b9392505050565b600060208201905081810360008301526107e281610668565b9050919050565b60006020820190508181036000830152610802816106a8565b9050919050565b60006020820190508181036000830152610822816106e8565b9050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600061085b82610878565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b838110156108c05780820151818401526020810190506108a5565b838111156108cf576000848401525b50505050565b6108de81610850565b81146108e957600080fd5b50565b6108f581610862565b811461090057600080fd5b50565b61090c8161086e565b811461091757600080fd5b50565b61092381610898565b811461092e57600080fd5b5056fea26469706673582212209e0d22c3aa23e37aa79cb08c8e64a56b7883ff304ca46ca389077af3a2e6ded264736f6c63430006040033"

// DeployERC721FeeHandler deploys a new Ethereum contract, binding an instance of ERC721FeeHandler to it.
func DeployERC721FeeHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *ERC721FeeHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721FeeHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721FeeHandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721FeeHandler{ERC721FeeHandlerCaller: ERC721FeeHandlerCaller{contract: contract}, ERC721FeeHandlerTransactor: ERC721FeeHandlerTransactor{contract: contract}, ERC721FeeHandlerFilterer: ERC721FeeHandlerFilterer{contract: contract}}, nil
}

// ERC721FeeHandler is an auto generated Go binding around an Ethereum contract.
type ERC721FeeHandler struct {
	ERC721FeeHandlerCaller     // Read-only binding to the contract
	ERC721FeeHandlerTransactor // Write-only binding to the contract
	ERC721FeeHandlerFilterer   // Log filterer for contract events
}

// ERC721FeeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721FeeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FeeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721FeeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FeeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721FeeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FeeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721FeeHandlerSession struct {
	Contract     *ERC721FeeHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721FeeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721FeeHandlerCallerSession struct {
	Contract *ERC721FeeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721FeeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721FeeHandlerTransactorSession struct {
	Contract     *ERC721FeeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721FeeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721FeeHandlerRaw struct {
	Contract *ERC721FeeHandler // Generic contract binding to access the raw methods on
}

// ERC721FeeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721FeeHandlerCallerRaw struct {
	Contract *ERC721FeeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721FeeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721FeeHandlerTransactorRaw struct {
	Contract *ERC721FeeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721FeeHandler creates a new instance of ERC721FeeHandler, bound to a specific deployed contract.
func NewERC721FeeHandler(address common.Address, backend bind.ContractBackend) (*ERC721FeeHandler, error) {
	contract, err := bindERC721FeeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721FeeHandler{ERC721FeeHandlerCaller: ERC721FeeHandlerCaller{contract: contract}, ERC721FeeHandlerTransactor: ERC721FeeHandlerTransactor{contract: contract}, ERC721FeeHandlerFilterer: ERC721FeeHandlerFilterer{contract: contract}}, nil
}

// NewERC721FeeHandlerCaller creates a new read-only instance of ERC721FeeHandler, bound to a specific deployed contract.
func NewERC721FeeHandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721FeeHandlerCaller, error) {
	contract, err := bindERC721FeeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FeeHandlerCaller{contract: contract}, nil
}

// NewERC721FeeHandlerTransactor creates a new write-only instance of ERC721FeeHandler, bound to a specific deployed contract.
func NewERC721FeeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721FeeHandlerTransactor, error) {
	contract, err := bindERC721FeeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FeeHandlerTransactor{contract: contract}, nil
}

// NewERC721FeeHandlerFilterer creates a new log filterer instance of ERC721FeeHandler, bound to a specific deployed contract.
func NewERC721FeeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721FeeHandlerFilterer, error) {
	contract, err := bindERC721FeeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721FeeHandlerFilterer{contract: contract}, nil
}

// bindERC721FeeHandler binds a generic wrapper to an already deployed contract.
func bindERC721FeeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721FeeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721FeeHandler *ERC721FeeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721FeeHandler.Contract.ERC721FeeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721FeeHandler *ERC721FeeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.ERC721FeeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721FeeHandler *ERC721FeeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.ERC721FeeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721FeeHandler *ERC721FeeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721FeeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721FeeHandler *ERC721FeeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721FeeHandler *ERC721FeeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721FeeHandler *ERC721FeeHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721FeeHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721FeeHandler *ERC721FeeHandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721FeeHandler.Contract.BridgeAddress(&_ERC721FeeHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721FeeHandler *ERC721FeeHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721FeeHandler.Contract.BridgeAddress(&_ERC721FeeHandler.CallOpts)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.FundERC20(&_ERC721FeeHandler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.FundERC20(&_ERC721FeeHandler.TransactOpts, tokenAddress, owner, amount)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactor) SetResourceFeeContract(opts *bind.TransactOpts, reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721FeeHandler.contract.Transact(opts, "setResourceFeeContract", reosurceId, tokenAddress)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerSession) SetResourceFeeContract(reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.SetResourceFeeContract(&_ERC721FeeHandler.TransactOpts, reosurceId, tokenAddress)
}

// SetResourceFeeContract is a paid mutator transaction binding the contract method 0xb581b552.
//
// Solidity: function setResourceFeeContract(bytes32 reosurceId, address tokenAddress) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactorSession) SetResourceFeeContract(reosurceId [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.SetResourceFeeContract(&_ERC721FeeHandler.TransactOpts, reosurceId, tokenAddress)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactor) TransferFee(opts *bind.TransactOpts, resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.contract.Transact(opts, "transferFee", resourceID, gasFeeRecipient, relayers, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.TransferFee(&_ERC721FeeHandler.TransactOpts, resourceID, gasFeeRecipient, relayers, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0x88617c2e.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, address[] relayers, uint256 amount) returns()
func (_ERC721FeeHandler *ERC721FeeHandlerTransactorSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, relayers []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC721FeeHandler.Contract.TransferFee(&_ERC721FeeHandler.TransactOpts, resourceID, gasFeeRecipient, relayers, amount)
}
