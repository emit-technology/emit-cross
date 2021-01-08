// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeCounter

import (
	"math/big"
	"strings"

	sero "github.com/sero-cash/go-sero"
	"github.com/sero-cash/go-sero/accounts/abi"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/core/types"
	"github.com/sero-cash/go-sero/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = sero.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeCounterABI is the input ABI used to generate the binding from.
const BridgeCounterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"}],\"name\":\"incr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"}],\"name\":\"setDestStartNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeCounterBin is the compiled bytecode used for deploying new contracts.
var BridgeCounterBin = "0x608060405234801561001057600080fd5b5060405161042738038061042783398101604081905261002f9161005d565b600080546001600160a01b039092166001600160a01b0319928316179055600180549091163317905561008b565b60006020828403121561006e578081fd5b81516001600160a01b0381168114610084578182fd5b9392505050565b61038d8061009a6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633b8e78111461006757806371af50991461007c5780638b8ddc7a146100a55780638da5cb5b146100b85780638dd14802146100cd578063e78cea92146100e0575b600080fd5b61007a610075366004610289565b6100e8565b005b61008f61008a366004610268565b610150565b60405161009c9190610342565b60405180910390f35b61008f6100b3366004610268565b6101b4565b6100c06101d0565b60405161009c91906102d2565b61007a6100db36600461023a565b6101df565b6100c061022b565b6001546001600160a01b0316331461011b5760405162461bcd60e51b8152600401610112906102e6565b60405180910390fd5b60ff919091166000908152600260205260409020805467ffffffffffffffff191667ffffffffffffffff909216919091179055565b600080546001600160a01b0316331461017b5760405162461bcd60e51b815260040161011290610314565b5060ff166000908152600260205260409020805467ffffffffffffffff198116600167ffffffffffffffff928316019182161790915590565b60026020526000908152604090205467ffffffffffffffff1681565b6001546001600160a01b031681565b6001546001600160a01b031633146102095760405162461bcd60e51b8152600401610112906102e6565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031681565b60006020828403121561024b578081fd5b81356001600160a01b0381168114610261578182fd5b9392505050565b600060208284031215610279578081fd5b813560ff81168114610261578182fd5b6000806040838503121561029b578081fd5b823560ff811681146102ab578182fd5b9150602083013567ffffffffffffffff811681146102c7578182fd5b809150509250929050565b6001600160a01b0391909116815260200190565b60208082526014908201527339b2b73232b91034b9903737ba101030b236b4b760611b604082015260600190565b60208082526014908201527373656e646572206973206e6f742062726964676560601b604082015260600190565b67ffffffffffffffff9190911681526020019056fea2646970667358221220deaf75bf405f150325612ca99edfd21ec57f003730fa1196489a85febd7e4edf64736f6c63430006040033"

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
func (_BridgeCounter *BridgeCounterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_BridgeCounter *BridgeCounterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_BridgeCounter *BridgeCounterCaller) Bridge(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _BridgeCounter.contract.Call(opts, out, "bridge")
	return *ret0, err
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeCounter *BridgeCounterSession) Bridge() (common.ContractAddress, error) {
	return _BridgeCounter.Contract.Bridge(&_BridgeCounter.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_BridgeCounter *BridgeCounterCallerSession) Bridge() (common.ContractAddress, error) {
	return _BridgeCounter.Contract.Bridge(&_BridgeCounter.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x8b8ddc7a.
//
// Solidity: function depositCounts(uint8 ) view returns(uint64)
func (_BridgeCounter *BridgeCounterCaller) DepositCounts(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _BridgeCounter.contract.Call(opts, out, "depositCounts", arg0)
	return *ret0, err
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
func (_BridgeCounter *BridgeCounterCaller) Owner(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _BridgeCounter.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCounter *BridgeCounterSession) Owner() (common.ContractAddress, error) {
	return _BridgeCounter.Contract.Owner(&_BridgeCounter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeCounter *BridgeCounterCallerSession) Owner() (common.ContractAddress, error) {
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
