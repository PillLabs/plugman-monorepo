// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumerzevm

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

// ZetaTokenConsumerZEVMErrorsMetaData contains all meta data concerning the ZetaTokenConsumerZEVMErrors contract.
var ZetaTokenConsumerZEVMErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidForZEVM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWZETAAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"}]",
}

// ZetaTokenConsumerZEVMErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerZEVMErrorsMetaData.ABI instead.
var ZetaTokenConsumerZEVMErrorsABI = ZetaTokenConsumerZEVMErrorsMetaData.ABI

// ZetaTokenConsumerZEVMErrors is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrors struct {
	ZetaTokenConsumerZEVMErrorsCaller     // Read-only binding to the contract
	ZetaTokenConsumerZEVMErrorsTransactor // Write-only binding to the contract
	ZetaTokenConsumerZEVMErrorsFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerZEVMErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerZEVMErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerZEVMErrorsSession struct {
	Contract     *ZetaTokenConsumerZEVMErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerZEVMErrorsCallerSession struct {
	Contract *ZetaTokenConsumerZEVMErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ZetaTokenConsumerZEVMErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerZEVMErrorsTransactorSession struct {
	Contract     *ZetaTokenConsumerZEVMErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrorsRaw struct {
	Contract *ZetaTokenConsumerZEVMErrors // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrorsCallerRaw struct {
	Contract *ZetaTokenConsumerZEVMErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMErrorsTransactorRaw struct {
	Contract *ZetaTokenConsumerZEVMErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerZEVMErrors creates a new instance of ZetaTokenConsumerZEVMErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMErrors(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerZEVMErrors, error) {
	contract, err := bindZetaTokenConsumerZEVMErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMErrors{ZetaTokenConsumerZEVMErrorsCaller: ZetaTokenConsumerZEVMErrorsCaller{contract: contract}, ZetaTokenConsumerZEVMErrorsTransactor: ZetaTokenConsumerZEVMErrorsTransactor{contract: contract}, ZetaTokenConsumerZEVMErrorsFilterer: ZetaTokenConsumerZEVMErrorsFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerZEVMErrorsCaller creates a new read-only instance of ZetaTokenConsumerZEVMErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerZEVMErrorsCaller, error) {
	contract, err := bindZetaTokenConsumerZEVMErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMErrorsCaller{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMErrorsTransactor creates a new write-only instance of ZetaTokenConsumerZEVMErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerZEVMErrorsTransactor, error) {
	contract, err := bindZetaTokenConsumerZEVMErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMErrorsTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMErrorsFilterer creates a new log filterer instance of ZetaTokenConsumerZEVMErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerZEVMErrorsFilterer, error) {
	contract, err := bindZetaTokenConsumerZEVMErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMErrorsFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerZEVMErrors binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerZEVMErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerZEVMErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVMErrors.Contract.ZetaTokenConsumerZEVMErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVMErrors.Contract.ZetaTokenConsumerZEVMErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVMErrors.Contract.ZetaTokenConsumerZEVMErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVMErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVMErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVMErrors *ZetaTokenConsumerZEVMErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVMErrors.Contract.contract.Transact(opts, method, params...)
}
