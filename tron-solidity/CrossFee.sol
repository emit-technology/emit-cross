pragma solidity 0.5.14;

pragma experimental ABIEncoderV2;

import "./utils/SafeMath.sol";

import "./interfaces/ICrossFee.sol";

contract CrossFee is ICrossFee {

    using SafeMath for uint256;

    address public owner;

    uint256 public MAX_GAS_PRICE = 80 * 1e9;

    uint256 public crossInGas;

    mapping(bytes32=>uint256) public resourceToFeeRate;

    mapping(bytes32=>uint256) public resourceToDefalutGasFee;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    constructor(uint256 crossInGas_) public {
        crossInGas = crossInGas_;
        owner = msg.sender;
    }


    function setResource(bytes32 resourceId,uint256 feeRate,uint256 defaultGasFee) public{
        require(msg.sender == owner,"not owner");
        resourceToFeeRate[resourceId] = feeRate;
        resourceToDefalutGasFee[resourceId] = defaultGasFee;
    }

    function setMaxGasPrice(uint256 max_) public onlyOwner {
        MAX_GAS_PRICE = max_;
    }

    function setDefaultGasFee(bytes32 resourceId,uint256 gasFee_) public onlyOwner {
        resourceToDefalutGasFee[resourceId] = gasFee_;
    }

    function setCrossFeeRate(bytes32 resourceId,uint256 feeRate_) public onlyOwner{
        resourceToFeeRate[resourceId] = feeRate_;
    }


    function calCrossFee(bytes32 resourceId,uint256 inputAmount,uint256 gasPrice) public  view returns(uint256 relayerFee,uint256 gasFee){
        uint256 feeRate = resourceToFeeRate[resourceId];
        relayerFee = inputAmount.mul(feeRate).div(10000);
        gasFee = resourceToDefalutGasFee[resourceId];
    }



    function estimateFee(bytes32 resourceId,uint256 inputAmount) public  view returns(uint256 fee){

        uint256 feeRate = resourceToFeeRate[resourceId];
        uint256 relayerFee = inputAmount.mul(feeRate).div(10000);
        uint256 gasFee = resourceToDefalutGasFee[resourceId];
        return relayerFee.add(gasFee);
    }


}