pragma solidity 0.6.4;

import "../interfaces/ISRCHandler.sol";

/**
    @title Function used across handler contracts.
    @author ChainSafe Systems.
    @notice This contract is intended to be used with the Bridge contract.
 */
contract HandlerHelpers is ISRCHandler {
    address public _bridgeAddress;

    // resourceID => token contract address
    mapping (bytes32 => address) public _resourceIDToTokenContractAddress;

    mapping (bytes32 => string) public _resourceIDToCurrency;

    mapping (string => bytes32) public  _currencyToResourceID;



    modifier onlyBridge() {
        _onlyBridge();
        _;
    }

    function _onlyBridge() view private {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
    }


    function setResourceCurrency(bytes32 resourceID, string calldata currency) external override  onlyBridge {

        _resourceIDToCurrency[resourceID] = currency;
        _currencyToResourceID[currency] = resourceID;
    }

    function setReourceTokenAddress(bytes32 resourceID,address contractAddress) external override onlyBridge {
        _resourceIDToTokenContractAddress[resourceID] = contractAddress;
    }


    /**
        @notice Used to manually release funds from ERC safes.
        @param currency currency of token  to release.
        @param recipient Address to release tokens to.
        @param amount Either the amount of SRC20 to release.
     */
    function withdraw(string calldata currency, address recipient, uint256 amount) external  virtual override {}

    function equals(string memory a, string memory b) internal pure returns (bool) {
        if (bytes(a).length != bytes(b).length) {
            return false;
        }
        for (uint i = 0; i < bytes(a).length; i ++) {
            if(bytes(a)[i] != bytes(b)[i]) {
                return false;
            }
        }
        return true;
    }
}
