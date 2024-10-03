// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornonnative

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

// ZetaConnectorNonNativeMetaData contains all meta data concerning the ZetaConnectorNonNative contract.
var ZetaConnectorNonNativeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zetaToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tssAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6002553480156200003557600080fd5b506040516200158d3803806200158d83398181016040528101906200005b919062000210565b8282826001600081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480620000ce5750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620001065750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156200013e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b8152505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050620002bf565b6000815190506200020a81620002a5565b92915050565b6000806000606084860312156200022c576200022b620002a0565b5b60006200023c86828701620001f9565b93505060206200024f86828701620001f9565b92505060406200026286828701620001f9565b9150509250925092565b6000620002798262000280565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620002b0816200026c565b8114620002bc57600080fd5b50565b60805160601c60a05160601c611247620003466000396000818161023001528181610311015281816103fc0152818161056601528181610647015281816107550152818161083101528181610912015281816109fd0152610b9d01526000818161034d015281816103c0015281816107310152818161094e01526109c101526112476000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635b112591116100665780635b1125911461010c5780635e3e9fef1461012a5780636f8b44b014610146578063743e0c9b14610162578063d5abeb011461017e57610093565b806302d5c89914610098578063106e6290146100b4578063116191b6146100d057806321e093b1146100ee575b600080fd5b6100b260048036038101906100ad9190610d88565b61019c565b005b6100ce60048036038101906100c99190610d35565b6104d2565b005b6100d861072f565b6040516100e59190610fbf565b60405180910390f35b6100f6610753565b6040516101039190610ef6565b60405180910390f35b610114610777565b6040516101219190610ef6565b60405180910390f35b610144600480360381019061013f9190610d88565b61079d565b005b610160600480360381019061015b9190610e10565b610ad3565b005b61017c60048036038101906101779190610e10565b610b9b565b005b610186610c2b565b6040516101939190610ffa565b60405180910390f35b6101a4610c31565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461022b576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561029457600080fd5b505afa1580156102a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cc9190610e3d565b856102d79190611069565b111561030f576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee7f000000000000000000000000000000000000000000000000000000000000000086846040518463ffffffff1660e01b815260040161038c93929190610f88565b600060405180830381600087803b1580156103a657600080fd5b505af11580156103ba573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b8969bd47f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b815260040161043f959493929190610f11565b600060405180830381600087803b15801561045957600080fd5b505af115801561046d573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe8585856040516104bb93929190611015565b60405180910390a26104cb610c81565b5050505050565b6104da610c31565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610561576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156105ca57600080fd5b505afa1580156105de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106029190610e3d565b8361060d9190611069565b1115610645576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee8484846040518463ffffffff1660e01b81526004016106a293929190610f88565b600060405180830381600087803b1580156106bc57600080fd5b505af11580156106d0573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161071a9190610ffa565b60405180910390a261072a610c81565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6107a5610c31565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461082c576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561089557600080fd5b505afa1580156108a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108cd9190610e3d565b856108d89190611069565b1115610910576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631e458bee7f000000000000000000000000000000000000000000000000000000000000000086846040518463ffffffff1660e01b815260040161098d93929190610f88565b600060405180830381600087803b1580156109a757600080fd5b505af11580156109bb573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b8152600401610a40959493929190610f11565b600060405180830381600087803b158015610a5a57600080fd5b505af1158015610a6e573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced858585604051610abc93929190611015565b60405180910390a2610acc610c81565b5050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b5a576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002819055507f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c81604051610b909190610ffa565b60405180910390a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379cc679033836040518363ffffffff1660e01b8152600401610bf6929190610f5f565b600060405180830381600087803b158015610c1057600080fd5b505af1158015610c24573d6000803e3d6000fd5b5050505050565b60025481565b60026000541415610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90610fda565b60405180910390fd5b6002600081905550565b6001600081905550565b600081359050610c9a816111cc565b92915050565b600081359050610caf816111e3565b92915050565b60008083601f840112610ccb57610cca61117e565b5b8235905067ffffffffffffffff811115610ce857610ce7611179565b5b602083019150836001820283011115610d0457610d03611183565b5b9250929050565b600081359050610d1a816111fa565b92915050565b600081519050610d2f816111fa565b92915050565b600080600060608486031215610d4e57610d4d61118d565b5b6000610d5c86828701610c8b565b9350506020610d6d86828701610d0b565b9250506040610d7e86828701610ca0565b9150509250925092565b600080600080600060808688031215610da457610da361118d565b5b6000610db288828901610c8b565b9550506020610dc388828901610d0b565b945050604086013567ffffffffffffffff811115610de457610de3611188565b5b610df088828901610cb5565b93509350506060610e0388828901610ca0565b9150509295509295909350565b600060208284031215610e2657610e2561118d565b5b6000610e3484828501610d0b565b91505092915050565b600060208284031215610e5357610e5261118d565b5b6000610e6184828501610d20565b91505092915050565b610e73816110bf565b82525050565b610e82816110d1565b82525050565b6000610e948385611047565b9350610ea183858461113b565b610eaa83611192565b840190509392505050565b610ebe81611105565b82525050565b6000610ed1601f83611058565b9150610edc826111a3565b602082019050919050565b610ef0816110fb565b82525050565b6000602082019050610f0b6000830184610e6a565b92915050565b6000608082019050610f266000830188610e6a565b610f336020830187610e6a565b610f406040830186610ee7565b8181036060830152610f53818486610e88565b90509695505050505050565b6000604082019050610f746000830185610e6a565b610f816020830184610ee7565b9392505050565b6000606082019050610f9d6000830186610e6a565b610faa6020830185610ee7565b610fb76040830184610e79565b949350505050565b6000602082019050610fd46000830184610eb5565b92915050565b60006020820190508181036000830152610ff381610ec4565b9050919050565b600060208201905061100f6000830184610ee7565b92915050565b600060408201905061102a6000830186610ee7565b818103602083015261103d818486610e88565b9050949350505050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611074826110fb565b915061107f836110fb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156110b4576110b361114a565b5b828201905092915050565b60006110ca826110db565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061111082611117565b9050919050565b600061112282611129565b9050919050565b6000611134826110db565b9050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6111d5816110bf565b81146111e057600080fd5b50565b6111ec816110d1565b81146111f757600080fd5b50565b611203816110fb565b811461120e57600080fd5b5056fea264697066735822122044af4e37ef2ad6dd470fc1c3af732b6c36970257d67f6756302aee086afdfa4d64736f6c63430008070033",
}

// ZetaConnectorNonNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNonNativeMetaData.ABI instead.
var ZetaConnectorNonNativeABI = ZetaConnectorNonNativeMetaData.ABI

// ZetaConnectorNonNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNonNativeMetaData.Bin instead.
var ZetaConnectorNonNativeBin = ZetaConnectorNonNativeMetaData.Bin

// DeployZetaConnectorNonNative deploys a new Ethereum contract, binding an instance of ZetaConnectorNonNative to it.
func DeployZetaConnectorNonNative(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address, _tssAddress common.Address) (common.Address, *types.Transaction, *ZetaConnectorNonNative, error) {
	parsed, err := ZetaConnectorNonNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNonNativeBin), backend, _gateway, _zetaToken, _tssAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNonNative{ZetaConnectorNonNativeCaller: ZetaConnectorNonNativeCaller{contract: contract}, ZetaConnectorNonNativeTransactor: ZetaConnectorNonNativeTransactor{contract: contract}, ZetaConnectorNonNativeFilterer: ZetaConnectorNonNativeFilterer{contract: contract}}, nil
}

// ZetaConnectorNonNative is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNonNative struct {
	ZetaConnectorNonNativeCaller     // Read-only binding to the contract
	ZetaConnectorNonNativeTransactor // Write-only binding to the contract
	ZetaConnectorNonNativeFilterer   // Log filterer for contract events
}

// ZetaConnectorNonNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNonNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNonNativeSession struct {
	Contract     *ZetaConnectorNonNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNonNativeCallerSession struct {
	Contract *ZetaConnectorNonNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaConnectorNonNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNonNativeTransactorSession struct {
	Contract     *ZetaConnectorNonNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNonNativeRaw struct {
	Contract *ZetaConnectorNonNative // Generic contract binding to access the raw methods on
}

// ZetaConnectorNonNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeCallerRaw struct {
	Contract *ZetaConnectorNonNativeCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNonNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeTransactorRaw struct {
	Contract *ZetaConnectorNonNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNonNative creates a new instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNative(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNonNative, error) {
	contract, err := bindZetaConnectorNonNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNative{ZetaConnectorNonNativeCaller: ZetaConnectorNonNativeCaller{contract: contract}, ZetaConnectorNonNativeTransactor: ZetaConnectorNonNativeTransactor{contract: contract}, ZetaConnectorNonNativeFilterer: ZetaConnectorNonNativeFilterer{contract: contract}}, nil
}

// NewZetaConnectorNonNativeCaller creates a new read-only instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNonNativeCaller, error) {
	contract, err := bindZetaConnectorNonNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeCaller{contract: contract}, nil
}

// NewZetaConnectorNonNativeTransactor creates a new write-only instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNonNativeTransactor, error) {
	contract, err := bindZetaConnectorNonNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeTransactor{contract: contract}, nil
}

// NewZetaConnectorNonNativeFilterer creates a new log filterer instance of ZetaConnectorNonNative, bound to a specific deployed contract.
func NewZetaConnectorNonNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNonNativeFilterer, error) {
	contract, err := bindZetaConnectorNonNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeFilterer{contract: contract}, nil
}

// bindZetaConnectorNonNative binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNonNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNonNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ZetaConnectorNonNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.Gateway(&_ZetaConnectorNonNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.Gateway(&_ZetaConnectorNonNative.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNative.Contract.MaxSupply(&_ZetaConnectorNonNative.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNative.Contract.MaxSupply(&_ZetaConnectorNonNative.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.TssAddress(&_ZetaConnectorNonNative.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.TssAddress(&_ZetaConnectorNonNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNative.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.ZetaToken(&_ZetaConnectorNonNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNative.Contract.ZetaToken(&_ZetaConnectorNonNative.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ReceiveTokens(&_ZetaConnectorNonNative.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.ReceiveTokens(&_ZetaConnectorNonNative.TransactOpts, amount)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.SetMaxSupply(&_ZetaConnectorNonNative.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.SetMaxSupply(&_ZetaConnectorNonNative.TransactOpts, _maxSupply)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Withdraw(&_ZetaConnectorNonNative.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.Withdraw(&_ZetaConnectorNonNative.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndCall(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndCall(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndRevert(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNative.Contract.WithdrawAndRevert(&_ZetaConnectorNonNative.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNonNativeMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeMaxSupplyUpdatedIterator struct {
	Event *ZetaConnectorNonNativeMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeMaxSupplyUpdated)
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
		it.Event = new(ZetaConnectorNonNativeMaxSupplyUpdated)
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
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeMaxSupplyUpdated struct {
	MaxSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonNativeMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeMaxSupplyUpdatedIterator{contract: _ZetaConnectorNonNative.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeMaxSupplyUpdated)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseMaxSupplyUpdated(log types.Log) (*ZetaConnectorNonNativeMaxSupplyUpdated, error) {
	event := new(ZetaConnectorNonNativeMaxSupplyUpdated)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawIterator struct {
	Event *ZetaConnectorNonNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdraw)
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
		it.Event = new(ZetaConnectorNonNativeWithdraw)
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
func (it *ZetaConnectorNonNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdraw represents a Withdraw event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawIterator{contract: _ZetaConnectorNonNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdraw)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNonNativeWithdraw, error) {
	event := new(ZetaConnectorNonNativeWithdraw)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndCallIterator struct {
	Event *ZetaConnectorNonNativeWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNonNativeWithdrawAndCall)
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
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawAndCallIterator{contract: _ZetaConnectorNonNative.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdrawAndCall)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNonNativeWithdrawAndCall, error) {
	event := new(ZetaConnectorNonNativeWithdrawAndCall)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndRevertIterator struct {
	Event *ZetaConnectorNonNativeWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorNonNativeWithdrawAndRevert)
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
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorNonNative contract.
type ZetaConnectorNonNativeWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeWithdrawAndRevertIterator{contract: _ZetaConnectorNonNative.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNative.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeWithdrawAndRevert)
				if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNative *ZetaConnectorNonNativeFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorNonNativeWithdrawAndRevert, error) {
	event := new(ZetaConnectorNonNativeWithdrawAndRevert)
	if err := _ZetaConnectorNonNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
