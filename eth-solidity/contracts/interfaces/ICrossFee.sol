
pragma solidity 0.6.4;


interface ICrossFee {

    function calCrossFee(bytes32 resourceId,uint256 inputAmount,uint256 gasPrice) external view returns(uint256 crossFee,uint256 gasFee);

    function estimateFee(bytes32 resourceId,uint256 inputAmount) external view returns(uint256 fee);
}

interface IFeeHander {

    function setResourceFeeContract(bytes32 reosurceId,address tokenAddress) external;

    function transferFee(bytes32 resourceID, address   gasFeeRecipient,address[] calldata relayers, uint256 amount) external;
}
