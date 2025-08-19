// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	Taker         common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int
	Nonce         *big.Int
	FeeRateBps    *big.Int
	Side          uint8
	SignatureType uint8
	Signature     []byte
}

// V2MetaData contains all meta data concerning the V2 contract.
var V2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_negRiskCtfExchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"feeCharged\",\"type\":\"uint256\"}],\"name\":\"FeeRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerReceiveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V2ABI is the input ABI used to generate the binding from.
// Deprecated: Use V2MetaData.ABI instead.
var V2ABI = V2MetaData.ABI

// V2 is an auto generated Go binding around an Ethereum contract.
type V2 struct {
	V2Caller     // Read-only binding to the contract
	V2Transactor // Write-only binding to the contract
	V2Filterer   // Log filterer for contract events
}

// V2Caller is an auto generated read-only Go binding around an Ethereum contract.
type V2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type V2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2Session struct {
	Contract     *V2               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2CallerSession struct {
	Contract *V2Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// V2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2TransactorSession struct {
	Contract     *V2Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2Raw is an auto generated low-level Go binding around an Ethereum contract.
type V2Raw struct {
	Contract *V2 // Generic contract binding to access the raw methods on
}

// V2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2CallerRaw struct {
	Contract *V2Caller // Generic read-only contract binding to access the raw methods on
}

// V2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2TransactorRaw struct {
	Contract *V2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewV2 creates a new instance of V2, bound to a specific deployed contract.
func NewV2(address common.Address, backend bind.ContractBackend) (*V2, error) {
	contract, err := bindV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2{V2Caller: V2Caller{contract: contract}, V2Transactor: V2Transactor{contract: contract}, V2Filterer: V2Filterer{contract: contract}}, nil
}

// NewV2Caller creates a new read-only instance of V2, bound to a specific deployed contract.
func NewV2Caller(address common.Address, caller bind.ContractCaller) (*V2Caller, error) {
	contract, err := bindV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2Caller{contract: contract}, nil
}

// NewV2Transactor creates a new write-only instance of V2, bound to a specific deployed contract.
func NewV2Transactor(address common.Address, transactor bind.ContractTransactor) (*V2Transactor, error) {
	contract, err := bindV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2Transactor{contract: contract}, nil
}

// NewV2Filterer creates a new log filterer instance of V2, bound to a specific deployed contract.
func NewV2Filterer(address common.Address, filterer bind.ContractFilterer) (*V2Filterer, error) {
	contract, err := bindV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2Filterer{contract: contract}, nil
}

// bindV2 binds a generic wrapper to an already deployed contract.
func bindV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2 *V2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2.Contract.V2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2 *V2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2.Contract.V2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2 *V2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2.Contract.V2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2 *V2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2 *V2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2 *V2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_V2 *V2Caller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V2.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_V2 *V2Session) Admins(arg0 common.Address) (*big.Int, error) {
	return _V2.Contract.Admins(&_V2.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_V2 *V2CallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _V2.Contract.Admins(&_V2.CallOpts, arg0)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_V2 *V2Caller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_V2 *V2Session) Collateral() (common.Address, error) {
	return _V2.Contract.Collateral(&_V2.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_V2 *V2CallerSession) Collateral() (common.Address, error) {
	return _V2.Contract.Collateral(&_V2.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_V2 *V2Caller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_V2 *V2Session) Ctf() (common.Address, error) {
	return _V2.Contract.Ctf(&_V2.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_V2 *V2CallerSession) Ctf() (common.Address, error) {
	return _V2.Contract.Ctf(&_V2.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_V2 *V2Caller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_V2 *V2Session) Exchange() (common.Address, error) {
	return _V2.Contract.Exchange(&_V2.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_V2 *V2CallerSession) Exchange() (common.Address, error) {
	return _V2.Contract.Exchange(&_V2.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_V2 *V2Caller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _V2.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_V2 *V2Session) IsAdmin(addr common.Address) (bool, error) {
	return _V2.Contract.IsAdmin(&_V2.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_V2 *V2CallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _V2.Contract.IsAdmin(&_V2.CallOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_V2 *V2Transactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_V2 *V2Session) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _V2.Contract.AddAdmin(&_V2.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_V2 *V2TransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _V2.Contract.AddAdmin(&_V2.TransactOpts, admin)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x2287e350.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256 takerReceiveAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_V2 *V2Transactor) MatchOrders(opts *bind.TransactOpts, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, takerReceiveAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "matchOrders", takerOrder, makerOrders, takerFillAmount, takerReceiveAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x2287e350.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256 takerReceiveAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_V2 *V2Session) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, takerReceiveAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _V2.Contract.MatchOrders(&_V2.TransactOpts, takerOrder, makerOrders, takerFillAmount, takerReceiveAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x2287e350.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256 takerReceiveAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_V2 *V2TransactorSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, takerReceiveAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _V2.Contract.MatchOrders(&_V2.TransactOpts, takerOrder, makerOrders, takerFillAmount, takerReceiveAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_V2 *V2Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_V2 *V2Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.Contract.OnERC1155BatchReceived(&_V2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_V2 *V2TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.Contract.OnERC1155BatchReceived(&_V2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_V2 *V2Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_V2 *V2Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.Contract.OnERC1155Received(&_V2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_V2 *V2TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _V2.Contract.OnERC1155Received(&_V2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_V2 *V2Transactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_V2 *V2Session) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _V2.Contract.RemoveAdmin(&_V2.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_V2 *V2TransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _V2.Contract.RemoveAdmin(&_V2.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_V2 *V2Transactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_V2 *V2Session) RenounceAdmin() (*types.Transaction, error) {
	return _V2.Contract.RenounceAdmin(&_V2.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_V2 *V2TransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _V2.Contract.RenounceAdmin(&_V2.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_V2 *V2Transactor) WithdrawFees(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _V2.contract.Transact(opts, "withdrawFees", to, id, amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_V2 *V2Session) WithdrawFees(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _V2.Contract.WithdrawFees(&_V2.TransactOpts, to, id, amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_V2 *V2TransactorSession) WithdrawFees(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _V2.Contract.WithdrawFees(&_V2.TransactOpts, to, id, amount)
}

// V2FeeRefundedIterator is returned from FilterFeeRefunded and is used to iterate over the raw logs and unpacked data for FeeRefunded events raised by the V2 contract.
type V2FeeRefundedIterator struct {
	Event *V2FeeRefunded // Event containing the contract specifics and raw log

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
func (it *V2FeeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2FeeRefunded)
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
		it.Event = new(V2FeeRefunded)
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
func (it *V2FeeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2FeeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2FeeRefunded represents a FeeRefunded event raised by the V2 contract.
type V2FeeRefunded struct {
	OrderHash  [32]byte
	To         common.Address
	Id         *big.Int
	Refund     *big.Int
	FeeCharged *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeRefunded is a free log retrieval operation binding the contract event 0xb608d2bf25d8b4b744ba23ce2ea9802ea955e216c064a62f42152fbf98958d24.
//
// Solidity: event FeeRefunded(bytes32 indexed orderHash, address indexed to, uint256 id, uint256 refund, uint256 indexed feeCharged)
func (_V2 *V2Filterer) FilterFeeRefunded(opts *bind.FilterOpts, orderHash [][32]byte, to []common.Address, feeCharged []*big.Int) (*V2FeeRefundedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var feeChargedRule []interface{}
	for _, feeChargedItem := range feeCharged {
		feeChargedRule = append(feeChargedRule, feeChargedItem)
	}

	logs, sub, err := _V2.contract.FilterLogs(opts, "FeeRefunded", orderHashRule, toRule, feeChargedRule)
	if err != nil {
		return nil, err
	}
	return &V2FeeRefundedIterator{contract: _V2.contract, event: "FeeRefunded", logs: logs, sub: sub}, nil
}

// WatchFeeRefunded is a free log subscription operation binding the contract event 0xb608d2bf25d8b4b744ba23ce2ea9802ea955e216c064a62f42152fbf98958d24.
//
// Solidity: event FeeRefunded(bytes32 indexed orderHash, address indexed to, uint256 id, uint256 refund, uint256 indexed feeCharged)
func (_V2 *V2Filterer) WatchFeeRefunded(opts *bind.WatchOpts, sink chan<- *V2FeeRefunded, orderHash [][32]byte, to []common.Address, feeCharged []*big.Int) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var feeChargedRule []interface{}
	for _, feeChargedItem := range feeCharged {
		feeChargedRule = append(feeChargedRule, feeChargedItem)
	}

	logs, sub, err := _V2.contract.WatchLogs(opts, "FeeRefunded", orderHashRule, toRule, feeChargedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2FeeRefunded)
				if err := _V2.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
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

// ParseFeeRefunded is a log parse operation binding the contract event 0xb608d2bf25d8b4b744ba23ce2ea9802ea955e216c064a62f42152fbf98958d24.
//
// Solidity: event FeeRefunded(bytes32 indexed orderHash, address indexed to, uint256 id, uint256 refund, uint256 indexed feeCharged)
func (_V2 *V2Filterer) ParseFeeRefunded(log types.Log) (*V2FeeRefunded, error) {
	event := new(V2FeeRefunded)
	if err := _V2.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2FeeWithdrawnIterator is returned from FilterFeeWithdrawn and is used to iterate over the raw logs and unpacked data for FeeWithdrawn events raised by the V2 contract.
type V2FeeWithdrawnIterator struct {
	Event *V2FeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *V2FeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2FeeWithdrawn)
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
		it.Event = new(V2FeeWithdrawn)
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
func (it *V2FeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2FeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2FeeWithdrawn represents a FeeWithdrawn event raised by the V2 contract.
type V2FeeWithdrawn struct {
	Token  common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawn is a free log retrieval operation binding the contract event 0x6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (_V2 *V2Filterer) FilterFeeWithdrawn(opts *bind.FilterOpts) (*V2FeeWithdrawnIterator, error) {

	logs, sub, err := _V2.contract.FilterLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &V2FeeWithdrawnIterator{contract: _V2.contract, event: "FeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawn is a free log subscription operation binding the contract event 0x6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (_V2 *V2Filterer) WatchFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *V2FeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _V2.contract.WatchLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2FeeWithdrawn)
				if err := _V2.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
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

// ParseFeeWithdrawn is a log parse operation binding the contract event 0x6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (_V2 *V2Filterer) ParseFeeWithdrawn(log types.Log) (*V2FeeWithdrawn, error) {
	event := new(V2FeeWithdrawn)
	if err := _V2.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the V2 contract.
type V2NewAdminIterator struct {
	Event *V2NewAdmin // Event containing the contract specifics and raw log

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
func (it *V2NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2NewAdmin)
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
		it.Event = new(V2NewAdmin)
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
func (it *V2NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2NewAdmin represents a NewAdmin event raised by the V2 contract.
type V2NewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_V2 *V2Filterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*V2NewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _V2.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &V2NewAdminIterator{contract: _V2.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_V2 *V2Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *V2NewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _V2.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2NewAdmin)
				if err := _V2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_V2 *V2Filterer) ParseNewAdmin(log types.Log) (*V2NewAdmin, error) {
	event := new(V2NewAdmin)
	if err := _V2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2RemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the V2 contract.
type V2RemovedAdminIterator struct {
	Event *V2RemovedAdmin // Event containing the contract specifics and raw log

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
func (it *V2RemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2RemovedAdmin)
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
		it.Event = new(V2RemovedAdmin)
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
func (it *V2RemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2RemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2RemovedAdmin represents a RemovedAdmin event raised by the V2 contract.
type V2RemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_V2 *V2Filterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*V2RemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _V2.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &V2RemovedAdminIterator{contract: _V2.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_V2 *V2Filterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *V2RemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _V2.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2RemovedAdmin)
				if err := _V2.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_V2 *V2Filterer) ParseRemovedAdmin(log types.Log) (*V2RemovedAdmin, error) {
	event := new(V2RemovedAdmin)
	if err := _V2.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
