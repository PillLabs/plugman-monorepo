// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plugman

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FullTraitValue is an auto generated low-level Go binding around an user-defined struct.
type FullTraitValue struct {
	TraitValue     [32]byte
	FullTraitValue string
}

// IERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// TraitLabel is an auto generated low-level Go binding around an user-defined struct.
type TraitLabel struct {
	FullTraitKey     string
	TraitLabel       string
	AcceptableValues []string
	FullTraitValues  []FullTraitValue
	DisplayType      uint8
	Editors          uint8
	Required         bool
}

// TraitLabelStorage is an auto generated low-level Go binding around an user-defined struct.
type TraitLabelStorage struct {
	AllowedEditors          uint8
	Required                bool
	ValuesRequireValidation bool
	StoredLabel             common.Address
}

// PlugmanMetaData contains all meta data concerning the Plugman contract.
var PlugmanMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"__name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"__symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"__mintVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"__batchMax\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PublicSaleNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WLLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WLNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WLStartTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchMax\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deleteTrait\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explicitOwnershipOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ownership\",\"type\":\"tuple\",\"internalType\":\"structIERC721A.TokenOwnership\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"burned\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"extraData\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"explicitOwnershipsOf\",\"inputs\":[{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIERC721A.TokenOwnership[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"burned\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"extraData\",\"type\":\"uint24\",\"internalType\":\"uint24\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCustomEditorAt\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCustomEditors\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCustomEditorsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTraitMetadataURI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTraitValue\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"traitValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTraitValues\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"traitValues\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCustomEditor\",\"inputs\":[{\"name\":\"editor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPublicSaleOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWLSaleOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockingOwnerTraitModify\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintOG\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintPublicCrossChain\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintPublicSale\",\"inputs\":[{\"name\":\"_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mintVerify\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_mintTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"_mintType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintWL\",\"inputs\":[{\"name\":\"_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mintWLCrossChain\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerFixTraits\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trait\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"plugable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicSaleLocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicSaleStartTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revealUntil\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revealedUntil\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBatchMax\",\"inputs\":[{\"name\":\"_batchMax\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setImageBaseUri\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublicPrice\",\"inputs\":[{\"name\":\"__publicSalePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublicSaleTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSafeRoomAddress\",\"inputs\":[{\"name\":\"__safeRoomAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShuttleMachine\",\"inputs\":[{\"name\":\"__shuttleMachineAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTrait\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trait\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTraitLabel\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_traitLabel\",\"type\":\"tuple\",\"internalType\":\"structTraitLabel\",\"components\":[{\"name\":\"fullTraitKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"traitLabel\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"acceptableValues\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"fullTraitValues\",\"type\":\"tuple[]\",\"internalType\":\"structFullTraitValue[]\",\"components\":[{\"name\":\"traitValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fullTraitValue\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"displayType\",\"type\":\"uint8\",\"internalType\":\"enumDisplayType\"},{\"name\":\"editors\",\"type\":\"uint8\",\"internalType\":\"Editors\"},{\"name\":\"required\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWLPrice\",\"inputs\":[{\"name\":\"__wlPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWLStartTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shuttleMachineAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"togglePlugable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toggleSale\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toggleWLSale\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensOfOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensOfOwnerIn\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stop\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"traitLabelStorage\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTraitLabelStorage\",\"components\":[{\"name\":\"allowedEditors\",\"type\":\"uint8\",\"internalType\":\"Editors\"},{\"name\":\"required\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"valuesRequireValidation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"storedLabel\",\"type\":\"address\",\"internalType\":\"StoredTraitLabel\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unsetShuttleMachine\",\"inputs\":[{\"name\":\"__shuttleMachineAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCustomEditor\",\"inputs\":[{\"name\":\"editor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"insert\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsecutiveTransfer\",\"inputs\":[{\"name\":\"fromTokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"firstTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitMetadataURIUpdated\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitUpdated\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitUpdatedList\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitUpdatedListUniformValue\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"traitValue\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitUpdatedRange\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TraitUpdatedRangeUniformValue\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"traitValue\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllTokenHasBeenMinted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalQueryForNonexistentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BalanceQueryForZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CanNotCallMintFromContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPrivilege\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBatchMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCoinValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMintAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQueryRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTraitKeys\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTraitValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsNotPlugable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintERC2309QuantityExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintZeroQuantity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonceIsTooOld\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCompatibleWithSpotMints\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotFixTraits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerQueryForNonexistentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipNotInitializedForExtraData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicSaleIsNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequentialMintExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequentialUpToTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SpotMintTokenIdTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TraitDoesNotExist\",\"inputs\":[{\"name\":\"traitKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TraitValueNotEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TraitValueUnchanged\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferCallerNotOwnerNorApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFromIncorrectOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToNonERC721ReceiverImplementer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"URIQueryForNonexistentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WLSaleIsNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawFailed\",\"inputs\":[]}]",
}

// PlugmanABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugmanMetaData.ABI instead.
var PlugmanABI = PlugmanMetaData.ABI

// Plugman is an auto generated Go binding around an Ethereum contract.
type Plugman struct {
	PlugmanCaller     // Read-only binding to the contract
	PlugmanTransactor // Write-only binding to the contract
	PlugmanFilterer   // Log filterer for contract events
}

// PlugmanCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugmanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugmanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugmanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugmanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugmanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugmanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugmanSession struct {
	Contract     *Plugman          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugmanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugmanCallerSession struct {
	Contract *PlugmanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PlugmanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugmanTransactorSession struct {
	Contract     *PlugmanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PlugmanRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugmanRaw struct {
	Contract *Plugman // Generic contract binding to access the raw methods on
}

// PlugmanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugmanCallerRaw struct {
	Contract *PlugmanCaller // Generic read-only contract binding to access the raw methods on
}

// PlugmanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugmanTransactorRaw struct {
	Contract *PlugmanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugman creates a new instance of Plugman, bound to a specific deployed contract.
func NewPlugman(address common.Address, backend bind.ContractBackend) (*Plugman, error) {
	contract, err := bindPlugman(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Plugman{PlugmanCaller: PlugmanCaller{contract: contract}, PlugmanTransactor: PlugmanTransactor{contract: contract}, PlugmanFilterer: PlugmanFilterer{contract: contract}}, nil
}

// NewPlugmanCaller creates a new read-only instance of Plugman, bound to a specific deployed contract.
func NewPlugmanCaller(address common.Address, caller bind.ContractCaller) (*PlugmanCaller, error) {
	contract, err := bindPlugman(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugmanCaller{contract: contract}, nil
}

// NewPlugmanTransactor creates a new write-only instance of Plugman, bound to a specific deployed contract.
func NewPlugmanTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugmanTransactor, error) {
	contract, err := bindPlugman(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugmanTransactor{contract: contract}, nil
}

// NewPlugmanFilterer creates a new log filterer instance of Plugman, bound to a specific deployed contract.
func NewPlugmanFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugmanFilterer, error) {
	contract, err := bindPlugman(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugmanFilterer{contract: contract}, nil
}

// bindPlugman binds a generic wrapper to an already deployed contract.
func bindPlugman(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugmanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plugman *PlugmanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Plugman.Contract.PlugmanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plugman *PlugmanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.Contract.PlugmanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plugman *PlugmanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plugman.Contract.PlugmanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Plugman *PlugmanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Plugman.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Plugman *PlugmanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Plugman *PlugmanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Plugman.Contract.contract.Transact(opts, method, params...)
}

// PublicSaleNonce is a free data retrieval call binding the contract method 0x8d4ac162.
//
// Solidity: function PublicSaleNonce(address ) view returns(uint256)
func (_Plugman *PlugmanCaller) PublicSaleNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "PublicSaleNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleNonce is a free data retrieval call binding the contract method 0x8d4ac162.
//
// Solidity: function PublicSaleNonce(address ) view returns(uint256)
func (_Plugman *PlugmanSession) PublicSaleNonce(arg0 common.Address) (*big.Int, error) {
	return _Plugman.Contract.PublicSaleNonce(&_Plugman.CallOpts, arg0)
}

// PublicSaleNonce is a free data retrieval call binding the contract method 0x8d4ac162.
//
// Solidity: function PublicSaleNonce(address ) view returns(uint256)
func (_Plugman *PlugmanCallerSession) PublicSaleNonce(arg0 common.Address) (*big.Int, error) {
	return _Plugman.Contract.PublicSaleNonce(&_Plugman.CallOpts, arg0)
}

// WLLocked is a free data retrieval call binding the contract method 0x3d959495.
//
// Solidity: function WLLocked() view returns(bool)
func (_Plugman *PlugmanCaller) WLLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "WLLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WLLocked is a free data retrieval call binding the contract method 0x3d959495.
//
// Solidity: function WLLocked() view returns(bool)
func (_Plugman *PlugmanSession) WLLocked() (bool, error) {
	return _Plugman.Contract.WLLocked(&_Plugman.CallOpts)
}

// WLLocked is a free data retrieval call binding the contract method 0x3d959495.
//
// Solidity: function WLLocked() view returns(bool)
func (_Plugman *PlugmanCallerSession) WLLocked() (bool, error) {
	return _Plugman.Contract.WLLocked(&_Plugman.CallOpts)
}

// WLNonce is a free data retrieval call binding the contract method 0x129d5b02.
//
// Solidity: function WLNonce(address ) view returns(uint256)
func (_Plugman *PlugmanCaller) WLNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "WLNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WLNonce is a free data retrieval call binding the contract method 0x129d5b02.
//
// Solidity: function WLNonce(address ) view returns(uint256)
func (_Plugman *PlugmanSession) WLNonce(arg0 common.Address) (*big.Int, error) {
	return _Plugman.Contract.WLNonce(&_Plugman.CallOpts, arg0)
}

// WLNonce is a free data retrieval call binding the contract method 0x129d5b02.
//
// Solidity: function WLNonce(address ) view returns(uint256)
func (_Plugman *PlugmanCallerSession) WLNonce(arg0 common.Address) (*big.Int, error) {
	return _Plugman.Contract.WLNonce(&_Plugman.CallOpts, arg0)
}

// WLStartTimestamp is a free data retrieval call binding the contract method 0x4d84a5ac.
//
// Solidity: function WLStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanCaller) WLStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "WLStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WLStartTimestamp is a free data retrieval call binding the contract method 0x4d84a5ac.
//
// Solidity: function WLStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanSession) WLStartTimestamp() (*big.Int, error) {
	return _Plugman.Contract.WLStartTimestamp(&_Plugman.CallOpts)
}

// WLStartTimestamp is a free data retrieval call binding the contract method 0x4d84a5ac.
//
// Solidity: function WLStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanCallerSession) WLStartTimestamp() (*big.Int, error) {
	return _Plugman.Contract.WLStartTimestamp(&_Plugman.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plugman *PlugmanCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plugman *PlugmanSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Plugman.Contract.BalanceOf(&_Plugman.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Plugman *PlugmanCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Plugman.Contract.BalanceOf(&_Plugman.CallOpts, owner)
}

// BatchMax is a free data retrieval call binding the contract method 0x589e47dc.
//
// Solidity: function batchMax() view returns(uint256)
func (_Plugman *PlugmanCaller) BatchMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "batchMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchMax is a free data retrieval call binding the contract method 0x589e47dc.
//
// Solidity: function batchMax() view returns(uint256)
func (_Plugman *PlugmanSession) BatchMax() (*big.Int, error) {
	return _Plugman.Contract.BatchMax(&_Plugman.CallOpts)
}

// BatchMax is a free data retrieval call binding the contract method 0x589e47dc.
//
// Solidity: function batchMax() view returns(uint256)
func (_Plugman *PlugmanCallerSession) BatchMax() (*big.Int, error) {
	return _Plugman.Contract.BatchMax(&_Plugman.CallOpts)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Plugman *PlugmanCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721ATokenOwnership)).(*IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Plugman *PlugmanSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Plugman.Contract.ExplicitOwnershipOf(&_Plugman.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Plugman *PlugmanCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Plugman.Contract.ExplicitOwnershipOf(&_Plugman.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Plugman *PlugmanCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721ATokenOwnership)).(*[]IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Plugman *PlugmanSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Plugman.Contract.ExplicitOwnershipsOf(&_Plugman.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Plugman *PlugmanCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Plugman.Contract.ExplicitOwnershipsOf(&_Plugman.CallOpts, tokenIds)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Plugman.Contract.GetApproved(&_Plugman.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Plugman.Contract.GetApproved(&_Plugman.CallOpts, tokenId)
}

// GetCustomEditorAt is a free data retrieval call binding the contract method 0x8e31261a.
//
// Solidity: function getCustomEditorAt(uint256 index) view returns(address)
func (_Plugman *PlugmanCaller) GetCustomEditorAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getCustomEditorAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCustomEditorAt is a free data retrieval call binding the contract method 0x8e31261a.
//
// Solidity: function getCustomEditorAt(uint256 index) view returns(address)
func (_Plugman *PlugmanSession) GetCustomEditorAt(index *big.Int) (common.Address, error) {
	return _Plugman.Contract.GetCustomEditorAt(&_Plugman.CallOpts, index)
}

// GetCustomEditorAt is a free data retrieval call binding the contract method 0x8e31261a.
//
// Solidity: function getCustomEditorAt(uint256 index) view returns(address)
func (_Plugman *PlugmanCallerSession) GetCustomEditorAt(index *big.Int) (common.Address, error) {
	return _Plugman.Contract.GetCustomEditorAt(&_Plugman.CallOpts, index)
}

// GetCustomEditors is a free data retrieval call binding the contract method 0xcc84c1b7.
//
// Solidity: function getCustomEditors() view returns(address[])
func (_Plugman *PlugmanCaller) GetCustomEditors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getCustomEditors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCustomEditors is a free data retrieval call binding the contract method 0xcc84c1b7.
//
// Solidity: function getCustomEditors() view returns(address[])
func (_Plugman *PlugmanSession) GetCustomEditors() ([]common.Address, error) {
	return _Plugman.Contract.GetCustomEditors(&_Plugman.CallOpts)
}

// GetCustomEditors is a free data retrieval call binding the contract method 0xcc84c1b7.
//
// Solidity: function getCustomEditors() view returns(address[])
func (_Plugman *PlugmanCallerSession) GetCustomEditors() ([]common.Address, error) {
	return _Plugman.Contract.GetCustomEditors(&_Plugman.CallOpts)
}

// GetCustomEditorsLength is a free data retrieval call binding the contract method 0xd7682b5a.
//
// Solidity: function getCustomEditorsLength() view returns(uint256)
func (_Plugman *PlugmanCaller) GetCustomEditorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getCustomEditorsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCustomEditorsLength is a free data retrieval call binding the contract method 0xd7682b5a.
//
// Solidity: function getCustomEditorsLength() view returns(uint256)
func (_Plugman *PlugmanSession) GetCustomEditorsLength() (*big.Int, error) {
	return _Plugman.Contract.GetCustomEditorsLength(&_Plugman.CallOpts)
}

// GetCustomEditorsLength is a free data retrieval call binding the contract method 0xd7682b5a.
//
// Solidity: function getCustomEditorsLength() view returns(uint256)
func (_Plugman *PlugmanCallerSession) GetCustomEditorsLength() (*big.Int, error) {
	return _Plugman.Contract.GetCustomEditorsLength(&_Plugman.CallOpts)
}

// GetTraitMetadataURI is a free data retrieval call binding the contract method 0xde475bf9.
//
// Solidity: function getTraitMetadataURI() view returns(string)
func (_Plugman *PlugmanCaller) GetTraitMetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getTraitMetadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTraitMetadataURI is a free data retrieval call binding the contract method 0xde475bf9.
//
// Solidity: function getTraitMetadataURI() view returns(string)
func (_Plugman *PlugmanSession) GetTraitMetadataURI() (string, error) {
	return _Plugman.Contract.GetTraitMetadataURI(&_Plugman.CallOpts)
}

// GetTraitMetadataURI is a free data retrieval call binding the contract method 0xde475bf9.
//
// Solidity: function getTraitMetadataURI() view returns(string)
func (_Plugman *PlugmanCallerSession) GetTraitMetadataURI() (string, error) {
	return _Plugman.Contract.GetTraitMetadataURI(&_Plugman.CallOpts)
}

// GetTraitValue is a free data retrieval call binding the contract method 0xa28eec87.
//
// Solidity: function getTraitValue(uint256 tokenId, bytes32 traitKey) view returns(bytes32 traitValue)
func (_Plugman *PlugmanCaller) GetTraitValue(opts *bind.CallOpts, tokenId *big.Int, traitKey [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getTraitValue", tokenId, traitKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTraitValue is a free data retrieval call binding the contract method 0xa28eec87.
//
// Solidity: function getTraitValue(uint256 tokenId, bytes32 traitKey) view returns(bytes32 traitValue)
func (_Plugman *PlugmanSession) GetTraitValue(tokenId *big.Int, traitKey [32]byte) ([32]byte, error) {
	return _Plugman.Contract.GetTraitValue(&_Plugman.CallOpts, tokenId, traitKey)
}

// GetTraitValue is a free data retrieval call binding the contract method 0xa28eec87.
//
// Solidity: function getTraitValue(uint256 tokenId, bytes32 traitKey) view returns(bytes32 traitValue)
func (_Plugman *PlugmanCallerSession) GetTraitValue(tokenId *big.Int, traitKey [32]byte) ([32]byte, error) {
	return _Plugman.Contract.GetTraitValue(&_Plugman.CallOpts, tokenId, traitKey)
}

// GetTraitValues is a free data retrieval call binding the contract method 0xf80ecba3.
//
// Solidity: function getTraitValues(uint256 tokenId, bytes32[] traitKeys) view returns(bytes32[] traitValues)
func (_Plugman *PlugmanCaller) GetTraitValues(opts *bind.CallOpts, tokenId *big.Int, traitKeys [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "getTraitValues", tokenId, traitKeys)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetTraitValues is a free data retrieval call binding the contract method 0xf80ecba3.
//
// Solidity: function getTraitValues(uint256 tokenId, bytes32[] traitKeys) view returns(bytes32[] traitValues)
func (_Plugman *PlugmanSession) GetTraitValues(tokenId *big.Int, traitKeys [][32]byte) ([][32]byte, error) {
	return _Plugman.Contract.GetTraitValues(&_Plugman.CallOpts, tokenId, traitKeys)
}

// GetTraitValues is a free data retrieval call binding the contract method 0xf80ecba3.
//
// Solidity: function getTraitValues(uint256 tokenId, bytes32[] traitKeys) view returns(bytes32[] traitValues)
func (_Plugman *PlugmanCallerSession) GetTraitValues(tokenId *big.Int, traitKeys [][32]byte) ([][32]byte, error) {
	return _Plugman.Contract.GetTraitValues(&_Plugman.CallOpts, tokenId, traitKeys)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plugman *PlugmanCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plugman *PlugmanSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Plugman.Contract.IsApprovedForAll(&_Plugman.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Plugman *PlugmanCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Plugman.Contract.IsApprovedForAll(&_Plugman.CallOpts, owner, operator)
}

// IsCustomEditor is a free data retrieval call binding the contract method 0xef9c9fe0.
//
// Solidity: function isCustomEditor(address editor) view returns(bool)
func (_Plugman *PlugmanCaller) IsCustomEditor(opts *bind.CallOpts, editor common.Address) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "isCustomEditor", editor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCustomEditor is a free data retrieval call binding the contract method 0xef9c9fe0.
//
// Solidity: function isCustomEditor(address editor) view returns(bool)
func (_Plugman *PlugmanSession) IsCustomEditor(editor common.Address) (bool, error) {
	return _Plugman.Contract.IsCustomEditor(&_Plugman.CallOpts, editor)
}

// IsCustomEditor is a free data retrieval call binding the contract method 0xef9c9fe0.
//
// Solidity: function isCustomEditor(address editor) view returns(bool)
func (_Plugman *PlugmanCallerSession) IsCustomEditor(editor common.Address) (bool, error) {
	return _Plugman.Contract.IsCustomEditor(&_Plugman.CallOpts, editor)
}

// IsPublicSaleOpen is a free data retrieval call binding the contract method 0x1a6949e3.
//
// Solidity: function isPublicSaleOpen() view returns(bool)
func (_Plugman *PlugmanCaller) IsPublicSaleOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "isPublicSaleOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSaleOpen is a free data retrieval call binding the contract method 0x1a6949e3.
//
// Solidity: function isPublicSaleOpen() view returns(bool)
func (_Plugman *PlugmanSession) IsPublicSaleOpen() (bool, error) {
	return _Plugman.Contract.IsPublicSaleOpen(&_Plugman.CallOpts)
}

// IsPublicSaleOpen is a free data retrieval call binding the contract method 0x1a6949e3.
//
// Solidity: function isPublicSaleOpen() view returns(bool)
func (_Plugman *PlugmanCallerSession) IsPublicSaleOpen() (bool, error) {
	return _Plugman.Contract.IsPublicSaleOpen(&_Plugman.CallOpts)
}

// IsWLSaleOpen is a free data retrieval call binding the contract method 0x116cb034.
//
// Solidity: function isWLSaleOpen() view returns(bool)
func (_Plugman *PlugmanCaller) IsWLSaleOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "isWLSaleOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWLSaleOpen is a free data retrieval call binding the contract method 0x116cb034.
//
// Solidity: function isWLSaleOpen() view returns(bool)
func (_Plugman *PlugmanSession) IsWLSaleOpen() (bool, error) {
	return _Plugman.Contract.IsWLSaleOpen(&_Plugman.CallOpts)
}

// IsWLSaleOpen is a free data retrieval call binding the contract method 0x116cb034.
//
// Solidity: function isWLSaleOpen() view returns(bool)
func (_Plugman *PlugmanCallerSession) IsWLSaleOpen() (bool, error) {
	return _Plugman.Contract.IsWLSaleOpen(&_Plugman.CallOpts)
}

// MintVerify is a free data retrieval call binding the contract method 0xb3f429c3.
//
// Solidity: function mintVerify(bytes _signature, address _mintTo, uint256 _amount, uint256 _nonce, uint256 _timestamp, bytes32[] traitValue, uint8 _mintType) view returns(bool)
func (_Plugman *PlugmanCaller) MintVerify(opts *bind.CallOpts, _signature []byte, _mintTo common.Address, _amount *big.Int, _nonce *big.Int, _timestamp *big.Int, traitValue [][32]byte, _mintType uint8) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "mintVerify", _signature, _mintTo, _amount, _nonce, _timestamp, traitValue, _mintType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintVerify is a free data retrieval call binding the contract method 0xb3f429c3.
//
// Solidity: function mintVerify(bytes _signature, address _mintTo, uint256 _amount, uint256 _nonce, uint256 _timestamp, bytes32[] traitValue, uint8 _mintType) view returns(bool)
func (_Plugman *PlugmanSession) MintVerify(_signature []byte, _mintTo common.Address, _amount *big.Int, _nonce *big.Int, _timestamp *big.Int, traitValue [][32]byte, _mintType uint8) (bool, error) {
	return _Plugman.Contract.MintVerify(&_Plugman.CallOpts, _signature, _mintTo, _amount, _nonce, _timestamp, traitValue, _mintType)
}

// MintVerify is a free data retrieval call binding the contract method 0xb3f429c3.
//
// Solidity: function mintVerify(bytes _signature, address _mintTo, uint256 _amount, uint256 _nonce, uint256 _timestamp, bytes32[] traitValue, uint8 _mintType) view returns(bool)
func (_Plugman *PlugmanCallerSession) MintVerify(_signature []byte, _mintTo common.Address, _amount *big.Int, _nonce *big.Int, _timestamp *big.Int, traitValue [][32]byte, _mintType uint8) (bool, error) {
	return _Plugman.Contract.MintVerify(&_Plugman.CallOpts, _signature, _mintTo, _amount, _nonce, _timestamp, traitValue, _mintType)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plugman *PlugmanCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plugman *PlugmanSession) Name() (string, error) {
	return _Plugman.Contract.Name(&_Plugman.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Plugman *PlugmanCallerSession) Name() (string, error) {
	return _Plugman.Contract.Name(&_Plugman.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Plugman *PlugmanCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Plugman *PlugmanSession) Owner() (common.Address, error) {
	return _Plugman.Contract.Owner(&_Plugman.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Plugman *PlugmanCallerSession) Owner() (common.Address, error) {
	return _Plugman.Contract.Owner(&_Plugman.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Plugman.Contract.OwnerOf(&_Plugman.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Plugman *PlugmanCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Plugman.Contract.OwnerOf(&_Plugman.CallOpts, tokenId)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Plugman *PlugmanCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Plugman *PlugmanSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Plugman.Contract.OwnershipHandoverExpiresAt(&_Plugman.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Plugman *PlugmanCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Plugman.Contract.OwnershipHandoverExpiresAt(&_Plugman.CallOpts, pendingOwner)
}

// Plugable is a free data retrieval call binding the contract method 0xff5757de.
//
// Solidity: function plugable() view returns(bool)
func (_Plugman *PlugmanCaller) Plugable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "plugable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Plugable is a free data retrieval call binding the contract method 0xff5757de.
//
// Solidity: function plugable() view returns(bool)
func (_Plugman *PlugmanSession) Plugable() (bool, error) {
	return _Plugman.Contract.Plugable(&_Plugman.CallOpts)
}

// Plugable is a free data retrieval call binding the contract method 0xff5757de.
//
// Solidity: function plugable() view returns(bool)
func (_Plugman *PlugmanCallerSession) Plugable() (bool, error) {
	return _Plugman.Contract.Plugable(&_Plugman.CallOpts)
}

// PublicSaleLocked is a free data retrieval call binding the contract method 0xdfd0f395.
//
// Solidity: function publicSaleLocked() view returns(bool)
func (_Plugman *PlugmanCaller) PublicSaleLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "publicSaleLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicSaleLocked is a free data retrieval call binding the contract method 0xdfd0f395.
//
// Solidity: function publicSaleLocked() view returns(bool)
func (_Plugman *PlugmanSession) PublicSaleLocked() (bool, error) {
	return _Plugman.Contract.PublicSaleLocked(&_Plugman.CallOpts)
}

// PublicSaleLocked is a free data retrieval call binding the contract method 0xdfd0f395.
//
// Solidity: function publicSaleLocked() view returns(bool)
func (_Plugman *PlugmanCallerSession) PublicSaleLocked() (bool, error) {
	return _Plugman.Contract.PublicSaleLocked(&_Plugman.CallOpts)
}

// PublicSaleStartTimestamp is a free data retrieval call binding the contract method 0xd7822c99.
//
// Solidity: function publicSaleStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanCaller) PublicSaleStartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "publicSaleStartTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSaleStartTimestamp is a free data retrieval call binding the contract method 0xd7822c99.
//
// Solidity: function publicSaleStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanSession) PublicSaleStartTimestamp() (*big.Int, error) {
	return _Plugman.Contract.PublicSaleStartTimestamp(&_Plugman.CallOpts)
}

// PublicSaleStartTimestamp is a free data retrieval call binding the contract method 0xd7822c99.
//
// Solidity: function publicSaleStartTimestamp() view returns(uint256)
func (_Plugman *PlugmanCallerSession) PublicSaleStartTimestamp() (*big.Int, error) {
	return _Plugman.Contract.PublicSaleStartTimestamp(&_Plugman.CallOpts)
}

// RevealedUntil is a free data retrieval call binding the contract method 0x2bc5e487.
//
// Solidity: function revealedUntil() view returns(uint256)
func (_Plugman *PlugmanCaller) RevealedUntil(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "revealedUntil")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealedUntil is a free data retrieval call binding the contract method 0x2bc5e487.
//
// Solidity: function revealedUntil() view returns(uint256)
func (_Plugman *PlugmanSession) RevealedUntil() (*big.Int, error) {
	return _Plugman.Contract.RevealedUntil(&_Plugman.CallOpts)
}

// RevealedUntil is a free data retrieval call binding the contract method 0x2bc5e487.
//
// Solidity: function revealedUntil() view returns(uint256)
func (_Plugman *PlugmanCallerSession) RevealedUntil() (*big.Int, error) {
	return _Plugman.Contract.RevealedUntil(&_Plugman.CallOpts)
}

// ShuttleMachineAddresses is a free data retrieval call binding the contract method 0xd3c1ac27.
//
// Solidity: function shuttleMachineAddresses(address ) view returns(bool)
func (_Plugman *PlugmanCaller) ShuttleMachineAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "shuttleMachineAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShuttleMachineAddresses is a free data retrieval call binding the contract method 0xd3c1ac27.
//
// Solidity: function shuttleMachineAddresses(address ) view returns(bool)
func (_Plugman *PlugmanSession) ShuttleMachineAddresses(arg0 common.Address) (bool, error) {
	return _Plugman.Contract.ShuttleMachineAddresses(&_Plugman.CallOpts, arg0)
}

// ShuttleMachineAddresses is a free data retrieval call binding the contract method 0xd3c1ac27.
//
// Solidity: function shuttleMachineAddresses(address ) view returns(bool)
func (_Plugman *PlugmanCallerSession) ShuttleMachineAddresses(arg0 common.Address) (bool, error) {
	return _Plugman.Contract.ShuttleMachineAddresses(&_Plugman.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plugman *PlugmanCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plugman *PlugmanSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Plugman.Contract.SupportsInterface(&_Plugman.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Plugman *PlugmanCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Plugman.Contract.SupportsInterface(&_Plugman.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plugman *PlugmanCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plugman *PlugmanSession) Symbol() (string, error) {
	return _Plugman.Contract.Symbol(&_Plugman.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Plugman *PlugmanCallerSession) Symbol() (string, error) {
	return _Plugman.Contract.Symbol(&_Plugman.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plugman *PlugmanCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plugman *PlugmanSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Plugman.Contract.TokenURI(&_Plugman.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Plugman *PlugmanCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Plugman.Contract.TokenURI(&_Plugman.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Plugman *PlugmanCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Plugman *PlugmanSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Plugman.Contract.TokensOfOwner(&_Plugman.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Plugman *PlugmanCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Plugman.Contract.TokensOfOwner(&_Plugman.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Plugman *PlugmanCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Plugman *PlugmanSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Plugman.Contract.TokensOfOwnerIn(&_Plugman.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Plugman *PlugmanCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Plugman.Contract.TokensOfOwnerIn(&_Plugman.CallOpts, owner, start, stop)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Plugman *PlugmanCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Plugman *PlugmanSession) TotalSupply() (*big.Int, error) {
	return _Plugman.Contract.TotalSupply(&_Plugman.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Plugman *PlugmanCallerSession) TotalSupply() (*big.Int, error) {
	return _Plugman.Contract.TotalSupply(&_Plugman.CallOpts)
}

// TraitLabelStorage is a free data retrieval call binding the contract method 0xf1c94c5c.
//
// Solidity: function traitLabelStorage(bytes32 traitKey) view returns((uint8,bool,bool,address))
func (_Plugman *PlugmanCaller) TraitLabelStorage(opts *bind.CallOpts, traitKey [32]byte) (TraitLabelStorage, error) {
	var out []interface{}
	err := _Plugman.contract.Call(opts, &out, "traitLabelStorage", traitKey)

	if err != nil {
		return *new(TraitLabelStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(TraitLabelStorage)).(*TraitLabelStorage)

	return out0, err

}

// TraitLabelStorage is a free data retrieval call binding the contract method 0xf1c94c5c.
//
// Solidity: function traitLabelStorage(bytes32 traitKey) view returns((uint8,bool,bool,address))
func (_Plugman *PlugmanSession) TraitLabelStorage(traitKey [32]byte) (TraitLabelStorage, error) {
	return _Plugman.Contract.TraitLabelStorage(&_Plugman.CallOpts, traitKey)
}

// TraitLabelStorage is a free data retrieval call binding the contract method 0xf1c94c5c.
//
// Solidity: function traitLabelStorage(bytes32 traitKey) view returns((uint8,bool,bool,address))
func (_Plugman *PlugmanCallerSession) TraitLabelStorage(traitKey [32]byte) (TraitLabelStorage, error) {
	return _Plugman.Contract.TraitLabelStorage(&_Plugman.CallOpts, traitKey)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.Approve(&_Plugman.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.Approve(&_Plugman.TransactOpts, to, tokenId)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Plugman *PlugmanTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Plugman *PlugmanSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Plugman.Contract.CancelOwnershipHandover(&_Plugman.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Plugman *PlugmanTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Plugman.Contract.CancelOwnershipHandover(&_Plugman.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Plugman *PlugmanTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Plugman *PlugmanSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.CompleteOwnershipHandover(&_Plugman.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Plugman *PlugmanTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.CompleteOwnershipHandover(&_Plugman.TransactOpts, pendingOwner)
}

// DeleteTrait is a paid mutator transaction binding the contract method 0xce914f02.
//
// Solidity: function deleteTrait(bytes32 traitKey, uint256 tokenId) returns()
func (_Plugman *PlugmanTransactor) DeleteTrait(opts *bind.TransactOpts, traitKey [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "deleteTrait", traitKey, tokenId)
}

// DeleteTrait is a paid mutator transaction binding the contract method 0xce914f02.
//
// Solidity: function deleteTrait(bytes32 traitKey, uint256 tokenId) returns()
func (_Plugman *PlugmanSession) DeleteTrait(traitKey [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.DeleteTrait(&_Plugman.TransactOpts, traitKey, tokenId)
}

// DeleteTrait is a paid mutator transaction binding the contract method 0xce914f02.
//
// Solidity: function deleteTrait(bytes32 traitKey, uint256 tokenId) returns()
func (_Plugman *PlugmanTransactorSession) DeleteTrait(traitKey [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.DeleteTrait(&_Plugman.TransactOpts, traitKey, tokenId)
}

// LockingOwnerTraitModify is a paid mutator transaction binding the contract method 0x714b95c3.
//
// Solidity: function lockingOwnerTraitModify() returns()
func (_Plugman *PlugmanTransactor) LockingOwnerTraitModify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "lockingOwnerTraitModify")
}

// LockingOwnerTraitModify is a paid mutator transaction binding the contract method 0x714b95c3.
//
// Solidity: function lockingOwnerTraitModify() returns()
func (_Plugman *PlugmanSession) LockingOwnerTraitModify() (*types.Transaction, error) {
	return _Plugman.Contract.LockingOwnerTraitModify(&_Plugman.TransactOpts)
}

// LockingOwnerTraitModify is a paid mutator transaction binding the contract method 0x714b95c3.
//
// Solidity: function lockingOwnerTraitModify() returns()
func (_Plugman *PlugmanTransactorSession) LockingOwnerTraitModify() (*types.Transaction, error) {
	return _Plugman.Contract.LockingOwnerTraitModify(&_Plugman.TransactOpts)
}

// MintOG is a paid mutator transaction binding the contract method 0xa68a5555.
//
// Solidity: function mintOG(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanTransactor) MintOG(opts *bind.TransactOpts, to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "mintOG", to, _count, _nonce, timestamp, traitValue, signature)
}

// MintOG is a paid mutator transaction binding the contract method 0xa68a5555.
//
// Solidity: function mintOG(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanSession) MintOG(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintOG(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// MintOG is a paid mutator transaction binding the contract method 0xa68a5555.
//
// Solidity: function mintOG(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanTransactorSession) MintOG(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintOG(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicCrossChain is a paid mutator transaction binding the contract method 0x9b309c55.
//
// Solidity: function mintPublicCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanTransactor) MintPublicCrossChain(opts *bind.TransactOpts, to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "mintPublicCrossChain", to, _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicCrossChain is a paid mutator transaction binding the contract method 0x9b309c55.
//
// Solidity: function mintPublicCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanSession) MintPublicCrossChain(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintPublicCrossChain(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicCrossChain is a paid mutator transaction binding the contract method 0x9b309c55.
//
// Solidity: function mintPublicCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) returns()
func (_Plugman *PlugmanTransactorSession) MintPublicCrossChain(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintPublicCrossChain(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicSale is a paid mutator transaction binding the contract method 0x9be24d8a.
//
// Solidity: function mintPublicSale(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactor) MintPublicSale(opts *bind.TransactOpts, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "mintPublicSale", _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicSale is a paid mutator transaction binding the contract method 0x9be24d8a.
//
// Solidity: function mintPublicSale(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanSession) MintPublicSale(_count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintPublicSale(&_Plugman.TransactOpts, _count, _nonce, timestamp, traitValue, signature)
}

// MintPublicSale is a paid mutator transaction binding the contract method 0x9be24d8a.
//
// Solidity: function mintPublicSale(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactorSession) MintPublicSale(_count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintPublicSale(&_Plugman.TransactOpts, _count, _nonce, timestamp, traitValue, signature)
}

// MintWL is a paid mutator transaction binding the contract method 0xd9ee2d7e.
//
// Solidity: function mintWL(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactor) MintWL(opts *bind.TransactOpts, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "mintWL", _count, _nonce, timestamp, traitValue, signature)
}

// MintWL is a paid mutator transaction binding the contract method 0xd9ee2d7e.
//
// Solidity: function mintWL(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanSession) MintWL(_count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintWL(&_Plugman.TransactOpts, _count, _nonce, timestamp, traitValue, signature)
}

// MintWL is a paid mutator transaction binding the contract method 0xd9ee2d7e.
//
// Solidity: function mintWL(uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactorSession) MintWL(_count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintWL(&_Plugman.TransactOpts, _count, _nonce, timestamp, traitValue, signature)
}

// MintWLCrossChain is a paid mutator transaction binding the contract method 0x6f488cb4.
//
// Solidity: function mintWLCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactor) MintWLCrossChain(opts *bind.TransactOpts, to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "mintWLCrossChain", to, _count, _nonce, timestamp, traitValue, signature)
}

// MintWLCrossChain is a paid mutator transaction binding the contract method 0x6f488cb4.
//
// Solidity: function mintWLCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanSession) MintWLCrossChain(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintWLCrossChain(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// MintWLCrossChain is a paid mutator transaction binding the contract method 0x6f488cb4.
//
// Solidity: function mintWLCrossChain(address to, uint256 _count, uint256 _nonce, uint256 timestamp, bytes32[] traitValue, bytes signature) payable returns()
func (_Plugman *PlugmanTransactorSession) MintWLCrossChain(to common.Address, _count *big.Int, _nonce *big.Int, timestamp *big.Int, traitValue [][32]byte, signature []byte) (*types.Transaction, error) {
	return _Plugman.Contract.MintWLCrossChain(&_Plugman.TransactOpts, to, _count, _nonce, timestamp, traitValue, signature)
}

// OwnerFixTraits is a paid mutator transaction binding the contract method 0x56417179.
//
// Solidity: function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanTransactor) OwnerFixTraits(opts *bind.TransactOpts, tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "ownerFixTraits", tokenId, traitKey, trait)
}

// OwnerFixTraits is a paid mutator transaction binding the contract method 0x56417179.
//
// Solidity: function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanSession) OwnerFixTraits(tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.Contract.OwnerFixTraits(&_Plugman.TransactOpts, tokenId, traitKey, trait)
}

// OwnerFixTraits is a paid mutator transaction binding the contract method 0x56417179.
//
// Solidity: function ownerFixTraits(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanTransactorSession) OwnerFixTraits(tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.Contract.OwnerFixTraits(&_Plugman.TransactOpts, tokenId, traitKey, trait)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Plugman *PlugmanTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Plugman *PlugmanSession) RenounceOwnership() (*types.Transaction, error) {
	return _Plugman.Contract.RenounceOwnership(&_Plugman.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Plugman *PlugmanTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Plugman.Contract.RenounceOwnership(&_Plugman.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Plugman *PlugmanTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Plugman *PlugmanSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Plugman.Contract.RequestOwnershipHandover(&_Plugman.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Plugman *PlugmanTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Plugman.Contract.RequestOwnershipHandover(&_Plugman.TransactOpts)
}

// RevealUntil is a paid mutator transaction binding the contract method 0x4f37ef0d.
//
// Solidity: function revealUntil(uint256 tokenId) returns()
func (_Plugman *PlugmanTransactor) RevealUntil(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "revealUntil", tokenId)
}

// RevealUntil is a paid mutator transaction binding the contract method 0x4f37ef0d.
//
// Solidity: function revealUntil(uint256 tokenId) returns()
func (_Plugman *PlugmanSession) RevealUntil(tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.RevealUntil(&_Plugman.TransactOpts, tokenId)
}

// RevealUntil is a paid mutator transaction binding the contract method 0x4f37ef0d.
//
// Solidity: function revealUntil(uint256 tokenId) returns()
func (_Plugman *PlugmanTransactorSession) RevealUntil(tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.RevealUntil(&_Plugman.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SafeTransferFrom(&_Plugman.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SafeTransferFrom(&_Plugman.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Plugman *PlugmanTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Plugman *PlugmanSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Plugman.Contract.SafeTransferFrom0(&_Plugman.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Plugman *PlugmanTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Plugman.Contract.SafeTransferFrom0(&_Plugman.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Plugman *PlugmanTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Plugman *PlugmanSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Plugman.Contract.SetApprovalForAll(&_Plugman.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Plugman *PlugmanTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Plugman.Contract.SetApprovalForAll(&_Plugman.TransactOpts, operator, approved)
}

// SetBatchMax is a paid mutator transaction binding the contract method 0x1fbbf119.
//
// Solidity: function setBatchMax(uint256 _batchMax) returns()
func (_Plugman *PlugmanTransactor) SetBatchMax(opts *bind.TransactOpts, _batchMax *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setBatchMax", _batchMax)
}

// SetBatchMax is a paid mutator transaction binding the contract method 0x1fbbf119.
//
// Solidity: function setBatchMax(uint256 _batchMax) returns()
func (_Plugman *PlugmanSession) SetBatchMax(_batchMax *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetBatchMax(&_Plugman.TransactOpts, _batchMax)
}

// SetBatchMax is a paid mutator transaction binding the contract method 0x1fbbf119.
//
// Solidity: function setBatchMax(uint256 _batchMax) returns()
func (_Plugman *PlugmanTransactorSession) SetBatchMax(_batchMax *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetBatchMax(&_Plugman.TransactOpts, _batchMax)
}

// SetImageBaseUri is a paid mutator transaction binding the contract method 0x39b92bbb.
//
// Solidity: function setImageBaseUri(string uri) returns()
func (_Plugman *PlugmanTransactor) SetImageBaseUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setImageBaseUri", uri)
}

// SetImageBaseUri is a paid mutator transaction binding the contract method 0x39b92bbb.
//
// Solidity: function setImageBaseUri(string uri) returns()
func (_Plugman *PlugmanSession) SetImageBaseUri(uri string) (*types.Transaction, error) {
	return _Plugman.Contract.SetImageBaseUri(&_Plugman.TransactOpts, uri)
}

// SetImageBaseUri is a paid mutator transaction binding the contract method 0x39b92bbb.
//
// Solidity: function setImageBaseUri(string uri) returns()
func (_Plugman *PlugmanTransactorSession) SetImageBaseUri(uri string) (*types.Transaction, error) {
	return _Plugman.Contract.SetImageBaseUri(&_Plugman.TransactOpts, uri)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 __publicSalePrice) returns()
func (_Plugman *PlugmanTransactor) SetPublicPrice(opts *bind.TransactOpts, __publicSalePrice *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setPublicPrice", __publicSalePrice)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 __publicSalePrice) returns()
func (_Plugman *PlugmanSession) SetPublicPrice(__publicSalePrice *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetPublicPrice(&_Plugman.TransactOpts, __publicSalePrice)
}

// SetPublicPrice is a paid mutator transaction binding the contract method 0xc6275255.
//
// Solidity: function setPublicPrice(uint256 __publicSalePrice) returns()
func (_Plugman *PlugmanTransactorSession) SetPublicPrice(__publicSalePrice *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetPublicPrice(&_Plugman.TransactOpts, __publicSalePrice)
}

// SetPublicSaleTimestamp is a paid mutator transaction binding the contract method 0x511a9605.
//
// Solidity: function setPublicSaleTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanTransactor) SetPublicSaleTimestamp(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setPublicSaleTimestamp", timestamp)
}

// SetPublicSaleTimestamp is a paid mutator transaction binding the contract method 0x511a9605.
//
// Solidity: function setPublicSaleTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanSession) SetPublicSaleTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetPublicSaleTimestamp(&_Plugman.TransactOpts, timestamp)
}

// SetPublicSaleTimestamp is a paid mutator transaction binding the contract method 0x511a9605.
//
// Solidity: function setPublicSaleTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanTransactorSession) SetPublicSaleTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetPublicSaleTimestamp(&_Plugman.TransactOpts, timestamp)
}

// SetSafeRoomAddress is a paid mutator transaction binding the contract method 0xa7043cd4.
//
// Solidity: function setSafeRoomAddress(address __safeRoomAddress) returns()
func (_Plugman *PlugmanTransactor) SetSafeRoomAddress(opts *bind.TransactOpts, __safeRoomAddress common.Address) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setSafeRoomAddress", __safeRoomAddress)
}

// SetSafeRoomAddress is a paid mutator transaction binding the contract method 0xa7043cd4.
//
// Solidity: function setSafeRoomAddress(address __safeRoomAddress) returns()
func (_Plugman *PlugmanSession) SetSafeRoomAddress(__safeRoomAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.SetSafeRoomAddress(&_Plugman.TransactOpts, __safeRoomAddress)
}

// SetSafeRoomAddress is a paid mutator transaction binding the contract method 0xa7043cd4.
//
// Solidity: function setSafeRoomAddress(address __safeRoomAddress) returns()
func (_Plugman *PlugmanTransactorSession) SetSafeRoomAddress(__safeRoomAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.SetSafeRoomAddress(&_Plugman.TransactOpts, __safeRoomAddress)
}

// SetShuttleMachine is a paid mutator transaction binding the contract method 0xea3d8859.
//
// Solidity: function setShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanTransactor) SetShuttleMachine(opts *bind.TransactOpts, __shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setShuttleMachine", __shuttleMachineAddress)
}

// SetShuttleMachine is a paid mutator transaction binding the contract method 0xea3d8859.
//
// Solidity: function setShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanSession) SetShuttleMachine(__shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.SetShuttleMachine(&_Plugman.TransactOpts, __shuttleMachineAddress)
}

// SetShuttleMachine is a paid mutator transaction binding the contract method 0xea3d8859.
//
// Solidity: function setShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanTransactorSession) SetShuttleMachine(__shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.SetShuttleMachine(&_Plugman.TransactOpts, __shuttleMachineAddress)
}

// SetTrait is a paid mutator transaction binding the contract method 0x2bf453e3.
//
// Solidity: function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanTransactor) SetTrait(opts *bind.TransactOpts, tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setTrait", tokenId, traitKey, trait)
}

// SetTrait is a paid mutator transaction binding the contract method 0x2bf453e3.
//
// Solidity: function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanSession) SetTrait(tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.Contract.SetTrait(&_Plugman.TransactOpts, tokenId, traitKey, trait)
}

// SetTrait is a paid mutator transaction binding the contract method 0x2bf453e3.
//
// Solidity: function setTrait(uint256 tokenId, bytes32 traitKey, bytes32 trait) returns()
func (_Plugman *PlugmanTransactorSession) SetTrait(tokenId *big.Int, traitKey [32]byte, trait [32]byte) (*types.Transaction, error) {
	return _Plugman.Contract.SetTrait(&_Plugman.TransactOpts, tokenId, traitKey, trait)
}

// SetTraitLabel is a paid mutator transaction binding the contract method 0x33e5b467.
//
// Solidity: function setTraitLabel(bytes32 traitKey, (string,string,string[],(bytes32,string)[],uint8,uint8,bool) _traitLabel) returns()
func (_Plugman *PlugmanTransactor) SetTraitLabel(opts *bind.TransactOpts, traitKey [32]byte, _traitLabel TraitLabel) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setTraitLabel", traitKey, _traitLabel)
}

// SetTraitLabel is a paid mutator transaction binding the contract method 0x33e5b467.
//
// Solidity: function setTraitLabel(bytes32 traitKey, (string,string,string[],(bytes32,string)[],uint8,uint8,bool) _traitLabel) returns()
func (_Plugman *PlugmanSession) SetTraitLabel(traitKey [32]byte, _traitLabel TraitLabel) (*types.Transaction, error) {
	return _Plugman.Contract.SetTraitLabel(&_Plugman.TransactOpts, traitKey, _traitLabel)
}

// SetTraitLabel is a paid mutator transaction binding the contract method 0x33e5b467.
//
// Solidity: function setTraitLabel(bytes32 traitKey, (string,string,string[],(bytes32,string)[],uint8,uint8,bool) _traitLabel) returns()
func (_Plugman *PlugmanTransactorSession) SetTraitLabel(traitKey [32]byte, _traitLabel TraitLabel) (*types.Transaction, error) {
	return _Plugman.Contract.SetTraitLabel(&_Plugman.TransactOpts, traitKey, _traitLabel)
}

// SetWLPrice is a paid mutator transaction binding the contract method 0xf6a5b8e6.
//
// Solidity: function setWLPrice(uint256 __wlPrice) returns()
func (_Plugman *PlugmanTransactor) SetWLPrice(opts *bind.TransactOpts, __wlPrice *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setWLPrice", __wlPrice)
}

// SetWLPrice is a paid mutator transaction binding the contract method 0xf6a5b8e6.
//
// Solidity: function setWLPrice(uint256 __wlPrice) returns()
func (_Plugman *PlugmanSession) SetWLPrice(__wlPrice *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetWLPrice(&_Plugman.TransactOpts, __wlPrice)
}

// SetWLPrice is a paid mutator transaction binding the contract method 0xf6a5b8e6.
//
// Solidity: function setWLPrice(uint256 __wlPrice) returns()
func (_Plugman *PlugmanTransactorSession) SetWLPrice(__wlPrice *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetWLPrice(&_Plugman.TransactOpts, __wlPrice)
}

// SetWLStartTimestamp is a paid mutator transaction binding the contract method 0xa50fef11.
//
// Solidity: function setWLStartTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanTransactor) SetWLStartTimestamp(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "setWLStartTimestamp", timestamp)
}

// SetWLStartTimestamp is a paid mutator transaction binding the contract method 0xa50fef11.
//
// Solidity: function setWLStartTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanSession) SetWLStartTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetWLStartTimestamp(&_Plugman.TransactOpts, timestamp)
}

// SetWLStartTimestamp is a paid mutator transaction binding the contract method 0xa50fef11.
//
// Solidity: function setWLStartTimestamp(uint256 timestamp) returns()
func (_Plugman *PlugmanTransactorSession) SetWLStartTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.SetWLStartTimestamp(&_Plugman.TransactOpts, timestamp)
}

// TogglePlugable is a paid mutator transaction binding the contract method 0xe2681e18.
//
// Solidity: function togglePlugable() returns()
func (_Plugman *PlugmanTransactor) TogglePlugable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "togglePlugable")
}

// TogglePlugable is a paid mutator transaction binding the contract method 0xe2681e18.
//
// Solidity: function togglePlugable() returns()
func (_Plugman *PlugmanSession) TogglePlugable() (*types.Transaction, error) {
	return _Plugman.Contract.TogglePlugable(&_Plugman.TransactOpts)
}

// TogglePlugable is a paid mutator transaction binding the contract method 0xe2681e18.
//
// Solidity: function togglePlugable() returns()
func (_Plugman *PlugmanTransactorSession) TogglePlugable() (*types.Transaction, error) {
	return _Plugman.Contract.TogglePlugable(&_Plugman.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Plugman *PlugmanTransactor) ToggleSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "toggleSale")
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Plugman *PlugmanSession) ToggleSale() (*types.Transaction, error) {
	return _Plugman.Contract.ToggleSale(&_Plugman.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Plugman *PlugmanTransactorSession) ToggleSale() (*types.Transaction, error) {
	return _Plugman.Contract.ToggleSale(&_Plugman.TransactOpts)
}

// ToggleWLSale is a paid mutator transaction binding the contract method 0xd57f8279.
//
// Solidity: function toggleWLSale() returns()
func (_Plugman *PlugmanTransactor) ToggleWLSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "toggleWLSale")
}

// ToggleWLSale is a paid mutator transaction binding the contract method 0xd57f8279.
//
// Solidity: function toggleWLSale() returns()
func (_Plugman *PlugmanSession) ToggleWLSale() (*types.Transaction, error) {
	return _Plugman.Contract.ToggleWLSale(&_Plugman.TransactOpts)
}

// ToggleWLSale is a paid mutator transaction binding the contract method 0xd57f8279.
//
// Solidity: function toggleWLSale() returns()
func (_Plugman *PlugmanTransactorSession) ToggleWLSale() (*types.Transaction, error) {
	return _Plugman.Contract.ToggleWLSale(&_Plugman.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.TransferFrom(&_Plugman.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Plugman *PlugmanTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Plugman.Contract.TransferFrom(&_Plugman.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Plugman *PlugmanTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Plugman *PlugmanSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.TransferOwnership(&_Plugman.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Plugman *PlugmanTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.TransferOwnership(&_Plugman.TransactOpts, newOwner)
}

// UnsetShuttleMachine is a paid mutator transaction binding the contract method 0xdd17e794.
//
// Solidity: function unsetShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanTransactor) UnsetShuttleMachine(opts *bind.TransactOpts, __shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "unsetShuttleMachine", __shuttleMachineAddress)
}

// UnsetShuttleMachine is a paid mutator transaction binding the contract method 0xdd17e794.
//
// Solidity: function unsetShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanSession) UnsetShuttleMachine(__shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.UnsetShuttleMachine(&_Plugman.TransactOpts, __shuttleMachineAddress)
}

// UnsetShuttleMachine is a paid mutator transaction binding the contract method 0xdd17e794.
//
// Solidity: function unsetShuttleMachine(address __shuttleMachineAddress) returns()
func (_Plugman *PlugmanTransactorSession) UnsetShuttleMachine(__shuttleMachineAddress common.Address) (*types.Transaction, error) {
	return _Plugman.Contract.UnsetShuttleMachine(&_Plugman.TransactOpts, __shuttleMachineAddress)
}

// UpdateCustomEditor is a paid mutator transaction binding the contract method 0xab7d7d7d.
//
// Solidity: function updateCustomEditor(address editor, bool insert) returns()
func (_Plugman *PlugmanTransactor) UpdateCustomEditor(opts *bind.TransactOpts, editor common.Address, insert bool) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "updateCustomEditor", editor, insert)
}

// UpdateCustomEditor is a paid mutator transaction binding the contract method 0xab7d7d7d.
//
// Solidity: function updateCustomEditor(address editor, bool insert) returns()
func (_Plugman *PlugmanSession) UpdateCustomEditor(editor common.Address, insert bool) (*types.Transaction, error) {
	return _Plugman.Contract.UpdateCustomEditor(&_Plugman.TransactOpts, editor, insert)
}

// UpdateCustomEditor is a paid mutator transaction binding the contract method 0xab7d7d7d.
//
// Solidity: function updateCustomEditor(address editor, bool insert) returns()
func (_Plugman *PlugmanTransactorSession) UpdateCustomEditor(editor common.Address, insert bool) (*types.Transaction, error) {
	return _Plugman.Contract.UpdateCustomEditor(&_Plugman.TransactOpts, editor, insert)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Plugman *PlugmanTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Plugman.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Plugman *PlugmanSession) Withdraw() (*types.Transaction, error) {
	return _Plugman.Contract.Withdraw(&_Plugman.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Plugman *PlugmanTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Plugman.Contract.Withdraw(&_Plugman.TransactOpts)
}

// PlugmanApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Plugman contract.
type PlugmanApprovalIterator struct {
	Event *PlugmanApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanApproval represents a Approval event raised by the Plugman contract.
type PlugmanApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PlugmanApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanApprovalIterator{contract: _Plugman.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PlugmanApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanApproval)
				if err := _Plugman.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) ParseApproval(log types.Log) (*PlugmanApproval, error) {
	event := new(PlugmanApproval)
	if err := _Plugman.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Plugman contract.
type PlugmanApprovalForAllIterator struct {
	Event *PlugmanApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanApprovalForAll represents a ApprovalForAll event raised by the Plugman contract.
type PlugmanApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Plugman *PlugmanFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PlugmanApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanApprovalForAllIterator{contract: _Plugman.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Plugman *PlugmanFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PlugmanApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanApprovalForAll)
				if err := _Plugman.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Plugman *PlugmanFilterer) ParseApprovalForAll(log types.Log) (*PlugmanApprovalForAll, error) {
	event := new(PlugmanApprovalForAll)
	if err := _Plugman.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Plugman contract.
type PlugmanConsecutiveTransferIterator struct {
	Event *PlugmanConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Plugman contract.
type PlugmanConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Plugman *PlugmanFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*PlugmanConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanConsecutiveTransferIterator{contract: _Plugman.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Plugman *PlugmanFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *PlugmanConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanConsecutiveTransfer)
				if err := _Plugman.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Plugman *PlugmanFilterer) ParseConsecutiveTransfer(log types.Log) (*PlugmanConsecutiveTransfer, error) {
	event := new(PlugmanConsecutiveTransfer)
	if err := _Plugman.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Plugman contract.
type PlugmanMintIterator struct {
	Event *PlugmanMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanMint represents a Mint event raised by the Plugman contract.
type PlugmanMint struct {
	To           common.Address
	Nonce        *big.Int
	FirstTokenId *big.Int
	MintType     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x542ce94ca9ffbc5633caeef7e530d07d13d22c3760be9a74d31657f90efdf985.
//
// Solidity: event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType)
func (_Plugman *PlugmanFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*PlugmanMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanMintIterator{contract: _Plugman.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x542ce94ca9ffbc5633caeef7e530d07d13d22c3760be9a74d31657f90efdf985.
//
// Solidity: event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType)
func (_Plugman *PlugmanFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PlugmanMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanMint)
				if err := _Plugman.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x542ce94ca9ffbc5633caeef7e530d07d13d22c3760be9a74d31657f90efdf985.
//
// Solidity: event Mint(address indexed to, uint256 nonce, uint256 firstTokenId, uint8 mintType)
func (_Plugman *PlugmanFilterer) ParseMint(log types.Log) (*PlugmanMint, error) {
	event := new(PlugmanMint)
	if err := _Plugman.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Plugman contract.
type PlugmanOwnershipHandoverCanceledIterator struct {
	Event *PlugmanOwnershipHandoverCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanOwnershipHandoverCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanOwnershipHandoverCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Plugman contract.
type PlugmanOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*PlugmanOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanOwnershipHandoverCanceledIterator{contract: _Plugman.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *PlugmanOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanOwnershipHandoverCanceled)
				if err := _Plugman.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*PlugmanOwnershipHandoverCanceled, error) {
	event := new(PlugmanOwnershipHandoverCanceled)
	if err := _Plugman.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Plugman contract.
type PlugmanOwnershipHandoverRequestedIterator struct {
	Event *PlugmanOwnershipHandoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanOwnershipHandoverRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanOwnershipHandoverRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Plugman contract.
type PlugmanOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*PlugmanOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanOwnershipHandoverRequestedIterator{contract: _Plugman.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *PlugmanOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanOwnershipHandoverRequested)
				if err := _Plugman.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Plugman *PlugmanFilterer) ParseOwnershipHandoverRequested(log types.Log) (*PlugmanOwnershipHandoverRequested, error) {
	event := new(PlugmanOwnershipHandoverRequested)
	if err := _Plugman.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Plugman contract.
type PlugmanOwnershipTransferredIterator struct {
	Event *PlugmanOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanOwnershipTransferred represents a OwnershipTransferred event raised by the Plugman contract.
type PlugmanOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Plugman *PlugmanFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PlugmanOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanOwnershipTransferredIterator{contract: _Plugman.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Plugman *PlugmanFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlugmanOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanOwnershipTransferred)
				if err := _Plugman.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Plugman *PlugmanFilterer) ParseOwnershipTransferred(log types.Log) (*PlugmanOwnershipTransferred, error) {
	event := new(PlugmanOwnershipTransferred)
	if err := _Plugman.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitMetadataURIUpdatedIterator is returned from FilterTraitMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for TraitMetadataURIUpdated events raised by the Plugman contract.
type PlugmanTraitMetadataURIUpdatedIterator struct {
	Event *PlugmanTraitMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitMetadataURIUpdated represents a TraitMetadataURIUpdated event raised by the Plugman contract.
type PlugmanTraitMetadataURIUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTraitMetadataURIUpdated is a free log retrieval operation binding the contract event 0x0c42722a91eb9b96ce65a38fc22054e4d2ab7ab642a9c0f92da35c99d965a748.
//
// Solidity: event TraitMetadataURIUpdated()
func (_Plugman *PlugmanFilterer) FilterTraitMetadataURIUpdated(opts *bind.FilterOpts) (*PlugmanTraitMetadataURIUpdatedIterator, error) {

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitMetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitMetadataURIUpdatedIterator{contract: _Plugman.contract, event: "TraitMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchTraitMetadataURIUpdated is a free log subscription operation binding the contract event 0x0c42722a91eb9b96ce65a38fc22054e4d2ab7ab642a9c0f92da35c99d965a748.
//
// Solidity: event TraitMetadataURIUpdated()
func (_Plugman *PlugmanFilterer) WatchTraitMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *PlugmanTraitMetadataURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitMetadataURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitMetadataURIUpdated)
				if err := _Plugman.contract.UnpackLog(event, "TraitMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitMetadataURIUpdated is a log parse operation binding the contract event 0x0c42722a91eb9b96ce65a38fc22054e4d2ab7ab642a9c0f92da35c99d965a748.
//
// Solidity: event TraitMetadataURIUpdated()
func (_Plugman *PlugmanFilterer) ParseTraitMetadataURIUpdated(log types.Log) (*PlugmanTraitMetadataURIUpdated, error) {
	event := new(PlugmanTraitMetadataURIUpdated)
	if err := _Plugman.contract.UnpackLog(event, "TraitMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitUpdatedIterator is returned from FilterTraitUpdated and is used to iterate over the raw logs and unpacked data for TraitUpdated events raised by the Plugman contract.
type PlugmanTraitUpdatedIterator struct {
	Event *PlugmanTraitUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitUpdated represents a TraitUpdated event raised by the Plugman contract.
type PlugmanTraitUpdated struct {
	TraitKey   [32]byte
	TokenId    *big.Int
	TraitValue [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTraitUpdated is a free log retrieval operation binding the contract event 0x8386f3b08e49490d0c5a9d2c401c091f13b01a17d75ce4a2f0f8f923b410ff7d.
//
// Solidity: event TraitUpdated(bytes32 indexed traitKey, uint256 tokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) FilterTraitUpdated(opts *bind.FilterOpts, traitKey [][32]byte) (*PlugmanTraitUpdatedIterator, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitUpdated", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitUpdatedIterator{contract: _Plugman.contract, event: "TraitUpdated", logs: logs, sub: sub}, nil
}

// WatchTraitUpdated is a free log subscription operation binding the contract event 0x8386f3b08e49490d0c5a9d2c401c091f13b01a17d75ce4a2f0f8f923b410ff7d.
//
// Solidity: event TraitUpdated(bytes32 indexed traitKey, uint256 tokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) WatchTraitUpdated(opts *bind.WatchOpts, sink chan<- *PlugmanTraitUpdated, traitKey [][32]byte) (event.Subscription, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitUpdated", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitUpdated)
				if err := _Plugman.contract.UnpackLog(event, "TraitUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitUpdated is a log parse operation binding the contract event 0x8386f3b08e49490d0c5a9d2c401c091f13b01a17d75ce4a2f0f8f923b410ff7d.
//
// Solidity: event TraitUpdated(bytes32 indexed traitKey, uint256 tokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) ParseTraitUpdated(log types.Log) (*PlugmanTraitUpdated, error) {
	event := new(PlugmanTraitUpdated)
	if err := _Plugman.contract.UnpackLog(event, "TraitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitUpdatedListIterator is returned from FilterTraitUpdatedList and is used to iterate over the raw logs and unpacked data for TraitUpdatedList events raised by the Plugman contract.
type PlugmanTraitUpdatedListIterator struct {
	Event *PlugmanTraitUpdatedList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitUpdatedListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitUpdatedList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitUpdatedList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitUpdatedListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitUpdatedListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitUpdatedList represents a TraitUpdatedList event raised by the Plugman contract.
type PlugmanTraitUpdatedList struct {
	TraitKey [32]byte
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTraitUpdatedList is a free log retrieval operation binding the contract event 0xa871bf661500ebea566fc58dd77242d75a9a3f659e5cb12e3bae0a05110561a1.
//
// Solidity: event TraitUpdatedList(bytes32 indexed traitKey, uint256[] tokenIds)
func (_Plugman *PlugmanFilterer) FilterTraitUpdatedList(opts *bind.FilterOpts, traitKey [][32]byte) (*PlugmanTraitUpdatedListIterator, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitUpdatedList", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitUpdatedListIterator{contract: _Plugman.contract, event: "TraitUpdatedList", logs: logs, sub: sub}, nil
}

// WatchTraitUpdatedList is a free log subscription operation binding the contract event 0xa871bf661500ebea566fc58dd77242d75a9a3f659e5cb12e3bae0a05110561a1.
//
// Solidity: event TraitUpdatedList(bytes32 indexed traitKey, uint256[] tokenIds)
func (_Plugman *PlugmanFilterer) WatchTraitUpdatedList(opts *bind.WatchOpts, sink chan<- *PlugmanTraitUpdatedList, traitKey [][32]byte) (event.Subscription, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitUpdatedList", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitUpdatedList)
				if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitUpdatedList is a log parse operation binding the contract event 0xa871bf661500ebea566fc58dd77242d75a9a3f659e5cb12e3bae0a05110561a1.
//
// Solidity: event TraitUpdatedList(bytes32 indexed traitKey, uint256[] tokenIds)
func (_Plugman *PlugmanFilterer) ParseTraitUpdatedList(log types.Log) (*PlugmanTraitUpdatedList, error) {
	event := new(PlugmanTraitUpdatedList)
	if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitUpdatedListUniformValueIterator is returned from FilterTraitUpdatedListUniformValue and is used to iterate over the raw logs and unpacked data for TraitUpdatedListUniformValue events raised by the Plugman contract.
type PlugmanTraitUpdatedListUniformValueIterator struct {
	Event *PlugmanTraitUpdatedListUniformValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitUpdatedListUniformValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitUpdatedListUniformValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitUpdatedListUniformValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitUpdatedListUniformValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitUpdatedListUniformValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitUpdatedListUniformValue represents a TraitUpdatedListUniformValue event raised by the Plugman contract.
type PlugmanTraitUpdatedListUniformValue struct {
	TraitKey   [32]byte
	TokenIds   []*big.Int
	TraitValue [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTraitUpdatedListUniformValue is a free log retrieval operation binding the contract event 0x77679e37d5806f0b2f7f023c0b6124bd74d7abb9154942ee1303c8700b31bda7.
//
// Solidity: event TraitUpdatedListUniformValue(bytes32 indexed traitKey, uint256[] tokenIds, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) FilterTraitUpdatedListUniformValue(opts *bind.FilterOpts, traitKey [][32]byte) (*PlugmanTraitUpdatedListUniformValueIterator, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitUpdatedListUniformValue", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitUpdatedListUniformValueIterator{contract: _Plugman.contract, event: "TraitUpdatedListUniformValue", logs: logs, sub: sub}, nil
}

// WatchTraitUpdatedListUniformValue is a free log subscription operation binding the contract event 0x77679e37d5806f0b2f7f023c0b6124bd74d7abb9154942ee1303c8700b31bda7.
//
// Solidity: event TraitUpdatedListUniformValue(bytes32 indexed traitKey, uint256[] tokenIds, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) WatchTraitUpdatedListUniformValue(opts *bind.WatchOpts, sink chan<- *PlugmanTraitUpdatedListUniformValue, traitKey [][32]byte) (event.Subscription, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitUpdatedListUniformValue", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitUpdatedListUniformValue)
				if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedListUniformValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitUpdatedListUniformValue is a log parse operation binding the contract event 0x77679e37d5806f0b2f7f023c0b6124bd74d7abb9154942ee1303c8700b31bda7.
//
// Solidity: event TraitUpdatedListUniformValue(bytes32 indexed traitKey, uint256[] tokenIds, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) ParseTraitUpdatedListUniformValue(log types.Log) (*PlugmanTraitUpdatedListUniformValue, error) {
	event := new(PlugmanTraitUpdatedListUniformValue)
	if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedListUniformValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitUpdatedRangeIterator is returned from FilterTraitUpdatedRange and is used to iterate over the raw logs and unpacked data for TraitUpdatedRange events raised by the Plugman contract.
type PlugmanTraitUpdatedRangeIterator struct {
	Event *PlugmanTraitUpdatedRange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitUpdatedRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitUpdatedRange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitUpdatedRange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitUpdatedRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitUpdatedRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitUpdatedRange represents a TraitUpdatedRange event raised by the Plugman contract.
type PlugmanTraitUpdatedRange struct {
	TraitKey    [32]byte
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTraitUpdatedRange is a free log retrieval operation binding the contract event 0x9a6c97c48b507ce8a38891ab26b440119b2f725c688fec730db4699bba6b24a7.
//
// Solidity: event TraitUpdatedRange(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId)
func (_Plugman *PlugmanFilterer) FilterTraitUpdatedRange(opts *bind.FilterOpts, traitKey [][32]byte) (*PlugmanTraitUpdatedRangeIterator, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitUpdatedRange", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitUpdatedRangeIterator{contract: _Plugman.contract, event: "TraitUpdatedRange", logs: logs, sub: sub}, nil
}

// WatchTraitUpdatedRange is a free log subscription operation binding the contract event 0x9a6c97c48b507ce8a38891ab26b440119b2f725c688fec730db4699bba6b24a7.
//
// Solidity: event TraitUpdatedRange(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId)
func (_Plugman *PlugmanFilterer) WatchTraitUpdatedRange(opts *bind.WatchOpts, sink chan<- *PlugmanTraitUpdatedRange, traitKey [][32]byte) (event.Subscription, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitUpdatedRange", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitUpdatedRange)
				if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedRange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitUpdatedRange is a log parse operation binding the contract event 0x9a6c97c48b507ce8a38891ab26b440119b2f725c688fec730db4699bba6b24a7.
//
// Solidity: event TraitUpdatedRange(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId)
func (_Plugman *PlugmanFilterer) ParseTraitUpdatedRange(log types.Log) (*PlugmanTraitUpdatedRange, error) {
	event := new(PlugmanTraitUpdatedRange)
	if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTraitUpdatedRangeUniformValueIterator is returned from FilterTraitUpdatedRangeUniformValue and is used to iterate over the raw logs and unpacked data for TraitUpdatedRangeUniformValue events raised by the Plugman contract.
type PlugmanTraitUpdatedRangeUniformValueIterator struct {
	Event *PlugmanTraitUpdatedRangeUniformValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTraitUpdatedRangeUniformValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTraitUpdatedRangeUniformValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTraitUpdatedRangeUniformValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTraitUpdatedRangeUniformValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTraitUpdatedRangeUniformValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTraitUpdatedRangeUniformValue represents a TraitUpdatedRangeUniformValue event raised by the Plugman contract.
type PlugmanTraitUpdatedRangeUniformValue struct {
	TraitKey    [32]byte
	FromTokenId *big.Int
	ToTokenId   *big.Int
	TraitValue  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTraitUpdatedRangeUniformValue is a free log retrieval operation binding the contract event 0xa96ef7c0130afeab54b0046e6e2d01169250d194c74036ac03d0e58ac32422bf.
//
// Solidity: event TraitUpdatedRangeUniformValue(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) FilterTraitUpdatedRangeUniformValue(opts *bind.FilterOpts, traitKey [][32]byte) (*PlugmanTraitUpdatedRangeUniformValueIterator, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "TraitUpdatedRangeUniformValue", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTraitUpdatedRangeUniformValueIterator{contract: _Plugman.contract, event: "TraitUpdatedRangeUniformValue", logs: logs, sub: sub}, nil
}

// WatchTraitUpdatedRangeUniformValue is a free log subscription operation binding the contract event 0xa96ef7c0130afeab54b0046e6e2d01169250d194c74036ac03d0e58ac32422bf.
//
// Solidity: event TraitUpdatedRangeUniformValue(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) WatchTraitUpdatedRangeUniformValue(opts *bind.WatchOpts, sink chan<- *PlugmanTraitUpdatedRangeUniformValue, traitKey [][32]byte) (event.Subscription, error) {

	var traitKeyRule []interface{}
	for _, traitKeyItem := range traitKey {
		traitKeyRule = append(traitKeyRule, traitKeyItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "TraitUpdatedRangeUniformValue", traitKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTraitUpdatedRangeUniformValue)
				if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedRangeUniformValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTraitUpdatedRangeUniformValue is a log parse operation binding the contract event 0xa96ef7c0130afeab54b0046e6e2d01169250d194c74036ac03d0e58ac32422bf.
//
// Solidity: event TraitUpdatedRangeUniformValue(bytes32 indexed traitKey, uint256 fromTokenId, uint256 toTokenId, bytes32 traitValue)
func (_Plugman *PlugmanFilterer) ParseTraitUpdatedRangeUniformValue(log types.Log) (*PlugmanTraitUpdatedRangeUniformValue, error) {
	event := new(PlugmanTraitUpdatedRangeUniformValue)
	if err := _Plugman.contract.UnpackLog(event, "TraitUpdatedRangeUniformValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlugmanTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Plugman contract.
type PlugmanTransferIterator struct {
	Event *PlugmanTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlugmanTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugmanTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlugmanTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlugmanTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugmanTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugmanTransfer represents a Transfer event raised by the Plugman contract.
type PlugmanTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PlugmanTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plugman.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PlugmanTransferIterator{contract: _Plugman.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PlugmanTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Plugman.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugmanTransfer)
				if err := _Plugman.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Plugman *PlugmanFilterer) ParseTransfer(log types.Log) (*PlugmanTransfer, error) {
	event := new(PlugmanTransfer)
	if err := _Plugman.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
