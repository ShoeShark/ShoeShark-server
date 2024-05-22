// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ShoeSharkRewardPointMetaData contains all meta data concerning the ShoeSharkRewardPoint contract.
var ShoeSharkRewardPointMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_shoeSharkToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractShoeSharkToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exchangeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserPoints\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeemAllPointsForTokens\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemPointsForTokensForAddress\",\"inputs\":[{\"name\":\"pointHolder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"points\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_pointHolders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_pointsMap\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setExchangeRate\",\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoints\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkRewardPoint_BatchPointSet\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShoeSharkRewardPoint_PointSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ShoeSharkRewardPoint__redeemPointsForTokensForAddress__NotEnoughPoints\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ShoeSharkRewardPoint__setPoints__NotEqualLength\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c3d380380610c3d83398101604081905261002f916100d9565b338061005557604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61005e81610089565b50600380546001600160a01b0319166001600160a01b03929092169190911790556001600255610109565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100eb57600080fd5b81516001600160a01b038116811461010257600080fd5b9392505050565b610b25806101186000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063aeefe31f1161008c578063db068e0e11610066578063db068e0e146101b8578063f2fde38b146101cb578063f4c90b58146101de578063f6f0e0e6146101e657600080fd5b8063aeefe31f14610169578063b8b770b014610192578063c7964257146101a557600080fd5b80632383a230146100d45780633ba0b9a914610104578063715018a61461011b5780638da5cb5b146101255780638ebb4f18146101365780639a80c4a114610156575b600080fd5b6003546100e7906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61010d60025481565b6040519081526020016100fb565b6101236101f9565b005b6000546001600160a01b03166100e7565b61010d610144366004610829565b60016020526000908152604090205481565b610123610164366004610921565b61020d565b61010d610177366004610829565b6001600160a01b031660009081526001602052604090205490565b6100e76101a03660046109e1565b610370565b6101236101b33660046109fa565b61039a565b6101236101c63660046109e1565b610472565b6101236101d9366004610829565b61047f565b6101236104c2565b6101236101f43660046109fa565b610539565b6102016105bf565b61020b60006105ec565b565b6102156105bf565b8051825114610237576040516328f564a160e11b815260040160405180910390fd5b60005b825181101561033157600083828151811061025757610257610a24565b60200260200101519050600083838151811061027557610275610a24565b6020026020010151905060016000836001600160a01b03166001600160a01b031681526020019081526020016000205460001480156102b45750600081115b1561030557600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b039091166000908152600160205260409020558061032981610a50565b91505061023a565b507f015fe5e6142fe7776168060d5fd0058d1dd80ae0a1163d17db7d650d143cadc9825160405161036491815260200190565b60405180910390a15050565b6004818154811061038057600080fd5b6000918252602090912001546001600160a01b0316905081565b6103a26105bf565b6001600160a01b0382166000908152600160205260409020541580156103c85750600081115b1561041957600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b03821660008181526001602052604090819020839055517fa7edf51bc67b4fa6b662aead0f2a92547f3f8d7f250f61ff4218fa017bf51083906104669084815260200190565b60405180910390a25050565b61047a6105bf565b600255565b6104876105bf565b6001600160a01b0381166104b657604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b6104bf816105ec565b50565b6104ca6105bf565b60005b6004548110156104bf576000600482815481106104ec576104ec610a24565b60009182526020808320909101546001600160a01b031680835260019091526040909120549091508015610524576105248282610539565b5050808061053190610a50565b9150506104cd565b6105416105bf565b6001600160a01b03821660009081526001602052604090205481111561057a576040516324555e8d60e01b815260040160405180910390fd5b610584828261063c565b6001600160a01b038216600090815260016020526040812080548392906105ac908490610a69565b909155506105bb9050826106e7565b5050565b6000546001600160a01b0316331461020b5760405163118cdaa760e01b81523360048201526024016104ad565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600254600090600161064e8285610a82565b6106589190610a69565b6106629190610a95565b6003546040516323b872dd60e01b81523360048201526001600160a01b038681166024830152604482018490529293509116906323b872dd906064016020604051808303816000875af11580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190610ab7565b50505050565b6001600160a01b03811660009081526001602052604081205490036104bf5760005b6004548110156105bb57816001600160a01b03166004828154811061073057610730610a24565b6000918252602090912001546001600160a01b0316036107fb576004805461075a90600190610a69565b8154811061076a5761076a610a24565b600091825260209091200154600480546001600160a01b03909216918390811061079657610796610a24565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060048054806107d5576107d5610ad9565b600082815260209020810160001990810180546001600160a01b03191690550190555050565b8061080581610a50565b915050610709565b80356001600160a01b038116811461082457600080fd5b919050565b60006020828403121561083b57600080fd5b6108448261080d565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561088a5761088a61084b565b604052919050565b600067ffffffffffffffff8211156108ac576108ac61084b565b5060051b60200190565b600082601f8301126108c757600080fd5b813560206108dc6108d783610892565b610861565b82815260059290921b840181019181810190868411156108fb57600080fd5b8286015b8481101561091657803583529183019183016108ff565b509695505050505050565b6000806040838503121561093457600080fd5b823567ffffffffffffffff8082111561094c57600080fd5b818501915085601f83011261096057600080fd5b813560206109706108d783610892565b82815260059290921b8401810191818101908984111561098f57600080fd5b948201945b838610156109b4576109a58661080d565b82529482019490820190610994565b965050860135925050808211156109ca57600080fd5b506109d7858286016108b6565b9150509250929050565b6000602082840312156109f357600080fd5b5035919050565b60008060408385031215610a0d57600080fd5b610a168361080d565b946020939093013593505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a6257610a62610a3a565b5060010190565b81810381811115610a7c57610a7c610a3a565b92915050565b80820180821115610a7c57610a7c610a3a565b600082610ab257634e487b7160e01b600052601260045260246000fd5b500490565b600060208284031215610ac957600080fd5b8151801515811461084457600080fd5b634e487b7160e01b600052603160045260246000fdfea2646970667358221220cac83eddb06f4f711217d20a74ca71f50aabdd02f526031c9aa066daa926889d64736f6c63430008140033",
}

// ShoeSharkRewardPointABI is the input ABI used to generate the binding from.
// Deprecated: Use ShoeSharkRewardPointMetaData.ABI instead.
var ShoeSharkRewardPointABI = ShoeSharkRewardPointMetaData.ABI

// ShoeSharkRewardPointBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShoeSharkRewardPointMetaData.Bin instead.
var ShoeSharkRewardPointBin = ShoeSharkRewardPointMetaData.Bin

// DeployShoeSharkRewardPoint deploys a new Ethereum contract, binding an instance of ShoeSharkRewardPoint to it.
func DeployShoeSharkRewardPoint(auth *bind.TransactOpts, backend bind.ContractBackend, _shoeSharkToken common.Address) (common.Address, *types.Transaction, *ShoeSharkRewardPoint, error) {
	parsed, err := ShoeSharkRewardPointMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShoeSharkRewardPointBin), backend, _shoeSharkToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShoeSharkRewardPoint{ShoeSharkRewardPointCaller: ShoeSharkRewardPointCaller{contract: contract}, ShoeSharkRewardPointTransactor: ShoeSharkRewardPointTransactor{contract: contract}, ShoeSharkRewardPointFilterer: ShoeSharkRewardPointFilterer{contract: contract}}, nil
}

// ShoeSharkRewardPoint is an auto generated Go binding around an Ethereum contract.
type ShoeSharkRewardPoint struct {
	ShoeSharkRewardPointCaller     // Read-only binding to the contract
	ShoeSharkRewardPointTransactor // Write-only binding to the contract
	ShoeSharkRewardPointFilterer   // Log filterer for contract events
}

// ShoeSharkRewardPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShoeSharkRewardPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkRewardPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShoeSharkRewardPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkRewardPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShoeSharkRewardPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShoeSharkRewardPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShoeSharkRewardPointSession struct {
	Contract     *ShoeSharkRewardPoint // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ShoeSharkRewardPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShoeSharkRewardPointCallerSession struct {
	Contract *ShoeSharkRewardPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ShoeSharkRewardPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShoeSharkRewardPointTransactorSession struct {
	Contract     *ShoeSharkRewardPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ShoeSharkRewardPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShoeSharkRewardPointRaw struct {
	Contract *ShoeSharkRewardPoint // Generic contract binding to access the raw methods on
}

// ShoeSharkRewardPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShoeSharkRewardPointCallerRaw struct {
	Contract *ShoeSharkRewardPointCaller // Generic read-only contract binding to access the raw methods on
}

// ShoeSharkRewardPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShoeSharkRewardPointTransactorRaw struct {
	Contract *ShoeSharkRewardPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShoeSharkRewardPoint creates a new instance of ShoeSharkRewardPoint, bound to a specific deployed contract.
func NewShoeSharkRewardPoint(address common.Address, backend bind.ContractBackend) (*ShoeSharkRewardPoint, error) {
	contract, err := bindShoeSharkRewardPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPoint{ShoeSharkRewardPointCaller: ShoeSharkRewardPointCaller{contract: contract}, ShoeSharkRewardPointTransactor: ShoeSharkRewardPointTransactor{contract: contract}, ShoeSharkRewardPointFilterer: ShoeSharkRewardPointFilterer{contract: contract}}, nil
}

// NewShoeSharkRewardPointCaller creates a new read-only instance of ShoeSharkRewardPoint, bound to a specific deployed contract.
func NewShoeSharkRewardPointCaller(address common.Address, caller bind.ContractCaller) (*ShoeSharkRewardPointCaller, error) {
	contract, err := bindShoeSharkRewardPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointCaller{contract: contract}, nil
}

// NewShoeSharkRewardPointTransactor creates a new write-only instance of ShoeSharkRewardPoint, bound to a specific deployed contract.
func NewShoeSharkRewardPointTransactor(address common.Address, transactor bind.ContractTransactor) (*ShoeSharkRewardPointTransactor, error) {
	contract, err := bindShoeSharkRewardPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointTransactor{contract: contract}, nil
}

// NewShoeSharkRewardPointFilterer creates a new log filterer instance of ShoeSharkRewardPoint, bound to a specific deployed contract.
func NewShoeSharkRewardPointFilterer(address common.Address, filterer bind.ContractFilterer) (*ShoeSharkRewardPointFilterer, error) {
	contract, err := bindShoeSharkRewardPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointFilterer{contract: contract}, nil
}

// bindShoeSharkRewardPoint binds a generic wrapper to an already deployed contract.
func bindShoeSharkRewardPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShoeSharkRewardPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkRewardPoint.Contract.ShoeSharkRewardPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.ShoeSharkRewardPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.ShoeSharkRewardPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShoeSharkRewardPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.contract.Transact(opts, method, params...)
}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) SST(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "SST")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SST() (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.SST(&_ShoeSharkRewardPoint.CallOpts)
}

// SST is a free data retrieval call binding the contract method 0x2383a230.
//
// Solidity: function SST() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) SST() (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.SST(&_ShoeSharkRewardPoint.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) ExchangeRate() (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.ExchangeRate(&_ShoeSharkRewardPoint.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) ExchangeRate() (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.ExchangeRate(&_ShoeSharkRewardPoint.CallOpts)
}

// GetUserPoints is a free data retrieval call binding the contract method 0xaeefe31f.
//
// Solidity: function getUserPoints(address user) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) GetUserPoints(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "getUserPoints", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserPoints is a free data retrieval call binding the contract method 0xaeefe31f.
//
// Solidity: function getUserPoints(address user) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) GetUserPoints(user common.Address) (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.GetUserPoints(&_ShoeSharkRewardPoint.CallOpts, user)
}

// GetUserPoints is a free data retrieval call binding the contract method 0xaeefe31f.
//
// Solidity: function getUserPoints(address user) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) GetUserPoints(user common.Address) (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.GetUserPoints(&_ShoeSharkRewardPoint.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) Owner() (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.Owner(&_ShoeSharkRewardPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) Owner() (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.Owner(&_ShoeSharkRewardPoint.CallOpts)
}

// SPointHolders is a free data retrieval call binding the contract method 0xb8b770b0.
//
// Solidity: function s_pointHolders(uint256 ) view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) SPointHolders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "s_pointHolders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPointHolders is a free data retrieval call binding the contract method 0xb8b770b0.
//
// Solidity: function s_pointHolders(uint256 ) view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SPointHolders(arg0 *big.Int) (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.SPointHolders(&_ShoeSharkRewardPoint.CallOpts, arg0)
}

// SPointHolders is a free data retrieval call binding the contract method 0xb8b770b0.
//
// Solidity: function s_pointHolders(uint256 ) view returns(address)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) SPointHolders(arg0 *big.Int) (common.Address, error) {
	return _ShoeSharkRewardPoint.Contract.SPointHolders(&_ShoeSharkRewardPoint.CallOpts, arg0)
}

// SPointsMap is a free data retrieval call binding the contract method 0x8ebb4f18.
//
// Solidity: function s_pointsMap(address ) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCaller) SPointsMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShoeSharkRewardPoint.contract.Call(opts, &out, "s_pointsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SPointsMap is a free data retrieval call binding the contract method 0x8ebb4f18.
//
// Solidity: function s_pointsMap(address ) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SPointsMap(arg0 common.Address) (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.SPointsMap(&_ShoeSharkRewardPoint.CallOpts, arg0)
}

// SPointsMap is a free data retrieval call binding the contract method 0x8ebb4f18.
//
// Solidity: function s_pointsMap(address ) view returns(uint256)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointCallerSession) SPointsMap(arg0 common.Address) (*big.Int, error) {
	return _ShoeSharkRewardPoint.Contract.SPointsMap(&_ShoeSharkRewardPoint.CallOpts, arg0)
}

// RedeemAllPointsForTokens is a paid mutator transaction binding the contract method 0xf4c90b58.
//
// Solidity: function redeemAllPointsForTokens() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) RedeemAllPointsForTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "redeemAllPointsForTokens")
}

// RedeemAllPointsForTokens is a paid mutator transaction binding the contract method 0xf4c90b58.
//
// Solidity: function redeemAllPointsForTokens() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) RedeemAllPointsForTokens() (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RedeemAllPointsForTokens(&_ShoeSharkRewardPoint.TransactOpts)
}

// RedeemAllPointsForTokens is a paid mutator transaction binding the contract method 0xf4c90b58.
//
// Solidity: function redeemAllPointsForTokens() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) RedeemAllPointsForTokens() (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RedeemAllPointsForTokens(&_ShoeSharkRewardPoint.TransactOpts)
}

// RedeemPointsForTokensForAddress is a paid mutator transaction binding the contract method 0xf6f0e0e6.
//
// Solidity: function redeemPointsForTokensForAddress(address pointHolder, uint256 points) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) RedeemPointsForTokensForAddress(opts *bind.TransactOpts, pointHolder common.Address, points *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "redeemPointsForTokensForAddress", pointHolder, points)
}

// RedeemPointsForTokensForAddress is a paid mutator transaction binding the contract method 0xf6f0e0e6.
//
// Solidity: function redeemPointsForTokensForAddress(address pointHolder, uint256 points) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) RedeemPointsForTokensForAddress(pointHolder common.Address, points *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RedeemPointsForTokensForAddress(&_ShoeSharkRewardPoint.TransactOpts, pointHolder, points)
}

// RedeemPointsForTokensForAddress is a paid mutator transaction binding the contract method 0xf6f0e0e6.
//
// Solidity: function redeemPointsForTokensForAddress(address pointHolder, uint256 points) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) RedeemPointsForTokensForAddress(pointHolder common.Address, points *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RedeemPointsForTokensForAddress(&_ShoeSharkRewardPoint.TransactOpts, pointHolder, points)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RenounceOwnership(&_ShoeSharkRewardPoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.RenounceOwnership(&_ShoeSharkRewardPoint.TransactOpts)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) SetExchangeRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "setExchangeRate", newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetExchangeRate(&_ShoeSharkRewardPoint.TransactOpts, newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetExchangeRate(&_ShoeSharkRewardPoint.TransactOpts, newRate)
}

// SetPoint is a paid mutator transaction binding the contract method 0xc7964257.
//
// Solidity: function setPoint(address account, uint256 amount) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) SetPoint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "setPoint", account, amount)
}

// SetPoint is a paid mutator transaction binding the contract method 0xc7964257.
//
// Solidity: function setPoint(address account, uint256 amount) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SetPoint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetPoint(&_ShoeSharkRewardPoint.TransactOpts, account, amount)
}

// SetPoint is a paid mutator transaction binding the contract method 0xc7964257.
//
// Solidity: function setPoint(address account, uint256 amount) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) SetPoint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetPoint(&_ShoeSharkRewardPoint.TransactOpts, account, amount)
}

// SetPoints is a paid mutator transaction binding the contract method 0x9a80c4a1.
//
// Solidity: function setPoints(address[] accounts, uint256[] amounts) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) SetPoints(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "setPoints", accounts, amounts)
}

// SetPoints is a paid mutator transaction binding the contract method 0x9a80c4a1.
//
// Solidity: function setPoints(address[] accounts, uint256[] amounts) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) SetPoints(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetPoints(&_ShoeSharkRewardPoint.TransactOpts, accounts, amounts)
}

// SetPoints is a paid mutator transaction binding the contract method 0x9a80c4a1.
//
// Solidity: function setPoints(address[] accounts, uint256[] amounts) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) SetPoints(accounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.SetPoints(&_ShoeSharkRewardPoint.TransactOpts, accounts, amounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.TransferOwnership(&_ShoeSharkRewardPoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShoeSharkRewardPoint.Contract.TransferOwnership(&_ShoeSharkRewardPoint.TransactOpts, newOwner)
}

// ShoeSharkRewardPointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointOwnershipTransferredIterator struct {
	Event *ShoeSharkRewardPointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShoeSharkRewardPointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkRewardPointOwnershipTransferred)
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
		it.Event = new(ShoeSharkRewardPointOwnershipTransferred)
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
func (it *ShoeSharkRewardPointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkRewardPointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkRewardPointOwnershipTransferred represents a OwnershipTransferred event raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShoeSharkRewardPointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkRewardPoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointOwnershipTransferredIterator{contract: _ShoeSharkRewardPoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShoeSharkRewardPointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShoeSharkRewardPoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkRewardPointOwnershipTransferred)
				if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) ParseOwnershipTransferred(log types.Log) (*ShoeSharkRewardPointOwnershipTransferred, error) {
	event := new(ShoeSharkRewardPointOwnershipTransferred)
	if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator is returned from FilterShoeSharkRewardPointBatchPointSet and is used to iterate over the raw logs and unpacked data for ShoeSharkRewardPointBatchPointSet events raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator struct {
	Event *ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet // Event containing the contract specifics and raw log

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
func (it *ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet)
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
		it.Event = new(ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet)
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
func (it *ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet represents a ShoeSharkRewardPointBatchPointSet event raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkRewardPointBatchPointSet is a free log retrieval operation binding the contract event 0x015fe5e6142fe7776168060d5fd0058d1dd80ae0a1163d17db7d650d143cadc9.
//
// Solidity: event ShoeSharkRewardPoint_BatchPointSet(uint256 length)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) FilterShoeSharkRewardPointBatchPointSet(opts *bind.FilterOpts) (*ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator, error) {

	logs, sub, err := _ShoeSharkRewardPoint.contract.FilterLogs(opts, "ShoeSharkRewardPoint_BatchPointSet")
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointShoeSharkRewardPointBatchPointSetIterator{contract: _ShoeSharkRewardPoint.contract, event: "ShoeSharkRewardPoint_BatchPointSet", logs: logs, sub: sub}, nil
}

// WatchShoeSharkRewardPointBatchPointSet is a free log subscription operation binding the contract event 0x015fe5e6142fe7776168060d5fd0058d1dd80ae0a1163d17db7d650d143cadc9.
//
// Solidity: event ShoeSharkRewardPoint_BatchPointSet(uint256 length)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) WatchShoeSharkRewardPointBatchPointSet(opts *bind.WatchOpts, sink chan<- *ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet) (event.Subscription, error) {

	logs, sub, err := _ShoeSharkRewardPoint.contract.WatchLogs(opts, "ShoeSharkRewardPoint_BatchPointSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet)
				if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "ShoeSharkRewardPoint_BatchPointSet", log); err != nil {
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

// ParseShoeSharkRewardPointBatchPointSet is a log parse operation binding the contract event 0x015fe5e6142fe7776168060d5fd0058d1dd80ae0a1163d17db7d650d143cadc9.
//
// Solidity: event ShoeSharkRewardPoint_BatchPointSet(uint256 length)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) ParseShoeSharkRewardPointBatchPointSet(log types.Log) (*ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet, error) {
	event := new(ShoeSharkRewardPointShoeSharkRewardPointBatchPointSet)
	if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "ShoeSharkRewardPoint_BatchPointSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator is returned from FilterShoeSharkRewardPointPointSet and is used to iterate over the raw logs and unpacked data for ShoeSharkRewardPointPointSet events raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator struct {
	Event *ShoeSharkRewardPointShoeSharkRewardPointPointSet // Event containing the contract specifics and raw log

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
func (it *ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShoeSharkRewardPointShoeSharkRewardPointPointSet)
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
		it.Event = new(ShoeSharkRewardPointShoeSharkRewardPointPointSet)
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
func (it *ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShoeSharkRewardPointShoeSharkRewardPointPointSet represents a ShoeSharkRewardPointPointSet event raised by the ShoeSharkRewardPoint contract.
type ShoeSharkRewardPointShoeSharkRewardPointPointSet struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShoeSharkRewardPointPointSet is a free log retrieval operation binding the contract event 0xa7edf51bc67b4fa6b662aead0f2a92547f3f8d7f250f61ff4218fa017bf51083.
//
// Solidity: event ShoeSharkRewardPoint_PointSet(address indexed account, uint256 amount)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) FilterShoeSharkRewardPointPointSet(opts *bind.FilterOpts, account []common.Address) (*ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShoeSharkRewardPoint.contract.FilterLogs(opts, "ShoeSharkRewardPoint_PointSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &ShoeSharkRewardPointShoeSharkRewardPointPointSetIterator{contract: _ShoeSharkRewardPoint.contract, event: "ShoeSharkRewardPoint_PointSet", logs: logs, sub: sub}, nil
}

// WatchShoeSharkRewardPointPointSet is a free log subscription operation binding the contract event 0xa7edf51bc67b4fa6b662aead0f2a92547f3f8d7f250f61ff4218fa017bf51083.
//
// Solidity: event ShoeSharkRewardPoint_PointSet(address indexed account, uint256 amount)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) WatchShoeSharkRewardPointPointSet(opts *bind.WatchOpts, sink chan<- *ShoeSharkRewardPointShoeSharkRewardPointPointSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ShoeSharkRewardPoint.contract.WatchLogs(opts, "ShoeSharkRewardPoint_PointSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShoeSharkRewardPointShoeSharkRewardPointPointSet)
				if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "ShoeSharkRewardPoint_PointSet", log); err != nil {
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

// ParseShoeSharkRewardPointPointSet is a log parse operation binding the contract event 0xa7edf51bc67b4fa6b662aead0f2a92547f3f8d7f250f61ff4218fa017bf51083.
//
// Solidity: event ShoeSharkRewardPoint_PointSet(address indexed account, uint256 amount)
func (_ShoeSharkRewardPoint *ShoeSharkRewardPointFilterer) ParseShoeSharkRewardPointPointSet(log types.Log) (*ShoeSharkRewardPointShoeSharkRewardPointPointSet, error) {
	event := new(ShoeSharkRewardPointShoeSharkRewardPointPointSet)
	if err := _ShoeSharkRewardPoint.contract.UnpackLog(event, "ShoeSharkRewardPoint_PointSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
