// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SRC20Handler

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

// SRC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type SRC20HandlerDepositRecord struct {
	TokenAddress                   common.ContractAddress
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.ContractAddress
	Amount                         *big.Int
}

// SRC20HandlerABI is the input ABI used to generate the binding from.
const SRC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_currencyToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToCurrency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structSRC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setReourceTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"}],\"name\":\"setResourceCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SRC20HandlerBin is the compiled bytecode used for deploying new contracts.
var SRC20HandlerBin = "0x60806040527f3be6bf24d822bcd6f6348f6f5a5c2d3108f04991ee63e80cde49a8c4746a0ef360001b6004557fcf19eb4256453a4e30b6a06d651f1970c223fb6bd1826a28ed861f0e602db9b860001b6005557f868bd6629e7c2e3d2ccf7b9968fad79b448e7a2bfb3ee20ed1acbc695c3c8b2360001b6006557f7c98e64bd943448b4e24ef8c2cdec7b8b1275970cfe10daf2a9bfa4b04dce90560001b6007557fa6a366f1a72e1aef5d8d52ee240a476f619d15be7bc62d3df37496025b83459f60001b6008557ff1964f6690a0536daa42e5c575091297d2479edcc96f721ad85b95358644d27660001b6009557f9ab0d7c07029f006485cf3468ce7811aa8743b5a108599f6bec9367c50ac6aad60001b600a557fa6cafc6282f61eff9032603a017e652f68410d3d3c69f0a3eeca8f181aec1d1760001b600b557f6800e94e36131c049eaeb631e4530829b0d3d20d5b637c8015a8dc9cedd70aed60001b600c553480156200017057600080fd5b5060405162002256380380620022568339818101604052810190620001969190620001f4565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200026e565b600081519050620001ee8162000254565b92915050565b6000602082840312156200020757600080fd5b60006200021784828501620001dd565b91505092915050565b60006200022d8262000234565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200025f8162000220565b81146200026b57600080fd5b50565b611fd8806200027e6000396000f3fe6080604052600436106100ab5760003560e01c80639201de55116100645780639201de55146101f15780639707bece1461022e578063ba40f28414610257578063ba484c0914610273578063c6a3f31f146102b0578063dfce8258146102ed576100b2565b80630396ec10146100b75780630a6d55d8146100e0578063318c136e1461011d5780634402027f146101485780635a73b0bf1461018b57806380b90744146101b4576100b2565b366100b257005b600080fd5b3480156100c357600080fd5b506100de60048036038101906100d99190611667565b610316565b005b3480156100ec57600080fd5b5061010760048036038101906101029190611602565b610524565b6040516101149190611c10565b60405180910390f35b34801561012957600080fd5b50610132610557565b60405161013f9190611c10565b60405180910390f35b34801561015457600080fd5b5061016f600480360381019061016a919061189e565b61057c565b6040516101829796959493929190611c2b565b60405180910390f35b34801561019757600080fd5b506101b260048036038101906101ad91906117b5565b6106bd565b005b3480156101c057600080fd5b506101db60048036038101906101d69190611821565b610759565b6040516101e89190611ca1565b60405180910390f35b3480156101fd57600080fd5b5061021860048036038101906102139190611602565b610787565b6040516102259190611cde565b60405180910390f35b34801561023a57600080fd5b506102556004803603810190610250919061162b565b61092e565b005b610271600480360381019061026c919061170e565b61098c565b005b34801561027f57600080fd5b5061029a60048036038101906102959190611862565b610d84565b6040516102a79190611d80565b60405180910390f35b3480156102bc57600080fd5b506102d760048036038101906102d29190611602565b610f79565b6040516102e49190611cbc565b60405180910390f35b3480156102f957600080fd5b50610314600480360381019061030f91906116b6565b611029565b005b61031e61107e565b600073ffffffffffffffffffffffffffffffffffffffff166001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461042a5760006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a0712d68836040518263ffffffff1660e01b81526004016103f69190611da2565b600060405180830381600087803b15801561041057600080fd5b505af1158015610424573d6000803e3d6000fd5b50505050505b6104e082600260008681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104d55780601f106104aa576101008083540402835291602001916104d5565b820191906000526020600020905b8154815290600101906020018083116104b857829003601f168201915b50505050508361110f565b61051f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051690611d40565b60405180910390fd5b505050565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600d602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060000160159054906101000a900460ff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106875780601f1061065c57610100808354040283529160200191610687565b820191906000526020600020905b81548152906001019060200180831161066a57829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905087565b6106c561107e565b6107148285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508361110f565b610753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074a90611d60565b60405180910390fd5b50505050565b6003818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b60606000809050606060206040519080825280601f01601f1916602001820160405280156107c45781602001600182028036833780820191505090505b50905060008090505b602081101561087c5760008160080260020a8660001c0260001b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461085f578083858151811061082357fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350838060010194505061086e565b6000841461086d575061087c565b5b5080806001019150506107cd565b506060826040519080825280601f01601f1916602001820160405280156108b25781602001600182028036833780820191505090505b50905060008090505b83811015610922578281815181106108cf57fe5b602001015160f81c60f81b8282815181106108e657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506108bb565b50809350505050919050565b61093661107e565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61099461107e565b60006001600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506060600260008a81526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a755780601f10610a4a57610100808354040283529160200191610a75565b820191906000526020600020905b815481529060010190602001808311610a5857829003601f168201915b505050505090506060610a86611139565b9050610a928282611190565b610ad1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac890611d20565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610b8a576000839050610b288386604051806020016040528060008152506000801b61124d565b8073ffffffffffffffffffffffffffffffffffffffff166344df8e706040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610b7057600080fd5b505af1158015610b84573d6000803e3d6000fd5b50505050505b6040518060e001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018787905060ff1681526020018a60ff1681526020018b815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018873ffffffffffffffffffffffffffffffffffffffff16815260200185815250600d60008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160000160156101000a81548160ff021916908360ff160217905550606082015181600101556080820151816002019080519060200190610d2392919061131a565b5060a08201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816004015590505050505050505050505050565b610d8c61139a565b600d60008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016000820160159054906101000a900460ff1660ff1660ff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f085780601f10610edd57610100808354040283529160200191610f08565b820191906000526020600020905b815481529060010190602001808311610eeb57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481525050905092915050565b60026020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110215780601f10610ff657610100808354040283529160200191611021565b820191906000526020600020905b81548152906001019060200180831161100457829003601f168201915b505050505081565b61103161107e565b818160026000868152602001908152602001600020919061105392919061140c565b508260038383604051611067929190611bf7565b908152602001604051809103902081905550505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461110d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110490611d00565b60405180910390fd5b565b6000611130848484604051806020016040528060008152506000801b6112a9565b90509392505050565b60608060206040519080825280601f01601f1916602001820160405280156111705781602001600182028036833780820191505090505b5090506000600754602083a18151905061118981610787565b9250505090565b600081518351146111a45760009050611247565b60008090505b8351811015611241578281815181106111bf57fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168482815181106111f857fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614611234576000915050611247565b80806001019150506111aa565b50600190505b92915050565b606060806040519080825280601f01601f1916602001820160405280156112835781602001600182028036833780820191505090505b509050848152836020820152826040820152816060820152600b54608082a15050505050565b6000606060a06040519080825280601f01601f1916602001820160405280156112e15781602001600182028036833780820191505090505b50905086815285602082015284604082015283606082015282608082015260065460a082a1608081015191508191505095945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061135b57805160ff1916838001178555611389565b82800160010185558215611389579182015b8281111561138857825182559160200191906001019061136d565b5b509050611396919061148c565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600060ff1681526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061144d57803560ff191683800117855561147b565b8280016001018555821561147b579182015b8281111561147a57823582559160200191906001019061145f565b5b509050611488919061148c565b5090565b6114ae91905b808211156114aa576000816000905550600101611492565b5090565b90565b6000813590506114c081611f2f565b92915050565b6000813590506114d581611f46565b92915050565b60008083601f8401126114ed57600080fd5b8235905067ffffffffffffffff81111561150657600080fd5b60208301915083600182028301111561151e57600080fd5b9250929050565b60008083601f84011261153757600080fd5b8235905067ffffffffffffffff81111561155057600080fd5b60208301915083600182028301111561156857600080fd5b9250929050565b600082601f83011261158057600080fd5b813561159361158e82611dea565b611dbd565b915080825260208301602083018583830111156115af57600080fd5b6115ba838284611edc565b50505092915050565b6000813590506115d281611f5d565b92915050565b6000813590506115e781611f74565b92915050565b6000813590506115fc81611f8b565b92915050565b60006020828403121561161457600080fd5b6000611622848285016114c6565b91505092915050565b6000806040838503121561163e57600080fd5b600061164c858286016114c6565b925050602061165d858286016114b1565b9150509250929050565b60008060006060848603121561167c57600080fd5b600061168a868287016114c6565b935050602061169b868287016114b1565b92505060406116ac868287016115c3565b9150509250925092565b6000806000604084860312156116cb57600080fd5b60006116d9868287016114c6565b935050602084013567ffffffffffffffff8111156116f657600080fd5b61170286828701611525565b92509250509250925092565b600080600080600080600060c0888a03121561172957600080fd5b60006117378a828b016114c6565b97505060206117488a828b016115ed565b96505060406117598a828b016115d8565b955050606061176a8a828b016114b1565b945050608088013567ffffffffffffffff81111561178757600080fd5b6117938a828b016114db565b935093505060a06117a68a828b016115c3565b91505092959891949750929550565b600080600080606085870312156117cb57600080fd5b600085013567ffffffffffffffff8111156117e557600080fd5b6117f187828801611525565b94509450506020611804878288016114b1565b9250506040611815878288016115c3565b91505092959194509250565b60006020828403121561183357600080fd5b600082013567ffffffffffffffff81111561184d57600080fd5b6118598482850161156f565b91505092915050565b6000806040838503121561187557600080fd5b6000611883858286016115d8565b9250506020611894858286016115ed565b9150509250929050565b600080604083850312156118b157600080fd5b60006118bf858286016115ed565b92505060206118d0858286016115d8565b9150509250929050565b6118e381611e75565b82525050565b6118f281611e75565b82525050565b61190181611e87565b82525050565b61191081611e87565b82525050565b600061192182611e16565b61192b8185611e37565b935061193b818560208601611eeb565b61194481611f1e565b840191505092915050565b600061195a82611e16565b6119648185611e48565b9350611974818560208601611eeb565b61197d81611f1e565b840191505092915050565b60006119948385611e6a565b93506119a1838584611edc565b82840190509392505050565b60006119b882611e2c565b6119c28185611e59565b93506119d2818560208601611eeb565b6119db81611f1e565b840191505092915050565b60006119f182611e21565b6119fb8185611e59565b9350611a0b818560208601611eeb565b611a1481611f1e565b840191505092915050565b6000611a2c601e83611e59565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611a6c601083611e59565b91507f696e76616c69642063757272656e6379000000000000000000000000000000006000830152602082019050919050565b6000611aac601d83611e59565b91507f73656e6420746f207265636569656e74416472657373206661696c65640000006000830152602082019050919050565b6000611aec601c83611e59565b91507f776974686472617720746f20726563697069656e74206661696c6564000000006000830152602082019050919050565b600060e083016000830151611b3760008601826118da565b506020830151611b4a6020860182611bd9565b506040830151611b5d6040860182611bd9565b506060830151611b7060608601826118f8565b5060808301518482036080860152611b888282611916565b91505060a0830151611b9d60a08601826118da565b5060c0830151611bb060c0860182611bbb565b508091505092915050565b611bc481611eb1565b82525050565b611bd381611eb1565b82525050565b611be281611ecf565b82525050565b611bf181611ecf565b82525050565b6000611c04828486611988565b91508190509392505050565b6000602082019050611c2560008301846118e9565b92915050565b600060e082019050611c40600083018a6118e9565b611c4d6020830189611be8565b611c5a6040830188611be8565b611c676060830187611907565b8181036080830152611c79818661194f565b9050611c8860a08301856118e9565b611c9560c0830184611bca565b98975050505050505050565b6000602082019050611cb66000830184611907565b92915050565b60006020820190508181036000830152611cd681846119e6565b905092915050565b60006020820190508181036000830152611cf881846119ad565b905092915050565b60006020820190508181036000830152611d1981611a1f565b9050919050565b60006020820190508181036000830152611d3981611a5f565b9050919050565b60006020820190508181036000830152611d5981611a9f565b9050919050565b60006020820190508181036000830152611d7981611adf565b9050919050565b60006020820190508181036000830152611d9a8184611b1f565b905092915050565b6000602082019050611db76000830184611bca565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611de057600080fd5b8060405250919050565b600067ffffffffffffffff821115611e0157600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611e8082611e91565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611f09578082015181840152602081019050611eee565b83811115611f18576000848401525b50505050565b6000601f19601f8301169050919050565b611f3881611e75565b8114611f4357600080fd5b50565b611f4f81611e87565b8114611f5a57600080fd5b50565b611f6681611eb1565b8114611f7157600080fd5b50565b611f7d81611ebb565b8114611f8857600080fd5b50565b611f9481611ecf565b8114611f9f57600080fd5b5056fea2646970667358221220526a5221a595d7a29fe9e9aa8f75ff1d38f54603157eaefce23cd8d478f38fd064736f6c63430006040033"

// DeploySRC20Handler deploys a new Ethereum contract, binding an instance of SRC20Handler to it.
func DeploySRC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *SRC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SRC20HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SRC20Handler{SRC20HandlerCaller: SRC20HandlerCaller{contract: contract}, SRC20HandlerTransactor: SRC20HandlerTransactor{contract: contract}, SRC20HandlerFilterer: SRC20HandlerFilterer{contract: contract}}, nil
}

// SRC20Handler is an auto generated Go binding around an Ethereum contract.
type SRC20Handler struct {
	SRC20HandlerCaller     // Read-only binding to the contract
	SRC20HandlerTransactor // Write-only binding to the contract
	SRC20HandlerFilterer   // Log filterer for contract events
}

// SRC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SRC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SRC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SRC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SRC20HandlerSession struct {
	Contract     *SRC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SRC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SRC20HandlerCallerSession struct {
	Contract *SRC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SRC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SRC20HandlerTransactorSession struct {
	Contract     *SRC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SRC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SRC20HandlerRaw struct {
	Contract *SRC20Handler // Generic contract binding to access the raw methods on
}

// SRC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SRC20HandlerCallerRaw struct {
	Contract *SRC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// SRC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SRC20HandlerTransactorRaw struct {
	Contract *SRC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSRC20Handler creates a new instance of SRC20Handler, bound to a specific deployed contract.
func NewSRC20Handler(address common.Address, backend bind.ContractBackend) (*SRC20Handler, error) {
	contract, err := bindSRC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SRC20Handler{SRC20HandlerCaller: SRC20HandlerCaller{contract: contract}, SRC20HandlerTransactor: SRC20HandlerTransactor{contract: contract}, SRC20HandlerFilterer: SRC20HandlerFilterer{contract: contract}}, nil
}

// NewSRC20HandlerCaller creates a new read-only instance of SRC20Handler, bound to a specific deployed contract.
func NewSRC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*SRC20HandlerCaller, error) {
	contract, err := bindSRC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SRC20HandlerCaller{contract: contract}, nil
}

// NewSRC20HandlerTransactor creates a new write-only instance of SRC20Handler, bound to a specific deployed contract.
func NewSRC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*SRC20HandlerTransactor, error) {
	contract, err := bindSRC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SRC20HandlerTransactor{contract: contract}, nil
}

// NewSRC20HandlerFilterer creates a new log filterer instance of SRC20Handler, bound to a specific deployed contract.
func NewSRC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*SRC20HandlerFilterer, error) {
	contract, err := bindSRC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SRC20HandlerFilterer{contract: contract}, nil
}

// bindSRC20Handler binds a generic wrapper to an already deployed contract.
func bindSRC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC20Handler *SRC20HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC20Handler.Contract.SRC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC20Handler *SRC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SRC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC20Handler *SRC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SRC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC20Handler *SRC20HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC20Handler *SRC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC20Handler *SRC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC20Handler *SRC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC20Handler *SRC20HandlerSession) BridgeAddress() (common.ContractAddress, error) {
	return _SRC20Handler.Contract.BridgeAddress(&_SRC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC20Handler *SRC20HandlerCallerSession) BridgeAddress() (common.ContractAddress, error) {
	return _SRC20Handler.Contract.BridgeAddress(&_SRC20Handler.CallOpts)
}

// CurrencyToResourceID is a free data retrieval call binding the contract method 0x80b90744.
//
// Solidity: function _currencyToResourceID(string ) view returns(bytes32)
func (_SRC20Handler *SRC20HandlerCaller) CurrencyToResourceID(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "_currencyToResourceID", arg0)
	return *ret0, err
}

// CurrencyToResourceID is a free data retrieval call binding the contract method 0x80b90744.
//
// Solidity: function _currencyToResourceID(string ) view returns(bytes32)
func (_SRC20Handler *SRC20HandlerSession) CurrencyToResourceID(arg0 string) ([32]byte, error) {
	return _SRC20Handler.Contract.CurrencyToResourceID(&_SRC20Handler.CallOpts, arg0)
}

// CurrencyToResourceID is a free data retrieval call binding the contract method 0x80b90744.
//
// Solidity: function _currencyToResourceID(string ) view returns(bytes32)
func (_SRC20Handler *SRC20HandlerCallerSession) CurrencyToResourceID(arg0 string) ([32]byte, error) {
	return _SRC20Handler.Contract.CurrencyToResourceID(&_SRC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_SRC20Handler *SRC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.ContractAddress
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.ContractAddress
	Amount                         *big.Int
}, error) {
	ret := new(struct {
		TokenAddress                   common.ContractAddress
		LenDestinationRecipientAddress uint8
		DestinationChainID             uint8
		ResourceID                     [32]byte
		DestinationRecipientAddress    []byte
		Depositer                      common.ContractAddress
		Amount                         *big.Int
	})
	out := ret
	err := _SRC20Handler.contract.Call(opts, out, "_depositRecords", arg0, arg1)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_SRC20Handler *SRC20HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.ContractAddress
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.ContractAddress
	Amount                         *big.Int
}, error) {
	return _SRC20Handler.Contract.DepositRecords(&_SRC20Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_SRC20Handler *SRC20HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.ContractAddress
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.ContractAddress
	Amount                         *big.Int
}, error) {
	return _SRC20Handler.Contract.DepositRecords(&_SRC20Handler.CallOpts, arg0, arg1)
}

// ResourceIDToCurrency is a free data retrieval call binding the contract method 0xc6a3f31f.
//
// Solidity: function _resourceIDToCurrency(bytes32 ) view returns(string)
func (_SRC20Handler *SRC20HandlerCaller) ResourceIDToCurrency(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "_resourceIDToCurrency", arg0)
	return *ret0, err
}

// ResourceIDToCurrency is a free data retrieval call binding the contract method 0xc6a3f31f.
//
// Solidity: function _resourceIDToCurrency(bytes32 ) view returns(string)
func (_SRC20Handler *SRC20HandlerSession) ResourceIDToCurrency(arg0 [32]byte) (string, error) {
	return _SRC20Handler.Contract.ResourceIDToCurrency(&_SRC20Handler.CallOpts, arg0)
}

// ResourceIDToCurrency is a free data retrieval call binding the contract method 0xc6a3f31f.
//
// Solidity: function _resourceIDToCurrency(bytes32 ) view returns(string)
func (_SRC20Handler *SRC20HandlerCallerSession) ResourceIDToCurrency(arg0 [32]byte) (string, error) {
	return _SRC20Handler.Contract.ResourceIDToCurrency(&_SRC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_SRC20Handler *SRC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "_resourceIDToTokenContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_SRC20Handler *SRC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC20Handler.Contract.ResourceIDToTokenContractAddress(&_SRC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_SRC20Handler *SRC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC20Handler.Contract.ResourceIDToTokenContractAddress(&_SRC20Handler.CallOpts, arg0)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Handler *SRC20HandlerCaller) Bytes32ToString(opts *bind.CallOpts, x [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "bytes32ToString", x)
	return *ret0, err
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Handler *SRC20HandlerSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC20Handler.Contract.Bytes32ToString(&_SRC20Handler.CallOpts, x)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC20Handler *SRC20HandlerCallerSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC20Handler.Contract.Bytes32ToString(&_SRC20Handler.CallOpts, x)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_SRC20Handler *SRC20HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (SRC20HandlerDepositRecord, error) {
	var (
		ret0 = new(SRC20HandlerDepositRecord)
	)
	out := ret0
	err := _SRC20Handler.contract.Call(opts, out, "getDepositRecord", depositNonce, destId)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_SRC20Handler *SRC20HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (SRC20HandlerDepositRecord, error) {
	return _SRC20Handler.Contract.GetDepositRecord(&_SRC20Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_SRC20Handler *SRC20HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (SRC20HandlerDepositRecord, error) {
	return _SRC20Handler.Contract.GetDepositRecord(&_SRC20Handler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) payable returns()
func (_SRC20Handler *SRC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) payable returns()
func (_SRC20Handler *SRC20HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.Deposit(&_SRC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) payable returns()
func (_SRC20Handler *SRC20HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.Deposit(&_SRC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, _recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.contract.Transact(opts, "executeProposal", resourceID, _recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerSession) ExecuteProposal(resourceID [32]byte, _recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.ExecuteProposal(&_SRC20Handler.TransactOpts, resourceID, _recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, _recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.ExecuteProposal(&_SRC20Handler.TransactOpts, resourceID, _recipient, amount)
}

// SetReourceTokenAddress is a paid mutator transaction binding the contract method 0x9707bece.
//
// Solidity: function setReourceTokenAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC20Handler *SRC20HandlerTransactor) SetReourceTokenAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC20Handler.contract.Transact(opts, "setReourceTokenAddress", resourceID, contractAddress)
}

// SetReourceTokenAddress is a paid mutator transaction binding the contract method 0x9707bece.
//
// Solidity: function setReourceTokenAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC20Handler *SRC20HandlerSession) SetReourceTokenAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SetReourceTokenAddress(&_SRC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetReourceTokenAddress is a paid mutator transaction binding the contract method 0x9707bece.
//
// Solidity: function setReourceTokenAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC20Handler *SRC20HandlerTransactorSession) SetReourceTokenAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SetReourceTokenAddress(&_SRC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceCurrency is a paid mutator transaction binding the contract method 0xdfce8258.
//
// Solidity: function setResourceCurrency(bytes32 resourceID, string currency) returns()
func (_SRC20Handler *SRC20HandlerTransactor) SetResourceCurrency(opts *bind.TransactOpts, resourceID [32]byte, currency string) (*types.Transaction, error) {
	return _SRC20Handler.contract.Transact(opts, "setResourceCurrency", resourceID, currency)
}

// SetResourceCurrency is a paid mutator transaction binding the contract method 0xdfce8258.
//
// Solidity: function setResourceCurrency(bytes32 resourceID, string currency) returns()
func (_SRC20Handler *SRC20HandlerSession) SetResourceCurrency(resourceID [32]byte, currency string) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SetResourceCurrency(&_SRC20Handler.TransactOpts, resourceID, currency)
}

// SetResourceCurrency is a paid mutator transaction binding the contract method 0xdfce8258.
//
// Solidity: function setResourceCurrency(bytes32 resourceID, string currency) returns()
func (_SRC20Handler *SRC20HandlerTransactorSession) SetResourceCurrency(resourceID [32]byte, currency string) (*types.Transaction, error) {
	return _SRC20Handler.Contract.SetResourceCurrency(&_SRC20Handler.TransactOpts, resourceID, currency)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5a73b0bf.
//
// Solidity: function withdraw(string currency, address recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, currency string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.contract.Transact(opts, "withdraw", currency, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5a73b0bf.
//
// Solidity: function withdraw(string currency, address recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerSession) Withdraw(currency string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.Withdraw(&_SRC20Handler.TransactOpts, currency, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5a73b0bf.
//
// Solidity: function withdraw(string currency, address recipient, uint256 amount) returns()
func (_SRC20Handler *SRC20HandlerTransactorSession) Withdraw(currency string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SRC20Handler.Contract.Withdraw(&_SRC20Handler.TransactOpts, currency, recipient, amount)
}
