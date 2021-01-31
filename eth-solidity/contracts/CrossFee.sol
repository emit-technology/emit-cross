pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import 'github.com/Uniswap/uniswap-v2-core//contracts/interfaces/IUniswapV2Pair.sol';
import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/math/SafeMath.sol";

import "./interfaces/ICrossFee.sol";

contract CrossFee is ICrossFee {

    using SafeMath for uint256;

    address public owner;

    address public weth;

    address public uni_factory;

    uint256 public MAX_GAS_PRICE = 80 * 1e9;

    uint256 public crossInGas;

    mapping(bytes32=>address) public resourceToTokenAddrss;
    mapping(bytes32=>uint256) public resourceToFeeRate;
    mapping(bytes32=>uint256) public resourceToDefalutGasFee;

    constructor(address weth_,address uni_,uint256 crossInGas_) public {
        weth = weth_;
        uni_factory = uni_;
        owner = msg.sender;
        crossInGas = crossInGas_;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }


    function setWETH(address weth_) public {
        require(msg.sender == owner,"not owner");
        weth = weth_;
    }

    function setUni(address uni_) public {
        require(msg.sender == owner,"not owner");
        uni_factory = uni_;
    }

    function setResource(bytes32 resourceId,address tokenAddress,uint256 defaltGasFee,uint256 feeRate) public{
        require(msg.sender == owner,"not owner");
        resourceToTokenAddrss[resourceId] = tokenAddress;
        resourceToFeeRate[resourceId] = feeRate;
        resourceToDefalutGasFee[resourceId] = defaltGasFee;
    }

    function setMaxGasPrice(uint256 max_) public onlyOwner{
        MAX_GAS_PRICE = max_;
    }

    function setCrossInGas(uint256 crossInGas_) public onlyOwner {
        crossInGas = crossInGas_;
    }

    function setDefaultGasFee(bytes32 resourceId,uint256 gasFee_) public onlyOwner{
        resourceToDefalutGasFee[resourceId] = gasFee_;
    }

    function setCorssFeeRate(bytes32 resourceId,uint256 feeRate_) public onlyOwner{
        resourceToFeeRate[resourceId] = feeRate_;
    }


    function calCrossFee(bytes32 resourceId,uint256 inputAmount,uint256 gasPrice) public override view returns(uint256 relayerFee,uint256 gasFee){
        return _calFee(resourceId,inputAmount,gasPrice);
    }

    function _calFee(bytes32 resourceId,uint256 inputAmount,uint256 gasPrice) internal view returns(uint256 relayerFee,uint256 gasFee){
        uint256 feeRate = resourceToFeeRate[resourceId];
        relayerFee = inputAmount.mul(feeRate).div(10000);
        gasFee = resourceToDefalutGasFee[resourceId];

        if (gasFee == 0){
            if (gasPrice > MAX_GAS_PRICE) {
                gasPrice = MAX_GAS_PRICE;
            }
            uint256 ethFee = crossInGas.mul(gasPrice);
            address tokenAddress = resourceToTokenAddrss[resourceId];
            if (weth==tokenAddress){
                gasFee = ethFee;
            }else{
                (uint reserveA, uint reserveB) =getReserves(uni_factory,weth,tokenAddress);
                gasFee = quote(ethFee,reserveA,reserveB);
            }

        }
        return (relayerFee,gasFee);
    }




    function estimateFee(bytes32 resourceId,uint256 inputAmount) public override view returns(uint256 fee){

        (uint256 relayerFee,uint256 gasFee) = _calFee(resourceId,inputAmount,MAX_GAS_PRICE);
        return relayerFee.add(gasFee);

    }

    // returns sorted token addresses, used to handle return values from pairs sorted in this order
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        require(tokenA != tokenB, 'UniswapV2Library: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'UniswapV2Library: ZERO_ADDRESS');
    }

    // calculates the CREATE2 address for a pair without making any external calls
    function pairFor(address factory, address tokenA, address tokenB) internal pure returns (address pair) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pair = address(uint(keccak256(abi.encodePacked(
                hex'ff',
                factory,
                keccak256(abi.encodePacked(token0, token1)),
                hex'96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f' // init code hash
            ))));
    }

    // fetches and sorts the reserves for a pair
    function getReserves(address factory, address tokenA, address tokenB) internal view returns (uint reserveA, uint reserveB) {
        (address token0,) = sortTokens(tokenA, tokenB);
        (uint reserve0, uint reserve1,) = IUniswapV2Pair(pairFor(factory, tokenA, tokenB)).getReserves();
        (reserveA, reserveB) = tokenA == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
    }

    function quote(uint amountA, uint reserveA, uint reserveB) internal pure returns (uint amountB) {
        require(amountA > 0, 'UniswapV2Library: INSUFFICIENT_AMOUNT');
        require(reserveA > 0 && reserveB > 0, 'UniswapV2Library: INSUFFICIENT_LIQUIDITY');
        amountB = amountA.mul(reserveB) / reserveA;
    }
}