// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buildersExecutorV2

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
)

// OrderExecutorV2LimitOrderFillData is an auto generated low-level Go binding around an user-defined struct.
type OrderExecutorV2LimitOrderFillData struct {
	Orders           []OrderTypesLimitOrder
	Signatures       [][]byte
	MakingAmounts    []*big.Int
	TakingAmounts    []*big.Int
	ThresholdAmounts []*big.Int
}

// OrderExecutorV2MarketOrderFillData is an auto generated low-level Go binding around an user-defined struct.
type OrderExecutorV2MarketOrderFillData struct {
	Order     OrderTypesMarketOrder
	Signature []byte
}

// OrderTypesLimitOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesLimitOrder struct {
	Salt           *big.Int
	MakerAsset     common.Address
	TakerAsset     common.Address
	MakerAssetData []byte
	TakerAssetData []byte
	GetMakerAmount []byte
	GetTakerAmount []byte
	Predicate      []byte
	Permit         []byte
	Interaction    []byte
	Signer         common.Address
	SigType        *big.Int
}

// OrderTypesMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesMarketOrder struct {
	Salt         *big.Int
	Signer       common.Address
	Maker        common.Address
	MakerAsset   common.Address
	MakerAmount  *big.Int
	MakerAssetID *big.Int
	TakerAsset   common.Address
	TakerAssetID *big.Int
	SigType      *big.Int
}

// OrderExecutorV2MetaData contains all meta data concerning the OrderExecutorV2 contract.
var OrderExecutorV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"limitOrderProtocolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CollateralTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155AssetWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20AssetWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"Filled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.MarketOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrderExecutorV2.MarketOrderFillData\",\"name\":\"marketOrderFillData\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"makingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"takingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"thresholdAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structOrderExecutorV2.LimitOrderFillData\",\"name\":\"fillData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"conditionalTokenAddress\",\"type\":\"address\"}],\"name\":\"fill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitOrderProtocolContract\",\"outputs\":[{\"internalType\":\"contractLimitOrderProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"name\":\"setCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OrderExecutorV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderExecutorV2MetaData.ABI instead.
var OrderExecutorV2ABI = OrderExecutorV2MetaData.ABI

// OrderExecutorV2 is an auto generated Go binding around an Ethereum contract.
type OrderExecutorV2 struct {
	OrderExecutorV2Caller     // Read-only binding to the contract
	OrderExecutorV2Transactor // Write-only binding to the contract
	OrderExecutorV2Filterer   // Log filterer for contract events
}

// OrderExecutorV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type OrderExecutorV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderExecutorV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderExecutorV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderExecutorV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderExecutorV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderExecutorV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderExecutorV2Session struct {
	Contract     *OrderExecutorV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderExecutorV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderExecutorV2CallerSession struct {
	Contract *OrderExecutorV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OrderExecutorV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderExecutorV2TransactorSession struct {
	Contract     *OrderExecutorV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OrderExecutorV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type OrderExecutorV2Raw struct {
	Contract *OrderExecutorV2 // Generic contract binding to access the raw methods on
}

// OrderExecutorV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderExecutorV2CallerRaw struct {
	Contract *OrderExecutorV2Caller // Generic read-only contract binding to access the raw methods on
}

// OrderExecutorV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderExecutorV2TransactorRaw struct {
	Contract *OrderExecutorV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderExecutorV2 creates a new instance of OrderExecutorV2, bound to a specific deployed contract.
func NewOrderExecutorV2(address common.Address, backend bind.ContractBackend) (*OrderExecutorV2, error) {
	contract, err := bindOrderExecutorV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2{OrderExecutorV2Caller: OrderExecutorV2Caller{contract: contract}, OrderExecutorV2Transactor: OrderExecutorV2Transactor{contract: contract}, OrderExecutorV2Filterer: OrderExecutorV2Filterer{contract: contract}}, nil
}

// NewOrderExecutorV2Caller creates a new read-only instance of OrderExecutorV2, bound to a specific deployed contract.
func NewOrderExecutorV2Caller(address common.Address, caller bind.ContractCaller) (*OrderExecutorV2Caller, error) {
	contract, err := bindOrderExecutorV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2Caller{contract: contract}, nil
}

// NewOrderExecutorV2Transactor creates a new write-only instance of OrderExecutorV2, bound to a specific deployed contract.
func NewOrderExecutorV2Transactor(address common.Address, transactor bind.ContractTransactor) (*OrderExecutorV2Transactor, error) {
	contract, err := bindOrderExecutorV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2Transactor{contract: contract}, nil
}

// NewOrderExecutorV2Filterer creates a new log filterer instance of OrderExecutorV2, bound to a specific deployed contract.
func NewOrderExecutorV2Filterer(address common.Address, filterer bind.ContractFilterer) (*OrderExecutorV2Filterer, error) {
	contract, err := bindOrderExecutorV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2Filterer{contract: contract}, nil
}

// bindOrderExecutorV2 binds a generic wrapper to an already deployed contract.
func bindOrderExecutorV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderExecutorV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderExecutorV2 *OrderExecutorV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderExecutorV2.Contract.OrderExecutorV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderExecutorV2 *OrderExecutorV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OrderExecutorV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderExecutorV2 *OrderExecutorV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OrderExecutorV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderExecutorV2 *OrderExecutorV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderExecutorV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderExecutorV2 *OrderExecutorV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderExecutorV2 *OrderExecutorV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.contract.Transact(opts, method, params...)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Caller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderExecutorV2.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Session) CollateralToken() (common.Address, error) {
	return _OrderExecutorV2.Contract.CollateralToken(&_OrderExecutorV2.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2CallerSession) CollateralToken() (common.Address, error) {
	return _OrderExecutorV2.Contract.CollateralToken(&_OrderExecutorV2.CallOpts)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2Caller) Executed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _OrderExecutorV2.contract.Call(opts, &out, "executed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2Session) Executed(arg0 [32]byte) (bool, error) {
	return _OrderExecutorV2.Contract.Executed(&_OrderExecutorV2.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2CallerSession) Executed(arg0 [32]byte) (bool, error) {
	return _OrderExecutorV2.Contract.Executed(&_OrderExecutorV2.CallOpts, arg0)
}

// LimitOrderProtocolContract is a free data retrieval call binding the contract method 0xf20166f7.
//
// Solidity: function limitOrderProtocolContract() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Caller) LimitOrderProtocolContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderExecutorV2.contract.Call(opts, &out, "limitOrderProtocolContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LimitOrderProtocolContract is a free data retrieval call binding the contract method 0xf20166f7.
//
// Solidity: function limitOrderProtocolContract() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Session) LimitOrderProtocolContract() (common.Address, error) {
	return _OrderExecutorV2.Contract.LimitOrderProtocolContract(&_OrderExecutorV2.CallOpts)
}

// LimitOrderProtocolContract is a free data retrieval call binding the contract method 0xf20166f7.
//
// Solidity: function limitOrderProtocolContract() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2CallerSession) LimitOrderProtocolContract() (common.Address, error) {
	return _OrderExecutorV2.Contract.LimitOrderProtocolContract(&_OrderExecutorV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderExecutorV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2Session) Owner() (common.Address, error) {
	return _OrderExecutorV2.Contract.Owner(&_OrderExecutorV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderExecutorV2 *OrderExecutorV2CallerSession) Owner() (common.Address, error) {
	return _OrderExecutorV2.Contract.Owner(&_OrderExecutorV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OrderExecutorV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OrderExecutorV2.Contract.SupportsInterface(&_OrderExecutorV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OrderExecutorV2 *OrderExecutorV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OrderExecutorV2.Contract.SupportsInterface(&_OrderExecutorV2.CallOpts, interfaceId)
}

// Fill is a paid mutator transaction binding the contract method 0x0b799f76.
//
// Solidity: function fill(((uint256,address,address,address,uint256,uint256,address,uint256,uint256),bytes) marketOrderFillData, ((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[],bytes[],uint256[],uint256[],uint256[]) fillData, address conditionalTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) Fill(opts *bind.TransactOpts, marketOrderFillData OrderExecutorV2MarketOrderFillData, fillData OrderExecutorV2LimitOrderFillData, conditionalTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "fill", marketOrderFillData, fillData, conditionalTokenAddress)
}

// Fill is a paid mutator transaction binding the contract method 0x0b799f76.
//
// Solidity: function fill(((uint256,address,address,address,uint256,uint256,address,uint256,uint256),bytes) marketOrderFillData, ((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[],bytes[],uint256[],uint256[],uint256[]) fillData, address conditionalTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) Fill(marketOrderFillData OrderExecutorV2MarketOrderFillData, fillData OrderExecutorV2LimitOrderFillData, conditionalTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.Fill(&_OrderExecutorV2.TransactOpts, marketOrderFillData, fillData, conditionalTokenAddress)
}

// Fill is a paid mutator transaction binding the contract method 0x0b799f76.
//
// Solidity: function fill(((uint256,address,address,address,uint256,uint256,address,uint256,uint256),bytes) marketOrderFillData, ((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[],bytes[],uint256[],uint256[],uint256[]) fillData, address conditionalTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) Fill(marketOrderFillData OrderExecutorV2MarketOrderFillData, fillData OrderExecutorV2LimitOrderFillData, conditionalTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.Fill(&_OrderExecutorV2.TransactOpts, marketOrderFillData, fillData, conditionalTokenAddress)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OnERC1155BatchReceived(&_OrderExecutorV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OnERC1155BatchReceived(&_OrderExecutorV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OnERC1155Received(&_OrderExecutorV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.OnERC1155Received(&_OrderExecutorV2.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.RenounceOwnership(&_OrderExecutorV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.RenounceOwnership(&_OrderExecutorV2.TransactOpts)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address collateralTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) SetCollateralToken(opts *bind.TransactOpts, collateralTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "setCollateralToken", collateralTokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address collateralTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) SetCollateralToken(collateralTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.SetCollateralToken(&_OrderExecutorV2.TransactOpts, collateralTokenAddress)
}

// SetCollateralToken is a paid mutator transaction binding the contract method 0x666181a9.
//
// Solidity: function setCollateralToken(address collateralTokenAddress) returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) SetCollateralToken(collateralTokenAddress common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.SetCollateralToken(&_OrderExecutorV2.TransactOpts, collateralTokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.TransferOwnership(&_OrderExecutorV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.TransferOwnership(&_OrderExecutorV2.TransactOpts, newOwner)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xa0edb48b.
//
// Solidity: function withdrawERC1155(address dest, address asset, uint256 id, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) WithdrawERC1155(opts *bind.TransactOpts, dest common.Address, asset common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "withdrawERC1155", dest, asset, id, amount)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xa0edb48b.
//
// Solidity: function withdrawERC1155(address dest, address asset, uint256 id, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) WithdrawERC1155(dest common.Address, asset common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.WithdrawERC1155(&_OrderExecutorV2.TransactOpts, dest, asset, id, amount)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xa0edb48b.
//
// Solidity: function withdrawERC1155(address dest, address asset, uint256 id, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) WithdrawERC1155(dest common.Address, asset common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.WithdrawERC1155(&_OrderExecutorV2.TransactOpts, dest, asset, id, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address dest, address asset, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2Transactor) WithdrawERC20(opts *bind.TransactOpts, dest common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.contract.Transact(opts, "withdrawERC20", dest, asset, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address dest, address asset, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2Session) WithdrawERC20(dest common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.WithdrawERC20(&_OrderExecutorV2.TransactOpts, dest, asset, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address dest, address asset, uint256 amount) returns()
func (_OrderExecutorV2 *OrderExecutorV2TransactorSession) WithdrawERC20(dest common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OrderExecutorV2.Contract.WithdrawERC20(&_OrderExecutorV2.TransactOpts, dest, asset, amount)
}

// OrderExecutorV2CollateralTokenChangedIterator is returned from FilterCollateralTokenChanged and is used to iterate over the raw logs and unpacked data for CollateralTokenChanged events raised by the OrderExecutorV2 contract.
type OrderExecutorV2CollateralTokenChangedIterator struct {
	Event *OrderExecutorV2CollateralTokenChanged // Event containing the contract specifics and raw log

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
func (it *OrderExecutorV2CollateralTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderExecutorV2CollateralTokenChanged)
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
		it.Event = new(OrderExecutorV2CollateralTokenChanged)
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
func (it *OrderExecutorV2CollateralTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderExecutorV2CollateralTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderExecutorV2CollateralTokenChanged represents a CollateralTokenChanged event raised by the OrderExecutorV2 contract.
type OrderExecutorV2CollateralTokenChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCollateralTokenChanged is a free log retrieval operation binding the contract event 0xa99c2dd2e7b321e9734aede5fc2a795ab9fc50063aad4d70060dfc01a35fa847.
//
// Solidity: event CollateralTokenChanged(address from, address to)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) FilterCollateralTokenChanged(opts *bind.FilterOpts) (*OrderExecutorV2CollateralTokenChangedIterator, error) {

	logs, sub, err := _OrderExecutorV2.contract.FilterLogs(opts, "CollateralTokenChanged")
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2CollateralTokenChangedIterator{contract: _OrderExecutorV2.contract, event: "CollateralTokenChanged", logs: logs, sub: sub}, nil
}

// WatchCollateralTokenChanged is a free log subscription operation binding the contract event 0xa99c2dd2e7b321e9734aede5fc2a795ab9fc50063aad4d70060dfc01a35fa847.
//
// Solidity: event CollateralTokenChanged(address from, address to)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) WatchCollateralTokenChanged(opts *bind.WatchOpts, sink chan<- *OrderExecutorV2CollateralTokenChanged) (event.Subscription, error) {

	logs, sub, err := _OrderExecutorV2.contract.WatchLogs(opts, "CollateralTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderExecutorV2CollateralTokenChanged)
				if err := _OrderExecutorV2.contract.UnpackLog(event, "CollateralTokenChanged", log); err != nil {
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

// ParseCollateralTokenChanged is a log parse operation binding the contract event 0xa99c2dd2e7b321e9734aede5fc2a795ab9fc50063aad4d70060dfc01a35fa847.
//
// Solidity: event CollateralTokenChanged(address from, address to)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) ParseCollateralTokenChanged(log types.Log) (*OrderExecutorV2CollateralTokenChanged, error) {
	event := new(OrderExecutorV2CollateralTokenChanged)
	if err := _OrderExecutorV2.contract.UnpackLog(event, "CollateralTokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderExecutorV2ERC1155AssetWithdrawnIterator is returned from FilterERC1155AssetWithdrawn and is used to iterate over the raw logs and unpacked data for ERC1155AssetWithdrawn events raised by the OrderExecutorV2 contract.
type OrderExecutorV2ERC1155AssetWithdrawnIterator struct {
	Event *OrderExecutorV2ERC1155AssetWithdrawn // Event containing the contract specifics and raw log

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
func (it *OrderExecutorV2ERC1155AssetWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderExecutorV2ERC1155AssetWithdrawn)
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
		it.Event = new(OrderExecutorV2ERC1155AssetWithdrawn)
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
func (it *OrderExecutorV2ERC1155AssetWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderExecutorV2ERC1155AssetWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderExecutorV2ERC1155AssetWithdrawn represents a ERC1155AssetWithdrawn event raised by the OrderExecutorV2 contract.
type OrderExecutorV2ERC1155AssetWithdrawn struct {
	To     common.Address
	Asset  common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC1155AssetWithdrawn is a free log retrieval operation binding the contract event 0x6da8aacda85cd97244895b9ffd96c4d3e70f3729060ea7b7cea217b9e015b0d6.
//
// Solidity: event ERC1155AssetWithdrawn(address to, address asset, uint256 id, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) FilterERC1155AssetWithdrawn(opts *bind.FilterOpts) (*OrderExecutorV2ERC1155AssetWithdrawnIterator, error) {

	logs, sub, err := _OrderExecutorV2.contract.FilterLogs(opts, "ERC1155AssetWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2ERC1155AssetWithdrawnIterator{contract: _OrderExecutorV2.contract, event: "ERC1155AssetWithdrawn", logs: logs, sub: sub}, nil
}

// WatchERC1155AssetWithdrawn is a free log subscription operation binding the contract event 0x6da8aacda85cd97244895b9ffd96c4d3e70f3729060ea7b7cea217b9e015b0d6.
//
// Solidity: event ERC1155AssetWithdrawn(address to, address asset, uint256 id, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) WatchERC1155AssetWithdrawn(opts *bind.WatchOpts, sink chan<- *OrderExecutorV2ERC1155AssetWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OrderExecutorV2.contract.WatchLogs(opts, "ERC1155AssetWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderExecutorV2ERC1155AssetWithdrawn)
				if err := _OrderExecutorV2.contract.UnpackLog(event, "ERC1155AssetWithdrawn", log); err != nil {
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

// ParseERC1155AssetWithdrawn is a log parse operation binding the contract event 0x6da8aacda85cd97244895b9ffd96c4d3e70f3729060ea7b7cea217b9e015b0d6.
//
// Solidity: event ERC1155AssetWithdrawn(address to, address asset, uint256 id, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) ParseERC1155AssetWithdrawn(log types.Log) (*OrderExecutorV2ERC1155AssetWithdrawn, error) {
	event := new(OrderExecutorV2ERC1155AssetWithdrawn)
	if err := _OrderExecutorV2.contract.UnpackLog(event, "ERC1155AssetWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderExecutorV2ERC20AssetWithdrawnIterator is returned from FilterERC20AssetWithdrawn and is used to iterate over the raw logs and unpacked data for ERC20AssetWithdrawn events raised by the OrderExecutorV2 contract.
type OrderExecutorV2ERC20AssetWithdrawnIterator struct {
	Event *OrderExecutorV2ERC20AssetWithdrawn // Event containing the contract specifics and raw log

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
func (it *OrderExecutorV2ERC20AssetWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderExecutorV2ERC20AssetWithdrawn)
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
		it.Event = new(OrderExecutorV2ERC20AssetWithdrawn)
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
func (it *OrderExecutorV2ERC20AssetWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderExecutorV2ERC20AssetWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderExecutorV2ERC20AssetWithdrawn represents a ERC20AssetWithdrawn event raised by the OrderExecutorV2 contract.
type OrderExecutorV2ERC20AssetWithdrawn struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20AssetWithdrawn is a free log retrieval operation binding the contract event 0x1ccb37f428c9c1cb96e49d15994e15cf07ba02648e3c28b284f4a737923aed53.
//
// Solidity: event ERC20AssetWithdrawn(address to, address asset, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) FilterERC20AssetWithdrawn(opts *bind.FilterOpts) (*OrderExecutorV2ERC20AssetWithdrawnIterator, error) {

	logs, sub, err := _OrderExecutorV2.contract.FilterLogs(opts, "ERC20AssetWithdrawn")
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2ERC20AssetWithdrawnIterator{contract: _OrderExecutorV2.contract, event: "ERC20AssetWithdrawn", logs: logs, sub: sub}, nil
}

// WatchERC20AssetWithdrawn is a free log subscription operation binding the contract event 0x1ccb37f428c9c1cb96e49d15994e15cf07ba02648e3c28b284f4a737923aed53.
//
// Solidity: event ERC20AssetWithdrawn(address to, address asset, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) WatchERC20AssetWithdrawn(opts *bind.WatchOpts, sink chan<- *OrderExecutorV2ERC20AssetWithdrawn) (event.Subscription, error) {

	logs, sub, err := _OrderExecutorV2.contract.WatchLogs(opts, "ERC20AssetWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderExecutorV2ERC20AssetWithdrawn)
				if err := _OrderExecutorV2.contract.UnpackLog(event, "ERC20AssetWithdrawn", log); err != nil {
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

// ParseERC20AssetWithdrawn is a log parse operation binding the contract event 0x1ccb37f428c9c1cb96e49d15994e15cf07ba02648e3c28b284f4a737923aed53.
//
// Solidity: event ERC20AssetWithdrawn(address to, address asset, uint256 amount)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) ParseERC20AssetWithdrawn(log types.Log) (*OrderExecutorV2ERC20AssetWithdrawn, error) {
	event := new(OrderExecutorV2ERC20AssetWithdrawn)
	if err := _OrderExecutorV2.contract.UnpackLog(event, "ERC20AssetWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderExecutorV2FilledIterator is returned from FilterFilled and is used to iterate over the raw logs and unpacked data for Filled events raised by the OrderExecutorV2 contract.
type OrderExecutorV2FilledIterator struct {
	Event *OrderExecutorV2Filled // Event containing the contract specifics and raw log

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
func (it *OrderExecutorV2FilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderExecutorV2Filled)
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
		it.Event = new(OrderExecutorV2Filled)
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
func (it *OrderExecutorV2FilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderExecutorV2FilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderExecutorV2Filled represents a Filled event raised by the OrderExecutorV2 contract.
type OrderExecutorV2Filled struct {
	Caller            common.Address
	Maker             common.Address
	MakerAsset        common.Address
	TakerAsset        common.Address
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFilled is a free log retrieval operation binding the contract event 0x338f9befc70d477e86571de0aa926cdd6a3b43fc5dba40993f76f8b83a10cc59.
//
// Solidity: event Filled(address caller, address maker, address makerAsset, address takerAsset, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) FilterFilled(opts *bind.FilterOpts) (*OrderExecutorV2FilledIterator, error) {

	logs, sub, err := _OrderExecutorV2.contract.FilterLogs(opts, "Filled")
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2FilledIterator{contract: _OrderExecutorV2.contract, event: "Filled", logs: logs, sub: sub}, nil
}

// WatchFilled is a free log subscription operation binding the contract event 0x338f9befc70d477e86571de0aa926cdd6a3b43fc5dba40993f76f8b83a10cc59.
//
// Solidity: event Filled(address caller, address maker, address makerAsset, address takerAsset, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) WatchFilled(opts *bind.WatchOpts, sink chan<- *OrderExecutorV2Filled) (event.Subscription, error) {

	logs, sub, err := _OrderExecutorV2.contract.WatchLogs(opts, "Filled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderExecutorV2Filled)
				if err := _OrderExecutorV2.contract.UnpackLog(event, "Filled", log); err != nil {
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

// ParseFilled is a log parse operation binding the contract event 0x338f9befc70d477e86571de0aa926cdd6a3b43fc5dba40993f76f8b83a10cc59.
//
// Solidity: event Filled(address caller, address maker, address makerAsset, address takerAsset, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) ParseFilled(log types.Log) (*OrderExecutorV2Filled, error) {
	event := new(OrderExecutorV2Filled)
	if err := _OrderExecutorV2.contract.UnpackLog(event, "Filled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderExecutorV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OrderExecutorV2 contract.
type OrderExecutorV2OwnershipTransferredIterator struct {
	Event *OrderExecutorV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrderExecutorV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderExecutorV2OwnershipTransferred)
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
		it.Event = new(OrderExecutorV2OwnershipTransferred)
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
func (it *OrderExecutorV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderExecutorV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderExecutorV2OwnershipTransferred represents a OwnershipTransferred event raised by the OrderExecutorV2 contract.
type OrderExecutorV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrderExecutorV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OrderExecutorV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderExecutorV2OwnershipTransferredIterator{contract: _OrderExecutorV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OrderExecutorV2 *OrderExecutorV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrderExecutorV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OrderExecutorV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderExecutorV2OwnershipTransferred)
				if err := _OrderExecutorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OrderExecutorV2 *OrderExecutorV2Filterer) ParseOwnershipTransferred(log types.Log) (*OrderExecutorV2OwnershipTransferred, error) {
	event := new(OrderExecutorV2OwnershipTransferred)
	if err := _OrderExecutorV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
