// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc1967upgradeable

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

// IERC1967UpgradeableMetaData contains all meta data concerning the IERC1967Upgradeable contract.
var IERC1967UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]",
}

// IERC1967UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1967UpgradeableMetaData.ABI instead.
var IERC1967UpgradeableABI = IERC1967UpgradeableMetaData.ABI

// IERC1967Upgradeable is an auto generated Go binding around an Ethereum contract.
type IERC1967Upgradeable struct {
	IERC1967UpgradeableCaller     // Read-only binding to the contract
	IERC1967UpgradeableTransactor // Write-only binding to the contract
	IERC1967UpgradeableFilterer   // Log filterer for contract events
}

// IERC1967UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1967UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1967UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1967UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1967UpgradeableSession struct {
	Contract     *IERC1967Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC1967UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1967UpgradeableCallerSession struct {
	Contract *IERC1967UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC1967UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1967UpgradeableTransactorSession struct {
	Contract     *IERC1967UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC1967UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1967UpgradeableRaw struct {
	Contract *IERC1967Upgradeable // Generic contract binding to access the raw methods on
}

// IERC1967UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1967UpgradeableCallerRaw struct {
	Contract *IERC1967UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1967UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1967UpgradeableTransactorRaw struct {
	Contract *IERC1967UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1967Upgradeable creates a new instance of IERC1967Upgradeable, bound to a specific deployed contract.
func NewIERC1967Upgradeable(address common.Address, backend bind.ContractBackend) (*IERC1967Upgradeable, error) {
	contract, err := bindIERC1967Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1967Upgradeable{IERC1967UpgradeableCaller: IERC1967UpgradeableCaller{contract: contract}, IERC1967UpgradeableTransactor: IERC1967UpgradeableTransactor{contract: contract}, IERC1967UpgradeableFilterer: IERC1967UpgradeableFilterer{contract: contract}}, nil
}

// NewIERC1967UpgradeableCaller creates a new read-only instance of IERC1967Upgradeable, bound to a specific deployed contract.
func NewIERC1967UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC1967UpgradeableCaller, error) {
	contract, err := bindIERC1967Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableCaller{contract: contract}, nil
}

// NewIERC1967UpgradeableTransactor creates a new write-only instance of IERC1967Upgradeable, bound to a specific deployed contract.
func NewIERC1967UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1967UpgradeableTransactor, error) {
	contract, err := bindIERC1967Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableTransactor{contract: contract}, nil
}

// NewIERC1967UpgradeableFilterer creates a new log filterer instance of IERC1967Upgradeable, bound to a specific deployed contract.
func NewIERC1967UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1967UpgradeableFilterer, error) {
	contract, err := bindIERC1967Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableFilterer{contract: contract}, nil
}

// bindIERC1967Upgradeable binds a generic wrapper to an already deployed contract.
func bindIERC1967Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1967UpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1967Upgradeable *IERC1967UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1967Upgradeable.Contract.IERC1967UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1967Upgradeable *IERC1967UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1967Upgradeable.Contract.IERC1967UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1967Upgradeable *IERC1967UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1967Upgradeable.Contract.IERC1967UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1967Upgradeable *IERC1967UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1967Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1967Upgradeable *IERC1967UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1967Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1967Upgradeable *IERC1967UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1967Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// IERC1967UpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableAdminChangedIterator struct {
	Event *IERC1967UpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *IERC1967UpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967UpgradeableAdminChanged)
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
		it.Event = new(IERC1967UpgradeableAdminChanged)
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
func (it *IERC1967UpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967UpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967UpgradeableAdminChanged represents a AdminChanged event raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*IERC1967UpgradeableAdminChangedIterator, error) {

	logs, sub, err := _IERC1967Upgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableAdminChangedIterator{contract: _IERC1967Upgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *IERC1967UpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _IERC1967Upgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967UpgradeableAdminChanged)
				if err := _IERC1967Upgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) ParseAdminChanged(log types.Log) (*IERC1967UpgradeableAdminChanged, error) {
	event := new(IERC1967UpgradeableAdminChanged)
	if err := _IERC1967Upgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1967UpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableBeaconUpgradedIterator struct {
	Event *IERC1967UpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *IERC1967UpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967UpgradeableBeaconUpgraded)
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
		it.Event = new(IERC1967UpgradeableBeaconUpgraded)
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
func (it *IERC1967UpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967UpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967UpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*IERC1967UpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IERC1967Upgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableBeaconUpgradedIterator{contract: _IERC1967Upgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *IERC1967UpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IERC1967Upgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967UpgradeableBeaconUpgraded)
				if err := _IERC1967Upgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*IERC1967UpgradeableBeaconUpgraded, error) {
	event := new(IERC1967UpgradeableBeaconUpgraded)
	if err := _IERC1967Upgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1967UpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableUpgradedIterator struct {
	Event *IERC1967UpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *IERC1967UpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967UpgradeableUpgraded)
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
		it.Event = new(IERC1967UpgradeableUpgraded)
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
func (it *IERC1967UpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967UpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967UpgradeableUpgraded represents a Upgraded event raised by the IERC1967Upgradeable contract.
type IERC1967UpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IERC1967UpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IERC1967Upgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradeableUpgradedIterator{contract: _IERC1967Upgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IERC1967UpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IERC1967Upgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967UpgradeableUpgraded)
				if err := _IERC1967Upgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IERC1967Upgradeable *IERC1967UpgradeableFilterer) ParseUpgraded(log types.Log) (*IERC1967UpgradeableUpgraded, error) {
	event := new(IERC1967UpgradeableUpgraded)
	if err := _IERC1967Upgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
