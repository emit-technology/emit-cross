pragma solidity 0.6.4;

interface ISRC721Handler {

    function setResourceCategory(bytes32 resourceID, string calldata catg) external;

    function setResourceFeeHandler(bytes32 resourceID, address src20handler) external;

    function setResourceContractAddress(bytes32 resourceID,address contractAddress) external;

    function setBurnable(address contractAddress) external;

    function withdraw(string calldata catg, address recipient, bytes32 tkt) external;
}
