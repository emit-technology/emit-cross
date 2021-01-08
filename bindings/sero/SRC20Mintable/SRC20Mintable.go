// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SRC20Mintable

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

// SRC20MintableABI is the input ABI used to generate the binding from.
const SRC20MintableABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dropMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SRC20MintableBin is the compiled bytecode used for deploying new contracts.
var SRC20MintableBin = "0x608060408190527f3be6bf24d822bcd6f6348f6f5a5c2d3108f04991ee63e80cde49a8c4746a0ef36000557fcf19eb4256453a4e30b6a06d651f1970c223fb6bd1826a28ed861f0e602db9b86001557f868bd6629e7c2e3d2ccf7b9968fad79b448e7a2bfb3ee20ed1acbc695c3c8b236002557f7c98e64bd943448b4e24ef8c2cdec7b8b1275970cfe10daf2a9bfa4b04dce9056003557fa6a366f1a72e1aef5d8d52ee240a476f619d15be7bc62d3df37496025b83459f6004557ff1964f6690a0536daa42e5c575091297d2479edcc96f721ad85b95358644d2766005557f9ab0d7c07029f006485cf3468ce7811aa8743b5a108599f6bec9367c50ac6aad6006557fa6cafc6282f61eff9032603a017e652f68410d3d3c69f0a3eeca8f181aec1d176007557f6800e94e36131c049eaeb631e4530829b0d3d20d5b637c8015a8dc9cedd70aed600855620011b038819003908190833981810160405260608110156200016c57600080fd5b81019080805160405193929190846401000000008211156200018d57600080fd5b908301906020820185811115620001a357600080fd5b8251640100000000811182820188101715620001be57600080fd5b82525081516020918201929091019080838360005b83811015620001ed578181015183820152602001620001d3565b50505050905090810190601f1680156200021b5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200023f57600080fd5b9083019060208201858111156200025557600080fd5b82516401000000008111828201881017156200027057600080fd5b82525081516020918201929091019080838360005b838110156200029f57818101518382015260200162000285565b50505050905090810190601f168015620002cd5780820380516001836020036101000a031916815260200191505b506040908152602091820151600980546001600160a01b031916339081179091556000908152600a845291909120805460ff1916600117905585519093506200031d9250600b918601906200043f565b5081516200033390600c9060208501906200043f565b50600d805460ff191660ff8316179055600e54600c805460408051602060026101006001861615026000190190941693909304601f8101849004840282018401909252818152620003ee94939092909190830182828015620003d95780601f10620003ad57610100808354040283529160200191620003d9565b820191906000526020600020905b815481529060010190602001808311620003bb57829003601f168201915b50506001600160e01b03620004011692505050565b620003f857600080fd5b505050620004e4565b6040805181815260608181018352600092909190602082018180368337019050509050828152836020820152600054604082a1602001519392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200048257805160ff1916838001178555620004b2565b82800160010185558215620004b2579182015b82811115620004b257825182559160200191906001019062000495565b50620004c0929150620004c4565b5090565b620004e191905b80821115620004c05760008155600101620004cb565b90565b610cbc80620004f46000396000f3fe6080604052600436106100a75760003560e01c80639201de55116100645780639201de55146101f657806395d89b4114610220578063983b2d5614610235578063a0712d6814610268578063d4ff21fe14610292578063f2fde38b146102a7576100a7565b806306fdde03146100ac57806318160ddd1461013657806323338b881461015d578063313ce5671461019257806344df8e70146101bd5780638da5cb5b146101c5575b600080fd5b3480156100b857600080fd5b506100c16102da565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100fb5781810151838201526020016100e3565b50505050905090810190601f1680156101285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014257600080fd5b5061014b610370565b60408051918252519081900360200190f35b34801561016957600080fd5b506101906004803603602081101561018057600080fd5b50356001600160a01b0316610376565b005b34801561019e57600080fd5b506101a76103ae565b6040805160ff9092168252519081900360200190f35b6101906103b7565b3480156101d157600080fd5b506101da6104ba565b604080516001600160a01b039092168252519081900360200190f35b34801561020257600080fd5b506100c16004803603602081101561021957600080fd5b50356104c9565b34801561022c57600080fd5b506100c16105df565b34801561024157600080fd5b506101906004803603602081101561025857600080fd5b50356001600160a01b0316610640565b34801561027457600080fd5b506101906004803603602081101561028b57600080fd5b503561067b565b34801561029e57600080fd5b506101906108e6565b3480156102b357600080fd5b50610190600480360360208110156102ca57600080fd5b50356001600160a01b0316610956565b600b8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103665780601f1061033b57610100808354040283529160200191610366565b820191906000526020600020905b81548152906001019060200180831161034957829003601f168201915b5050505050905090565b600e5490565b6009546001600160a01b0316331461038d57600080fd5b6001600160a01b03166000908152600a60205260409020805460ff19169055565b600d5460ff1690565b60606103c161099b565b600c8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815293945061045d938593909290918301828280156104535780601f1061042857610100808354040283529160200191610453565b820191906000526020600020905b81548152906001019060200180831161043657829003601f168201915b50505050506109db565b6104a1576040805162461bcd60e51b815260206004820152601060248201526f6e6f7420737570706f7274656420637960801b604482015290519081900360640190fd5b600e546104b4903463ffffffff610a5916565b600e5550565b6009546001600160a01b031681565b604080516020808252818301909252606091600091839160208201818036833701905050905060005b6020811015610558576008810260020a85026001600160f81b0319811615610543578083858151811061052157fe5b60200101906001600160f81b031916908160001a90535060019093019261054f565b831561054f5750610558565b506001016104f2565b506060826040519080825280601f01601f191660200182016040528015610586576020820181803683370190505b50905060005b838110156105d6578281815181106105a057fe5b602001015160f81c60f81b8282815181106105b757fe5b60200101906001600160f81b031916908160001a90535060010161058c565b50949350505050565b600c8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103665780601f1061033b57610100808354040283529160200191610366565b6009546001600160a01b0316331461065757600080fd5b6001600160a01b03166000908152600a60205260409020805460ff19166001179055565b336000908152600a602052604090205460ff166106ce576040805162461bcd60e51b815260206004820152600c60248201526b1b9bdd08185c1c1c9bdd995960a21b604482015290519081900360640190fd5b600c805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152600093610767939192909183018282801561075d5780601f106107325761010080835404028352916020019161075d565b820191906000526020600020905b81548152906001019060200180831161074057829003601f168201915b5050505050610aa2565b90508181101561082a576000610783838363ffffffff610a5916565b600c8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815293945061081f938593909290918301828280156108155780601f106107ea57610100808354040283529160200191610815565b820191906000526020600020905b8154815290600101906020018083116107f857829003601f168201915b5050505050610ad8565b61082857600080fd5b505b600e5461083d908363ffffffff610b1616565b600e55600c805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181526108d9933393919290918301828280156108ce5780601f106108a3576101008083540402835291602001916108ce565b820191906000526020600020905b8154815290600101906020018083116108b157829003601f168201915b505050505084610b70565b6108e257600080fd5b5050565b336000908152600a602052604090205460ff1661093d576040805162461bcd60e51b815260206004820152601060248201526f63616e206e6f74206d696e7461626c6560801b604482015290519081900360640190fd5b336000908152600a60205260409020805460ff19169055565b6009546001600160a01b0316331461096d57600080fd5b6001600160a01b0381161561099857600980546001600160a01b0319166001600160a01b0383161790555b50565b6040805160208082528183019092526060918291906020820181803683370190505090506000600354602083a15080516109d4816104c9565b9250505090565b600081518351146109ee57506000610a53565b60005b8351811015610a4d57828181518110610a0657fe5b602001015160f81c60f81b6001600160f81b031916848281518110610a2757fe5b01602001516001600160f81b03191614610a45576000915050610a53565b6001016109f1565b50600190505b92915050565b6000610a9b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610b99565b9392505050565b60408051602080825281830190925260009160609190602082018180368337019050509050828152600154602082a15192915050565b6040805181815260608181018352600092909190602082018180368337019050509050828152836020820152600054604082a1602001519392505050565b600082820183811015610a9b576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000610b91848484604051806020016040528060008152506000801b610c30565b949350505050565b60008184841115610c285760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610bed578181015183820152602001610bd5565b50505050905090810190601f168015610c1a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6040805160a080825260c082019092526000916060919060208201818036833701905050905086815285602082015284604082015283606082015282608082015260025460a082a160800151969550505050505056fea264697066735822122056c55a4434c6d365b54e2ddc39b0ae77c4a2cc05858552f23705f08649f4dc6564736f6c63430006040033"

// DeploySRC20Mintable deploys a new Ethereum contract, binding an instance of SRC20Mintable to it.
func DeploySRC20Mintable(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *SRC20Mintable, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC20MintableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SRC20MintableBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SRC20Mintable{SRC20MintableCaller: SRC20MintableCaller{contract: contract}, SRC20MintableTransactor: SRC20MintableTransactor{contract: contract}, SRC20MintableFilterer: SRC20MintableFilterer{contract: contract}}, nil
}

// SRC20Mintable is an auto generated Go binding around an Ethereum contract.
type SRC20Mintable struct {
	SRC20MintableCaller     // Read-only binding to the contract
	SRC20MintableTransactor // Write-only binding to the contract
	SRC20MintableFilterer   // Log filterer for contract events
}

// SRC20MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SRC20MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SRC20MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SRC20MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SRC20MintableSession struct {
	Contract     *SRC20Mintable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SRC20MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SRC20MintableCallerSession struct {
	Contract *SRC20MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SRC20MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SRC20MintableTransactorSession struct {
	Contract     *SRC20MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SRC20MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SRC20MintableRaw struct {
	Contract *SRC20Mintable // Generic contract binding to access the raw methods on
}

// SRC20MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SRC20MintableCallerRaw struct {
	Contract *SRC20MintableCaller // Generic read-only contract binding to access the raw methods on
}

// SRC20MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SRC20MintableTransactorRaw struct {
	Contract *SRC20MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSRC20Mintable creates a new instance of SRC20Mintable, bound to a specific deployed contract.
func NewSRC20Mintable(address common.Address, backend bind.ContractBackend) (*SRC20Mintable, error) {
	contract, err := bindSRC20Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SRC20Mintable{SRC20MintableCaller: SRC20MintableCaller{contract: contract}, SRC20MintableTransactor: SRC20MintableTransactor{contract: contract}, SRC20MintableFilterer: SRC20MintableFilterer{contract: contract}}, nil
}

// NewSRC20MintableCaller creates a new read-only instance of SRC20Mintable, bound to a specific deployed contract.
func NewSRC20MintableCaller(address common.Address, caller bind.ContractCaller) (*SRC20MintableCaller, error) {
	contract, err := bindSRC20Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SRC20MintableCaller{contract: contract}, nil
}

// NewSRC20MintableTransactor creates a new write-only instance of SRC20Mintable, bound to a specific deployed contract.
func NewSRC20MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*SRC20MintableTransactor, error) {
	contract, err := bindSRC20Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SRC20MintableTransactor{contract: contract}, nil
}

// NewSRC20MintableFilterer creates a new log filterer instance of SRC20Mintable, bound to a specific deployed contract.
func NewSRC20MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*SRC20MintableFilterer, error) {
	contract, err := bindSRC20Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SRC20MintableFilterer{contract: contract}, nil
}

// bindSRC20Mintable binds a generic wrapper to an already deployed contract.
func bindSRC20Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC20MintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC20Mintable *SRC20MintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC20Mintable.Contract.SRC20MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC20Mintable *SRC20MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.SRC20MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC20Mintable *SRC20MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.SRC20MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC20Mintable *SRC20MintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC20Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC20Mintable *SRC20MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC20Mintable *SRC20MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Mintable *SRC20MintableCaller) Bytes32ToString(opts *bind.CallOpts, x [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "bytes32ToString", x)
	return *ret0, err
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Mintable *SRC20MintableSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC20Mintable.Contract.Bytes32ToString(&_SRC20Mintable.CallOpts, x)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Mintable *SRC20MintableCallerSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC20Mintable.Contract.Bytes32ToString(&_SRC20Mintable.CallOpts, x)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SRC20Mintable *SRC20MintableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SRC20Mintable *SRC20MintableSession) Decimals() (uint8, error) {
	return _SRC20Mintable.Contract.Decimals(&_SRC20Mintable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SRC20Mintable *SRC20MintableCallerSession) Decimals() (uint8, error) {
	return _SRC20Mintable.Contract.Decimals(&_SRC20Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SRC20Mintable *SRC20MintableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SRC20Mintable *SRC20MintableSession) Name() (string, error) {
	return _SRC20Mintable.Contract.Name(&_SRC20Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SRC20Mintable *SRC20MintableCallerSession) Name() (string, error) {
	return _SRC20Mintable.Contract.Name(&_SRC20Mintable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SRC20Mintable *SRC20MintableCaller) Owner(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SRC20Mintable *SRC20MintableSession) Owner() (common.ContractAddress, error) {
	return _SRC20Mintable.Contract.Owner(&_SRC20Mintable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SRC20Mintable *SRC20MintableCallerSession) Owner() (common.ContractAddress, error) {
	return _SRC20Mintable.Contract.Owner(&_SRC20Mintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SRC20Mintable *SRC20MintableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SRC20Mintable *SRC20MintableSession) Symbol() (string, error) {
	return _SRC20Mintable.Contract.Symbol(&_SRC20Mintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SRC20Mintable *SRC20MintableCallerSession) Symbol() (string, error) {
	return _SRC20Mintable.Contract.Symbol(&_SRC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SRC20Mintable *SRC20MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SRC20Mintable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SRC20Mintable *SRC20MintableSession) TotalSupply() (*big.Int, error) {
	return _SRC20Mintable.Contract.TotalSupply(&_SRC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SRC20Mintable *SRC20MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _SRC20Mintable.Contract.TotalSupply(&_SRC20Mintable.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableTransactor) AddMinter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "addMinter", user)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableSession) AddMinter(user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.AddMinter(&_SRC20Mintable.TransactOpts, user)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) AddMinter(user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.AddMinter(&_SRC20Mintable.TransactOpts, user)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_SRC20Mintable *SRC20MintableTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_SRC20Mintable *SRC20MintableSession) Burn() (*types.Transaction, error) {
	return _SRC20Mintable.Contract.Burn(&_SRC20Mintable.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() payable returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) Burn() (*types.Transaction, error) {
	return _SRC20Mintable.Contract.Burn(&_SRC20Mintable.TransactOpts)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableTransactor) DelMinter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "delMinter", user)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableSession) DelMinter(user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.DelMinter(&_SRC20Mintable.TransactOpts, user)
}

// DelMinter is a paid mutator transaction binding the contract method 0x23338b88.
//
// Solidity: function delMinter(address user) returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) DelMinter(user common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.DelMinter(&_SRC20Mintable.TransactOpts, user)
}

// DropMintable is a paid mutator transaction binding the contract method 0xd4ff21fe.
//
// Solidity: function dropMintable() returns()
func (_SRC20Mintable *SRC20MintableTransactor) DropMintable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "dropMintable")
}

// DropMintable is a paid mutator transaction binding the contract method 0xd4ff21fe.
//
// Solidity: function dropMintable() returns()
func (_SRC20Mintable *SRC20MintableSession) DropMintable() (*types.Transaction, error) {
	return _SRC20Mintable.Contract.DropMintable(&_SRC20Mintable.TransactOpts)
}

// DropMintable is a paid mutator transaction binding the contract method 0xd4ff21fe.
//
// Solidity: function dropMintable() returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) DropMintable() (*types.Transaction, error) {
	return _SRC20Mintable.Contract.DropMintable(&_SRC20Mintable.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 total) returns()
func (_SRC20Mintable *SRC20MintableTransactor) Mint(opts *bind.TransactOpts, total *big.Int) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "mint", total)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 total) returns()
func (_SRC20Mintable *SRC20MintableSession) Mint(total *big.Int) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.Mint(&_SRC20Mintable.TransactOpts, total)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 total) returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) Mint(total *big.Int) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.Mint(&_SRC20Mintable.TransactOpts, total)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SRC20Mintable *SRC20MintableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SRC20Mintable *SRC20MintableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.TransferOwnership(&_SRC20Mintable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SRC20Mintable *SRC20MintableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SRC20Mintable.Contract.TransferOwnership(&_SRC20Mintable.TransactOpts, newOwner)
}
