// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SRC721Handler

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

// SRC721HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type SRC721HandlerDepositRecord struct {
	TokenAddress                common.ContractAddress
	DestinationChainID          uint8
	TicketResourceID            [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.ContractAddress
	TokenID                     *big.Int
	MetaData                    []byte
	Amount                      *big.Int
}

// SRC721HandlerABI is the input ABI used to generate the binding from.
const SRC721HandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_categoryToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_ticketResourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToCategory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToFeeCurrency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToFeeHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToFeeResoruceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ticketId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metaData\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_ticketResourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSRC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"catg\",\"type\":\"string\"}],\"name\":\"setResourceCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResourceContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cy\",\"type\":\"string\"}],\"name\":\"setResourceFeeCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"scr20handler\",\"type\":\"address\"}],\"name\":\"setResourceFeeHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tkt\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SRC721HandlerBin is the compiled bytecode used for deploying new contracts.
var SRC721HandlerBin = "0x60806040527f3be6bf24d822bcd6f6348f6f5a5c2d3108f04991ee63e80cde49a8c4746a0ef36008557fcf19eb4256453a4e30b6a06d651f1970c223fb6bd1826a28ed861f0e602db9b86009557f868bd6629e7c2e3d2ccf7b9968fad79b448e7a2bfb3ee20ed1acbc695c3c8b23600a557f7c98e64bd943448b4e24ef8c2cdec7b8b1275970cfe10daf2a9bfa4b04dce905600b557fa6a366f1a72e1aef5d8d52ee240a476f619d15be7bc62d3df37496025b83459f600c557ff1964f6690a0536daa42e5c575091297d2479edcc96f721ad85b95358644d276600d557f9ab0d7c07029f006485cf3468ce7811aa8743b5a108599f6bec9367c50ac6aad600e557fa6cafc6282f61eff9032603a017e652f68410d3d3c69f0a3eeca8f181aec1d17600f557f6800e94e36131c049eaeb631e4530829b0d3d20d5b637c8015a8dc9cedd70aed60105534801561015457600080fd5b50604051611def380380611def83398101604081905261017391610198565b600080546001600160a01b0319166001600160a01b03929092169190911790556101c6565b6000602082840312156101a9578081fd5b81516001600160a01b03811681146101bf578182fd5b9392505050565b611c1a806101d56000396000f3fe6080604052600436106101235760003560e01c806385aa92a7116100a0578063afe609b411610064578063afe609b41461033f578063ba484c091461035f578063c13e91971461038c578063c54c2a11146103ac578063da19cc99146103cc5761012a565b806385aa92a71461029d5780638bdf1591146102b25780638cc795ef146102d25780639201de55146102f257806395d291381461031f5761012a565b806365265cad116100e757806365265cad146101e35780636a70d08114610203578063728e218f14610230578063761e3d11146102505780637cdbdd5d146102705761012a565b806307b7ed991461012f578063318c136e1461015157806338995da91461017c5780633e756d671461018f5780634402027f146101af5761012a565b3661012a57005b600080fd5b34801561013b57600080fd5b5061014f61014a36600461157e565b6103ec565b005b34801561015d57600080fd5b50610166610400565b6040516101739190611939565b60405180910390f35b61014f61018a3660046116af565b61040f565b34801561019b57600080fd5b5061014f6101aa366004611666565b61084e565b3480156101bb57600080fd5b506101cf6101ca3660046118ac565b610898565b604051610173989796959493929190611966565b3480156101ef57600080fd5b506101666101fe3660046115a0565b610a0e565b34801561020f57600080fd5b5061022361021e36600461157e565b610a29565b60405161017391906119d4565b34801561023c57600080fd5b5061014f61024b3660046115ff565b610a3e565b34801561025c57600080fd5b5061014f61026b3660046115d0565b610c95565b34801561027c57600080fd5b5061029061028b3660046115a0565b610ccb565b60405161017391906119df565b3480156102a957600080fd5b50610166610cdd565b3480156102be57600080fd5b5061014f6102cd3660046115d0565b610cec565b3480156102de57600080fd5b5061014f6102ed366004611666565b610d22565b3480156102fe57600080fd5b5061031261030d3660046115a0565b610d49565b6040516101739190611a1e565b34801561032b57600080fd5b5061029061033a366004611788565b610e5f565b34801561034b57600080fd5b5061031261035a3660046115a0565b610e7c565b34801561036b57600080fd5b5061037f61037a366004611878565b610f17565b6040516101739190611abb565b34801561039857600080fd5b506103126103a73660046115a0565b6110d4565b3480156103b857600080fd5b506101666103c73660046115a0565b61113a565b3480156103d857600080fd5b5061014f6103e736600461172e565b611155565b6103f4611159565b6103fd81611185565b50565b6000546001600160a01b031681565b610417611159565b600086815260016020818152604080842054600280845294829020805483516101009682161596909602600019011695909504601f81018490048402850184019092528184526001600160a01b031693606093929091908301828280156104bf5780601f10610494576101008083540402835291602001916104bf565b820191906000526020600020905b8154815290600101906020018083116104a257829003601f168201915b5050505050905060606104d06111a9565b905060006104dc6111ea565b90506104e8838361121c565b61050d5760405162461bcd60e51b815260040161050490611a91565b60405180910390fd5b6001600160a01b038416600090815260076020526040902054849060ff16156105a05761054c6040518060200160405280600081525060008585611298565b806001600160a01b03166344df8e706040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561058757600080fd5b505af115801561059b573d6000803e3d6000fd5b505050505b34156105ec5760008b8152600660205260409020546105d0906001600160a01b03166105ca6112dd565b34611316565b6105ec5760405162461bcd60e51b815260040161050490611a31565b604051806101000160405280866001600160a01b031681526020018b60ff1681526020018c815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160a01b03808b1660208301526040805163200d6fb360e11b81529201919084169063401adf66906106889087906004016119df565b60206040518083038186803b1580156106a057600080fd5b505afa1580156106b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d891906115b8565b8152602001826001600160a01b031663f8ac7f64856040518263ffffffff1660e01b815260040161070991906119df565b60006040518083038186803b15801561072157600080fd5b505afa158015610735573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261075d9190810190611803565b81523460209182015260ff808d1660009081526011835260408082206001600160401b038f1683528452908190208451815486860151909416600160a01b0260ff60a01b196001600160a01b039092166001600160a01b0319909516949094171692909217825583015160018201556060830151805191926107e7926002850192909101906113b0565b5060808201516003820180546001600160a01b0319166001600160a01b0390921691909117905560a0820151600482015560c082015180516108339160058401916020909101906113b0565b5060e082015181600601559050505050505050505050505050565b610856611159565b600083815260026020526040902061086f90838361142e565b508260058383604051610883929190611929565b90815260405190819003602001902055505050565b601160209081526000928352604080842082529183529181902080546001808301546002808501805487516101009582161595909502600019011691909104601f81018890048802840188019096528583526001600160a01b03841696600160a01b90940460ff16959194939091908301828280156109585780601f1061092d57610100808354040283529160200191610958565b820191906000526020600020905b81548152906001019060200180831161093b57829003601f168201915b505050506003830154600484015460058501805460408051602060026101006001861615026000190190941693909304601f810184900484028201840190925281815296976001600160a01b0390951696939550908301828280156109fe5780601f106109d3576101008083540402835291602001916109fe565b820191906000526020600020905b8154815290600101906020018083116109e157829003601f168201915b5050505050908060060154905088565b6006602052600090815260409020546001600160a01b031681565b60076020526000908152604090205460ff1681565b610a46611159565b6000858152600160209081526040808320546001600160a01b031680845260079092528220549091829160ff1615610b62576040516340c10f1960e01b81526001600160a01b038316906340c10f1990610aa6908a908a9060040161194d565b602060405180830381600087803b158015610ac057600080fd5b505af1158015610ad4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af891906115b8565b604051637672bad760e01b81529091506001600160a01b03831690637672bad790610b2b908490899089906004016119e8565b600060405180830381600087803b158015610b4557600080fd5b505af1158015610b59573d6000803e3d6000fd5b50505050610c8b565b604051636094ff9960e11b81526001600160a01b0383169063c129ff3290610b8e9089906004016119df565b60206040518083038186803b158015610ba657600080fd5b505afa158015610bba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bde91906115b8565b600089815260026020818152604092839020805484516001821615610100026000190190911693909304601f8101839004830284018301909452838352939450610c82938b939091830182828015610c775780601f10610c4c57610100808354040283529160200191610c77565b820191906000526020600020905b815481529060010190602001808311610c5a57829003601f168201915b50505050508361133f565b610c8b57600080fd5b5050505050505050565b610c9d611159565b60009182526001602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b60046020526000908152604090205481565b6012546001600160a01b031681565b610cf4611159565b60009182526006602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b610d2a611159565b6000838152600360205260409020610d4390838361142e565b50505050565b604080516020808252818301909252606091600091839160208201818036833701905050905060005b6020811015610dd8576008810260020a85026001600160f81b0319811615610dc35780838581518110610da157fe5b60200101906001600160f81b031916908160001a905350600190930192610dcf565b8315610dcf5750610dd8565b50600101610d72565b506060826040519080825280601f01601f191660200182016040528015610e06576020820181803683370190505b50905060005b83811015610e5657828181518110610e2057fe5b602001015160f81c60f81b828281518110610e3757fe5b60200101906001600160f81b031916908160001a905350600101610e0c565b50949350505050565b805160208183018101805160058252928201919093012091525481565b60036020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610f0f5780601f10610ee457610100808354040283529160200191610f0f565b820191906000526020600020905b815481529060010190602001808311610ef257829003601f168201915b505050505081565b610f1f61149c565b60ff82811660009081526011602090815260408083206001600160401b03881684528252918290208251610100808201855282546001600160a01b0381168352600160a01b90049095168184015260018083015482860152600280840180548751938116159098026000190190971604601f8101859004850282018501909552848152909491936060860193919291830182828015610fff5780601f10610fd457610100808354040283529160200191610fff565b820191906000526020600020905b815481529060010190602001808311610fe257829003601f168201915b505050918352505060038201546001600160a01b03166020808301919091526004830154604080840191909152600584018054825160026101006001841615026000190190921691909104601f8101859004850282018501909352828152606090940193929091908301828280156110b85780601f1061108d576101008083540402835291602001916110b8565b820191906000526020600020905b81548152906001019060200180831161109b57829003601f168201915b5050505050815260200160068201548152505090505b92915050565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f8101839004830284018301909452838352919290830182828015610f0f5780601f10610ee457610100808354040283529160200191610f0f565b6001602052600090815260409020546001600160a01b031681565b610d435b6000546001600160a01b031633146111835760405162461bcd60e51b815260040161050490611a5a565b565b6001600160a01b03166000908152600760205260409020805460ff19166001179055565b6040805160208082528183019092526060918291906020820181803683370190505090506000600d54602083a15080516111e281610d49565b925050505b90565b60408051602080825281830190925260009160609190602082018180368337019050509050600e54602082a151905090565b6000815183511461122f575060006110ce565b60005b835181101561128e5782818151811061124757fe5b602001015160f81c60f81b6001600160f81b03191684828151811061126857fe5b01602001516001600160f81b031916146112865760009150506110ce565b600101611232565b5060019392505050565b60408051608080825260a08201909252606091602082018180368337019050509050848152836020820152826040820152816060820152600f54608082a15050505050565b6040805160208082528183019092526060918291906020820181803683370190505090506000600b54602083a15080516111e281610d49565b6000611337848484604051806020016040528060008152506000801b61135a565b949350505050565b60006113378460405180602001604052806000815250600086865b6040805160a080825260c0820190925260009160609190602082018180368337019050509050868152856020820152846040820152836060820152826080820152600a5460a082a1608001519695505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113f157805160ff191683800117855561141e565b8280016001018555821561141e579182015b8281111561141e578251825591602001919060010190611403565b5061142a9291506114df565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061146f5782800160ff1982351617855561141e565b8280016001018555821561141e579182015b8281111561141e578235825591602001919060010190611481565b604080516101008101825260008082526020820181905291810182905260608082018190526080820183905260a0820183905260c082015260e081019190915290565b6111e791905b8082111561142a57600081556001016114e5565b80356001600160a01b03811681146110ce57600080fd5b60008083601f840112611521578182fd5b5081356001600160401b03811115611537578182fd5b60208301915083602082850101111561154f57600080fd5b9250929050565b80356001600160401b03811681146110ce57600080fd5b803560ff811681146110ce57600080fd5b60006020828403121561158f578081fd5b61159983836114f9565b9392505050565b6000602082840312156115b1578081fd5b5035919050565b6000602082840312156115c9578081fd5b5051919050565b600080604083850312156115e2578081fd5b8235915060208301356115f481611bcf565b809150509250929050565b600080600080600060808688031215611616578081fd5b85359450602086013561162881611bcf565b93506040860135925060608601356001600160401b03811115611649578182fd5b61165588828901611510565b969995985093965092949392505050565b60008060006040848603121561167a578283fd5b8335925060208401356001600160401b03811115611696578283fd5b6116a286828701611510565b9497909650939450505050565b60008060008060008060a087890312156116c7578081fd5b863595506116d8886020890161156d565b94506116e78860408901611556565b93506116f688606089016114f9565b925060808701356001600160401b03811115611710578182fd5b61171c89828a01611510565b979a9699509497509295939492505050565b60008060008060608587031215611743578384fd5b84356001600160401b03811115611758578485fd5b61176487828801611510565b9095509350611778905086602087016114f9565b9396929550929360400135925050565b600060208284031215611799578081fd5b81356001600160401b038111156117ae578182fd5b80830184601f8201126117bf578283fd5b803591506117d46117cf83611b80565b611b5a565b8281528560208484010111156117e8578384fd5b82602083016020830137918201602001929092529392505050565b600060208284031215611814578081fd5b81516001600160401b03811115611829578182fd5b80830184601f82011261183a578283fd5b8051915061184a6117cf83611b80565b82815285602084840101111561185e578384fd5b61186f836020830160208501611ba3565b95945050505050565b6000806040838503121561188a578182fd5b6118948484611556565b91506118a3846020850161156d565b90509250929050565b600080604083850312156118be578182fd5b823560ff811681146118ce578283fd5b915060208301356001600160401b03811681146115f4578182fd5b6001600160a01b03169052565b6000815180845261190e816020860160208601611ba3565b601f01601f19169290920160200192915050565b60ff169052565b6000828483379101908152919050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b03898116825260ff89166020830152604082018890526101006060830181905260009161199c8483018a6118f6565b81891660808601528760a086015284810360c08601526119bc81886118f6565b93505050508260e08301529998505050505050505050565b901515815260200190565b90815260200190565b60008482526040602083015282604083015282846060840137818301606090810191909152601f909201601f1916010192915050565b60006020825261159960208301846118f6565b6020808252600f908201526e1cd95b99081999594819985a5b1959608a1b604082015260600190565b6020808252601e908201527f73656e646572206d7573742062652062726964676520636f6e74726163740000604082015260600190565b60208082526010908201526f696e76616c69642063617465676f727960801b604082015260600190565b602080825282516001600160a01b031682820152820151600090611ae26040840182611922565b50604083015160608301526060830151610100806080850152611b096101208501836118f6565b60808601519250611b1d60a08601846118e9565b60a086015160c086015260c08601519250601f198582030160e0860152611b4481846118f6565b60e0870151838701528094505050505092915050565b6040518181016001600160401b0381118282101715611b7857600080fd5b604052919050565b60006001600160401b03821115611b95578081fd5b50601f01601f191660200190565b60005b83811015611bbe578181015183820152602001611ba6565b83811115610d435750506000910152565b6001600160a01b03811681146103fd57600080fdfea264697066735822122009622463f73b5c773689e08fee48edc061e6b4b7be212f006e89ab9549388b3064736f6c63430006040033"

// DeploySRC721Handler deploys a new Ethereum contract, binding an instance of SRC721Handler to it.
func DeploySRC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *SRC721Handler, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC721HandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SRC721HandlerBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SRC721Handler{SRC721HandlerCaller: SRC721HandlerCaller{contract: contract}, SRC721HandlerTransactor: SRC721HandlerTransactor{contract: contract}, SRC721HandlerFilterer: SRC721HandlerFilterer{contract: contract}}, nil
}

// SRC721Handler is an auto generated Go binding around an Ethereum contract.
type SRC721Handler struct {
	SRC721HandlerCaller     // Read-only binding to the contract
	SRC721HandlerTransactor // Write-only binding to the contract
	SRC721HandlerFilterer   // Log filterer for contract events
}

// SRC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SRC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SRC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SRC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SRC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SRC721HandlerSession struct {
	Contract     *SRC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SRC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SRC721HandlerCallerSession struct {
	Contract *SRC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SRC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SRC721HandlerTransactorSession struct {
	Contract     *SRC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SRC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SRC721HandlerRaw struct {
	Contract *SRC721Handler // Generic contract binding to access the raw methods on
}

// SRC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SRC721HandlerCallerRaw struct {
	Contract *SRC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// SRC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SRC721HandlerTransactorRaw struct {
	Contract *SRC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSRC721Handler creates a new instance of SRC721Handler, bound to a specific deployed contract.
func NewSRC721Handler(address common.Address, backend bind.ContractBackend) (*SRC721Handler, error) {
	contract, err := bindSRC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SRC721Handler{SRC721HandlerCaller: SRC721HandlerCaller{contract: contract}, SRC721HandlerTransactor: SRC721HandlerTransactor{contract: contract}, SRC721HandlerFilterer: SRC721HandlerFilterer{contract: contract}}, nil
}

// NewSRC721HandlerCaller creates a new read-only instance of SRC721Handler, bound to a specific deployed contract.
func NewSRC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*SRC721HandlerCaller, error) {
	contract, err := bindSRC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SRC721HandlerCaller{contract: contract}, nil
}

// NewSRC721HandlerTransactor creates a new write-only instance of SRC721Handler, bound to a specific deployed contract.
func NewSRC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*SRC721HandlerTransactor, error) {
	contract, err := bindSRC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SRC721HandlerTransactor{contract: contract}, nil
}

// NewSRC721HandlerFilterer creates a new log filterer instance of SRC721Handler, bound to a specific deployed contract.
func NewSRC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*SRC721HandlerFilterer, error) {
	contract, err := bindSRC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SRC721HandlerFilterer{contract: contract}, nil
}

// bindSRC721Handler binds a generic wrapper to an already deployed contract.
func bindSRC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SRC721HandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC721Handler *SRC721HandlerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC721Handler.Contract.SRC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC721Handler *SRC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SRC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC721Handler *SRC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SRC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SRC721Handler *SRC721HandlerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SRC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SRC721Handler *SRC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SRC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SRC721Handler *SRC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SRC721Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_bridgeAddress")
	return *ret0, err
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerSession) BridgeAddress() (common.ContractAddress, error) {
	return _SRC721Handler.Contract.BridgeAddress(&_SRC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerCallerSession) BridgeAddress() (common.ContractAddress, error) {
	return _SRC721Handler.Contract.BridgeAddress(&_SRC721Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_SRC721Handler *SRC721HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_burnList", arg0)
	return *ret0, err
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_SRC721Handler *SRC721HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _SRC721Handler.Contract.BurnList(&_SRC721Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_SRC721Handler *SRC721HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _SRC721Handler.Contract.BurnList(&_SRC721Handler.CallOpts, arg0)
}

// CategoryToResourceID is a free data retrieval call binding the contract method 0x95d29138.
//
// Solidity: function _categoryToResourceID(string ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerCaller) CategoryToResourceID(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_categoryToResourceID", arg0)
	return *ret0, err
}

// CategoryToResourceID is a free data retrieval call binding the contract method 0x95d29138.
//
// Solidity: function _categoryToResourceID(string ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerSession) CategoryToResourceID(arg0 string) ([32]byte, error) {
	return _SRC721Handler.Contract.CategoryToResourceID(&_SRC721Handler.CallOpts, arg0)
}

// CategoryToResourceID is a free data retrieval call binding the contract method 0x95d29138.
//
// Solidity: function _categoryToResourceID(string ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerCallerSession) CategoryToResourceID(arg0 string) ([32]byte, error) {
	return _SRC721Handler.Contract.CategoryToResourceID(&_SRC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _ticketResourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData, uint256 amount)
func (_SRC721Handler *SRC721HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.ContractAddress
	DestinationChainID          uint8
	TicketResourceID            [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.ContractAddress
	TokenID                     *big.Int
	MetaData                    []byte
	Amount                      *big.Int
}, error) {
	ret := new(struct {
		TokenAddress                common.ContractAddress
		DestinationChainID          uint8
		TicketResourceID            [32]byte
		DestinationRecipientAddress []byte
		Depositer                   common.ContractAddress
		TokenID                     *big.Int
		MetaData                    []byte
		Amount                      *big.Int
	})
	out := ret
	err := _SRC721Handler.contract.Call(opts, out, "_depositRecords", arg0, arg1)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _ticketResourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData, uint256 amount)
func (_SRC721Handler *SRC721HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.ContractAddress
	DestinationChainID          uint8
	TicketResourceID            [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.ContractAddress
	TokenID                     *big.Int
	MetaData                    []byte
	Amount                      *big.Int
}, error) {
	return _SRC721Handler.Contract.DepositRecords(&_SRC721Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _destinationChainID, bytes32 _ticketResourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData, uint256 amount)
func (_SRC721Handler *SRC721HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                common.ContractAddress
	DestinationChainID          uint8
	TicketResourceID            [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.ContractAddress
	TokenID                     *big.Int
	MetaData                    []byte
	Amount                      *big.Int
}, error) {
	return _SRC721Handler.Contract.DepositRecords(&_SRC721Handler.CallOpts, arg0, arg1)
}

// ResourceIDToCategory is a free data retrieval call binding the contract method 0xc13e9197.
//
// Solidity: function _resourceIDToCategory(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerCaller) ResourceIDToCategory(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_resourceIDToCategory", arg0)
	return *ret0, err
}

// ResourceIDToCategory is a free data retrieval call binding the contract method 0xc13e9197.
//
// Solidity: function _resourceIDToCategory(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerSession) ResourceIDToCategory(arg0 [32]byte) (string, error) {
	return _SRC721Handler.Contract.ResourceIDToCategory(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToCategory is a free data retrieval call binding the contract method 0xc13e9197.
//
// Solidity: function _resourceIDToCategory(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerCallerSession) ResourceIDToCategory(arg0 [32]byte) (string, error) {
	return _SRC721Handler.Contract.ResourceIDToCategory(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_resourceIDToContractAddress", arg0)
	return *ret0, err
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC721Handler.Contract.ResourceIDToContractAddress(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC721Handler.Contract.ResourceIDToContractAddress(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeCurrency is a free data retrieval call binding the contract method 0xafe609b4.
//
// Solidity: function _resourceIDToFeeCurrency(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerCaller) ResourceIDToFeeCurrency(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_resourceIDToFeeCurrency", arg0)
	return *ret0, err
}

// ResourceIDToFeeCurrency is a free data retrieval call binding the contract method 0xafe609b4.
//
// Solidity: function _resourceIDToFeeCurrency(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerSession) ResourceIDToFeeCurrency(arg0 [32]byte) (string, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeCurrency(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeCurrency is a free data retrieval call binding the contract method 0xafe609b4.
//
// Solidity: function _resourceIDToFeeCurrency(bytes32 ) view returns(string)
func (_SRC721Handler *SRC721HandlerCallerSession) ResourceIDToFeeCurrency(arg0 [32]byte) (string, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeCurrency(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeHandler is a free data retrieval call binding the contract method 0x65265cad.
//
// Solidity: function _resourceIDToFeeHandler(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerCaller) ResourceIDToFeeHandler(opts *bind.CallOpts, arg0 [32]byte) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_resourceIDToFeeHandler", arg0)
	return *ret0, err
}

// ResourceIDToFeeHandler is a free data retrieval call binding the contract method 0x65265cad.
//
// Solidity: function _resourceIDToFeeHandler(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerSession) ResourceIDToFeeHandler(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeHandler(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeHandler is a free data retrieval call binding the contract method 0x65265cad.
//
// Solidity: function _resourceIDToFeeHandler(bytes32 ) view returns(address)
func (_SRC721Handler *SRC721HandlerCallerSession) ResourceIDToFeeHandler(arg0 [32]byte) (common.ContractAddress, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeHandler(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeResoruceID is a free data retrieval call binding the contract method 0x7cdbdd5d.
//
// Solidity: function _resourceIDToFeeResoruceID(bytes32 ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerCaller) ResourceIDToFeeResoruceID(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "_resourceIDToFeeResoruceID", arg0)
	return *ret0, err
}

// ResourceIDToFeeResoruceID is a free data retrieval call binding the contract method 0x7cdbdd5d.
//
// Solidity: function _resourceIDToFeeResoruceID(bytes32 ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerSession) ResourceIDToFeeResoruceID(arg0 [32]byte) ([32]byte, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeResoruceID(&_SRC721Handler.CallOpts, arg0)
}

// ResourceIDToFeeResoruceID is a free data retrieval call binding the contract method 0x7cdbdd5d.
//
// Solidity: function _resourceIDToFeeResoruceID(bytes32 ) view returns(bytes32)
func (_SRC721Handler *SRC721HandlerCallerSession) ResourceIDToFeeResoruceID(arg0 [32]byte) ([32]byte, error) {
	return _SRC721Handler.Contract.ResourceIDToFeeResoruceID(&_SRC721Handler.CallOpts, arg0)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC721Handler *SRC721HandlerCaller) Bytes32ToString(opts *bind.CallOpts, x [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "bytes32ToString", x)
	return *ret0, err
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC721Handler *SRC721HandlerSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC721Handler.Contract.Bytes32ToString(&_SRC721Handler.CallOpts, x)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 x) pure returns(string)
func (_SRC721Handler *SRC721HandlerCallerSession) Bytes32ToString(x [32]byte) (string, error) {
	return _SRC721Handler.Contract.Bytes32ToString(&_SRC721Handler.CallOpts, x)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes,uint256))
func (_SRC721Handler *SRC721HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (SRC721HandlerDepositRecord, error) {
	var (
		ret0 = new(SRC721HandlerDepositRecord)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "getDepositRecord", depositNonce, destId)
	return *ret0, err
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes,uint256))
func (_SRC721Handler *SRC721HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (SRC721HandlerDepositRecord, error) {
	return _SRC721Handler.Contract.GetDepositRecord(&_SRC721Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,bytes32,bytes,address,uint256,bytes,uint256))
func (_SRC721Handler *SRC721HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (SRC721HandlerDepositRecord, error) {
	return _SRC721Handler.Contract.GetDepositRecord(&_SRC721Handler.CallOpts, depositNonce, destId)
}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerCaller) StorageAddress(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SRC721Handler.contract.Call(opts, out, "storageAddress")
	return *ret0, err
}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerSession) StorageAddress() (common.ContractAddress, error) {
	return _SRC721Handler.Contract.StorageAddress(&_SRC721Handler.CallOpts)
}

// StorageAddress is a free data retrieval call binding the contract method 0x85aa92a7.
//
// Solidity: function storageAddress() view returns(address)
func (_SRC721Handler *SRC721HandlerCallerSession) StorageAddress() (common.ContractAddress, error) {
	return _SRC721Handler.Contract.StorageAddress(&_SRC721Handler.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient) payable returns()
func (_SRC721Handler *SRC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient) payable returns()
func (_SRC721Handler *SRC721HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.Deposit(&_SRC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes recipient) payable returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, recipient []byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.Deposit(&_SRC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, recipient)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 ticketId, bytes metaData) returns()
func (_SRC721Handler *SRC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, _recipient common.Address, ticketId *big.Int, metaData []byte) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "executeProposal", resourceID, _recipient, ticketId, metaData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 ticketId, bytes metaData) returns()
func (_SRC721Handler *SRC721HandlerSession) ExecuteProposal(resourceID [32]byte, _recipient common.Address, ticketId *big.Int, metaData []byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.ExecuteProposal(&_SRC721Handler.TransactOpts, resourceID, _recipient, ticketId, metaData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x728e218f.
//
// Solidity: function executeProposal(bytes32 resourceID, address _recipient, uint256 ticketId, bytes metaData) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, _recipient common.Address, ticketId *big.Int, metaData []byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.ExecuteProposal(&_SRC721Handler.TransactOpts, resourceID, _recipient, ticketId, metaData)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetBurnable(&_SRC721Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetBurnable(&_SRC721Handler.TransactOpts, contractAddress)
}

// SetResourceCategory is a paid mutator transaction binding the contract method 0x3e756d67.
//
// Solidity: function setResourceCategory(bytes32 resourceID, string catg) returns()
func (_SRC721Handler *SRC721HandlerTransactor) SetResourceCategory(opts *bind.TransactOpts, resourceID [32]byte, catg string) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "setResourceCategory", resourceID, catg)
}

// SetResourceCategory is a paid mutator transaction binding the contract method 0x3e756d67.
//
// Solidity: function setResourceCategory(bytes32 resourceID, string catg) returns()
func (_SRC721Handler *SRC721HandlerSession) SetResourceCategory(resourceID [32]byte, catg string) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceCategory(&_SRC721Handler.TransactOpts, resourceID, catg)
}

// SetResourceCategory is a paid mutator transaction binding the contract method 0x3e756d67.
//
// Solidity: function setResourceCategory(bytes32 resourceID, string catg) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) SetResourceCategory(resourceID [32]byte, catg string) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceCategory(&_SRC721Handler.TransactOpts, resourceID, catg)
}

// SetResourceContractAddress is a paid mutator transaction binding the contract method 0x761e3d11.
//
// Solidity: function setResourceContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerTransactor) SetResourceContractAddress(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "setResourceContractAddress", resourceID, contractAddress)
}

// SetResourceContractAddress is a paid mutator transaction binding the contract method 0x761e3d11.
//
// Solidity: function setResourceContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerSession) SetResourceContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceContractAddress(&_SRC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceContractAddress is a paid mutator transaction binding the contract method 0x761e3d11.
//
// Solidity: function setResourceContractAddress(bytes32 resourceID, address contractAddress) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) SetResourceContractAddress(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceContractAddress(&_SRC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResourceFeeCurrency is a paid mutator transaction binding the contract method 0x8cc795ef.
//
// Solidity: function setResourceFeeCurrency(bytes32 resourceID, string cy) returns()
func (_SRC721Handler *SRC721HandlerTransactor) SetResourceFeeCurrency(opts *bind.TransactOpts, resourceID [32]byte, cy string) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "setResourceFeeCurrency", resourceID, cy)
}

// SetResourceFeeCurrency is a paid mutator transaction binding the contract method 0x8cc795ef.
//
// Solidity: function setResourceFeeCurrency(bytes32 resourceID, string cy) returns()
func (_SRC721Handler *SRC721HandlerSession) SetResourceFeeCurrency(resourceID [32]byte, cy string) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceFeeCurrency(&_SRC721Handler.TransactOpts, resourceID, cy)
}

// SetResourceFeeCurrency is a paid mutator transaction binding the contract method 0x8cc795ef.
//
// Solidity: function setResourceFeeCurrency(bytes32 resourceID, string cy) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) SetResourceFeeCurrency(resourceID [32]byte, cy string) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceFeeCurrency(&_SRC721Handler.TransactOpts, resourceID, cy)
}

// SetResourceFeeHandler is a paid mutator transaction binding the contract method 0x8bdf1591.
//
// Solidity: function setResourceFeeHandler(bytes32 resourceID, address scr20handler) returns()
func (_SRC721Handler *SRC721HandlerTransactor) SetResourceFeeHandler(opts *bind.TransactOpts, resourceID [32]byte, scr20handler common.Address) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "setResourceFeeHandler", resourceID, scr20handler)
}

// SetResourceFeeHandler is a paid mutator transaction binding the contract method 0x8bdf1591.
//
// Solidity: function setResourceFeeHandler(bytes32 resourceID, address scr20handler) returns()
func (_SRC721Handler *SRC721HandlerSession) SetResourceFeeHandler(resourceID [32]byte, scr20handler common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceFeeHandler(&_SRC721Handler.TransactOpts, resourceID, scr20handler)
}

// SetResourceFeeHandler is a paid mutator transaction binding the contract method 0x8bdf1591.
//
// Solidity: function setResourceFeeHandler(bytes32 resourceID, address scr20handler) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) SetResourceFeeHandler(resourceID [32]byte, scr20handler common.Address) (*types.Transaction, error) {
	return _SRC721Handler.Contract.SetResourceFeeHandler(&_SRC721Handler.TransactOpts, resourceID, scr20handler)
}

// Withdraw is a paid mutator transaction binding the contract method 0xda19cc99.
//
// Solidity: function withdraw(string category, address recipient, bytes32 tkt) returns()
func (_SRC721Handler *SRC721HandlerTransactor) Withdraw(opts *bind.TransactOpts, category string, recipient common.Address, tkt [32]byte) (*types.Transaction, error) {
	return _SRC721Handler.contract.Transact(opts, "withdraw", category, recipient, tkt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xda19cc99.
//
// Solidity: function withdraw(string category, address recipient, bytes32 tkt) returns()
func (_SRC721Handler *SRC721HandlerSession) Withdraw(category string, recipient common.Address, tkt [32]byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.Withdraw(&_SRC721Handler.TransactOpts, category, recipient, tkt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xda19cc99.
//
// Solidity: function withdraw(string category, address recipient, bytes32 tkt) returns()
func (_SRC721Handler *SRC721HandlerTransactorSession) Withdraw(category string, recipient common.Address, tkt [32]byte) (*types.Transaction, error) {
	return _SRC721Handler.Contract.Withdraw(&_SRC721Handler.TransactOpts, category, recipient, tkt)
}
