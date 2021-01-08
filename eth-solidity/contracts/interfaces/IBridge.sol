pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

/**
    @title Interface for Bridge contract.
    @author ChainSafe Systems.
 */
interface IBridge {
    /**
        @notice Exposing getter for {_chainID} instead of forcing the use of call.
        @return uint8 The {_chainID} that is currently set for the Bridge contract.
     */
    function _chainID() external returns (uint8);

    function minCrossAmount(bytes32 resourceId) external view returns(uint256);

    function resourceIDToLimit(bytes32 resourceId) external view returns(uint256 minAmount,uint256 maxAmount);


    function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient,uint256 amount) external payable ;


}