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

// IGatewayZEVMEventsMetaData contains all meta data concerning the IGatewayZEVMEvents contract.
var IGatewayZEVMEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasfee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Withdrawal\",\"type\":\"event\"}]",
}

// IGatewayZEVMEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayZEVMEventsMetaData.ABI instead.
var IGatewayZEVMEventsABI = IGatewayZEVMEventsMetaData.ABI

// IGatewayZEVMEvents is an auto generated Go binding around an Ethereum contract.
type IGatewayZEVMEvents struct {
	IGatewayZEVMEventsCaller     // Read-only binding to the contract
	IGatewayZEVMEventsTransactor // Write-only binding to the contract
	IGatewayZEVMEventsFilterer   // Log filterer for contract events
}

// IGatewayZEVMEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayZEVMEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayZEVMEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayZEVMEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayZEVMEventsSession struct {
	Contract     *IGatewayZEVMEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGatewayZEVMEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayZEVMEventsCallerSession struct {
	Contract *IGatewayZEVMEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IGatewayZEVMEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayZEVMEventsTransactorSession struct {
	Contract     *IGatewayZEVMEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IGatewayZEVMEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayZEVMEventsRaw struct {
	Contract *IGatewayZEVMEvents // Generic contract binding to access the raw methods on
}

// IGatewayZEVMEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayZEVMEventsCallerRaw struct {
	Contract *IGatewayZEVMEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayZEVMEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayZEVMEventsTransactorRaw struct {
	Contract *IGatewayZEVMEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayZEVMEvents creates a new instance of IGatewayZEVMEvents, bound to a specific deployed contract.
func NewIGatewayZEVMEvents(address common.Address, backend bind.ContractBackend) (*IGatewayZEVMEvents, error) {
	contract, err := bindIGatewayZEVMEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEvents{IGatewayZEVMEventsCaller: IGatewayZEVMEventsCaller{contract: contract}, IGatewayZEVMEventsTransactor: IGatewayZEVMEventsTransactor{contract: contract}, IGatewayZEVMEventsFilterer: IGatewayZEVMEventsFilterer{contract: contract}}, nil
}

// NewIGatewayZEVMEventsCaller creates a new read-only instance of IGatewayZEVMEvents, bound to a specific deployed contract.
func NewIGatewayZEVMEventsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayZEVMEventsCaller, error) {
	contract, err := bindIGatewayZEVMEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEventsCaller{contract: contract}, nil
}

// NewIGatewayZEVMEventsTransactor creates a new write-only instance of IGatewayZEVMEvents, bound to a specific deployed contract.
func NewIGatewayZEVMEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayZEVMEventsTransactor, error) {
	contract, err := bindIGatewayZEVMEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEventsTransactor{contract: contract}, nil
}

// NewIGatewayZEVMEventsFilterer creates a new log filterer instance of IGatewayZEVMEvents, bound to a specific deployed contract.
func NewIGatewayZEVMEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayZEVMEventsFilterer, error) {
	contract, err := bindIGatewayZEVMEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEventsFilterer{contract: contract}, nil
}

// bindIGatewayZEVMEvents binds a generic wrapper to an already deployed contract.
func bindIGatewayZEVMEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayZEVMEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVMEvents.Contract.IGatewayZEVMEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVMEvents.Contract.IGatewayZEVMEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVMEvents.Contract.IGatewayZEVMEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVMEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVMEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVMEvents *IGatewayZEVMEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVMEvents.Contract.contract.Transact(opts, method, params...)
}

// IGatewayZEVMEventsCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the IGatewayZEVMEvents contract.
type IGatewayZEVMEventsCallIterator struct {
	Event *IGatewayZEVMEventsCall // Event containing the contract specifics and raw log

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
func (it *IGatewayZEVMEventsCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayZEVMEventsCall)
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
		it.Event = new(IGatewayZEVMEventsCall)
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
func (it *IGatewayZEVMEventsCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayZEVMEventsCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayZEVMEventsCall represents a Call event raised by the IGatewayZEVMEvents contract.
type IGatewayZEVMEventsCall struct {
	Sender   common.Address
	Receiver []byte
	Message  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) FilterCall(opts *bind.FilterOpts, sender []common.Address) (*IGatewayZEVMEventsCallIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IGatewayZEVMEvents.contract.FilterLogs(opts, "Call", senderRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEventsCallIterator{contract: _IGatewayZEVMEvents.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *IGatewayZEVMEventsCall, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IGatewayZEVMEvents.contract.WatchLogs(opts, "Call", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayZEVMEventsCall)
				if err := _IGatewayZEVMEvents.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x2b5af078ce280d812dc2241658dc5435c93408020e5418eef55a2b536de51c0f.
//
// Solidity: event Call(address indexed sender, bytes receiver, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) ParseCall(log types.Log) (*IGatewayZEVMEventsCall, error) {
	event := new(IGatewayZEVMEventsCall)
	if err := _IGatewayZEVMEvents.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayZEVMEventsWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the IGatewayZEVMEvents contract.
type IGatewayZEVMEventsWithdrawalIterator struct {
	Event *IGatewayZEVMEventsWithdrawal // Event containing the contract specifics and raw log

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
func (it *IGatewayZEVMEventsWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayZEVMEventsWithdrawal)
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
		it.Event = new(IGatewayZEVMEventsWithdrawal)
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
func (it *IGatewayZEVMEventsWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayZEVMEventsWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayZEVMEventsWithdrawal represents a Withdrawal event raised by the IGatewayZEVMEvents contract.
type IGatewayZEVMEventsWithdrawal struct {
	From            common.Address
	Zrc20           common.Address
	To              []byte
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*IGatewayZEVMEventsWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGatewayZEVMEvents.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMEventsWithdrawalIterator{contract: _IGatewayZEVMEvents.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *IGatewayZEVMEventsWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGatewayZEVMEvents.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayZEVMEventsWithdrawal)
				if err := _IGatewayZEVMEvents.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x2265ce9ec38ea098a1143406678482665a6e1ccd82ab22d37eea3a78abc57716.
//
// Solidity: event Withdrawal(address indexed from, address zrc20, bytes to, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message)
func (_IGatewayZEVMEvents *IGatewayZEVMEventsFilterer) ParseWithdrawal(log types.Log) (*IGatewayZEVMEventsWithdrawal, error) {
	event := new(IGatewayZEVMEventsWithdrawal)
	if err := _IGatewayZEVMEvents.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
