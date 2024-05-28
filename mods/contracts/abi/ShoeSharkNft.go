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

// ShoeSharkNftMetaData contains all meta data concerning the ShoeSharkNft contract.
var ShoeSharkNftMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"merkleroot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"metadataUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subscriptionId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mintMaxTotal\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"vrfCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"keyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataUri\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintWhiteList\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rawFulfillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_HasMinted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_IsMinting\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_Numbers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_RemainingNumbers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_RequestIdToTokenId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_Root\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsMinting\",\"inputs\":[{\"name\":\"isMinting\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataUri\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoot\",\"inputs\":[{\"name\":\"newRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenByIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNft_NftBurned\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNft_NftMinted\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNft_Withdraw\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkNft_tokenIdtoTokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"metadateIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721OutOfBoundsIndex\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyCoordinatorCanFulfill\",\"inputs\":[{\"name\":\"have\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"want\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__burn__OnlyBurnByTokenOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__byOwner__OnlyOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mintWhitelist__HasMinted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mintWhitelist__InvalidMintCost\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mintWhitelist__NotInWhitelist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mintWhitelist__TransferFaild\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mint__MintingIsNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mint__NoMoreNft\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__mint__TokenIdOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__tokenURI__TokenUriNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__withdraw__BalanceIsZeroWhenWithdrawing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkNft__withdraw__WithdrawFailed\",\"inputs\":[]}]",
}

// ShoeSharkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use ShoeSharkNftMetaData.ABI instead.
var ShoeSharkNftABI = ShoeSharkNftMetaData.ABI

// ShoeSharkNft is an auto generated Go binding around an Ethereum contract.
type ShoeSharkNft struct {
	ShoeSharkNftCaller     // Read-only binding to the contract
	ShoeSharkNftTransactor // Write-only binding to the contract
	ShoeSharkNftFilterer   // Log filterer for contract events
}

// ShoeSharkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShoeSharkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShoeSharkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShoeSharkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShoeSharkNftSession struct {
	Contract     *ShoeSharkNft     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShoeSharkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShoeSharkNftCallerSession struct {
	Contract *ShoeSharkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ShoeSharkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShoeSharkNftTransactorSession struct {
	Contract     *ShoeSharkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ShoeSharkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShoeSharkNftRaw struct {
	Contract *ShoeSharkNft // Generic contract binding to access the raw methods on
}

// ShoeSharkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShoeSharkNftCallerRaw struct {
	Contract *ShoeSharkNftCaller // Generic read-only contract binding to access the raw methods on
}

// ShoeSharkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShoeSharkNftTransactorRaw struct {
	Contract *ShoeSharkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShoeSharkNft creates a new instance of ShoeSharkNft, bound to a specific deployed contract.
func NewShoeSharkNft(address common.Address, backend bind.ContractBackend) (*ShoeSharkNft, error) {
	contract, err := bindShoeSharkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNft{ShoeSharkNftCaller: ShoeSharkNftCaller{contract: contract}, ShoeSharkNftTransactor: ShoeSharkNftTransactor{contract: contract}, ShoeSharkNftFilterer: ShoeSharkNftFilterer{contract: contract}}, nil
}

// NewShoeSharkNftCaller creates a new read-only instance of ShoeSharkNft, bound to a specific deployed contract.
func NewShoeSharkNftCaller(address common.Address, caller bind.ContractCaller) (*ShoeSharkNftCaller, error) {
	contract, err := bindShoeSharkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftCaller{contract: contract}, nil
}

// NewShoeSharkNftTransactor creates a new write-only instance of ShoeSharkNft, bound to a specific deployed contract.
func NewShoeSharkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*ShoeSharkNftTransactor, error) {
	contract, err := bindShoeSharkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftTransactor{contract: contract}, nil
}

// NewShoeSharkNftFilterer creates a new log filterer instance of ShoeSharkNft, bound to a specific deployed contract.
func NewShoeSharkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*ShoeSharkNftFilterer, error) {
	contract, err := bindShoeSharkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftFilterer{contract: contract}, nil
}

// bindShoeSharkNft binds a generic wrapper to an already deployed contract.
func bindShoeSharkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShoeSharkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkNft *ShoeSharkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkNft.Contract.ShoeSharkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkNft *ShoeSharkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.ShoeSharkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkNft *ShoeSharkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.ShoeSharkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkNft *ShoeSharkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkNft *ShoeSharkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkNft *ShoeSharkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ShoeSharkNft.Contract.BalanceOf(&_ShoeSharkNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ShoeSharkNft.Contract.BalanceOf(&_ShoeSharkNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ShoeSharkNft.Contract.GetApproved(&_ShoeSharkNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ShoeSharkNft.Contract.GetApproved(&_ShoeSharkNft.CallOpts, tokenId)
}

// GetMetadataUri is a free data retrieval call binding the contract method 0xa0bd5c40.
//
// Solidity: function getMetadataUri() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCaller) GetMetadataUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "getMetadataUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataUri is a free data retrieval call binding the contract method 0xa0bd5c40.
//
// Solidity: function getMetadataUri() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftSession) GetMetadataUri() (string, error) {
	return _ShoeSharkNft.Contract.GetMetadataUri(&_ShoeSharkNft.CallOpts)
}

// GetMetadataUri is a free data retrieval call binding the contract method 0xa0bd5c40.
//
// Solidity: function getMetadataUri() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) GetMetadataUri() (string, error) {
	return _ShoeSharkNft.Contract.GetMetadataUri(&_ShoeSharkNft.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftSession) GetOwner() (common.Address, error) {
	return _ShoeSharkNft.Contract.GetOwner(&_ShoeSharkNft.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) GetOwner() (common.Address, error) {
	return _ShoeSharkNft.Contract.GetOwner(&_ShoeSharkNft.CallOpts)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) GetTokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "getTokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) GetTokenCounter() (*big.Int, error) {
	return _ShoeSharkNft.Contract.GetTokenCounter(&_ShoeSharkNft.CallOpts)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) GetTokenCounter() (*big.Int, error) {
	return _ShoeSharkNft.Contract.GetTokenCounter(&_ShoeSharkNft.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ShoeSharkNft.Contract.IsApprovedForAll(&_ShoeSharkNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ShoeSharkNft.Contract.IsApprovedForAll(&_ShoeSharkNft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftSession) Name() (string, error) {
	return _ShoeSharkNft.Contract.Name(&_ShoeSharkNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) Name() (string, error) {
	return _ShoeSharkNft.Contract.Name(&_ShoeSharkNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftSession) Owner() (common.Address, error) {
	return _ShoeSharkNft.Contract.Owner(&_ShoeSharkNft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) Owner() (common.Address, error) {
	return _ShoeSharkNft.Contract.Owner(&_ShoeSharkNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ShoeSharkNft.Contract.OwnerOf(&_ShoeSharkNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ShoeSharkNft.Contract.OwnerOf(&_ShoeSharkNft.CallOpts, tokenId)
}

// SHasMinted is a free data retrieval call binding the contract method 0x4d82503c.
//
// Solidity: function s_HasMinted(address ) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCaller) SHasMinted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_HasMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SHasMinted is a free data retrieval call binding the contract method 0x4d82503c.
//
// Solidity: function s_HasMinted(address ) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftSession) SHasMinted(arg0 common.Address) (bool, error) {
	return _ShoeSharkNft.Contract.SHasMinted(&_ShoeSharkNft.CallOpts, arg0)
}

// SHasMinted is a free data retrieval call binding the contract method 0x4d82503c.
//
// Solidity: function s_HasMinted(address ) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SHasMinted(arg0 common.Address) (bool, error) {
	return _ShoeSharkNft.Contract.SHasMinted(&_ShoeSharkNft.CallOpts, arg0)
}

// SIsMinting is a free data retrieval call binding the contract method 0x6054930d.
//
// Solidity: function s_IsMinting() view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCaller) SIsMinting(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_IsMinting")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SIsMinting is a free data retrieval call binding the contract method 0x6054930d.
//
// Solidity: function s_IsMinting() view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftSession) SIsMinting() (bool, error) {
	return _ShoeSharkNft.Contract.SIsMinting(&_ShoeSharkNft.CallOpts)
}

// SIsMinting is a free data retrieval call binding the contract method 0x6054930d.
//
// Solidity: function s_IsMinting() view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SIsMinting() (bool, error) {
	return _ShoeSharkNft.Contract.SIsMinting(&_ShoeSharkNft.CallOpts)
}

// SNumbers is a free data retrieval call binding the contract method 0xa072b445.
//
// Solidity: function s_Numbers(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) SNumbers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_Numbers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SNumbers is a free data retrieval call binding the contract method 0xa072b445.
//
// Solidity: function s_Numbers(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) SNumbers(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.SNumbers(&_ShoeSharkNft.CallOpts, arg0)
}

// SNumbers is a free data retrieval call binding the contract method 0xa072b445.
//
// Solidity: function s_Numbers(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SNumbers(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.SNumbers(&_ShoeSharkNft.CallOpts, arg0)
}

// SRemainingNumbers is a free data retrieval call binding the contract method 0xb545b319.
//
// Solidity: function s_RemainingNumbers() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) SRemainingNumbers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_RemainingNumbers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRemainingNumbers is a free data retrieval call binding the contract method 0xb545b319.
//
// Solidity: function s_RemainingNumbers() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) SRemainingNumbers() (*big.Int, error) {
	return _ShoeSharkNft.Contract.SRemainingNumbers(&_ShoeSharkNft.CallOpts)
}

// SRemainingNumbers is a free data retrieval call binding the contract method 0xb545b319.
//
// Solidity: function s_RemainingNumbers() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SRemainingNumbers() (*big.Int, error) {
	return _ShoeSharkNft.Contract.SRemainingNumbers(&_ShoeSharkNft.CallOpts)
}

// SRequestIdToTokenId is a free data retrieval call binding the contract method 0x0717300d.
//
// Solidity: function s_RequestIdToTokenId(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) SRequestIdToTokenId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_RequestIdToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRequestIdToTokenId is a free data retrieval call binding the contract method 0x0717300d.
//
// Solidity: function s_RequestIdToTokenId(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) SRequestIdToTokenId(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.SRequestIdToTokenId(&_ShoeSharkNft.CallOpts, arg0)
}

// SRequestIdToTokenId is a free data retrieval call binding the contract method 0x0717300d.
//
// Solidity: function s_RequestIdToTokenId(uint256 ) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SRequestIdToTokenId(arg0 *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.SRequestIdToTokenId(&_ShoeSharkNft.CallOpts, arg0)
}

// SRoot is a free data retrieval call binding the contract method 0x1dc81078.
//
// Solidity: function s_Root() view returns(bytes32)
func (_ShoeSharkNft *ShoeSharkNftCaller) SRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "s_Root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SRoot is a free data retrieval call binding the contract method 0x1dc81078.
//
// Solidity: function s_Root() view returns(bytes32)
func (_ShoeSharkNft *ShoeSharkNftSession) SRoot() ([32]byte, error) {
	return _ShoeSharkNft.Contract.SRoot(&_ShoeSharkNft.CallOpts)
}

// SRoot is a free data retrieval call binding the contract method 0x1dc81078.
//
// Solidity: function s_Root() view returns(bytes32)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SRoot() ([32]byte, error) {
	return _ShoeSharkNft.Contract.SRoot(&_ShoeSharkNft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShoeSharkNft.Contract.SupportsInterface(&_ShoeSharkNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShoeSharkNft.Contract.SupportsInterface(&_ShoeSharkNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftSession) Symbol() (string, error) {
	return _ShoeSharkNft.Contract.Symbol(&_ShoeSharkNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) Symbol() (string, error) {
	return _ShoeSharkNft.Contract.Symbol(&_ShoeSharkNft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.TokenByIndex(&_ShoeSharkNft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.TokenByIndex(&_ShoeSharkNft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.TokenOfOwnerByIndex(&_ShoeSharkNft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ShoeSharkNft.Contract.TokenOfOwnerByIndex(&_ShoeSharkNft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShoeSharkNft *ShoeSharkNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ShoeSharkNft.Contract.TokenURI(&_ShoeSharkNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ShoeSharkNft.Contract.TokenURI(&_ShoeSharkNft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkNft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftSession) TotalSupply() (*big.Int, error) {
	return _ShoeSharkNft.Contract.TotalSupply(&_ShoeSharkNft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShoeSharkNft *ShoeSharkNftCallerSession) TotalSupply() (*big.Int, error) {
	return _ShoeSharkNft.Contract.TotalSupply(&_ShoeSharkNft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Approve(&_ShoeSharkNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Approve(&_ShoeSharkNft.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Burn(&_ShoeSharkNft.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Burn(&_ShoeSharkNft.TransactOpts, tokenId)
}

// MintWhiteList is a paid mutator transaction binding the contract method 0x45e081b7.
//
// Solidity: function mintWhiteList(address player, bytes32[] proof) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) MintWhiteList(opts *bind.TransactOpts, player common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "mintWhiteList", player, proof)
}

// MintWhiteList is a paid mutator transaction binding the contract method 0x45e081b7.
//
// Solidity: function mintWhiteList(address player, bytes32[] proof) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) MintWhiteList(player common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.MintWhiteList(&_ShoeSharkNft.TransactOpts, player, proof)
}

// MintWhiteList is a paid mutator transaction binding the contract method 0x45e081b7.
//
// Solidity: function mintWhiteList(address player, bytes32[] proof) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) MintWhiteList(player common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.MintWhiteList(&_ShoeSharkNft.TransactOpts, player, proof)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.RawFulfillRandomWords(&_ShoeSharkNft.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.RawFulfillRandomWords(&_ShoeSharkNft.TransactOpts, requestId, randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkNft *ShoeSharkNftSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.RenounceOwnership(&_ShoeSharkNft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.RenounceOwnership(&_ShoeSharkNft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SafeTransferFrom(&_ShoeSharkNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SafeTransferFrom(&_ShoeSharkNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SafeTransferFrom0(&_ShoeSharkNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SafeTransferFrom0(&_ShoeSharkNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetApprovalForAll(&_ShoeSharkNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetApprovalForAll(&_ShoeSharkNft.TransactOpts, operator, approved)
}

// SetIsMinting is a paid mutator transaction binding the contract method 0x7b93d1e5.
//
// Solidity: function setIsMinting(bool isMinting) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SetIsMinting(opts *bind.TransactOpts, isMinting bool) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "setIsMinting", isMinting)
}

// SetIsMinting is a paid mutator transaction binding the contract method 0x7b93d1e5.
//
// Solidity: function setIsMinting(bool isMinting) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SetIsMinting(isMinting bool) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetIsMinting(&_ShoeSharkNft.TransactOpts, isMinting)
}

// SetIsMinting is a paid mutator transaction binding the contract method 0x7b93d1e5.
//
// Solidity: function setIsMinting(bool isMinting) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SetIsMinting(isMinting bool) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetIsMinting(&_ShoeSharkNft.TransactOpts, isMinting)
}

// SetMetadataUri is a paid mutator transaction binding the contract method 0x1130630c.
//
// Solidity: function setMetadataUri(string uri) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SetMetadataUri(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "setMetadataUri", uri)
}

// SetMetadataUri is a paid mutator transaction binding the contract method 0x1130630c.
//
// Solidity: function setMetadataUri(string uri) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SetMetadataUri(uri string) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetMetadataUri(&_ShoeSharkNft.TransactOpts, uri)
}

// SetMetadataUri is a paid mutator transaction binding the contract method 0x1130630c.
//
// Solidity: function setMetadataUri(string uri) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SetMetadataUri(uri string) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetMetadataUri(&_ShoeSharkNft.TransactOpts, uri)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 newRoot) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) SetRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "setRoot", newRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 newRoot) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) SetRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetRoot(&_ShoeSharkNft.TransactOpts, newRoot)
}

// SetRoot is a paid mutator transaction binding the contract method 0xdab5f340.
//
// Solidity: function setRoot(bytes32 newRoot) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) SetRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.SetRoot(&_ShoeSharkNft.TransactOpts, newRoot)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.TransferFrom(&_ShoeSharkNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.TransferFrom(&_ShoeSharkNft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkNft *ShoeSharkNftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.TransferOwnership(&_ShoeSharkNft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.TransferOwnership(&_ShoeSharkNft.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_ShoeSharkNft *ShoeSharkNftTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkNft.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_ShoeSharkNft *ShoeSharkNftSession) Withdraw() (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Withdraw(&_ShoeSharkNft.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_ShoeSharkNft *ShoeSharkNftTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ShoeSharkNft.Contract.Withdraw(&_ShoeSharkNft.TransactOpts)
}

// ShoeSharkNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShoeSharkNft contract.
type ShoeSharkNftApprovalIterator struct {
	Event *ShoeSharkNftApproval // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftApproval)
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
		it.Event = new(ShoeSharkNftApproval)
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
func (it *ShoeSharkNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftApproval represents a Approval event raised by the ShoeSharkNft contract.
type ShoeSharkNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ShoeSharkNftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftApprovalIterator{contract: _ShoeSharkNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftApproval)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseApproval(log types.Log) (*ShoeSharkNftApproval, error) {
	event := new(ShoeSharkNftApproval)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ShoeSharkNft contract.
type ShoeSharkNftApprovalForAllIterator struct {
	Event *ShoeSharkNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftApprovalForAll)
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
		it.Event = new(ShoeSharkNftApprovalForAll)
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
func (it *ShoeSharkNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftApprovalForAll represents a ApprovalForAll event raised by the ShoeSharkNft contract.
type ShoeSharkNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ShoeSharkNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftApprovalForAllIterator{contract: _ShoeSharkNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftApprovalForAll)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseApprovalForAll(log types.Log) (*ShoeSharkNftApprovalForAll, error) {
	event := new(ShoeSharkNftApprovalForAll)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ShoeSharkNft contract.
type ShoeSharkNftBatchMetadataUpdateIterator struct {
	Event *ShoeSharkNftBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftBatchMetadataUpdate)
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
		it.Event = new(ShoeSharkNftBatchMetadataUpdate)
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
func (it *ShoeSharkNftBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ShoeSharkNft contract.
type ShoeSharkNftBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ShoeSharkNftBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftBatchMetadataUpdateIterator{contract: _ShoeSharkNft.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftBatchMetadataUpdate)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseBatchMetadataUpdate(log types.Log) (*ShoeSharkNftBatchMetadataUpdate, error) {
	event := new(ShoeSharkNftBatchMetadataUpdate)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ShoeSharkNft contract.
type ShoeSharkNftMetadataUpdateIterator struct {
	Event *ShoeSharkNftMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftMetadataUpdate)
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
		it.Event = new(ShoeSharkNftMetadataUpdate)
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
func (it *ShoeSharkNftMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftMetadataUpdate represents a MetadataUpdate event raised by the ShoeSharkNft contract.
type ShoeSharkNftMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ShoeSharkNftMetadataUpdateIterator, error) {

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftMetadataUpdateIterator{contract: _ShoeSharkNft.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftMetadataUpdate)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseMetadataUpdate(log types.Log) (*ShoeSharkNftMetadataUpdate, error) {
	event := new(ShoeSharkNftMetadataUpdate)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShoeSharkNft contract.
type ShoeSharkNftOwnershipTransferredIterator struct {
	Event *ShoeSharkNftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftOwnershipTransferred)
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
		it.Event = new(ShoeSharkNftOwnershipTransferred)
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
func (it *ShoeSharkNftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftOwnershipTransferred represents a OwnershipTransferred event raised by the ShoeSharkNft contract.
type ShoeSharkNftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShoeSharkNftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftOwnershipTransferredIterator{contract: _ShoeSharkNft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftOwnershipTransferred)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseOwnershipTransferred(log types.Log) (*ShoeSharkNftOwnershipTransferred, error) {
	event := new(ShoeSharkNftOwnershipTransferred)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftShoeSharkNftNftBurnedIterator is returned from FilterShoeSharkNftNftBurned and is used to iterate over the raw logs and unpacked data for ShoeSharkNftNftBurned events raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftNftBurnedIterator struct {
	Event *ShoeSharkNftShoeSharkNftNftBurned // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftShoeSharkNftNftBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftShoeSharkNftNftBurned)
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
		it.Event = new(ShoeSharkNftShoeSharkNftNftBurned)
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
func (it *ShoeSharkNftShoeSharkNftNftBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftShoeSharkNftNftBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftShoeSharkNftNftBurned represents a ShoeSharkNftNftBurned event raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftNftBurned struct {
	Player  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftNftBurned is a free log retrieval operation binding the contract event 0xb781a8ff7579d9ce6007f40b74e6cc669642276819b1d6376e5c7d45bf8ae207.
//
// Solidity: event ShoeSharkNft_NftBurned(address indexed player, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterShoeSharkNftNftBurned(opts *bind.FilterOpts, player []common.Address, tokenId []*big.Int) (*ShoeSharkNftShoeSharkNftNftBurnedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "ShoeSharkNft_NftBurned", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftShoeSharkNftNftBurnedIterator{contract: _ShoeSharkNft.contract, event: "ShoeSharkNft_NftBurned", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftNftBurned is a free log subscription operation binding the contract event 0xb781a8ff7579d9ce6007f40b74e6cc669642276819b1d6376e5c7d45bf8ae207.
//
// Solidity: event ShoeSharkNft_NftBurned(address indexed player, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchShoeSharkNftNftBurned(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftShoeSharkNftNftBurned, player []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "ShoeSharkNft_NftBurned", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftShoeSharkNftNftBurned)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_NftBurned", log); err != nil {
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

// ParseShoeSharkNftNftBurned is a log parse operation binding the contract event 0xb781a8ff7579d9ce6007f40b74e6cc669642276819b1d6376e5c7d45bf8ae207.
//
// Solidity: event ShoeSharkNft_NftBurned(address indexed player, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseShoeSharkNftNftBurned(log types.Log) (*ShoeSharkNftShoeSharkNftNftBurned, error) {
	event := new(ShoeSharkNftShoeSharkNftNftBurned)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_NftBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftShoeSharkNftNftMintedIterator is returned from FilterShoeSharkNftNftMinted and is used to iterate over the raw logs and unpacked data for ShoeSharkNftNftMinted events raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftNftMintedIterator struct {
	Event *ShoeSharkNftShoeSharkNftNftMinted // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftShoeSharkNftNftMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftShoeSharkNftNftMinted)
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
		it.Event = new(ShoeSharkNftShoeSharkNftNftMinted)
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
func (it *ShoeSharkNftShoeSharkNftNftMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftShoeSharkNftNftMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftShoeSharkNftNftMinted represents a ShoeSharkNftNftMinted event raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftNftMinted struct {
	Player    common.Address
	TokenId   *big.Int
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftNftMinted is a free log retrieval operation binding the contract event 0x7f3fb005410316e1c12471b0a260ddfa89d456cf131b9877365b0eeba79c120c.
//
// Solidity: event ShoeSharkNft_NftMinted(address indexed player, uint256 indexed tokenId, uint256 requestId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterShoeSharkNftNftMinted(opts *bind.FilterOpts, player []common.Address, tokenId []*big.Int) (*ShoeSharkNftShoeSharkNftNftMintedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "ShoeSharkNft_NftMinted", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftShoeSharkNftNftMintedIterator{contract: _ShoeSharkNft.contract, event: "ShoeSharkNft_NftMinted", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftNftMinted is a free log subscription operation binding the contract event 0x7f3fb005410316e1c12471b0a260ddfa89d456cf131b9877365b0eeba79c120c.
//
// Solidity: event ShoeSharkNft_NftMinted(address indexed player, uint256 indexed tokenId, uint256 requestId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchShoeSharkNftNftMinted(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftShoeSharkNftNftMinted, player []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "ShoeSharkNft_NftMinted", playerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftShoeSharkNftNftMinted)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_NftMinted", log); err != nil {
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

// ParseShoeSharkNftNftMinted is a log parse operation binding the contract event 0x7f3fb005410316e1c12471b0a260ddfa89d456cf131b9877365b0eeba79c120c.
//
// Solidity: event ShoeSharkNft_NftMinted(address indexed player, uint256 indexed tokenId, uint256 requestId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseShoeSharkNftNftMinted(log types.Log) (*ShoeSharkNftShoeSharkNftNftMinted, error) {
	event := new(ShoeSharkNftShoeSharkNftNftMinted)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_NftMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftShoeSharkNftWithdrawIterator is returned from FilterShoeSharkNftWithdraw and is used to iterate over the raw logs and unpacked data for ShoeSharkNftWithdraw events raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftWithdrawIterator struct {
	Event *ShoeSharkNftShoeSharkNftWithdraw // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftShoeSharkNftWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftShoeSharkNftWithdraw)
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
		it.Event = new(ShoeSharkNftShoeSharkNftWithdraw)
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
func (it *ShoeSharkNftShoeSharkNftWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftShoeSharkNftWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftShoeSharkNftWithdraw represents a ShoeSharkNftWithdraw event raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftWithdraw struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftWithdraw is a free log retrieval operation binding the contract event 0x52d1e090427e43f39f0a1936063f6d6f8619d7cb57685fd665f38d8ca12ba668.
//
// Solidity: event ShoeSharkNft_Withdraw()
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterShoeSharkNftWithdraw(opts *bind.FilterOpts) (*ShoeSharkNftShoeSharkNftWithdrawIterator, error) {

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "ShoeSharkNft_Withdraw")
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftShoeSharkNftWithdrawIterator{contract: _ShoeSharkNft.contract, event: "ShoeSharkNft_Withdraw", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftWithdraw is a free log subscription operation binding the contract event 0x52d1e090427e43f39f0a1936063f6d6f8619d7cb57685fd665f38d8ca12ba668.
//
// Solidity: event ShoeSharkNft_Withdraw()
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchShoeSharkNftWithdraw(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftShoeSharkNftWithdraw) (event.Subscription, error) {

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "ShoeSharkNft_Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftShoeSharkNftWithdraw)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_Withdraw", log); err != nil {
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

// ParseShoeSharkNftWithdraw is a log parse operation binding the contract event 0x52d1e090427e43f39f0a1936063f6d6f8619d7cb57685fd665f38d8ca12ba668.
//
// Solidity: event ShoeSharkNft_Withdraw()
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseShoeSharkNftWithdraw(log types.Log) (*ShoeSharkNftShoeSharkNftWithdraw, error) {
	event := new(ShoeSharkNftShoeSharkNftWithdraw)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator is returned from FilterShoeSharkNftTokenIdtoTokenURI and is used to iterate over the raw logs and unpacked data for ShoeSharkNftTokenIdtoTokenURI events raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator struct {
	Event *ShoeSharkNftShoeSharkNftTokenIdtoTokenURI // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftShoeSharkNftTokenIdtoTokenURI)
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
		it.Event = new(ShoeSharkNftShoeSharkNftTokenIdtoTokenURI)
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
func (it *ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftShoeSharkNftTokenIdtoTokenURI represents a ShoeSharkNftTokenIdtoTokenURI event raised by the ShoeSharkNft contract.
type ShoeSharkNftShoeSharkNftTokenIdtoTokenURI struct {
	TokenId       *big.Int
	MetadateIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkNftTokenIdtoTokenURI is a free log retrieval operation binding the contract event 0x6f47fe54448c9bbe840f8436b97e684310c20e652fb8be073e176feeb991bc67.
//
// Solidity: event ShoeSharkNft_tokenIdtoTokenURI(uint256 indexed tokenId, uint256 metadateIndex)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterShoeSharkNftTokenIdtoTokenURI(opts *bind.FilterOpts, tokenId []*big.Int) (*ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "ShoeSharkNft_tokenIdtoTokenURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftShoeSharkNftTokenIdtoTokenURIIterator{contract: _ShoeSharkNft.contract, event: "ShoeSharkNft_tokenIdtoTokenURI", logs: logs, sub: sub}, nil
}

// WatchShoeSharkNftTokenIdtoTokenURI is a free log subscription operation binding the contract event 0x6f47fe54448c9bbe840f8436b97e684310c20e652fb8be073e176feeb991bc67.
//
// Solidity: event ShoeSharkNft_tokenIdtoTokenURI(uint256 indexed tokenId, uint256 metadateIndex)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchShoeSharkNftTokenIdtoTokenURI(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftShoeSharkNftTokenIdtoTokenURI, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "ShoeSharkNft_tokenIdtoTokenURI", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftShoeSharkNftTokenIdtoTokenURI)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_tokenIdtoTokenURI", log); err != nil {
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

// ParseShoeSharkNftTokenIdtoTokenURI is a log parse operation binding the contract event 0x6f47fe54448c9bbe840f8436b97e684310c20e652fb8be073e176feeb991bc67.
//
// Solidity: event ShoeSharkNft_tokenIdtoTokenURI(uint256 indexed tokenId, uint256 metadateIndex)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseShoeSharkNftTokenIdtoTokenURI(log types.Log) (*ShoeSharkNftShoeSharkNftTokenIdtoTokenURI, error) {
	event := new(ShoeSharkNftShoeSharkNftTokenIdtoTokenURI)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "ShoeSharkNft_tokenIdtoTokenURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShoeSharkNft contract.
type ShoeSharkNftTransferIterator struct {
	Event *ShoeSharkNftTransfer // Event containing the contract specifics and raw log

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
func (it *ShoeSharkNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkNftTransfer)
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
		it.Event = new(ShoeSharkNftTransfer)
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
func (it *ShoeSharkNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkNftTransfer represents a Transfer event raised by the ShoeSharkNft contract.
type ShoeSharkNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ShoeSharkNftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkNftTransferIterator{contract: _ShoeSharkNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShoeSharkNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShoeSharkNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkNftTransfer)
				if err := _ShoeSharkNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShoeSharkNft *ShoeSharkNftFilterer) ParseTransfer(log types.Log) (*ShoeSharkNftTransfer, error) {
	event := new(ShoeSharkNftTransfer)
	if err := _ShoeSharkNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
