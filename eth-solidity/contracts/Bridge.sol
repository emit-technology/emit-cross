pragma solidity 0.6.4;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "./utils/Pausable.sol";
import "./interfaces/IDepositExecute.sol";
import "./interfaces/IBridge.sol";
import "./interfaces/IERCHandler.sol";
import "./interfaces/ICrossFee.sol";
import "./interfaces/IBridgeCounter.sol";


/**
    @title Facilitates deposits, creation and votiing of deposit proposals, and deposit executions.
    @author ChainSafe Systems.
 */
contract Bridge is Pausable, AccessControl {

    using SafeMath for uint256;


    address public crossFeeAddreess;

    address public bridgeCounterAddress;

    uint8   public _chainID;
    uint256 public _relayerThreshold;
    uint256 public _totalRelayers;
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

    // resourceID => handler address
    mapping(bytes32 => address) public _resourceIDToHandlerAddress;

    // depositNonce => destinationChainID => bytes
    mapping(uint64 => mapping(uint8 => bytes)) public _depositRecords;

    // originChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => Proposal)) public _proposals;

    // originChainID + depositNonce => dataHash => relayerAddress => bool
    mapping(uint72 => mapping(bytes32 => mapping(address => bool))) public _hasVotedOnProposal;


    struct Limit {

        uint256 min;

        uint256 max;
    }

    mapping (bytes32 => Limit) public resourceIDToLimit;

    event RelayerThresholdChanged(uint indexed newThreshold);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);
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

    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    modifier onlyAdmin() {
        _onlyAdmin();
        _;
    }


    modifier onlyRelayers() {
        _onlyRelayers();
        _;
    }


    function _onlyAdmin() view private {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "sender doesn't have admin role");
    }

    function _onlyRelayers() view private {
        require(hasRole(RELAYER_ROLE, msg.sender), "sender doesn't have relayer role");
    }

    /**
        @notice Initializes Bridge, creates and grants {msg.sender} the admin role,
        creates and grants {initialRelayers} the relayer role.
        @param chainID ID of chain the Bridge contract exists on.
        @param initialRelayers Addresses that should be initially granted the relayer role.
        @param initialRelayerThreshold Number of votes needed for a deposit proposal to be considered passed.
     */
    constructor (uint8 chainID, address[] memory initialRelayers, uint initialRelayerThreshold, uint256 expiry) public {
        _chainID = chainID;
        _relayerThreshold = initialRelayerThreshold;
        _expiry = expiry;

        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setRoleAdmin(RELAYER_ROLE, DEFAULT_ADMIN_ROLE);

        for (uint i; i < initialRelayers.length; i++) {
            grantRole(RELAYER_ROLE, initialRelayers[i]);
            _totalRelayers++;
        }

    }

    /**
        @notice Returns true if {relayer} has the relayer role.
        @param relayer Address to check.
     */
    function isRelayer(address relayer) external view returns (bool) {
        return hasRole(RELAYER_ROLE, relayer);
    }

    /**
        @notice Removes admin role from {msg.sender} and grants it to {newAdmin}.
        @notice Only callable by an address that currently has the admin role.
        @param newAdmin Address that admin role will be granted to.
     */
    function renounceAdmin(address newAdmin) external onlyAdmin {
        grantRole(DEFAULT_ADMIN_ROLE, newAdmin);
        renounceRole(DEFAULT_ADMIN_ROLE, msg.sender);
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

    function adminSetCrossFeeContract(address crossFee_) external onlyAdmin {
        crossFeeAddreess = crossFee_;
    }

    function adminSetBridgeCounterContract(address bridgeCounter_) external onlyAdmin {
        bridgeCounterAddress = bridgeCounter_;
    }

    /**
        @notice Modifies the number of votes required for a proposal to be considered passed.
        @notice Only callable by an address that currently has the admin role.
        @param newThreshold Value {_relayerThreshold} will be changed to.
        @notice Emits {RelayerThresholdChanged} event.
     */
    function adminChangeRelayerThreshold(uint newThreshold) external onlyAdmin {
        _relayerThreshold = newThreshold;
        emit RelayerThresholdChanged(newThreshold);
    }

    /**
        @notice Grants {relayerAddress} the relayer role and increases {_totalRelayer} count.
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be added.
        @notice Emits {RelayerAdded} event.
     */
    function adminAddRelayer(address relayerAddress) external onlyAdmin {
        require(!hasRole(RELAYER_ROLE, relayerAddress), "addr already has relayer role!");
        grantRole(RELAYER_ROLE, relayerAddress);
        emit RelayerAdded(relayerAddress);
        _totalRelayers++;
    }

    /**
        @notice Removes relayer role for {relayerAddress} and decreases {_totalRelayer} count.
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be removed.
        @notice Emits {RelayerRemoved} event.
     */
    function adminRemoveRelayer(address relayerAddress) external onlyAdmin {
        require(hasRole(RELAYER_ROLE, relayerAddress), "addr doesn't have relayer role!");
        revokeRole(RELAYER_ROLE, relayerAddress);
        emit RelayerRemoved(relayerAddress);
        _totalRelayers--;
    }

    /**
        @notice Sets a new resource for handler contracts that use the IERCHandler interface,
        and maps the {handlerAddress} to {resourceID} in {_resourceIDToHandlerAddress}.
        @notice Only callable by an address that currently has the admin role.
        @param handlerAddress Address of handler resource will be set for.
        @param resourceID ResourceID to be used when making deposits.
        @param tokenAddress Address of contract to be called when a deposit is made and a deposited is executed.
        @param tokenAddress Minimum amount of deposits.
     */
    function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress,uint256 minCrossAmount_,uint256 maxCrossAmount_) external onlyAdmin {

        _resourceIDToHandlerAddress[resourceID] = handlerAddress;

        IERCHandler handler = IERCHandler(handlerAddress);

        handler.setResource(resourceID, tokenAddress);

        resourceIDToLimit[resourceID].min = minCrossAmount_;

        resourceIDToLimit[resourceID].max = maxCrossAmount_;

    }


    function minCrossAmount(bytes32 resourceId) external view returns(uint256){
        return resourceIDToLimit[resourceId].min;
    }


    /**
        @notice Sets a resource as burnable for handler contracts that use the IERCHandler interface.
        @notice Only callable by an address that currently has the admin role.
        @param resourceID .
        @param tokenAddress Address of contract to be called when a deposit is made and a deposited is executed.
     */
    function adminSetBurnable(bytes32 resourceID, address tokenAddress) external onlyAdmin {
        address handlerAddress = _resourceIDToHandlerAddress[resourceID];

        IERCHandler handler = IERCHandler(handlerAddress);
        handler.setBurnable(tokenAddress);
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
        @param handlerAddress Address of handler to withdraw from.
        @param tokenAddress Address of token to withdraw.
        @param recipient Address to withdraw tokens to.
        @param amount the amount of ERC20 tokens.
     */
    function adminWithdraw(
        address handlerAddress,
        address tokenAddress,
        address payable recipient,
        uint256 amount
    ) external onlyAdmin {
        IERCHandler handler = IERCHandler(handlerAddress);
        handler.withdraw(tokenAddress, recipient, amount);
    }

    /**
        @notice Initiates a transfer using a specified handler contract.
        @notice Only callable when Bridge is not paused.
        @param destinationChainID ID of chain deposit will be bridged to.
        @param resourceID ResourceID used to find address of handler to be used for deposit.
        @param recipient Receipt address on the dedesstination chain.
        @param amount the amount of ERC20 tokens.
        @notice Emits {Deposit} event.
     */

    function depositFT(uint8 destinationChainID, bytes32 resourceID, bytes calldata recipient,uint256 amount) external whenNotPaused {

        require(destinationChainID != _chainID,"destinationChainID is self");

        require(amount >= resourceIDToLimit[resourceID].min &&  amount <= resourceIDToLimit[resourceID].max,"cross amount is to low");

        require(bridgeCounterAddress != address(0),"not set bridgeCounter address");

        address handler = _resourceIDToHandlerAddress[resourceID];

        require(handler != address(0), "resourceID not mapped to handler");

        uint64 depositNonce = IBridgeCounter(bridgeCounterAddress).incr(destinationChainID);

        _depositRecords[depositNonce][destinationChainID] = abi.encodePacked(recipient,amount);

        IDepositFTExecute depositHandler = IDepositFTExecute(handler);

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

        address handler = _resourceIDToHandlerAddress[resourceID];

        require(handler != address(0), "no handler for resourceID");

        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);

        bytes32 dataHash = keccak256(abi.encodePacked(handler, recipient,amount));

        Proposal storage proposal = _proposals[nonceAndID][dataHash];

        require(uint(proposal._status) <= 1, "proposal already passed/executed/cancelled");


        for(uint256 i = 0; i < signs.length; i++ ){

            address relayer = ecrecoverDecode(signHash,signs[i]);

            require(hasRole(RELAYER_ROLE, relayer),"not relayer signs");

            if (!_hasVotedOnProposal[nonceAndID][dataHash][relayer]){

                proposal._yesVotes.push(relayer);

                _hasVotedOnProposal[nonceAndID][dataHash][relayer] =true;
            }
        }


        if (uint(proposal._status) == 0) {

            ++_totalProposals;

            proposal._resourceID = resourceID;

            proposal._dataHash = dataHash;

            proposal._proposedBlock = block.number;

            if (_relayerThreshold <= 1 || proposal._yesVotes.length >= _relayerThreshold) {

                proposal._status = ProposalStatus.Executed;

            }else {
                proposal._status = ProposalStatus.Active;

            }

        } else {
            if (block.number.sub(proposal._proposedBlock) > _expiry) {
                // if the number of blocks that has passed since this proposal was
                // submitted exceeds the expiry threshold set, cancel the proposal
                proposal._status = ProposalStatus.Cancelled;


            } else {

                require(dataHash == proposal._dataHash, "datahash mismatch");

                if (_relayerThreshold <= 1 || proposal._yesVotes.length >= _relayerThreshold) {

                    proposal._status = ProposalStatus.Executed;

                }
            }

        }

        if (proposal._status == ProposalStatus.Executed){

            _excueteProposal(msg.sender,handler,proposal._resourceID,recipient,amount,proposal._yesVotes);

        }

        emit ProposalEvent(originChainID, depositNonce,proposal._status, resourceID, dataHash);

    }

    function _excueteProposal(address payable caller,address handler,bytes32 resourceID,address  recipient,uint256 amount,address[] memory relaysers) internal {

        IDepositFTExecute depositHandler = IDepositFTExecute(handler);

        (uint256 relayerFee,uint256 gasFee) = calGasFee(resourceID,amount,tx.gasprice);

        uint256 recipientAmount = amount.sub(relayerFee);

        if (recipientAmount > gasFee){

            recipientAmount = recipientAmount.sub(gasFee);

        }else {
            gasFee = recipientAmount;

            recipientAmount = 0;
        }

        if(recipientAmount >0){
            depositHandler.executeProposal(resourceID, recipient,recipientAmount);
        }

        depositHandler.transferFee(resourceID, caller, gasFee,relaysers,relayerFee);


    }

    function calGasFee(bytes32 resourceID,uint256 inputAmount,uint256 gasPrice) public view returns(uint256 relayerFee,uint256 gasFee) {

        if (crossFeeAddreess !=address(0)){
            (relayerFee, gasFee) = ICrossFee(crossFeeAddreess).calCrossFee(resourceID,inputAmount,gasPrice);
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

    /**
        @notice Transfers eth in the contract to the specified addresses. The parameters addrs and amounts are mapped 1-1.
        This means that the address at index 0 for addrs will receive the amount (in WEI) from amounts at index 0.
        @param addrs Array of addresses to transfer {amounts} to.
        @param amounts Array of amonuts to transfer to {addrs}.
     */
    function transferFunds(address payable[] calldata addrs, uint[] calldata amounts) external onlyAdmin {
        for (uint i = 0; i < addrs.length; i++) {
            addrs[i].transfer(amounts[i]);
        }
    }

}
