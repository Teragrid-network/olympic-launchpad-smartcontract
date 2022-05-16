//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Context.sol";

contract PrivateSale is Ownable {
    struct User {
        uint256 buyLimit;
        uint256 boughtAmount;
        uint256 claimAmount;
    }

    IERC20 idoToken;
    //IERC20 token;
    uint256 openSaleTime;
    uint256 endSaleTime;
    uint256 claimableTime;
    //int256 periodAmount;
    //int256 periodDuration;
    uint256 tokenToIDOTokenRatio; // 1AVAX = 100 IDOToken -> tokenToIDOTokenRatio = 100

    mapping(address => User) addressToUserDetail;
    uint256 totalBuy;

    constructor(
        address _idoTokenAddress,
        uint256 _openSaleTime,
        uint256 _endSaleTime,
        uint256 _claimableTime,
        uint256 _tokenToIDOTokenRatio,
        address[] memory _accounts,
        uint256[] memory _limitAmount
    ) {
        require(
            block.timestamp <= _openSaleTime,
            "PrivateSale/constructor: openSaleTime must after current time"
        );
        require(
            _openSaleTime < _endSaleTime,
            "PrivateSale/constructor: endSaleTime must after openSaleTime"
        );

        require(
            _endSaleTime <= _claimableTime,
            "PrivateSale/constructor: claimableTime must after endSaleTime"
        );
        require(
            _tokenToIDOTokenRatio > 0,
            "PrivateSale/constructor: invalid ratio"
        );

        idoToken = IERC20(_idoTokenAddress);
        openSaleTime = _openSaleTime;
        endSaleTime = _endSaleTime;
        claimableTime = _claimableTime;
        tokenToIDOTokenRatio = _tokenToIDOTokenRatio;

        //SNAPSHOT
        require(
            _accounts.length > 0 && _limitAmount.length > 0,
            "PrivateSale/constructor: null array"
        );
        require(
            _accounts.length == _limitAmount.length,
            "PrivateSale/constructor: not equal array"
        );

        for (uint256 i = 0; i < _accounts.length; i++) {
            require(
                _limitAmount[i] > 0,
                "PrivateSale/constructor: invalid limitAmount"
            );
            require(
                addressToUserDetail[_accounts[i]].buyLimit == 0,
                "PrivateSale/constructor: user is already add"
            );

            addressToUserDetail[_accounts[i]].buyLimit = _limitAmount[i];
        }
    }

    function buy() external payable onlyUser {
        require(
            block.timestamp >= openSaleTime && block.timestamp <= endSaleTime,
            "PrivateSale/buy: not time to buy"
        );

        uint256 buyAmount = msg.value * tokenToIDOTokenRatio;

        require(
            addressToUserDetail[msg.sender].buyLimit >=
                addressToUserDetail[msg.sender].boughtAmount + buyAmount,
            "PrivateSale/buy: max buylimit"
        );

        totalBuy += buyAmount;
        addressToUserDetail[msg.sender].boughtAmount += buyAmount;
    }

    function claim() external onlyUser {
        require(
            block.timestamp >= claimableTime,
            "PrivateSale/claim: not time to claim"
        );
        uint256 tokenCanClaimedAmount = getClaimableAmount(msg.sender);
        // maybe not need this require
        require(
            tokenCanClaimedAmount > 0,
            "PrivateSale/claim: not enough to claim"
        );
        idoToken.transfer(msg.sender, tokenCanClaimedAmount);

        addressToUserDetail[msg.sender].claimAmount += tokenCanClaimedAmount;
    }

    //TODO: function to check if amount in the contract is enough for all users to claim
    //TODO: special case: user send shit token to contract
    function withdraw() external onlyOwner {
        require(
            idoToken.balanceOf(address(this)) - totalBuy >= 0,
            "PrivateSale/withdraw: not enough to withdraw"
        );

        payable(msg.sender).transfer(address(this).balance);
        uint256 remainToken = idoToken.balanceOf(address(this)) - totalBuy;
        idoToken.transfer(owner(), remainToken);
    }

    function getClaimableAmount(address _account)
        public
        view
        returns (uint256)
    {
        if (block.timestamp < claimableTime) {
            return 0;
        }
        return
            addressToUserDetail[_account].boughtAmount -
            addressToUserDetail[_account].claimAmount;
    }

    function getUserDetail(address _account)
        external
        view
        returns (
            uint256 buyLimit,
            uint256 boughtAmount,
            uint256 claimAmount
        )
    {
        return (
            addressToUserDetail[_account].buyLimit,
            addressToUserDetail[_account].boughtAmount,
            addressToUserDetail[_account].claimAmount
        );
    }

    modifier onlyUser() {
        require(
            addressToUserDetail[msg.sender].buyLimit > 0,
            "PrivateSale/onlyUser: not user"
        );
        _;
    }
}
