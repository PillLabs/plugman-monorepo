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

// IERC20CustodyNewEventsMetaData contains all meta data concerning the IERC20CustodyNewEvents contract.
var IERC20CustodyNewEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"}]",
}

// IERC20CustodyNewEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyNewEventsMetaData.ABI instead.
var IERC20CustodyNewEventsABI = IERC20CustodyNewEventsMetaData.ABI

// IERC20CustodyNewEvents is an auto generated Go binding around an Ethereum contract.
type IERC20CustodyNewEvents struct {
	IERC20CustodyNewEventsCaller     // Read-only binding to the contract
	IERC20CustodyNewEventsTransactor // Write-only binding to the contract
	IERC20CustodyNewEventsFilterer   // Log filterer for contract events
}

// IERC20CustodyNewEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyNewEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyNewEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyNewEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyNewEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodyNewEventsSession struct {
	Contract     *IERC20CustodyNewEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20CustodyNewEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyNewEventsCallerSession struct {
	Contract *IERC20CustodyNewEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IERC20CustodyNewEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyNewEventsTransactorSession struct {
	Contract     *IERC20CustodyNewEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IERC20CustodyNewEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyNewEventsRaw struct {
	Contract *IERC20CustodyNewEvents // Generic contract binding to access the raw methods on
}

// IERC20CustodyNewEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyNewEventsCallerRaw struct {
	Contract *IERC20CustodyNewEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyNewEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyNewEventsTransactorRaw struct {
	Contract *IERC20CustodyNewEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20CustodyNewEvents creates a new instance of IERC20CustodyNewEvents, bound to a specific deployed contract.
func NewIERC20CustodyNewEvents(address common.Address, backend bind.ContractBackend) (*IERC20CustodyNewEvents, error) {
	contract, err := bindIERC20CustodyNewEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEvents{IERC20CustodyNewEventsCaller: IERC20CustodyNewEventsCaller{contract: contract}, IERC20CustodyNewEventsTransactor: IERC20CustodyNewEventsTransactor{contract: contract}, IERC20CustodyNewEventsFilterer: IERC20CustodyNewEventsFilterer{contract: contract}}, nil
}

// NewIERC20CustodyNewEventsCaller creates a new read-only instance of IERC20CustodyNewEvents, bound to a specific deployed contract.
func NewIERC20CustodyNewEventsCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyNewEventsCaller, error) {
	contract, err := bindIERC20CustodyNewEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsCaller{contract: contract}, nil
}

// NewIERC20CustodyNewEventsTransactor creates a new write-only instance of IERC20CustodyNewEvents, bound to a specific deployed contract.
func NewIERC20CustodyNewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyNewEventsTransactor, error) {
	contract, err := bindIERC20CustodyNewEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsTransactor{contract: contract}, nil
}

// NewIERC20CustodyNewEventsFilterer creates a new log filterer instance of IERC20CustodyNewEvents, bound to a specific deployed contract.
func NewIERC20CustodyNewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyNewEventsFilterer, error) {
	contract, err := bindIERC20CustodyNewEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsFilterer{contract: contract}, nil
}

// bindIERC20CustodyNewEvents binds a generic wrapper to an already deployed contract.
func bindIERC20CustodyNewEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyNewEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyNewEvents.Contract.IERC20CustodyNewEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyNewEvents.Contract.IERC20CustodyNewEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyNewEvents.Contract.IERC20CustodyNewEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyNewEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyNewEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyNewEvents.Contract.contract.Transact(opts, method, params...)
}

// IERC20CustodyNewEventsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdrawIterator struct {
	Event *IERC20CustodyNewEventsWithdraw // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyNewEventsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyNewEventsWithdraw)
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
		it.Event = new(IERC20CustodyNewEventsWithdraw)
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
func (it *IERC20CustodyNewEventsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyNewEventsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyNewEventsWithdraw represents a Withdraw event raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyNewEventsWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.FilterLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsWithdrawIterator{contract: _IERC20CustodyNewEvents.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IERC20CustodyNewEventsWithdraw, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.WatchLogs(opts, "Withdraw", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyNewEventsWithdraw)
				if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed token, address indexed to, uint256 amount)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) ParseWithdraw(log types.Log) (*IERC20CustodyNewEventsWithdraw, error) {
	event := new(IERC20CustodyNewEventsWithdraw)
	if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyNewEventsWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdrawAndCallIterator struct {
	Event *IERC20CustodyNewEventsWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyNewEventsWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyNewEventsWithdrawAndCall)
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
		it.Event = new(IERC20CustodyNewEventsWithdrawAndCall)
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
func (it *IERC20CustodyNewEventsWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyNewEventsWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyNewEventsWithdrawAndCall represents a WithdrawAndCall event raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdrawAndCall struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyNewEventsWithdrawAndCallIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.FilterLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsWithdrawAndCallIterator{contract: _IERC20CustodyNewEvents.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *IERC20CustodyNewEventsWithdrawAndCall, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.WatchLogs(opts, "WithdrawAndCall", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyNewEventsWithdrawAndCall)
				if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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

// ParseWithdrawAndCall is a log parse operation binding the contract event 0x85b5be9cf454e05e0bddf49315178102227c312078eefa3c00294fb4d912ae4e.
//
// Solidity: event WithdrawAndCall(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) ParseWithdrawAndCall(log types.Log) (*IERC20CustodyNewEventsWithdrawAndCall, error) {
	event := new(IERC20CustodyNewEventsWithdrawAndCall)
	if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyNewEventsWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdrawAndRevertIterator struct {
	Event *IERC20CustodyNewEventsWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyNewEventsWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyNewEventsWithdrawAndRevert)
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
		it.Event = new(IERC20CustodyNewEventsWithdrawAndRevert)
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
func (it *IERC20CustodyNewEventsWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyNewEventsWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyNewEventsWithdrawAndRevert represents a WithdrawAndRevert event raised by the IERC20CustodyNewEvents contract.
type IERC20CustodyNewEventsWithdrawAndRevert struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*IERC20CustodyNewEventsWithdrawAndRevertIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.FilterLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyNewEventsWithdrawAndRevertIterator{contract: _IERC20CustodyNewEvents.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *IERC20CustodyNewEventsWithdrawAndRevert, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20CustodyNewEvents.contract.WatchLogs(opts, "WithdrawAndRevert", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyNewEventsWithdrawAndRevert)
				if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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

// ParseWithdrawAndRevert is a log parse operation binding the contract event 0xb9d4efa96044e5f5e03e696fa9ae2ff66911cc27e8a637c3627c75bc5b2241c8.
//
// Solidity: event WithdrawAndRevert(address indexed token, address indexed to, uint256 amount, bytes data)
func (_IERC20CustodyNewEvents *IERC20CustodyNewEventsFilterer) ParseWithdrawAndRevert(log types.Log) (*IERC20CustodyNewEventsWithdrawAndRevert, error) {
	event := new(IERC20CustodyNewEventsWithdrawAndRevert)
	if err := _IERC20CustodyNewEvents.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
