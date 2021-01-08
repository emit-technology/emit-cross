pragma solidity 0.6.4;


interface IBridgeCounter {

    function incr(uint8 destinationChainID) external returns(uint64);
}

