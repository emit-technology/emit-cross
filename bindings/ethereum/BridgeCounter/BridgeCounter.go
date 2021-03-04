// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeCounter

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

// BridgeCounterABI is the input ABI used to generate the binding from.
const BridgeCounterABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"}],\"name\":\"incr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"addBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"removeBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"}],\"name\":\"setDestStartNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeCounterBin is the compiled bytecode used for deploying new contracts.
var BridgeCounterBin = "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610872806100616000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638b8ddc7a1161005b5780638b8ddc7a146100ea5780638da5cb5b1461011a5780639712fdf814610138578063b11fd910146101545761007d565b806304df017d146100825780633b8e78111461009e57806371af5099146100ba575b600080fd5b61009c600480360381019061009791906105bb565b610184565b005b6100b860048036038101906100b3919061060d565b610265565b005b6100d460048036038101906100cf91906105e4565b610339565b6040516100e1919061076c565b60405180910390f35b61010460048036038101906100ff91906105e4565b610425565b604051610111919061076c565b60405180910390f35b61012261044c565b60405161012f91906106f6565b60405180910390f35b610152600480360381019061014d91906105bb565b610472565b005b61016e600480360381019061016991906105bb565b61055c565b60405161017b9190610711565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610214576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020b9061072c565b60405180910390fd5b6000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff021916905550565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ec9061072c565b60405180910390fd5b80600260008460ff1660ff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bd9061074c565b60405180910390fd5b600260008360ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790559050919050565b60026020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610502576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f99061072c565b60405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915054906101000a900460ff1681565b60008135905061058b816107f7565b92915050565b6000813590506105a08161080e565b92915050565b6000813590506105b581610825565b92915050565b6000602082840312156105cd57600080fd5b60006105db8482850161057c565b91505092915050565b6000602082840312156105f657600080fd5b6000610604848285016105a6565b91505092915050565b6000806040838503121561062057600080fd5b600061062e858286016105a6565b925050602061063f85828601610591565b9150509250929050565b61065281610798565b82525050565b610661816107aa565b82525050565b6000610674601483610787565b91507f73656e646572206973206e6f74202061646d696e0000000000000000000000006000830152602082019050919050565b60006106b4601483610787565b91507f73656e646572206973206e6f74206272696467650000000000000000000000006000830152602082019050919050565b6106f0816107d6565b82525050565b600060208201905061070b6000830184610649565b92915050565b60006020820190506107266000830184610658565b92915050565b6000602082019050818103600083015261074581610667565b9050919050565b60006020820190508181036000830152610765816106a7565b9050919050565b600060208201905061078160008301846106e7565b92915050565b600082825260208201905092915050565b60006107a3826107b6565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b61080081610798565b811461080b57600080fd5b50565b610817816107d6565b811461082257600080fd5b50565b61082e816107ea565b811461083957600080fd5b5056fea2646970667358221220fd9e38515c12a0e52d0df66762ddd60f8c5e6ac0fa85c39f1fd48e083428c4cc64736f6c63430006040033"

// DeployBridgeCounter deploys a new Ethereum contract, binding an instance of BridgeCounter to it.
func DeployBridgeCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeCounter, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeCounterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeCounter{BridgeCounterCaller: BridgeCounterCaller{contract: contract}, BridgeCounterTransactor: BridgeCounterTransactor{contract: contract}, BridgeCounterFilterer: BridgeCounterFilterer{contract: contract}}, nil
}

// BridgeCounter is an auto generated Go binding around an Ethereum contract.
type BridgeCounter struct {
	BridgeCounterCaller     // Read-only binding to the contract
	BridgeCounterTransactor // Write-only binding to the contract
	BridgeCounterFilterer   // Log filterer for contract events
}

// BridgeCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeCounterSession struct {
	Contract     *BridgeCounter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCounterCallerSession struct {
	Contract *BridgeCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeCounterTransactorSession struct {
	Contract     *BridgeCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeCounterRaw struct {
	Contract *BridgeCounter // Generic contract binding to access the raw methods on
}

// BridgeCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCounterCallerRaw struct {
	Contract *BridgeCounterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeCounterTransactorRaw struct {
	Contract *BridgeCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeCounter creates a new instance of BridgeCounter, bound to a specific deployed contract.
func NewBridgeCounter(address common.Address, backend bind.ContractBackend) (*BridgeCounter, error) {
	contract, err := bindBridgeCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeCounter{BridgeCounterCaller: BridgeCounterCaller{contract: contract}, BridgeCounterTransactor: BridgeCounterTransactor{contract: contract}, BridgeCounterFilterer: BridgeCounterFilterer{contract: contract}}, nil
}

// NewBridgeCounterCaller creates a new read-only instance of BridgeCounter, bound to a specific deployed contract.
func NewBridgeCounterCaller(address common.Address, caller bind.ContractCaller) (*BridgeCounterCaller, error) {
	contract, err := bindBridgeCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCounterCaller{contract: contract}, nil
}

// NewBridgeCounterTransactor creates a new write-only instance of BridgeCounter, bound to a specific deployed contract.
func NewBridgeCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeCounterTransactor, error) {
	contract, err := bindBridgeCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCounterTransactor{contract: contract}, nil
}

// NewBridgeCounterFilterer creates a new log filterer instance of BridgeCounter, bound to a specific deployed contract.
func NewBridgeCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeCounterFilterer, error) {
	contract, err := bindBridgeCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeCounterFilterer{contract: contract}, nil
}

// bindBridgeCounter binds a generic wrapper to an already deployed contract.
func bindBridgeCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeCounterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCounter *BridgeCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCounter.Contract.BridgeCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCounter *BridgeCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCounter.Contract.BridgeCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCounter *BridgeCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCounter.Contract.BridgeCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeCounter *BridgeCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeCounter *BridgeCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeCounter *BridgeCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeCounter.Contract.contract.Transact(opts, method, params...)
}

// ApprovedBridge is a free data retrieval call binding the contract method 0xb11fd910.
//
// Solidity: function approvedBridge(address ) view returns(bool)
func (_BridgeCounter *BridgeCounterCaller) ApprovedBridge(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeCounter.contract.Call(opts, &out, "approvedBridge", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedBridge is a free data retrieval call binding the contract method 0xb11fd910.
//
// Solidity: function approvedBridge(address ) view returns(bool)
func (_BridgeCounter *BridgeCounterSession) ApprovedBridge(arg0 common.Address) (bool, error) {
	return _BridgeCounter.Contract.ApprovedBridge(&_BridgeCounter.CallOpts, arg0)
}

// ApprovedBridge is a free data retrieval call binding the contract method 0xb11fd910.
//
// Solidity: function approvedBridge(address ) view returns(bool)
func (_BridgeCounter *BridgeCounterCallerSession) ApprovedBridge(arg0 common.Address) (bool, error) {
	return _BridgeCounter.Contract.ApprovedBridge(&_BridgeCounter.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x8b8ddc7a.
//
// Solidity: function depositCounts(uint8 ) view returns(uint64)
func (_BridgeCounter *BridgeCounterCaller) DepositCounts(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _BridgeCounter.contract.Call(opts, &out, "depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x8b8ddc7a.
//
// Solidity: function depositCounts(uint8 ) view returns(uint64)
func (_BridgeCounter *BridgeCounterSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _BridgeCounter.Contract.DepositCounts(&_BridgeCounter.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x8b8ddc7a.
//
// Solidity: function depositCounts(uint8 ) view returns(uint64)
func (_BridgeCounter *BridgeCounterCallerSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _BridgeCounter.Contract.DepositCounts(&_BridgeCounter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCounter *BridgeCounterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCounter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCounter *BridgeCounterSession) Owner() (common.Address, error) {
	return _BridgeCounter.Contract.Owner(&_BridgeCounter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCounter *BridgeCounterCallerSession) Owner() (common.Address, error) {
	return _BridgeCounter.Contract.Owner(&_BridgeCounter.CallOpts)
}

// AddBridge is a paid mutator transaction binding the contract method 0x9712fdf8.
//
// Solidity: function addBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactor) AddBridge(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.contract.Transact(opts, "addBridge", bridge_)
}

// AddBridge is a paid mutator transaction binding the contract method 0x9712fdf8.
//
// Solidity: function addBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterSession) AddBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.AddBridge(&_BridgeCounter.TransactOpts, bridge_)
}

// AddBridge is a paid mutator transaction binding the contract method 0x9712fdf8.
//
// Solidity: function addBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactorSession) AddBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.AddBridge(&_BridgeCounter.TransactOpts, bridge_)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_BridgeCounter *BridgeCounterTransactor) Incr(opts *bind.TransactOpts, destinationChainID uint8) (*types.Transaction, error) {
	return _BridgeCounter.contract.Transact(opts, "incr", destinationChainID)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_BridgeCounter *BridgeCounterSession) Incr(destinationChainID uint8) (*types.Transaction, error) {
	return _BridgeCounter.Contract.Incr(&_BridgeCounter.TransactOpts, destinationChainID)
}

// Incr is a paid mutator transaction binding the contract method 0x71af5099.
//
// Solidity: function incr(uint8 destinationChainID) returns(uint64)
func (_BridgeCounter *BridgeCounterTransactorSession) Incr(destinationChainID uint8) (*types.Transaction, error) {
	return _BridgeCounter.Contract.Incr(&_BridgeCounter.TransactOpts, destinationChainID)
}

// RemoveBridge is a paid mutator transaction binding the contract method 0x04df017d.
//
// Solidity: function removeBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactor) RemoveBridge(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.contract.Transact(opts, "removeBridge", bridge_)
}

// RemoveBridge is a paid mutator transaction binding the contract method 0x04df017d.
//
// Solidity: function removeBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterSession) RemoveBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.RemoveBridge(&_BridgeCounter.TransactOpts, bridge_)
}

// RemoveBridge is a paid mutator transaction binding the contract method 0x04df017d.
//
// Solidity: function removeBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactorSession) RemoveBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.RemoveBridge(&_BridgeCounter.TransactOpts, bridge_)
}

// SetDestStartNonce is a paid mutator transaction binding the contract method 0x3b8e7811.
//
// Solidity: function setDestStartNonce(uint8 destinationChainID, uint64 start) returns()
func (_BridgeCounter *BridgeCounterTransactor) SetDestStartNonce(opts *bind.TransactOpts, destinationChainID uint8, start uint64) (*types.Transaction, error) {
	return _BridgeCounter.contract.Transact(opts, "setDestStartNonce", destinationChainID, start)
}

// SetDestStartNonce is a paid mutator transaction binding the contract method 0x3b8e7811.
//
// Solidity: function setDestStartNonce(uint8 destinationChainID, uint64 start) returns()
func (_BridgeCounter *BridgeCounterSession) SetDestStartNonce(destinationChainID uint8, start uint64) (*types.Transaction, error) {
	return _BridgeCounter.Contract.SetDestStartNonce(&_BridgeCounter.TransactOpts, destinationChainID, start)
}

// SetDestStartNonce is a paid mutator transaction binding the contract method 0x3b8e7811.
//
// Solidity: function setDestStartNonce(uint8 destinationChainID, uint64 start) returns()
func (_BridgeCounter *BridgeCounterTransactorSession) SetDestStartNonce(destinationChainID uint8, start uint64) (*types.Transaction, error) {
	return _BridgeCounter.Contract.SetDestStartNonce(&_BridgeCounter.TransactOpts, destinationChainID, start)
}
