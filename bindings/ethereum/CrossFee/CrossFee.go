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
var CrossFeeBin = "0x60806040526412a05f20006003553480156200001a57600080fd5b5060405162001b1c38038062001b1c833981810160405281019062000040919062000140565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060048190555050505062000208565b6000815190506200012381620001d4565b92915050565b6000815190506200013a81620001ee565b92915050565b6000806000606084860312156200015657600080fd5b6000620001668682870162000112565b9350506020620001798682870162000112565b92505060406200018c8682870162000129565b9150509250925092565b6000620001a382620001aa565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b620001df8162000196565b8114620001eb57600080fd5b50565b620001f981620001ca565b81146200020557600080fd5b50565b61190480620002186000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063d2fa635e11610071578063d2fa635e146102b7578063d702ffd0146102d3578063e3bbb4f1146102f1578063eab44ff61461030f578063ec1bef0c1461033f5761010b565b80638da5cb5b1461021d57806392c212d81461023b578063a8e31a681461026b578063d090223e1461029b5761010b565b80633fc8cef3116100de5780633fc8cef3146101a95780635b769f3c146101c757806369213d03146101e35780636f3ca55d146102015761010b565b80630100a5551461011057806302e3588f146101405780631120e6551461015c578063171cb89614610178575b600080fd5b61012a6004803603810190610125919061108a565b61035b565b6040516101379190611701565b60405180910390f35b61015a600480360381019061015591906110b3565b610373565b005b610176600480360381019061017191906111f0565b61048a565b005b610192600480360381019061018d9190611152565b6104ed565b6040516101a092919061171c565b60405180910390f35b6101b1610622565b6040516101be91906115e4565b60405180910390f35b6101e160048036038101906101dc9190611061565b610648565b005b6101eb61071b565b6040516101f89190611701565b60405180910390f35b61021b60048036038101906102169190611061565b610721565b005b6102256107f4565b60405161023291906115e4565b60405180910390f35b61025560048036038101906102509190611116565b610819565b6040516102629190611701565b60405180910390f35b6102856004803603810190610280919061108a565b61095a565b6040516102929190611701565b60405180910390f35b6102b560048036038101906102b09190611116565b610972565b005b6102d160048036038101906102cc91906111f0565b6109e7565b005b6102db610a4a565b6040516102e891906115e4565b60405180910390f35b6102f9610a70565b6040516103069190611701565b60405180910390f35b6103296004803603810190610324919061108a565b610a76565b60405161033691906115e4565b60405180910390f35b61035960048036038101906103549190611116565b610aa9565b005b60066020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610402576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f9906116e1565b60405180910390fd5b826005600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660008681526020019081526020016000208190555081600760008681526020019081526020016000208190555050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104e357600080fd5b8060048190555050565b60008060006006600087815260200190815260200160002054905061052f6127106105218388610b1e90919063ffffffff16565b610b8e90919063ffffffff16565b9250600760008781526020019081526020016000205491506000821415610619576003548411156105605760035493505b600061057785600454610b1e90919063ffffffff16565b905060006005600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600080610603600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685610bd8565b91509150610612848383610cea565b9550505050505b50935093915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ce906116e1565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a7906116e1565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060066000858152602001908152602001600020549050600061085b61271061084d8487610b1e90919063ffffffff16565b610b8e90919063ffffffff16565b90506000600760008781526020019081526020016000205490506000811415610939576000610897600354600454610b1e90919063ffffffff16565b905060006005600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600080610923600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685610bd8565b91509150610932848383610cea565b9450505050505b61094c8183610da290919063ffffffff16565b935083935050505092915050565b60076020528060005260406000206000915090505481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109cb57600080fd5b8060076000848152602001908152602001600020819055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4057600080fd5b8060038190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b0257600080fd5b8060066000848152602001908152602001600020819055505050565b600080831415610b315760009050610b88565b6000828402905082848281610b4257fe5b0414610b83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7a90611681565b60405180910390fd5b809150505b92915050565b6000610bd083836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610df7565b905092915050565b6000806000610be78585610e58565b509050600080610bf8888888610f86565b73ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610c3d57600080fd5b505afa158015610c51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7591906111a1565b506dffffffffffffffffffffffffffff1691506dffffffffffffffffffffffffffff1691508273ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610cd4578082610cd7565b81815b8095508196505050505050935093915050565b6000808411610d2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d25906116a1565b60405180910390fd5b600083118015610d3e5750600082115b610d7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7490611661565b60405180910390fd5b82610d918386610b1e90919063ffffffff16565b81610d9857fe5b0490509392505050565b600080828401905083811015610ded576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de490611621565b60405180910390fd5b8091505092915050565b60008083118290610e3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3591906115ff565b60405180910390fd5b506000838581610e4a57fe5b049050809150509392505050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610eca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec190611641565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610f04578284610f07565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f76906116c1565b60405180910390fd5b9250929050565b6000806000610f958585610e58565b91509150858282604051602001610fad929190611576565b60405160208183030381529060405280519060200120604051602001610fd49291906115a2565b6040516020818303038152906040528051906020012060001c925050509392505050565b6000813590506110078161185b565b92915050565b60008135905061101c81611872565b92915050565b60008151905061103181611889565b92915050565b600081359050611046816118a0565b92915050565b60008151905061105b816118b7565b92915050565b60006020828403121561107357600080fd5b600061108184828501610ff8565b91505092915050565b60006020828403121561109c57600080fd5b60006110aa8482850161100d565b91505092915050565b600080600080608085870312156110c957600080fd5b60006110d78782880161100d565b94505060206110e887828801610ff8565b93505060406110f987828801611037565b925050606061110a87828801611037565b91505092959194509250565b6000806040838503121561112957600080fd5b60006111378582860161100d565b925050602061114885828601611037565b9150509250929050565b60008060006060848603121561116757600080fd5b60006111758682870161100d565b935050602061118686828701611037565b925050604061119786828701611037565b9150509250925092565b6000806000606084860312156111b657600080fd5b60006111c486828701611022565b93505060206111d586828701611022565b92505060406111e68682870161104c565b9150509250925092565b60006020828403121561120257600080fd5b600061121084828501611037565b91505092915050565b6112228161176c565b82525050565b6112396112348261176c565b61180f565b82525050565b61125061124b8261177e565b611821565b82525050565b600061126182611745565b61126b8185611750565b935061127b8185602086016117dc565b6112848161183d565b840191505092915050565b600061129c601b83611750565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006112dc602583611750565b91507f556e697377617056324c6962726172793a204944454e544943414c5f4144445260008301527f45535345530000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611342602083611761565b91507f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f6000830152602082019050919050565b6000611382602883611750565b91507f556e697377617056324c6962726172793a20494e53554646494349454e545f4c60008301527f49515549444954590000000000000000000000000000000000000000000000006020830152604082019050919050565b60006113e8600183611761565b91507fff000000000000000000000000000000000000000000000000000000000000006000830152600182019050919050565b6000611428602183611750565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061148e602583611750565b91507f556e697377617056324c6962726172793a20494e53554646494349454e545f4160008301527f4d4f554e540000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006114f4601e83611750565b91507f556e697377617056324c6962726172793a205a45524f5f4144445245535300006000830152602082019050919050565b6000611534600983611750565b91507f6e6f74206f776e657200000000000000000000000000000000000000000000006000830152602082019050919050565b611570816117c2565b82525050565b60006115828285611228565b6014820191506115928284611228565b6014820191508190509392505050565b60006115ad826113db565b91506115b98285611228565b6014820191506115c9828461123f565b6020820191506115d882611335565b91508190509392505050565b60006020820190506115f96000830184611219565b92915050565b600060208201905081810360008301526116198184611256565b905092915050565b6000602082019050818103600083015261163a8161128f565b9050919050565b6000602082019050818103600083015261165a816112cf565b9050919050565b6000602082019050818103600083015261167a81611375565b9050919050565b6000602082019050818103600083015261169a8161141b565b9050919050565b600060208201905081810360008301526116ba81611481565b9050919050565b600060208201905081810360008301526116da816114e7565b9050919050565b600060208201905081810360008301526116fa81611527565b9050919050565b60006020820190506117166000830184611567565b92915050565b60006040820190506117316000830185611567565b61173e6020830184611567565b9392505050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000611777826117a2565b9050919050565b6000819050919050565b60006dffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b60005b838110156117fa5780820151818401526020810190506117df565b83811115611809576000848401525b50505050565b600061181a8261182b565b9050919050565b6000819050919050565b60006118368261184e565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b6118648161176c565b811461186f57600080fd5b50565b61187b8161177e565b811461188657600080fd5b50565b61189281611788565b811461189d57600080fd5b50565b6118a9816117c2565b81146118b457600080fd5b50565b6118c0816117cc565b81146118cb57600080fd5b5056fea26469706673582212204ba05d19ed614b57cc51ab3335d12fb75a3bf00ef3f8401089ba798cfa6a8be264736f6c63430006040033"

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
