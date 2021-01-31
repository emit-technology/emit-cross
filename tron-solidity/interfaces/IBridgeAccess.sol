
pragma solidity 0.5.14;


interface IBridgeAccess {

    function isRelayer(address relayer) external view returns (bool);

    function relayerThreshold() external view returns(uint256);

    function totalRelayers() external view returns(uint256);
}

