// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {EIP712} from '@openzeppelin/contracts/utils/cryptography/EIP712.sol';
import {ECDSA} from '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import {Errors} from "./library/Errors.sol";

/**
 * @title MintVerifier
 * @author Plugman
 * @dev This contract implements an EIP712 verifier.
 * Plugman adopts dynamic traits stored on-chain. Trait metadata is
 * generated off-chain and brought on-chain during minting.
 * MintVerifier is used to verify if the traits are generated and signed by the
 * authorized generator.
 */

contract MintVerifier is EIP712 {

    // metadataCreator The address of the off-chain generator
    address private immutable metadataCreator;

    /**
    * @dev Constructor.
    * MintVerifier is constructed using the EIP712 constructor function.
    * @param __name Name of the EIP712 Domain
    * @param __metadataCreator Address of the off-chain generator
    */
    constructor(string memory __name, address __metadataCreator ) EIP712(__name, '1') {
        if (__metadataCreator == address(0)) revert Errors.ZeroAddressNotAllowed();
        metadataCreator = __metadataCreator;
    }


    /**
    * @dev EIP712 signature verification function. It only verifies the signature,
    * returning a boolean indicating if the signature is valid. Checks for other parameters
    * are handled in the Plugman NFT contract.
    * @param _signature Signature of the metadata from minting, signed by the off-chain trait generator
    * @param _mintTo Address that should receive the minted NFTs
    * @param _amount Amount of NFTs to be minted
    * @param _nonce Nonce maintained by the trait generator, used to track minting
    * @param _timestamp Expiry time of the signature; timestamp checks are performed in the Plugman NFT contract
    * @param traitValue Array of metadata recording the traits of the minted NFT
    * @param _mintType Minting type
    */
    function mintVerify(
        bytes calldata _signature,
        address _mintTo,
        uint256 _amount,
        uint256 _nonce,
        uint256 _timestamp,
        bytes32[] calldata traitValue,
        uint8   _mintType
    ) external view returns (bool) {
        bytes32 digest = _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256(
                        'PlugmanMint(address mintTo,uint256 amount,uint256 nonce,uint256 timestamp,bytes32 traitHash,uint8 mintType)'
                    ),
                    _mintTo,
                    _amount,
                    _nonce,
                    _timestamp,
                    keccak256(abi.encodePacked(traitValue)),
                    _mintType
                )
            )
        );
        return metadataCreator == ECDSA.recover(digest, _signature);
    }
}