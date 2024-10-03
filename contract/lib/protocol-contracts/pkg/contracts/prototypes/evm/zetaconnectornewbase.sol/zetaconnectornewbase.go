// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornewbase

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

// ZetaConnectorNewBaseMetaData contains all meta data concerning the ZetaConnectorNewBase contract.
var ZetaConnectorNewBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawAndRevert\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"contractIGatewayEVM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"internalSendHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZetaConnectorNewBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNewBaseMetaData.ABI instead.
var ZetaConnectorNewBaseABI = ZetaConnectorNewBaseMetaData.ABI

// ZetaConnectorNewBase is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNewBase struct {
	ZetaConnectorNewBaseCaller     // Read-only binding to the contract
	ZetaConnectorNewBaseTransactor // Write-only binding to the contract
	ZetaConnectorNewBaseFilterer   // Log filterer for contract events
}

// ZetaConnectorNewBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNewBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNewBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNewBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNewBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNewBaseSession struct {
	Contract     *ZetaConnectorNewBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZetaConnectorNewBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNewBaseCallerSession struct {
	Contract *ZetaConnectorNewBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ZetaConnectorNewBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNewBaseTransactorSession struct {
	Contract     *ZetaConnectorNewBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZetaConnectorNewBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNewBaseRaw struct {
	Contract *ZetaConnectorNewBase // Generic contract binding to access the raw methods on
}

// ZetaConnectorNewBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNewBaseCallerRaw struct {
	Contract *ZetaConnectorNewBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNewBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNewBaseTransactorRaw struct {
	Contract *ZetaConnectorNewBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNewBase creates a new instance of ZetaConnectorNewBase, bound to a specific deployed contract.
func NewZetaConnectorNewBase(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNewBase, error) {
	contract, err := bindZetaConnectorNewBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBase{ZetaConnectorNewBaseCaller: ZetaConnectorNewBaseCaller{contract: contract}, ZetaConnectorNewBaseTransactor: ZetaConnectorNewBaseTransactor{contract: contract}, ZetaConnectorNewBaseFilterer: ZetaConnectorNewBaseFilterer{contract: contract}}, nil
}

// NewZetaConnectorNewBaseCaller creates a new read-only instance of ZetaConnectorNewBase, bound to a specific deployed contract.
func NewZetaConnectorNewBaseCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNewBaseCaller, error) {
	contract, err := bindZetaConnectorNewBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseCaller{contract: contract}, nil
}

// NewZetaConnectorNewBaseTransactor creates a new write-only instance of ZetaConnectorNewBase, bound to a specific deployed contract.
func NewZetaConnectorNewBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNewBaseTransactor, error) {
	contract, err := bindZetaConnectorNewBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseTransactor{contract: contract}, nil
}

// NewZetaConnectorNewBaseFilterer creates a new log filterer instance of ZetaConnectorNewBase, bound to a specific deployed contract.
func NewZetaConnectorNewBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNewBaseFilterer, error) {
	contract, err := bindZetaConnectorNewBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseFilterer{contract: contract}, nil
}

// bindZetaConnectorNewBase binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNewBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNewBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewBase.Contract.ZetaConnectorNewBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.ZetaConnectorNewBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.ZetaConnectorNewBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNewBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewBase.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.Gateway(&_ZetaConnectorNewBase.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.Gateway(&_ZetaConnectorNewBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewBase.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.TssAddress(&_ZetaConnectorNewBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.TssAddress(&_ZetaConnectorNewBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNewBase.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.ZetaToken(&_ZetaConnectorNewBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNewBase.Contract.ZetaToken(&_ZetaConnectorNewBase.CallOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.ReceiveTokens(&_ZetaConnectorNewBase.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.ReceiveTokens(&_ZetaConnectorNewBase.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.Withdraw(&_ZetaConnectorNewBase.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.Withdraw(&_ZetaConnectorNewBase.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactor) WithdrawAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.contract.Transact(opts, "withdrawAndCall", to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.WithdrawAndCall(&_ZetaConnectorNewBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x5e3e9fef.
//
// Solidity: function withdrawAndCall(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorSession) WithdrawAndCall(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.WithdrawAndCall(&_ZetaConnectorNewBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.WithdrawAndRevert(&_ZetaConnectorNewBase.TransactOpts, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x02d5c899.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNewBase.Contract.WithdrawAndRevert(&_ZetaConnectorNewBase.TransactOpts, to, amount, data, internalSendHash)
}

// ZetaConnectorNewBaseWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdrawIterator struct {
	Event *ZetaConnectorNewBaseWithdraw // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewBaseWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewBaseWithdraw)
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
		it.Event = new(ZetaConnectorNewBaseWithdraw)
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
func (it *ZetaConnectorNewBaseWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewBaseWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewBaseWithdraw represents a Withdraw event raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewBaseWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseWithdrawIterator{contract: _ZetaConnectorNewBase.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewBaseWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewBaseWithdraw)
				if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) ParseWithdraw(log types.Log) (*ZetaConnectorNewBaseWithdraw, error) {
	event := new(ZetaConnectorNewBaseWithdraw)
	if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNewBaseWithdrawAndCallIterator is returned from FilterWithdrawAndCall and is used to iterate over the raw logs and unpacked data for WithdrawAndCall events raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdrawAndCallIterator struct {
	Event *ZetaConnectorNewBaseWithdrawAndCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewBaseWithdrawAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewBaseWithdrawAndCall)
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
		it.Event = new(ZetaConnectorNewBaseWithdrawAndCall)
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
func (it *ZetaConnectorNewBaseWithdrawAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewBaseWithdrawAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewBaseWithdrawAndCall represents a WithdrawAndCall event raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdrawAndCall struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCall is a free log retrieval operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) FilterWithdrawAndCall(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewBaseWithdrawAndCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.FilterLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseWithdrawAndCallIterator{contract: _ZetaConnectorNewBase.contract, event: "WithdrawAndCall", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCall is a free log subscription operation binding the contract event 0x7772f56296d3a5202974a45c61c9188d844ab4d6eeb18c851e4b8d5384ca6ced.
//
// Solidity: event WithdrawAndCall(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) WatchWithdrawAndCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewBaseWithdrawAndCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.WatchLogs(opts, "WithdrawAndCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewBaseWithdrawAndCall)
				if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
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
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) ParseWithdrawAndCall(log types.Log) (*ZetaConnectorNewBaseWithdrawAndCall, error) {
	event := new(ZetaConnectorNewBaseWithdrawAndCall)
	if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "WithdrawAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNewBaseWithdrawAndRevertIterator is returned from FilterWithdrawAndRevert and is used to iterate over the raw logs and unpacked data for WithdrawAndRevert events raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdrawAndRevertIterator struct {
	Event *ZetaConnectorNewBaseWithdrawAndRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNewBaseWithdrawAndRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNewBaseWithdrawAndRevert)
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
		it.Event = new(ZetaConnectorNewBaseWithdrawAndRevert)
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
func (it *ZetaConnectorNewBaseWithdrawAndRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNewBaseWithdrawAndRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNewBaseWithdrawAndRevert represents a WithdrawAndRevert event raised by the ZetaConnectorNewBase contract.
type ZetaConnectorNewBaseWithdrawAndRevert struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndRevert is a free log retrieval operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) FilterWithdrawAndRevert(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNewBaseWithdrawAndRevertIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.FilterLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNewBaseWithdrawAndRevertIterator{contract: _ZetaConnectorNewBase.contract, event: "WithdrawAndRevert", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndRevert is a free log subscription operation binding the contract event 0xba96f26bdda53eb8c8ba39045dfb4ff39753fbc7a6edcf250a88e75e78d102fe.
//
// Solidity: event WithdrawAndRevert(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) WatchWithdrawAndRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNewBaseWithdrawAndRevert, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNewBase.contract.WatchLogs(opts, "WithdrawAndRevert", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNewBaseWithdrawAndRevert)
				if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
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
func (_ZetaConnectorNewBase *ZetaConnectorNewBaseFilterer) ParseWithdrawAndRevert(log types.Log) (*ZetaConnectorNewBaseWithdrawAndRevert, error) {
	event := new(ZetaConnectorNewBaseWithdrawAndRevert)
	if err := _ZetaConnectorNewBase.contract.UnpackLog(event, "WithdrawAndRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
