pragma solidity 0.6.4;

/**
    @title Interface to be used with handlers that support ERC20s and ERC721s.
    @author ChainSafe Systems.
 */
interface ISRCHandler {

    function setResourceCurrency(bytes32 resourceID, string calldata currency) external;

    function setReourceTokenAddress(bytes32 resourceID,address contractAddress) external;

    function withdraw(string calldata currency, address recipient, uint256 amountOrTokenID) external;
}
