// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.4;

interface IMintVerifier {
    error ECDSAInvalidSignature();
    error ECDSAInvalidSignatureLength(uint256 length);
    error ECDSAInvalidSignatureS(bytes32 s);
    error InvalidShortString();
    error StringTooLong(string str);

    event EIP712DomainChanged();

    function eip712Domain()
    external
    view
    returns (
        bytes1 fields,
        string memory name,
        string memory version,
        uint256 chainId,
        address verifyingContract,
        bytes32 salt,
        uint256[] memory extensions
    );
    function mintVerify(
        bytes memory _signature,
        address _mintTo,
        uint256 _amount,
        uint256 _nonce,
        uint256 _timestamp,
        bytes32[] memory traitValue,
        uint8 _mintType
    ) external view returns (bool);
}
