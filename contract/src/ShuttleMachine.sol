// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "@zetachain/protocol-contracts/contracts/zevm/SystemContract.sol";
import "@zetachain/protocol-contracts/contracts/zevm/interfaces/zContract.sol";
import "@zetachain/toolkit/contracts/BytesHelperLib.sol";
import "@zetachain/toolkit/contracts/OnlySystem.sol";
import "./interface/IPlugman.sol";
import {Errors} from "./library/Errors.sol";

/**
 * @title ShuttleMachine
 * @author Plugman
 * @dev This contract that handles cross-chain minting of tokens,
 * inheriting from zContract and OnlySystem.
 */
contract ShuttleMachine is zContract, OnlySystem {

    // Event that logs cross-chain minting operations
    event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType);

    // cross chain zeta contract
    SystemContract public immutable systemContract;
    // Specific chain ID where minting is permitted
    uint256 private immutable permitChainId;
    // Address of the Plugman contract which handles the actual token minting
    address public immutable plugmanAddress;
    // Address of the treasury where funds are collected
    address public immutable treasuryAddress;

    // Mint types identifiers
    uint8 private constant WL_SALE = 2;
    uint8 private constant PUBLIC_SALE = 3;

    // Fixed prices of the crossed-chain native coin for whitelist sales
    // and public sales
    uint256 public immutable wlPrice;
    uint256 public immutable publicSalePrice;

    // Struct to hold parameters for cross-chain minting operations
    struct CrossChainMintParams {
        address to;
        uint256 count;
        uint256 nonce;
        uint256 timestamp;
        uint8 mintType;
        bytes32[] traitValue;
        bytes signature;
    }

    /**
    * @dev Constructor.
    * Initialize contract with specific addresses and prices
    * @param __systemContractAddress Address of ZetaChain system contract
    * @param __plugmanAddress Address of Plugman contract
    * @param __treasuryAddress Address of the treasury where funds are collected
    * @param __permitChainId Specific chain ID where minting is permitted
    * @param __wlPrice Fixed prices of the crossed-chain native coin
    * @param __publicSalePrice Fixed prices of the crossed-chain native coin
    */
    constructor(
        address __systemContractAddress,
        address __plugmanAddress,
        address __treasuryAddress,
        uint256 __permitChainId,
        uint256 __wlPrice,
        uint256 __publicSalePrice) {
        if (__systemContractAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        if (__plugmanAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        if (__treasuryAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        systemContract = SystemContract(__systemContractAddress);
        plugmanAddress = __plugmanAddress;
        treasuryAddress = __treasuryAddress;
        permitChainId = __permitChainId;
        wlPrice = __wlPrice;
        publicSalePrice = __publicSalePrice;
    }

    // Function to handle incoming cross-chain calls, restricted to the system context
    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    ) external override onlySystem(systemContract) {
        CrossChainMintParams memory params;

        if (context.chainID != permitChainId) {
            revert Errors.OriginChainNotSupported();
        }

        address zrc20__ = systemContract.gasCoinZRC20ByChainId(
            context.chainID
        );
        if (zrc20 != zrc20__) revert Errors.InvalidCoinAddress();

        bytes memory data = abi.decode(message, (bytes));
        params = abi.decode(data, (CrossChainMintParams));

        if (params.mintType == WL_SALE) {
            mintWL(params, zrc20, amount);
        } else if (params.mintType == PUBLIC_SALE) {
            mintPublic(params, zrc20, amount);
        } else {
            revert Errors.NotSupportThisMintType();
        }

        emit CrossChainMint(params.to, context.sender, params.count, params.nonce, params.mintType);
    }

    function mintWL(CrossChainMintParams memory params, address zrc20, uint256 amount) internal {
        if (amount < params.count * wlPrice) revert Errors.InvalidCoinValue();
        transferToTreasury(zrc20, amount);
        PlugmanInterface(plugmanAddress).mintWLCrossChain(params.to, params.count, params.nonce, params.timestamp, params.traitValue, params.signature);
    }

    function mintPublic(CrossChainMintParams memory params, address zrc20, uint256 amount) internal {
        if (amount < params.count * publicSalePrice) revert Errors.InvalidCoinValue();
        transferToTreasury(zrc20, amount);
        PlugmanInterface(plugmanAddress).mintPublicCrossChain(params.to, params.count, params.nonce, params.timestamp, params.traitValue, params.signature);
    }

    function transferToTreasury(address zrc20, uint256 amount) internal {
        uint256 balance = IZRC20(zrc20).balanceOf(address(this));
        if (balance < amount) revert Errors.InvalidBalance();
        bool transferred = IZRC20(zrc20).transfer(treasuryAddress, balance);
        if (!transferred) revert Errors.TransferToTreasuryFailed();
    }
}