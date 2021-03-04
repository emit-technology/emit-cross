pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;


import "./ISRC721CrossFee.sol";

contract CrossFee is ISRC721CrossFee {


    address public owner;


    mapping(bytes32=>string) public resourceToFeeCy;

    mapping(bytes32=>uint256) public resourceToMinAmount;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    constructor() public {
        owner = msg.sender;
    }

    function getFeeInfo(bytes32 resourceID) external override view returns(string memory cy,uint256 minAmount) {
        return ( resourceToFeeCy[resourceID],resourceToMinAmount[resourceID]);
    }

    function setFeeInfo(bytes32 resourceID,string  calldata cy,uint256 minAmount) external override onlyOwner{
        resourceToFeeCy[resourceID] = cy;
        resourceToMinAmount[resourceID] = minAmount;
    }


}