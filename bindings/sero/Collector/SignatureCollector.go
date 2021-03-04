// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Collector

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

// SignatureCollectorABI is the input ABI used to generate the binding from.
const SignatureCollectorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAccessAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"resourceChainId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"destinationChainId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumSignatureCollector.TransferType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumSignatureCollector.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"SignProposalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"}],\"name\":\"_destinationStartNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAccess\",\"outputs\":[{\"internalType\":\"contractIBridgeAccess\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAccessAddr\",\"type\":\"address\"}],\"name\":\"changeBridgeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationId\",\"type\":\"uint8\"}],\"name\":\"coId\",\"outputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"destinationProposals\",\"outputs\":[{\"internalType\":\"enumSignatureCollector.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"srcId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDestinationStartNoncee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"srcId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"getProposalSignatures\",\"outputs\":[{\"internalType\":\"enumSignatureCollector.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"srcId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"hasSignedProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"srcId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"startNonce\",\"type\":\"uint64\"}],\"name\":\"setDestinationStartNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"srcId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"enumSignatureCollector.TransferType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"signProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SignatureCollectorBin is the compiled bytecode used for deploying new contracts.
var SignatureCollectorBin = "0x608060405234801561001057600080fd5b50604051610f4a380380610f4a83398101604081905261002f91610062565b60018054336001600160a01b031991821617909155600080549091166001600160a01b0392909216919091179055610090565b600060208284031215610073578081fd5b81516001600160a01b0381168114610089578182fd5b9392505050565b610eab8061009f6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638b98dbad116100715780638b98dbad146101605780638da5cb5b14610180578063af4f625714610188578063d1b0628c146101a8578063ea1d71c1146101bb578063f2fde38b146101ce576100b4565b80631b350ac1146100b95780633ee27649146100d7578063533703da146100f7578063602b5d7e1461010c5780637de0ac851461012c57806380c22a241461013f575b600080fd5b6100c16101e1565b6040516100ce9190610c1b565b60405180910390f35b6100ea6100e5366004610a51565b6101f0565b6040516100ce9190610c3a565b61010a610105366004610b66565b610213565b005b61011f61011a366004610a35565b6105fb565b6040516100ce9190610e19565b61010a61013a3660046109cf565b610616565b61015261014d366004610abd565b61064f565b6040516100ce929190610c4d565b61017361016e366004610a89565b6107a8565b6040516100ce9190610e2d565b6100c16107bd565b61019b610196366004610b02565b6107cc565b6040516100ce9190610c2f565b61010a6101b6366004610abd565b61082a565b61011f6101c9366004610a89565b61088d565b61010a6101dc3660046109cf565b6108c7565b6000546001600160a01b031681565b600460209081526000928352604080842090915290825290206001015460ff1681565b600054604051630a83aaa960e31b81526001600160a01b039091169063541d554890610243903390600401610c1b565b60206040518083038186803b15801561025b57600080fd5b505afa15801561026f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029391906109fd565b6102b85760405162461bcd60e51b81526004016102af90610d45565b60405180910390fd5b60008360028111156102c657fe5b60ff16116102e65760405162461bcd60e51b81526004016102af90610d6a565b600360006102f488886107a8565b6001600160481b031681526020810191909152604001600020546001600160401b03908116908516101561033a5760405162461bcd60e51b81526004016102af90610d1b565b60006004600061034a89896107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b038916825290925290206001808201549192509060ff16600281111561039157fe5b11156103af5760405162461bcd60e51b81526004016102af90610d98565b600260006103bd89896107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b0389168252835281812033825290925290205460ff16156104155760405162461bcd60e51b81526004016102af90610de2565b6001600260006104258a8a6107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b038a16825283528181203382529092529020805460ff1916911515919091179055600181015460ff16600281111561047e57fe5b6104e1576001818101805460ff191682800217905550846001600160401b03168660ff168860ff167ff3e1949d167ac18929f02f5eb320d0379bac1f3196f1bf0ce9ececdd44b6cf458760016040516104d8929190610cf5565b60405180910390a45b80546001810182556000828152602090206104fe9101848461090c565b5060008060009054906101000a90046001600160a01b03166001600160a01b031663f0e9b0b86040518163ffffffff1660e01b815260040160206040518083038186803b15801561054e57600080fd5b505afa158015610562573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105869190610a1d565b825490915081116105f15760018201805460ff191660029081179091556040516001600160401b0388169160ff808b1692908c16917ff3e1949d167ac18929f02f5eb320d0379bac1f3196f1bf0ce9ececdd44b6cf45916105e8918b91610cf5565b60405180910390a45b5050505050505050565b6003602052600090815260409020546001600160401b031681565b6001546001600160a01b0316331461062d57600080fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600060606004600061066187876107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b038716825290925281206001015460ff1692506004906106a687876107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b03871682528352818120805483518186028101860190945280845292939092919084015b8282101561079a5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156107865780601f1061075b57610100808354040283529160200191610786565b820191906000526020600020905b81548152906001019060200180831161076957829003601f168201915b5050505050815260200190600101906106ef565b505050509050935093915050565b61ff00600883901b1660ff8216175b92915050565b6001546001600160a01b031681565b6000600260006107dc87876107a8565b6001600160481b03168152602080820192909252604090810160009081206001600160401b038716825283528181206001600160a01b038616825290925290205460ff169050949350505050565b6001546001600160a01b0316331461084157600080fd5b806003600061085086866107a8565b6001600160481b031681526020810191909152604001600020805467ffffffffffffffff19166001600160401b0392909216919091179055505050565b60006003600061089d85856107a8565b6001600160481b031681526020810191909152604001600020546001600160401b03169392505050565b6001546001600160a01b031633146108de57600080fd5b6001600160a01b0381161561090957600180546001600160a01b0319166001600160a01b0383161790555b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061094d5782800160ff1982351617855561097a565b8280016001018555821561097a579182015b8281111561097a57823582559160200191906001019061095f565b5061098692915061098a565b5090565b6109a491905b808211156109865760008155600101610990565b90565b80356001600160401b03811681146107b757600080fd5b803560ff811681146107b757600080fd5b6000602082840312156109e0578081fd5b81356001600160a01b03811681146109f6578182fd5b9392505050565b600060208284031215610a0e578081fd5b815180151581146109f6578182fd5b600060208284031215610a2e578081fd5b5051919050565b600060208284031215610a46578081fd5b81356109f681610e60565b60008060408385031215610a63578081fd5b8235610a6e81610e60565b91506020830135610a7e81610e4b565b809150509250929050565b60008060408385031215610a9b578182fd5b610aa584846109be565b9150610ab484602085016109be565b90509250929050565b600080600060608486031215610ad1578081fd5b610adb85856109be565b9250610aea85602086016109be565b9150610af985604086016109a7565b90509250925092565b60008060008060808587031215610b17578081fd5b610b2186866109be565b9350610b3086602087016109be565b92506040850135610b4081610e4b565b915060608501356001600160a01b0381168114610b5b578182fd5b939692955090935050565b60008060008060008060a08789031215610b7e578182fd5b610b8888886109be565b9550610b9788602089016109be565b9450610ba688604089016109a7565b9350606087013560038110610bb9578283fd5b925060808701356001600160401b0380821115610bd4578384fd5b8189018a601f820112610be5578485fd5b8035925081831115610bf5578485fd5b8a6020848301011115610c06578485fd5b979a9699509497509295602001949293505050565b6001600160a01b0391909116815260200190565b901515815260200190565b60208101610c4783610e41565b91905290565b600060408201610c5c85610e41565b848352602060408185015281855180845260608601915060608382028701019350828701855b82811015610ce757878603605f1901845281518051808852885b81811015610cb7578281018801518982018901528701610c9c565b81811115610cc7578988838b0101525b50601f01601f191696909601850195509284019290840190600101610c82565b509398975050505050505050565b60408101610d0284610e41565b838252610d0e83610e41565b8260208301529392505050565b60208082526010908201526f6f6c64206465706f7369744e6f6e636560801b604082015260600190565b6020808252600b908201526a3737ba103932b630bcb2b960a91b604082015260600190565b602080825260149082015273696e76616c6964207472616e736665725479706560601b604082015260600190565b6020808252602a908201527f70726f706f73616c20616c7265616479207061737365642f65786563757465646040820152690bd8d85b98d95b1b195960b21b606082015260800190565b6020808252601b908201527f72656c6179657220616c726561647920636f6d6d6974207369676e0000000000604082015260600190565b6001600160401b0391909116815260200190565b6001600160481b0391909116815260200190565b6003811061090957fe5b6001600160401b038116811461090957600080fd5b6001600160481b038116811461090957600080fdfea2646970667358221220f3d5f4de1647187d6287f960ab0af17eed769f371462d28642010afb08daf83164736f6c63430006040033"

// DeploySignatureCollector deploys a new Ethereum contract, binding an instance of SignatureCollector to it.
func DeploySignatureCollector(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAccessAddr common.Address) (common.Address, *types.Transaction, *SignatureCollector, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureCollectorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignatureCollectorBin), backend, bridgeAccessAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignatureCollector{SignatureCollectorCaller: SignatureCollectorCaller{contract: contract}, SignatureCollectorTransactor: SignatureCollectorTransactor{contract: contract}, SignatureCollectorFilterer: SignatureCollectorFilterer{contract: contract}}, nil
}

// SignatureCollector is an auto generated Go binding around an Ethereum contract.
type SignatureCollector struct {
	SignatureCollectorCaller     // Read-only binding to the contract
	SignatureCollectorTransactor // Write-only binding to the contract
	SignatureCollectorFilterer   // Log filterer for contract events
}

// SignatureCollectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureCollectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureCollectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureCollectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureCollectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureCollectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureCollectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureCollectorSession struct {
	Contract     *SignatureCollector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SignatureCollectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureCollectorCallerSession struct {
	Contract *SignatureCollectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SignatureCollectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureCollectorTransactorSession struct {
	Contract     *SignatureCollectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SignatureCollectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureCollectorRaw struct {
	Contract *SignatureCollector // Generic contract binding to access the raw methods on
}

// SignatureCollectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureCollectorCallerRaw struct {
	Contract *SignatureCollectorCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureCollectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureCollectorTransactorRaw struct {
	Contract *SignatureCollectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureCollector creates a new instance of SignatureCollector, bound to a specific deployed contract.
func NewSignatureCollector(address common.Address, backend bind.ContractBackend) (*SignatureCollector, error) {
	contract, err := bindSignatureCollector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureCollector{SignatureCollectorCaller: SignatureCollectorCaller{contract: contract}, SignatureCollectorTransactor: SignatureCollectorTransactor{contract: contract}, SignatureCollectorFilterer: SignatureCollectorFilterer{contract: contract}}, nil
}

// NewSignatureCollectorCaller creates a new read-only instance of SignatureCollector, bound to a specific deployed contract.
func NewSignatureCollectorCaller(address common.Address, caller bind.ContractCaller) (*SignatureCollectorCaller, error) {
	contract, err := bindSignatureCollector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureCollectorCaller{contract: contract}, nil
}

// NewSignatureCollectorTransactor creates a new write-only instance of SignatureCollector, bound to a specific deployed contract.
func NewSignatureCollectorTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureCollectorTransactor, error) {
	contract, err := bindSignatureCollector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureCollectorTransactor{contract: contract}, nil
}

// NewSignatureCollectorFilterer creates a new log filterer instance of SignatureCollector, bound to a specific deployed contract.
func NewSignatureCollectorFilterer(address common.Address, filterer bind.ContractFilterer) (*SignatureCollectorFilterer, error) {
	contract, err := bindSignatureCollector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureCollectorFilterer{contract: contract}, nil
}

// bindSignatureCollector binds a generic wrapper to an already deployed contract.
func bindSignatureCollector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureCollectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureCollector *SignatureCollectorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignatureCollector.Contract.SignatureCollectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureCollector *SignatureCollectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SignatureCollectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureCollector *SignatureCollectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SignatureCollectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureCollector *SignatureCollectorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignatureCollector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureCollector *SignatureCollectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureCollector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureCollector *SignatureCollectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureCollector.Contract.contract.Transact(opts, method, params...)
}

// DestinationStartNonce is a free data retrieval call binding the contract method 0x602b5d7e.
//
// Solidity: function _destinationStartNonce(uint72 ) view returns(uint64)
func (_SignatureCollector *SignatureCollectorCaller) DestinationStartNonce(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "_destinationStartNonce", arg0)
	return *ret0, err
}

// DestinationStartNonce is a free data retrieval call binding the contract method 0x602b5d7e.
//
// Solidity: function _destinationStartNonce(uint72 ) view returns(uint64)
func (_SignatureCollector *SignatureCollectorSession) DestinationStartNonce(arg0 *big.Int) (uint64, error) {
	return _SignatureCollector.Contract.DestinationStartNonce(&_SignatureCollector.CallOpts, arg0)
}

// DestinationStartNonce is a free data retrieval call binding the contract method 0x602b5d7e.
//
// Solidity: function _destinationStartNonce(uint72 ) view returns(uint64)
func (_SignatureCollector *SignatureCollectorCallerSession) DestinationStartNonce(arg0 *big.Int) (uint64, error) {
	return _SignatureCollector.Contract.DestinationStartNonce(&_SignatureCollector.CallOpts, arg0)
}

// BridgeAccess is a free data retrieval call binding the contract method 0x1b350ac1.
//
// Solidity: function bridgeAccess() view returns(address)
func (_SignatureCollector *SignatureCollectorCaller) BridgeAccess(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "bridgeAccess")
	return *ret0, err
}

// BridgeAccess is a free data retrieval call binding the contract method 0x1b350ac1.
//
// Solidity: function bridgeAccess() view returns(address)
func (_SignatureCollector *SignatureCollectorSession) BridgeAccess() (common.ContractAddress, error) {
	return _SignatureCollector.Contract.BridgeAccess(&_SignatureCollector.CallOpts)
}

// BridgeAccess is a free data retrieval call binding the contract method 0x1b350ac1.
//
// Solidity: function bridgeAccess() view returns(address)
func (_SignatureCollector *SignatureCollectorCallerSession) BridgeAccess() (common.ContractAddress, error) {
	return _SignatureCollector.Contract.BridgeAccess(&_SignatureCollector.CallOpts)
}

// CoId is a free data retrieval call binding the contract method 0x8b98dbad.
//
// Solidity: function coId(uint8 originChainId, uint8 destinationId) pure returns(uint72)
func (_SignatureCollector *SignatureCollectorCaller) CoId(opts *bind.CallOpts, originChainId uint8, destinationId uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "coId", originChainId, destinationId)
	return *ret0, err
}

// CoId is a free data retrieval call binding the contract method 0x8b98dbad.
//
// Solidity: function coId(uint8 originChainId, uint8 destinationId) pure returns(uint72)
func (_SignatureCollector *SignatureCollectorSession) CoId(originChainId uint8, destinationId uint8) (*big.Int, error) {
	return _SignatureCollector.Contract.CoId(&_SignatureCollector.CallOpts, originChainId, destinationId)
}

// CoId is a free data retrieval call binding the contract method 0x8b98dbad.
//
// Solidity: function coId(uint8 originChainId, uint8 destinationId) pure returns(uint72)
func (_SignatureCollector *SignatureCollectorCallerSession) CoId(originChainId uint8, destinationId uint8) (*big.Int, error) {
	return _SignatureCollector.Contract.CoId(&_SignatureCollector.CallOpts, originChainId, destinationId)
}

// DestinationProposals is a free data retrieval call binding the contract method 0x3ee27649.
//
// Solidity: function destinationProposals(uint72 , uint64 ) view returns(uint8 status)
func (_SignatureCollector *SignatureCollectorCaller) DestinationProposals(opts *bind.CallOpts, arg0 *big.Int, arg1 uint64) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "destinationProposals", arg0, arg1)
	return *ret0, err
}

// DestinationProposals is a free data retrieval call binding the contract method 0x3ee27649.
//
// Solidity: function destinationProposals(uint72 , uint64 ) view returns(uint8 status)
func (_SignatureCollector *SignatureCollectorSession) DestinationProposals(arg0 *big.Int, arg1 uint64) (uint8, error) {
	return _SignatureCollector.Contract.DestinationProposals(&_SignatureCollector.CallOpts, arg0, arg1)
}

// DestinationProposals is a free data retrieval call binding the contract method 0x3ee27649.
//
// Solidity: function destinationProposals(uint72 , uint64 ) view returns(uint8 status)
func (_SignatureCollector *SignatureCollectorCallerSession) DestinationProposals(arg0 *big.Int, arg1 uint64) (uint8, error) {
	return _SignatureCollector.Contract.DestinationProposals(&_SignatureCollector.CallOpts, arg0, arg1)
}

// GetDestinationStartNoncee is a free data retrieval call binding the contract method 0xea1d71c1.
//
// Solidity: function getDestinationStartNoncee(uint8 srcId, uint8 destId) view returns(uint64)
func (_SignatureCollector *SignatureCollectorCaller) GetDestinationStartNoncee(opts *bind.CallOpts, srcId uint8, destId uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "getDestinationStartNoncee", srcId, destId)
	return *ret0, err
}

// GetDestinationStartNoncee is a free data retrieval call binding the contract method 0xea1d71c1.
//
// Solidity: function getDestinationStartNoncee(uint8 srcId, uint8 destId) view returns(uint64)
func (_SignatureCollector *SignatureCollectorSession) GetDestinationStartNoncee(srcId uint8, destId uint8) (uint64, error) {
	return _SignatureCollector.Contract.GetDestinationStartNoncee(&_SignatureCollector.CallOpts, srcId, destId)
}

// GetDestinationStartNoncee is a free data retrieval call binding the contract method 0xea1d71c1.
//
// Solidity: function getDestinationStartNoncee(uint8 srcId, uint8 destId) view returns(uint64)
func (_SignatureCollector *SignatureCollectorCallerSession) GetDestinationStartNoncee(srcId uint8, destId uint8) (uint64, error) {
	return _SignatureCollector.Contract.GetDestinationStartNoncee(&_SignatureCollector.CallOpts, srcId, destId)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x80c22a24.
//
// Solidity: function getProposalSignatures(uint8 srcId, uint8 destId, uint64 depositNonce) view returns(uint8 status, bytes[] signatures)
func (_SignatureCollector *SignatureCollectorCaller) GetProposalSignatures(opts *bind.CallOpts, srcId uint8, destId uint8, depositNonce uint64) (struct {
	Status     uint8
	Signatures [][]byte
}, error) {
	ret := new(struct {
		Status     uint8
		Signatures [][]byte
	})
	out := ret
	err := _SignatureCollector.contract.Call(opts, out, "getProposalSignatures", srcId, destId, depositNonce)
	return *ret, err
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x80c22a24.
//
// Solidity: function getProposalSignatures(uint8 srcId, uint8 destId, uint64 depositNonce) view returns(uint8 status, bytes[] signatures)
func (_SignatureCollector *SignatureCollectorSession) GetProposalSignatures(srcId uint8, destId uint8, depositNonce uint64) (struct {
	Status     uint8
	Signatures [][]byte
}, error) {
	return _SignatureCollector.Contract.GetProposalSignatures(&_SignatureCollector.CallOpts, srcId, destId, depositNonce)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x80c22a24.
//
// Solidity: function getProposalSignatures(uint8 srcId, uint8 destId, uint64 depositNonce) view returns(uint8 status, bytes[] signatures)
func (_SignatureCollector *SignatureCollectorCallerSession) GetProposalSignatures(srcId uint8, destId uint8, depositNonce uint64) (struct {
	Status     uint8
	Signatures [][]byte
}, error) {
	return _SignatureCollector.Contract.GetProposalSignatures(&_SignatureCollector.CallOpts, srcId, destId, depositNonce)
}

// HasSignedProposal is a free data retrieval call binding the contract method 0xaf4f6257.
//
// Solidity: function hasSignedProposal(uint8 srcId, uint8 destId, uint64 depositNonce, address sender) view returns(bool)
func (_SignatureCollector *SignatureCollectorCaller) HasSignedProposal(opts *bind.CallOpts, srcId uint8, destId uint8, depositNonce uint64, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "hasSignedProposal", srcId, destId, depositNonce, sender)
	return *ret0, err
}

// HasSignedProposal is a free data retrieval call binding the contract method 0xaf4f6257.
//
// Solidity: function hasSignedProposal(uint8 srcId, uint8 destId, uint64 depositNonce, address sender) view returns(bool)
func (_SignatureCollector *SignatureCollectorSession) HasSignedProposal(srcId uint8, destId uint8, depositNonce uint64, sender common.Address) (bool, error) {
	return _SignatureCollector.Contract.HasSignedProposal(&_SignatureCollector.CallOpts, srcId, destId, depositNonce, sender)
}

// HasSignedProposal is a free data retrieval call binding the contract method 0xaf4f6257.
//
// Solidity: function hasSignedProposal(uint8 srcId, uint8 destId, uint64 depositNonce, address sender) view returns(bool)
func (_SignatureCollector *SignatureCollectorCallerSession) HasSignedProposal(srcId uint8, destId uint8, depositNonce uint64, sender common.Address) (bool, error) {
	return _SignatureCollector.Contract.HasSignedProposal(&_SignatureCollector.CallOpts, srcId, destId, depositNonce, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SignatureCollector *SignatureCollectorCaller) Owner(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SignatureCollector *SignatureCollectorSession) Owner() (common.ContractAddress, error) {
	return _SignatureCollector.Contract.Owner(&_SignatureCollector.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SignatureCollector *SignatureCollectorCallerSession) Owner() (common.ContractAddress, error) {
	return _SignatureCollector.Contract.Owner(&_SignatureCollector.CallOpts)
}

// ChangeBridgeAccess is a paid mutator transaction binding the contract method 0x7de0ac85.
//
// Solidity: function changeBridgeAccess(address bridgeAccessAddr) returns()
func (_SignatureCollector *SignatureCollectorTransactor) ChangeBridgeAccess(opts *bind.TransactOpts, bridgeAccessAddr common.Address) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "changeBridgeAccess", bridgeAccessAddr)
}

// ChangeBridgeAccess is a paid mutator transaction binding the contract method 0x7de0ac85.
//
// Solidity: function changeBridgeAccess(address bridgeAccessAddr) returns()
func (_SignatureCollector *SignatureCollectorSession) ChangeBridgeAccess(bridgeAccessAddr common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.ChangeBridgeAccess(&_SignatureCollector.TransactOpts, bridgeAccessAddr)
}

// ChangeBridgeAccess is a paid mutator transaction binding the contract method 0x7de0ac85.
//
// Solidity: function changeBridgeAccess(address bridgeAccessAddr) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) ChangeBridgeAccess(bridgeAccessAddr common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.ChangeBridgeAccess(&_SignatureCollector.TransactOpts, bridgeAccessAddr)
}

// SetDestinationStartNonce is a paid mutator transaction binding the contract method 0xd1b0628c.
//
// Solidity: function setDestinationStartNonce(uint8 srcId, uint8 destId, uint64 startNonce) returns()
func (_SignatureCollector *SignatureCollectorTransactor) SetDestinationStartNonce(opts *bind.TransactOpts, srcId uint8, destId uint8, startNonce uint64) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "setDestinationStartNonce", srcId, destId, startNonce)
}

// SetDestinationStartNonce is a paid mutator transaction binding the contract method 0xd1b0628c.
//
// Solidity: function setDestinationStartNonce(uint8 srcId, uint8 destId, uint64 startNonce) returns()
func (_SignatureCollector *SignatureCollectorSession) SetDestinationStartNonce(srcId uint8, destId uint8, startNonce uint64) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SetDestinationStartNonce(&_SignatureCollector.TransactOpts, srcId, destId, startNonce)
}

// SetDestinationStartNonce is a paid mutator transaction binding the contract method 0xd1b0628c.
//
// Solidity: function setDestinationStartNonce(uint8 srcId, uint8 destId, uint64 startNonce) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) SetDestinationStartNonce(srcId uint8, destId uint8, startNonce uint64) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SetDestinationStartNonce(&_SignatureCollector.TransactOpts, srcId, destId, startNonce)
}

// SignProposal is a paid mutator transaction binding the contract method 0x533703da.
//
// Solidity: function signProposal(uint8 srcId, uint8 destId, uint64 depositNonce, uint8 typ, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorTransactor) SignProposal(opts *bind.TransactOpts, srcId uint8, destId uint8, depositNonce uint64, typ uint8, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "signProposal", srcId, destId, depositNonce, typ, signature)
}

// SignProposal is a paid mutator transaction binding the contract method 0x533703da.
//
// Solidity: function signProposal(uint8 srcId, uint8 destId, uint64 depositNonce, uint8 typ, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorSession) SignProposal(srcId uint8, destId uint8, depositNonce uint64, typ uint8, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SignProposal(&_SignatureCollector.TransactOpts, srcId, destId, depositNonce, typ, signature)
}

// SignProposal is a paid mutator transaction binding the contract method 0x533703da.
//
// Solidity: function signProposal(uint8 srcId, uint8 destId, uint64 depositNonce, uint8 typ, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) SignProposal(srcId uint8, destId uint8, depositNonce uint64, typ uint8, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.Contract.SignProposal(&_SignatureCollector.TransactOpts, srcId, destId, depositNonce, typ, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SignatureCollector *SignatureCollectorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SignatureCollector *SignatureCollectorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.TransferOwnership(&_SignatureCollector.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.TransferOwnership(&_SignatureCollector.TransactOpts, newOwner)
}

// SignatureCollectorSignProposalEventIterator is returned from FilterSignProposalEvent and is used to iterate over the raw logs and unpacked data for SignProposalEvent events raised by the SignatureCollector contract.
type SignatureCollectorSignProposalEventIterator struct {
	Event *SignatureCollectorSignProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  sero.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SignatureCollectorSignProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureCollectorSignProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SignatureCollectorSignProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SignatureCollectorSignProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureCollectorSignProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureCollectorSignProposalEvent represents a SignProposalEvent event raised by the SignatureCollector contract.
type SignatureCollectorSignProposalEvent struct {
	ResourceChainId    uint8
	DestinationChainId uint8
	DepositNonce       uint64
	Typ                uint8
	Status             uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSignProposalEvent is a free log retrieval operation binding the contract event 0xf3e1949d167ac18929f02f5eb320d0379bac1f3196f1bf0ce9ececdd44b6cf45.
//
// Solidity: event SignProposalEvent(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, uint8 typ, uint8 status)
func (_SignatureCollector *SignatureCollectorFilterer) FilterSignProposalEvent(opts *bind.FilterOpts, resourceChainId []uint8, destinationChainId []uint8, depositNonce []uint64) (*SignatureCollectorSignProposalEventIterator, error) {

	var resourceChainIdRule []interface{}
	for _, resourceChainIdItem := range resourceChainId {
		resourceChainIdRule = append(resourceChainIdRule, resourceChainIdItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _SignatureCollector.contract.FilterLogs(opts, "SignProposalEvent", resourceChainIdRule, destinationChainIdRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &SignatureCollectorSignProposalEventIterator{contract: _SignatureCollector.contract, event: "SignProposalEvent", logs: logs, sub: sub}, nil
}

// WatchSignProposalEvent is a free log subscription operation binding the contract event 0xf3e1949d167ac18929f02f5eb320d0379bac1f3196f1bf0ce9ececdd44b6cf45.
//
// Solidity: event SignProposalEvent(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, uint8 typ, uint8 status)
func (_SignatureCollector *SignatureCollectorFilterer) WatchSignProposalEvent(opts *bind.WatchOpts, sink chan<- *SignatureCollectorSignProposalEvent, resourceChainId []uint8, destinationChainId []uint8, depositNonce []uint64) (event.Subscription, error) {

	var resourceChainIdRule []interface{}
	for _, resourceChainIdItem := range resourceChainId {
		resourceChainIdRule = append(resourceChainIdRule, resourceChainIdItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _SignatureCollector.contract.WatchLogs(opts, "SignProposalEvent", resourceChainIdRule, destinationChainIdRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureCollectorSignProposalEvent)
				if err := _SignatureCollector.contract.UnpackLog(event, "SignProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignProposalEvent is a log parse operation binding the contract event 0xf3e1949d167ac18929f02f5eb320d0379bac1f3196f1bf0ce9ececdd44b6cf45.
//
// Solidity: event SignProposalEvent(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, uint8 typ, uint8 status)
func (_SignatureCollector *SignatureCollectorFilterer) ParseSignProposalEvent(log types.Log) (*SignatureCollectorSignProposalEvent, error) {
	event := new(SignatureCollectorSignProposalEvent)
	if err := _SignatureCollector.contract.UnpackLog(event, "SignProposalEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
