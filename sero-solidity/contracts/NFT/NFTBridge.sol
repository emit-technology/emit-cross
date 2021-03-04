// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <0.8.0;pragma experimental ABIEncoderV2;

import "../utils/SafeMath.sol";
import "../utils/Pausable.sol";
import  "../utils/seroInterface.sol";

import "./ISRC721Handler.sol";
import "./IDepositExecute.sol";
import "../interfaces/IBridgeCounter.sol";
import "./ISRC721CrossFee.sol";
import "../interfaces/IBridgeAccess.sol";

/**
    @title Facilitates deposits, creation and votiing of deposit proposals, and deposit executions.
    @author ChainSafe Systems.
 */
contract NFTBridge is Pausable,SeroInterface {

    using SafeMath for uint256;

    address public owner;

    address public bridgeCounterAddress;

    address public crossFeeAddress;

    address public bridgeAccessAddress;

    uint8   public _chainID;

    uint256 public _totalProposals;

    enum Vote {No, Yes}

    enum ProposalStatus {Inactive, Active, Passed, Executed, Cancelled}

    struct Proposal {
        bytes32 _resourceID;
        bytes32 _dataHash;
        address[] _yesVotes;
        address[] _noVotes;
        ProposalStatus _status;
        uint256 _proposedBlock;
    }

    address public src721HandlerAddress;

    // depositNonce => destinationChainID => bytes
    mapping(uint64 => mapping(uint8 => bytes)) public _depositRecords;
    // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => Proposal)) public _proposals;

    // destinationChainID + depositNonce => dataHash => relayerAddress => bool
    mapping(uint72 => mapping(bytes32 => mapping(address => bool))) public _hasVotedOnProposal;

    mapping (bytes32 => string) public _resourceIDToCategory;



    event Deposit(
        uint8   indexed destinationChainID,
        bytes32 indexed resourceID,
        uint64  indexed depositNonce
    );
    event ProposalEvent(
        uint8           indexed originChainID,
        uint64          indexed depositNonce,
        ProposalStatus  indexed status,
        bytes32 resourceID,
        bytes32 dataHash
    );

    event ProposalVote(
        uint8   indexed originChainID,
        uint64  indexed depositNonce,
        ProposalStatus indexed status,
        bytes32 resourceID
    );

    modifier onlyAdmin() {
        _onlyAdmin();
        _;
    }



    modifier onlyRelayers() {
        _onlyRelayers();
        _;
    }


    function _onlyAdmin() view private {
        require( msg.sender == owner, "sender not admin ");
    }

    function _onlyRelayers() view private {
        require(IBridgeAccess(bridgeAccessAddress).isRelayer(msg.sender), "sender doesn't have relayer role");
    }


    /**
        @notice Initializes Bridge, creates and grants {msg.sender} the admin role,
        creates and grants {initialRelayers} the relayer role.
        @param chainID ID of chain the Bridge contract exists on.
     */
    constructor (uint8 chainID,address brigeAccessAddress_) public {
        owner = msg.sender;
        _chainID = chainID;
        bridgeAccessAddress = brigeAccessAddress_;

    }


    function transferOwnership(address newOwner) public onlyAdmin {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }

    /**
        @notice Pauses deposits, proposal creation and voting, and deposit executions.
        @notice Only callable by an address that currently has the admin role.
     */
    function adminPauseTransfers() external onlyAdmin {
        _pause();
    }

    /**
        @notice Unpauses deposits, proposal creation and voting, and deposit executions.
        @notice Only callable by an address that currently has the admin role.
     */
    function adminUnpauseTransfers() external onlyAdmin {
        _unpause();
    }

    /**
        @notice Only callable by an address that currently has the admin role.
        @param bridgeCounter_ bridgeCounter contract address.

     */
    function adminSetBridgeCounterContract(address bridgeCounter_) external onlyAdmin {
        bridgeCounterAddress = bridgeCounter_;
    }

    /**
       @notice Only callable by an address that currently has the admin role.
       @param crossFee_ CrossFee contract address.

    */
    function adminSetCrossFeeContract(address crossFee_) external onlyAdmin {
        crossFeeAddress = crossFee_;
    }

    /**
      @notice Only callable by an address that currently has the admin role.
      @param handler src20Handler contract address.

   */
    function adminSetSrc721HandlerContract(address handler) external onlyAdmin {
        src721HandlerAddress = handler;
    }


    /**
        @notice Sets a new resource for handler contracts that use the ISRCHandler interface,
        @notice Only callable by an address that currently has the admin role.
        @param resourceID ResourceID to be used when making deposits.
        @param ticketAddress Address of contract to be called when a deposit is made and a deposited is executed.
     */
    function adminSetResourceInfo(bytes32 resourceID, address ticketAddress,string calldata category,bool burnable) external onlyAdmin {
        ISRC721Handler handler = ISRC721Handler(src721HandlerAddress);
        handler.setResourceCategory(resourceID, category);
        handler.setResourceContractAddress(resourceID, ticketAddress);
        _resourceIDToCategory[resourceID] = category;
        if (burnable){
            handler.setBurnable(ticketAddress);
        }
    }

    function adminSetResourceFeeHandler(bytes32 resourceID,address scr20handler) external onlyAdmin {
        ISRC721Handler handler = ISRC721Handler(src721HandlerAddress);
        handler.setResourceFeeHandler(resourceID,scr20handler);

    }


    /**
        @notice Returns a proposal.
        @param originChainID Chain ID deposit originated from.
        @param depositNonce ID of proposal generated by proposal's origin Bridge contract.
        @param dataHash Hash of data to be provided when deposit proposal is executed.
        @return Proposal which consists of:
        - _dataHash Hash of data to be provided when deposit proposal is executed.
        - _yesVotes Number of votes in favor of proposal.
        - _noVotes Number of votes against proposal.
        - _status Current status of proposal.
     */
    function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) external view returns (Proposal memory) {
        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);
        return _proposals[nonceAndID][dataHash];
    }



    /**
        @notice Used to manually withdraw funds from ERC safes.
        @param catg symbol of token to withdraw.
        @param recipient Address to withdraw tokens to.
        @param tkt  the amount of SRC20 tokens to withdraw.
     */
    function adminWithdraw(
        string calldata catg,
        address recipient,
        bytes32 tkt
    ) external onlyAdmin {
        ISRC721Handler handler = ISRC721Handler(src721HandlerAddress);
        handler.withdraw(catg, recipient, tkt);
    }

    /**
       @notice Initiates a transfer using a specified handler contract.
       @notice Only callable when Bridge is not paused.
       @param destinationChainID ID of chain deposit will be bridged to.
       @param resourceID ResourceID used to find address of handler to be used for deposit.
       @param recipient Receipt address on the dedesstination chain.
       @notice Emits {Deposit} event.
    */

    function depositNFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient) external payable whenNotPaused {

        require(destinationChainID != _chainID,"destinationChainID is self");

        require(bridgeCounterAddress != address(0),"not set bridgeCounter address");

        string memory catg = sero_msg_category();

        require(equals(catg,_resourceIDToCategory[resourceID]),"invalid resource category");


        if (crossFeeAddress !=address(0)){
            (string memory feecy,uint256 minValue) = ISRC721CrossFee(crossFeeAddress).getFeeInfo(resourceID);
            if (minValue > 0){
                require(msg.value >= minValue,"send value to low");
                string memory cy = sero_msg_currency();
                require(equals(cy,feecy),"invalid resource fee currency");
            }
        }else {
            require(msg.value == 0,"not need send value");
        }


        require(src721HandlerAddress != address(0), "resourceID not mapped to handler");

        uint64 depositNonce = IBridgeCounter(bridgeCounterAddress).incr(destinationChainID);

        _depositRecords[depositNonce][destinationChainID] = abi.encodePacked(recipient,msg.value);

        IDepositExecute depositHandler = IDepositExecute(src721HandlerAddress);

        sero_setCallValues(sero_msg_currency(),msg.value,_resourceIDToCategory[resourceID],sero_msg_ticket());

        depositHandler.deposit(resourceID, destinationChainID, depositNonce, msg.sender,recipient);

        emit Deposit(destinationChainID, resourceID, depositNonce);
    }

    /**
          @notice When called, {msg.sender} will be marked as voting in favor of proposal.
          @notice Only callable by relayers when Bridge is not paused.
          @param chainID ID of chain deposit originated from.
          @param depositNonce ID of deposited generated by origin Bridge contract.
          @param dataHash Hash of data provided when deposit was made.
          @notice Proposal must not have already been passed or executed.
          @notice {msg.sender} must not have already voted on proposal.
          @notice Emits {ProposalEvent} event with status indicating the proposal status.
          @notice Emits {ProposalVote} event.
       */
    function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) external onlyRelayers whenNotPaused {

        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(chainID);

        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(src721HandlerAddress != address(0), "no handler for resourceID");

        require(uint(proposal._status) <= 1, "proposal already passed/executed/cancelled");

        require(!_hasVotedOnProposal[nonceAndID][dataHash][msg.sender], "relayer already voted");

        uint256 _relayerThreshold = IBridgeAccess(bridgeAccessAddress).relayerThreshold();

        if (uint(proposal._status) == 0) {
            ++_totalProposals;
            _proposals[nonceAndID][dataHash] = Proposal({
                _resourceID : resourceID,
                _dataHash : dataHash,
                _yesVotes : new address[](1),
                _noVotes : new address[](0),
                _status : ProposalStatus.Active,
                _proposedBlock : block.number
                });

            proposal._yesVotes[0] = msg.sender;
            emit ProposalEvent(chainID, depositNonce, ProposalStatus.Active, resourceID, dataHash);
        } else {

            require(dataHash == proposal._dataHash, "datahash mismatch");
            proposal._yesVotes.push(msg.sender);

        }
        if (proposal._status != ProposalStatus.Cancelled) {
            _hasVotedOnProposal[nonceAndID][dataHash][msg.sender] = true;
            emit ProposalVote(chainID, depositNonce, proposal._status, resourceID);

            // If _depositThreshold is set to 1, then auto finalize
            // or if _relayerThreshold has been exceeded
            if (_relayerThreshold  <= 1 || proposal._yesVotes.length >= _relayerThreshold) {
                proposal._status = ProposalStatus.Passed;

                emit ProposalEvent(chainID, depositNonce, ProposalStatus.Passed, resourceID, dataHash);
            }
        }

    }



    /**
        @notice Executes a deposit proposal that is considered passed using a specified handler contract.
        @notice Only callable by relayers when Bridge is not paused.
        @param originChainID ID of chain deposit originated from.
        @param resourceID ResourceID to be used when making deposits.
        @param depositNonce ID of deposited generated by origin Bridge contract.
        @param recipient Recipient originally provided when deposit was made.
        @param tokenId tokenId originally provided when deposit was made
        @notice Proposal must have Passed status.
        @notice Hash of {data} must equal proposal's {dataHash}.
        @notice Emits {ProposalEvent} event with status {Executed}.
     */
    function executeProposal(uint8 originChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 tokenId,bytes calldata metaData) external onlyRelayers whenNotPaused {

        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);

        bytes32 dataHash = keccak256(abi.encodePacked(src721HandlerAddress, recipient,tokenId,metaData));

        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(proposal._status != ProposalStatus.Inactive, "proposal is not active");

        require(proposal._status == ProposalStatus.Passed, "proposal already transferred");

        require(dataHash == proposal._dataHash, "data doesn't match datahash");

        proposal._status = ProposalStatus.Executed;

        _excueteProposal(resourceID,recipient,tokenId,metaData);

        emit ProposalEvent(originChainID, depositNonce, proposal._status, proposal._resourceID, proposal._dataHash);

    }
    function _excueteProposal(bytes32 resourceID,address  recipient,uint256 tokenId,bytes memory metaData) internal {

        IDepositExecute depositHandler = IDepositExecute(src721HandlerAddress);

        depositHandler.executeProposal(resourceID, recipient,tokenId,metaData);

    }




    /**
        @notice Executes a deposit proposal that is considered passed using a specified handler contract.
        @notice Only callable by relayers when Bridge is not paused.
        @param chainID ID of chain deposit originated from.
        @param depositNonce ID of deposited generated by origin Bridge contract.
        @param dataHash Hash of data originally provided when deposit was made.
        @notice Proposal must be past expiry threshold.
        @notice Emits {ProposalEvent} event with status {Cancelled}.
     */
    function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) public onlyAdmin {
        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(chainID);
        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(proposal._status != ProposalStatus.Cancelled, "Proposal already cancelled");

        proposal._status = ProposalStatus.Cancelled;
        emit ProposalEvent(chainID, depositNonce, ProposalStatus.Cancelled, proposal._resourceID, proposal._dataHash);

    }


    function equals(string memory a, string memory b) internal pure returns (bool) {
        if (bytes(a).length != bytes(b).length) {
            return false;
        }
        for (uint i = 0; i < bytes(a).length; i ++) {
            if(bytes(a)[i] != bytes(b)[i]) {
                return false;
            }
        }
        return true;
    }

}
