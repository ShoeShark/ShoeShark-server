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

// ShoeSharkNftMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type ShoeSharkNftMarketOrder struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}

// ShoeSharkNftMarketMetaData contains all meta data concerning the ShoeSharkNftMarket contract.
var ShoeSharkNftMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_SST\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_NFT\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NFT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buy\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changePrice\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllNFTs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structShoeSharkNftMarket.Order[]\",\"components\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMyNFTs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structShoeSharkNftMarket.Order[]\",\"components\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isListed\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_orders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_ordersOfId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenUri\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_tokenIdToOrderIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toUint256\",\"inputs\":[{\"name\":\"_bytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_start\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ShoeSharkNftMarket_Deal\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNftMarket_NewOrder\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNftMarket_OrderCancelled\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNftMarket_PriceChanged\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"previousPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__ToUint256__OutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__ToUint256__OverFlow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__buy__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__cancelOrder__NotSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__changePrice__NotSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__constructor__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNftMarket__onERC721Received__PriceMustBeGreaterThanZero\",\"inputs\":[]}]",
}

// ShoeSharkNftMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use ShoeSharkNftMarketMetaData.ABI instead.
var ShoeSharkNftMarketABI = ShoeSharkNftMarketMetaData.ABI

// ShoeSharkNftMarket is an auto generated Go binding around an Ethereum contract.
type ShoeSharkNftMarket struct {
	ShoeSharkNftMarketCaller     // Read-only binding to the contract
	ShoeSharkNftMarketTransactor // Write-only binding to the contract
	ShoeSharkNftMarketFilterer   // Log filterer for contract events
}

// ShoeSharkNftMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShoeSharkNftMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShoeSharkNftMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShoeSharkNftMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShoeSharkNftMarketSession struct {
	Contract     *ShoeSharkNftMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ShoeSharkNftMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShoeSharkNftMarketCallerSession struct {
	Contract *ShoeSharkNftMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ShoeSharkNftMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShoeSharkNftMarketTransactorSession struct {
	Contract     *ShoeSharkNftMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ShoeSharkNftMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShoeSharkNftMarketRaw struct {
	Contract *ShoeSharkNftMarket // Generic contract binding to access the raw methods on
}

// ShoeSharkNftMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShoeSharkNftMarketCallerRaw struct {
	Contract *ShoeSharkNftMarketCaller // Generic read-only contract binding to access the raw methods on
}

// ShoeSharkNftMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShoeSharkNftMarketTransactorRaw struct {
	Contract *ShoeSharkNftMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShoeSharkNftMarket creates a new instance of ShoeSharkNftMarket, bound to a specific deployed contract.
func NewShoeSharkNftMarket(address common.Address, backend bind.ContractBackend) (*ShoeSharkNftMarket, error) {
	contract, err := bindShoeSharkNftMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarket{ShoeSharkNftMarketCaller: ShoeSharkNftMarketCaller{contract: contract}, ShoeSharkNftMarketTransactor: ShoeSharkNftMarketTransactor{contract: contract}, ShoeSharkNftMarketFilterer: ShoeSharkNftMarketFilterer{contract: contract}}, nil
}

// NewShoeSharkNftMarketCaller creates a new read-only instance of ShoeSharkNftMarket, bound to a specific deployed contract.
func NewShoeSharkNftMarketCaller(address common.Address, caller bind.ContractCaller) (*ShoeSharkNftMarketCaller, error) {
	contract, err := bindShoeSharkNftMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketCaller{contract: contract}, nil
}

// NewShoeSharkNftMarketTransactor creates a new write-only instance of ShoeSharkNftMarket, bound to a specific deployed contract.
func NewShoeSharkNftMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*ShoeSharkNftMarketTransactor, error) {
	contract, err := bindShoeSharkNftMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketTransactor{contract: contract}, nil
}

// NewShoeSharkNftMarketFilterer creates a new log filterer instance of ShoeSharkNftMarket, bound to a specific deployed contract.
func NewShoeSharkNftMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*ShoeSharkNftMarketFilterer, error) {
	contract, err := bindShoeSharkNftMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketFilterer{contract: contract}, nil
}

// bindShoeSharkNftMarket binds a generic wrapper to an already deployed contract.
func bindShoeSharkNftMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShoeSharkNftMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkNftMarket.Contract.ShoeSharkNftMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.ShoeSharkNftMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.ShoeSharkNftMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkNftMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.contract.Transact(opts, method, params...)
}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) NFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "NFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) NFT() (common.Address, error) {
	return _ShoeSharkNftMarket.Contract.NFT(&_ShoeSharkNftMarket.CallOpts)
}

// NFT is a free data retrieval call binding the contract method 0x7c0b8de2.
//
// Solidity: function NFT() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) NFT() (common.Address, error) {
	return _ShoeSharkNftMarket.Contract.NFT(&_ShoeSharkNftMarket.CallOpts)
}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) SST(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "SST")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) SST() (common.Address, error) {
	return _ShoeSharkNftMarket.Contract.SST(&_ShoeSharkNftMarket.CallOpts)
}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) SST() (common.Address, error) {
	return _ShoeSharkNftMarket.Contract.SST(&_ShoeSharkNftMarket.CallOpts)
}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) GetAllNFTs(opts *bind.CallOpts) ([]ShoeSharkNftMarketOrder, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "getAllNFTs")

	if err != nil {
		return *new([]ShoeSharkNftMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ShoeSharkNftMarketOrder)).(*[]ShoeSharkNftMarketOrder)

	return out0, err

}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) GetAllNFTs() ([]ShoeSharkNftMarketOrder, error) {
	return _ShoeSharkNftMarket.Contract.GetAllNFTs(&_ShoeSharkNftMarket.CallOpts)
}

// GetAllNFTs is a free data retrieval call binding the contract method 0xe0391b09.
//
// Solidity: function getAllNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) GetAllNFTs() ([]ShoeSharkNftMarketOrder, error) {
	return _ShoeSharkNftMarket.Contract.GetAllNFTs(&_ShoeSharkNftMarket.CallOpts)
}

// GetMyNFTs is a free data retrieval call binding the contract method 0x629cb2e4.
//
// Solidity: function getMyNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) GetMyNFTs(opts *bind.CallOpts) ([]ShoeSharkNftMarketOrder, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "getMyNFTs")

	if err != nil {
		return *new([]ShoeSharkNftMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]ShoeSharkNftMarketOrder)).(*[]ShoeSharkNftMarketOrder)

	return out0, err

}

// GetMyNFTs is a free data retrieval call binding the contract method 0x629cb2e4.
//
// Solidity: function getMyNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) GetMyNFTs() ([]ShoeSharkNftMarketOrder, error) {
	return _ShoeSharkNftMarket.Contract.GetMyNFTs(&_ShoeSharkNftMarket.CallOpts)
}

// GetMyNFTs is a free data retrieval call binding the contract method 0x629cb2e4.
//
// Solidity: function getMyNFTs() view returns((uint256,uint256,address,string)[])
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) GetMyNFTs() ([]ShoeSharkNftMarketOrder, error) {
	return _ShoeSharkNftMarket.Contract.GetMyNFTs(&_ShoeSharkNftMarket.CallOpts)
}

// GetOrderLength is a free data retrieval call binding the contract method 0x9d4971b7.
//
// Solidity: function getOrderLength() view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) GetOrderLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "getOrderLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderLength is a free data retrieval call binding the contract method 0x9d4971b7.
//
// Solidity: function getOrderLength() view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) GetOrderLength() (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.GetOrderLength(&_ShoeSharkNftMarket.CallOpts)
}

// GetOrderLength is a free data retrieval call binding the contract method 0x9d4971b7.
//
// Solidity: function getOrderLength() view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) GetOrderLength() (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.GetOrderLength(&_ShoeSharkNftMarket.CallOpts)
}

// IsListed is a free data retrieval call binding the contract method 0xfcce4883.
//
// Solidity: function isListed(uint256 _tokenId) view returns(bool)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) IsListed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "isListed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsListed is a free data retrieval call binding the contract method 0xfcce4883.
//
// Solidity: function isListed(uint256 _tokenId) view returns(bool)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) IsListed(_tokenId *big.Int) (bool, error) {
	return _ShoeSharkNftMarket.Contract.IsListed(&_ShoeSharkNftMarket.CallOpts, _tokenId)
}

// IsListed is a free data retrieval call binding the contract method 0xfcce4883.
//
// Solidity: function isListed(uint256 _tokenId) view returns(bool)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) IsListed(_tokenId *big.Int) (bool, error) {
	return _ShoeSharkNftMarket.Contract.IsListed(&_ShoeSharkNftMarket.CallOpts, _tokenId)
}

// SOrders is a free data retrieval call binding the contract method 0xa6f3c50f.
//
// Solidity: function s_orders(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) SOrders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "s_orders", arg0)

	outstruct := new(struct {
		TokenId  *big.Int
		Price    *big.Int
		Seller   common.Address
		TokenUri string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenUri = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// SOrders is a free data retrieval call binding the contract method 0xa6f3c50f.
//
// Solidity: function s_orders(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) SOrders(arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	return _ShoeSharkNftMarket.Contract.SOrders(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// SOrders is a free data retrieval call binding the contract method 0xa6f3c50f.
//
// Solidity: function s_orders(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) SOrders(arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	return _ShoeSharkNftMarket.Contract.SOrders(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// SOrdersOfId is a free data retrieval call binding the contract method 0x19ebc23e.
//
// Solidity: function s_ordersOfId(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) SOrdersOfId(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "s_ordersOfId", arg0)

	outstruct := new(struct {
		TokenId  *big.Int
		Price    *big.Int
		Seller   common.Address
		TokenUri string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenUri = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// SOrdersOfId is a free data retrieval call binding the contract method 0x19ebc23e.
//
// Solidity: function s_ordersOfId(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) SOrdersOfId(arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	return _ShoeSharkNftMarket.Contract.SOrdersOfId(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// SOrdersOfId is a free data retrieval call binding the contract method 0x19ebc23e.
//
// Solidity: function s_ordersOfId(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, string tokenUri)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) SOrdersOfId(arg0 *big.Int) (struct {
	TokenId  *big.Int
	Price    *big.Int
	Seller   common.Address
	TokenUri string
}, error) {
	return _ShoeSharkNftMarket.Contract.SOrdersOfId(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// STokenIdToOrderIndex is a free data retrieval call binding the contract method 0xb0f2d808.
//
// Solidity: function s_tokenIdToOrderIndex(uint256 ) view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) STokenIdToOrderIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "s_tokenIdToOrderIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STokenIdToOrderIndex is a free data retrieval call binding the contract method 0xb0f2d808.
//
// Solidity: function s_tokenIdToOrderIndex(uint256 ) view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) STokenIdToOrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.STokenIdToOrderIndex(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// STokenIdToOrderIndex is a free data retrieval call binding the contract method 0xb0f2d808.
//
// Solidity: function s_tokenIdToOrderIndex(uint256 ) view returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) STokenIdToOrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.STokenIdToOrderIndex(&_ShoeSharkNftMarket.CallOpts, arg0)
}

// ToUint256 is a free data retrieval call binding the contract method 0xb5cdf924.
//
// Solidity: function toUint256(bytes _bytes, uint256 _start) pure returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCaller) ToUint256(opts *bind.CallOpts, _bytes []byte, _start *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNftMarket.contract.Call(opts, &out, "toUint256", _bytes, _start)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint256 is a free data retrieval call binding the contract method 0xb5cdf924.
//
// Solidity: function toUint256(bytes _bytes, uint256 _start) pure returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) ToUint256(_bytes []byte, _start *big.Int) (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.ToUint256(&_ShoeSharkNftMarket.CallOpts, _bytes, _start)
}

// ToUint256 is a free data retrieval call binding the contract method 0xb5cdf924.
//
// Solidity: function toUint256(bytes _bytes, uint256 _start) pure returns(uint256)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketCallerSession) ToUint256(_bytes []byte, _start *big.Int) (*big.Int, error) {
	return _ShoeSharkNftMarket.Contract.ToUint256(&_ShoeSharkNftMarket.CallOpts, _bytes, _start)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactor) Buy(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.contract.Transact(opts, "buy", _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.Buy(&_ShoeSharkNftMarket.TransactOpts, _tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorSession) Buy(_tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.Buy(&_ShoeSharkNftMarket.TransactOpts, _tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactor) CancelOrder(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.contract.Transact(opts, "cancelOrder", _tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) CancelOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.CancelOrder(&_ShoeSharkNftMarket.TransactOpts, _tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorSession) CancelOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.CancelOrder(&_ShoeSharkNftMarket.TransactOpts, _tokenId)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 _tokenId, uint256 _price) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactor) ChangePrice(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.contract.Transact(opts, "changePrice", _tokenId, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 _tokenId, uint256 _price) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) ChangePrice(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.ChangePrice(&_ShoeSharkNftMarket.TransactOpts, _tokenId, _price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 _tokenId, uint256 _price) returns()
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorSession) ChangePrice(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.ChangePrice(&_ShoeSharkNftMarket.TransactOpts, _tokenId, _price)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.OnERC721Received(&_ShoeSharkNftMarket.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNftMarket.Contract.OnERC721Received(&_ShoeSharkNftMarket.TransactOpts, operator, from, tokenId, data)
}

// ShoeSharkNftMarketShoeSharkNftMarketDealIterator is returned from FilterShoeSharkNftMarketDeal and is used to iterate over the raw logs and unpacked data for ShoeSharkNftMarketDeal events raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketDealIterator struct {
	Event *ShoeSharkNftMarketShoeSharkNftMarketDeal // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftMarketShoeSharkNftMarketDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketDeal)
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
		it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketDeal)
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
func (it *ShoeSharkNftMarketShoeSharkNftMarketDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftMarketShoeSharkNftMarketDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftMarketShoeSharkNftMarketDeal represents a ShoeSharkNftMarketDeal event raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketDeal struct {
	Seller  common.Address
	Buyer   common.Address
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftMarketDeal is a free log retrieval operation binding the contract event 0x4fea1276970c7033d8991207c04fc931d2c3d7a676f4bc25de79f045007f078b.
//
// Solidity: event ShoeSharkNftMarket_Deal(address indexed seller, address indexed buyer, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) FilterShoeSharkNftMarketDeal(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, tokenId []*big.Int) (*ShoeSharkNftMarketShoeSharkNftMarketDealIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.FilterLogs(opts, "ShoeSharkNftMarket_Deal", sellerRule, buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketShoeSharkNftMarketDealIterator{contract: _ShoeSharkNftMarket.contract, event: "ShoeSharkNftMarket_Deal", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftMarketDeal is a free log subscription operation binding the contract event 0x4fea1276970c7033d8991207c04fc931d2c3d7a676f4bc25de79f045007f078b.
//
// Solidity: event ShoeSharkNftMarket_Deal(address indexed seller, address indexed buyer, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) WatchShoeSharkNftMarketDeal(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftMarketShoeSharkNftMarketDeal, seller []common.Address, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.WatchLogs(opts, "ShoeSharkNftMarket_Deal", sellerRule, buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftMarketShoeSharkNftMarketDeal)
				if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_Deal", log); err != nil {
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

// ParseShoeSharkNftMarketDeal is a log parse operation binding the contract event 0x4fea1276970c7033d8991207c04fc931d2c3d7a676f4bc25de79f045007f078b.
//
// Solidity: event ShoeSharkNftMarket_Deal(address indexed seller, address indexed buyer, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) ParseShoeSharkNftMarketDeal(log types.Log) (*ShoeSharkNftMarketShoeSharkNftMarketDeal, error) {
	event := new(ShoeSharkNftMarketShoeSharkNftMarketDeal)
	if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_Deal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator is returned from FilterShoeSharkNftMarketNewOrder and is used to iterate over the raw logs and unpacked data for ShoeSharkNftMarketNewOrder events raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator struct {
	Event *ShoeSharkNftMarketShoeSharkNftMarketNewOrder // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketNewOrder)
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
		it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketNewOrder)
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
func (it *ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftMarketShoeSharkNftMarketNewOrder represents a ShoeSharkNftMarketNewOrder event raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketNewOrder struct {
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftMarketNewOrder is a free log retrieval operation binding the contract event 0x8abf3892f00c603f3ffbf6555afe57f893e67c9fc751c0dc571b27530eeebd6f.
//
// Solidity: event ShoeSharkNftMarket_NewOrder(address indexed seller, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) FilterShoeSharkNftMarketNewOrder(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.FilterLogs(opts, "ShoeSharkNftMarket_NewOrder", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketShoeSharkNftMarketNewOrderIterator{contract: _ShoeSharkNftMarket.contract, event: "ShoeSharkNftMarket_NewOrder", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftMarketNewOrder is a free log subscription operation binding the contract event 0x8abf3892f00c603f3ffbf6555afe57f893e67c9fc751c0dc571b27530eeebd6f.
//
// Solidity: event ShoeSharkNftMarket_NewOrder(address indexed seller, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) WatchShoeSharkNftMarketNewOrder(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftMarketShoeSharkNftMarketNewOrder, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.WatchLogs(opts, "ShoeSharkNftMarket_NewOrder", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftMarketShoeSharkNftMarketNewOrder)
				if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_NewOrder", log); err != nil {
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

// ParseShoeSharkNftMarketNewOrder is a log parse operation binding the contract event 0x8abf3892f00c603f3ffbf6555afe57f893e67c9fc751c0dc571b27530eeebd6f.
//
// Solidity: event ShoeSharkNftMarket_NewOrder(address indexed seller, uint256 indexed tokenId, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) ParseShoeSharkNftMarketNewOrder(log types.Log) (*ShoeSharkNftMarketShoeSharkNftMarketNewOrder, error) {
	event := new(ShoeSharkNftMarketShoeSharkNftMarketNewOrder)
	if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_NewOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator is returned from FilterShoeSharkNftMarketOrderCancelled and is used to iterate over the raw logs and unpacked data for ShoeSharkNftMarketOrderCancelled events raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator struct {
	Event *ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled)
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
		it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled)
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
func (it *ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled represents a ShoeSharkNftMarketOrderCancelled event raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled struct {
	Seller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftMarketOrderCancelled is a free log retrieval operation binding the contract event 0x7069e1cf2d809da72f139aa524e4ec923bb51f62aa877dc7ff8e313a75f85238.
//
// Solidity: event ShoeSharkNftMarket_OrderCancelled(address indexed seller, uint256 indexed tokenId)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) FilterShoeSharkNftMarketOrderCancelled(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.FilterLogs(opts, "ShoeSharkNftMarket_OrderCancelled", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketShoeSharkNftMarketOrderCancelledIterator{contract: _ShoeSharkNftMarket.contract, event: "ShoeSharkNftMarket_OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftMarketOrderCancelled is a free log subscription operation binding the contract event 0x7069e1cf2d809da72f139aa524e4ec923bb51f62aa877dc7ff8e313a75f85238.
//
// Solidity: event ShoeSharkNftMarket_OrderCancelled(address indexed seller, uint256 indexed tokenId)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) WatchShoeSharkNftMarketOrderCancelled(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.WatchLogs(opts, "ShoeSharkNftMarket_OrderCancelled", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled)
				if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_OrderCancelled", log); err != nil {
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

// ParseShoeSharkNftMarketOrderCancelled is a log parse operation binding the contract event 0x7069e1cf2d809da72f139aa524e4ec923bb51f62aa877dc7ff8e313a75f85238.
//
// Solidity: event ShoeSharkNftMarket_OrderCancelled(address indexed seller, uint256 indexed tokenId)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) ParseShoeSharkNftMarketOrderCancelled(log types.Log) (*ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled, error) {
	event := new(ShoeSharkNftMarketShoeSharkNftMarketOrderCancelled)
	if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator is returned from FilterShoeSharkNftMarketPriceChanged and is used to iterate over the raw logs and unpacked data for ShoeSharkNftMarketPriceChanged events raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator struct {
	Event *ShoeSharkNftMarketShoeSharkNftMarketPriceChanged // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketPriceChanged)
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
		it.Event = new(ShoeSharkNftMarketShoeSharkNftMarketPriceChanged)
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
func (it *ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftMarketShoeSharkNftMarketPriceChanged represents a ShoeSharkNftMarketPriceChanged event raised by the ShoeSharkNftMarket contract.
type ShoeSharkNftMarketShoeSharkNftMarketPriceChanged struct {
	Seller        common.Address
	TokenId       *big.Int
	PreviousPrice *big.Int
	Price         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftMarketPriceChanged is a free log retrieval operation binding the contract event 0x407ca485131450a05093468b333f5cdb756b470bfb0ad00fdeba1a39342663f3.
//
// Solidity: event ShoeSharkNftMarket_PriceChanged(address indexed seller, uint256 indexed tokenId, uint256 previousPrice, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) FilterShoeSharkNftMarketPriceChanged(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.FilterLogs(opts, "ShoeSharkNftMarket_PriceChanged", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMarketShoeSharkNftMarketPriceChangedIterator{contract: _ShoeSharkNftMarket.contract, event: "ShoeSharkNftMarket_PriceChanged", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftMarketPriceChanged is a free log subscription operation binding the contract event 0x407ca485131450a05093468b333f5cdb756b470bfb0ad00fdeba1a39342663f3.
//
// Solidity: event ShoeSharkNftMarket_PriceChanged(address indexed seller, uint256 indexed tokenId, uint256 previousPrice, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) WatchShoeSharkNftMarketPriceChanged(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftMarketShoeSharkNftMarketPriceChanged, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNftMarket.contract.WatchLogs(opts, "ShoeSharkNftMarket_PriceChanged", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftMarketShoeSharkNftMarketPriceChanged)
				if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_PriceChanged", log); err != nil {
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

// ParseShoeSharkNftMarketPriceChanged is a log parse operation binding the contract event 0x407ca485131450a05093468b333f5cdb756b470bfb0ad00fdeba1a39342663f3.
//
// Solidity: event ShoeSharkNftMarket_PriceChanged(address indexed seller, uint256 indexed tokenId, uint256 previousPrice, uint256 price)
func (_ShoeSharkNftMarket *ShoeSharkNftMarketFilterer) ParseShoeSharkNftMarketPriceChanged(log types.Log) (*ShoeSharkNftMarketShoeSharkNftMarketPriceChanged, error) {
	event := new(ShoeSharkNftMarketShoeSharkNftMarketPriceChanged)
	if err := _ShoeSharkNftMarket.contract.UnpackLog(event, "ShoeSharkNftMarket_PriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
