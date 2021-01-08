pragma solidity 0.6.4;

import "./interfaces/IBridgeAccess.sol";

contract BridgeAccess is IBridgeAccess {


    event RelayerThresholdChanged(uint indexed newThreshold);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);

    address public owner;

    mapping(address=>bool) relayers;

    uint256 public _relayerThreshold;

    uint256 public _totalRelayers;

    constructor(address[] memory initialRelayers, uint initialRelayerThreshold) public {

        owner = msg.sender;

        _relayerThreshold = initialRelayerThreshold;

        for (uint i; i < initialRelayers.length; i++) {

            relayers[initialRelayers[i]] = true;

            _totalRelayers++;
        }

    }

    function relayerThreshold() public override view returns(uint256){
        return _relayerThreshold;
    }

    function totalRelayers() public override view returns(uint256){
        return _totalRelayers;
    }

    function isRelayer(address relayer) public override view returns(bool){
        return relayers[relayer];
    }


    /**
	 * @dev Throws if called by any account other than the owner.
	 */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }


    /**
	 * @dev Allows the current owner to transfer control of the contract to a newOwner.
	 * @param newOwner The address to transfer ownership to.
	 */
    function transferOwnership(address newOwner) public onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }


    function addRelayer(address user) public onlyOwner{
        require(!relayers[user],"is a relayer");
        relayers[user] =true;
        _totalRelayers++;
        emit RelayerAdded(user);

    }

    function delRelayer(address user) public onlyOwner {
        require(relayers[user],'not a relayer');
        delete relayers[user];
        _totalRelayers--;
        emit RelayerRemoved(user);
    }

    function dropRalayer() public {
        require(relayers[msg.sender],"not a rlayer");
        delete relayers[msg.sender];
        _totalRelayers--;
        emit RelayerRemoved(msg.sender);

    }

    function ChangeRelayerThreshold(uint newThreshold) external onlyOwner {
        _relayerThreshold = newThreshold;
        emit RelayerThresholdChanged(newThreshold);
    }

}