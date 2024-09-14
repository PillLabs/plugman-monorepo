// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shuttleMachine

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

// ShuttleMachineMetaData contains all meta data concerning the ShuttleMachine contract.
var ShuttleMachineMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"__systemContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"__plugmanAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"__treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"__permitChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"__wlPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"__publicSalePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCrossChainCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structzContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"plugmanAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicSalePrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSystemContract\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasuryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wlPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CrossChainMint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotOwnerNotApproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCoinAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCoinValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSupportThisMintType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"OriginChainNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToTreasuryFailed\",\"inputs\":[]}]",
}

// ShuttleMachineABI is the input ABI used to generate the binding from.
// Deprecated: Use ShuttleMachineMetaData.ABI instead.
var ShuttleMachineABI = ShuttleMachineMetaData.ABI

// ShuttleMachine is an auto generated Go binding around an Ethereum contract.
type ShuttleMachine struct {
	ShuttleMachineCaller     // Read-only binding to the contract
	ShuttleMachineTransactor // Write-only binding to the contract
	ShuttleMachineFilterer   // Log filterer for contract events
}

// ShuttleMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShuttleMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttleMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShuttleMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttleMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShuttleMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttleMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShuttleMachineSession struct {
	Contract     *ShuttleMachine   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShuttleMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShuttleMachineCallerSession struct {
	Contract *ShuttleMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ShuttleMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShuttleMachineTransactorSession struct {
	Contract     *ShuttleMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ShuttleMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShuttleMachineRaw struct {
	Contract *ShuttleMachine // Generic contract binding to access the raw methods on
}

// ShuttleMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShuttleMachineCallerRaw struct {
	Contract *ShuttleMachineCaller // Generic read-only contract binding to access the raw methods on
}

// ShuttleMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShuttleMachineTransactorRaw struct {
	Contract *ShuttleMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShuttleMachine creates a new instance of ShuttleMachine, bound to a specific deployed contract.
func NewShuttleMachine(address common.Address, backend bind.ContractBackend) (*ShuttleMachine, error) {
	contract, err := bindShuttleMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShuttleMachine{ShuttleMachineCaller: ShuttleMachineCaller{contract: contract}, ShuttleMachineTransactor: ShuttleMachineTransactor{contract: contract}, ShuttleMachineFilterer: ShuttleMachineFilterer{contract: contract}}, nil
}

// NewShuttleMachineCaller creates a new read-only instance of ShuttleMachine, bound to a specific deployed contract.
func NewShuttleMachineCaller(address common.Address, caller bind.ContractCaller) (*ShuttleMachineCaller, error) {
	contract, err := bindShuttleMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShuttleMachineCaller{contract: contract}, nil
}

// NewShuttleMachineTransactor creates a new write-only instance of ShuttleMachine, bound to a specific deployed contract.
func NewShuttleMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*ShuttleMachineTransactor, error) {
	contract, err := bindShuttleMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShuttleMachineTransactor{contract: contract}, nil
}

// NewShuttleMachineFilterer creates a new log filterer instance of ShuttleMachine, bound to a specific deployed contract.
func NewShuttleMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*ShuttleMachineFilterer, error) {
	contract, err := bindShuttleMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShuttleMachineFilterer{contract: contract}, nil
}

// bindShuttleMachine binds a generic wrapper to an already deployed contract.
func bindShuttleMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShuttleMachineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShuttleMachine *ShuttleMachineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShuttleMachine.Contract.ShuttleMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShuttleMachine *ShuttleMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.ShuttleMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShuttleMachine *ShuttleMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.ShuttleMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShuttleMachine *ShuttleMachineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShuttleMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShuttleMachine *ShuttleMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShuttleMachine *ShuttleMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.contract.Transact(opts, method, params...)
}

// PlugmanAddress is a free data retrieval call binding the contract method 0x16d2207d.
//
// Solidity: function plugmanAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineCaller) PlugmanAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShuttleMachine.contract.Call(opts, &out, "plugmanAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlugmanAddress is a free data retrieval call binding the contract method 0x16d2207d.
//
// Solidity: function plugmanAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineSession) PlugmanAddress() (common.Address, error) {
	return _ShuttleMachine.Contract.PlugmanAddress(&_ShuttleMachine.CallOpts)
}

// PlugmanAddress is a free data retrieval call binding the contract method 0x16d2207d.
//
// Solidity: function plugmanAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineCallerSession) PlugmanAddress() (common.Address, error) {
	return _ShuttleMachine.Contract.PlugmanAddress(&_ShuttleMachine.CallOpts)
}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineCaller) PublicSalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShuttleMachine.contract.Call(opts, &out, "publicSalePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineSession) PublicSalePrice() (*big.Int, error) {
	return _ShuttleMachine.Contract.PublicSalePrice(&_ShuttleMachine.CallOpts)
}

// PublicSalePrice is a free data retrieval call binding the contract method 0x9b6860c8.
//
// Solidity: function publicSalePrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineCallerSession) PublicSalePrice() (*big.Int, error) {
	return _ShuttleMachine.Contract.PublicSalePrice(&_ShuttleMachine.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ShuttleMachine *ShuttleMachineCaller) SystemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShuttleMachine.contract.Call(opts, &out, "systemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ShuttleMachine *ShuttleMachineSession) SystemContract() (common.Address, error) {
	return _ShuttleMachine.Contract.SystemContract(&_ShuttleMachine.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ShuttleMachine *ShuttleMachineCallerSession) SystemContract() (common.Address, error) {
	return _ShuttleMachine.Contract.SystemContract(&_ShuttleMachine.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShuttleMachine.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineSession) TreasuryAddress() (common.Address, error) {
	return _ShuttleMachine.Contract.TreasuryAddress(&_ShuttleMachine.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_ShuttleMachine *ShuttleMachineCallerSession) TreasuryAddress() (common.Address, error) {
	return _ShuttleMachine.Contract.TreasuryAddress(&_ShuttleMachine.CallOpts)
}

// WlPrice is a free data retrieval call binding the contract method 0xc7f8d01a.
//
// Solidity: function wlPrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineCaller) WlPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShuttleMachine.contract.Call(opts, &out, "wlPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlPrice is a free data retrieval call binding the contract method 0xc7f8d01a.
//
// Solidity: function wlPrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineSession) WlPrice() (*big.Int, error) {
	return _ShuttleMachine.Contract.WlPrice(&_ShuttleMachine.CallOpts)
}

// WlPrice is a free data retrieval call binding the contract method 0xc7f8d01a.
//
// Solidity: function wlPrice() view returns(uint256)
func (_ShuttleMachine *ShuttleMachineCallerSession) WlPrice() (*big.Int, error) {
	return _ShuttleMachine.Contract.WlPrice(&_ShuttleMachine.CallOpts)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ShuttleMachine *ShuttleMachineTransactor) OnCrossChainCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ShuttleMachine.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ShuttleMachine *ShuttleMachineSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.OnCrossChainCall(&_ShuttleMachine.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ShuttleMachine *ShuttleMachineTransactorSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ShuttleMachine.Contract.OnCrossChainCall(&_ShuttleMachine.TransactOpts, context, zrc20, amount, message)
}

// ShuttleMachineCrossChainMintIterator is returned from FilterCrossChainMint and is used to iterate over the raw logs and unpacked data for CrossChainMint events raised by the ShuttleMachine contract.
type ShuttleMachineCrossChainMintIterator struct {
	Event *ShuttleMachineCrossChainMint // Event containing the contract specifics and raw log

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
func (it *ShuttleMachineCrossChainMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShuttleMachineCrossChainMint)
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
		it.Event = new(ShuttleMachineCrossChainMint)
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
func (it *ShuttleMachineCrossChainMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShuttleMachineCrossChainMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShuttleMachineCrossChainMint represents a CrossChainMint event raised by the ShuttleMachine contract.
type ShuttleMachineCrossChainMint struct {
	To       common.Address
	Sender   common.Address
	Count    *big.Int
	Nonce    *big.Int
	MintType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCrossChainMint is a free log retrieval operation binding the contract event 0xf01ca84cb640fe7629d97577c869a60819c89a5d8b46aed04d0473843a646351.
//
// Solidity: event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType)
func (_ShuttleMachine *ShuttleMachineFilterer) FilterCrossChainMint(opts *bind.FilterOpts, to []common.Address) (*ShuttleMachineCrossChainMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShuttleMachine.contract.FilterLogs(opts, "CrossChainMint", toRule)
	if err != nil {
		return nil, err
	}
	return &ShuttleMachineCrossChainMintIterator{contract: _ShuttleMachine.contract, event: "CrossChainMint", logs: logs, sub: sub}, nil
}

// WatchCrossChainMint is a free log subscription operation binding the contract event 0xf01ca84cb640fe7629d97577c869a60819c89a5d8b46aed04d0473843a646351.
//
// Solidity: event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType)
func (_ShuttleMachine *ShuttleMachineFilterer) WatchCrossChainMint(opts *bind.WatchOpts, sink chan<- *ShuttleMachineCrossChainMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShuttleMachine.contract.WatchLogs(opts, "CrossChainMint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShuttleMachineCrossChainMint)
				if err := _ShuttleMachine.contract.UnpackLog(event, "CrossChainMint", log); err != nil {
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

// ParseCrossChainMint is a log parse operation binding the contract event 0xf01ca84cb640fe7629d97577c869a60819c89a5d8b46aed04d0473843a646351.
//
// Solidity: event CrossChainMint(address indexed to, address sender, uint256 count, uint256 nonce, uint8 mintType)
func (_ShuttleMachine *ShuttleMachineFilterer) ParseCrossChainMint(log types.Log) (*ShuttleMachineCrossChainMint, error) {
	event := new(ShuttleMachineCrossChainMint)
	if err := _ShuttleMachine.contract.UnpackLog(event, "CrossChainMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
