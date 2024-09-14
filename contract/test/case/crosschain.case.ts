import { expect } from 'chai';
import axios from 'axios';
import dotenv from 'dotenv';
import { prepareData } from "@zetachain/toolkit/client";
import {
    createPublicClient,
    createWalletClient,
    http,
    Address,
    parseSignature,
    parseAbi,
    parseUnits,
    encodeAbiParameters, parseAbiParameters,
    BaseError, ContractFunctionRevertedError,
    getAddress,
    ByteArray, stringToBytes, toBytes, stringToHex, Hex, toHex, parseEther, parseGwei,
} from 'viem';
import {privateKeyToAccount} from 'viem/accounts';
import {zetachainAthensTestnet, zetachain, sepolia, polygonAmoy, mainnet, polygon, anvil} from 'viem/chains';

dotenv.config();

describe('Cross Chain Mint Test', function() {
    const ORDER_INIT = 1;
    const ORDER_SUCCESS = 2
    const ORDER_FAILED  = 3
    const PLUGMAN_ABI = [{"type":"constructor","inputs":[{"name":"__name","type":"string","internalType":"string"},{"name":"__symbol","type":"string","internalType":"string"},{"name":"__mintVerifier","type":"address","internalType":"address"},{"name":"__batchMax","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"PublicSaleNonce","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"WLLocked","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"WLNonce","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"WLStartTimestamp","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"approve","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"balanceOf","inputs":[{"name":"owner","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"batchMax","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"cancelOwnershipHandover","inputs":[],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"completeOwnershipHandover","inputs":[{"name":"pendingOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"crossChainMinting","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"deleteTrait","inputs":[{"name":"traitKey","type":"bytes32","internalType":"bytes32"},{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"explicitOwnershipOf","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"ownership","type":"tuple","internalType":"struct IERC721A.TokenOwnership","components":[{"name":"addr","type":"address","internalType":"address"},{"name":"startTimestamp","type":"uint64","internalType":"uint64"},{"name":"burned","type":"bool","internalType":"bool"},{"name":"extraData","type":"uint24","internalType":"uint24"}]}],"stateMutability":"view"},{"type":"function","name":"explicitOwnershipsOf","inputs":[{"name":"tokenIds","type":"uint256[]","internalType":"uint256[]"}],"outputs":[{"name":"","type":"tuple[]","internalType":"struct IERC721A.TokenOwnership[]","components":[{"name":"addr","type":"address","internalType":"address"},{"name":"startTimestamp","type":"uint64","internalType":"uint64"},{"name":"burned","type":"bool","internalType":"bool"},{"name":"extraData","type":"uint24","internalType":"uint24"}]}],"stateMutability":"view"},{"type":"function","name":"getApproved","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getCustomEditorAt","inputs":[{"name":"index","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getCustomEditors","inputs":[],"outputs":[{"name":"","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"getCustomEditorsLength","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getTraitMetadataURI","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"getTraitValue","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"},{"name":"traitKey","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"traitValue","type":"bytes32","internalType":"bytes32"}],"stateMutability":"view"},{"type":"function","name":"getTraitValues","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"},{"name":"traitKeys","type":"bytes32[]","internalType":"bytes32[]"}],"outputs":[{"name":"traitValues","type":"bytes32[]","internalType":"bytes32[]"}],"stateMutability":"view"},{"type":"function","name":"isApprovedForAll","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"operator","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isCustomEditor","inputs":[{"name":"editor","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isPublicSaleOpen","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isWLSaleOpen","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"lockingOwnerTraitModify","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"mintOG","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"_count","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"signature","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"mintPublicCrossChain","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"_count","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"signature","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"mintPublicSale","inputs":[{"name":"_count","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"signature","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"mintVerify","inputs":[{"name":"_signature","type":"bytes","internalType":"bytes"},{"name":"_mintTo","type":"address","internalType":"address"},{"name":"_amount","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"_timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"_mintType","type":"uint8","internalType":"uint8"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"mintWL","inputs":[{"name":"_count","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"signature","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"mintWLCrossChain","inputs":[{"name":"to","type":"address","internalType":"address"},{"name":"_count","type":"uint256","internalType":"uint256"},{"name":"_nonce","type":"uint256","internalType":"uint256"},{"name":"timestamp","type":"uint256","internalType":"uint256"},{"name":"traitValue","type":"bytes32[]","internalType":"bytes32[]"},{"name":"signature","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"name","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"result","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"ownerFixTraits","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"},{"name":"traitKey","type":"bytes32","internalType":"bytes32"},{"name":"trait","type":"bytes32","internalType":"bytes32"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"ownerOf","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"ownershipHandoverExpiresAt","inputs":[{"name":"pendingOwner","type":"address","internalType":"address"}],"outputs":[{"name":"result","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"plugable","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"publicSaleLocked","inputs":[],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"publicSaleStartTimestamp","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"renounceOwnership","inputs":[],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"requestOwnershipHandover","inputs":[],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"reveal","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"safeTransferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"safeTransferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"tokenId","type":"uint256","internalType":"uint256"},{"name":"_data","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"setApprovalForAll","inputs":[{"name":"operator","type":"address","internalType":"address"},{"name":"approved","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setBatchMax","inputs":[{"name":"_batchMax","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setImageBaseUri","inputs":[{"name":"uri","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPublicPrice","inputs":[{"name":"__publicSalePrice","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPublicSaleTimestamp","inputs":[{"name":"timestamp","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSafeRoomAddress","inputs":[{"name":"__safeRoomAddress","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setShuttleMachine","inputs":[{"name":"__shuttleMachineAddress","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTrait","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"},{"name":"traitKey","type":"bytes32","internalType":"bytes32"},{"name":"trait","type":"bytes32","internalType":"bytes32"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTraitLabel","inputs":[{"name":"traitKey","type":"bytes32","internalType":"bytes32"},{"name":"_traitLabel","type":"tuple","internalType":"struct TraitLabel","components":[{"name":"fullTraitKey","type":"string","internalType":"string"},{"name":"traitLabel","type":"string","internalType":"string"},{"name":"acceptableValues","type":"string[]","internalType":"string[]"},{"name":"fullTraitValues","type":"tuple[]","internalType":"struct FullTraitValue[]","components":[{"name":"traitValue","type":"bytes32","internalType":"bytes32"},{"name":"fullTraitValue","type":"string","internalType":"string"}]},{"name":"displayType","type":"uint8","internalType":"enum DisplayType"},{"name":"editors","type":"uint8","internalType":"Editors"},{"name":"required","type":"bool","internalType":"bool"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setWLPrice","inputs":[{"name":"__wlPrice","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setWLStartTimestamp","inputs":[{"name":"timestamp","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"supportsInterface","inputs":[{"name":"interfaceId","type":"bytes4","internalType":"bytes4"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"symbol","inputs":[],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"toggleCrossChain","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"togglePlugable","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"toggleSale","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"toggleWLSale","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"tokenURI","inputs":[{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"string","internalType":"string"}],"stateMutability":"view"},{"type":"function","name":"tokensOfOwner","inputs":[{"name":"owner","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256[]","internalType":"uint256[]"}],"stateMutability":"view"},{"type":"function","name":"tokensOfOwnerIn","inputs":[{"name":"owner","type":"address","internalType":"address"},{"name":"start","type":"uint256","internalType":"uint256"},{"name":"stop","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256[]","internalType":"uint256[]"}],"stateMutability":"view"},{"type":"function","name":"totalSupply","inputs":[],"outputs":[{"name":"result","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"traitLabelStorage","inputs":[{"name":"traitKey","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"","type":"tuple","internalType":"struct TraitLabelStorage","components":[{"name":"allowedEditors","type":"uint8","internalType":"Editors"},{"name":"required","type":"bool","internalType":"bool"},{"name":"valuesRequireValidation","type":"bool","internalType":"bool"},{"name":"storedLabel","type":"address","internalType":"StoredTraitLabel"}]}],"stateMutability":"view"},{"type":"function","name":"transferFrom","inputs":[{"name":"from","type":"address","internalType":"address"},{"name":"to","type":"address","internalType":"address"},{"name":"tokenId","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"transferOwnership","inputs":[{"name":"newOwner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"payable"},{"type":"function","name":"updateCustomEditor","inputs":[{"name":"editor","type":"address","internalType":"address"},{"name":"insert","type":"bool","internalType":"bool"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"withdraw","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"Approval","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"approved","type":"address","indexed":true,"internalType":"address"},{"name":"tokenId","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"ApprovalForAll","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"operator","type":"address","indexed":true,"internalType":"address"},{"name":"approved","type":"bool","indexed":false,"internalType":"bool"}],"anonymous":false},{"type":"event","name":"ConsecutiveTransfer","inputs":[{"name":"fromTokenId","type":"uint256","indexed":true,"internalType":"uint256"},{"name":"toTokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"Mint","inputs":[{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"nonce","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"firstTokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"mintType","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"OwnershipHandoverCanceled","inputs":[{"name":"pendingOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"OwnershipHandoverRequested","inputs":[{"name":"pendingOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"OwnershipTransferred","inputs":[{"name":"oldOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"TraitMetadataURIUpdated","inputs":[],"anonymous":false},{"type":"event","name":"TraitUpdated","inputs":[{"name":"traitKey","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"tokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"traitValue","type":"bytes32","indexed":false,"internalType":"bytes32"}],"anonymous":false},{"type":"event","name":"TraitUpdatedList","inputs":[{"name":"traitKey","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"tokenIds","type":"uint256[]","indexed":false,"internalType":"uint256[]"}],"anonymous":false},{"type":"event","name":"TraitUpdatedListUniformValue","inputs":[{"name":"traitKey","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"tokenIds","type":"uint256[]","indexed":false,"internalType":"uint256[]"},{"name":"traitValue","type":"bytes32","indexed":false,"internalType":"bytes32"}],"anonymous":false},{"type":"event","name":"TraitUpdatedRange","inputs":[{"name":"traitKey","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"fromTokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"toTokenId","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"TraitUpdatedRangeUniformValue","inputs":[{"name":"traitKey","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"fromTokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"toTokenId","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"traitValue","type":"bytes32","indexed":false,"internalType":"bytes32"}],"anonymous":false},{"type":"event","name":"Transfer","inputs":[{"name":"from","type":"address","indexed":true,"internalType":"address"},{"name":"to","type":"address","indexed":true,"internalType":"address"},{"name":"tokenId","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"error","name":"AllTokenHasBeenMinted","inputs":[]},{"type":"error","name":"AlreadyInitialized","inputs":[]},{"type":"error","name":"ApprovalCallerNotOwnerNorApproved","inputs":[]},{"type":"error","name":"ApprovalQueryForNonexistentToken","inputs":[]},{"type":"error","name":"BalanceQueryForZeroAddress","inputs":[]},{"type":"error","name":"CallerNotApproved","inputs":[]},{"type":"error","name":"CanNotCallMintFromContract","inputs":[]},{"type":"error","name":"InsufficientPrivilege","inputs":[]},{"type":"error","name":"InvalidBatchMax","inputs":[]},{"type":"error","name":"InvalidCoinValue","inputs":[]},{"type":"error","name":"InvalidMintAmount","inputs":[]},{"type":"error","name":"InvalidQueryRange","inputs":[]},{"type":"error","name":"InvalidSignature","inputs":[]},{"type":"error","name":"InvalidTimestamp","inputs":[]},{"type":"error","name":"InvalidTraitKeys","inputs":[]},{"type":"error","name":"InvalidTraitValue","inputs":[]},{"type":"error","name":"IsNotPlugable","inputs":[]},{"type":"error","name":"MintERC2309QuantityExceedsLimit","inputs":[]},{"type":"error","name":"MintToZeroAddress","inputs":[]},{"type":"error","name":"MintZeroQuantity","inputs":[]},{"type":"error","name":"NewOwnerIsZeroAddress","inputs":[]},{"type":"error","name":"NoHandoverRequest","inputs":[]},{"type":"error","name":"NonceIsTooOld","inputs":[]},{"type":"error","name":"NotCompatibleWithSpotMints","inputs":[]},{"type":"error","name":"NotMeetCrossMintingStatus","inputs":[]},{"type":"error","name":"OwnerCannotFixTraits","inputs":[]},{"type":"error","name":"OwnerQueryForNonexistentToken","inputs":[]},{"type":"error","name":"OwnershipNotInitializedForExtraData","inputs":[]},{"type":"error","name":"PublicSaleIsNotActive","inputs":[]},{"type":"error","name":"ReentrancyGuardReentrantCall","inputs":[]},{"type":"error","name":"SequentialMintExceedsLimit","inputs":[]},{"type":"error","name":"SequentialUpToTooSmall","inputs":[]},{"type":"error","name":"SpotMintTokenIdTooSmall","inputs":[]},{"type":"error","name":"TokenAlreadyExists","inputs":[]},{"type":"error","name":"TraitDoesNotExist","inputs":[{"name":"traitKey","type":"bytes32","internalType":"bytes32"}]},{"type":"error","name":"TraitValueNotEmpty","inputs":[]},{"type":"error","name":"TraitValueUnchanged","inputs":[]},{"type":"error","name":"TransferCallerNotOwnerNorApproved","inputs":[]},{"type":"error","name":"TransferFromIncorrectOwner","inputs":[]},{"type":"error","name":"TransferToNonERC721ReceiverImplementer","inputs":[]},{"type":"error","name":"TransferToZeroAddress","inputs":[]},{"type":"error","name":"URIQueryForNonexistentToken","inputs":[]},{"type":"error","name":"Unauthorized","inputs":[]},{"type":"error","name":"WLSaleIsNotActive","inputs":[]},{"type":"error","name":"WithdrawFailed","inputs":[]}]
    const PLUGMAN_CONTRACT = "0x67Ace24E885F1ec2e3FA33d1BA375520cf5C22FF";
    const SHUTTLE_MACHINE_ABI = [
        {
            "type": "function",
            "name": "mockOnCrossChainCall",
            "inputs": [
                {
                    "name": "message",
                    "type": "bytes",
                    "internalType": "bytes"
                }
            ],
            "outputs": [],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "onCrossChainCall",
            "inputs": [
                {
                    "name": "context",
                    "type": "tuple",
                    "internalType": "struct Interface.zContext",
                    "components": [
                        {
                            "name": "origin",
                            "type": "bytes",
                            "internalType": "bytes"
                        },
                        {
                            "name": "sender",
                            "type": "address",
                            "internalType": "address"
                        },
                        {
                            "name": "chainID",
                            "type": "uint256",
                            "internalType": "uint256"
                        }
                    ]
                },
                {
                    "name": "zrc20",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                },
                {
                    "name": "message",
                    "type": "bytes",
                    "internalType": "bytes"
                }
            ],
            "outputs": [],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "plugmanAddress",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "address",
                    "internalType": "address"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "systemContract",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "address",
                    "internalType": "address"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "treasuryAddress",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "address",
                    "internalType": "address"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "event",
            "name": "CrossChainMint",
            "inputs": [
                {
                    "name": "to",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "sender",
                    "type": "address",
                    "indexed": false,
                    "internalType": "address"
                },
                {
                    "name": "count",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                },
                {
                    "name": "nonce",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                },
                {
                    "name": "mintType",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "error",
            "name": "CallerNotOwnerNotApproved",
            "inputs": []
        },
        {
            "type": "error",
            "name": "InvalidBalance",
            "inputs": []
        },
        {
            "type": "error",
            "name": "InvalidCoinAddress",
            "inputs": []
        },
        {
            "type": "error",
            "name": "InvalidCoinValue",
            "inputs": []
        },
        {
            "type": "error",
            "name": "InvalidSender",
            "inputs": []
        },
        {
            "type": "error",
            "name": "NotSupportThisMintType",
            "inputs": []
        },
        {
            "type": "error",
            "name": "OnlySystemContract",
            "inputs": [
                {
                    "name": "",
                    "type": "string",
                    "internalType": "string"
                }
            ]
        },
        {
            "type": "error",
            "name": "OriginChainNotSupported",
            "inputs": []
        },
        {
            "type": "error",
            "name": "TransferToTreasuryFailed",
            "inputs": []
        }
    ]
    const ETHEREUM_SHUTTLE_MACHINE_CONTRACT = "0x4C89a64E762a3Ff9589027a2189d8FF9c76Bd1FE";
    const POLYGON_SHUTTLE_MACHINE_CONTRACT = "0xddf3375d8aeD7B230Fdf81aBbBd85d77215971Ab";
    const ZRC20_ABI = [
        {
            "type": "function",
            "name": "PROTOCOL_FEE",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "allowance",
            "inputs": [
                {
                    "name": "owner",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "spender",
                    "type": "address",
                    "internalType": "address"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "approve",
            "inputs": [
                {
                    "name": "spender",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "balanceOf",
            "inputs": [
                {
                    "name": "account",
                    "type": "address",
                    "internalType": "address"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "burn",
            "inputs": [
                {
                    "name": "account",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "decreaseAllowance",
            "inputs": [
                {
                    "name": "spender",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "deposit",
            "inputs": [
                {
                    "name": "to",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "increaseAllowance",
            "inputs": [
                {
                    "name": "spender",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "totalSupply",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "function",
            "name": "transfer",
            "inputs": [
                {
                    "name": "recipient",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "transferFrom",
            "inputs": [
                {
                    "name": "sender",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "recipient",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "withdraw",
            "inputs": [
                {
                    "name": "to",
                    "type": "bytes",
                    "internalType": "bytes"
                },
                {
                    "name": "amount",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "outputs": [
                {
                    "name": "",
                    "type": "bool",
                    "internalType": "bool"
                }
            ],
            "stateMutability": "nonpayable"
        },
        {
            "type": "function",
            "name": "withdrawGasFee",
            "inputs": [],
            "outputs": [
                {
                    "name": "",
                    "type": "address",
                    "internalType": "address"
                },
                {
                    "name": "",
                    "type": "uint256",
                    "internalType": "uint256"
                }
            ],
            "stateMutability": "view"
        },
        {
            "type": "event",
            "name": "Approval",
            "inputs": [
                {
                    "name": "owner",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "spender",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "value",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "Deposit",
            "inputs": [
                {
                    "name": "from",
                    "type": "bytes",
                    "indexed": false,
                    "internalType": "bytes"
                },
                {
                    "name": "to",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "value",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "Transfer",
            "inputs": [
                {
                    "name": "from",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "to",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "value",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "UpdatedGasLimit",
            "inputs": [
                {
                    "name": "gasLimit",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "UpdatedProtocolFlatFee",
            "inputs": [
                {
                    "name": "protocolFlatFee",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "UpdatedSystemContract",
            "inputs": [
                {
                    "name": "systemContract",
                    "type": "address",
                    "indexed": false,
                    "internalType": "address"
                }
            ],
            "anonymous": false
        },
        {
            "type": "event",
            "name": "Withdrawal",
            "inputs": [
                {
                    "name": "from",
                    "type": "address",
                    "indexed": true,
                    "internalType": "address"
                },
                {
                    "name": "to",
                    "type": "bytes",
                    "indexed": false,
                    "internalType": "bytes"
                },
                {
                    "name": "value",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                },
                {
                    "name": "gasFee",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                },
                {
                    "name": "protocolFlatFee",
                    "type": "uint256",
                    "indexed": false,
                    "internalType": "uint256"
                }
            ],
            "anonymous": false
        }
    ]
    const ZRC20_ETH = "0x9fd96203f7b22bCF72d9DCb40ff98302376cE09c"
    const TREASURY_ADDRESS = "0x6F881627057b37B12163118e09F7e7901096Ccea";
    const ZETA_TSS_ADDRESS = "0x8531a5aB847ff5B22D855633C25ED1DA3255247e";
    const WL_MINT_TYPE = 2;
    const WL_SALE_PRICE = parseEther('0.029');
    const PUBLIC_MINT_TYPE = 3;
    const PUBLIC_SALE_PRICE = parseEther('0.049');
    const privateKey = "0x".concat(process.env.PRIVATE_KEY);
    const privateTestKey = "0x".concat(process.env.PRIVATE_TEST_KEY);
    const account = privateKeyToAccount(privateKey as `0x${string}`);
    const testAccount = privateKeyToAccount(privateTestKey as `0x${string}`);
    const apiUrl = "http://127.0.0.1:8080/api/v1";
    const walletSepoliaClient = createWalletClient({
        account: account,
        chain: sepolia,
        transport: http("https://responsive-multi-pool.ethereum-sepolia.quiknode.pro/ce6e72ffa3ed18c4158cb0fc6c757f9bd7d2b0ba/")
    });

    const publicSepoliaClient = createPublicClient({
        chain: sepolia,
        transport: http()
    });

    const walletAmoyClient = createWalletClient({
        account: account,
        chain: polygonAmoy,
        transport: http()
    });
    const publicAmoyClient = createPublicClient({
        chain: polygonAmoy,
        transport: http()
    });

    const walletAthensClient = createWalletClient({
        account: account,
        chain: zetachainAthensTestnet,
        transport: http()
    });
    const publicAthensClient = createPublicClient({
        chain: zetachainAthensTestnet,
        transport: http()
    });

    interface GetSignatureResponseData {
        count: string;
        nonce: string;
        timestamp: string;
        traitValue: string[];
        signature: string;
        mintType: number;
    }

    interface GetSignatureResponse {
        data: GetSignatureResponseData;
        message: string;
        returnCode: number;
    }

    interface PostOrderRequest {
        chain_id: string;
        tx_hash: string;
        from_address: string;
        mint_type: number;
        nonce: number;
    }

    interface TssTransfer {
        from_address: string;
        amount: string;
        tx_hash: string;
        mint_type: number;
        nonce: number;
        order_status: number;
    }

    interface GetOrderResponseData{
        tss_transfer: TssTransfer;
        token_id: string[];
    }

    interface GetOrderResponse {
        data: GetOrderResponseData;
        message: string;
        returnCode: number;
    }

    const getSignature = async function(to: Address, mintType: number, count: number): Promise<GetSignatureResponseData | undefined>  {
        const url = `${apiUrl}/mint/${to}/${mintType}/${count}/getSignature`;
        try {
            const response = await axios.get<GetSignatureResponse>(url);
            if (response.status === 200) {
                console.log('Response: ', response.data);
                if (response.data.returnCode === 20000) {
                    return response.data.data
                }
            }
        } catch(error) {
            console.error('Error fetching signature: ', error);
            return undefined
        }
    }

    const postAndGetOrder = async function(address: Address, hash: string, mintType: number, nonce: number, chain_id: string) {
        const order_url = `${apiUrl}/order/${account.address.toString()}`
        const payload: PostOrderRequest = {
            chain_id:     chain_id,
            tx_hash:      hash,
            from_address: address,
            mint_type:    mintType,
            nonce:        nonce
        }
        const post_order_response = await axios.post(order_url, payload)
        if (post_order_response.status !== 200) {
            expect.fail(post_order_response.statusText);
        }
        if (post_order_response.data.returnCode !== 20000) {
            expect.fail('The post order return code is not 20000')
        }
        console.log("Post order successfully")

        const get_order_response = await axios.get<GetOrderResponse>(order_url)
        if (get_order_response.status !== 200) {
            expect.fail(get_order_response.statusText);
        }
        if (get_order_response.data.returnCode !== 20000) {
            expect.fail('The get orders return code is not 20000')
        }
        console.log("Get order successfully", JSON.stringify(get_order_response.data.data))
    }

    const prepareCrossChainData = function(signatureData: GetSignatureResponseData, shuttleMachine: Address) : string {
        const count = parseUnits(signatureData.count, 0);
        const nonce = parseUnits(signatureData.nonce, 0);
        const timestamp = parseUnits(signatureData.timestamp, 0);
        const mintType = BigInt(signatureData.mintType);
        const traits: Hex[] = signatureData.traitValue.map(str =>
            toHex(toBytes(str, { size: 32 })));
        const signature: Hex = toHex(toBytes(signatureData.signature));
        const data = encodeAbiParameters(
            [
                {
                    components: [
                        { type: "address", name: "to" },
                        { type: "uint256", name: "count" },
                        { type: "uint256", name: "nonce" },
                        { type: "uint256", name: "timestamp" },
                        { type: "uint256", name: "mintType" },
                        { type: "bytes32[]", name: "traitValue" },
                        { type: "bytes", name: "signature" }
                    ],
                    type: "tuple"
                }
            ],
            [{
                to: account.address,
                count: count,
                nonce: nonce,
                timestamp: timestamp,
                mintType: mintType,
                traitValue: traits,
                signature: signature}
            ]
        )
        return prepareData(
            shuttleMachine,
            ['bytes'],
            [data.toString()]
        )
    }

    const openWLSale = async function() {
        let isOpen = await publicAthensClient.readContract({
            address: PLUGMAN_CONTRACT,
            abi: PLUGMAN_ABI,
            functionName: 'isWLSaleOpen',
        })
        if (!isOpen) {
            const isPublicOpen = await publicAthensClient.readContract({
                address: PLUGMAN_CONTRACT,
                abi: PLUGMAN_ABI,
                functionName: 'isPublicSaleOpen',
            })
            if (isPublicOpen) {
                // toggle public sale
                const hash = await walletAthensClient.writeContract({
                    address: PLUGMAN_CONTRACT,
                    abi: PLUGMAN_ABI,
                    functionName: 'toggleSale',
                    chain: zetachainAthensTestnet,
                    account: account,
                });
                console.log("Hash: ", hash as string);
                const transactionReceipt = await publicAthensClient.waitForTransactionReceipt({
                    hash: hash,
                })
                console.log("TransactionReceipt for toggleSale: ", transactionReceipt)
            }
            isOpen = await publicAthensClient.readContract({
                address: PLUGMAN_CONTRACT,
                abi: PLUGMAN_ABI,
                functionName: 'isWLSaleOpen',
            })
            if (!isOpen) {
                const hash = await walletAthensClient.writeContract({
                    address: PLUGMAN_CONTRACT,
                    abi: PLUGMAN_ABI,
                    functionName: 'toggleWLSale',
                    chain: zetachainAthensTestnet,
                    account: account,
                });
                console.log("Hash: ", hash as string);
                const transactionReceipt = await publicAthensClient.waitForTransactionReceipt({
                    hash: hash,
                })
                console.log("TransactionReceipt for toggleWLSale: ", transactionReceipt)
                isOpen = await publicAthensClient.readContract({
                    address: PLUGMAN_CONTRACT,
                    abi: PLUGMAN_ABI,
                    functionName: 'isWLSaleOpen',
                })
                expect(isOpen as boolean).to.eq(true);
            }
        }
    }

    it('Should response to http request', async function() {
        const response = await axios.get(apiUrl);
        expect(response.status).to.eq(200);
    })

    it('Should start public sales', async function() {
        this.timeout(180000);
        let isOpen = await publicAthensClient.readContract({
            address: PLUGMAN_CONTRACT,
            abi: PLUGMAN_ABI,
            functionName: 'isPublicSaleOpen',
        })
        let isWlOpen = await publicAthensClient.readContract({
            address: PLUGMAN_CONTRACT,
            abi: PLUGMAN_ABI,
            functionName: 'isWLSaleOpen',
        })
        if (isOpen as boolean === false && isWlOpen as boolean === false) {
            console.log("Sale not started")
            let hash = await walletAthensClient.writeContract({
                address: PLUGMAN_CONTRACT,
                abi: PLUGMAN_ABI,
                functionName: 'setWLStartTimestamp',
                args: [1718117400],
                chain: zetachainAthensTestnet,
                account: account,
            });
            console.log("Hash: ", hash as string);
            let transactionReceipt = await publicAthensClient.waitForTransactionReceipt({
                hash: hash,
            })
            console.log("TransactionReceipt: ", transactionReceipt)
            hash = await walletAthensClient.writeContract({
                address: PLUGMAN_CONTRACT,
                abi: PLUGMAN_ABI,
                functionName: 'setPublicSaleTimestamp',
                args: [1718117500],
                chain: zetachainAthensTestnet,
                account: account,
            });
            console.log("Hash: ", hash as string);
            transactionReceipt = await publicAthensClient.waitForTransactionReceipt({
                hash: hash,
            })
            console.log("TransactionReceipt: ", transactionReceipt)
        }

        isOpen = await publicAthensClient.readContract({
            address: PLUGMAN_CONTRACT,
            abi: PLUGMAN_ABI,
            functionName: 'isPublicSaleOpen',
        })
        expect(isOpen as boolean).to.eq(true);
    })

//     it('Should mint WL cross chain via polygon', async function(){
//         this.timeout(180000);
//         await openWLSale();
//         const count = 1;
//         const signatureData = await getSignature(account.address, WL_MINT_TYPE, count);
//         if (signatureData !== undefined) {
//             const prepareData = prepareCrossChainData(signatureData, POLYGON_SHUTTLE_MACHINE_CONTRACT)
//             const value = WL_SALE_PRICE * BigInt(count);
//             try {
//                 const hash = await walletAmoyClient.sendTransaction({
//                     data: prepareData as `0x${string}`,
//                     to: getAddress(ZETA_TSS_ADDRESS),
//                     value: value,
//                     kzg: undefined as any,
//                     chain: polygonAmoy,
//                     account: account,
//                 })
//                 console.log("Hash: ", hash as string);
//                 const transactionReceipt = await publicAmoyClient.waitForTransactionReceipt({
//                     hash: hash,
//                 })
//                 console.log("TransactionReceipt: ", transactionReceipt)
//                 await postAndGetOrder(account.address, hash as string, WL_MINT_TYPE, Number(signatureData.nonce), "80002")
//             } catch (error: any) {
//                 console.error('Unexpected error:', error);
//                 expect.fail('Unexpected error: ' + error);
//             }
//         } else {
//             console.error('Failed to retrieve the signature.');
//             expect.fail('Unexpected error: get signature error');
//         }
//     })
//
//     // it('Should mint cross chain via polygon', async function() {
//     //     this.timeout(60000);
//     //     const count = 1;
//     //     const signatureData = await getSignature(testAccount.address, PUBLIC_MINT_TYPE, count);
//     //     if (signatureData !== undefined) {
//     //         const prepareData = prepareCrossChainData(signatureData, POLYGON_SHUTTLE_MACHINE_CONTRACT)
//     //         const value = PUBLIC_SALE_PRICE * BigInt(count);
//     //         try {
//     //             const hash = await walletSepoliaClient.sendTransaction({
//     //                 data: prepareData as `0x${string}`,
//     //                 to: getAddress(ZETA_TSS_ADDRESS),
//     //                 value: value,
//     //                 kzg: undefined as any,
//     //                 chain: sepolia,
//     //                 account: testAccount,
//     //             })
//     //             console.log("Hash: ", hash as string);
//     //             const transactionReceipt = await publicSepoliaClient.waitForTransactionReceipt({
//     //                 hash: hash,
//     //             })
//     //             console.log("TransactionReceipt: ", transactionReceipt)
//     //             await postAndGetOrder(testAccount.address, hash as string, WL_MINT_TYPE, Number(signatureData.nonce), "80002")
//     //         } catch (error: any) {
//     //             console.error('Unexpected error:', error);
//     //             expect.fail('Unexpected error: ' + error);
//     //         }
//     //     } else {
//     //         console.error('Failed to retrieve the signature.');
//     //         expect.fail('Unexpected error: get signature error');
//     //     }
//     // })
//
//     it('should own some NFTs',  async function()  {
//         const amount = await publicAthensClient.readContract({
//             address: PLUGMAN_CONTRACT,
//             abi: PLUGMAN_ABI,
//             functionName: 'balanceOf',
//             args: [account.address]
//         })
//         const amountBig = amount as bigint;
//         console.log("User owns NFT amount: ", amountBig.toString());
//     });
})