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
const SignatureCollectorABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"relayerThreshold_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"resourceChainId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"destinationChainId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"SignProposal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRelayerThreshold_\",\"type\":\"uint256\"}],\"name\":\"changeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"sourceChainId\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"commitSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasSignedProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isRalyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"soruce\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"destination\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"signatureCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SignatureCollectorBin is the compiled bytecode used for deploying new contracts.
var SignatureCollectorBin = "0x608060405234801561001057600080fd5b50604051610ae1380380610ae183398101604081905261002f9161010d565b600080546001600160a01b031916331781555b82518110156100755761006d83828151811061005a57fe5b602002602001015161007f60201b60201c565b600101610042565b5060035550610210565b6001600160a01b03811660009081526001602052604090205460ff16156100c15760405162461bcd60e51b81526004016100b8906101bf565b60405180910390fd5b6001600160a01b03166000908152600160208190526040909120805460ff191682179055600280549091019055565b80516001600160a01b038116811461010757600080fd5b92915050565b6000806040838503121561011f578182fd5b82516001600160401b0380821115610135578384fd5b81850186601f820112610146578485fd5b8051925081831115610156578485fd5b602091508183026101688382016101ea565b8481528381019083850183850186018b1015610182578889fd5b8894505b868510156101ac576101988b826100f0565b835260019490940193918501918501610186565b5097909301519698969750505050505050565b60208082526011908201527030b63932b0b23c9030903932b630bcb2b960791b604082015260600190565b6040518181016001600160401b038111828210171561020857600080fd5b604052919050565b6108c28061021f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063abd3636211610071578063abd3636214610129578063ac57fc4b1461014d578063dd39f00d14610160578063e85b479014610173578063f0e9b0b814610186578063f2fde38b1461018e576100a9565b80631c99fd03146100ae5780632bceaea6146100c3578063862159ab146100d65780638da5cb5b146100f45780639f856e8714610109575b600080fd5b6100c16100bc366004610625565b6101a1565b005b6100c16100d136600461057d565b610339565b6100de610355565b6040516100eb9190610801565b60405180910390f35b6100fc61035b565b6040516100eb91906106e9565b61011c61011736600461055b565b61036a565b6040516100eb91906106fd565b61013c610137366004610595565b610388565b6040516100eb95949392919061080a565b6100c161015b36600461055b565b6103c7565b6100c161016e36600461055b565b610441565b61011c6101813660046105cd565b610464565b6100de61048a565b6100c161019c36600461055b565b610490565b3360009081526001602052604090205460ff166101d95760405162461bcd60e51b81526004016101d0906107dc565b60405180910390fd5b60ff808416600090815260046020908152604080832067ffffffffffffffff87168452825280832033845290915290205416156102285760405162461bcd60e51b81526004016101d09061075b565b60ff808416600090815260056020908152604080832067ffffffffffffffff871684529091529020600381015490911661029157805460ff8581166101000261ff001991881660ff19938416179190911617825560038201805490911660011790554360028201555b6001808201805490910190819055600354116102b75760038101805460ff191660021790555b60ff808516600081815260046020908152604080832067ffffffffffffffff891680855290835281842033855290925291829020805460ff19166001179055905190928816907f3783137594f14f0d0c60f4fc0f22bdf23861713bd61875f7d6bdef551afefee79061032a908790610708565b60405180910390a45050505050565b6000546001600160a01b0316331461035057600080fd5b600355565b60025481565b6000546001600160a01b031681565b6001600160a01b031660009081526001602052604090205460ff1690565b6005602090815260009283526040808420909152908252902080546001820154600283015460039093015460ff80841694610100909404811693911685565b6000546001600160a01b031633146103de57600080fd5b6001600160a01b03811660009081526001602052604090205460ff166104165760405162461bcd60e51b81526004016101d09061078a565b6001600160a01b03166000908152600160205260409020805460ff1916905560028054600019019055565b6000546001600160a01b0316331461045857600080fd5b610461816104d6565b50565b600460209081526000938452604080852082529284528284209052825290205460ff1681565b60035481565b6000546001600160a01b031633146104a757600080fd5b6001600160a01b0381161561046157600080546001600160a01b0383166001600160a01b031990911617905550565b6001600160a01b03811660009081526001602052604090205460ff161561050f5760405162461bcd60e51b81526004016101d0906107b1565b6001600160a01b03166000908152600160208190526040909120805460ff191682179055600280549091019055565b80356001600160a01b038116811461055557600080fd5b92915050565b60006020828403121561056c578081fd5b610576838361053e565b9392505050565b60006020828403121561058e578081fd5b5035919050565b600080604083850312156105a7578081fd5b82356105b28161087d565b915060208301356105c281610867565b809150509250929050565b6000806000606084860312156105e1578081fd5b833560ff811681146105f1578182fd5b9250602084013567ffffffffffffffff8116811461060d578182fd5b915061061c856040860161053e565b90509250925092565b6000806000806080858703121561063a578081fd5b84356106458161087d565b935060208501356106558161087d565b9250604085013561066581610867565b9150606085013567ffffffffffffffff80821115610681578283fd5b81870188601f820112610692578384fd5b80359250818311156106a2578384fd5b6106b5601f8401601f1916602001610834565b91508282528860208483010111156106cb578384fd5b6106dc83602084016020840161085b565b5094979396509194505050565b6001600160a01b0391909116815260200190565b901515815260200190565b6000602080835283518082850152825b8181101561073457858101830151858201604001528201610718565b818111156107455783604083870101525b50601f01601f1916929092016040019392505050565b6020808252601590820152741c995b185e595c88185b1c9958591e481d9bdd1959605a1b604082015260600190565b6020808252600d908201526c3737ba1030903932b630bcb2b960991b604082015260600190565b60208082526011908201527030b63932b0b23c9030903932b630bcb2b960791b604082015260600190565b6020808252600b908201526a3737ba103932b630bcb2b960a91b604082015260600190565b90815260200190565b60ff9586168152938516602085015260408401929092526060830152909116608082015260a00190565b60405181810167ffffffffffffffff8111828210171561085357600080fd5b604052919050565b82818337506000910152565b67ffffffffffffffff8116811461046157600080fd5b60ff8116811461046157600080fdfea2646970667358221220576410bbfecfc5f2f901ca8c056f538b8f13744dfbae70322f8d039084c325d464736f6c63430006040033"

// DeploySignatureCollector deploys a new Ethereum contract, binding an instance of SignatureCollector to it.
func DeploySignatureCollector(auth *bind.TransactOpts, backend bind.ContractBackend, initialRelayers []common.Address, relayerThreshold_ *big.Int) (common.Address, *types.Transaction, *SignatureCollector, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureCollectorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignatureCollectorBin), backend, initialRelayers, relayerThreshold_)
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

// HasSignedProposal is a free data retrieval call binding the contract method 0xe85b4790.
//
// Solidity: function hasSignedProposal(uint8 , uint64 , address ) view returns(bool)
func (_SignatureCollector *SignatureCollectorCaller) HasSignedProposal(opts *bind.CallOpts, arg0 uint8, arg1 uint64, arg2 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "hasSignedProposal", arg0, arg1, arg2)
	return *ret0, err
}

// HasSignedProposal is a free data retrieval call binding the contract method 0xe85b4790.
//
// Solidity: function hasSignedProposal(uint8 , uint64 , address ) view returns(bool)
func (_SignatureCollector *SignatureCollectorSession) HasSignedProposal(arg0 uint8, arg1 uint64, arg2 common.Address) (bool, error) {
	return _SignatureCollector.Contract.HasSignedProposal(&_SignatureCollector.CallOpts, arg0, arg1, arg2)
}

// HasSignedProposal is a free data retrieval call binding the contract method 0xe85b4790.
//
// Solidity: function hasSignedProposal(uint8 , uint64 , address ) view returns(bool)
func (_SignatureCollector *SignatureCollectorCallerSession) HasSignedProposal(arg0 uint8, arg1 uint64, arg2 common.Address) (bool, error) {
	return _SignatureCollector.Contract.HasSignedProposal(&_SignatureCollector.CallOpts, arg0, arg1, arg2)
}

// IsRalyer is a free data retrieval call binding the contract method 0x9f856e87.
//
// Solidity: function isRalyer(address user) view returns(bool)
func (_SignatureCollector *SignatureCollectorCaller) IsRalyer(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "isRalyer", user)
	return *ret0, err
}

// IsRalyer is a free data retrieval call binding the contract method 0x9f856e87.
//
// Solidity: function isRalyer(address user) view returns(bool)
func (_SignatureCollector *SignatureCollectorSession) IsRalyer(user common.Address) (bool, error) {
	return _SignatureCollector.Contract.IsRalyer(&_SignatureCollector.CallOpts, user)
}

// IsRalyer is a free data retrieval call binding the contract method 0x9f856e87.
//
// Solidity: function isRalyer(address user) view returns(bool)
func (_SignatureCollector *SignatureCollectorCallerSession) IsRalyer(user common.Address) (bool, error) {
	return _SignatureCollector.Contract.IsRalyer(&_SignatureCollector.CallOpts, user)
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

// Proposals is a free data retrieval call binding the contract method 0xabd36362.
//
// Solidity: function proposals(uint8 , uint64 ) view returns(uint8 soruce, uint8 destination, uint256 signatureCount, uint256 proposedBlock, uint8 status)
func (_SignatureCollector *SignatureCollectorCaller) Proposals(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	Soruce         uint8
	Destination    uint8
	SignatureCount *big.Int
	ProposedBlock  *big.Int
	Status         uint8
}, error) {
	ret := new(struct {
		Soruce         uint8
		Destination    uint8
		SignatureCount *big.Int
		ProposedBlock  *big.Int
		Status         uint8
	})
	out := ret
	err := _SignatureCollector.contract.Call(opts, out, "proposals", arg0, arg1)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0xabd36362.
//
// Solidity: function proposals(uint8 , uint64 ) view returns(uint8 soruce, uint8 destination, uint256 signatureCount, uint256 proposedBlock, uint8 status)
func (_SignatureCollector *SignatureCollectorSession) Proposals(arg0 uint8, arg1 uint64) (struct {
	Soruce         uint8
	Destination    uint8
	SignatureCount *big.Int
	ProposedBlock  *big.Int
	Status         uint8
}, error) {
	return _SignatureCollector.Contract.Proposals(&_SignatureCollector.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0xabd36362.
//
// Solidity: function proposals(uint8 , uint64 ) view returns(uint8 soruce, uint8 destination, uint256 signatureCount, uint256 proposedBlock, uint8 status)
func (_SignatureCollector *SignatureCollectorCallerSession) Proposals(arg0 uint8, arg1 uint64) (struct {
	Soruce         uint8
	Destination    uint8
	SignatureCount *big.Int
	ProposedBlock  *big.Int
	Status         uint8
}, error) {
	return _SignatureCollector.Contract.Proposals(&_SignatureCollector.CallOpts, arg0, arg1)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xf0e9b0b8.
//
// Solidity: function relayerThreshold() view returns(uint256)
func (_SignatureCollector *SignatureCollectorCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "relayerThreshold")
	return *ret0, err
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xf0e9b0b8.
//
// Solidity: function relayerThreshold() view returns(uint256)
func (_SignatureCollector *SignatureCollectorSession) RelayerThreshold() (*big.Int, error) {
	return _SignatureCollector.Contract.RelayerThreshold(&_SignatureCollector.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xf0e9b0b8.
//
// Solidity: function relayerThreshold() view returns(uint256)
func (_SignatureCollector *SignatureCollectorCallerSession) RelayerThreshold() (*big.Int, error) {
	return _SignatureCollector.Contract.RelayerThreshold(&_SignatureCollector.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x862159ab.
//
// Solidity: function totalRelayers() view returns(uint256)
func (_SignatureCollector *SignatureCollectorCaller) TotalRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignatureCollector.contract.Call(opts, out, "totalRelayers")
	return *ret0, err
}

// TotalRelayers is a free data retrieval call binding the contract method 0x862159ab.
//
// Solidity: function totalRelayers() view returns(uint256)
func (_SignatureCollector *SignatureCollectorSession) TotalRelayers() (*big.Int, error) {
	return _SignatureCollector.Contract.TotalRelayers(&_SignatureCollector.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x862159ab.
//
// Solidity: function totalRelayers() view returns(uint256)
func (_SignatureCollector *SignatureCollectorCallerSession) TotalRelayers() (*big.Int, error) {
	return _SignatureCollector.Contract.TotalRelayers(&_SignatureCollector.CallOpts)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorTransactor) AddRelayer(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "addRelayer", user)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorSession) AddRelayer(user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.AddRelayer(&_SignatureCollector.TransactOpts, user)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) AddRelayer(user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.AddRelayer(&_SignatureCollector.TransactOpts, user)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newRelayerThreshold_) returns()
func (_SignatureCollector *SignatureCollectorTransactor) ChangeRelayerThreshold(opts *bind.TransactOpts, newRelayerThreshold_ *big.Int) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "changeRelayerThreshold", newRelayerThreshold_)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newRelayerThreshold_) returns()
func (_SignatureCollector *SignatureCollectorSession) ChangeRelayerThreshold(newRelayerThreshold_ *big.Int) (*types.Transaction, error) {
	return _SignatureCollector.Contract.ChangeRelayerThreshold(&_SignatureCollector.TransactOpts, newRelayerThreshold_)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newRelayerThreshold_) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) ChangeRelayerThreshold(newRelayerThreshold_ *big.Int) (*types.Transaction, error) {
	return _SignatureCollector.Contract.ChangeRelayerThreshold(&_SignatureCollector.TransactOpts, newRelayerThreshold_)
}

// CommitSign is a paid mutator transaction binding the contract method 0x1c99fd03.
//
// Solidity: function commitSign(uint8 sourceChainId, uint8 destinationChainId, uint64 depositNonce, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorTransactor) CommitSign(opts *bind.TransactOpts, sourceChainId uint8, destinationChainId uint8, depositNonce uint64, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "commitSign", sourceChainId, destinationChainId, depositNonce, signature)
}

// CommitSign is a paid mutator transaction binding the contract method 0x1c99fd03.
//
// Solidity: function commitSign(uint8 sourceChainId, uint8 destinationChainId, uint64 depositNonce, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorSession) CommitSign(sourceChainId uint8, destinationChainId uint8, depositNonce uint64, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.Contract.CommitSign(&_SignatureCollector.TransactOpts, sourceChainId, destinationChainId, depositNonce, signature)
}

// CommitSign is a paid mutator transaction binding the contract method 0x1c99fd03.
//
// Solidity: function commitSign(uint8 sourceChainId, uint8 destinationChainId, uint64 depositNonce, bytes signature) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) CommitSign(sourceChainId uint8, destinationChainId uint8, depositNonce uint64, signature []byte) (*types.Transaction, error) {
	return _SignatureCollector.Contract.CommitSign(&_SignatureCollector.TransactOpts, sourceChainId, destinationChainId, depositNonce, signature)
}

// DelRelayer is a paid mutator transaction binding the contract method 0xac57fc4b.
//
// Solidity: function delRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorTransactor) DelRelayer(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.contract.Transact(opts, "delRelayer", user)
}

// DelRelayer is a paid mutator transaction binding the contract method 0xac57fc4b.
//
// Solidity: function delRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorSession) DelRelayer(user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.DelRelayer(&_SignatureCollector.TransactOpts, user)
}

// DelRelayer is a paid mutator transaction binding the contract method 0xac57fc4b.
//
// Solidity: function delRelayer(address user) returns()
func (_SignatureCollector *SignatureCollectorTransactorSession) DelRelayer(user common.Address) (*types.Transaction, error) {
	return _SignatureCollector.Contract.DelRelayer(&_SignatureCollector.TransactOpts, user)
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

// SignatureCollectorSignProposalIterator is returned from FilterSignProposal and is used to iterate over the raw logs and unpacked data for SignProposal events raised by the SignatureCollector contract.
type SignatureCollectorSignProposalIterator struct {
	Event *SignatureCollectorSignProposal // Event containing the contract specifics and raw log

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
func (it *SignatureCollectorSignProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureCollectorSignProposal)
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
		it.Event = new(SignatureCollectorSignProposal)
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
func (it *SignatureCollectorSignProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureCollectorSignProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureCollectorSignProposal represents a SignProposal event raised by the SignatureCollector contract.
type SignatureCollectorSignProposal struct {
	ResourceChainId    uint8
	DestinationChainId uint8
	DepositNonce       uint64
	Signature          []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSignProposal is a free log retrieval operation binding the contract event 0x3783137594f14f0d0c60f4fc0f22bdf23861713bd61875f7d6bdef551afefee7.
//
// Solidity: event SignProposal(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, bytes signature)
func (_SignatureCollector *SignatureCollectorFilterer) FilterSignProposal(opts *bind.FilterOpts, resourceChainId []uint8, destinationChainId []uint8, depositNonce []uint64) (*SignatureCollectorSignProposalIterator, error) {

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

	logs, sub, err := _SignatureCollector.contract.FilterLogs(opts, "SignProposal", resourceChainIdRule, destinationChainIdRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &SignatureCollectorSignProposalIterator{contract: _SignatureCollector.contract, event: "SignProposal", logs: logs, sub: sub}, nil
}

// WatchSignProposal is a free log subscription operation binding the contract event 0x3783137594f14f0d0c60f4fc0f22bdf23861713bd61875f7d6bdef551afefee7.
//
// Solidity: event SignProposal(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, bytes signature)
func (_SignatureCollector *SignatureCollectorFilterer) WatchSignProposal(opts *bind.WatchOpts, sink chan<- *SignatureCollectorSignProposal, resourceChainId []uint8, destinationChainId []uint8, depositNonce []uint64) (event.Subscription, error) {

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

	logs, sub, err := _SignatureCollector.contract.WatchLogs(opts, "SignProposal", resourceChainIdRule, destinationChainIdRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureCollectorSignProposal)
				if err := _SignatureCollector.contract.UnpackLog(event, "SignProposal", log); err != nil {
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

// ParseSignProposal is a log parse operation binding the contract event 0x3783137594f14f0d0c60f4fc0f22bdf23861713bd61875f7d6bdef551afefee7.
//
// Solidity: event SignProposal(uint8 indexed resourceChainId, uint8 indexed destinationChainId, uint64 indexed depositNonce, bytes signature)
func (_SignatureCollector *SignatureCollectorFilterer) ParseSignProposal(log types.Log) (*SignatureCollectorSignProposal, error) {
	event := new(SignatureCollectorSignProposal)
	if err := _SignatureCollector.contract.UnpackLog(event, "SignProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}
