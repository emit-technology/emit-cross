// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <0.8.0;


interface ISRC721 {
    /**
   * @dev Returns the token collection name.
   */
    function name() external view returns (string memory);

    /**
     * @dev Returns the token collection symbol.
     */
    function symbol() external view returns (string memory);

    /**
     * @dev Returns the Uniform Resource Identifier (URI) for `tokenId` token.
     */
    function ticketURI(bytes32 tkt) external view  returns (string memory);


    /**
     * @dev Returns the total amount of tokens stored by the contract.
     */
    function totalSupply() external view returns (uint256);

}



interface ISRC721Mintable  is ISRC721 {

    function mint(address to,uint256 ticketId) external returns(bytes32);
    function burn() external ;
    function ticketId(bytes32 tkt) external view returns(uint256);
    function ticket(uint256 tktId) external view returns(bytes32);
    function setTicketURI(bytes32 tokenId, string calldata _tokenURI) external;
}