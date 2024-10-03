// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayzevm

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

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// IGatewayZEVMMetaData contains all meta data concerning the IGatewayZEVM contract.
var IGatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"depositAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayZEVMMetaData.ABI instead.
var IGatewayZEVMABI = IGatewayZEVMMetaData.ABI

// IGatewayZEVM is an auto generated Go binding around an Ethereum contract.
type IGatewayZEVM struct {
	IGatewayZEVMCaller     // Read-only binding to the contract
	IGatewayZEVMTransactor // Write-only binding to the contract
	IGatewayZEVMFilterer   // Log filterer for contract events
}

// IGatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayZEVMSession struct {
	Contract     *IGatewayZEVM     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayZEVMCallerSession struct {
	Contract *IGatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayZEVMTransactorSession struct {
	Contract     *IGatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayZEVMRaw struct {
	Contract *IGatewayZEVM // Generic contract binding to access the raw methods on
}

// IGatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayZEVMCallerRaw struct {
	Contract *IGatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayZEVMTransactorRaw struct {
	Contract *IGatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayZEVM creates a new instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVM(address common.Address, backend bind.ContractBackend) (*IGatewayZEVM, error) {
	contract, err := bindIGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVM{IGatewayZEVMCaller: IGatewayZEVMCaller{contract: contract}, IGatewayZEVMTransactor: IGatewayZEVMTransactor{contract: contract}, IGatewayZEVMFilterer: IGatewayZEVMFilterer{contract: contract}}, nil
}

// NewIGatewayZEVMCaller creates a new read-only instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*IGatewayZEVMCaller, error) {
	contract, err := bindIGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMCaller{contract: contract}, nil
}

// NewIGatewayZEVMTransactor creates a new write-only instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayZEVMTransactor, error) {
	contract, err := bindIGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMTransactor{contract: contract}, nil
}

// NewIGatewayZEVMFilterer creates a new log filterer instance of IGatewayZEVM, bound to a specific deployed contract.
func NewIGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayZEVMFilterer, error) {
	contract, err := bindIGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMFilterer{contract: contract}, nil
}

// bindIGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindIGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVM *IGatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVM.Contract.IGatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVM *IGatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.IGatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVM *IGatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.IGatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVM *IGatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVM *IGatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVM *IGatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "call", receiver, message)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Call(receiver []byte, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Call(&_IGatewayZEVM.TransactOpts, receiver, message)
}

// Call is a paid mutator transaction binding the contract method 0x0ac7c44c.
//
// Solidity: function call(bytes receiver, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Call(receiver []byte, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Call(&_IGatewayZEVM.TransactOpts, receiver, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Deposit(&_IGatewayZEVM.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Deposit(&_IGatewayZEVM.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "depositAndCall", context, zrc20, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) DepositAndCall(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) DepositAndCall(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.DepositAndCall(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Execute(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Execute(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Execute(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Execute(&_IGatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20)
}

// Withdraw is a paid mutator transaction binding the contract method 0x135390f9.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.Withdraw(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, message)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7993c1e0.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message) returns()
func (_IGatewayZEVM *IGatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayZEVM.Contract.WithdrawAndCall(&_IGatewayZEVM.TransactOpts, receiver, amount, zrc20, message)
}
