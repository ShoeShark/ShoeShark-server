// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ShoeSharkTokenMetaData contains all meta data concerning the ShoeSharkToken contract.
var ShoeSharkTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkToken_Burned\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkToken_Minted\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ShoeSharkToken__InvalidMintAmount\",\"inputs\":[]}]",
}

// ShoeSharkTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ShoeSharkTokenMetaData.ABI instead.
var ShoeSharkTokenABI = ShoeSharkTokenMetaData.ABI

// ShoeSharkToken is an auto generated Go binding around an Ethereum contract.
type ShoeSharkToken struct {
	ShoeSharkTokenCaller     // Read-only binding to the contract
	ShoeSharkTokenTransactor // Write-only binding to the contract
	ShoeSharkTokenFilterer   // Log filterer for contract events
}

// ShoeSharkTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShoeSharkTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShoeSharkTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShoeSharkTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShoeSharkTokenSession struct {
	Contract     *ShoeSharkToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShoeSharkTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShoeSharkTokenCallerSession struct {
	Contract *ShoeSharkTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ShoeSharkTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShoeSharkTokenTransactorSession struct {
	Contract     *ShoeSharkTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ShoeSharkTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShoeSharkTokenRaw struct {
	Contract *ShoeSharkToken // Generic contract binding to access the raw methods on
}

// ShoeSharkTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShoeSharkTokenCallerRaw struct {
	Contract *ShoeSharkTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ShoeSharkTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShoeSharkTokenTransactorRaw struct {
	Contract *ShoeSharkTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShoeSharkToken creates a new instance of ShoeSharkToken, bound to a specific deployed contract.
func NewShoeSharkToken(address common.Address, backend bind.ContractBackend) (*ShoeSharkToken, error) {
	contract, err := bindShoeSharkToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkToken{ShoeSharkTokenCaller: ShoeSharkTokenCaller{contract: contract}, ShoeSharkTokenTransactor: ShoeSharkTokenTransactor{contract: contract}, ShoeSharkTokenFilterer: ShoeSharkTokenFilterer{contract: contract}}, nil
}

// NewShoeSharkTokenCaller creates a new read-only instance of ShoeSharkToken, bound to a specific deployed contract.
func NewShoeSharkTokenCaller(address common.Address, caller bind.ContractCaller) (*ShoeSharkTokenCaller, error) {
	contract, err := bindShoeSharkToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenCaller{contract: contract}, nil
}

// NewShoeSharkTokenTransactor creates a new write-only instance of ShoeSharkToken, bound to a specific deployed contract.
func NewShoeSharkTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ShoeSharkTokenTransactor, error) {
	contract, err := bindShoeSharkToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenTransactor{contract: contract}, nil
}

// NewShoeSharkTokenFilterer creates a new log filterer instance of ShoeSharkToken, bound to a specific deployed contract.
func NewShoeSharkTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ShoeSharkTokenFilterer, error) {
	contract, err := bindShoeSharkToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenFilterer{contract: contract}, nil
}

// bindShoeSharkToken binds a generic wrapper to an already deployed contract.
func bindShoeSharkToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShoeSharkTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkToken *ShoeSharkTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkToken.Contract.ShoeSharkTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkToken *ShoeSharkTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.ShoeSharkTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkToken *ShoeSharkTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.ShoeSharkTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkToken *ShoeSharkTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkToken *ShoeSharkTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkToken *ShoeSharkTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShoeSharkToken.Contract.Allowance(&_ShoeSharkToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ShoeSharkToken.Contract.Allowance(&_ShoeSharkToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShoeSharkToken.Contract.BalanceOf(&_ShoeSharkToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ShoeSharkToken.Contract.BalanceOf(&_ShoeSharkToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShoeSharkToken *ShoeSharkTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShoeSharkToken *ShoeSharkTokenSession) Decimals() (uint8, error) {
	return _ShoeSharkToken.Contract.Decimals(&_ShoeSharkToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) Decimals() (uint8, error) {
	return _ShoeSharkToken.Contract.Decimals(&_ShoeSharkToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenSession) Name() (string, error) {
	return _ShoeSharkToken.Contract.Name(&_ShoeSharkToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) Name() (string, error) {
	return _ShoeSharkToken.Contract.Name(&_ShoeSharkToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkToken *ShoeSharkTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkToken *ShoeSharkTokenSession) Owner() (common.Address, error) {
	return _ShoeSharkToken.Contract.Owner(&_ShoeSharkToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) Owner() (common.Address, error) {
	return _ShoeSharkToken.Contract.Owner(&_ShoeSharkToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenSession) Symbol() (string, error) {
	return _ShoeSharkToken.Contract.Symbol(&_ShoeSharkToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) Symbol() (string, error) {
	return _ShoeSharkToken.Contract.Symbol(&_ShoeSharkToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenSession) TotalSupply() (*big.Int, error) {
	return _ShoeSharkToken.Contract.TotalSupply(&_ShoeSharkToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkToken *ShoeSharkTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ShoeSharkToken.Contract.TotalSupply(&_ShoeSharkToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Approve(&_ShoeSharkToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Approve(&_ShoeSharkToken.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ShoeSharkToken *ShoeSharkTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Mint(&_ShoeSharkToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Mint(&_ShoeSharkToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkToken *ShoeSharkTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.RenounceOwnership(&_ShoeSharkToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.RenounceOwnership(&_ShoeSharkToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Transfer(&_ShoeSharkToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.Transfer(&_ShoeSharkToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.TransferFrom(&_ShoeSharkToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.TransferFrom(&_ShoeSharkToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkToken *ShoeSharkTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.TransferOwnership(&_ShoeSharkToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkToken *ShoeSharkTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkToken.Contract.TransferOwnership(&_ShoeSharkToken.TransactOpts, newOwner)
}

// ShoeSharkTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShoeSharkToken contract.
type ShoeSharkTokenApprovalIterator struct {
	Event *ShoeSharkTokenApproval // Event containing the contract specifics and raw log

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
func (it *ShoeSharkTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkTokenApproval)
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
		it.Event = new(ShoeSharkTokenApproval)
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
func (it *ShoeSharkTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkTokenApproval represents a Approval event raised by the ShoeSharkToken contract.
type ShoeSharkTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ShoeSharkTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenApprovalIterator{contract: _ShoeSharkToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShoeSharkTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkTokenApproval)
				if err := _ShoeSharkToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ShoeSharkToken *ShoeSharkTokenFilterer) ParseApproval(log types.Log) (*ShoeSharkTokenApproval, error) {
	event := new(ShoeSharkTokenApproval)
	if err := _ShoeSharkToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShoeSharkToken contract.
type ShoeSharkTokenOwnershipTransferredIterator struct {
	Event *ShoeSharkTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShoeSharkTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkTokenOwnershipTransferred)
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
		it.Event = new(ShoeSharkTokenOwnershipTransferred)
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
func (it *ShoeSharkTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ShoeSharkToken contract.
type ShoeSharkTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShoeSharkTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenOwnershipTransferredIterator{contract: _ShoeSharkToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShoeSharkTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkTokenOwnershipTransferred)
				if err := _ShoeSharkToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ShoeSharkTokenOwnershipTransferred, error) {
	event := new(ShoeSharkTokenOwnershipTransferred)
	if err := _ShoeSharkToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkTokenShoeSharkTokenBurnedIterator is returned from FilterShoeSharkTokenBurned and is used to iterate over the raw logs and unpacked data for ShoeSharkTokenBurned events raised by the ShoeSharkToken contract.
type ShoeSharkTokenShoeSharkTokenBurnedIterator struct {
	Event *ShoeSharkTokenShoeSharkTokenBurned // Event containing the contract specifics and raw log

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
func (it *ShoeSharkTokenShoeSharkTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkTokenShoeSharkTokenBurned)
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
		it.Event = new(ShoeSharkTokenShoeSharkTokenBurned)
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
func (it *ShoeSharkTokenShoeSharkTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkTokenShoeSharkTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkTokenShoeSharkTokenBurned represents a ShoeSharkTokenBurned event raised by the ShoeSharkToken contract.
type ShoeSharkTokenShoeSharkTokenBurned struct {
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkTokenBurned is a free log retrieval operation binding the contract event 0xf3aaaebee0e6b4b3197b10213391480b5fd80f118f7549df5ac1b209f938ef7b.
//
// Solidity: event ShoeSharkToken_Burned(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) FilterShoeSharkTokenBurned(opts *bind.FilterOpts, player []common.Address) (*ShoeSharkTokenShoeSharkTokenBurnedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.FilterLogs(opts, "ShoeSharkToken_Burned", playerRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenShoeSharkTokenBurnedIterator{contract: _ShoeSharkToken.contract, event: "ShoeSharkToken_Burned", logs: logs, sub: sub}, nil
}

// WatchShoeSharkTokenBurned is a free log subscription operation binding the contract event 0xf3aaaebee0e6b4b3197b10213391480b5fd80f118f7549df5ac1b209f938ef7b.
//
// Solidity: event ShoeSharkToken_Burned(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) WatchShoeSharkTokenBurned(opts *bind.WatchOpts, sink chan<- *ShoeSharkTokenShoeSharkTokenBurned, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.WatchLogs(opts, "ShoeSharkToken_Burned", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkTokenShoeSharkTokenBurned)
				if err := _ShoeSharkToken.contract.UnpackLog(event, "ShoeSharkToken_Burned", log); err != nil {
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

// ParseShoeSharkTokenBurned is a log parse operation binding the contract event 0xf3aaaebee0e6b4b3197b10213391480b5fd80f118f7549df5ac1b209f938ef7b.
//
// Solidity: event ShoeSharkToken_Burned(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) ParseShoeSharkTokenBurned(log types.Log) (*ShoeSharkTokenShoeSharkTokenBurned, error) {
	event := new(ShoeSharkTokenShoeSharkTokenBurned)
	if err := _ShoeSharkToken.contract.UnpackLog(event, "ShoeSharkToken_Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkTokenShoeSharkTokenMintedIterator is returned from FilterShoeSharkTokenMinted and is used to iterate over the raw logs and unpacked data for ShoeSharkTokenMinted events raised by the ShoeSharkToken contract.
type ShoeSharkTokenShoeSharkTokenMintedIterator struct {
	Event *ShoeSharkTokenShoeSharkTokenMinted // Event containing the contract specifics and raw log

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
func (it *ShoeSharkTokenShoeSharkTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkTokenShoeSharkTokenMinted)
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
		it.Event = new(ShoeSharkTokenShoeSharkTokenMinted)
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
func (it *ShoeSharkTokenShoeSharkTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkTokenShoeSharkTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkTokenShoeSharkTokenMinted represents a ShoeSharkTokenMinted event raised by the ShoeSharkToken contract.
type ShoeSharkTokenShoeSharkTokenMinted struct {
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkTokenMinted is a free log retrieval operation binding the contract event 0xa77f95bc932b4b33dcbc835a7a5fca2f8fabced6af929af36e91d1931512b965.
//
// Solidity: event ShoeSharkToken_Minted(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) FilterShoeSharkTokenMinted(opts *bind.FilterOpts, player []common.Address) (*ShoeSharkTokenShoeSharkTokenMintedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.FilterLogs(opts, "ShoeSharkToken_Minted", playerRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenShoeSharkTokenMintedIterator{contract: _ShoeSharkToken.contract, event: "ShoeSharkToken_Minted", logs: logs, sub: sub}, nil
}

// WatchShoeSharkTokenMinted is a free log subscription operation binding the contract event 0xa77f95bc932b4b33dcbc835a7a5fca2f8fabced6af929af36e91d1931512b965.
//
// Solidity: event ShoeSharkToken_Minted(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) WatchShoeSharkTokenMinted(opts *bind.WatchOpts, sink chan<- *ShoeSharkTokenShoeSharkTokenMinted, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.WatchLogs(opts, "ShoeSharkToken_Minted", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkTokenShoeSharkTokenMinted)
				if err := _ShoeSharkToken.contract.UnpackLog(event, "ShoeSharkToken_Minted", log); err != nil {
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

// ParseShoeSharkTokenMinted is a log parse operation binding the contract event 0xa77f95bc932b4b33dcbc835a7a5fca2f8fabced6af929af36e91d1931512b965.
//
// Solidity: event ShoeSharkToken_Minted(address indexed player)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) ParseShoeSharkTokenMinted(log types.Log) (*ShoeSharkTokenShoeSharkTokenMinted, error) {
	event := new(ShoeSharkTokenShoeSharkTokenMinted)
	if err := _ShoeSharkToken.contract.UnpackLog(event, "ShoeSharkToken_Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShoeSharkToken contract.
type ShoeSharkTokenTransferIterator struct {
	Event *ShoeSharkTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ShoeSharkTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkTokenTransfer)
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
		it.Event = new(ShoeSharkTokenTransfer)
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
func (it *ShoeSharkTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkTokenTransfer represents a Transfer event raised by the ShoeSharkToken contract.
type ShoeSharkTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShoeSharkTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkTokenTransferIterator{contract: _ShoeSharkToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ShoeSharkToken *ShoeSharkTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShoeSharkTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShoeSharkToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkTokenTransfer)
				if err := _ShoeSharkToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ShoeSharkToken *ShoeSharkTokenFilterer) ParseTransfer(log types.Log) (*ShoeSharkTokenTransfer, error) {
	event := new(ShoeSharkTokenTransfer)
	if err := _ShoeSharkToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
