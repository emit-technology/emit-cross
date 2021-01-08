pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../interfaces/IBridgeAccess.sol";

contract SignatureCollector  {

    IBridgeAccess public bridgeAccess ;

    address public owner;



    event SignProposal(uint8 indexed resourceChainId,uint8 indexed destinationChainId ,uint64 indexed depositNonce,bytes signature);


    mapping(uint8 => mapping(uint64 => mapping(address => bool))) public hasSignedProposal;



    /**
   * @dev Throws if called by any account other than the owner.
   */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyRelayer() {
        require(bridgeAccess.isRelayer(msg.sender),"not relayer");
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

    constructor (address bridgeAccessAddr) public {

        bridgeAccess = IBridgeAccess(bridgeAccessAddr);

    }



    function changeBridgeAccess (address bridgeAccessAddr) public onlyOwner {
        bridgeAccess = IBridgeAccess(bridgeAccessAddr);
    }

    function commitSign(uint8 sourceChainId,uint8 destinationChainId, uint64 depositNonce, bytes memory signature ) public onlyRelayer {

        require(!hasSignedProposal[destinationChainId][depositNonce][msg.sender], "relayer already commit sign");


        hasSignedProposal[destinationChainId][depositNonce][msg.sender] = true;

        emit SignProposal(sourceChainId,destinationChainId,depositNonce,signature);
    }

}