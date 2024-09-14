// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {EIP712} from '@openzeppelin/contracts/utils/cryptography/EIP712.sol';
import {ECDSA} from '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';

contract MintVerifier is EIP712 {

    address private immutable metadataCreator;

    constructor(string memory __name, address __metadataCreator ) EIP712(__name, '1') {
        metadataCreator = __metadataCreator;
    }

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