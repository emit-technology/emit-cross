pragma solidity 0.5.14;


interface IBridgeCounter {

    function incr(uint8 destinationChainID) external returns(uint64);
}

