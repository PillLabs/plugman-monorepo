// SPDX-License-Identifier: MIT
// Plugman@ZetaChain
pragma solidity ^0.8.17;

import {Test, console} from "forge-std/Test.sol";
import {Plugman} from "../src/Plugman.sol";
import {MintVerifier} from "../src/MintVerifier.sol";
import {ShuttleMachine} from "../src/ShuttleMachine.sol";
import "../src/interface/IMintVerifier.sol";
import "../src/interface/IShuttleMachine.sol";
import "../src/interface/IPlugman.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

struct CrossChainMintParams {
    address to;
    uint256 count;
    uint256 nonce;
    uint256 timestamp;
    uint8 mintType;
    bytes32[] traitValue;
    bytes signature;
}

contract PlugmanTest is Test {
    event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType);

    PlugmanInterface public plugman;
    IMintVerifier public mintVerifier;
    IShuttleMachine public shuttleMachine;
    uint256 internal verificationPrivateKey;
    address internal verificationAddress;
    address internal shuttleMachineAddress;
    address internal plugmanAddress;
    address internal verifierAddress;
    address internal treasuryAddress;
    uint256 internal permitChainId;

    uint256 constant internal BATCH_MAX = 7;

    function setUp() public {
        verificationAddress = 0x6F881627057b37B12163118e09F7e7901096Ccea;
        treasuryAddress = 0x6F881627057b37B12163118e09F7e7901096Ccea;
        permitChainId = 31337;
        // change me corresponding to "verificationAddress"
        verificationPrivateKey = 0x123;
        MintVerifier verifierContract = new MintVerifier("Plugman", verificationAddress);
        verifierAddress = address(verifierContract);
        mintVerifier = IMintVerifier(verifierAddress);
        Plugman plugmanContract = new Plugman("Plugman", "PLGM", verifierAddress, BATCH_MAX);
        plugmanAddress = address(plugmanContract);
        plugman = PlugmanInterface(plugmanAddress);
        uint256 crossingWlPrice = 0.029 ether;
        uint256 crossingPublicSalePrice = 0.049 ether;
        ShuttleMachine shuttleMachineContract = new ShuttleMachine(
            address(0x123), plugmanAddress, treasuryAddress, permitChainId, crossingWlPrice, crossingPublicSalePrice);
        shuttleMachineAddress = address(shuttleMachineContract);
        shuttleMachine = IShuttleMachine(shuttleMachine);

        uint256 _timestamp = 1717844000;
        uint256 mockBlockchainTimestamp = 1717840000;
        vm.warp(mockBlockchainTimestamp);
        mintOG(address(this), address(0x123), 1, 0, _timestamp);
    }

    function test_batchMax() public {
        uint256 batchMax = plugman.batchMax();
        assertEqUint(batchMax, BATCH_MAX);
        uint256 newBatchMax = batchMax - 1;
        vm.expectRevert(InvalidBatchMax.selector);
        plugman.setBatchMax(newBatchMax);
        newBatchMax = 3370;
        vm.expectRevert(InvalidBatchMax.selector);
        plugman.setBatchMax(newBatchMax);
        newBatchMax = batchMax + 1;
        plugman.setBatchMax(newBatchMax);
        assertEqUint(newBatchMax, plugman.batchMax());
    }

    function test_isPublicSaleOpen() public {
        assert(!plugman.isPublicSaleOpen());
        uint256 publicSaleTimestamp = 1717844525;
        uint256 mockBlockchainTimestamp = 1717844000;
        plugman.setPublicSaleTimestamp(publicSaleTimestamp);
        vm.warp(mockBlockchainTimestamp);
        assert(!plugman.isPublicSaleOpen());
        mockBlockchainTimestamp = 1717849000;
        vm.warp(mockBlockchainTimestamp);
        assert(plugman.isPublicSaleOpen());
        plugman.toggleSale();
        assert(!plugman.isPublicSaleOpen());
    }

    function test_isWLSaleOpen() public {
        assert(!plugman.isWLSaleOpen());
        uint256 WLSaleTimestamp = 1717844520;
        uint256 publicSaleTimestamp = 1717844525;
        uint256 mockBlockchainTimestamp = 1717844000;
        plugman.setWLStartTimestamp(WLSaleTimestamp);
        vm.warp(mockBlockchainTimestamp);
        assert(!plugman.isWLSaleOpen());
        mockBlockchainTimestamp = 1717849000;
        vm.warp(mockBlockchainTimestamp);
        assert(plugman.isWLSaleOpen());
        plugman.toggleWLSale();
        assert(!plugman.isWLSaleOpen());
        plugman.toggleWLSale();
        assert(plugman.isWLSaleOpen());
        plugman.setPublicSaleTimestamp(publicSaleTimestamp);
        assert(!plugman.isWLSaleOpen());
    }

    function test_withdraw() public {
        test_mintPublicSale();
        vm.prank(address(0x123));
        vm.expectRevert(Unauthorized.selector);
        uint256 balance = plugmanAddress.balance;
        plugman.withdraw();
        assertEqUint(balance, 0.049 ether);
        balance = plugmanAddress.balance;
        plugman.withdraw();
        assertEqUint(balance, 0.049 ether);
        balance = plugmanAddress.balance;
        assertEqUint(balance, 0);
    }

    function test_mintOG() public {
        address _from = address(this);
        address _to = 0x82459c94a2724EEE83e369847F27De9EDaD373e4;
        uint256 _count = 3;
        uint256 _nonce = 0;
        uint256 _timestamp = 1717844000;
        uint256 mockBlockchainTimestamp = 1717840000;
        vm.warp(mockBlockchainTimestamp);
        mintOG(_from, _to, _count, _nonce, _timestamp);
    }

    function test_tokenURI() public {
        string memory imageBaseUri = "https://test.plugman.nft/images/";
        plugman.setImageBaseUri(imageBaseUri);
        string memory base64DecodedTokenUri = plugman.tokenURI(1);
        string memory expectedBase64 = "data:application/json;base64,eyJuYW1lIjoiUGx1Z21hbiAjMSIsImRlc2NyaXB0aW9uIjoiUGx1Z21hbiAjMS4iLCJpbWFnZSI6Imh0dHBzOi8vdGVzdC5wbHVnbWFuLm5mdC9pbWFnZXMvMSIsImF0dHJpYnV0ZXMiOlt7InRyYWl0X3R5cGUiOiJmb3JtIiwidmFsdWUiOiJNeXN0ZXJ5IEJveCIsImRpc3BsYXlfdHlwZSI6InN0cmluZyJ9XX0=";
        assertEq0(bytes(base64DecodedTokenUri), bytes(expectedBase64));
        plugman.revealUntil(7);
        base64DecodedTokenUri = plugman.tokenURI(1);
        expectedBase64 = "data:application/json;base64,eyJuYW1lIjoiUGx1Z21hbiAjMSIsImRlc2NyaXB0aW9uIjoiUGx1Z21hbiAjMS4iLCJpbWFnZSI6Imh0dHBzOi8vdGVzdC5wbHVnbWFuLm5mdC9pbWFnZXMvMSIsImF0dHJpYnV0ZXMiOlt7InRyYWl0X3R5cGUiOiJmb3JtIiwidmFsdWUiOiJQbHVnbWFuIiwiZGlzcGxheV90eXBlIjoic3RyaW5nIn0seyJ0cmFpdF90eXBlIjoiQm9keSIsInZhbHVlIjoiYWEiLCJkaXNwbGF5X3R5cGUiOiJzdHJpbmcifSx7InRyYWl0X3R5cGUiOiJCb2R5IENvbG9yIiwidmFsdWUiOiJiYiIsImRpc3BsYXlfdHlwZSI6InN0cmluZyJ9LHsidHJhaXRfdHlwZSI6IkJhY2tncm91bmQiLCJ2YWx1ZSI6ImNjIiwiZGlzcGxheV90eXBlIjoic3RyaW5nIn0seyJ0cmFpdF90eXBlIjoiRnJvbnQiLCJ2YWx1ZSI6ImRkIiwiZGlzcGxheV90eXBlIjoic3RyaW5nIn0seyJ0cmFpdF90eXBlIjoiU2lkZSIsInZhbHVlIjoiZWUiLCJkaXNwbGF5X3R5cGUiOiJzdHJpbmcifSx7InRyYWl0X3R5cGUiOiJUb3AiLCJ2YWx1ZSI6ImZmIiwiZGlzcGxheV90eXBlIjoic3RyaW5nIn1dfQ==";
        assertEq0(bytes(base64DecodedTokenUri), bytes(expectedBase64));
    }

    function test_safeRoomSetTraits() public {
        bytes32 sideTraitKey = bytes32("side");
        vm.expectRevert(IsNotPlugable.selector);
        plugman.setTrait(1, sideTraitKey, bytes32("NGMI"));
        address safeRoomAddress = address(0x123);
        plugman.setSafeRoomAddress(safeRoomAddress);
        vm.prank(safeRoomAddress);
        vm.expectRevert(IsNotPlugable.selector);
        plugman.setTrait(1, sideTraitKey, bytes32("NGMI"));
        plugman.togglePlugable();
        vm.prank(safeRoomAddress);
        vm.expectRevert(TraitValueNotEmpty.selector);
        plugman.setTrait(1, sideTraitKey, bytes32("NGMI"));
        bytes32 noneSenseTraitKey = bytes32("noneSense");
        vm.startPrank(safeRoomAddress);
        vm.expectRevert(InvalidTraitKeys.selector);
        plugman.setTrait(1, noneSenseTraitKey, bytes32("NGMI"));
        plugman.deleteTrait(sideTraitKey, 1);
        plugman.setTrait(1, sideTraitKey, bytes32("WAGMI"));
        assertEq32(plugman.getTraitValue(1, sideTraitKey), bytes32("WAGMI"));
    }

    function test_traits() public {
        bytes32 frontTraitKey = bytes32("front");
        bytes32 frontTraitValue = plugman.getTraitValue(1, frontTraitKey);
        assertEq32(frontTraitValue, bytes32("dd"));
        bytes32 newFrontTraitValue = bytes32("dddd");
        plugman.ownerFixTraits(1, frontTraitKey, newFrontTraitValue);
        frontTraitValue = plugman.getTraitValue(1, frontTraitKey);
        assertEq32(frontTraitValue, newFrontTraitValue);
        plugman.lockingOwnerTraitModify();
        vm.expectRevert(OwnerCannotFixTraits.selector);
        newFrontTraitValue = bytes32("ddddddd");
        plugman.ownerFixTraits(1, frontTraitKey, newFrontTraitValue);
    }

    function test_SetPublicStartTimestamp() public {
        uint256 timestamp = 1717844525;
        plugman.setPublicSaleTimestamp(timestamp);
    }

    function test_mintPublicSale() public payable {
        uint256 depositAmount = 0.049 ether;
        vm.warp(1717844600);
        uint256 count = 1;
        uint256 nonce = 1;
        uint256 publicTimestamp = 1717844524;
        uint256 timestamp = 1717844620;
        bytes32[] memory traitValue = getTraitValues(count);
        plugman.setPublicSaleTimestamp(publicTimestamp);
        address user = address(0x6F881627057b37B12163118e09F7e7901096Ccea);

        bytes memory signature = signMintEIP712(user, count, nonce, timestamp, traitValue, 3);
        vm.prank(user, user);
        vm.deal(user, depositAmount);
        plugman.mintPublicSale{value: depositAmount}(count, nonce, timestamp, traitValue, signature);
    }

    function test_mintWL() public {
        uint256 depositAmount = 0.029 ether * 5;
        uint256 _count = 5;
        uint256 _nonce = 5;
        address _to = address(0x123);
        bytes32[] memory traitValue = getTraitValues(_count);
        uint256 WLTimestamp = 1717844524;
        uint256 _timestamp = 1717844600;
        uint256 blockchainTimestamp = 1717844510;
        vm.deal(_to, 1 ether);
        vm.warp(blockchainTimestamp);
        plugman.setWLStartTimestamp(WLTimestamp);
        vm.startPrank(_to, _to);
        bytes memory signature = signMintEIP712(_to, _count, _nonce, _timestamp, traitValue, 2);

        vm.expectRevert(WLSaleIsNotActive.selector);
        plugman.mintWL{value: depositAmount}(_count, _nonce, _timestamp, traitValue, signature);

        blockchainTimestamp = 1717844580;
        vm.warp(blockchainTimestamp);
        vm.expectRevert(InvalidMintAmount.selector);
        plugman.mintWL{value: depositAmount}(9, _nonce, _timestamp, traitValue, signature);

        vm.expectRevert(InvalidCoinValue.selector);
        plugman.mintWL{value: 0.1 ether}(_count, _nonce, _timestamp, traitValue, signature);

        uint256 totalSupplyBeforeMint = plugman.totalSupply();
        uint256 nextTokenId = totalSupplyBeforeMint + 1;
        vm.expectEmit(true, false, false, true);
        emit Mint(_to, _nonce, nextTokenId, 2);
        plugman.mintWL{value: depositAmount}(_count, _nonce, _timestamp, traitValue, signature);

        vm.expectRevert(AllTokenHasBeenMinted.selector);
        plugman.mintWL{value: depositAmount}(_count, 6, _timestamp, traitValue, signature);

        vm.expectRevert(NonceIsTooOld.selector);
        plugman.mintWL{value: 0.029 ether}(1, 3, _timestamp, traitValue, signature);

        _count = 1;
        _nonce = 7;
        signature = signMintEIP712(_to, _count, _nonce, _timestamp, traitValue, 2);
        vm.expectRevert(InvalidTraitValue.selector);
        plugman.mintWL{value: 0.029 ether}(_count, _nonce, _timestamp, traitValue, signature);

        traitValue = getTraitValues(_count);
        signature = signMintEIP712(_to, _count, _nonce, _timestamp, traitValue, 2);
        totalSupplyBeforeMint = plugman.totalSupply();
        nextTokenId = totalSupplyBeforeMint + 1;
        vm.expectEmit(true, false, false, true);
        emit Mint(_to, _nonce, nextTokenId, 2);
        plugman.mintWL{value: 0.029 ether}(_count, _nonce, _timestamp, traitValue, signature);
        uint256[] memory tokenIds = plugman.tokensOfOwner(_to);
        uint256 len = tokenIds.length;
        assertEqUint(tokenIds[len - 1], 7);
    }

    function mintOG(address _from, address _to, uint256 _count, uint256 _nonce, uint256 _timestamp) public {
        bytes32[] memory traitValue = getTraitValues(_count);
        bytes memory signature = signMintEIP712(_to, _count, _nonce, _timestamp, traitValue, 1);
        uint256 totalSupplyBeforeMint = plugman.totalSupply();
        uint256 nextTokenId = totalSupplyBeforeMint + 1;
        vm.prank(_from, _from);
        vm.expectEmit(true, false, false, true);
        emit Mint(_to, _nonce, nextTokenId, 1);
        plugman.mintOG(_to, _count, _nonce, _timestamp, traitValue, signature);
    }

    function getTraitValues(uint256 _count) public pure returns (bytes32[] memory){
        uint256 length = _count * 6;
        bytes32[] memory traitValues = new bytes32[](length);
        for (uint256 i = 0; i < _count; i++) {
            uint256 offset = i * 6;
            traitValues[0 + offset] = bytes32("aa");
            traitValues[1 + offset] = bytes32("bb");
            traitValues[2 + offset] = bytes32("cc");
            traitValues[3 + offset] = bytes32("dd");
            traitValues[4 + offset] = bytes32("ee");
            traitValues[5 + offset] = bytes32("ff");
        }
        return traitValues;
    }

    function signMintEIP712(
        address _mintTo, uint256 _count, uint256 _nonce, uint256 _timestamp, bytes32[] memory _traitValues, uint8 _mintType)
        internal view returns (bytes memory) {

        // Domain separator
        string memory name = plugman.name();
        bytes32 TYPE_HASH = keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
        bytes32 _hashedName = keccak256(bytes(name));
        bytes32 _hashedVersion = keccak256(bytes('1'));
        bytes32 domainSeparator =  keccak256(abi.encode(TYPE_HASH, _hashedName, _hashedVersion, block.chainid, address(verifierAddress)));

        // typed data
        bytes32 digest = _hashTypedDataV4(
            domainSeparator,
            keccak256(
                abi.encode(
                    keccak256(
                        'PlugmanMint(address mintTo,uint256 amount,uint256 nonce,uint256 timestamp,bytes32 traitHash,uint8 mintType)'
                    ),
                    _mintTo,
                    _count,
                    _nonce,
                    _timestamp,
                    keccak256(abi.encodePacked(_traitValues)),
                    _mintType
                )
            )
        );
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(verificationPrivateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    function _hashTypedDataV4(bytes32 domainSeparator, bytes32 structHash) internal view virtual returns (bytes32) {
        return MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
    }

    receive() external payable {}
}