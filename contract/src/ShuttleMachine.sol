// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
import "@zetachain/toolkit/contracts/BytesHelperLib.sol";
import "@zetachain/toolkit/contracts/OnlySystem.sol";
import "./interface/IPlugman.sol";


contract ShuttleMachine is zContract, OnlySystem {

    event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType);

    // cross chain zeta contract
    SystemContract public immutable systemContract;
    uint256 private immutable permitChainId;
    address public immutable plugmanAddress;
    address public immutable treasuryAddress;

    // mint types
    uint8 private constant WL_SALE = 2;
    uint8 private constant PUBLIC_SALE = 3;

    // prices
    uint256 public immutable wlPrice;
    uint256 public immutable publicSalePrice;

    error CallerNotOwnerNotApproved();
    error OriginChainNotSupported();
    error InvalidCoinValue();
    error InvalidCoinAddress();
    error InvalidSender();
    error InvalidBalance();
    error TransferToTreasuryFailed();
    error NotSupportThisMintType();

    struct CrossChainMintParams {
        address to;
        uint256 count;
        uint256 nonce;
        uint256 timestamp;
        uint8 mintType;
        bytes32[] traitValue;
        bytes signature;
    }

    constructor(
        address __systemContractAddress,
        address __plugmanAddress,
        address __treasuryAddress,
        uint256 __permitChainId,
        uint256 __wlPrice,
        uint256 __publicSalePrice) {
        systemContract = SystemContract(__systemContractAddress);
        plugmanAddress = __plugmanAddress;
        treasuryAddress = __treasuryAddress;
        permitChainId = __permitChainId;
        wlPrice = __wlPrice;
        publicSalePrice = __publicSalePrice;
    }

    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override onlySystem(systemContract) {
        CrossChainMintParams memory params;

        if (context.chainID != permitChainId) {
            revert OriginChainNotSupported();
        }

        address zrc20__ = systemContract.gasCoinZRC20ByChainId(
            context.chainID
        );
        if (zrc20 != zrc20__) revert InvalidCoinAddress();

        // if (context.sender != params.to) revert InvalidSender();

        bytes memory data = abi.decode(message, (bytes));
        params = abi.decode(data, (CrossChainMintParams));

        if (params.mintType == WL_SALE) {
            mintWL(params, zrc20, amount);
        } else if (params.mintType == PUBLIC_SALE) {
            mintPublic(params, zrc20, amount);
        } else {
            revert NotSupportThisMintType();
        }

        emit CrossChainMint(params.to, context.sender, params.count, params.nonce, params.mintType);
    }

    function mintWL(CrossChainMintParams memory params, address zrc20, uint256 amount) internal {
        if (amount < params.count * wlPrice) revert InvalidCoinValue();
        transferToTreasury(zrc20, amount);
        PlugmanInterface(plugmanAddress).mintWLCrossChain(params.to, params.count, params.nonce, params.timestamp, params.traitValue, params.signature);
    }

    function mintPublic(CrossChainMintParams memory params, address zrc20, uint256 amount) internal {
        if (amount < params.count * publicSalePrice) revert InvalidCoinValue();
        transferToTreasury(zrc20, amount);
        PlugmanInterface(plugmanAddress).mintPublicCrossChain(params.to, params.count, params.nonce, params.timestamp, params.traitValue, params.signature);
    }

    function transferToTreasury(address zrc20, uint256 amount) internal {
        uint256 balance = IZRC20(zrc20).balanceOf(address(this));
        if (balance < amount) revert InvalidBalance();
        bool transferred = IZRC20(zrc20).transfer(treasuryAddress, balance);
        if (!transferred) revert TransferToTreasuryFailed();
    }
}