// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange_fees

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

// ExchangeFeesMetaData contains all meta data concerning the ExchangeFees contract.
var ExchangeFeesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchange\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeRate\",\"type\":\"uint256\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeFeesABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeFeesMetaData.ABI instead.
var ExchangeFeesABI = ExchangeFeesMetaData.ABI

// ExchangeFees is an auto generated Go binding around an Ethereum contract.
type ExchangeFees struct {
	ExchangeFeesCaller     // Read-only binding to the contract
	ExchangeFeesTransactor // Write-only binding to the contract
	ExchangeFeesFilterer   // Log filterer for contract events
}

// ExchangeFeesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeFeesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFeesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeFeesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFeesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFeesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFeesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeFeesSession struct {
	Contract     *ExchangeFees     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeFeesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeFeesCallerSession struct {
	Contract *ExchangeFeesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExchangeFeesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeFeesTransactorSession struct {
	Contract     *ExchangeFeesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExchangeFeesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeFeesRaw struct {
	Contract *ExchangeFees // Generic contract binding to access the raw methods on
}

// ExchangeFeesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeFeesCallerRaw struct {
	Contract *ExchangeFeesCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeFeesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeFeesTransactorRaw struct {
	Contract *ExchangeFeesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeFees creates a new instance of ExchangeFees, bound to a specific deployed contract.
func NewExchangeFees(address common.Address, backend bind.ContractBackend) (*ExchangeFees, error) {
	contract, err := bindExchangeFees(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeFees{ExchangeFeesCaller: ExchangeFeesCaller{contract: contract}, ExchangeFeesTransactor: ExchangeFeesTransactor{contract: contract}, ExchangeFeesFilterer: ExchangeFeesFilterer{contract: contract}}, nil
}

// NewExchangeFeesCaller creates a new read-only instance of ExchangeFees, bound to a specific deployed contract.
func NewExchangeFeesCaller(address common.Address, caller bind.ContractCaller) (*ExchangeFeesCaller, error) {
	contract, err := bindExchangeFees(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesCaller{contract: contract}, nil
}

// NewExchangeFeesTransactor creates a new write-only instance of ExchangeFees, bound to a specific deployed contract.
func NewExchangeFeesTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeFeesTransactor, error) {
	contract, err := bindExchangeFees(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesTransactor{contract: contract}, nil
}

// NewExchangeFeesFilterer creates a new log filterer instance of ExchangeFees, bound to a specific deployed contract.
func NewExchangeFeesFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFeesFilterer, error) {
	contract, err := bindExchangeFees(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesFilterer{contract: contract}, nil
}

// bindExchangeFees binds a generic wrapper to an already deployed contract.
func bindExchangeFees(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeFeesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeFees *ExchangeFeesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeFees.Contract.ExchangeFeesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeFees *ExchangeFeesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeFees.Contract.ExchangeFeesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeFees *ExchangeFeesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeFees.Contract.ExchangeFeesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeFees *ExchangeFeesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeFees.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeFees *ExchangeFeesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeFees.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeFees *ExchangeFeesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeFees.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_ExchangeFees *ExchangeFeesCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeFees.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_ExchangeFees *ExchangeFeesSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _ExchangeFees.Contract.Admins(&_ExchangeFees.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_ExchangeFees *ExchangeFeesCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _ExchangeFees.Contract.Admins(&_ExchangeFees.CallOpts, arg0)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_ExchangeFees *ExchangeFeesCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeFees.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_ExchangeFees *ExchangeFeesSession) Collateral() (common.Address, error) {
	return _ExchangeFees.Contract.Collateral(&_ExchangeFees.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_ExchangeFees *ExchangeFeesCallerSession) Collateral() (common.Address, error) {
	return _ExchangeFees.Contract.Collateral(&_ExchangeFees.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExchangeFees *ExchangeFeesCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeFees.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExchangeFees *ExchangeFeesSession) Ctf() (common.Address, error) {
	return _ExchangeFees.Contract.Ctf(&_ExchangeFees.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExchangeFees *ExchangeFeesCallerSession) Ctf() (common.Address, error) {
	return _ExchangeFees.Contract.Ctf(&_ExchangeFees.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ExchangeFees *ExchangeFeesCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeFees.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ExchangeFees *ExchangeFeesSession) Exchange() (common.Address, error) {
	return _ExchangeFees.Contract.Exchange(&_ExchangeFees.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_ExchangeFees *ExchangeFeesCallerSession) Exchange() (common.Address, error) {
	return _ExchangeFees.Contract.Exchange(&_ExchangeFees.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_ExchangeFees *ExchangeFeesCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangeFees.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_ExchangeFees *ExchangeFeesSession) IsAdmin(addr common.Address) (bool, error) {
	return _ExchangeFees.Contract.IsAdmin(&_ExchangeFees.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_ExchangeFees *ExchangeFeesCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _ExchangeFees.Contract.IsAdmin(&_ExchangeFees.CallOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.Contract.AddAdmin(&_ExchangeFees.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.Contract.AddAdmin(&_ExchangeFees.TransactOpts, admin)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xd2539b37.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 makerFeeRate) returns()
func (_ExchangeFees *ExchangeFeesTransactor) MatchOrders(opts *bind.TransactOpts, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, makerFeeRate *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts, makerFeeRate)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xd2539b37.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 makerFeeRate) returns()
func (_ExchangeFees *ExchangeFeesSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, makerFeeRate *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.Contract.MatchOrders(&_ExchangeFees.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, makerFeeRate)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xd2539b37.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 makerFeeRate) returns()
func (_ExchangeFees *ExchangeFeesTransactorSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, makerFeeRate *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.Contract.MatchOrders(&_ExchangeFees.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, makerFeeRate)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.Contract.OnERC1155BatchReceived(&_ExchangeFees.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.Contract.OnERC1155BatchReceived(&_ExchangeFees.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.Contract.OnERC1155Received(&_ExchangeFees.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ExchangeFees *ExchangeFeesTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ExchangeFees.Contract.OnERC1155Received(&_ExchangeFees.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.Contract.RemoveAdmin(&_ExchangeFees.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ExchangeFees *ExchangeFeesTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExchangeFees.Contract.RemoveAdmin(&_ExchangeFees.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_ExchangeFees *ExchangeFeesTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_ExchangeFees *ExchangeFeesSession) RenounceAdmin() (*types.Transaction, error) {
	return _ExchangeFees.Contract.RenounceAdmin(&_ExchangeFees.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_ExchangeFees *ExchangeFeesTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _ExchangeFees.Contract.RenounceAdmin(&_ExchangeFees.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_ExchangeFees *ExchangeFeesTransactor) WithdrawFees(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.contract.Transact(opts, "withdrawFees", to, id, amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_ExchangeFees *ExchangeFeesSession) WithdrawFees(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.Contract.WithdrawFees(&_ExchangeFees.TransactOpts, to, id, amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x425c2096.
//
// Solidity: function withdrawFees(address to, uint256 id, uint256 amount) returns()
func (_ExchangeFees *ExchangeFeesTransactorSession) WithdrawFees(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeFees.Contract.WithdrawFees(&_ExchangeFees.TransactOpts, to, id, amount)
}

// ExchangeFeesFeeRefundedIterator is returned from FilterFeeRefunded and is used to iterate over the raw logs and unpacked data for FeeRefunded events raised by the ExchangeFees contract.
type ExchangeFeesFeeRefundedIterator struct {
	Event *ExchangeFeesFeeRefunded // Event containing the contract specifics and raw log

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
func (it *ExchangeFeesFeeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeFeesFeeRefunded)
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
		it.Event = new(ExchangeFeesFeeRefunded)
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
func (it *ExchangeFeesFeeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeFeesFeeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeFeesFeeRefunded represents a FeeRefunded event raised by the ExchangeFees contract.
type ExchangeFeesFeeRefunded struct {
	Token  common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeRefunded is a free log retrieval operation binding the contract event 0x18fe0464eb77016dc4e227eb0d690e4002756d82b442143bbfb874548952b5f2.
//
// Solidity: event FeeRefunded(address token, address to, uint256 id, uint256 amount)
func (_ExchangeFees *ExchangeFeesFilterer) FilterFeeRefunded(opts *bind.FilterOpts) (*ExchangeFeesFeeRefundedIterator, error) {

	logs, sub, err := _ExchangeFees.contract.FilterLogs(opts, "FeeRefunded")
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesFeeRefundedIterator{contract: _ExchangeFees.contract, event: "FeeRefunded", logs: logs, sub: sub}, nil
}

// WatchFeeRefunded is a free log subscription operation binding the contract event 0x18fe0464eb77016dc4e227eb0d690e4002756d82b442143bbfb874548952b5f2.
//
// Solidity: event FeeRefunded(address token, address to, uint256 id, uint256 amount)
func (_ExchangeFees *ExchangeFeesFilterer) WatchFeeRefunded(opts *bind.WatchOpts, sink chan<- *ExchangeFeesFeeRefunded) (event.Subscription, error) {

	logs, sub, err := _ExchangeFees.contract.WatchLogs(opts, "FeeRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeFeesFeeRefunded)
				if err := _ExchangeFees.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
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

// ParseFeeRefunded is a log parse operation binding the contract event 0x18fe0464eb77016dc4e227eb0d690e4002756d82b442143bbfb874548952b5f2.
//
// Solidity: event FeeRefunded(address token, address to, uint256 id, uint256 amount)
func (_ExchangeFees *ExchangeFeesFilterer) ParseFeeRefunded(log types.Log) (*ExchangeFeesFeeRefunded, error) {
	event := new(ExchangeFeesFeeRefunded)
	if err := _ExchangeFees.contract.UnpackLog(event, "FeeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeFeesFeeWithdrawnIterator is returned from FilterFeeWithdrawn and is used to iterate over the raw logs and unpacked data for FeeWithdrawn events raised by the ExchangeFees contract.
type ExchangeFeesFeeWithdrawnIterator struct {
	Event *ExchangeFeesFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ExchangeFeesFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeFeesFeeWithdrawn)
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
		it.Event = new(ExchangeFeesFeeWithdrawn)
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
func (it *ExchangeFeesFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeFeesFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeFeesFeeWithdrawn represents a FeeWithdrawn event raised by the ExchangeFees contract.
type ExchangeFeesFeeWithdrawn struct {
	Token  common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawn is a free log retrieval operation binding the contract event 0x6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (_ExchangeFees *ExchangeFeesFilterer) FilterFeeWithdrawn(opts *bind.FilterOpts) (*ExchangeFeesFeeWithdrawnIterator, error) {

	logs, sub, err := _ExchangeFees.contract.FilterLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesFeeWithdrawnIterator{contract: _ExchangeFees.contract, event: "FeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawn is a free log subscription operation binding the contract event 0x6ce49f8691a80db5eb4f60cd55b14640529346a7ddf9bf8f77a423fa6a10bfdb.
//
// Solidity: event FeeWithdrawn(address token, address to, uint256 id, uint256 amount)
func (_ExchangeFees *ExchangeFeesFilterer) WatchFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *ExchangeFeesFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _ExchangeFees.contract.WatchLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeFeesFeeWithdrawn)
				if err := _ExchangeFees.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
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
func (_ExchangeFees *ExchangeFeesFilterer) ParseFeeWithdrawn(log types.Log) (*ExchangeFeesFeeWithdrawn, error) {
	event := new(ExchangeFeesFeeWithdrawn)
	if err := _ExchangeFees.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeFeesNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the ExchangeFees contract.
type ExchangeFeesNewAdminIterator struct {
	Event *ExchangeFeesNewAdmin // Event containing the contract specifics and raw log

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
func (it *ExchangeFeesNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeFeesNewAdmin)
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
		it.Event = new(ExchangeFeesNewAdmin)
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
func (it *ExchangeFeesNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeFeesNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeFeesNewAdmin represents a NewAdmin event raised by the ExchangeFees contract.
type ExchangeFeesNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_ExchangeFees *ExchangeFeesFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*ExchangeFeesNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _ExchangeFees.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesNewAdminIterator{contract: _ExchangeFees.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_ExchangeFees *ExchangeFeesFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *ExchangeFeesNewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _ExchangeFees.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeFeesNewAdmin)
				if err := _ExchangeFees.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_ExchangeFees *ExchangeFeesFilterer) ParseNewAdmin(log types.Log) (*ExchangeFeesNewAdmin, error) {
	event := new(ExchangeFeesNewAdmin)
	if err := _ExchangeFees.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeFeesRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the ExchangeFees contract.
type ExchangeFeesRemovedAdminIterator struct {
	Event *ExchangeFeesRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *ExchangeFeesRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeFeesRemovedAdmin)
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
		it.Event = new(ExchangeFeesRemovedAdmin)
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
func (it *ExchangeFeesRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeFeesRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeFeesRemovedAdmin represents a RemovedAdmin event raised by the ExchangeFees contract.
type ExchangeFeesRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_ExchangeFees *ExchangeFeesFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*ExchangeFeesRemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _ExchangeFees.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeFeesRemovedAdminIterator{contract: _ExchangeFees.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_ExchangeFees *ExchangeFeesFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *ExchangeFeesRemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _ExchangeFees.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeFeesRemovedAdmin)
				if err := _ExchangeFees.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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
func (_ExchangeFees *ExchangeFeesFilterer) ParseRemovedAdmin(log types.Log) (*ExchangeFeesRemovedAdmin, error) {
	event := new(ExchangeFeesRemovedAdmin)
	if err := _ExchangeFees.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
