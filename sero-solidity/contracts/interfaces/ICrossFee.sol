
pragma solidity 0.6.4;


interface ICrossFee {

    function calCrossFee(bytes32 resourceId,uint256 inputAmount,uint256 gasPrice) external view returns(uint256 crossFee,uint256 gasFee);

    function estimateFee(bytes32 resourceId,uint256 inputAmount) external view returns(uint256 fee);
}

