pragma solidity >=0.6.0 <0.8.0;

import "github.com/OpenZeppelin/openzeppelin-contracts/contracts/math/SafeMath.sol";

import "../utils/seroInterface.sol";
import "./AccessControl.sol";

contract SRC20Mintable  is SeroInterface,AccessControl  {


    using SafeMath for uint256;

    string private _name;
    string private _symbol;
    uint8 private _decimals;
    uint256 private _totalSupply;

    constructor(string memory name_,string memory symbol_,uint8 decimals_) public payable{

        _name = name_;

        _symbol = symbol_;

        _decimals = decimals_;

        require(sero_issueToken(_totalSupply,_symbol));

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



    function mint(uint256 total) public onlyMinter {

        uint256 currentBalance = sero_balanceOf(_symbol);

        if (currentBalance<total){

            uint256 _mintAmount =total.sub(currentBalance);

            require(sero_issueToken(_mintAmount,_symbol),"mint failed");

        }
        _totalSupply = _totalSupply.add(total);

        require(sero_send_token(msg.sender,_symbol,total),"send mint failed");

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


    function burn()public payable {

        string memory cy = sero_msg_currency();

        require(equals(cy,_symbol),"not supported cy");

        _totalSupply = _totalSupply.sub(msg.value);

    }
}