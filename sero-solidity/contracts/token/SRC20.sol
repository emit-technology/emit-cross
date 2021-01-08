pragma solidity >=0.6.0 <0.8.0;

import "../utils/seroInterface.sol";

contract SRC20  is SeroInterface  {


    string private _name;
    string private _symbol;
    uint8 private _decimals;
    uint256 private _totalSupply;

    constructor(string memory name_,string memory symbol_,uint256 totalSupply_) public payable{

        _totalSupply = totalSupply_* 10**18;

        _name = name_;

        _symbol = symbol_;

        _decimals = 18;

        require(sero_issueToken(_totalSupply,_symbol));

        require(sero_send_token(msg.sender,_symbol,_totalSupply));

    }

    /**
     * @return the name of the token.
     */
    function name() public view returns (string memory) {
        return _name;
    }

    /**
     * @return the symbol of the token.
     */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /**
     * @return the number of decimals of the token.
     */
    function decimals() public view returns (uint8) {
        return _decimals;
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }


}