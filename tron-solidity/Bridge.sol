pragma solidity 0.5.14;

pragma experimental ABIEncoderV2;


import "./utils/Pausable.sol";
import "./utils/SafeMath.sol";
import "./interfaces/IDepositExecute.sol";
import "./interfaces/ITRCHandler.sol";
import "./interfaces/IBridgeAccess.sol";
import "./interfaces/IBridgeCounter.sol";
import "./interfaces/ICrossFee.sol";

/**
    @title Facilitates deposits, creation and voting of deposit proposals, and deposit executions.
    @author ChainSafe Systems.
 */
contract Bridge is Pausable {

    using SafeMath for *;

    address  public bridgeAccessAddress;

    address public bridgeCounterAddress;

    address public crossFeeAddress;

    address public owner;

    uint8   public _chainID;

    enum ProposalStatus {Inactive, Active, Passed, Executed, Cancelled}

    struct Proposal {
        ProposalStatus _status;
        address[] _yesVotes;
        uint64  _proposedBlock;
    }


    address public trc20HandlerAddress;


    // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => Proposal)) private _proposals;

    mapping(uint72 => mapping(bytes32 => mapping(address => bool))) public _hasVotedOnProposal;

    struct Limit {

        uint256 min;

        uint256 max;
    }
    mapping (bytes32 => Limit) public resourceIDToLimit;


    event Deposit(
        uint8   destinationChainID,
        bytes32 resourceID,
        uint64  depositNonce
    );
    event ProposalEvent(
        uint8           indexed originChainID,
        uint64          indexed depositNonce,
        ProposalStatus  indexed status,
        bytes32 resourceID,
        bytes32 dataHash
    );

    modifier onlyAdmin() {
        _onlyAdmin();
        _;
    }


    modifier onlyRelayers() {
        _onlyRelayers();
        _;
    }



    function _onlyAdmin() private view {
        require(owner == msg.sender, "sender doesn't have admin role");
    }

    function _onlyRelayers() private view {

        require( IBridgeAccess(bridgeAccessAddress).isRelayer(msg.sender), "sender doesn't have relayer role");
    }



    /**
        @notice Initializes Bridge, creates and grants {msg.sender} the admin role,
        creates and grants {initialRelayers} the relayer role.
        @param chainID ID of chain the Bridge contract exists on.
     */
    constructor (uint8 chainID,address bridgeAccess_) public {
        owner =  msg.sender;
        _chainID = chainID;
        bridgeAccessAddress = bridgeAccess_;
    }


    /**
        @notice Returns true if {relayer} has the relayer role.
        @param relayer Address to check.
     */
    function isRelayer(address relayer) public view returns (bool) {
        return IBridgeAccess(bridgeAccessAddress).isRelayer(relayer);
    }

    /**
        @notice Removes admin role from {msg.sender} and grants it to {newAdmin}.
        @notice Only callable by an address that currently has the admin role.
        @param newAdmin Address that admin role will be granted to.
     */
    function renounceAdmin(address newAdmin) external onlyAdmin {
        require(msg.sender != newAdmin, 'Cannot renounce oneself');
        owner = newAdmin;
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
        @notice Sets a new resource for handler contracts that use the IERCHandler interface,
        and maps the {handlerAddress} to {resourceID} in {_resourceIDToHandlerAddress}.
        @notice Only callable by an address that currently has the admin role.
        @param handlerAddress Address of handler resource will be set for.
        @param resourceID ResourceID to be used when making deposits.
        @param tokenAddress Address of contract to be called when a deposit is made and a deposited is executed.
     */
    function adminSetResource(address handlerAddress,bytes32 resourceID, address tokenAddress,uint256 minCrossAmount_,uint256 maxCrossAmount_) external onlyAdmin {
        trc20HandlerAddress = handlerAddress;
        IERCHandler handler = IERCHandler(trc20HandlerAddress);
        handler.setResource(resourceID, tokenAddress);
        resourceIDToLimit[resourceID] = Limit(minCrossAmount_,maxCrossAmount_);
    }


    function adminSetCrossFeeContract(address crossFee_) external onlyAdmin {
        crossFeeAddress = crossFee_;
    }

    function adminSetBridgeCounterContract(address bridgeCounter_) external onlyAdmin {
        bridgeCounterAddress = bridgeCounter_;
    }

    function adminSetBridgeAccessContract(address bridgeAccess_) external onlyAdmin {
        bridgeAccessAddress = bridgeAccess_;
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
        @notice Returns total relayers number.
        @notice Added for backwards compatibility.
     */
    function _totalRelayers() public view returns (uint) {
        return IBridgeAccess(bridgeAccessAddress).totalRelayers();
    }

    function _relayerThreshold() public view returns(uint){
        return IBridgeAccess(bridgeAccessAddress).relayerThreshold();
    }



    /**
        @notice Used to manually withdraw funds from ERC safes.
        @param recipient Address to withdraw tokens to.
        @param amountOrTokenID Either the amount of ERC20 tokens or the ERC721 token ID to withdraw.
     */
    function adminWithdraw(
        address tokenAddress,
        address recipient,
        uint256 amountOrTokenID
    ) external onlyAdmin {
        IERCHandler handler = IERCHandler(trc20HandlerAddress);
        handler.withdraw(tokenAddress, recipient, amountOrTokenID);
    }

    /**
        @notice Initiates a transfer using a specified handler contract.
        @notice Only callable when Bridge is not paused.
        @param destinationChainID ID of chain deposit will be bridged to.
        @param resourceID ResourceID used to find address of handler to be used for deposit.
       @notice Emits {Deposit} event.
     */
    function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient ,uint256 amount) external whenNotPaused {

        require(trc20HandlerAddress != address(0), "resourceID not mapped to handler");

        uint64 depositNonce = IBridgeCounter(bridgeCounterAddress).incr(destinationChainID);

        IDepositExecute depositHandler = IDepositExecute(trc20HandlerAddress);

        depositHandler.deposit(resourceID, destinationChainID, depositNonce, msg.sender, recipient,amount);

        emit Deposit(destinationChainID, resourceID, depositNonce);
    }




    function bytesToBytes32(bytes memory source) public pure returns (bytes32 result) {
        assembly {
            result := mload(add(source, 32))
        }
    }

    function slice(bytes memory data, uint start, uint len) public pure returns (bytes memory){
        bytes memory b = new bytes(len);

        for(uint i = 0; i < len; i++){
            b[i] = data[i + start];
        }

        return b;
    }

    function ecrecoverDecode(bytes32 signHash,bytes memory sign) public pure returns (address addr){
        bytes32  r = bytesToBytes32(slice(sign, 0, 32));
        bytes32  s = bytesToBytes32(slice(sign, 32, 32));
        byte  v1 = slice(sign, 64, 1)[0];
        uint8 v = uint8(v1) + 27;
        addr = ecrecover(signHash, v, r, s);
        return addr;
    }


    /**
       @notice Only callable by relayers when Bridge is not paused.
       @param originChainID ID of chain deposit originated from.
       @param depositNonce ID of deposited generated by origin Bridge contract.
       @param recipient receipt address .
       @param amount received amount .
       @param signs relayer's signature to the proposale.
       @notice Proposal must not have already been passed or executed.
       @notice {msg.sender} must not have already voted on proposal.
       @notice Emits {ProposalEvent} event with status indicating the proposal status.
    */
    function commitVotes(uint8 originChainID, uint64 depositNonce,bytes32 resourceID, address  recipient,uint256 amount,bytes[] calldata signs) external  whenNotPaused{

        require(signs.length>0,"invalid signs");

        bytes32 signHash =  keccak256(abi.encodePacked(address(this),originChainID,_chainID, depositNonce,resourceID,recipient,amount));

        require(trc20HandlerAddress != address(0), "no handler for resourceID");

        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);

        bytes32 dataHash = keccak256(abi.encodePacked(trc20HandlerAddress, recipient,amount));

        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(uint(proposal._status) <= 1, "proposal already passed/executed/cancelled");


        for(uint256 i = 0; i < signs.length; i++ ){

            address relayer = ecrecoverDecode(signHash,signs[i]);

            require(isRelayer(relayer),"not relayer signs");

            if (!_hasVotedOnProposal[nonceAndID][dataHash][relayer]){

                proposal._yesVotes.push(relayer);

                _hasVotedOnProposal[nonceAndID][dataHash][relayer] =true;
            }
        }


        if (uint(proposal._status) == 0) {

            proposal._proposedBlock = uint64(block.number);

            if (proposal._yesVotes.length >= _relayerThreshold()) {

                proposal._status = ProposalStatus.Executed;

            }else {
                proposal._status = ProposalStatus.Active;

            }

        } else {


            if (proposal._yesVotes.length >= _relayerThreshold()) {

                proposal._status = ProposalStatus.Executed;

            }


        }

        if (proposal._status == ProposalStatus.Executed){

            _excueteProposal(msg.sender,resourceID,recipient,amount,proposal._yesVotes);

        }

        emit ProposalEvent(originChainID, depositNonce,proposal._status, resourceID, dataHash);

    }

    /**
        @notice Cancels a deposit proposal that has not been executed yet.
        @notice Only callable by relayers when Bridge is not paused.
        @param chainID ID of chain deposit originated from.
        @param depositNonce ID of deposited generated by origin Bridge contract.
        @param dataHash Hash of data originally provided when deposit was made.
        @notice Proposal must be past expiry threshold.
        @notice Emits {ProposalEvent} event with status {Cancelled}.
     */
    function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID,bytes32 dataHash) public onlyAdmin {
        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(chainID);
        Proposal memory proposal = _proposals[nonceAndID][dataHash];
        ProposalStatus currentStatus = proposal._status;

        require(currentStatus == ProposalStatus.Active || currentStatus == ProposalStatus.Passed,
            "Proposal cannot be cancelled");

        proposal._status = ProposalStatus.Cancelled;

        _proposals[nonceAndID][dataHash] = proposal;

        emit ProposalEvent(chainID, depositNonce, ProposalStatus.Cancelled,resourceID, dataHash);
    }


    function _excueteProposal(address  caller,bytes32 resourceID,address  recipient,uint256 amount,address[] memory relaysers) internal {

        IDepositExecute depositHandler = IDepositExecute(trc20HandlerAddress);

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
}
