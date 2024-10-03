// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornative

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

// ZetaConnectorNativeMetaData contains all meta data concerning the ZetaConnectorNative contract.
var ZetaConnectorNativeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zetaToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tssAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001638380380620016388339818101604052810190620000379190620001ec565b8282826001600081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480620000aa5750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b80620000e25750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156200011a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b8152505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050506200029b565b600081519050620001e68162000281565b92915050565b6000806000606084860312156200020857620002076200027c565b5b60006200021886828701620001d5565b93505060206200022b86828701620001d5565b92505060406200023e86828701620001d5565b9150509250925092565b600062000255826200025c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200028c8162000248565b81146200029857600080fd5b50565b60805160601c60a05160601c61132b6200030d6000396000818161020201528181610284015281816103f0015281816104b5015281816105b30152818161063501526107130152600081816101e001528181610248015281816104910152818161059101526105f9015261132b6000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806321e093b11161005b57806321e093b1146100d85780635b112591146100f65780635e3e9fef14610114578063743e0c9b146101305761007d565b806302d5c89914610082578063106e62901461009e578063116191b6146100ba575b600080fd5b61009c60048036038101906100979190610c56565b61014c565b005b6100b860048036038101906100b39190610c03565b61035a565b005b6100c261048f565b6040516100cf9190610f68565b60405180910390f35b6100e06104b3565b6040516100ed9190610e9f565b60405180910390f35b6100fe6104d7565b60405161010b9190610e9f565b60405180910390f35b61012e60048036038101906101299190610c56565b6104fd565b005b61014a60048036038101906101459190610d0b565b61070b565b005b61015461075b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101db576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102467f0000000000000000000000000000000000000000000000000000000000000000857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166107ab9092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b8969bd47f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b81526004016102c7959493929190610ef1565b600060405180830381600087803b1580156102e157600080fd5b505af11580156102f5573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167fba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe85858560405161034393929190611040565b60405180910390a2610353610831565b5050505050565b61036261075b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61043483837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166107ab9092919063ffffffff16565b8273ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161047a9190611025565b60405180910390a261048a610831565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61050561075b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461058c576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105f77f0000000000000000000000000000000000000000000000000000000000000000857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166107ab9092919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635131ab597f0000000000000000000000000000000000000000000000000000000000000000878787876040518663ffffffff1660e01b8152600401610678959493929190610ef1565b600060405180830381600087803b15801561069257600080fd5b505af11580156106a6573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff167f7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced8585856040516106f493929190611040565b60405180910390a2610704610831565b5050505050565b6107583330837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661083b909392919063ffffffff16565b50565b600260005414156107a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079890611005565b60405180910390fd5b6002600081905550565b61082c8363a9059cbb60e01b84846040516024016107ca929190610f3f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506108c4565b505050565b6001600081905550565b6108be846323b872dd60e01b85858560405160240161085c93929190610eba565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506108c4565b50505050565b6000610926826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661098b9092919063ffffffff16565b905060008151111561098657808060200190518101906109469190610cde565b610985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097c90610fe5565b60405180910390fd5b5b505050565b606061099a84846000856109a3565b90509392505050565b6060824710156109e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109df90610fa5565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610a119190610e88565b60006040518083038185875af1925050503d8060008114610a4e576040519150601f19603f3d011682016040523d82523d6000602084013e610a53565b606091505b5091509150610a6487838387610a70565b92505050949350505050565b60608315610ad357600083511415610acb57610a8b85610ae6565b610aca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac190610fc5565b60405180910390fd5b5b829050610ade565b610add8383610b09565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115610b1c5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b509190610f83565b60405180910390fd5b600081359050610b6881611299565b92915050565b600081519050610b7d816112b0565b92915050565b600081359050610b92816112c7565b92915050565b60008083601f840112610bae57610bad611184565b5b8235905067ffffffffffffffff811115610bcb57610bca61117f565b5b602083019150836001820283011115610be757610be6611189565b5b9250929050565b600081359050610bfd816112de565b92915050565b600080600060608486031215610c1c57610c1b611193565b5b6000610c2a86828701610b59565b9350506020610c3b86828701610bee565b9250506040610c4c86828701610b83565b9150509250925092565b600080600080600060808688031215610c7257610c71611193565b5b6000610c8088828901610b59565b9550506020610c9188828901610bee565b945050604086013567ffffffffffffffff811115610cb257610cb161118e565b5b610cbe88828901610b98565b93509350506060610cd188828901610b83565b9150509295509295909350565b600060208284031215610cf457610cf3611193565b5b6000610d0284828501610b6e565b91505092915050565b600060208284031215610d2157610d20611193565b5b6000610d2f84828501610bee565b91505092915050565b610d41816110b5565b82525050565b6000610d538385611088565b9350610d6083858461113d565b610d6983611198565b840190509392505050565b6000610d7f82611072565b610d898185611099565b9350610d9981856020860161114c565b80840191505092915050565b610dae81611107565b82525050565b6000610dbf8261107d565b610dc981856110a4565b9350610dd981856020860161114c565b610de281611198565b840191505092915050565b6000610dfa6026836110a4565b9150610e05826111a9565b604082019050919050565b6000610e1d601d836110a4565b9150610e28826111f8565b602082019050919050565b6000610e40602a836110a4565b9150610e4b82611221565b604082019050919050565b6000610e63601f836110a4565b9150610e6e82611270565b602082019050919050565b610e82816110fd565b82525050565b6000610e948284610d74565b915081905092915050565b6000602082019050610eb46000830184610d38565b92915050565b6000606082019050610ecf6000830186610d38565b610edc6020830185610d38565b610ee96040830184610e79565b949350505050565b6000608082019050610f066000830188610d38565b610f136020830187610d38565b610f206040830186610e79565b8181036060830152610f33818486610d47565b90509695505050505050565b6000604082019050610f546000830185610d38565b610f616020830184610e79565b9392505050565b6000602082019050610f7d6000830184610da5565b92915050565b60006020820190508181036000830152610f9d8184610db4565b905092915050565b60006020820190508181036000830152610fbe81610ded565b9050919050565b60006020820190508181036000830152610fde81610e10565b9050919050565b60006020820190508181036000830152610ffe81610e33565b9050919050565b6000602082019050818103600083015261101e81610e56565b9050919050565b600060208201905061103a6000830184610e79565b92915050565b60006040820190506110556000830186610e79565b8181036020830152611068818486610d47565b9050949350505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006110c0826110dd565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061111282611119565b9050919050565b60006111248261112b565b9050919050565b6000611136826110dd565b9050919050565b82818337600083830152505050565b60005b8381101561116a57808201518184015260208101905061114f565b83811115611179576000848401525b50505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6112a2816110b5565b81146112ad57600080fd5b50565b6112b9816110c7565b81146112c457600080fd5b50565b6112d0816110d3565b81146112db57600080fd5b50565b6112e7816110fd565b81146112f257600080fd5b5056fea2646970667358221220655dba4ad3270e6153f071458a8ebcb80a69a23cb0bf8a44ed4a1f4df173d65664736f6c63430008070033",
}

// ZetaConnectorNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeMetaData.ABI instead.
var ZetaConnectorNativeABI = ZetaConnectorNativeMetaData.ABI

// ZetaConnectorNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeMetaData.Bin instead.
var ZetaConnectorNativeBin = ZetaConnectorNativeMetaData.Bin

// DeployZetaConnectorNative deploys a new Ethereum contract, binding an instance of ZetaConnectorNative to it.
func DeployZetaConnectorNative(auth *bind.TransactOpts, backend bind.ContractBackend, _gateway common.Address, _zetaToken common.Address, _tssAddress common.Address) (common.Address, *types.Transaction, *ZetaConnectorNative, error) {
	parsed, err := ZetaConnectorNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeBin), backend, _gateway, _zetaToken, _tssAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNative{ZetaConnectorNativeCaller: ZetaConnectorNativeCaller{contract: contract}, ZetaConnectorNativeTransactor: ZetaConnectorNativeTransactor{contract: contract}, ZetaConnectorNativeFilterer: ZetaConnectorNativeFilterer{contract: contract}}, nil
}

// ZetaConnectorNative is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNative struct {
	ZetaConnectorNativeCaller     // Read-only binding to the contract
	ZetaConnectorNativeTransactor // Write-only binding to the contract
	ZetaConnectorNativeFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeSession struct {
	Contract     *ZetaConnectorNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeCallerSession struct {
	Contract *ZetaConnectorNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZetaConnectorNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeTransactorSession struct {
	Contract     *ZetaConnectorNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeRaw struct {
	Contract *ZetaConnectorNative // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeCallerRaw struct {
	Contract *ZetaConnectorNativeCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTransactorRaw struct {
	Contract *ZetaConnectorNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNative creates a new instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNative(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNative, error) {
	contract, err := bindZetaConnectorNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNative{ZetaConnectorNativeCaller: ZetaConnectorNativeCaller{contract: contract}, ZetaConnectorNativeTransactor: ZetaConnectorNativeTransactor{contract: contract}, ZetaConnectorNativeFilterer: ZetaConnectorNativeFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeCaller creates a new read-only instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeCaller, error) {
	contract, err := bindZetaConnectorNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeCaller{contract: contract}, nil
}

// NewZetaConnectorNativeTransactor creates a new write-only instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeTransactor, error) {
	contract, err := bindZetaConnectorNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeFilterer creates a new log filterer instance of ZetaConnectorNative, bound to a specific deployed contract.
func NewZetaConnectorNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeFilterer, error) {
	contract, err := bindZetaConnectorNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeFilterer{contract: contract}, nil
}

// bindZetaConnectorNative binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNative *ZetaConnectorNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ZetaConnectorNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNative *ZetaConnectorNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNative.Contract.Gateway(&_ZetaConnectorNative.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNative.Contract.Gateway(&_ZetaConnectorNative.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNative.Contract.TssAddress(&_ZetaConnectorNative.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNative.Contract.TssAddress(&_ZetaConnectorNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNative.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNative.Contract.ZetaToken(&_ZetaConnectorNative.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNative *ZetaConnectorNativeCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNative.Contract.ZetaToken(&_ZetaConnectorNative.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ReceiveTokens(&_ZetaConnectorNative.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.ReceiveTokens(&_ZetaConnectorNative.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Withdraw(&_ZetaConnectorNative.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.Withdraw(&_ZetaConnectorNative.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndCall(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndCall(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndRevert(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNative *ZetaConnectorNativeTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNative.Contract.WithdrawAndRevert(&_ZetaConnectorNative.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawIterator struct {
	Event *ZetaConnectorNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdraw)
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
		it.Event = new(ZetaConnectorNativeWithdraw)
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
func (it *ZetaConnectorNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdraw represents a Withdraw event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawIterator{contract: _ZetaConnectorNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdraw)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNativeWithdraw, error) {
	event := new(ZetaConnectorNativeWithdraw)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndCallIterator struct {
	Event *ZetaConnectorNativeWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNativeWithdrawAndCall)
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
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawAndCallIterator{contract: _ZetaConnectorNative.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdrawAndCall)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNativeWithdrawAndCall, error) {
	event := new(ZetaConnectorNativeWithdrawAndCall)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndRevertIterator struct {
	Event *ZetaConnectorNativeWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorNativeWithdrawAndRevert)
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
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorNative contract.
type ZetaConnectorNativeWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeWithdrawAndRevertIterator{contract: _ZetaConnectorNative.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNative.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeWithdrawAndRevert)
				if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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
func (_ZetaConnectorNative *ZetaConnectorNativeFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorNativeWithdrawAndRevert, error) {
	event := new(ZetaConnectorNativeWithdrawAndRevert)
	if err := _ZetaConnectorNative.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
