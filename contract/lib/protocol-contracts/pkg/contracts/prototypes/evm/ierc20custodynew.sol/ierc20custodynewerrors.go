// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20custodynew

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

// IERC20CustodyNewErrorsMetaData contains all meta data concerning the IERC20CustodyNewErrors contract.
var IERC20CustodyNewErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"}]",
}

// IERC20CustodyNewErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyNewErrorsMetaData.ABI instead.
var IERC20CustodyNewErrorsABI = IERC20CustodyNewErrorsMetaData.ABI

// IERC20CustodyNewErrors is an auto generated Go binding around an Ethereum contract.
type IERC20CustodyNewErrors struct {
	IERC20CustodyNewErrorsCaller     // Read-only binding to the contract
	IERC20CustodyNewErrorsTransactor // Write-only binding to the contract
	IERC20CustodyNewErrorsFilterer   // Log filterer for contract events
}

// IERC20CustodyNewErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyNewErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyNewErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyNewErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodyNewErrorsSession struct {
	Contract     *IERC20CustodyNewErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20CustodyNewErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyNewErrorsCallerSession struct {
	Contract *IERC20CustodyNewErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IERC20CustodyNewErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyNewErrorsTransactorSession struct {
	Contract     *IERC20CustodyNewErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IERC20CustodyNewErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyNewErrorsRaw struct {
	Contract *IERC20CustodyNewErrors // Generic contract binding to access the raw methods on
}

// IERC20CustodyNewErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyNewErrorsCallerRaw struct {
	Contract *IERC20CustodyNewErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyNewErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyNewErrorsTransactorRaw struct {
	Contract *IERC20CustodyNewErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20CustodyNewErrors creates a new instance of IERC20CustodyNewErrors, bound to a specific deployed contract.
func NewIERC20CustodyNewErrors(address common.Address, backend bind.ContractBackend) (*IERC20CustodyNewErrors, error) {
	contract, err := bindIERC20CustodyNewErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewErrors{IERC20CustodyNewErrorsCaller: IERC20CustodyNewErrorsCaller{contract: contract}, IERC20CustodyNewErrorsTransactor: IERC20CustodyNewErrorsTransactor{contract: contract}, IERC20CustodyNewErrorsFilterer: IERC20CustodyNewErrorsFilterer{contract: contract}}, nil
}

// NewIERC20CustodyNewErrorsCaller creates a new read-only instance of IERC20CustodyNewErrors, bound to a specific deployed contract.
func NewIERC20CustodyNewErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyNewErrorsCaller, error) {
	contract, err := bindIERC20CustodyNewErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewErrorsCaller{contract: contract}, nil
}

// NewIERC20CustodyNewErrorsTransactor creates a new write-only instance of IERC20CustodyNewErrors, bound to a specific deployed contract.
func NewIERC20CustodyNewErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyNewErrorsTransactor, error) {
	contract, err := bindIERC20CustodyNewErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewErrorsTransactor{contract: contract}, nil
}

// NewIERC20CustodyNewErrorsFilterer creates a new log filterer instance of IERC20CustodyNewErrors, bound to a specific deployed contract.
func NewIERC20CustodyNewErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyNewErrorsFilterer, error) {
	contract, err := bindIERC20CustodyNewErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewErrorsFilterer{contract: contract}, nil
}

// bindIERC20CustodyNewErrors binds a generic wrapper to an already deployed contract.
func bindIERC20CustodyNewErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyNewErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyNewErrors.Contract.IERC20CustodyNewErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyNewErrors.Contract.IERC20CustodyNewErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyNewErrors.Contract.IERC20CustodyNewErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyNewErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyNewErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyNewErrors *IERC20CustodyNewErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyNewErrors.Contract.contract.Transact(opts, method, params...)
}
