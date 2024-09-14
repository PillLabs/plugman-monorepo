// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.8;

type DisplayType is uint8;
type Editors is uint8;
type StoredTraitLabel is address;

struct FullTraitValue {
    bytes32 traitValue;
    string fullTraitValue;
}

struct TokenOwnership {
    address addr;
    uint64 startTimestamp;
    bool burned;
    uint24 extraData;
}

struct TraitLabel {
    string fullTraitKey;
    string traitLabel;
    string[] acceptableValues;
    FullTraitValue[] fullTraitValues;
    DisplayType displayType;
    Editors editors;
    bool required;
}

struct TraitLabelStorage {
    Editors allowedEditors;
    bool required;
    bool valuesRequireValidation;
    StoredTraitLabel storedLabel;
}

error AllTokenHasBeenMinted();
error AlreadyInitialized();
error ApprovalCallerNotOwnerNorApproved();
error ApprovalQueryForNonexistentToken();
error BalanceQueryForZeroAddress();
error CallerNotApproved();
error CanNotCallMintFromContract();
error InsufficientPrivilege();
error InvalidBatchMax();
error InvalidCoinValue();
error InvalidMintAmount();
error InvalidQueryRange();
error InvalidSignature();
error InvalidTimestamp();
error InvalidTraitKeys();
error InvalidTraitValue();
error IsNotPlugable();
error MintERC2309QuantityExceedsLimit();
error MintToZeroAddress();
error MintZeroQuantity();
error NewOwnerIsZeroAddress();
error NoHandoverRequest();
error NonceIsTooOld();
error NotCompatibleWithSpotMints();
error OwnerCannotFixTraits();
error OwnerQueryForNonexistentToken();
error OwnershipNotInitializedForExtraData();
error PublicSaleIsNotActive();
error ReentrancyGuardReentrantCall();
error SequentialMintExceedsLimit();
error SequentialUpToTooSmall();
error SpotMintTokenIdTooSmall();
error TokenAlreadyExists();
error TraitDoesNotExist(bytes32 traitKey);
error TraitValueNotEmpty();
error TraitValueUnchanged();
error TransferCallerNotOwnerNorApproved();
error TransferFromIncorrectOwner();
error TransferToNonERC721ReceiverImplementer();
error TransferToZeroAddress();
error URIQueryForNonexistentToken();
error Unauthorized();
error WLSaleIsNotActive();
error WithdrawFailed();

interface PlugmanInterface {

    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);
    event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to);
    event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType);
    event OwnershipHandoverCanceled(address indexed pendingOwner);
    event OwnershipHandoverRequested(address indexed pendingOwner);
    event OwnershipTransferred(address indexed oldOwner, address indexed newOwner);
    event TraitMetadataURIUpdated();
    event TraitUpdated(bytes32 indexed traitKey, uint256 tokenId, bytes32 traitValue);
    event TraitUpdatedList(bytes32 indexed traitKey, uint256[] tokenIds);
    event TraitUpdatedListUniformValue(bytes32 indexed traitKey, uint256[] tokenIds, bytes32 traitValue);
    event TraitUpdatedRange(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId);
    event TraitUpdatedRangeUniformValue(
        bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId, bytes32 traitValue
    );
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);

    function PublicSaleNonce(address) external view returns (uint256);
    function WLLocked() external view returns (bool);
    function WLNonce(address) external view returns (uint256);
    function WLStartTimestamp() external view returns (uint256);
    function approve(address to, uint256 tokenId) external payable;
    function balanceOf(address owner) external view returns (uint256);
    function batchMax() external view returns (uint256);
    function cancelOwnershipHandover() external payable;
    function completeOwnershipHandover(address pendingOwner) external payable;
    function deleteTrait(bytes32 traitKey, uint256 tokenId) external;
    function explicitOwnershipOf(uint256 tokenId) external view returns (TokenOwnership memory ownership);
    function explicitOwnershipsOf(uint256[] memory tokenIds) external view returns (TokenOwnership[] memory);
    function getApproved(uint256 tokenId) external view returns (address);
    function getCustomEditorAt(uint256 index) external view returns (address);
    function getCustomEditors() external view returns (address[] memory);
    function getCustomEditorsLength() external view returns (uint256);
    function getTraitMetadataURI() external view returns (string memory);
    function getTraitValue(uint256 tokenId, bytes32 traitKey) external view returns (bytes32 traitValue);
    function getTraitValues(uint256 tokenId, bytes32[] memory traitKeys)
    external
    view
    returns (bytes32[] memory traitValues);
    function isApprovedForAll(address owner, address operator) external view returns (bool);
    function isCustomEditor(address editor) external view returns (bool);
    function isPublicSaleOpen() external view returns (bool);
    function isWLSaleOpen() external view returns (bool);
    function lockingOwnerTraitModify() external;
    function mintOG(
        address to,
        uint256 _count,
        uint256 _nonce,
        uint256 timestamp,
        bytes32[] memory traitValue,
        bytes memory signature
    ) external;
    function mintPublicCrossChain(
        address to,
        uint256 _count,
        uint256 _nonce,
        uint256 timestamp,
        bytes32[] memory traitValue,
        bytes memory signature
    ) external;
    function mintPublicSale(
        uint256 _count,
        uint256 _nonce,
        uint256 timestamp,
        bytes32[] memory traitValue,
        bytes memory signature
    ) external payable;
    function mintVerify(
        bytes memory _signature,
        address _mintTo,
        uint256 _amount,
        uint256 _nonce,
        uint256 _timestamp,
        bytes32[] memory traitValue,
        uint8 _mintType
    ) external view returns (bool);
    function mintWL(
        uint256 _count,
        uint256 _nonce,
        uint256 timestamp,
        bytes32[] memory traitValue,
        bytes memory signature
    ) external payable;
    function mintWLCrossChain(
        address to,
        uint256 _count,
        uint256 _nonce,
        uint256 timestamp,
        bytes32[] memory traitValue,
        bytes memory signature
    ) external payable;
    function name() external view returns (string memory);
    function owner() external view returns (address result);
    function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) external;
    function ownerOf(uint256 tokenId) external view returns (address);
    function ownershipHandoverExpiresAt(address pendingOwner) external view returns (uint256 result);
    function plugable() external view returns (bool);
    function publicSaleLocked() external view returns (bool);
    function publicSaleStartTimestamp() external view returns (uint256);
    function renounceOwnership() external payable;
    function requestOwnershipHandover() external payable;
    function revealUntil(uint256 tokenId) external;
    function revealedUntil() external view returns (uint256);
    function safeTransferFrom(address from, address to, uint256 tokenId) external payable;
    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) external payable;
    function setApprovalForAll(address operator, bool approved) external;
    function setBatchMax(uint256 _batchMax) external;
    function setImageBaseUri(string memory uri) external;
    function setPublicPrice(uint256 __publicSalePrice) external;
    function setPublicSaleTimestamp(uint256 timestamp) external;
    function setSafeRoomAddress(address __safeRoomAddress) external;
    function setShuttleMachine(address __shuttleMachineAddress) external;
    function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) external;
    function setTraitLabel(bytes32 traitKey, TraitLabel memory _traitLabel) external;
    function setWLPrice(uint256 __wlPrice) external;
    function setWLStartTimestamp(uint256 timestamp) external;
    function shuttleMachineAddresses(address) external view returns (bool);
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
    function symbol() external view returns (string memory);
    function togglePlugable() external;
    function toggleSale() external;
    function toggleWLSale() external;
    function tokenURI(uint256 tokenId) external view returns (string memory);
    function tokensOfOwner(address owner) external view returns (uint256[] memory);
    function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) external view returns (uint256[] memory);
    function totalSupply() external view returns (uint256 result);
    function traitLabelStorage(bytes32 traitKey) external view returns (TraitLabelStorage memory);
    function transferFrom(address from, address to, uint256 tokenId) external payable;
    function transferOwnership(address newOwner) external payable;
    function unsetShuttleMachine(address __shuttleMachineAddress) external;
    function updateCustomEditor(address editor, bool insert) external;
    function withdraw() external;
}
