// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.25;

library Errors {
    error PublicSaleIsNotActive();
    error WLSaleIsNotActive();
    error IsNotPlugable();
    error TraitValueNotEmpty();
    error CanNotCallMintFromContract();
    error InvalidMintAmount();
    error AllTokenHasBeenMinted();
    error InvalidCoinValue();
    error InvalidSignature();
    error InvalidTimestamp();
    error InvalidBatchMax();
    error InvalidTraitValue();
    error InvalidTraitKeys();
    error NonceIsTooOld();
    error OwnerCannotFixTraits();
    error WithdrawFailed();
    error CallerNotApproved();
    error RefundExcessCoinFailed();
    error ZeroAddressNotAllowed();
    error OriginChainNotSupported();
    error InvalidCoinAddress();
    error InvalidBalance();
    error TransferToTreasuryFailed();
    error NotSupportThisMintType();
}
