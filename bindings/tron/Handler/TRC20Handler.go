package Handler

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const TRC20HandlerABI = `
[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "bridgeAddress",
				"type": "address"
			},
			{
				"internalType": "bytes32[]",
				"name": "initialResourceIDs",
				"type": "bytes32[]"
			},
			{
				"internalType": "address[]",
				"name": "initialContractAddresses",
				"type": "address[]"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "_bridgeAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "_contractWhitelist",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			},
			{
				"internalType": "uint64",
				"name": "",
				"type": "uint64"
			}
		],
		"name": "_depositRecords",
		"outputs": [
			{
				"internalType": "address",
				"name": "_tokenAddress",
				"type": "address"
			},
			{
				"internalType": "uint8",
				"name": "_destinationChainID",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "_resourceID",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "_destinationRecipientAddress",
				"type": "bytes"
			},
			{
				"internalType": "address",
				"name": "_depositer",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_amount",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"name": "_resourceIDToTokenContractAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "_tokenContractAddressToResourceID",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "resourceID",
				"type": "bytes32"
			},
			{
				"internalType": "uint8",
				"name": "destinationChainID",
				"type": "uint8"
			},
			{
				"internalType": "uint64",
				"name": "depositNonce",
				"type": "uint64"
			},
			{
				"internalType": "address",
				"name": "depositer",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "recipientAddress",
				"type": "bytes"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "deposit",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "resourceID",
				"type": "bytes32"
			},
			{
				"internalType": "address",
				"name": "recipientAddress",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "executeProposal",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "address",
				"name": "tokenAddress",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "fundERC20",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "uint64",
				"name": "depositNonce",
				"type": "uint64"
			},
			{
				"internalType": "uint8",
				"name": "destId",
				"type": "uint8"
			}
		],
		"name": "getDepositRecord",
		"outputs": [
			{
				"components": [
					{
						"internalType": "address",
						"name": "_tokenAddress",
						"type": "address"
					},
					{
						"internalType": "uint8",
						"name": "_destinationChainID",
						"type": "uint8"
					},
					{
						"internalType": "bytes32",
						"name": "_resourceID",
						"type": "bytes32"
					},
					{
						"internalType": "bytes",
						"name": "_destinationRecipientAddress",
						"type": "bytes"
					},
					{
						"internalType": "address",
						"name": "_depositer",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "_amount",
						"type": "uint256"
					}
				],
				"internalType": "struct ERC20Handler.DepositRecord",
				"name": "",
				"type": "tuple"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "resourceID",
				"type": "bytes32"
			},
			{
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			}
		],
		"name": "setResource",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "resourceID",
				"type": "bytes32"
			},
			{
				"internalType": "address",
				"name": "gasFeeRecipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "gasFee",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "relayers",
				"type": "address[]"
			},
			{
				"internalType": "uint256",
				"name": "relayFee",
				"type": "uint256"
			}
		],
		"name": "transferFee",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "address",
				"name": "tokenAddress",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "withdraw",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
`

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type TRC20HandlerDepositRecord struct {
	TokenAddress                common.Address
	DestinationChainID          uint8
	ResourceID                  [32]byte
	DestinationRecipientAddress []byte
	Depositer                   common.Address
	Amount                      *big.Int
}

type Trc20HandlerABI struct {
	abi abi.ABI
}

func NewTrc20HanderABI() *Trc20HandlerABI {
	parsed, err := abi.JSON(strings.NewReader(TRC20HandlerABI))
	if err != nil {
		panic(err)
	}
	return &Trc20HandlerABI{parsed}
}

func (t *Trc20HandlerABI) GetDepositRecord(output []byte) (TRC20HandlerDepositRecord, error) {
	res, err := t.abi.Unpack("getDepositRecord", output)
	if err != nil {
		return *new(TRC20HandlerDepositRecord), err
	}

	//err = t.abi.UnpackIntoInterface(res[0], "getDepositRecord", output)
	//if err != nil {
	//	return *new(TRC20HandlerDepositRecord), err
	//}
	out0 := *abi.ConvertType(res[0], new(TRC20HandlerDepositRecord)).(*TRC20HandlerDepositRecord)

	return out0, nil
}
