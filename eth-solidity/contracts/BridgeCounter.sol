pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "./interfaces/IBridgeCounter.sol";

contract BridgeCounter is IBridgeCounter{

    address public bridge;

    address public owner;

    // destinationChainID => number of deposits
    mapping(uint8 => uint64) public depositCounts;

    constructor(address bridge_) public {
        bridge = bridge_;
        owner = msg.sender;
    }


    modifier onlyOwner() {
        require(msg.sender == owner,
            "sender is not  admin");
        _;
    }

    modifier onlyBridge() {
        require(msg.sender == bridge,
            "sender is not bridge");
        _;
    }

    function incr(uint8 destinationChainID) public override onlyBridge returns(uint64) {
        return ++depositCounts[destinationChainID];
    }

    function setBridge(address bridge_) public onlyOwner {
        bridge = bridge_;
    }

    function setDestStartNonce(uint8 destinationChainID,uint64 start) public onlyOwner {
        depositCounts[destinationChainID] = start;
    }
}
