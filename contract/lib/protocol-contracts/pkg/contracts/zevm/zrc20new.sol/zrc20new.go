// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrc20new

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

// ZRC20NewMetaData contains all meta data concerning the ZRC20New contract.
var ZRC20NewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chainid_\",\"type\":\"uint256\"},{\"internalType\":\"enumCoinType\",\"name\":\"coinType_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"systemContractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gatewayContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasCoin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasPrice\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"from\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"}],\"name\":\"UpdatedProtocolFlatFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"systemContract\",\"type\":\"address\"}],\"name\":\"UpdatedSystemContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COIN_TYPE\",\"outputs\":[{\"internalType\":\"enumCoinType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GATEWAY_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROTOCOL_FLAT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"updateGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolFlatFee\",\"type\":\"uint256\"}],\"name\":\"updateProtocolFlatFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"updateSystemContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawGasFee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200282d3803806200282d83398181016040528101906200003791906200035b565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8760079080519060200190620000c9929190620001d1565b508660089080519060200190620000e2929190620001d1565b5085600960006101000a81548160ff021916908360ff16021790555084608081815250508360028111156200011c576200011b620005ae565b5b60a0816002811115620001345762000133620005ae565b5b60f81b8152505082600281905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050620006bf565b828054620001df9062000542565b90600052602060002090601f0160209004810192826200020357600085556200024f565b82601f106200021e57805160ff19168380011785556200024f565b828001600101855582156200024f579182015b828111156200024e57825182559160200191906001019062000231565b5b5090506200025e919062000262565b5090565b5b808211156200027d57600081600090555060010162000263565b5090565b60006200029862000292846200048b565b62000462565b905082815260208101848484011115620002b757620002b662000640565b5b620002c48482856200050c565b509392505050565b600081519050620002dd8162000660565b92915050565b600081519050620002f4816200067a565b92915050565b600082601f8301126200031257620003116200063b565b5b81516200032484826020860162000281565b91505092915050565b6000815190506200033e816200068b565b92915050565b6000815190506200035581620006a5565b92915050565b600080600080600080600080610100898b0312156200037f576200037e6200064a565b5b600089015167ffffffffffffffff811115620003a0576200039f62000645565b5b620003ae8b828c01620002fa565b985050602089015167ffffffffffffffff811115620003d257620003d162000645565b5b620003e08b828c01620002fa565b9750506040620003f38b828c0162000344565b9650506060620004068b828c016200032d565b9550506080620004198b828c01620002e3565b94505060a06200042c8b828c016200032d565b93505060c06200043f8b828c01620002cc565b92505060e0620004528b828c01620002cc565b9150509295985092959890939650565b60006200046e62000481565b90506200047c828262000578565b919050565b6000604051905090565b600067ffffffffffffffff821115620004a957620004a86200060c565b5b620004b4826200064f565b9050602081019050919050565b6000620004ce82620004d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156200052c5780820151818401526020810190506200050f565b838111156200053c576000848401525b50505050565b600060028204905060018216806200055b57607f821691505b60208210811415620005725762000571620005dd565b5b50919050565b62000583826200064f565b810181811067ffffffffffffffff82111715620005a557620005a46200060c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6200066b81620004c1565b81146200067757600080fd5b50565b600381106200068857600080fd5b50565b6200069681620004f5565b8114620006a257600080fd5b50565b620006b081620004ff565b8114620006bc57600080fd5b50565b60805160a05160f81c612137620006f660003960006109580152600081816108a201528181610c250152610d5a01526121376000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806385e1f4d0116100c3578063d9eeebed1161007c578063d9eeebed146103cc578063dd62ed3e146103eb578063e2f535b81461041b578063eddeb12314610439578063f2441b3214610455578063f687d12a146104735761014d565b806385e1f4d0146102f657806395d89b4114610314578063a3413d0314610332578063a9059cbb14610350578063c701262614610380578063c835d7cc146103b05761014d565b8063313ce56711610115578063313ce5671461020c5780633ce4a5bc1461022a57806342966c681461024857806347e7ef24146102785780634d8943bb146102a857806370a08231146102c65761014d565b806306fdde0314610152578063091d278814610170578063095ea7b31461018e57806318160ddd146101be57806323b872dd146101dc575b600080fd5b61015a61048f565b6040516101679190611cad565b60405180910390f35b610178610521565b6040516101859190611ccf565b60405180910390f35b6101a860048036038101906101a3919061196e565b610527565b6040516101b59190611bfb565b60405180910390f35b6101c6610545565b6040516101d39190611ccf565b60405180910390f35b6101f660048036038101906101f1919061191b565b61054f565b6040516102039190611bfb565b60405180910390f35b610214610647565b6040516102219190611cea565b60405180910390f35b61023261065e565b60405161023f9190611b80565b60405180910390f35b610262600480360381019061025d9190611a37565b610676565b60405161026f9190611bfb565b60405180910390f35b610292600480360381019061028d919061196e565b61068b565b60405161029f9190611bfb565b60405180910390f35b6102b0610851565b6040516102bd9190611ccf565b60405180910390f35b6102e060048036038101906102db9190611881565b610857565b6040516102ed9190611ccf565b60405180910390f35b6102fe6108a0565b60405161030b9190611ccf565b60405180910390f35b61031c6108c4565b6040516103299190611cad565b60405180910390f35b61033a610956565b6040516103479190611c92565b60405180910390f35b61036a6004803603810190610365919061196e565b61097a565b6040516103779190611bfb565b60405180910390f35b61039a600480360381019061039591906119db565b610998565b6040516103a79190611bfb565b60405180910390f35b6103ca60048036038101906103c59190611881565b610aee565b005b6103d4610be1565b6040516103e2929190611bd2565b60405180910390f35b610405600480360381019061040091906118db565b610e4e565b6040516104129190611ccf565b60405180910390f35b610423610ed5565b6040516104309190611b80565b60405180910390f35b610453600480360381019061044e9190611a37565b610efb565b005b61045d610fb5565b60405161046a9190611b80565b60405180910390f35b61048d60048036038101906104889190611a37565b610fd9565b005b60606007805461049e90611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546104ca90611f33565b80156105175780601f106104ec57610100808354040283529160200191610517565b820191906000526020600020905b8154815290600101906020018083116104fa57829003601f168201915b5050505050905090565b60025481565b600061053b610534611093565b848461109b565b6001905092915050565b6000600654905090565b600061055c848484611254565b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006105a7611093565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561061e576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61063b8561062a611093565b85846106369190611e43565b61109b565b60019150509392505050565b6000600960009054906101000a900460ff16905090565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b600061068233836114b0565b60019050919050565b600073735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610729575060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b80156107835750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156107ba576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107c48383611668565b8273ffffffffffffffffffffffffffffffffffffffff167f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab373735b14bb79463307aacbed86daf3322b1e6226ab6040516020016108219190611b65565b6040516020818303038152906040528460405161083f929190611c16565b60405180910390a26001905092915050565b60035481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060600880546108d390611f33565b80601f01602080910402602001604051908101604052809291908181526020018280546108ff90611f33565b801561094c5780601f106109215761010080835404028352916020019161094c565b820191906000526020600020905b81548152906001019060200180831161092f57829003601f168201915b5050505050905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061098e610987611093565b8484611254565b6001905092915050565b60008060006109a5610be1565b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3373735b14bb79463307aacbed86daf3322b1e6226ab846040518463ffffffff1660e01b81526004016109fa93929190611b9b565b602060405180830381600087803b158015610a1457600080fd5b505af1158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c91906119ae565b610a82576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8c33856114b0565b3373ffffffffffffffffffffffffffffffffffffffff167f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955868684600354604051610ada9493929190611c46565b60405180910390a260019250505092915050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b67576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae81604051610bd69190611b80565b60405180910390a150565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630be155477f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610c609190611ccf565b60206040518083038186803b158015610c7857600080fd5b505afa158015610c8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb091906118ae565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d19576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7fd7afb7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610d959190611ccf565b60206040518083038186803b158015610dad57600080fd5b505afa158015610dc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de59190611a64565b90506000811415610e22576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060035460025483610e359190611de9565b610e3f9190611d93565b90508281945094505050509091565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f74576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003819055507fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f81604051610faa9190611ccf565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611052576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806002819055507fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a816040516110889190611ccf565b60405180910390a150565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611102576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611169576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516112479190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156112bb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611322576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156113a0576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816113ac9190611e43565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461143e9190611d93565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114a29190611ccf565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611517576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611595576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81816115a19190611e43565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600660008282546115f69190611e43565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161165b9190611ccf565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116cf576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600660008282546116e19190611d93565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546117379190611d93565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161179c9190611ccf565b60405180910390a35050565b60006117bb6117b684611d2a565b611d05565b9050828152602081018484840111156117d7576117d661207b565b5b6117e2848285611ef1565b509392505050565b6000813590506117f9816120bc565b92915050565b60008151905061180e816120bc565b92915050565b600081519050611823816120d3565b92915050565b600082601f83011261183e5761183d612076565b5b813561184e8482602086016117a8565b91505092915050565b600081359050611866816120ea565b92915050565b60008151905061187b816120ea565b92915050565b60006020828403121561189757611896612085565b5b60006118a5848285016117ea565b91505092915050565b6000602082840312156118c4576118c3612085565b5b60006118d2848285016117ff565b91505092915050565b600080604083850312156118f2576118f1612085565b5b6000611900858286016117ea565b9250506020611911858286016117ea565b9150509250929050565b60008060006060848603121561193457611933612085565b5b6000611942868287016117ea565b9350506020611953868287016117ea565b925050604061196486828701611857565b9150509250925092565b6000806040838503121561198557611984612085565b5b6000611993858286016117ea565b92505060206119a485828601611857565b9150509250929050565b6000602082840312156119c4576119c3612085565b5b60006119d284828501611814565b91505092915050565b600080604083850312156119f2576119f1612085565b5b600083013567ffffffffffffffff811115611a1057611a0f612080565b5b611a1c85828601611829565b9250506020611a2d85828601611857565b9150509250929050565b600060208284031215611a4d57611a4c612085565b5b6000611a5b84828501611857565b91505092915050565b600060208284031215611a7a57611a79612085565b5b6000611a888482850161186c565b91505092915050565b611a9a81611e77565b82525050565b611ab1611aac82611e77565b611f96565b82525050565b611ac081611e89565b82525050565b6000611ad182611d5b565b611adb8185611d71565b9350611aeb818560208601611f00565b611af48161208a565b840191505092915050565b611b0881611edf565b82525050565b6000611b1982611d66565b611b238185611d82565b9350611b33818560208601611f00565b611b3c8161208a565b840191505092915050565b611b5081611ec8565b82525050565b611b5f81611ed2565b82525050565b6000611b718284611aa0565b60148201915081905092915050565b6000602082019050611b956000830184611a91565b92915050565b6000606082019050611bb06000830186611a91565b611bbd6020830185611a91565b611bca6040830184611b47565b949350505050565b6000604082019050611be76000830185611a91565b611bf46020830184611b47565b9392505050565b6000602082019050611c106000830184611ab7565b92915050565b60006040820190508181036000830152611c308185611ac6565b9050611c3f6020830184611b47565b9392505050565b60006080820190508181036000830152611c608187611ac6565b9050611c6f6020830186611b47565b611c7c6040830185611b47565b611c896060830184611b47565b95945050505050565b6000602082019050611ca76000830184611aff565b92915050565b60006020820190508181036000830152611cc78184611b0e565b905092915050565b6000602082019050611ce46000830184611b47565b92915050565b6000602082019050611cff6000830184611b56565b92915050565b6000611d0f611d20565b9050611d1b8282611f65565b919050565b6000604051905090565b600067ffffffffffffffff821115611d4557611d44612047565b5b611d4e8261208a565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611d9e82611ec8565b9150611da983611ec8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611dde57611ddd611fba565b5b828201905092915050565b6000611df482611ec8565b9150611dff83611ec8565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e3857611e37611fba565b5b828202905092915050565b6000611e4e82611ec8565b9150611e5983611ec8565b925082821015611e6c57611e6b611fba565b5b828203905092915050565b6000611e8282611ea8565b9050919050565b60008115159050919050565b6000819050611ea3826120a8565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611eea82611e95565b9050919050565b82818337600083830152505050565b60005b83811015611f1e578082015181840152602081019050611f03565b83811115611f2d576000848401525b50505050565b60006002820490506001821680611f4b57607f821691505b60208210811415611f5f57611f5e612018565b5b50919050565b611f6e8261208a565b810181811067ffffffffffffffff82111715611f8d57611f8c612047565b5b80604052505050565b6000611fa182611fa8565b9050919050565b6000611fb38261209b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b600381106120b9576120b8611fe9565b5b50565b6120c581611e77565b81146120d057600080fd5b50565b6120dc81611e89565b81146120e757600080fd5b50565b6120f381611ec8565b81146120fe57600080fd5b5056fea264697066735822122088cc6797a637e9f7f0d8220bcf745a632c2a8fcb9c479e8f3633f88f9d369ecc64736f6c63430008070033",
}

// ZRC20NewABI is the input ABI used to generate the binding from.
// Deprecated: Use ZRC20NewMetaData.ABI instead.
var ZRC20NewABI = ZRC20NewMetaData.ABI

// ZRC20NewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZRC20NewMetaData.Bin instead.
var ZRC20NewBin = ZRC20NewMetaData.Bin

// DeployZRC20New deploys a new Ethereum contract, binding an instance of ZRC20New to it.
func DeployZRC20New(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, chainid_ *big.Int, coinType_ uint8, gasLimit_ *big.Int, systemContractAddress_ common.Address, gatewayContractAddress_ common.Address) (common.Address, *types.Transaction, *ZRC20New, error) {
	parsed, err := ZRC20NewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZRC20NewBin), backend, name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_, gatewayContractAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZRC20New{ZRC20NewCaller: ZRC20NewCaller{contract: contract}, ZRC20NewTransactor: ZRC20NewTransactor{contract: contract}, ZRC20NewFilterer: ZRC20NewFilterer{contract: contract}}, nil
}

// ZRC20New is an auto generated Go binding around an Ethereum contract.
type ZRC20New struct {
	ZRC20NewCaller     // Read-only binding to the contract
	ZRC20NewTransactor // Write-only binding to the contract
	ZRC20NewFilterer   // Log filterer for contract events
}

// ZRC20NewCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZRC20NewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20NewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZRC20NewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20NewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZRC20NewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20NewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZRC20NewSession struct {
	Contract     *ZRC20New         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20NewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZRC20NewCallerSession struct {
	Contract *ZRC20NewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZRC20NewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZRC20NewTransactorSession struct {
	Contract     *ZRC20NewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZRC20NewRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZRC20NewRaw struct {
	Contract *ZRC20New // Generic contract binding to access the raw methods on
}

// ZRC20NewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZRC20NewCallerRaw struct {
	Contract *ZRC20NewCaller // Generic read-only contract binding to access the raw methods on
}

// ZRC20NewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZRC20NewTransactorRaw struct {
	Contract *ZRC20NewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZRC20New creates a new instance of ZRC20New, bound to a specific deployed contract.
func NewZRC20New(address common.Address, backend bind.ContractBackend) (*ZRC20New, error) {
	contract, err := bindZRC20New(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZRC20New{ZRC20NewCaller: ZRC20NewCaller{contract: contract}, ZRC20NewTransactor: ZRC20NewTransactor{contract: contract}, ZRC20NewFilterer: ZRC20NewFilterer{contract: contract}}, nil
}

// NewZRC20NewCaller creates a new read-only instance of ZRC20New, bound to a specific deployed contract.
func NewZRC20NewCaller(address common.Address, caller bind.ContractCaller) (*ZRC20NewCaller, error) {
	contract, err := bindZRC20New(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewCaller{contract: contract}, nil
}

// NewZRC20NewTransactor creates a new write-only instance of ZRC20New, bound to a specific deployed contract.
func NewZRC20NewTransactor(address common.Address, transactor bind.ContractTransactor) (*ZRC20NewTransactor, error) {
	contract, err := bindZRC20New(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewTransactor{contract: contract}, nil
}

// NewZRC20NewFilterer creates a new log filterer instance of ZRC20New, bound to a specific deployed contract.
func NewZRC20NewFilterer(address common.Address, filterer bind.ContractFilterer) (*ZRC20NewFilterer, error) {
	contract, err := bindZRC20New(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewFilterer{contract: contract}, nil
}

// bindZRC20New binds a generic wrapper to an already deployed contract.
func bindZRC20New(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZRC20NewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20New *ZRC20NewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20New.Contract.ZRC20NewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20New *ZRC20NewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20New.Contract.ZRC20NewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20New *ZRC20NewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20New.Contract.ZRC20NewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20New *ZRC20NewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20New.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20New *ZRC20NewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20New.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20New *ZRC20NewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20New.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20New *ZRC20NewSession) CHAINID() (*big.Int, error) {
	return _ZRC20New.Contract.CHAINID(&_ZRC20New.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) CHAINID() (*big.Int, error) {
	return _ZRC20New.Contract.CHAINID(&_ZRC20New.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20New *ZRC20NewCaller) COINTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "COIN_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20New *ZRC20NewSession) COINTYPE() (uint8, error) {
	return _ZRC20New.Contract.COINTYPE(&_ZRC20New.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20New *ZRC20NewCallerSession) COINTYPE() (uint8, error) {
	return _ZRC20New.Contract.COINTYPE(&_ZRC20New.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20New.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20New.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20New *ZRC20NewSession) GASLIMIT() (*big.Int, error) {
	return _ZRC20New.Contract.GASLIMIT(&_ZRC20New.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) GASLIMIT() (*big.Int, error) {
	return _ZRC20New.Contract.GASLIMIT(&_ZRC20New.CallOpts)
}

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCaller) GATEWAYCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "GATEWAY_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewSession) GATEWAYCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.GATEWAYCONTRACTADDRESS(&_ZRC20New.CallOpts)
}

// GATEWAYCONTRACTADDRESS is a free data retrieval call binding the contract method 0xe2f535b8.
//
// Solidity: function GATEWAY_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCallerSession) GATEWAYCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.GATEWAYCONTRACTADDRESS(&_ZRC20New.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20New *ZRC20NewSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20New.Contract.PROTOCOLFLATFEE(&_ZRC20New.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20New.Contract.PROTOCOLFLATFEE(&_ZRC20New.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCaller) SYSTEMCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "SYSTEM_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewSession) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20New.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20New *ZRC20NewCallerSession) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20New.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20New.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20New *ZRC20NewSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20New.Contract.Allowance(&_ZRC20New.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20New.Contract.Allowance(&_ZRC20New.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20New *ZRC20NewSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20New.Contract.BalanceOf(&_ZRC20New.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20New.Contract.BalanceOf(&_ZRC20New.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20New *ZRC20NewCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20New *ZRC20NewSession) Decimals() (uint8, error) {
	return _ZRC20New.Contract.Decimals(&_ZRC20New.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20New *ZRC20NewCallerSession) Decimals() (uint8, error) {
	return _ZRC20New.Contract.Decimals(&_ZRC20New.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20New *ZRC20NewCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20New *ZRC20NewSession) Name() (string, error) {
	return _ZRC20New.Contract.Name(&_ZRC20New.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20New *ZRC20NewCallerSession) Name() (string, error) {
	return _ZRC20New.Contract.Name(&_ZRC20New.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20New *ZRC20NewCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20New *ZRC20NewSession) Symbol() (string, error) {
	return _ZRC20New.Contract.Symbol(&_ZRC20New.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20New *ZRC20NewCallerSession) Symbol() (string, error) {
	return _ZRC20New.Contract.Symbol(&_ZRC20New.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20New *ZRC20NewCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20New *ZRC20NewSession) TotalSupply() (*big.Int, error) {
	return _ZRC20New.Contract.TotalSupply(&_ZRC20New.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20New *ZRC20NewCallerSession) TotalSupply() (*big.Int, error) {
	return _ZRC20New.Contract.TotalSupply(&_ZRC20New.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20New *ZRC20NewCaller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ZRC20New.contract.Call(opts, &out, "withdrawGasFee")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20New *ZRC20NewSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20New.Contract.WithdrawGasFee(&_ZRC20New.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20New *ZRC20NewCallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20New.Contract.WithdrawGasFee(&_ZRC20New.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Approve(&_ZRC20New.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Approve(&_ZRC20New.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Burn(&_ZRC20New.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Burn(&_ZRC20New.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Deposit(&_ZRC20New.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Deposit(&_ZRC20New.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Transfer(&_ZRC20New.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Transfer(&_ZRC20New.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.TransferFrom(&_ZRC20New.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.TransferFrom(&_ZRC20New.TransactOpts, sender, recipient, amount)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20New *ZRC20NewTransactor) UpdateGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "updateGasLimit", gasLimit)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20New *ZRC20NewSession) UpdateGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateGasLimit(&_ZRC20New.TransactOpts, gasLimit)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit) returns()
func (_ZRC20New *ZRC20NewTransactorSession) UpdateGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateGasLimit(&_ZRC20New.TransactOpts, gasLimit)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20New *ZRC20NewTransactor) UpdateProtocolFlatFee(opts *bind.TransactOpts, protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "updateProtocolFlatFee", protocolFlatFee)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20New *ZRC20NewSession) UpdateProtocolFlatFee(protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateProtocolFlatFee(&_ZRC20New.TransactOpts, protocolFlatFee)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee) returns()
func (_ZRC20New *ZRC20NewTransactorSession) UpdateProtocolFlatFee(protocolFlatFee *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateProtocolFlatFee(&_ZRC20New.TransactOpts, protocolFlatFee)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20New *ZRC20NewTransactor) UpdateSystemContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "updateSystemContractAddress", addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20New *ZRC20NewSession) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateSystemContractAddress(&_ZRC20New.TransactOpts, addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20New *ZRC20NewTransactorSession) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20New.Contract.UpdateSystemContractAddress(&_ZRC20New.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Withdraw(&_ZRC20New.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20New *ZRC20NewTransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20New.Contract.Withdraw(&_ZRC20New.TransactOpts, to, amount)
}

// ZRC20NewApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZRC20New contract.
type ZRC20NewApprovalIterator struct {
	Event *ZRC20NewApproval // Event containing the contract specifics and raw log

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
func (it *ZRC20NewApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewApproval)
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
		it.Event = new(ZRC20NewApproval)
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
func (it *ZRC20NewApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewApproval represents a Approval event raised by the ZRC20New contract.
type ZRC20NewApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZRC20NewApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewApprovalIterator{contract: _ZRC20New.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZRC20NewApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewApproval)
				if err := _ZRC20New.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) ParseApproval(log types.Log) (*ZRC20NewApproval, error) {
	event := new(ZRC20NewApproval)
	if err := _ZRC20New.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZRC20New contract.
type ZRC20NewDepositIterator struct {
	Event *ZRC20NewDeposit // Event containing the contract specifics and raw log

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
func (it *ZRC20NewDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewDeposit)
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
		it.Event = new(ZRC20NewDeposit)
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
func (it *ZRC20NewDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewDeposit represents a Deposit event raised by the ZRC20New contract.
type ZRC20NewDeposit struct {
	From  []byte
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) FilterDeposit(opts *bind.FilterOpts, to []common.Address) (*ZRC20NewDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewDepositIterator{contract: _ZRC20New.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZRC20NewDeposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewDeposit)
				if err := _ZRC20New.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) ParseDeposit(log types.Log) (*ZRC20NewDeposit, error) {
	event := new(ZRC20NewDeposit)
	if err := _ZRC20New.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZRC20New contract.
type ZRC20NewTransferIterator struct {
	Event *ZRC20NewTransfer // Event containing the contract specifics and raw log

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
func (it *ZRC20NewTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewTransfer)
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
		it.Event = new(ZRC20NewTransfer)
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
func (it *ZRC20NewTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewTransfer represents a Transfer event raised by the ZRC20New contract.
type ZRC20NewTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZRC20NewTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewTransferIterator{contract: _ZRC20New.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZRC20NewTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewTransfer)
				if err := _ZRC20New.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20New *ZRC20NewFilterer) ParseTransfer(log types.Log) (*ZRC20NewTransfer, error) {
	event := new(ZRC20NewTransfer)
	if err := _ZRC20New.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewUpdatedGasLimitIterator is returned from FilterUpdatedGasLimit and is used to iterate over the raw logs and unpacked data for UpdatedGasLimit events raised by the ZRC20New contract.
type ZRC20NewUpdatedGasLimitIterator struct {
	Event *ZRC20NewUpdatedGasLimit // Event containing the contract specifics and raw log

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
func (it *ZRC20NewUpdatedGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewUpdatedGasLimit)
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
		it.Event = new(ZRC20NewUpdatedGasLimit)
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
func (it *ZRC20NewUpdatedGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewUpdatedGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewUpdatedGasLimit represents a UpdatedGasLimit event raised by the ZRC20New contract.
type ZRC20NewUpdatedGasLimit struct {
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasLimit is a free log retrieval operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20New *ZRC20NewFilterer) FilterUpdatedGasLimit(opts *bind.FilterOpts) (*ZRC20NewUpdatedGasLimitIterator, error) {

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return &ZRC20NewUpdatedGasLimitIterator{contract: _ZRC20New.contract, event: "UpdatedGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasLimit is a free log subscription operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20New *ZRC20NewFilterer) WatchUpdatedGasLimit(opts *bind.WatchOpts, sink chan<- *ZRC20NewUpdatedGasLimit) (event.Subscription, error) {

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewUpdatedGasLimit)
				if err := _ZRC20New.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
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

// ParseUpdatedGasLimit is a log parse operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20New *ZRC20NewFilterer) ParseUpdatedGasLimit(log types.Log) (*ZRC20NewUpdatedGasLimit, error) {
	event := new(ZRC20NewUpdatedGasLimit)
	if err := _ZRC20New.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewUpdatedProtocolFlatFeeIterator is returned from FilterUpdatedProtocolFlatFee and is used to iterate over the raw logs and unpacked data for UpdatedProtocolFlatFee events raised by the ZRC20New contract.
type ZRC20NewUpdatedProtocolFlatFeeIterator struct {
	Event *ZRC20NewUpdatedProtocolFlatFee // Event containing the contract specifics and raw log

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
func (it *ZRC20NewUpdatedProtocolFlatFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewUpdatedProtocolFlatFee)
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
		it.Event = new(ZRC20NewUpdatedProtocolFlatFee)
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
func (it *ZRC20NewUpdatedProtocolFlatFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewUpdatedProtocolFlatFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewUpdatedProtocolFlatFee represents a UpdatedProtocolFlatFee event raised by the ZRC20New contract.
type ZRC20NewUpdatedProtocolFlatFee struct {
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProtocolFlatFee is a free log retrieval operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) FilterUpdatedProtocolFlatFee(opts *bind.FilterOpts) (*ZRC20NewUpdatedProtocolFlatFeeIterator, error) {

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return &ZRC20NewUpdatedProtocolFlatFeeIterator{contract: _ZRC20New.contract, event: "UpdatedProtocolFlatFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedProtocolFlatFee is a free log subscription operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) WatchUpdatedProtocolFlatFee(opts *bind.WatchOpts, sink chan<- *ZRC20NewUpdatedProtocolFlatFee) (event.Subscription, error) {

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewUpdatedProtocolFlatFee)
				if err := _ZRC20New.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
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

// ParseUpdatedProtocolFlatFee is a log parse operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) ParseUpdatedProtocolFlatFee(log types.Log) (*ZRC20NewUpdatedProtocolFlatFee, error) {
	event := new(ZRC20NewUpdatedProtocolFlatFee)
	if err := _ZRC20New.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewUpdatedSystemContractIterator is returned from FilterUpdatedSystemContract and is used to iterate over the raw logs and unpacked data for UpdatedSystemContract events raised by the ZRC20New contract.
type ZRC20NewUpdatedSystemContractIterator struct {
	Event *ZRC20NewUpdatedSystemContract // Event containing the contract specifics and raw log

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
func (it *ZRC20NewUpdatedSystemContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewUpdatedSystemContract)
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
		it.Event = new(ZRC20NewUpdatedSystemContract)
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
func (it *ZRC20NewUpdatedSystemContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewUpdatedSystemContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewUpdatedSystemContract represents a UpdatedSystemContract event raised by the ZRC20New contract.
type ZRC20NewUpdatedSystemContract struct {
	SystemContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSystemContract is a free log retrieval operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20New *ZRC20NewFilterer) FilterUpdatedSystemContract(opts *bind.FilterOpts) (*ZRC20NewUpdatedSystemContractIterator, error) {

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return &ZRC20NewUpdatedSystemContractIterator{contract: _ZRC20New.contract, event: "UpdatedSystemContract", logs: logs, sub: sub}, nil
}

// WatchUpdatedSystemContract is a free log subscription operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20New *ZRC20NewFilterer) WatchUpdatedSystemContract(opts *bind.WatchOpts, sink chan<- *ZRC20NewUpdatedSystemContract) (event.Subscription, error) {

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewUpdatedSystemContract)
				if err := _ZRC20New.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
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

// ParseUpdatedSystemContract is a log parse operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20New *ZRC20NewFilterer) ParseUpdatedSystemContract(log types.Log) (*ZRC20NewUpdatedSystemContract, error) {
	event := new(ZRC20NewUpdatedSystemContract)
	if err := _ZRC20New.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20NewWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZRC20New contract.
type ZRC20NewWithdrawalIterator struct {
	Event *ZRC20NewWithdrawal // Event containing the contract specifics and raw log

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
func (it *ZRC20NewWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20NewWithdrawal)
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
		it.Event = new(ZRC20NewWithdrawal)
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
func (it *ZRC20NewWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20NewWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20NewWithdrawal represents a Withdrawal event raised by the ZRC20New contract.
type ZRC20NewWithdrawal struct {
	From            common.Address
	To              []byte
	Value           *big.Int
	GasFee          *big.Int
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*ZRC20NewWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20New.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20NewWithdrawalIterator{contract: _ZRC20New.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZRC20NewWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20New.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20NewWithdrawal)
				if err := _ZRC20New.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20New *ZRC20NewFilterer) ParseWithdrawal(log types.Log) (*ZRC20NewWithdrawal, error) {
	event := new(ZRC20NewWithdrawal)
	if err := _ZRC20New.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
