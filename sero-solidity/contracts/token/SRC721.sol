// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <0.8.0;

import "../utils/seroInterface.sol";
import "../utils/Strings.sol";
import "./ISRC721.sol";
import "./AccessControl.sol";


contract SRC721  is ISRC721Mintable,SeroInterface,AccessControl{

    using Strings for uint256;

    event Mint(address indexed to, bytes32 indexed ticket);

    string private _name;

    string private _symbol;

    uint256 private _totalSupply;

    uint256 public counter;


    // Optional mapping for ticket URIs
    mapping (bytes32 => string) private _ticketURIs;


    mapping(bytes32 => uint256) private _ticketToId;

    mapping(uint256 => bytes32) private _idToTicket;

    // Base URI
    string private _baseURI;

    bool public minting = true;


    constructor(string memory name_,string memory symbol_,string memory baseURI_) public{

        _name = name_;

        _symbol = symbol_;

        _baseURI = baseURI_;

    }

    /**
     * @return the name of the token.
     */
    function name() public view override returns (string memory) {
        return _name;
    }

    /**
     * @return the symbol of the token.
     */
    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }

    function baseURI() public view virtual returns (string memory) {
        return _baseURI;
    }

    function setBaseURI(string calldata baseURI_) external onlyOwner {
        _baseURI = baseURI_;
    }

    function stopMint() external onlyOwner {
        require(minting,"hash stoped");
        minting = false;
    }


    function ticketURI(bytes32 tkt) public view override returns (string memory) {

        string memory _ticketURI = _ticketURIs[tkt];

        // If there is no base URI, return the token URI.
        if (bytes(_baseURI).length == 0) {
            return _ticketURI;
        }
        // If both are set, concatenate the baseURI and tokenURI (via abi.encodePacked).
        if (bytes(_ticketURI).length > 0) {
            return string(abi.encodePacked(_baseURI, _ticketURI));
        }
        // If there is a baseURI but no tokenURI, concatenate the tokenID to the baseURI.
        return string(abi.encodePacked(_baseURI, _ticketToId[tkt].toString()));
    }


    function ticketId(bytes32 tkt) public view override returns(uint256) {
        return _ticketToId[tkt];
    }

    function ticket(uint256 tktId) external view override returns(bytes32){
        return _idToTicket[tktId];
    }


    function mint(address to,uint256 tktId) external override onlyMinter returns(bytes32){

        require(minting,"has stopped");

        require(to != address(0), "SRC721: mint to the zero address");

        bytes32 tkt = _idToTicket[tktId];

        if (tkt == bytes32(0)){

            tkt = _mint(to);

            _idToTicket[tktId]= tkt;

            _ticketToId[tkt] = tktId;

        }else {

            _totalSupply++;

            require(sero_send_ticket(to,_symbol,tkt),"send ticket failed");
        }

        return tkt;

    }

    function setTicketURI(bytes32 tokenId, string calldata _tokenURI) external override onlyMinter {
        _setTicketURI(tokenId,_tokenURI);
    }

    function _mint(address to) internal returns(bytes32){

        bytes32 tkt = sero_allotTicket(to,0,_symbol);

        require(tkt != bytes32(0),"allot failed");

        _totalSupply++;

        counter++;

        emit Mint(to, tkt);

        return tkt;
    }

    function burn() external override {

        string memory catg = sero_msg_category();

        _totalSupply--;

        require(equals(catg,_symbol),"not supported category");
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


    function _setTicketURI(bytes32 tokenId, string memory _ticketURI) internal virtual {
        _ticketURIs[tokenId] = _ticketURI;
    }


}