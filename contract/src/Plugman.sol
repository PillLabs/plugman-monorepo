// SPDX-License-Identifier: MIT
// Plugman@ZetaChain
pragma solidity ^0.8.7;

import {ERC721A, IERC721A} from "ERC721A/ERC721A.sol";
import {ERC721AQueryable} from "ERC721A/extensions/ERC721AQueryable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {Base64} from '@openzeppelin/contracts/utils/Base64.sol';
import {LibString} from "solady/src/utils/LibString.sol";
import {Solarray} from "solarray/Solarray.sol";
import "./interface/IMintVerifier.sol";
import {
    AllowedEditor,
    DisplayType,
    Editors,
    EditorsLib,
    FullTraitValue,
    TraitLabel,
    TraitLabelStorage,
    StoredTraitLabel,
    StoredTraitLabelLib
} from "shipyard-core/src/dynamic-traits/lib/TraitLabelLib.sol";
import {Metadata, DisplayType} from "shipyard-core/src/onchain/Metadata.sol";
import {json} from "shipyard-core/src/onchain/json.sol";
import {OnchainTraits, DynamicTraits} from "shipyard-core/src/dynamic-traits/OnchainTraits.sol";


/*
░▒▓███████▓▒░  ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓██████▓▒░  ░▒▓██████████████▓▒░   ░▒▓██████▓▒░  ░▒▓███████▓▒░
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓███████▓▒░  ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒▒▓███▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓████████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓████████▓▒░  ░▒▓██████▓▒░   ░▒▓██████▓▒░  ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
*/


contract Plugman is ERC721AQueryable, OnchainTraits, ReentrancyGuard {

    event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType);

    using LibString for uint256;

    bool public WLLocked;
    bool public publicSaleLocked;
    bool public plugable;

    bool private lockOwnerTraitModify;

    uint8 constant OG = 1;
    uint8 constant WL = 2;
    uint8 constant PublicSale = 3;

    // max supply
    uint256 private immutable PLUGMAN_MAX = 3369;
    uint256 private immutable MAX_PER_MINT = 5;

    // prices
    uint256 private wlPrice = 0.029 ether;
    uint256 private publicSalePrice = 0.049 ether;

    uint256 public WLStartTimestamp;
    uint256 public publicSaleStartTimestamp;
    uint256 public batchMax;
    uint256 public revealedUntil;

    mapping(address => uint256) public WLNonce;
    mapping(address => uint256) public PublicSaleNonce;
    mapping(address => bool) public shuttleMachineAddresses;

    address private immutable mintVerifier;
    address private safeRoomAddress;

    // Dynamic trait key
    bytes32 private constant BodyTraitKey = bytes32("body");
    bytes32 private constant BodyColorTraitKey = bytes32("bodyColor");
    bytes32 private constant BackgroundTraitKey = bytes32("background");
    bytes32 private constant FrontTraitKey = bytes32("front");
    bytes32 private constant SideTraitKey = bytes32("side");
    bytes32 private constant TopTraitKey = bytes32("top");

    string private imageBaseUri;

    // Errors
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


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        Constructor                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor(string memory __name, string memory __symbol, address __mintVerifier, uint256 __batchMax)
        ERC721A(__name, __symbol)
        OnchainTraits()
    {
        mintVerifier = __mintVerifier;
        batchMax = __batchMax;

        _setTraitLabel(
            BodyTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Body",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: true
            })
        );

        _setTraitLabel(
            BodyColorTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Body Color",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: true
            })
        );

        _setTraitLabel(
            BackgroundTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Background",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: true
            })
        );

        _setTraitLabel(
            FrontTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Front",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: false
            })
        );

        _setTraitLabel(
            SideTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Side",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: false
            })
        );

        _setTraitLabel(
            TopTraitKey,
            TraitLabel({
                fullTraitKey: "",
                traitLabel: "Top",
                acceptableValues: new string[](0),
                fullTraitValues: new FullTraitValue[](0),
                displayType: DisplayType.String,
                editors: Editors.wrap(EditorsLib.toBitMap(AllowedEditor.Anyone)),
                required: false
            })
        );

    }

    /**
    * @dev Returns the starting token ID for sequential mints. Overrides ERC721A
    */
    function _startTokenId() internal pure override(ERC721A) returns (uint256) {
        return 1;
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        Modifiers                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    modifier isPlugable() {
        if (safeRoomAddress == address(0) || msg.sender != safeRoomAddress || !plugable) revert IsNotPlugable();
        _;
    }

    modifier isValidTraitKey(bytes32 traitKey) {
        if (traitKey != FrontTraitKey &&
            traitKey != SideTraitKey &&
            traitKey != TopTraitKey) revert InvalidTraitKeys();
        _;
    }

    modifier isSignatureTimeout(uint256 timestamp) {
        if (block.timestamp >=  timestamp) revert InvalidTimestamp();
        _;
    }

    function isPublicSaleOpen() public view returns (bool) {
        return block.timestamp >= publicSaleStartTimestamp && publicSaleStartTimestamp != 0 && !publicSaleLocked;
    }

    function isWLSaleOpen() public view returns (bool) {
        return !isPublicSaleOpen() && block.timestamp >= WLStartTimestamp && WLStartTimestamp != 0 && !WLLocked;
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           IERC165                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721A, IERC721A, DynamicTraits) returns (bool) {
        return ERC721A.supportsInterface(interfaceId) || DynamicTraits.supportsInterface(interfaceId);
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   Only Owner Functions                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function setWLStartTimestamp(uint256 timestamp) external onlyOwner {
        WLStartTimestamp = timestamp;
    }

    function setPublicSaleTimestamp(uint256 timestamp) external onlyOwner {
        publicSaleStartTimestamp = timestamp;
    }

    function toggleWLSale() external onlyOwner {
        WLLocked = !WLLocked;
    }

    function toggleSale() external onlyOwner {
        publicSaleLocked = !publicSaleLocked;
    }

    function setBatchMax(uint256 _batchMax) external onlyOwner {
        if (_batchMax < batchMax || _batchMax > PLUGMAN_MAX) revert InvalidBatchMax();
        batchMax = _batchMax;
    }
    
    function setWLPrice(uint256 __wlPrice) external onlyOwner {
        wlPrice = __wlPrice;
    }

    function setPublicPrice(uint256 __publicSalePrice) external onlyOwner {
        publicSalePrice = __publicSalePrice;
    }

    function setSafeRoomAddress(address __safeRoomAddress) external onlyOwner {
        safeRoomAddress = __safeRoomAddress;
    }

    function setShuttleMachine(address __shuttleMachineAddress) external onlyOwner {
        shuttleMachineAddresses[__shuttleMachineAddress] = true;
    }

    function unsetShuttleMachine(address __shuttleMachineAddress) external onlyOwner {
        shuttleMachineAddresses[__shuttleMachineAddress] = false;
    }

    function togglePlugable() external onlyOwner {
        plugable = !plugable;
    }

    function setImageBaseUri(string calldata uri) external onlyOwner {
        imageBaseUri = uri;
    }

    function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) external onlyOwner {
        if (lockOwnerTraitModify) revert OwnerCannotFixTraits();
        _setTrait(tokenId, traitKey, trait);
        emit TraitUpdated(traitKey, tokenId, trait);
    }

    function revealUntil(uint256 tokenId) external onlyOwner {
        revealedUntil = tokenId;
    }

    function lockingOwnerTraitModify() external onlyOwner {
        lockOwnerTraitModify = true;
    }

    function withdraw() external onlyOwner {
        (bool success,) = msg.sender.call{value : address(this).balance}('');
        if (!success) revert WithdrawFailed();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       Mint Functions                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function mintOG(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external nonReentrant onlyOwner isSignatureTimeout(timestamp)
    {
        if (tx.origin != msg.sender) revert CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert AllTokenHasBeenMinted();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, OG, firstTokenId);
        emit Mint(to, _nonce, firstTokenId, OG);
    }

    function mintWL(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external payable nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isWLSaleOpen()) revert WLSaleIsNotActive();
        if (tx.origin != msg.sender) revert CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert AllTokenHasBeenMinted();
        if (msg.value < _count * wlPrice) revert InvalidCoinValue();
        if (_nonce <= WLNonce[msg.sender]) revert NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(msg.sender, _count, _nonce, timestamp, traitValue, signature, WL, firstTokenId);
        WLNonce[msg.sender] = _nonce;
        emit Mint(msg.sender, _nonce, firstTokenId, WL);
    }

    function mintPublicSale(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external payable nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isPublicSaleOpen()) revert PublicSaleIsNotActive();
        if (tx.origin != msg.sender) revert CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert AllTokenHasBeenMinted();
        if (msg.value < _count * publicSalePrice) revert InvalidCoinValue();
        if (_nonce <= PublicSaleNonce[msg.sender]) revert NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(msg.sender, _count, _nonce, timestamp, traitValue, signature, PublicSale, firstTokenId);
        PublicSaleNonce[msg.sender] = _nonce;
        emit Mint(msg.sender, _nonce, firstTokenId, PublicSale);
    }

    function mintWLCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
    external payable nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isWLSaleOpen()) revert WLSaleIsNotActive();
        if (!shuttleMachineAddresses[msg.sender]) revert CallerNotApproved();
        if (_count <= 0 || _count > MAX_PER_MINT) revert InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert AllTokenHasBeenMinted();
        if (_nonce <= WLNonce[to]) revert NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, WL, firstTokenId);
        WLNonce[to] = _nonce;
        emit Mint(to, _nonce, firstTokenId, WL);
    }

    function mintPublicCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
    external nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isPublicSaleOpen()) revert PublicSaleIsNotActive();
        if (!shuttleMachineAddresses[msg.sender]) revert CallerNotApproved();
        if (_count <= 0 || _count > MAX_PER_MINT) revert InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert AllTokenHasBeenMinted();
        if (_nonce <= PublicSaleNonce[to]) revert NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, PublicSale, firstTokenId);
        PublicSaleNonce[to] = _nonce;
        emit Mint(to, _nonce, firstTokenId, PublicSale);
    }

    function _verifyAndMint(
            address to,
            uint256 _count,
            uint256 _nonce,
            uint256 _timestamp,
            bytes32[] calldata traitValue,
            bytes calldata signature,
            uint8 mintType,
            uint256 firstTokenId)
            internal
    {
        // verify metadata source
        if (!mintVerify(signature, to, _count, _nonce, _timestamp, traitValue, mintType)) revert InvalidSignature();
        // ERC721A mint
        _mint(to, _count);
        // set dynamic traits
        _setTraitsAfterMint(_count, firstTokenId, traitValue);
    }

    function _setTraitsAfterMint(uint256 _count, uint256 firstTokenId, bytes32[] calldata traitValue) internal {
        if (_count * 6 != traitValue.length) revert InvalidTraitValue();
        for (uint256 i = 0; i < _count; i++) {
            _setTrait(firstTokenId + i, BodyTraitKey, traitValue[i * 6 + 0]);
            _setTrait(firstTokenId + i, BodyColorTraitKey, traitValue[i * 6 + 1]);
            _setTrait(firstTokenId + i, BackgroundTraitKey, traitValue[i * 6 + 2]);
            _setTrait(firstTokenId + i, FrontTraitKey, traitValue[i * 6 + 3]);
            _setTrait(firstTokenId + i, SideTraitKey, traitValue[i * 6 + 4]);
            _setTrait(firstTokenId + i, TopTraitKey, traitValue[i * 6 + 5]);
        }
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                  EIP-712 Verifier Functions                */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function mintVerify(
        bytes calldata _signature,
        address _mintTo,
        uint256 _amount,
        uint256 _nonce,
        uint256 _timestamp,
        bytes32[] calldata traitValue,
        uint8   _mintType
    ) public view returns (bool) {
        return IMintVerifier(mintVerifier).mintVerify(_signature, _mintTo, _amount, _nonce, _timestamp, traitValue, _mintType);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                  ERC-721 Metadata Functions                */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/


    function tokenURI(uint256 tokenId) public view override(ERC721A, IERC721A) returns (string memory) {
        string[] memory staticTraits = _staticAttributes(tokenId);
        string[] memory dynamicTraits = _dynamicAttributes(tokenId);

        string memory uri = revealedUntil > tokenId ?
            json.objectOf(
                Solarray.strings(
                    json.property("name", string.concat("Plugman #", tokenId.toString())),
                    json.property("description", string.concat("Plugman #", tokenId.toString(), ".")),
                    json.property("image", string.concat(imageBaseUri, tokenId.toString())),
                    json.rawProperty("attributes", json.arrayOf(staticTraits, dynamicTraits))
                )
            ) :
            json.objectOf(
                Solarray.strings(
                    json.property("name", string.concat("Plugman #", tokenId.toString())),
                    json.property("description", string.concat("Plugman #", tokenId.toString(), ".")),
                    json.property("image", string.concat(imageBaseUri, tokenId.toString())),
                    json.rawProperty("attributes", json.arrayOf(staticTraits))
                )
            );

        return string(abi.encodePacked('data:application/json;base64,', Base64.encode(bytes(uri))));
    }

    function _staticAttributes(uint256 tokenId) internal view returns (string[] memory) {

        return Solarray.strings(
            Metadata.attribute({traitType: "form", value: revealedUntil > tokenId ? "Plugman" : "Mystery Box", displayType: DisplayType.String})
        );
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ERC-7496 Functions                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Overrides OnchainTraits' setTrait
     * Due to plug/unplug feature, the traits can be set only when the value is empty
     */

    function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) public override isPlugable isValidTraitKey(traitKey) {
        bytes32 traitValue = getTraitValue(tokenId, traitKey);
        // only can set empty trait
        if (traitValue != bytes32(0)) revert TraitValueNotEmpty();
        _setTrait(tokenId, traitKey, trait);
        emit TraitUpdated(traitKey, tokenId, trait);
    }

    function deleteTrait(bytes32 traitKey, uint256 tokenId) public isPlugable isValidTraitKey(traitKey) {
        _setTrait(tokenId, traitKey, bytes32(0));
        emit TraitUpdated(traitKey, tokenId, bytes32(0));
    }

    /**
    * @dev Internal function that's used in Dynamic Traits to determine whether
     *      a given address is allowed to set or delete a trait for a given
     *      token ID. This function must be overridden, or the contract will not
     *      compile.
     */
    function _isOwnerOrApproved(uint256 tokenId, address addr) internal view virtual override returns (bool) {
        return ownerOf(tokenId) == addr || isApprovedForAll(ownerOf(tokenId), addr) || getApproved(tokenId) == addr;
    }

}
