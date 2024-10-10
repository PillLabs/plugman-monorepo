// SPDX-License-Identifier: MIT
// Plugman@ZetaChain
pragma solidity 0.8.25;

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
    TraitLabel
} from "shipyard-core/src/dynamic-traits/lib/TraitLabelLib.sol";
import {Metadata, DisplayType} from "shipyard-core/src/onchain/Metadata.sol";
import {json} from "shipyard-core/src/onchain/json.sol";
import {OnchainTraits, DynamicTraits} from "shipyard-core/src/dynamic-traits/OnchainTraits.sol";
import {Errors} from "./library/Errors.sol";


/*
░▒▓███████▓▒░  ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓██████▓▒░  ░▒▓██████████████▓▒░   ░▒▓██████▓▒░  ░▒▓███████▓▒░
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓███████▓▒░  ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒▒▓███▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓████████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓█▓▒░        ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
░▒▓█▓▒░        ░▒▓████████▓▒░  ░▒▓██████▓▒░   ░▒▓██████▓▒░  ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░░▒▓█▓▒░
*/

/**
 * @title Plugman
 * @author Plugman
 * @dev This contract implements an ERC-721A NFT contract compatible with
 * the ERC-7496 interface.
 * It includes the necessary minting implementations and interfaces for
 * future contract integrations.
 * Clients on the ZetaChain can interact directly with this contract to mint.
 * Clients on other chains (Ethereum and Polygon) need to interact through
 * ZetaChain's system contract to call the ShuttleMachine contract (for
 * relaying cross-chain minting requests).
 * This contract uses the MintVerifier contract to verify the metadata
 * signature.
 */
contract Plugman is ERC721AQueryable, OnchainTraits, ReentrancyGuard {

    /**
    * @dev Emitted after minting
    * @param to Address that should receive the minted NFTs
    * @param nonce Used by the off-chain trait generator to maintain and
    * track minting
    * @param firstTokenId Token ID of the first token minted in this transaction
    * @param mintType Type of minting
    */
    event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType);
    event SetTimestamp(uint256 timestamp, uint8 mintType);
    event ToggleSaleStatus(uint8 mintType);
    event SetSalePrice(uint256 price, uint8 mintType);
    event SetBatchMax(uint256 batchMax);
    event SetSafeRoom(address newSafeRoomAddress);
    event SetShuttleMachine(address shuttleMachineAddress);
    event UnsetShuttleMachine(address shuttleMachineAddress);
    event TogglePlugable();
    event SetImageBaseUri();
    event Withdraw(address receiver, uint256 amount);
    event RevealUntil(uint256 tokenId);
    event LockOwnerTraitModifying();


    using LibString for uint256;

    // Indicates if the whitelist sale is locked
    bool public WLLocked;
    // Indicates if the public sale is locked
    bool public publicSaleLocked;
    // Indicates if Plugman NFTs are plugable.
    // One Plugman NFT owns three plugable traits which will also be NFTs in
    // the future.
    bool public plugable;

    // Restricts the contract owner from modifying traits metadata
    // This may be necessary to fix metadata bugs after minting, if any
    bool private lockOwnerTraitModify;

    // MintType definitions
    uint8 constant OG = 1;
    uint8 constant WL = 2;
    uint8 constant PublicSale = 3;

    // Maximum supply of NFTs
    uint256 private immutable PLUGMAN_MAX = 3369;
    // Maximum number of NFTs mintable per transaction
    uint256 private immutable MAX_PER_MINT = 5;

    // Pricing for different sales
    uint256 private wlPrice = 0.029 ether;
    uint256 private publicSalePrice = 0.049 ether;

    // Timestamp when the whitelist sale starts
    uint256 public WLStartTimestamp;
    // Timestamp when the public sale starts
    uint256 public publicSaleStartTimestamp;
    // Token ID at which the current batch of minting ends.
    // Plugman mints in batches.
    uint256 public batchMax;
    // Timestamp until which the Plugman metadata is revealed
    uint256 public revealedUntil;

    // Tracking nonce for the whitelist sale
    mapping(address => uint256) public WLNonce;
    // Tracking nonce for the public sale
    mapping(address => uint256) public PublicSaleNonce;
    // Addresses of "ShuttleMachine Contracts" - serve as cross-chain
    // minting gateways.
    mapping(address => bool) public shuttleMachineAddresses;

    // Address of "MintVerifier Contract" - verifies metadata generated
    // and signed by the off-chain generator
    address private immutable mintVerifier;
    // Address of "SafeRoom Contract" - future feature for plugging and
    // unplugging trait NFTs
    address private safeRoomAddress;

    // Keys for dynamic traits
    bytes32 private constant BodyTraitKey = bytes32("body");
    bytes32 private constant BodyColorTraitKey = bytes32("bodyColor");
    bytes32 private constant BackgroundTraitKey = bytes32("background");
    bytes32 private constant FrontTraitKey = bytes32("front");
    bytes32 private constant SideTraitKey = bytes32("side");
    bytes32 private constant TopTraitKey = bytes32("top");

    // Base URI for NFT images
    string private imageBaseUri;


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        Constructor                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
    * @dev Constructor.
    * Plugman is constructed using the ERC721A constructor function and
    * the ERC7496 constructor function.
    * This constructor also create storage for trait labels.
    * @param __name Name of ERC721 NFT
    * @param __symbol Symbol of ERC721 NFT
    * @param __mintVerifier Address of "MintVerifier Contract"
    * @param __batchMax Token ID at which the current batch of minting ends
    */
    constructor(string memory __name, string memory __symbol, address __mintVerifier, uint256 __batchMax)
        ERC721A(__name, __symbol)
        OnchainTraits()
    {
        if (__mintVerifier == address(0)) revert Errors.ZeroAddressNotAllowed();
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

    /**
    * @dev Modifier to check if plugging or unplugging operations are allowed.
    */
    modifier isPlugable() {
        if (safeRoomAddress == address(0) || msg.sender != safeRoomAddress || !plugable) revert Errors.IsNotPlugable();
        _;
    }

    /**
    * @dev Modifier to check if a given trait key is valid.
    */
    modifier isValidTraitKey(bytes32 traitKey) {
        if (traitKey != FrontTraitKey &&
            traitKey != SideTraitKey &&
            traitKey != TopTraitKey) revert Errors.InvalidTraitKeys();
        _;
    }

    /**
    * @dev Modifier to ensure that the signature has not expired.
    */
    modifier isSignatureTimeout(uint256 timestamp) {
        if (block.timestamp >=  timestamp) revert Errors.InvalidTimestamp();
        _;
    }

    /**
    * @notice Returns true if the public sale is currently active.
    * @return True if the public sale is ongoing.
    */
    function isPublicSaleOpen() public view returns (bool) {
        return block.timestamp >= publicSaleStartTimestamp && publicSaleStartTimestamp != 0 && !publicSaleLocked;
    }

    /**
    * @notice Returns true if the whitelist sale is currently active,
    * automatically ends when the public sale starts.
    * @return True if the whitelist sale is ongoing.
    */
    function isWLSaleOpen() public view returns (bool) {
        return !isPublicSaleOpen() && block.timestamp >= WLStartTimestamp && WLStartTimestamp != 0 && !WLLocked;
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           IERC165                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
    * @dev Implements IERC165 interface detection in a way compatible
    * with ERC721A and DynamicTraits.
    */
    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721A, IERC721A, DynamicTraits) returns (bool) {
        return ERC721A.supportsInterface(interfaceId) || DynamicTraits.supportsInterface(interfaceId);
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   Only Owner Functions                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
    * @dev Sets the start timestamp for the whitelist sale.
    * Only callable by the owner.
    */
    function setWLStartTimestamp(uint256 timestamp) external onlyOwner {
        WLStartTimestamp = timestamp;
        emit SetTimestamp(timestamp, WL);
    }

    /**
    * @dev Sets the start timestamp for the public sale.
    * Only callable by the owner.
    */
    function setPublicSaleTimestamp(uint256 timestamp) external onlyOwner {
        publicSaleStartTimestamp = timestamp;
        emit SetTimestamp(timestamp, PublicSale);
    }

    /**
    * @dev Toggles the status of the whitelist sale, allowing manual
    * control for special circumstances.
    * Only callable by the owner.
    */
    function toggleWLSale() external onlyOwner {
        WLLocked = !WLLocked;
        emit ToggleSaleStatus(WL);
    }

    /**
    * @dev Toggles the status of the public sale, allowing manual
    * control for special circumstances.
    * Only callable by the owner.
    */
    function toggleSale() external onlyOwner {
        publicSaleLocked = !publicSaleLocked;
        emit ToggleSaleStatus(PublicSale);
    }

    /**
    * @dev Sets the maximum number of batch mintable tokens.
    * Will be used in following batches.
    * Only callable by the owner, within predefined limits.
    */
    function setBatchMax(uint256 _batchMax) external onlyOwner {
        if (_batchMax < batchMax || _batchMax > PLUGMAN_MAX) revert Errors.InvalidBatchMax();
        batchMax = _batchMax;
        emit SetBatchMax(_batchMax);
    }

    /**
    * @dev Sets the price for the whitelist sale.
    * Will be used in following batches.
    * Only callable by the owner.
    */
    function setWLPrice(uint256 __wlPrice) external onlyOwner {
        wlPrice = __wlPrice;
        emit SetSalePrice(__wlPrice, WL);
    }

    /**
    * @dev Sets the price for the public sale.
    * Will be used in following batches.
    * Only callable by the owner.
    */
    function setPublicPrice(uint256 __publicSalePrice) external onlyOwner {
        publicSalePrice = __publicSalePrice;
        emit SetSalePrice(__publicSalePrice, PublicSale);
    }

    /**
    * @dev Sets the address of the Safe Room Contract.
    * Only callable by the owner.
    */
    function setSafeRoomAddress(address __safeRoomAddress) external onlyOwner {
        if (__safeRoomAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        safeRoomAddress = __safeRoomAddress;
        emit SetSafeRoom(__safeRoomAddress);
    }

    /**
    * @dev Registers an address as a ShuttleMachine Contract.
    * Only callable by the owner.
    */
    function setShuttleMachine(address __shuttleMachineAddress) external onlyOwner {
        if (__shuttleMachineAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        shuttleMachineAddresses[__shuttleMachineAddress] = true;
        emit SetShuttleMachine(__shuttleMachineAddress);
    }

    /**
    * @dev Unregisters an address as a ShuttleMachine Contract.
    * Only callable by the owner.
    */
    function unsetShuttleMachine(address __shuttleMachineAddress) external onlyOwner {
        if (__shuttleMachineAddress == address(0)) revert Errors.ZeroAddressNotAllowed();
        shuttleMachineAddresses[__shuttleMachineAddress] = false;
        emit UnsetShuttleMachine(__shuttleMachineAddress);
    }

    /**
    * @dev Toggles the global configuration allowing or disallowing
    * plugging/unplugging.
    * Only callable by the owner.
    */
    function togglePlugable() external onlyOwner {
        plugable = !plugable;
        emit TogglePlugable();
    }

    /**
    * @dev Sets the base URI for NFT images.
    * Only callable by the owner.
    */
    function setImageBaseUri(string calldata uri) external onlyOwner {
        imageBaseUri = uri;
        emit SetImageBaseUri();
    }

    /**
    * @dev Corrects metadata bugs for a specific token and trait.
    * Only callable by the owner.
    */
    function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) external onlyOwner {
        if (lockOwnerTraitModify) revert Errors.OwnerCannotFixTraits();
        _setTrait(tokenId, traitKey, trait);
        emit TraitUpdated(traitKey, tokenId, trait);
    }

    /**
    * @dev Sets the last token ID for which the reveal is allowed
    * (because Plugman mints in batches)
    * Only callable by the owner.
    */
    function revealUntil(uint256 tokenId) external onlyOwner {
        revealedUntil = tokenId;
        emit RevealUntil(tokenId);
    }

    /**
    * @dev Locks the ability to modify NFT metadata.
    * Only callable by the owner.
    */
    function lockingOwnerTraitModify() external onlyOwner {
        lockOwnerTraitModify = true;
        emit LockOwnerTraitModifying();
    }

    /**
    * @dev Withdraws minting fees collected on the ZetaChain network.
    * Only callable by the owner.
    */
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        (bool success,) = msg.sender.call{value : balance}('');
        if (!success) revert Errors.WithdrawFailed();
        emit Withdraw(msg.sender, balance);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       Mint Functions                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
    * @notice mintOG For the native chain client, used to airdrop NFTs
    * to a specific address without charging the price of the NFT.
    * Emits a Mint event on successful call.
    * Only callable by the owner.
    * @param to Address that should receive the minted NFTs
    * @param _count Number of NFTs to mint
    * @param _nonce Used by the metadata generator to track minting outcomes
    * @param timestamp Deadline for the metadata generator's signature validity
    * @param traitValue Trait information in the metadata
    * @param signature EIP712 signature by the metadata generator for the metadata
    * and related content
    */
    function mintOG(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external nonReentrant onlyOwner isSignatureTimeout(timestamp)
    {
        if (tx.origin != msg.sender) revert Errors.CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert Errors.InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert Errors.AllTokenHasBeenMinted();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, OG, firstTokenId);
        emit Mint(to, _nonce, firstTokenId, OG);
    }

    /**
    * @notice mintWL For the native chain client, used for whitelist sales.
    * Emits a Mint event on successful call.
    * @param _count Number of NFTs to mint
    * @param _nonce Used by the metadata generator to track minting outcomes
    * @param timestamp Deadline for the metadata generator's signature validity
    * @param traitValue Trait information in the metadata
    * @param signature EIP712 signature by the metadata generator for the metadata
    * and related content
    */
    function mintWL(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external payable nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isWLSaleOpen()) revert Errors.WLSaleIsNotActive();
        if (tx.origin != msg.sender) revert Errors.CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert Errors.InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert Errors.AllTokenHasBeenMinted();
        if (msg.value != _count * wlPrice) revert Errors.InvalidCoinValue();
        if (_nonce <= WLNonce[msg.sender]) revert Errors.NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(msg.sender, _count, _nonce, timestamp, traitValue, signature, WL, firstTokenId);
        WLNonce[msg.sender] = _nonce;
        emit Mint(msg.sender, _nonce, firstTokenId, WL);
    }

    /**
    * @notice mintPublicSale For the native chain client, used for public sales.
    * Emits a Mint event on successful call.
    * @param _count Number of NFTs to mint
    * @param _nonce Used by the metadata generator to track minting outcomes
    * @param timestamp Deadline for the metadata generator's signature validity
    * @param traitValue Trait information in the metadata
    * @param signature EIP712 signature by the metadata generator for the metadata
    * and related content
    */
    function mintPublicSale(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
            external payable nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isPublicSaleOpen()) revert Errors.PublicSaleIsNotActive();
        if (tx.origin != msg.sender) revert Errors.CanNotCallMintFromContract();
        if (_count <= 0 || _count > MAX_PER_MINT) revert Errors.InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert Errors.AllTokenHasBeenMinted();
        if (msg.value != _count * publicSalePrice) revert Errors.InvalidCoinValue();
        if (_nonce <= PublicSaleNonce[msg.sender]) revert Errors.NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(msg.sender, _count, _nonce, timestamp, traitValue, signature, PublicSale, firstTokenId);
        PublicSaleNonce[msg.sender] = _nonce;
        emit Mint(msg.sender, _nonce, firstTokenId, PublicSale);
    }

    /**
    * @notice mintWLCrossChain For the ShuttleMachine Contract, used in
    * cross-chain whitelist sales.
    * Emits a Mint event on successful call.
    * @param to Address that should receive the minted NFTs
    * @param _count Number of NFTs to mint
    * @param _nonce Used by the metadata generator to track minting outcomes
    * @param timestamp Deadline for the metadata generator's signature validity
    * @param traitValue Trait information in the metadata
    * @param signature EIP712 signature by the metadata generator for the metadata
    * and related content
    */
    function mintWLCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
    external nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isWLSaleOpen()) revert Errors.WLSaleIsNotActive();
        if (!shuttleMachineAddresses[msg.sender]) revert Errors.CallerNotApproved();
        if (_count <= 0 || _count > MAX_PER_MINT) revert Errors.InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert Errors.AllTokenHasBeenMinted();
        if (_nonce <= WLNonce[to]) revert Errors.NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, WL, firstTokenId);
        WLNonce[to] = _nonce;
        emit Mint(to, _nonce, firstTokenId, WL);
    }

    /**
    * @notice mintPublicCrossChain For the ShuttleMachine Contract, used in
    * cross-chain public sales
    * Emits a Mint event on successful call.
    * @param to Address that should receive the minted NFTs
    * @param _count Number of NFTs to mint
    * @param _nonce Used by the metadata generator to track minting outcomes
    * @param timestamp Deadline for the metadata generator's signature validity
    * @param traitValue Trait information in the metadata
    * @param signature EIP712 signature by the metadata generator for the metadata
    * and related content
    */
    function mintPublicCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] calldata traitValue, bytes calldata signature)
    external nonReentrant isSignatureTimeout(timestamp)
    {
        if (!isPublicSaleOpen()) revert Errors.PublicSaleIsNotActive();
        if (!shuttleMachineAddresses[msg.sender]) revert Errors.CallerNotApproved();
        if (_count <= 0 || _count > MAX_PER_MINT) revert Errors.InvalidMintAmount();
        if (totalSupply() + _count > batchMax) revert Errors.AllTokenHasBeenMinted();
        if (_nonce <= PublicSaleNonce[to]) revert Errors.NonceIsTooOld();
        uint256 firstTokenId = _nextTokenId();
        _verifyAndMint(to, _count, _nonce, timestamp, traitValue, signature, PublicSale, firstTokenId);
        PublicSaleNonce[to] = _nonce;
        emit Mint(to, _nonce, firstTokenId, PublicSale);
    }

    /**
    * @dev _verifyAndMint Verifies metadata generator's signature and mints new tokens.
    */
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
        if (!mintVerify(signature, to, _count, _nonce, _timestamp, traitValue, mintType)) revert Errors.InvalidSignature();
        // ERC721A mint
        _safeMint(to, _count);
        // set dynamic traits
        _setTraitsAfterMint(_count, firstTokenId, traitValue);
    }

    /**
    * @dev _setTraitsAfterMint Sets dynamic on-chain traits after minting.
    */
    function _setTraitsAfterMint(uint256 _count, uint256 firstTokenId, bytes32[] calldata traitValue) internal {
        if (_count * 6 != traitValue.length) revert Errors.InvalidTraitValue();
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

    /**
    * @dev mintVerify Calls the MintVerifier contract to verify the EIP712 signature.
    */
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


    /**
    * @dev Provides a token URI with metadata for a given token ID. Overrides ERC721A and IERC721A implementations.
    * @param tokenId The ID of the token for which metadata is requested.
    * @return A string representing a URI that points to the token's metadata.
    * The metadata includes the token's name, description, image URL, and an array of attributes.
    * Depending on whether the token ID is less than or equal to `revealedUntil`, it will show either all attributes (static and dynamic)
    * or only static attributes. The metadata is encoded in Base64 and prefixed with the appropriate data URI scheme.
    */
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

    /**
    * @dev Generates an array of static attribute strings for a given token ID.
    */
    function _staticAttributes(uint256 tokenId) internal view returns (string[] memory) {

        return Solarray.strings(
            Metadata.attribute({traitType: "form", value: revealedUntil > tokenId ? "Plugman" : "Mystery Box", displayType: DisplayType.String})
        );
    }


    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ERC-7496 Functions                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Overrides OnchainTraits' setTrait to include plug/unplug feature.
     * @param tokenId ID of the token to update traits.
     * @param traitKey Key for the trait to update.
     * @param trait New trait value to set; only empty trait values can be updated.
     */

    function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) public override isPlugable isValidTraitKey(traitKey) {
        bytes32 traitValue = getTraitValue(tokenId, traitKey);
        // can set empty trait only
        if (traitValue != bytes32(0)) revert Errors.TraitValueNotEmpty();
        _setTrait(tokenId, traitKey, trait);
        emit TraitUpdated(traitKey, tokenId, trait);
    }

    /**
     * @dev Deletes a specific trait for a token.
     * @param traitKey Key of the trait to delete.
     * @param tokenId ID of the token.
     */
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
