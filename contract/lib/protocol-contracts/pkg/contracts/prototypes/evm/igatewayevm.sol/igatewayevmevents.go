// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayevm

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

// IGatewayEVMEventsMetaData contains all meta data concerning the IGatewayEVMEvents contract.
var IGatewayEVMEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedWithERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RevertedWithERC20\",\"type\":\"event\"}]",
}

// IGatewayEVMEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayEVMEventsMetaData.ABI instead.
var IGatewayEVMEventsABI = IGatewayEVMEventsMetaData.ABI

// IGatewayEVMEvents is an auto generated Go binding around an Ethereum contract.
type IGatewayEVMEvents struct {
	IGatewayEVMEventsCaller     // Read-only binding to the contract
	IGatewayEVMEventsTransactor // Write-only binding to the contract
	IGatewayEVMEventsFilterer   // Log filterer for contract events
}

// IGatewayEVMEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayEVMEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayEVMEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayEVMEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayEVMEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayEVMEventsSession struct {
	Contract     *IGatewayEVMEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IGatewayEVMEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayEVMEventsCallerSession struct {
	Contract *IGatewayEVMEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IGatewayEVMEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayEVMEventsTransactorSession struct {
	Contract     *IGatewayEVMEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IGatewayEVMEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayEVMEventsRaw struct {
	Contract *IGatewayEVMEvents // Generic contract binding to access the raw methods on
}

// IGatewayEVMEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayEVMEventsCallerRaw struct {
	Contract *IGatewayEVMEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayEVMEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayEVMEventsTransactorRaw struct {
	Contract *IGatewayEVMEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayEVMEvents creates a new instance of IGatewayEVMEvents, bound to a specific deployed contract.
func NewIGatewayEVMEvents(address common.Address, backend bind.ContractBackend) (*IGatewayEVMEvents, error) {
	contract, err := bindIGatewayEVMEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEvents{IGatewayEVMEventsCaller: IGatewayEVMEventsCaller{contract: contract}, IGatewayEVMEventsTransactor: IGatewayEVMEventsTransactor{contract: contract}, IGatewayEVMEventsFilterer: IGatewayEVMEventsFilterer{contract: contract}}, nil
}

// NewIGatewayEVMEventsCaller creates a new read-only instance of IGatewayEVMEvents, bound to a specific deployed contract.
func NewIGatewayEVMEventsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayEVMEventsCaller, error) {
	contract, err := bindIGatewayEVMEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsCaller{contract: contract}, nil
}

// NewIGatewayEVMEventsTransactor creates a new write-only instance of IGatewayEVMEvents, bound to a specific deployed contract.
func NewIGatewayEVMEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayEVMEventsTransactor, error) {
	contract, err := bindIGatewayEVMEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsTransactor{contract: contract}, nil
}

// NewIGatewayEVMEventsFilterer creates a new log filterer instance of IGatewayEVMEvents, bound to a specific deployed contract.
func NewIGatewayEVMEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayEVMEventsFilterer, error) {
	contract, err := bindIGatewayEVMEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsFilterer{contract: contract}, nil
}

// bindIGatewayEVMEvents binds a generic wrapper to an already deployed contract.
func bindIGatewayEVMEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayEVMEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVMEvents *IGatewayEVMEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVMEvents.Contract.IGatewayEVMEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVMEvents *IGatewayEVMEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVMEvents.Contract.IGatewayEVMEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVMEvents *IGatewayEVMEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVMEvents.Contract.IGatewayEVMEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayEVMEvents *IGatewayEVMEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayEVMEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayEVMEvents *IGatewayEVMEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayEVMEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayEVMEvents *IGatewayEVMEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayEVMEvents.Contract.contract.Transact(opts, method, params...)
}

// IGatewayEVMEventsCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsCallIterator struct {
	Event *IGatewayEVMEventsCall // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsCall)
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
		it.Event = new(IGatewayEVMEventsCall)
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
func (it *IGatewayEVMEventsCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsCall represents a Call event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsCall struct {
	Sender   common.Address
	Receiver common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*IGatewayEVMEventsCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsCallIterator{contract: _IGatewayEVMEvents.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsCall, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "Call", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsCall)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2a21062ee9199c2e205622999eeb7c3da73153674f36a0acd3f74fa6af67bde3.
//
// Solidity: event Call(address indexed sender, address indexed receiver, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseCall(log types.Log) (*IGatewayEVMEventsCall, error) {
	event := new(IGatewayEVMEventsCall)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMEventsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsDepositIterator struct {
	Event *IGatewayEVMEventsDeposit // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsDeposit)
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
		it.Event = new(IGatewayEVMEventsDeposit)
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
func (it *IGatewayEVMEventsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsDeposit represents a Deposit event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsDeposit struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Asset    common.Address
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*IGatewayEVMEventsDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsDepositIterator{contract: _IGatewayEVMEvents.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsDeposit, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "Deposit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsDeposit)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2103daedac6c1eee9e5bfbd02064d751c9ec3c03fb9bc3e4f94ca41afa38c1a4.
//
// Solidity: event Deposit(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseDeposit(log types.Log) (*IGatewayEVMEventsDeposit, error) {
	event := new(IGatewayEVMEventsDeposit)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMEventsExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsExecutedIterator struct {
	Event *IGatewayEVMEventsExecuted // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsExecuted)
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
		it.Event = new(IGatewayEVMEventsExecuted)
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
func (it *IGatewayEVMEventsExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsExecuted represents a Executed event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*IGatewayEVMEventsExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsExecutedIterator{contract: _IGatewayEVMEvents.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsExecuted)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseExecuted(log types.Log) (*IGatewayEVMEventsExecuted, error) {
	event := new(IGatewayEVMEventsExecuted)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMEventsExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsExecutedWithERC20Iterator struct {
	Event *IGatewayEVMEventsExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsExecutedWithERC20)
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
		it.Event = new(IGatewayEVMEventsExecutedWithERC20)
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
func (it *IGatewayEVMEventsExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsExecutedWithERC20 represents a ExecutedWithERC20 event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IGatewayEVMEventsExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsExecutedWithERC20Iterator{contract: _IGatewayEVMEvents.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsExecutedWithERC20)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseExecutedWithERC20(log types.Log) (*IGatewayEVMEventsExecutedWithERC20, error) {
	event := new(IGatewayEVMEventsExecutedWithERC20)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMEventsRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsRevertedIterator struct {
	Event *IGatewayEVMEventsReverted // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsReverted)
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
		it.Event = new(IGatewayEVMEventsReverted)
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
func (it *IGatewayEVMEventsRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsReverted represents a Reverted event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsReverted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterReverted(opts *bind.FilterOpts, destination []common.Address) (*IGatewayEVMEventsRevertedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsRevertedIterator{contract: _IGatewayEVMEvents.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsReverted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "Reverted", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsReverted)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xd5d7616b1678354a0dea9d7e57e6a090bff5babe9f8d6381fdbad16e89ba311c.
//
// Solidity: event Reverted(address indexed destination, uint256 value, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseReverted(log types.Log) (*IGatewayEVMEventsReverted, error) {
	event := new(IGatewayEVMEventsReverted)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayEVMEventsRevertedWithERC20Iterator is returned from FilterRevertedWithERC20 and is used to iterate over the raw logs and unpacked data for RevertedWithERC20 events raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsRevertedWithERC20Iterator struct {
	Event *IGatewayEVMEventsRevertedWithERC20 // Event containing the contract specifics and raw log

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
func (it *IGatewayEVMEventsRevertedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayEVMEventsRevertedWithERC20)
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
		it.Event = new(IGatewayEVMEventsRevertedWithERC20)
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
func (it *IGatewayEVMEventsRevertedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayEVMEventsRevertedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayEVMEventsRevertedWithERC20 represents a RevertedWithERC20 event raised by the IGatewayEVMEvents contract.
type IGatewayEVMEventsRevertedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevertedWithERC20 is a free log retrieval operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) FilterRevertedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IGatewayEVMEventsRevertedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.FilterLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayEVMEventsRevertedWithERC20Iterator{contract: _IGatewayEVMEvents.contract, event: "RevertedWithERC20", logs: logs, sub: sub}, nil
}

// WatchRevertedWithERC20 is a free log subscription operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) WatchRevertedWithERC20(opts *bind.WatchOpts, sink chan<- *IGatewayEVMEventsRevertedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IGatewayEVMEvents.contract.WatchLogs(opts, "RevertedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayEVMEventsRevertedWithERC20)
				if err := _IGatewayEVMEvents.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
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

// ParseRevertedWithERC20 is a log parse operation binding the contract event 0x723fc7be2448075379e4fdf1e6bf5fead954d2668d2da05dcb44ccfec4beeda7.
//
// Solidity: event RevertedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IGatewayEVMEvents *IGatewayEVMEventsFilterer) ParseRevertedWithERC20(log types.Log) (*IGatewayEVMEventsRevertedWithERC20, error) {
	event := new(IGatewayEVMEventsRevertedWithERC20)
	if err := _IGatewayEVMEvents.contract.UnpackLog(event, "RevertedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
