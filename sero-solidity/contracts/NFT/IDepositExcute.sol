pragma solidity 0.6.4;

/**
    @title Interface for handler contracts that support deposits and deposit executions.
 */
interface IDepositExecute {
    /**
        @notice It is intended that deposit are made using the Bridge contract.
        @param destinationChainID Chain ID deposit is expected to be bridged to.
        @param depositNonce This value is generated as an ID by the Bridge contract.
        @param depositer Address of account making the deposit in the Bridge contract.
     */
    function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer,bytes calldata recipient) external payable;


    function executeProposal(bytes32 resourceID,address _recipient,uint256 ticketId,bytes calldata metaData) external;

}
