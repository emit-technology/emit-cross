// SPDX-License-Identifier: MIT

pragma solidity >=0.6.0 <0.8.0;

abstract contract AccessControl{

    event MinterAdded(address indexed minter);
    event MinterRemoved(address indexed minter);

    address public owner;

    address public executor;


    mapping(address=>bool) minters;

    constructor() public {
        owner = msg.sender;
        executor = msg.sender;
    }


    modifier onlyOwner() {
        require(msg.sender == owner,"not owner");
        _;
    }

    modifier onlyExecutor() {
        require(msg.sender == executor,"not executor");
        _;
    }

    modifier onlyMinter() {
        require(minters[msg.sender],"not minter");
        _;
    }

    function transferOwnership(address newOwner) public onlyOwner {
        if (newOwner != address(0)) {
            owner = newOwner;
        }
    }

    function delExecutor() public onlyOwner {
        executor = address(0);
    }

    function changeExecutor(address newExecutor) public onlyExecutor {
        executor = newExecutor;
    }


    function addMinter(address user) public onlyExecutor{
        require(!minters[user],"is a minter");
        minters[user] =true;
        emit MinterAdded(user);

    }

    function delMinter(address user) public onlyExecutor {
        require(minters[user],"not a minter");
        delete minters[user];
        emit MinterRemoved(user);
    }

    function dropMintable() public {
        require(minters[msg.sender],"not is minter");
        delete minters[msg.sender];
        emit MinterRemoved(msg.sender);
    }
}
