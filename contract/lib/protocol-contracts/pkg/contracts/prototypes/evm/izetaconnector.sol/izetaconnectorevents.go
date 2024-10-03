// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetaconnector

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

// IZetaConnectorEventsMetaData contains all meta data concerning the IZetaConnectorEvents contract.
var IZetaConnectorEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"}]",
}

// IZetaConnectorEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaConnectorEventsMetaData.ABI instead.
var IZetaConnectorEventsABI = IZetaConnectorEventsMetaData.ABI

// IZetaConnectorEvents is an auto generated Go binding around an Ethereum contract.
type IZetaConnectorEvents struct {
	IZetaConnectorEventsCaller     // Read-only binding to the contract
	IZetaConnectorEventsTransactor // Write-only binding to the contract
	IZetaConnectorEventsFilterer   // Log filterer for contract events
}

// IZetaConnectorEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaConnectorEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaConnectorEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaConnectorEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaConnectorEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaConnectorEventsSession struct {
	Contract     *IZetaConnectorEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IZetaConnectorEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaConnectorEventsCallerSession struct {
	Contract *IZetaConnectorEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IZetaConnectorEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaConnectorEventsTransactorSession struct {
	Contract     *IZetaConnectorEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IZetaConnectorEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaConnectorEventsRaw struct {
	Contract *IZetaConnectorEvents // Generic contract binding to access the raw methods on
}

// IZetaConnectorEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaConnectorEventsCallerRaw struct {
	Contract *IZetaConnectorEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaConnectorEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaConnectorEventsTransactorRaw struct {
	Contract *IZetaConnectorEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaConnectorEvents creates a new instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEvents(address common.Address, backend bind.ContractBackend) (*IZetaConnectorEvents, error) {
	contract, err := bindIZetaConnectorEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEvents{IZetaConnectorEventsCaller: IZetaConnectorEventsCaller{contract: contract}, IZetaConnectorEventsTransactor: IZetaConnectorEventsTransactor{contract: contract}, IZetaConnectorEventsFilterer: IZetaConnectorEventsFilterer{contract: contract}}, nil
}

// NewIZetaConnectorEventsCaller creates a new read-only instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsCaller(address common.Address, caller bind.ContractCaller) (*IZetaConnectorEventsCaller, error) {
	contract, err := bindIZetaConnectorEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsCaller{contract: contract}, nil
}

// NewIZetaConnectorEventsTransactor creates a new write-only instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaConnectorEventsTransactor, error) {
	contract, err := bindIZetaConnectorEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsTransactor{contract: contract}, nil
}

// NewIZetaConnectorEventsFilterer creates a new log filterer instance of IZetaConnectorEvents, bound to a specific deployed contract.
func NewIZetaConnectorEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaConnectorEventsFilterer, error) {
	contract, err := bindIZetaConnectorEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsFilterer{contract: contract}, nil
}

// bindIZetaConnectorEvents binds a generic wrapper to an already deployed contract.
func bindIZetaConnectorEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaConnectorEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaConnectorEvents *IZetaConnectorEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.IZetaConnectorEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaConnectorEvents *IZetaConnectorEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaConnectorEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaConnectorEvents *IZetaConnectorEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaConnectorEvents *IZetaConnectorEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaConnectorEvents.Contract.contract.Transact(opts, method, params...)
}

// IZetaConnectorEventsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawIterator struct {
	Event *IZetaConnectorEventsWithdraw // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdraw)
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
		it.Event = new(IZetaConnectorEventsWithdraw)
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
func (it *IZetaConnectorEventsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdraw represents a Withdraw event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawIterator{contract: _IZetaConnectorEvents.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdraw)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdraw(log types.Log) (*IZetaConnectorEventsWithdraw, error) {
	event := new(IZetaConnectorEventsWithdraw)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaConnectorEventsWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawAndCallIterator struct {
	Event *IZetaConnectorEventsWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdrawAndCall)
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
		it.Event = new(IZetaConnectorEventsWithdrawAndCall)
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
func (it *IZetaConnectorEventsWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdrawAndCall represents a WithdrawAndCall event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawAndCallIterator{contract: _IZetaConnectorEvents.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdrawAndCall)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdrawAndCall(log types.Log) (*IZetaConnectorEventsWithdrawAndCall, error) {
	event := new(IZetaConnectorEventsWithdrawAndCall)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaConnectorEventsWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawAndRevertIterator struct {
	Event *IZetaConnectorEventsWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *IZetaConnectorEventsWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaConnectorEventsWithdrawAndRevert)
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
		it.Event = new(IZetaConnectorEventsWithdrawAndRevert)
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
func (it *IZetaConnectorEventsWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaConnectorEventsWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaConnectorEventsWithdrawAndRevert represents a WithdrawAndRevert event raised by the IZetaConnectorEvents contract.
type IZetaConnectorEventsWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*IZetaConnectorEventsWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaConnectorEventsWithdrawAndRevertIterator{contract: _IZetaConnectorEvents.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *IZetaConnectorEventsWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaConnectorEvents.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaConnectorEventsWithdrawAndRevert)
				if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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
func (_IZetaConnectorEvents *IZetaConnectorEventsFilterer) ParseWithdrawAndRevert(log types.Log) (*IZetaConnectorEventsWithdrawAndRevert, error) {
	event := new(IZetaConnectorEventsWithdrawAndRevert)
	if err := _IZetaConnectorEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
