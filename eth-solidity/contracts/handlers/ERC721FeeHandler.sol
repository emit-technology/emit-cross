pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../ERC20Safe.sol";

contract ERC721FeeHandler is  ERC20Safe {

    address public _bridgeAddress;

    mapping(bytes32 => address) resourceIdToTokenContractAddress;

    modifier onlyBridge() {
        _onlyBridge();
        _;
    }

    function _onlyBridge() view private {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
    }


    constructor(
        address          bridgeAddress
    ) public {

        _bridgeAddress = bridgeAddress;

    }

    function setResourceFeeContract(bytes32 reosurceId,address tokenAddress) external onlyBridge {
        resourceIdToTokenContractAddress[reosurceId] =  tokenAddress;

    }


    function transferFee(bytes32 resourceID, address   gasFeeRecipient,address[] calldata relayers, uint256 amount) external  onlyBridge {

        address tokenAddress = resourceIdToTokenContractAddress[resourceID];
        mintERC20(tokenAddress,gasFeeRecipient, amount);


    }


}
