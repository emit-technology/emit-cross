// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CrossFee

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

// CrossFeeABI is the input ABI used to generate the binding from.
const CrossFeeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uni_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"crossInGas_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"MAX_GAS_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossInGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"resourceToDefalutGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"resourceToFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"resourceToTokenAddrss\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uni_\",\"type\":\"address\"}],\"name\":\"setUni\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"defaltGasFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"crossInGas_\",\"type\":\"uint256\"}],\"name\":\"setCrossInGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"gasFee_\",\"type\":\"uint256\"}],\"name\":\"setDefaultGasFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"feeRate_\",\"type\":\"uint256\"}],\"name\":\"setCorssFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"calCrossFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"relayerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CrossFeeBin is the compiled bytecode used for deploying new contracts.
var CrossFeeBin = "0x60806040526412a05f20006003553480156200001a57600080fd5b5060405162001a8d38038062001a8d833981810160405281019062000040919062000140565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060048190555050505062000208565b6000815190506200012381620001d4565b92915050565b6000815190506200013a81620001ee565b92915050565b6000806000606084860312156200015657600080fd5b6000620001668682870162000112565b9350506020620001798682870162000112565b92505060406200018c8682870162000129565b9150509250925092565b6000620001a382620001aa565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b620001df8162000196565b8114620001eb57600080fd5b50565b620001f981620001ca565b81146200020557600080fd5b50565b61187580620002186000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063d2fa635e11610071578063d2fa635e146102b7578063d702ffd0146102d3578063e3bbb4f1146102f1578063eab44ff61461030f578063ec1bef0c1461033f5761010b565b80638da5cb5b1461021d57806392c212d81461023b578063a8e31a681461026b578063d090223e1461029b5761010b565b80633fc8cef3116100de5780633fc8cef3146101a95780635b769f3c146101c757806369213d03146101e35780636f3ca55d146102015761010b565b80630100a5551461011057806302e3588f146101405780631120e6551461015c578063171cb89614610178575b600080fd5b61012a60048036038101906101259190610ffb565b61035b565b6040516101379190611672565b60405180910390f35b61015a60048036038101906101559190611024565b610373565b005b61017660048036038101906101719190611161565b61048a565b005b610192600480360381019061018d91906110c3565b6104ed565b6040516101a092919061168d565b60405180910390f35b6101b1610507565b6040516101be9190611555565b60405180910390f35b6101e160048036038101906101dc9190610fd2565b61052d565b005b6101eb610600565b6040516101f89190611672565b60405180910390f35b61021b60048036038101906102169190610fd2565b610606565b005b6102256106d9565b6040516102329190611555565b60405180910390f35b61025560048036038101906102509190611087565b6106fe565b6040516102629190611672565b60405180910390f35b61028560048036038101906102809190610ffb565b610731565b6040516102929190611672565b60405180910390f35b6102b560048036038101906102b09190611087565b610749565b005b6102d160048036038101906102cc9190611161565b6107be565b005b6102db610821565b6040516102e89190611555565b60405180910390f35b6102f9610847565b6040516103069190611672565b60405180910390f35b61032960048036038101906103249190610ffb565b61084d565b6040516103369190611555565b60405180910390f35b61035960048036038101906103549190611087565b610880565b005b60066020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f990611652565b60405180910390fd5b826005600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660008681526020019081526020016000208190555081600760008681526020019081526020016000208190555050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104e357600080fd5b8060048190555050565b6000806104fb8585856108f5565b91509150935093915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b390611652565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068c90611652565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600061071085856003546108f5565b915091506107278183610a8f90919063ffffffff16565b9250505092915050565b60076020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107a257600080fd5b8060076000848152602001908152602001600020819055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081757600080fd5b8060038190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d957600080fd5b8060066000848152602001908152602001600020819055505050565b6000806000600660008781526020019081526020016000205490506109376127106109298388610ae490919063ffffffff16565b610b5490919063ffffffff16565b9250600760008781526020019081526020016000205491506000821415610a80576003548411156109685760035493505b600061097f85600454610ae490919063ffffffff16565b905060006005600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610a1757819350610a7d565b600080610a69600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685610b9e565b91509150610a78848383610cb0565b955050505b50505b82829250925050935093915050565b600080828401905083811015610ada576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad190611592565b60405180910390fd5b8091505092915050565b600080831415610af75760009050610b4e565b6000828402905082848281610b0857fe5b0414610b49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b40906115f2565b60405180910390fd5b809150505b92915050565b6000610b9683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610d68565b905092915050565b6000806000610bad8585610dc9565b509050600080610bbe888888610ef7565b73ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610c0357600080fd5b505afa158015610c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3b9190611112565b506dffffffffffffffffffffffffffff1691506dffffffffffffffffffffffffffff1691508273ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610c9a578082610c9d565b81815b8095508196505050505050935093915050565b6000808411610cf4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ceb90611612565b60405180910390fd5b600083118015610d045750600082115b610d43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3a906115d2565b60405180910390fd5b82610d578386610ae490919063ffffffff16565b81610d5e57fe5b0490509392505050565b60008083118290610daf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da69190611570565b60405180910390fd5b506000838581610dbb57fe5b049050809150509392505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610e3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e32906115b2565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610e75578284610e78565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ef0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee790611632565b60405180910390fd5b9250929050565b6000806000610f068585610dc9565b91509150858282604051602001610f1e9291906114e7565b60405160208183030381529060405280519060200120604051602001610f45929190611513565b6040516020818303038152906040528051906020012060001c925050509392505050565b600081359050610f78816117cc565b92915050565b600081359050610f8d816117e3565b92915050565b600081519050610fa2816117fa565b92915050565b600081359050610fb781611811565b92915050565b600081519050610fcc81611828565b92915050565b600060208284031215610fe457600080fd5b6000610ff284828501610f69565b91505092915050565b60006020828403121561100d57600080fd5b600061101b84828501610f7e565b91505092915050565b6000806000806080858703121561103a57600080fd5b600061104887828801610f7e565b945050602061105987828801610f69565b935050604061106a87828801610fa8565b925050606061107b87828801610fa8565b91505092959194509250565b6000806040838503121561109a57600080fd5b60006110a885828601610f7e565b92505060206110b985828601610fa8565b9150509250929050565b6000806000606084860312156110d857600080fd5b60006110e686828701610f7e565b93505060206110f786828701610fa8565b925050604061110886828701610fa8565b9150509250925092565b60008060006060848603121561112757600080fd5b600061113586828701610f93565b935050602061114686828701610f93565b925050604061115786828701610fbd565b9150509250925092565b60006020828403121561117357600080fd5b600061118184828501610fa8565b91505092915050565b611193816116dd565b82525050565b6111aa6111a5826116dd565b611780565b82525050565b6111c16111bc826116ef565b611792565b82525050565b60006111d2826116b6565b6111dc81856116c1565b93506111ec81856020860161174d565b6111f5816117ae565b840191505092915050565b600061120d601b836116c1565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b600061124d6025836116c1565b91507f556e697377617056324c6962726172793a204944454e544943414c5f4144445260008301527f45535345530000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006112b36020836116d2565b91507f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f6000830152602082019050919050565b60006112f36028836116c1565b91507f556e697377617056324c6962726172793a20494e53554646494349454e545f4c60008301527f49515549444954590000000000000000000000000000000000000000000000006020830152604082019050919050565b60006113596001836116d2565b91507fff000000000000000000000000000000000000000000000000000000000000006000830152600182019050919050565b60006113996021836116c1565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006113ff6025836116c1565b91507f556e697377617056324c6962726172793a20494e53554646494349454e545f4160008301527f4d4f554e540000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611465601e836116c1565b91507f556e697377617056324c6962726172793a205a45524f5f4144445245535300006000830152602082019050919050565b60006114a56009836116c1565b91507f6e6f74206f776e657200000000000000000000000000000000000000000000006000830152602082019050919050565b6114e181611733565b82525050565b60006114f38285611199565b6014820191506115038284611199565b6014820191508190509392505050565b600061151e8261134c565b915061152a8285611199565b60148201915061153a82846111b0565b602082019150611549826112a6565b91508190509392505050565b600060208201905061156a600083018461118a565b92915050565b6000602082019050818103600083015261158a81846111c7565b905092915050565b600060208201905081810360008301526115ab81611200565b9050919050565b600060208201905081810360008301526115cb81611240565b9050919050565b600060208201905081810360008301526115eb816112e6565b9050919050565b6000602082019050818103600083015261160b8161138c565b9050919050565b6000602082019050818103600083015261162b816113f2565b9050919050565b6000602082019050818103600083015261164b81611458565b9050919050565b6000602082019050818103600083015261166b81611498565b9050919050565b600060208201905061168760008301846114d8565b92915050565b60006040820190506116a260008301856114d8565b6116af60208301846114d8565b9392505050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006116e882611713565b9050919050565b6000819050919050565b60006dffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b60005b8381101561176b578082015181840152602081019050611750565b8381111561177a576000848401525b50505050565b600061178b8261179c565b9050919050565b6000819050919050565b60006117a7826117bf565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b6117d5816116dd565b81146117e057600080fd5b50565b6117ec816116ef565b81146117f757600080fd5b50565b611803816116f9565b811461180e57600080fd5b50565b61181a81611733565b811461182557600080fd5b50565b6118318161173d565b811461183c57600080fd5b5056fea2646970667358221220c4f4691c55a6f60952d2b16a65d33f6c76a735ea01650c190d9b01c189648dd664736f6c63430006040033"

// DeployCrossFee deploys a new Ethereum contract, binding an instance of CrossFee to it.
func DeployCrossFee(auth *bind.TransactOpts, backend bind.ContractBackend, weth_ common.Address, uni_ common.Address, crossInGas_ *big.Int) (common.Address, *types.Transaction, *CrossFee, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossFeeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrossFeeBin), backend, weth_, uni_, crossInGas_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossFee{CrossFeeCaller: CrossFeeCaller{contract: contract}, CrossFeeTransactor: CrossFeeTransactor{contract: contract}, CrossFeeFilterer: CrossFeeFilterer{contract: contract}}, nil
}

// CrossFee is an auto generated Go binding around an Ethereum contract.
type CrossFee struct {
	CrossFeeCaller     // Read-only binding to the contract
	CrossFeeTransactor // Write-only binding to the contract
	CrossFeeFilterer   // Log filterer for contract events
}

// CrossFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossFeeSession struct {
	Contract     *CrossFee         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossFeeCallerSession struct {
	Contract *CrossFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CrossFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossFeeTransactorSession struct {
	Contract     *CrossFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CrossFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossFeeRaw struct {
	Contract *CrossFee // Generic contract binding to access the raw methods on
}

// CrossFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossFeeCallerRaw struct {
	Contract *CrossFeeCaller // Generic read-only contract binding to access the raw methods on
}

// CrossFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossFeeTransactorRaw struct {
	Contract *CrossFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossFee creates a new instance of CrossFee, bound to a specific deployed contract.
func NewCrossFee(address common.Address, backend bind.ContractBackend) (*CrossFee, error) {
	contract, err := bindCrossFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossFee{CrossFeeCaller: CrossFeeCaller{contract: contract}, CrossFeeTransactor: CrossFeeTransactor{contract: contract}, CrossFeeFilterer: CrossFeeFilterer{contract: contract}}, nil
}

// NewCrossFeeCaller creates a new read-only instance of CrossFee, bound to a specific deployed contract.
func NewCrossFeeCaller(address common.Address, caller bind.ContractCaller) (*CrossFeeCaller, error) {
	contract, err := bindCrossFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossFeeCaller{contract: contract}, nil
}

// NewCrossFeeTransactor creates a new write-only instance of CrossFee, bound to a specific deployed contract.
func NewCrossFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossFeeTransactor, error) {
	contract, err := bindCrossFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossFeeTransactor{contract: contract}, nil
}

// NewCrossFeeFilterer creates a new log filterer instance of CrossFee, bound to a specific deployed contract.
func NewCrossFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossFeeFilterer, error) {
	contract, err := bindCrossFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossFeeFilterer{contract: contract}, nil
}

// bindCrossFee binds a generic wrapper to an already deployed contract.
func bindCrossFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossFeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossFee *CrossFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossFee.Contract.CrossFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossFee *CrossFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossFee.Contract.CrossFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossFee *CrossFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossFee.Contract.CrossFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossFee *CrossFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossFee *CrossFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossFee *CrossFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossFee.Contract.contract.Transact(opts, method, params...)
}

// MAXGASPRICE is a free data retrieval call binding the contract method 0xe3bbb4f1.
//
// Solidity: function MAX_GAS_PRICE() view returns(uint256)
func (_CrossFee *CrossFeeCaller) MAXGASPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "MAX_GAS_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXGASPRICE is a free data retrieval call binding the contract method 0xe3bbb4f1.
//
// Solidity: function MAX_GAS_PRICE() view returns(uint256)
func (_CrossFee *CrossFeeSession) MAXGASPRICE() (*big.Int, error) {
	return _CrossFee.Contract.MAXGASPRICE(&_CrossFee.CallOpts)
}

// MAXGASPRICE is a free data retrieval call binding the contract method 0xe3bbb4f1.
//
// Solidity: function MAX_GAS_PRICE() view returns(uint256)
func (_CrossFee *CrossFeeCallerSession) MAXGASPRICE() (*big.Int, error) {
	return _CrossFee.Contract.MAXGASPRICE(&_CrossFee.CallOpts)
}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 relayerFee, uint256 gasFee)
func (_CrossFee *CrossFeeCaller) CalCrossFee(opts *bind.CallOpts, resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	RelayerFee *big.Int
	GasFee     *big.Int
}, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "calCrossFee", resourceId, inputAmount, gasPrice)

	outstruct := new(struct {
		RelayerFee *big.Int
		GasFee     *big.Int
	})

	outstruct.RelayerFee = out[0].(*big.Int)
	outstruct.GasFee = out[1].(*big.Int)

	return *outstruct, err

}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 relayerFee, uint256 gasFee)
func (_CrossFee *CrossFeeSession) CalCrossFee(resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	RelayerFee *big.Int
	GasFee     *big.Int
}, error) {
	return _CrossFee.Contract.CalCrossFee(&_CrossFee.CallOpts, resourceId, inputAmount, gasPrice)
}

// CalCrossFee is a free data retrieval call binding the contract method 0x171cb896.
//
// Solidity: function calCrossFee(bytes32 resourceId, uint256 inputAmount, uint256 gasPrice) view returns(uint256 relayerFee, uint256 gasFee)
func (_CrossFee *CrossFeeCallerSession) CalCrossFee(resourceId [32]byte, inputAmount *big.Int, gasPrice *big.Int) (struct {
	RelayerFee *big.Int
	GasFee     *big.Int
}, error) {
	return _CrossFee.Contract.CalCrossFee(&_CrossFee.CallOpts, resourceId, inputAmount, gasPrice)
}

// CrossInGas is a free data retrieval call binding the contract method 0x69213d03.
//
// Solidity: function crossInGas() view returns(uint256)
func (_CrossFee *CrossFeeCaller) CrossInGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "crossInGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossInGas is a free data retrieval call binding the contract method 0x69213d03.
//
// Solidity: function crossInGas() view returns(uint256)
func (_CrossFee *CrossFeeSession) CrossInGas() (*big.Int, error) {
	return _CrossFee.Contract.CrossInGas(&_CrossFee.CallOpts)
}

// CrossInGas is a free data retrieval call binding the contract method 0x69213d03.
//
// Solidity: function crossInGas() view returns(uint256)
func (_CrossFee *CrossFeeCallerSession) CrossInGas() (*big.Int, error) {
	return _CrossFee.Contract.CrossInGas(&_CrossFee.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_CrossFee *CrossFeeCaller) EstimateFee(opts *bind.CallOpts, resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "estimateFee", resourceId, inputAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_CrossFee *CrossFeeSession) EstimateFee(resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	return _CrossFee.Contract.EstimateFee(&_CrossFee.CallOpts, resourceId, inputAmount)
}

// EstimateFee is a free data retrieval call binding the contract method 0x92c212d8.
//
// Solidity: function estimateFee(bytes32 resourceId, uint256 inputAmount) view returns(uint256 fee)
func (_CrossFee *CrossFeeCallerSession) EstimateFee(resourceId [32]byte, inputAmount *big.Int) (*big.Int, error) {
	return _CrossFee.Contract.EstimateFee(&_CrossFee.CallOpts, resourceId, inputAmount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossFee *CrossFeeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossFee *CrossFeeSession) Owner() (common.Address, error) {
	return _CrossFee.Contract.Owner(&_CrossFee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossFee *CrossFeeCallerSession) Owner() (common.Address, error) {
	return _CrossFee.Contract.Owner(&_CrossFee.CallOpts)
}

// ResourceToDefalutGasFee is a free data retrieval call binding the contract method 0xa8e31a68.
//
// Solidity: function resourceToDefalutGasFee(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeCaller) ResourceToDefalutGasFee(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "resourceToDefalutGasFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResourceToDefalutGasFee is a free data retrieval call binding the contract method 0xa8e31a68.
//
// Solidity: function resourceToDefalutGasFee(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeSession) ResourceToDefalutGasFee(arg0 [32]byte) (*big.Int, error) {
	return _CrossFee.Contract.ResourceToDefalutGasFee(&_CrossFee.CallOpts, arg0)
}

// ResourceToDefalutGasFee is a free data retrieval call binding the contract method 0xa8e31a68.
//
// Solidity: function resourceToDefalutGasFee(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeCallerSession) ResourceToDefalutGasFee(arg0 [32]byte) (*big.Int, error) {
	return _CrossFee.Contract.ResourceToDefalutGasFee(&_CrossFee.CallOpts, arg0)
}

// ResourceToFeeRate is a free data retrieval call binding the contract method 0x0100a555.
//
// Solidity: function resourceToFeeRate(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeCaller) ResourceToFeeRate(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "resourceToFeeRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResourceToFeeRate is a free data retrieval call binding the contract method 0x0100a555.
//
// Solidity: function resourceToFeeRate(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeSession) ResourceToFeeRate(arg0 [32]byte) (*big.Int, error) {
	return _CrossFee.Contract.ResourceToFeeRate(&_CrossFee.CallOpts, arg0)
}

// ResourceToFeeRate is a free data retrieval call binding the contract method 0x0100a555.
//
// Solidity: function resourceToFeeRate(bytes32 ) view returns(uint256)
func (_CrossFee *CrossFeeCallerSession) ResourceToFeeRate(arg0 [32]byte) (*big.Int, error) {
	return _CrossFee.Contract.ResourceToFeeRate(&_CrossFee.CallOpts, arg0)
}

// ResourceToTokenAddrss is a free data retrieval call binding the contract method 0xeab44ff6.
//
// Solidity: function resourceToTokenAddrss(bytes32 ) view returns(address)
func (_CrossFee *CrossFeeCaller) ResourceToTokenAddrss(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "resourceToTokenAddrss", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceToTokenAddrss is a free data retrieval call binding the contract method 0xeab44ff6.
//
// Solidity: function resourceToTokenAddrss(bytes32 ) view returns(address)
func (_CrossFee *CrossFeeSession) ResourceToTokenAddrss(arg0 [32]byte) (common.Address, error) {
	return _CrossFee.Contract.ResourceToTokenAddrss(&_CrossFee.CallOpts, arg0)
}

// ResourceToTokenAddrss is a free data retrieval call binding the contract method 0xeab44ff6.
//
// Solidity: function resourceToTokenAddrss(bytes32 ) view returns(address)
func (_CrossFee *CrossFeeCallerSession) ResourceToTokenAddrss(arg0 [32]byte) (common.Address, error) {
	return _CrossFee.Contract.ResourceToTokenAddrss(&_CrossFee.CallOpts, arg0)
}

// UniFactory is a free data retrieval call binding the contract method 0xd702ffd0.
//
// Solidity: function uni_factory() view returns(address)
func (_CrossFee *CrossFeeCaller) UniFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "uni_factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniFactory is a free data retrieval call binding the contract method 0xd702ffd0.
//
// Solidity: function uni_factory() view returns(address)
func (_CrossFee *CrossFeeSession) UniFactory() (common.Address, error) {
	return _CrossFee.Contract.UniFactory(&_CrossFee.CallOpts)
}

// UniFactory is a free data retrieval call binding the contract method 0xd702ffd0.
//
// Solidity: function uni_factory() view returns(address)
func (_CrossFee *CrossFeeCallerSession) UniFactory() (common.Address, error) {
	return _CrossFee.Contract.UniFactory(&_CrossFee.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CrossFee *CrossFeeCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossFee.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CrossFee *CrossFeeSession) Weth() (common.Address, error) {
	return _CrossFee.Contract.Weth(&_CrossFee.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CrossFee *CrossFeeCallerSession) Weth() (common.Address, error) {
	return _CrossFee.Contract.Weth(&_CrossFee.CallOpts)
}

// SetCorssFeeRate is a paid mutator transaction binding the contract method 0xec1bef0c.
//
// Solidity: function setCorssFeeRate(bytes32 resourceId, uint256 feeRate_) returns()
func (_CrossFee *CrossFeeTransactor) SetCorssFeeRate(opts *bind.TransactOpts, resourceId [32]byte, feeRate_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setCorssFeeRate", resourceId, feeRate_)
}

// SetCorssFeeRate is a paid mutator transaction binding the contract method 0xec1bef0c.
//
// Solidity: function setCorssFeeRate(bytes32 resourceId, uint256 feeRate_) returns()
func (_CrossFee *CrossFeeSession) SetCorssFeeRate(resourceId [32]byte, feeRate_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetCorssFeeRate(&_CrossFee.TransactOpts, resourceId, feeRate_)
}

// SetCorssFeeRate is a paid mutator transaction binding the contract method 0xec1bef0c.
//
// Solidity: function setCorssFeeRate(bytes32 resourceId, uint256 feeRate_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetCorssFeeRate(resourceId [32]byte, feeRate_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetCorssFeeRate(&_CrossFee.TransactOpts, resourceId, feeRate_)
}

// SetCrossInGas is a paid mutator transaction binding the contract method 0x1120e655.
//
// Solidity: function setCrossInGas(uint256 crossInGas_) returns()
func (_CrossFee *CrossFeeTransactor) SetCrossInGas(opts *bind.TransactOpts, crossInGas_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setCrossInGas", crossInGas_)
}

// SetCrossInGas is a paid mutator transaction binding the contract method 0x1120e655.
//
// Solidity: function setCrossInGas(uint256 crossInGas_) returns()
func (_CrossFee *CrossFeeSession) SetCrossInGas(crossInGas_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetCrossInGas(&_CrossFee.TransactOpts, crossInGas_)
}

// SetCrossInGas is a paid mutator transaction binding the contract method 0x1120e655.
//
// Solidity: function setCrossInGas(uint256 crossInGas_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetCrossInGas(crossInGas_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetCrossInGas(&_CrossFee.TransactOpts, crossInGas_)
}

// SetDefaultGasFee is a paid mutator transaction binding the contract method 0xd090223e.
//
// Solidity: function setDefaultGasFee(bytes32 resourceId, uint256 gasFee_) returns()
func (_CrossFee *CrossFeeTransactor) SetDefaultGasFee(opts *bind.TransactOpts, resourceId [32]byte, gasFee_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setDefaultGasFee", resourceId, gasFee_)
}

// SetDefaultGasFee is a paid mutator transaction binding the contract method 0xd090223e.
//
// Solidity: function setDefaultGasFee(bytes32 resourceId, uint256 gasFee_) returns()
func (_CrossFee *CrossFeeSession) SetDefaultGasFee(resourceId [32]byte, gasFee_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetDefaultGasFee(&_CrossFee.TransactOpts, resourceId, gasFee_)
}

// SetDefaultGasFee is a paid mutator transaction binding the contract method 0xd090223e.
//
// Solidity: function setDefaultGasFee(bytes32 resourceId, uint256 gasFee_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetDefaultGasFee(resourceId [32]byte, gasFee_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetDefaultGasFee(&_CrossFee.TransactOpts, resourceId, gasFee_)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 max_) returns()
func (_CrossFee *CrossFeeTransactor) SetMaxGasPrice(opts *bind.TransactOpts, max_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setMaxGasPrice", max_)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 max_) returns()
func (_CrossFee *CrossFeeSession) SetMaxGasPrice(max_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetMaxGasPrice(&_CrossFee.TransactOpts, max_)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 max_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetMaxGasPrice(max_ *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetMaxGasPrice(&_CrossFee.TransactOpts, max_)
}

// SetResource is a paid mutator transaction binding the contract method 0x02e3588f.
//
// Solidity: function setResource(bytes32 resourceId, address tokenAddress, uint256 defaltGasFee, uint256 feeRate) returns()
func (_CrossFee *CrossFeeTransactor) SetResource(opts *bind.TransactOpts, resourceId [32]byte, tokenAddress common.Address, defaltGasFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setResource", resourceId, tokenAddress, defaltGasFee, feeRate)
}

// SetResource is a paid mutator transaction binding the contract method 0x02e3588f.
//
// Solidity: function setResource(bytes32 resourceId, address tokenAddress, uint256 defaltGasFee, uint256 feeRate) returns()
func (_CrossFee *CrossFeeSession) SetResource(resourceId [32]byte, tokenAddress common.Address, defaltGasFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetResource(&_CrossFee.TransactOpts, resourceId, tokenAddress, defaltGasFee, feeRate)
}

// SetResource is a paid mutator transaction binding the contract method 0x02e3588f.
//
// Solidity: function setResource(bytes32 resourceId, address tokenAddress, uint256 defaltGasFee, uint256 feeRate) returns()
func (_CrossFee *CrossFeeTransactorSession) SetResource(resourceId [32]byte, tokenAddress common.Address, defaltGasFee *big.Int, feeRate *big.Int) (*types.Transaction, error) {
	return _CrossFee.Contract.SetResource(&_CrossFee.TransactOpts, resourceId, tokenAddress, defaltGasFee, feeRate)
}

// SetUni is a paid mutator transaction binding the contract method 0x6f3ca55d.
//
// Solidity: function setUni(address uni_) returns()
func (_CrossFee *CrossFeeTransactor) SetUni(opts *bind.TransactOpts, uni_ common.Address) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setUni", uni_)
}

// SetUni is a paid mutator transaction binding the contract method 0x6f3ca55d.
//
// Solidity: function setUni(address uni_) returns()
func (_CrossFee *CrossFeeSession) SetUni(uni_ common.Address) (*types.Transaction, error) {
	return _CrossFee.Contract.SetUni(&_CrossFee.TransactOpts, uni_)
}

// SetUni is a paid mutator transaction binding the contract method 0x6f3ca55d.
//
// Solidity: function setUni(address uni_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetUni(uni_ common.Address) (*types.Transaction, error) {
	return _CrossFee.Contract.SetUni(&_CrossFee.TransactOpts, uni_)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address weth_) returns()
func (_CrossFee *CrossFeeTransactor) SetWETH(opts *bind.TransactOpts, weth_ common.Address) (*types.Transaction, error) {
	return _CrossFee.contract.Transact(opts, "setWETH", weth_)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address weth_) returns()
func (_CrossFee *CrossFeeSession) SetWETH(weth_ common.Address) (*types.Transaction, error) {
	return _CrossFee.Contract.SetWETH(&_CrossFee.TransactOpts, weth_)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address weth_) returns()
func (_CrossFee *CrossFeeTransactorSession) SetWETH(weth_ common.Address) (*types.Transaction, error) {
	return _CrossFee.Contract.SetWETH(&_CrossFee.TransactOpts, weth_)
}
