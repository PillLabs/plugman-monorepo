// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.25;

interface IShuttleMachine {
    struct zContext {
        bytes origin;
        address sender;
        uint256 chainID;
    }

    error InvalidBalance();
    error InvalidCoinAddress();
    error InvalidCoinValue();
    error NotSupportThisMintType();
    error OnlySystemContract(string);
    error OriginChainNotSupported();
    error TransferToTreasuryFailed();
    error ZeroAddressNotAllowed();

    event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType);

    function onCrossChainCall(zContext memory context, address zrc20, uint256 amount, bytes memory message) external;
    function plugmanAddress() external view returns (address);
    function publicSalePrice() external view returns (uint256);
    function systemContract() external view returns (address);
    function treasuryAddress() external view returns (address);
    function wlPrice() external view returns (uint256);
}
