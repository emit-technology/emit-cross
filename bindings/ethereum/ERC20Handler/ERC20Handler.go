// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC20HandlerDepositRecord struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
const ERC20HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"gasFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"relayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"transferFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
var ERC20HandlerBin = "0x60806040523480156200001157600080fd5b50604051620026f6380380620026f6833981810160405281019062000037919062000489565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000632565b60405180910390fd5b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001165762000108848281518110620000df57fe5b6020026020010151848381518110620000f457fe5b60200260200101516200016560201b60201c565b8080600101915050620000c4565b5060008090505b81518110156200015a576200014c8282815181106200013857fe5b60200260200101516200025760201b60201c565b80806001019150506200011d565b505050505062000757565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16620002e6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002dd9062000610565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600081519050620003528162000723565b92915050565b600082601f8301126200036a57600080fd5b8151620003816200037b8262000682565b62000654565b91508181835260208401935060208101905083856020840282011115620003a757600080fd5b60005b83811015620003db5781620003c0888262000341565b845260208401935060208301925050600181019050620003aa565b5050505092915050565b600082601f830112620003f757600080fd5b81516200040e6200040882620006ab565b62000654565b915081818352602084019350602081019050838560208402820111156200043457600080fd5b60005b838110156200046857816200044d888262000472565b84526020840193506020830192505060018101905062000437565b5050505092915050565b60008151905062000483816200073d565b92915050565b60008060008060808587031215620004a057600080fd5b6000620004b08782880162000341565b945050602085015167ffffffffffffffff811115620004ce57600080fd5b620004dc87828801620003e5565b935050604085015167ffffffffffffffff811115620004fa57600080fd5b620005088782880162000358565b925050606085015167ffffffffffffffff8111156200052657600080fd5b620005348782880162000358565b91505092959194509250565b60006200054f602483620006d4565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005b7603c83620006d4565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b600060208201905081810360008301526200062b8162000540565b9050919050565b600060208201905081810360008301526200064d81620005a8565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200067857600080fd5b8060405250919050565b600067ffffffffffffffff8211156200069a57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620006c357600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620006f28262000703565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200072e81620006e5565b81146200073a57600080fd5b50565b6200074881620006f9565b81146200075457600080fd5b50565b611f8f80620007676000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806395601f091161008c578063ba484c0911610066578063ba484c091461025f578063c8ba6c871461028f578063d9caed12146102bf578063fe61f0e1146102db576100ea565b806395601f091461020b578063b8fa373614610227578063ba40f28414610243576100ea565b8063318c136e116100c8578063318c136e146101575780634402027f146101755780636a70d081146101ab5780637f79bea8146101db576100ea565b80630396ec10146100ef57806307b7ed991461010b5780630a6d55d814610127575b600080fd5b61010960048036038101906101049190611672565b6102f7565b005b6101256004803603810190610120919061156c565b610437565b005b610141600480360381019061013c919061160d565b61044b565b60405161014e9190611bdb565b60405180910390f35b61015f61047e565b60405161016c9190611bdb565b60405180910390f35b61018f600480360381019061018a9190611836565b6104a3565b6040516101a29796959493929190611c56565b60405180910390f35b6101c560048036038101906101c0919061156c565b6105e4565b6040516101d29190611ccc565b60405180910390f35b6101f560048036038101906101f0919061156c565b610604565b6040516102029190611ccc565b60405180910390f35b61022560048036038101906102209190611595565b610624565b005b610241600480360381019061023c9190611636565b61063b565b005b61025d60048036038101906102589190611753565b610651565b005b610279600480360381019061027491906117fa565b610984565b6040516102869190611da2565b60405180910390f35b6102a960048036038101906102a4919061156c565b610b79565b6040516102b69190611ce7565b60405180910390f35b6102d960048036038101906102d49190611595565b610b91565b005b6102f560048036038101906102f091906116c1565b610ba9565b005b6102ff610d98565b60006001600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ba90611d82565b60405180910390fd5b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561042557610420818484610e29565b610431565b610430818484610ea1565b5b50505050565b61043f610d98565b61044881610eb7565b50565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6005602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060000160159054906101000a900460ff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105ae5780601f10610583576101008083540402835291602001916105ae565b820191906000526020600020905b81548152906001019060200180831161059157829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905087565b60046020528060005260406000206000915054906101000a900460ff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b600083905061063581843085610f9e565b50505050565b610643610d98565b61064d8282611027565b5050565b610659610d98565b60006001600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661071d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071490611d82565b60405180910390fd5b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561077f5761077a818684611119565b61078c565b61078b81863085611191565b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018585905060ff1681526020018860ff16815260200189815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018673ffffffffffffffffffffffffffffffffffffffff16815260200183815250600560008960ff1660ff16815260200190815260200160002060008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160000160156101000a81548160ff021916908360ff160217905550606082015181600101556080820151816002019080519060200190610925929190611343565b5060a08201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c082015181600401559050505050505050505050565b61098c6113c3565b600560008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016000820160159054906101000a900460ff1660ff1660ff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b085780601f10610add57610100808354040283529160200191610b08565b820191906000526020600020905b815481529060010190602001808311610aeb57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481525050905092915050565b60026020528060005260406000206000915090505481565b610b99610d98565b610ba4838383610ea1565b505050565b610bb1610d98565b60006001600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6c90611d82565b60405180910390fd5b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610d2d57610cd2818787610e29565b60008090505b84849050811015610d2757610d1a82868684818110610cf357fe5b9050602002016020810190610d08919061156c565b878790508681610d1457fe5b04610e29565b8080600101915050610cd8565b50610d8f565b610d38818787610ea1565b60008090505b84849050811015610d8d57610d8082868684818110610d5957fe5b9050602002016020810190610d6e919061156c565b878790508681610d7a57fe5b04610ea1565b8080600101915050610d3e565b505b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1e90611d22565b60405180910390fd5b565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b8152600401610e69929190611c2d565b600060405180830381600087803b158015610e8357600080fd5b505af1158015610e97573d6000803e3d6000fd5b5050505050505050565b6000839050610eb18184846111a9565b50505050565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610f43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3a90611d42565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b611021846323b872dd60e01b858585604051602401610fbf93929190611bf6565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061122f565b50505050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b8152600401611159929190611c2d565b600060405180830381600087803b15801561117357600080fd5b505af1158015611187573d6000803e3d6000fd5b5050505050505050565b60008490506111a281858585610f9e565b5050505050565b61122a8363a9059cbb60e01b84846040516024016111c8929190611c2d565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061122f565b505050565b600060608373ffffffffffffffffffffffffffffffffffffffff16836040516112589190611bc4565b6000604051808303816000865af19150503d8060008114611295576040519150601f19603f3d011682016040523d82523d6000602084013e61129a565b606091505b5091509150816112df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d690611d62565b60405180910390fd5b60008151111561133d57808060200190518101906112fd91906115e4565b61133c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133390611d02565b60405180910390fd5b5b50505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061138457805160ff19168380011785556113b2565b828001600101855582156113b2579182015b828111156113b1578251825591602001919060010190611396565b5b5090506113bf9190611435565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600060ff1681526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b61145791905b8082111561145357600081600090555060010161143b565b5090565b90565b60008135905061146981611ecf565b92915050565b60008083601f84011261148157600080fd5b8235905067ffffffffffffffff81111561149a57600080fd5b6020830191508360208202830111156114b257600080fd5b9250929050565b6000815190506114c881611ee6565b92915050565b6000813590506114dd81611efd565b92915050565b60008083601f8401126114f557600080fd5b8235905067ffffffffffffffff81111561150e57600080fd5b60208301915083600182028301111561152657600080fd5b9250929050565b60008135905061153c81611f14565b92915050565b60008135905061155181611f2b565b92915050565b60008135905061156681611f42565b92915050565b60006020828403121561157e57600080fd5b600061158c8482850161145a565b91505092915050565b6000806000606084860312156115aa57600080fd5b60006115b88682870161145a565b93505060206115c98682870161145a565b92505060406115da8682870161152d565b9150509250925092565b6000602082840312156115f657600080fd5b6000611604848285016114b9565b91505092915050565b60006020828403121561161f57600080fd5b600061162d848285016114ce565b91505092915050565b6000806040838503121561164957600080fd5b6000611657858286016114ce565b92505060206116688582860161145a565b9150509250929050565b60008060006060848603121561168757600080fd5b6000611695868287016114ce565b93505060206116a68682870161145a565b92505060406116b78682870161152d565b9150509250925092565b60008060008060008060a087890312156116da57600080fd5b60006116e889828a016114ce565b96505060206116f989828a0161145a565b955050604061170a89828a0161152d565b945050606087013567ffffffffffffffff81111561172757600080fd5b61173389828a0161146f565b9350935050608061174689828a0161152d565b9150509295509295509295565b600080600080600080600060c0888a03121561176e57600080fd5b600061177c8a828b016114ce565b975050602061178d8a828b01611557565b965050604061179e8a828b01611542565b95505060606117af8a828b0161145a565b945050608088013567ffffffffffffffff8111156117cc57600080fd5b6117d88a828b016114e3565b935093505060a06117eb8a828b0161152d565b91505092959891949750929550565b6000806040838503121561180d57600080fd5b600061181b85828601611542565b925050602061182c85828601611557565b9150509250929050565b6000806040838503121561184957600080fd5b600061185785828601611557565b925050602061186885828601611542565b9150509250929050565b61187b81611e18565b82525050565b61188a81611e18565b82525050565b61189981611e2a565b82525050565b6118a881611e36565b82525050565b6118b781611e36565b82525050565b60006118c882611dcf565b6118d28185611dfc565b93506118e2818560208601611e8b565b80840191505092915050565b60006118f982611dc4565b6119038185611dda565b9350611913818560208601611e8b565b61191c81611ebe565b840191505092915050565b600061193282611dc4565b61193c8185611deb565b935061194c818560208601611e8b565b61195581611ebe565b840191505092915050565b600061196d602083611e07565b91507f45524332303a206f7065726174696f6e20646964206e6f7420737563636565646000830152602082019050919050565b60006119ad601e83611e07565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006119ed602483611e07565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611a53601283611e07565b91507f45524332303a2063616c6c206661696c656400000000000000000000000000006000830152602082019050919050565b6000611a93602883611e07565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b600060e083016000830151611b046000860182611872565b506020830151611b176020860182611ba6565b506040830151611b2a6040860182611ba6565b506060830151611b3d606086018261189f565b5060808301518482036080860152611b5582826118ee565b91505060a0830151611b6a60a0860182611872565b5060c0830151611b7d60c0860182611b88565b508091505092915050565b611b9181611e60565b82525050565b611ba081611e60565b82525050565b611baf81611e7e565b82525050565b611bbe81611e7e565b82525050565b6000611bd082846118bd565b915081905092915050565b6000602082019050611bf06000830184611881565b92915050565b6000606082019050611c0b6000830186611881565b611c186020830185611881565b611c256040830184611b97565b949350505050565b6000604082019050611c426000830185611881565b611c4f6020830184611b97565b9392505050565b600060e082019050611c6b600083018a611881565b611c786020830189611bb5565b611c856040830188611bb5565b611c9260608301876118ae565b8181036080830152611ca48186611927565b9050611cb360a0830185611881565b611cc060c0830184611b97565b98975050505050505050565b6000602082019050611ce16000830184611890565b92915050565b6000602082019050611cfc60008301846118ae565b92915050565b60006020820190508181036000830152611d1b81611960565b9050919050565b60006020820190508181036000830152611d3b816119a0565b9050919050565b60006020820190508181036000830152611d5b816119e0565b9050919050565b60006020820190508181036000830152611d7b81611a46565b9050919050565b60006020820190508181036000830152611d9b81611a86565b9050919050565b60006020820190508181036000830152611dbc8184611aec565b905092915050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611e2382611e40565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b83811015611ea9578082015181840152602081019050611e8e565b83811115611eb8576000848401525b50505050565b6000601f19601f8301169050919050565b611ed881611e18565b8114611ee357600080fd5b50565b611eef81611e2a565b8114611efa57600080fd5b50565b611f0681611e36565b8114611f1157600080fd5b50565b611f1d81611e60565b8114611f2857600080fd5b50565b611f3481611e6a565b8114611f3f57600080fd5b50565b611f4b81611e7e565b8114611f5657600080fd5b5056fea264697066735822122004484474acd7079e4f37adc198b8cccc0b05dc11bd1bd73e2e990f9da5734efe64736f6c63430006040033"

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                   common.Address
		LenDestinationRecipientAddress uint8
		DestinationChainID             uint8
		ResourceID                     [32]byte
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		Amount                         *big.Int
	})

	outstruct.TokenAddress = out[0].(common.Address)
	outstruct.LenDestinationRecipientAddress = out[1].(uint8)
	outstruct.DestinationChainID = out[2].(uint8)
	outstruct.ResourceID = out[3].([32]byte)
	outstruct.DestinationRecipientAddress = out[4].([]byte)
	outstruct.Depositer = out[5].(common.Address)
	outstruct.Amount = out[6].(*big.Int)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0, arg1)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC20HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20HandlerDepositRecord)).(*ERC20HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xba40f284.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeProposal", resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0396ec10.
//
// Solidity: function executeProposal(bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, recipientAddress, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_ERC20Handler *ERC20HandlerTransactor) TransferFee(opts *bind.TransactOpts, resourceID [32]byte, gasFeeRecipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "transferFee", resourceID, gasFeeRecipient, gasFee, relayers, relayFee)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_ERC20Handler *ERC20HandlerSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferFee(&_ERC20Handler.TransactOpts, resourceID, gasFeeRecipient, gasFee, relayers, relayFee)
}

// TransferFee is a paid mutator transaction binding the contract method 0xfe61f0e1.
//
// Solidity: function transferFee(bytes32 resourceID, address gasFeeRecipient, uint256 gasFee, address[] relayers, uint256 relayFee) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) TransferFee(resourceID [32]byte, gasFeeRecipient common.Address, gasFee *big.Int, relayers []common.Address, relayFee *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferFee(&_ERC20Handler.TransactOpts, resourceID, gasFeeRecipient, gasFee, relayers, relayFee)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}
