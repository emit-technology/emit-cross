pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/math/SafeMath.sol";
import "./utils/Pausable.sol";
import  "./utils/seroInterface.sol";

import "./interfaces/IDepositExecute.sol";
import "./interfaces/ISRCHandler.sol";
import "./interfaces/IBridgeCounter.sol";
import "./interfaces/ICrossFee.sol";
import "./interfaces/IBridgeAccess.sol";

/**
    @title Facilitates deposits, creation and votiing of deposit proposals, and deposit executions.
    @author ChainSafe Systems.
 */
contract Bridge is Pausable,SeroInterface {

    using SafeMath for uint256;

    address public owner;

    address public bridgeCounterAddress;

    address public crossFeeAddress;

    address public bridgeAccessAddress;

    uint8   public _chainID;

    uint256 public _totalProposals;

    uint256 public _expiry;

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

    address public src20HandlerAddress;

    // depositNonce => destinationChainID => bytes
    mapping(uint64 => mapping(uint8 => bytes)) public _depositRecords;
    // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => Proposal)) public _proposals;

    // destinationChainID + depositNonce => dataHash => relayerAddress => bool
    mapping(uint72 => mapping(bytes32 => mapping(address => bool))) public _hasVotedOnProposal;

    mapping (bytes32 => string) public _resourceIDToCurrency;

    struct Limit {

        uint256 min;

        uint256 max;
    }
    mapping (bytes32 => Limit) public resourceIDToLimit;

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


    function minCrossAmount(bytes32 resourceId) external view returns(uint256){
        return resourceIDToLimit[resourceId].min;
    }



    /**
        @notice Initializes Bridge, creates and grants {msg.sender} the admin role,
        creates and grants {initialRelayers} the relayer role.
        @param chainID ID of chain the Bridge contract exists on.
     */
    constructor (uint8 chainID, uint256 expiry,address brigeAccessAddress_) public {
        owner = msg.sender;
        _chainID = chainID;
        _expiry = expiry;
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
    function adminSetSrc20HandlerContract(address handler) external onlyAdmin {
        src20HandlerAddress = handler;
    }


    /**
        @notice Sets a new resource for handler contracts that use the ISRCHandler interface,
        @notice Only callable by an address that currently has the admin role.
        @param resourceID ResourceID to be used when making deposits.
        @param tokenAddress Address of contract to be called when a deposit is made and a deposited is executed.
     */
    function adminSetResourceTokenAddress(bytes32 resourceID, address tokenAddress) external onlyAdmin {
        ISRCHandler handler = ISRCHandler(src20HandlerAddress);
        handler.setReourceTokenAddress(resourceID, tokenAddress);
    }

    /**
        @notice Sets a new resource for handler contracts that use the ISRCHandler interface,
        @notice Only callable by an address that currently has the admin role.
        @param resourceID ResourceID to be used when making deposits.
        @param currency the currency to be deposited.
        @param minCrossAmount_ minimum deposited amount.
     */
    function adminSetResourceCurrency(bytes32 resourceID, string calldata currency,uint256 minCrossAmount_,uint256 maxCrossAmount_) external onlyAdmin {
        ISRCHandler handler = ISRCHandler(src20HandlerAddress);
        handler.setResourceCurrency(resourceID, currency);
        _resourceIDToCurrency[resourceID] = currency;
        resourceIDToLimit[resourceID].min = minCrossAmount_;
        resourceIDToLimit[resourceID].max = maxCrossAmount_;
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
        @param currency symbol of token to withdraw.
        @param recipient Address to withdraw tokens to.
        @param amount  the amount of SRC20 tokens to withdraw.
     */
    function adminWithdraw(
        string calldata currency,
        address recipient,
        uint256 amount
    ) external onlyAdmin {
        ISRCHandler handler = ISRCHandler(src20HandlerAddress);
        handler.withdraw(currency, recipient, amount);
    }

    /**
       @notice Initiates a transfer using a specified handler contract.
       @notice Only callable when Bridge is not paused.
       @param destinationChainID ID of chain deposit will be bridged to.
       @param resourceID ResourceID used to find address of handler to be used for deposit.
       @param recipient Receipt address on the dedesstination chain.
       @notice Emits {Deposit} event.
    */

    function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient) external payable whenNotPaused {

        require(destinationChainID != _chainID,"destinationChainID is self");

        require(bridgeCounterAddress != address(0),"not set bridgeCounter address");

        string memory cy = sero_msg_currency();

        require(equals(cy,_resourceIDToCurrency[resourceID]),"invalid resource currency");

        require(msg.value >=resourceIDToLimit[resourceID].min && msg.value <= resourceIDToLimit[resourceID].max,"invalid cross amount" );

        require(src20HandlerAddress != address(0), "resourceID not mapped to handler");

        uint64 depositNonce = IBridgeCounter(bridgeCounterAddress).incr(destinationChainID);

        _depositRecords[depositNonce][destinationChainID] = abi.encodePacked(recipient,msg.value);

        IDepositExecute depositHandler = IDepositExecute(src20HandlerAddress);

        sero_setCallValues(_resourceIDToCurrency[resourceID],msg.value,"",0);

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

        require(src20HandlerAddress != address(0), "no handler for resourceID");

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
            if (block.number.sub( proposal._proposedBlock) > _expiry) {
                // if the number of blocks that has passed since this proposal was
                // submitted exceeds the expiry threshold set, cancel the proposal
                proposal._status = ProposalStatus.Cancelled;
                emit ProposalEvent(chainID, depositNonce, ProposalStatus.Cancelled, resourceID, dataHash);
            } else {
                require(dataHash == proposal._dataHash, "datahash mismatch");
                proposal._yesVotes.push(msg.sender);


            }

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
        @param amount Amount originally provided when deposit was made
        @notice Proposal must have Passed status.
        @notice Hash of {data} must equal proposal's {dataHash}.
        @notice Emits {ProposalEvent} event with status {Executed}.
     */
    function executeProposal(uint8 originChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount) external onlyRelayers whenNotPaused {

        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);

        bytes32 dataHash = keccak256(abi.encodePacked(src20HandlerAddress, recipient,amount));

        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(proposal._status != ProposalStatus.Inactive, "proposal is not active");

        require(proposal._status == ProposalStatus.Passed, "proposal already transferred");

        require(dataHash == proposal._dataHash, "data doesn't match datahash");

        proposal._status = ProposalStatus.Executed;

        _excueteProposal(msg.sender,resourceID,recipient,amount,proposal._yesVotes);

        emit ProposalEvent(originChainID, depositNonce, proposal._status, proposal._resourceID, proposal._dataHash);

    }
    function _excueteProposal(address  caller,bytes32 resourceID,address  recipient,uint256 amount,address[] memory relaysers) internal {

        IDepositExecute depositHandler = IDepositExecute(src20HandlerAddress);

        (uint256 relayeFee,uint256 gasFee) = calGasFee(resourceID,amount,tx.gasprice);

        uint256 recipientAmount = amount.sub(relayeFee);

        if (recipientAmount > gasFee){

            recipientAmount = recipientAmount.sub(gasFee);

        }else {
            gasFee = recipientAmount;

            recipientAmount = 0;
        }

        if(recipientAmount >0){
            depositHandler.executeProposal(resourceID, recipient,recipientAmount);
        }

        if(gasFee.add(relayeFee) >0){
            depositHandler.transferFee(resourceID, caller, gasFee,relaysers,relayeFee);
        }

    }

    function calGasFee(bytes32 resourceID,uint256 inputAmount,uint256 gasPrice) public view returns(uint256 relayerFee,uint256 gasFee) {

        if (crossFeeAddress !=address(0)){
            (relayerFee, gasFee) = ICrossFee(crossFeeAddress).calCrossFee(resourceID,inputAmount,gasPrice);
        }

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
        require(block.number.sub(proposal._proposedBlock) > _expiry, "Proposal not at expiry threshold");

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
