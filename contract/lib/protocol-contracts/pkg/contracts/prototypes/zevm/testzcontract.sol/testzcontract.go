// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testzcontract

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// TestZContractMetaData contains all meta data concerning the TestZContract contract.
var TestZContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ContextData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ContextDataRevert\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structrevertContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506107e0806100206000396000f3fe60806040526004361061002d5760003560e01c806369582bee14610036578063de43156e1461005f57610034565b3661003457005b005b34801561004257600080fd5b5061005d60048036038101906100589190610346565b610088565b005b34801561006b57600080fd5b50610086600480360381019061008191906103ea565b610115565b005b606060008383905011156100a85782828101906100a591906102fd565b90505b7ffdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e999488680600001906100d99190610575565b8860200160208101906100ec91906102d0565b8960400135338660405161010596959493929190610512565b60405180910390a1505050505050565b6060600083839050111561013557828281019061013291906102fd565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e8680600001906101669190610575565b88602001602081019061017991906102d0565b8960400135338660405161019296959493929190610512565b60405180910390a1505050505050565b60006101b56101b0846105fd565b6105d8565b9050828152602081018484840111156101d1576101d061075c565b5b6101dc848285610697565b509392505050565b6000813590506101f38161077c565b92915050565b60008083601f84011261020f5761020e61073e565b5b8235905067ffffffffffffffff81111561022c5761022b610739565b5b60208301915083600182028301111561024857610247610752565b5b9250929050565b600082601f8301126102645761026361073e565b5b81356102748482602086016101a2565b91505092915050565b60006060828403121561029357610292610748565b5b81905092915050565b6000606082840312156102b2576102b1610748565b5b81905092915050565b6000813590506102ca81610793565b92915050565b6000602082840312156102e6576102e5610766565b5b60006102f4848285016101e4565b91505092915050565b60006020828403121561031357610312610766565b5b600082013567ffffffffffffffff81111561033157610330610761565b5b61033d8482850161024f565b91505092915050565b60008060008060006080868803121561036257610361610766565b5b600086013567ffffffffffffffff8111156103805761037f610761565b5b61038c8882890161027d565b955050602061039d888289016101e4565b94505060406103ae888289016102bb565b935050606086013567ffffffffffffffff8111156103cf576103ce610761565b5b6103db888289016101f9565b92509250509295509295909350565b60008060008060006080868803121561040657610405610766565b5b600086013567ffffffffffffffff81111561042457610423610761565b5b6104308882890161029c565b9550506020610441888289016101e4565b9450506040610452888289016102bb565b935050606086013567ffffffffffffffff81111561047357610472610761565b5b61047f888289016101f9565b92509250509295509295909350565b6104978161065b565b82525050565b60006104a98385610639565b93506104b6838584610697565b6104bf8361076b565b840190509392505050565b60006104d58261062e565b6104df818561064a565b93506104ef8185602086016106a6565b6104f88161076b565b840191505092915050565b61050c8161068d565b82525050565b600060a082019050818103600083015261052d81888a61049d565b905061053c602083018761048e565b6105496040830186610503565b610556606083018561048e565b818103608083015261056881846104ca565b9050979650505050505050565b600080833560016020038436030381126105925761059161074d565b5b80840192508235915067ffffffffffffffff8211156105b4576105b3610743565b5b6020830192506001820236038313156105d0576105cf610757565b5b509250929050565b60006105e26105f3565b90506105ee82826106d9565b919050565b6000604051905090565b600067ffffffffffffffff8211156106185761061761070a565b5b6106218261076b565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006106668261066d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156106c45780820151818401526020810190506106a9565b838111156106d3576000848401525b50505050565b6106e28261076b565b810181811067ffffffffffffffff821117156107015761070061070a565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6107858161065b565b811461079057600080fd5b50565b61079c8161068d565b81146107a757600080fd5b5056fea2646970667358221220febf3f8cf5fd0742329aa443cb341b916e26ee1be7a429c48a3a28b956bd845f64736f6c63430008070033",
}

// TestZContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestZContractMetaData.ABI instead.
var TestZContractABI = TestZContractMetaData.ABI

// TestZContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestZContractMetaData.Bin instead.
var TestZContractBin = TestZContractMetaData.Bin

// DeployTestZContract deploys a new Ethereum contract, binding an instance of TestZContract to it.
func DeployTestZContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestZContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestZContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// TestZContract is an auto generated Go binding around an Ethereum contract.
type TestZContract struct {
	TestZContractCaller     // Read-only binding to the contract
	TestZContractTransactor // Write-only binding to the contract
	TestZContractFilterer   // Log filterer for contract events
}

// TestZContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestZContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestZContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestZContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestZContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestZContractSession struct {
	Contract     *TestZContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestZContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestZContractCallerSession struct {
	Contract *TestZContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestZContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestZContractTransactorSession struct {
	Contract     *TestZContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestZContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestZContractRaw struct {
	Contract *TestZContract // Generic contract binding to access the raw methods on
}

// TestZContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestZContractCallerRaw struct {
	Contract *TestZContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestZContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestZContractTransactorRaw struct {
	Contract *TestZContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestZContract creates a new instance of TestZContract, bound to a specific deployed contract.
func NewTestZContract(address common.Address, backend bind.ContractBackend) (*TestZContract, error) {
	contract, err := bindTestZContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestZContract{TestZContractCaller: TestZContractCaller{contract: contract}, TestZContractTransactor: TestZContractTransactor{contract: contract}, TestZContractFilterer: TestZContractFilterer{contract: contract}}, nil
}

// NewTestZContractCaller creates a new read-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractCaller(address common.Address, caller bind.ContractCaller) (*TestZContractCaller, error) {
	contract, err := bindTestZContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractCaller{contract: contract}, nil
}

// NewTestZContractTransactor creates a new write-only instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestZContractTransactor, error) {
	contract, err := bindTestZContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestZContractTransactor{contract: contract}, nil
}

// NewTestZContractFilterer creates a new log filterer instance of TestZContract, bound to a specific deployed contract.
func NewTestZContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestZContractFilterer, error) {
	contract, err := bindTestZContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestZContractFilterer{contract: contract}, nil
}

// bindTestZContract binds a generic wrapper to an already deployed contract.
func bindTestZContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestZContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.TestZContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.TestZContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestZContract *TestZContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestZContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestZContract *TestZContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestZContract *TestZContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestZContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnCrossChainCall(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactor) OnRevert(opts *bind.TransactOpts, context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.contract.Transact(opts, "onRevert", context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0x69582bee.
//
// Solidity: function onRevert((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_TestZContract *TestZContractTransactorSession) OnRevert(context RevertContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.OnRevert(&_TestZContract.TransactOpts, context, zrc20, amount, message)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestZContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.Fallback(&_TestZContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestZContract *TestZContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestZContract.Contract.Fallback(&_TestZContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestZContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractSession) Receive() (*types.Transaction, error) {
	return _TestZContract.Contract.Receive(&_TestZContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestZContract *TestZContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TestZContract.Contract.Receive(&_TestZContract.TransactOpts)
}

// TestZContractContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the TestZContract contract.
type TestZContractContextDataIterator struct {
	Event *TestZContractContextData // Event containing the contract specifics and raw log

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
func (it *TestZContractContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestZContractContextData)
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
		it.Event = new(TestZContractContextData)
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
func (it *TestZContractContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestZContractContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestZContractContextData represents a ContextData event raised by the TestZContract contract.
type TestZContractContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) FilterContextData(opts *bind.FilterOpts) (*TestZContractContextDataIterator, error) {

	logs, sub, err := _TestZContract.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &TestZContractContextDataIterator{contract: _TestZContract.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *TestZContractContextData) (event.Subscription, error) {

	logs, sub, err := _TestZContract.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestZContractContextData)
				if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) ParseContextData(log types.Log) (*TestZContractContextData, error) {
	event := new(TestZContractContextData)
	if err := _TestZContract.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestZContractContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the TestZContract contract.
type TestZContractContextDataRevertIterator struct {
	Event *TestZContractContextDataRevert // Event containing the contract specifics and raw log

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
func (it *TestZContractContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestZContractContextDataRevert)
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
		it.Event = new(TestZContractContextDataRevert)
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
func (it *TestZContractContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestZContractContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestZContractContextDataRevert represents a ContextDataRevert event raised by the TestZContract contract.
type TestZContractContextDataRevert struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*TestZContractContextDataRevertIterator, error) {

	logs, sub, err := _TestZContract.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &TestZContractContextDataRevertIterator{contract: _TestZContract.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *TestZContractContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _TestZContract.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestZContractContextDataRevert)
				if err := _TestZContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xfdc887992b033668833927e252058e468fac0b6bd196d520f09c61b740e99948.
//
// Solidity: event ContextDataRevert(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestZContract *TestZContractFilterer) ParseContextDataRevert(log types.Log) (*TestZContractContextDataRevert, error) {
	event := new(TestZContractContextDataRevert)
	if err := _TestZContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
