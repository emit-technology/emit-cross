pragma solidity 0.6.4;

/**
    @title Interface for handler contracts that support deposits and deposit executions.
    @author ChainSafe Systems.
 */
interface IDepositFTExecute {

    function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes calldata recipient,uint256 amount) external;

    /**
        @notice It is intended that proposals are executed by the Bridge contract.
        @param resourceID resourceID
     */
    function executeProposal(bytes32 resourceID, address recipient,uint256 amount) external;


    function transferFee(bytes32 resourceID, address  recipient,uint256 gasFee,address[] calldata relayers, uint256 relayFee) external;

}

