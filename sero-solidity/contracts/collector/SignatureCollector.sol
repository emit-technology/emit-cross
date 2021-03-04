pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "../interfaces/IBridgeAccess.sol";

contract SignatureCollector  {

    IBridgeAccess public bridgeAccess ;

    address public owner;

    enum ProposalStatus {Inactive, Active, Passed}


    enum TransferType {_,Fungible,NonFungible}

    struct Proposal {
        bytes[] signatures;
        ProposalStatus status;
    }

    event SignProposalEvent(uint8 indexed resourceChainId,uint8 indexed destinationChainId ,uint64 indexed depositNonce,TransferType typ, ProposalStatus  status);


    mapping(uint72 => mapping(uint64 => mapping(address => bool))) private _hasSignedProposal;

    mapping(uint72 =>uint64) public _destinationStartNonce;

    mapping(uint72 => mapping(uint64 => Proposal)) public destinationProposals;



    function coId(uint8 originChainId, uint8 destinationId) public pure returns(uint72) {
        return ((uint72(originChainId) << 8) | uint72(destinationId));
    }

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

        owner = msg.sender;

        bridgeAccess = IBridgeAccess(bridgeAccessAddr);

    }

    function hasSignedProposal(uint8 srcId,uint8 destId,uint64 depositNonce,address sender )public view returns(bool){
        return _hasSignedProposal[coId(srcId,destId)][depositNonce][sender];
    }

    function getProposalSignatures(uint8 srcId,uint8 destId, uint64 depositNonce) public view returns(ProposalStatus status,bytes[] memory signatures){
        status = destinationProposals[coId(srcId,destId)][depositNonce].status;
        signatures = destinationProposals[coId(srcId,destId)][depositNonce].signatures;
    }



    function changeBridgeAccess (address bridgeAccessAddr) public onlyOwner {
        bridgeAccess = IBridgeAccess(bridgeAccessAddr);
    }

    function setDestinationStartNonce(uint8 srcId,uint8 destId,uint64 startNonce) public onlyOwner {
        _destinationStartNonce[coId(srcId,destId)] = startNonce;
    }
    function getDestinationStartNoncee(uint8 srcId,uint8 destId)public view returns(uint64){
        return _destinationStartNonce[coId(srcId,destId)];
    }

    function signProposal(uint8 srcId,uint8 destId, uint64 depositNonce, TransferType typ,bytes calldata signature) external onlyRelayer  {

        require(uint8(typ) > 0,"invalid transferType");

        require(depositNonce >= _destinationStartNonce[coId(srcId,destId)], "old depositNonce");

        Proposal storage proposal = destinationProposals[coId(srcId,destId)][depositNonce];

        require(uint(proposal.status) <= 1, "proposal already passed/executed/cancelled");

        require(!_hasSignedProposal[coId(srcId,destId)][depositNonce][msg.sender], "relayer already commit sign");

        _hasSignedProposal[coId(srcId,destId)][depositNonce][msg.sender] = true;

        if (uint(proposal.status) == 0) {
            proposal.status = ProposalStatus.Active;
            emit SignProposalEvent(srcId,destId,depositNonce, typ,ProposalStatus.Active);
        }
        proposal.signatures.push(signature);

        uint256 _relayerThreshold = bridgeAccess.relayerThreshold();

        if (proposal.signatures.length >= _relayerThreshold) {

            proposal.status = ProposalStatus.Passed;

            emit SignProposalEvent(srcId,destId, depositNonce,typ, ProposalStatus.Passed);
        }

    }



}