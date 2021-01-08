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
const BridgeCounterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"}],\"name\":\"incr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"}],\"name\":\"setDestStartNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeCounterBin is the compiled bytecode used for deploying new contracts.
var BridgeCounterBin = "0x608060405234801561001057600080fd5b50604051610853380380610853833981810160405281019061003291906100ce565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610140565b6000815190506100c881610129565b92915050565b6000602082840312156100e057600080fd5b60006100ee848285016100b9565b91505092915050565b600061010282610109565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b610132816100f7565b811461013d57600080fd5b50565b6107048061014f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633b8e78111461006757806371af5099146100835780638b8ddc7a146100b35780638da5cb5b146100e35780638dd1480214610101578063e78cea921461011d575b600080fd5b610081600480360381019061007c91906104d5565b61013b565b005b61009d600480360381019061009891906104ac565b61020f565b6040516100aa919061060a565b60405180910390f35b6100cd60048036038101906100c891906104ac565b6102ff565b6040516100da919061060a565b60405180910390f35b6100eb610326565b6040516100f891906105af565b60405180910390f35b61011b60048036038101906101169190610483565b61034c565b005b61012561041f565b60405161013291906105af565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c2906105ca565b60405180910390fd5b80600260008460ff1660ff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610297906105ea565b60405180910390fd5b600260008360ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790559050919050565b60026020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d3906105ca565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008135905061045381610689565b92915050565b600081359050610468816106a0565b92915050565b60008135905061047d816106b7565b92915050565b60006020828403121561049557600080fd5b60006104a384828501610444565b91505092915050565b6000602082840312156104be57600080fd5b60006104cc8482850161046e565b91505092915050565b600080604083850312156104e857600080fd5b60006104f68582860161046e565b925050602061050785828601610459565b9150509250929050565b61051a81610636565b82525050565b600061052d601483610625565b91507f73656e646572206973206e6f74202061646d696e0000000000000000000000006000830152602082019050919050565b600061056d601483610625565b91507f73656e646572206973206e6f74206272696467650000000000000000000000006000830152602082019050919050565b6105a981610668565b82525050565b60006020820190506105c46000830184610511565b92915050565b600060208201905081810360008301526105e381610520565b9050919050565b6000602082019050818103600083015261060381610560565b9050919050565b600060208201905061061f60008301846105a0565b92915050565b600082825260208201905092915050565b600061064182610648565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b61069281610636565b811461069d57600080fd5b50565b6106a981610668565b81146106b457600080fd5b50565b6106c08161067c565b81146106cb57600080fd5b5056fea2646970667358221220369f717b31e8f3ad8503ca4e75de37f6886ece37c499cfef9e157e4e3d0fc73064736f6c63430006040033"

// DeployBridgeCounter deploys a new Ethereum contract, binding an instance of BridgeCounter to it.
func DeployBridgeCounter(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address) (common.Address, *types.Transaction, *BridgeCounter, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeCounterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeCounterBin), backend, bridge_)
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

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeCounter *BridgeCounterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeCounter.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeCounter *BridgeCounterSession) Bridge() (common.Address, error) {
	return _BridgeCounter.Contract.Bridge(&_BridgeCounter.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeCounter *BridgeCounterCallerSession) Bridge() (common.Address, error) {
	return _BridgeCounter.Contract.Bridge(&_BridgeCounter.CallOpts)
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

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactor) SetBridge(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.contract.Transact(opts, "setBridge", bridge_)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterSession) SetBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.SetBridge(&_BridgeCounter.TransactOpts, bridge_)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridge_) returns()
func (_BridgeCounter *BridgeCounterTransactorSession) SetBridge(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeCounter.Contract.SetBridge(&_BridgeCounter.TransactOpts, bridge_)
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
