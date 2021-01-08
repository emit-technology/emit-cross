pragma solidity 0.6.4;

interface IEUSDTAgent {

    function tokenRate() external view returns(uint256 eusdtAmount,uint256 seroAmount);

    function  transfer(address to) external payable;

    function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient) external payable ;

}