pragma solidity 0.6.4;

interface ISRC721CrossFee {

    function getFeeInfo(bytes32 resourceID) external view returns(string memory cy,uint256 minAmount) ;

    function setFeeInfo(bytes32 resourceID,string  calldata cy,uint256 minAmount) external;

}
