pragma solidity 0.6.4;

import "./ISRC721Handler.sol";

/**
    @title Function used across handler contracts.
    @notice This contract is intended to be used with the Bridge contract.
 */
contract HandlerHelpers is ISRC721Handler {

    address public _bridgeAddress;

    // resourceID => token contract address
    mapping (bytes32 => address) public _resourceIDToContractAddress;

    mapping (bytes32 => string) public _resourceIDToCategory;


    mapping (bytes32 => bytes32) public _resourceIDToFeeResoruceID;

    mapping (string => bytes32) public  _categoryToResourceID;

    mapping(bytes32 => address) public _resourceIDToFeeHandler;


    mapping (address => bool) public _burnList;




    modifier onlyBridge() {
        _onlyBridge();
        _;
    }

    function _onlyBridge() view private {
        require(msg.sender == _bridgeAddress, "sender must be bridge contract");
    }


    function setResourceCategory(bytes32 resourceID, string calldata catg) external override  onlyBridge {

        _resourceIDToCategory[resourceID] = catg;
        _categoryToResourceID[catg] = resourceID;
    }

    function setResourceContractAddress(bytes32 resourceID,address contractAddress) external override onlyBridge {
        _resourceIDToContractAddress[resourceID] = contractAddress;
    }



    function setResourceFeeHandler(bytes32 resourceID,address scr20handler )external override onlyBridge{
        _resourceIDToFeeHandler[resourceID] = scr20handler;
    }

    function setBurnable(address contractAddress) external override onlyBridge{
        _setBurnable(contractAddress);
    }

    function _setBurnable(address contractAddress) internal {
        _burnList[contractAddress] = true;
    }


    /**
        @notice Used to manually release funds from ERC safes.
        @param catg category of ticket  to release.
        @param recipient Address to release tokens to.
        @param tkt Either the ticket of SRC721 to release.
     */
    function withdraw(string calldata catg, address recipient, bytes32 tkt) external  virtual override {}

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
