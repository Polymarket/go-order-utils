// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neg_risk

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

// NegRiskMetaData contains all meta data concerning the NegRisk contract.
var NegRiskMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ctf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FEE_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NO_TOKEN_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfBatch\",\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"col\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertPositions\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_indexSet\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ctf\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIConditionalTokens\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConditionId\",\"inputs\":[{\"name\":\"_questionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDetermined\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBips\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketData\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"MarketData\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOracle\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionId\",\"inputs\":[{\"name\":\"_questionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_outcome\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuestionCount\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResult\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mergePositions\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mergePositions\",\"inputs\":[{\"name\":\"_conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prepareMarket\",\"inputs\":[{\"name\":\"_feeBips\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prepareQuestion\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemPositions\",\"inputs\":[{\"name\":\"_conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reportOutcome\",\"inputs\":[{\"name\":\"_questionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_outcome\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"splitPosition\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"splitPosition\",\"inputs\":[{\"name\":\"_conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wcol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractWrappedCollateral\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MarketPrepared\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeBips\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdminAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutcomeReported\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"questionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"outcome\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayoutRedemption\",\"inputs\":[{\"name\":\"redeemer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"conditionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"payout\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionSplit\",\"inputs\":[{\"name\":\"stakeholder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"conditionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionsConverted\",\"inputs\":[{\"name\":\"stakeholder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"marketId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"indexSet\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionsMerge\",\"inputs\":[{\"name\":\"stakeholder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"conditionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuestionPrepared\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"questionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"removedAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DeterminedFlagAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeBipsOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IndexOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIndexSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketAlreadyDetermined\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketAlreadyPrepared\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketNotPrepared\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoConvertiblePositions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotApprovedForAll\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedCollateralToken\",\"inputs\":[]}]",
}

// NegRiskABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskMetaData.ABI instead.
var NegRiskABI = NegRiskMetaData.ABI

// NegRisk is an auto generated Go binding around an Ethereum contract.
type NegRisk struct {
	NegRiskCaller     // Read-only binding to the contract
	NegRiskTransactor // Write-only binding to the contract
	NegRiskFilterer   // Log filterer for contract events
}

// NegRiskCaller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskSession struct {
	Contract     *NegRisk          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NegRiskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskCallerSession struct {
	Contract *NegRiskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NegRiskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskTransactorSession struct {
	Contract     *NegRiskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NegRiskRaw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskRaw struct {
	Contract *NegRisk // Generic contract binding to access the raw methods on
}

// NegRiskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskCallerRaw struct {
	Contract *NegRiskCaller // Generic read-only contract binding to access the raw methods on
}

// NegRiskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskTransactorRaw struct {
	Contract *NegRiskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRisk creates a new instance of NegRisk, bound to a specific deployed contract.
func NewNegRisk(address common.Address, backend bind.ContractBackend) (*NegRisk, error) {
	contract, err := bindNegRisk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRisk{NegRiskCaller: NegRiskCaller{contract: contract}, NegRiskTransactor: NegRiskTransactor{contract: contract}, NegRiskFilterer: NegRiskFilterer{contract: contract}}, nil
}

// NewNegRiskCaller creates a new read-only instance of NegRisk, bound to a specific deployed contract.
func NewNegRiskCaller(address common.Address, caller bind.ContractCaller) (*NegRiskCaller, error) {
	contract, err := bindNegRisk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCaller{contract: contract}, nil
}

// NewNegRiskTransactor creates a new write-only instance of NegRisk, bound to a specific deployed contract.
func NewNegRiskTransactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskTransactor, error) {
	contract, err := bindNegRisk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskTransactor{contract: contract}, nil
}

// NewNegRiskFilterer creates a new log filterer instance of NegRisk, bound to a specific deployed contract.
func NewNegRiskFilterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskFilterer, error) {
	contract, err := bindNegRisk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskFilterer{contract: contract}, nil
}

// bindNegRisk binds a generic wrapper to an already deployed contract.
func bindNegRisk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRisk *NegRiskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRisk.Contract.NegRiskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRisk *NegRiskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRisk.Contract.NegRiskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRisk *NegRiskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRisk.Contract.NegRiskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRisk *NegRiskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRisk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRisk *NegRiskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRisk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRisk *NegRiskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRisk.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRisk *NegRiskCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRisk *NegRiskSession) FEEDENOMINATOR() (*big.Int, error) {
	return _NegRisk.Contract.FEEDENOMINATOR(&_NegRisk.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_NegRisk *NegRiskCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _NegRisk.Contract.FEEDENOMINATOR(&_NegRisk.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRisk *NegRiskCaller) NOTOKENBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "NO_TOKEN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRisk *NegRiskSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _NegRisk.Contract.NOTOKENBURNADDRESS(&_NegRisk.CallOpts)
}

// NOTOKENBURNADDRESS is a free data retrieval call binding the contract method 0x7ad7fe36.
//
// Solidity: function NO_TOKEN_BURN_ADDRESS() view returns(address)
func (_NegRisk *NegRiskCallerSession) NOTOKENBURNADDRESS() (common.Address, error) {
	return _NegRisk.Contract.NOTOKENBURNADDRESS(&_NegRisk.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRisk *NegRiskCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRisk *NegRiskSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRisk.Contract.Admins(&_NegRisk.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRisk.Contract.Admins(&_NegRisk.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRisk *NegRiskCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRisk *NegRiskSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _NegRisk.Contract.BalanceOf(&_NegRisk.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _NegRisk.Contract.BalanceOf(&_NegRisk.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRisk *NegRiskCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRisk *NegRiskSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _NegRisk.Contract.BalanceOfBatch(&_NegRisk.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_NegRisk *NegRiskCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _NegRisk.Contract.BalanceOfBatch(&_NegRisk.CallOpts, _owners, _ids)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRisk *NegRiskCaller) Col(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "col")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRisk *NegRiskSession) Col() (common.Address, error) {
	return _NegRisk.Contract.Col(&_NegRisk.CallOpts)
}

// Col is a free data retrieval call binding the contract method 0xa78695b0.
//
// Solidity: function col() view returns(address)
func (_NegRisk *NegRiskCallerSession) Col() (common.Address, error) {
	return _NegRisk.Contract.Col(&_NegRisk.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRisk *NegRiskCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRisk *NegRiskSession) Ctf() (common.Address, error) {
	return _NegRisk.Contract.Ctf(&_NegRisk.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_NegRisk *NegRiskCallerSession) Ctf() (common.Address, error) {
	return _NegRisk.Contract.Ctf(&_NegRisk.CallOpts)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRisk *NegRiskCaller) GetConditionId(opts *bind.CallOpts, _questionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getConditionId", _questionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRisk *NegRiskSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _NegRisk.Contract.GetConditionId(&_NegRisk.CallOpts, _questionId)
}

// GetConditionId is a free data retrieval call binding the contract method 0x04329c03.
//
// Solidity: function getConditionId(bytes32 _questionId) view returns(bytes32)
func (_NegRisk *NegRiskCallerSession) GetConditionId(_questionId [32]byte) ([32]byte, error) {
	return _NegRisk.Contract.GetConditionId(&_NegRisk.CallOpts, _questionId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRisk *NegRiskCaller) GetDetermined(opts *bind.CallOpts, _marketId [32]byte) (bool, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getDetermined", _marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRisk *NegRiskSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _NegRisk.Contract.GetDetermined(&_NegRisk.CallOpts, _marketId)
}

// GetDetermined is a free data retrieval call binding the contract method 0x7ae2e67b.
//
// Solidity: function getDetermined(bytes32 _marketId) view returns(bool)
func (_NegRisk *NegRiskCallerSession) GetDetermined(_marketId [32]byte) (bool, error) {
	return _NegRisk.Contract.GetDetermined(&_NegRisk.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCaller) GetFeeBips(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getFeeBips", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetFeeBips(&_NegRisk.CallOpts, _marketId)
}

// GetFeeBips is a free data retrieval call binding the contract method 0x2582cb5e.
//
// Solidity: function getFeeBips(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) GetFeeBips(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetFeeBips(&_NegRisk.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRisk *NegRiskCaller) GetMarketData(opts *bind.CallOpts, _marketId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getMarketData", _marketId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRisk *NegRiskSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _NegRisk.Contract.GetMarketData(&_NegRisk.CallOpts, _marketId)
}

// GetMarketData is a free data retrieval call binding the contract method 0x30f4f4bb.
//
// Solidity: function getMarketData(bytes32 _marketId) view returns(bytes32)
func (_NegRisk *NegRiskCallerSession) GetMarketData(_marketId [32]byte) ([32]byte, error) {
	return _NegRisk.Contract.GetMarketData(&_NegRisk.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRisk *NegRiskCaller) GetOracle(opts *bind.CallOpts, _marketId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getOracle", _marketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRisk *NegRiskSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _NegRisk.Contract.GetOracle(&_NegRisk.CallOpts, _marketId)
}

// GetOracle is a free data retrieval call binding the contract method 0xdafaf94a.
//
// Solidity: function getOracle(bytes32 _marketId) view returns(address)
func (_NegRisk *NegRiskCallerSession) GetOracle(_marketId [32]byte) (common.Address, error) {
	return _NegRisk.Contract.GetOracle(&_NegRisk.CallOpts, _marketId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRisk *NegRiskCaller) GetPositionId(opts *bind.CallOpts, _questionId [32]byte, _outcome bool) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getPositionId", _questionId, _outcome)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRisk *NegRiskSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _NegRisk.Contract.GetPositionId(&_NegRisk.CallOpts, _questionId, _outcome)
}

// GetPositionId is a free data retrieval call binding the contract method 0x752b5ba5.
//
// Solidity: function getPositionId(bytes32 _questionId, bool _outcome) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) GetPositionId(_questionId [32]byte, _outcome bool) (*big.Int, error) {
	return _NegRisk.Contract.GetPositionId(&_NegRisk.CallOpts, _questionId, _outcome)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCaller) GetQuestionCount(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getQuestionCount", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetQuestionCount(&_NegRisk.CallOpts, _marketId)
}

// GetQuestionCount is a free data retrieval call binding the contract method 0xb7f75d2c.
//
// Solidity: function getQuestionCount(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) GetQuestionCount(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetQuestionCount(&_NegRisk.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCaller) GetResult(opts *bind.CallOpts, _marketId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "getResult", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetResult(&_NegRisk.CallOpts, _marketId)
}

// GetResult is a free data retrieval call binding the contract method 0xadd4c784.
//
// Solidity: function getResult(bytes32 _marketId) view returns(uint256)
func (_NegRisk *NegRiskCallerSession) GetResult(_marketId [32]byte) (*big.Int, error) {
	return _NegRisk.Contract.GetResult(&_NegRisk.CallOpts, _marketId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRisk *NegRiskCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "isAdmin", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRisk *NegRiskSession) IsAdmin(addr common.Address) (bool, error) {
	return _NegRisk.Contract.IsAdmin(&_NegRisk.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address addr) view returns(bool)
func (_NegRisk *NegRiskCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _NegRisk.Contract.IsAdmin(&_NegRisk.CallOpts, addr)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRisk *NegRiskCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRisk *NegRiskSession) Vault() (common.Address, error) {
	return _NegRisk.Contract.Vault(&_NegRisk.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_NegRisk *NegRiskCallerSession) Vault() (common.Address, error) {
	return _NegRisk.Contract.Vault(&_NegRisk.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRisk *NegRiskCaller) Wcol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRisk.contract.Call(opts, &out, "wcol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRisk *NegRiskSession) Wcol() (common.Address, error) {
	return _NegRisk.Contract.Wcol(&_NegRisk.CallOpts)
}

// Wcol is a free data retrieval call binding the contract method 0x7e3b74c3.
//
// Solidity: function wcol() view returns(address)
func (_NegRisk *NegRiskCallerSession) Wcol() (common.Address, error) {
	return _NegRisk.Contract.Wcol(&_NegRisk.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRisk *NegRiskTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRisk *NegRiskSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRisk.Contract.AddAdmin(&_NegRisk.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_NegRisk *NegRiskTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRisk.Contract.AddAdmin(&_NegRisk.TransactOpts, admin)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactor) ConvertPositions(opts *bind.TransactOpts, _marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "convertPositions", _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRisk *NegRiskSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.ConvertPositions(&_NegRisk.TransactOpts, _marketId, _indexSet, _amount)
}

// ConvertPositions is a paid mutator transaction binding the contract method 0xc64748c4.
//
// Solidity: function convertPositions(bytes32 _marketId, uint256 _indexSet, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactorSession) ConvertPositions(_marketId [32]byte, _indexSet *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.ConvertPositions(&_NegRisk.TransactOpts, _marketId, _indexSet, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskTransactor) MergePositions(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "mergePositions", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.MergePositions(&_NegRisk.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskTransactorSession) MergePositions(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.MergePositions(&_NegRisk.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactor) MergePositions0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "mergePositions0", _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.MergePositions0(&_NegRisk.TransactOpts, _conditionId, _amount)
}

// MergePositions0 is a paid mutator transaction binding the contract method 0xb10c5c17.
//
// Solidity: function mergePositions(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactorSession) MergePositions0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.MergePositions0(&_NegRisk.TransactOpts, _conditionId, _amount)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRisk *NegRiskTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRisk *NegRiskSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.OnERC1155BatchReceived(&_NegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRisk *NegRiskTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.OnERC1155BatchReceived(&_NegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRisk *NegRiskTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRisk *NegRiskSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.OnERC1155Received(&_NegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRisk *NegRiskTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.OnERC1155Received(&_NegRisk.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskTransactor) PrepareMarket(opts *bind.TransactOpts, _feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "prepareMarket", _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.PrepareMarket(&_NegRisk.TransactOpts, _feeBips, _metadata)
}

// PrepareMarket is a paid mutator transaction binding the contract method 0x8a0db615.
//
// Solidity: function prepareMarket(uint256 _feeBips, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskTransactorSession) PrepareMarket(_feeBips *big.Int, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.PrepareMarket(&_NegRisk.TransactOpts, _feeBips, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskTransactor) PrepareQuestion(opts *bind.TransactOpts, _marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "prepareQuestion", _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.PrepareQuestion(&_NegRisk.TransactOpts, _marketId, _metadata)
}

// PrepareQuestion is a paid mutator transaction binding the contract method 0x1d69b48d.
//
// Solidity: function prepareQuestion(bytes32 _marketId, bytes _metadata) returns(bytes32)
func (_NegRisk *NegRiskTransactorSession) PrepareQuestion(_marketId [32]byte, _metadata []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.PrepareQuestion(&_NegRisk.TransactOpts, _marketId, _metadata)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRisk *NegRiskTransactor) RedeemPositions(opts *bind.TransactOpts, _conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "redeemPositions", _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRisk *NegRiskSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.RedeemPositions(&_NegRisk.TransactOpts, _conditionId, _amounts)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0xdbeccb23.
//
// Solidity: function redeemPositions(bytes32 _conditionId, uint256[] _amounts) returns()
func (_NegRisk *NegRiskTransactorSession) RedeemPositions(_conditionId [32]byte, _amounts []*big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.RedeemPositions(&_NegRisk.TransactOpts, _conditionId, _amounts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRisk *NegRiskTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRisk *NegRiskSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRisk.Contract.RemoveAdmin(&_NegRisk.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRisk *NegRiskTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRisk.Contract.RemoveAdmin(&_NegRisk.TransactOpts, admin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRisk *NegRiskTransactor) RenounceAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "renounceAdmin")
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRisk *NegRiskSession) RenounceAdmin() (*types.Transaction, error) {
	return _NegRisk.Contract.RenounceAdmin(&_NegRisk.TransactOpts)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x8bad0c0a.
//
// Solidity: function renounceAdmin() returns()
func (_NegRisk *NegRiskTransactorSession) RenounceAdmin() (*types.Transaction, error) {
	return _NegRisk.Contract.RenounceAdmin(&_NegRisk.TransactOpts)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRisk *NegRiskTransactor) ReportOutcome(opts *bind.TransactOpts, _questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "reportOutcome", _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRisk *NegRiskSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRisk.Contract.ReportOutcome(&_NegRisk.TransactOpts, _questionId, _outcome)
}

// ReportOutcome is a paid mutator transaction binding the contract method 0xe200affd.
//
// Solidity: function reportOutcome(bytes32 _questionId, bool _outcome) returns()
func (_NegRisk *NegRiskTransactorSession) ReportOutcome(_questionId [32]byte, _outcome bool) (*types.Transaction, error) {
	return _NegRisk.Contract.ReportOutcome(&_NegRisk.TransactOpts, _questionId, _outcome)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRisk *NegRiskTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRisk *NegRiskSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.SafeTransferFrom(&_NegRisk.TransactOpts, _from, _to, _id, _value, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _value, bytes _data) returns()
func (_NegRisk *NegRiskTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _NegRisk.Contract.SafeTransferFrom(&_NegRisk.TransactOpts, _from, _to, _id, _value, _data)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskTransactor) SplitPosition(opts *bind.TransactOpts, _collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "splitPosition", _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.SplitPosition(&_NegRisk.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address _collateralToken, bytes32 , bytes32 _conditionId, uint256[] , uint256 _amount) returns()
func (_NegRisk *NegRiskTransactorSession) SplitPosition(_collateralToken common.Address, arg1 [32]byte, _conditionId [32]byte, arg3 []*big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.SplitPosition(&_NegRisk.TransactOpts, _collateralToken, arg1, _conditionId, arg3, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactor) SplitPosition0(opts *bind.TransactOpts, _conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.contract.Transact(opts, "splitPosition0", _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.SplitPosition0(&_NegRisk.TransactOpts, _conditionId, _amount)
}

// SplitPosition0 is a paid mutator transaction binding the contract method 0xa3d7da1d.
//
// Solidity: function splitPosition(bytes32 _conditionId, uint256 _amount) returns()
func (_NegRisk *NegRiskTransactorSession) SplitPosition0(_conditionId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _NegRisk.Contract.SplitPosition0(&_NegRisk.TransactOpts, _conditionId, _amount)
}

// NegRiskMarketPreparedIterator is returned from FilterMarketPrepared and is used to iterate over the raw logs and unpacked data for MarketPrepared events raised by the NegRisk contract.
type NegRiskMarketPreparedIterator struct {
	Event *NegRiskMarketPrepared // Event containing the contract specifics and raw log

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
func (it *NegRiskMarketPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskMarketPrepared)
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
		it.Event = new(NegRiskMarketPrepared)
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
func (it *NegRiskMarketPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskMarketPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskMarketPrepared represents a MarketPrepared event raised by the NegRisk contract.
type NegRiskMarketPrepared struct {
	MarketId [32]byte
	Oracle   common.Address
	FeeBips  *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketPrepared is a free log retrieval operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRisk *NegRiskFilterer) FilterMarketPrepared(opts *bind.FilterOpts, marketId [][32]byte, oracle []common.Address) (*NegRiskMarketPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskMarketPreparedIterator{contract: _NegRisk.contract, event: "MarketPrepared", logs: logs, sub: sub}, nil
}

// WatchMarketPrepared is a free log subscription operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRisk *NegRiskFilterer) WatchMarketPrepared(opts *bind.WatchOpts, sink chan<- *NegRiskMarketPrepared, marketId [][32]byte, oracle []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "MarketPrepared", marketIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskMarketPrepared)
				if err := _NegRisk.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
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

// ParseMarketPrepared is a log parse operation binding the contract event 0xf059ab16d1ca60e123eab60e3c02b68faf060347c701a5d14885a8e1def7b3a8.
//
// Solidity: event MarketPrepared(bytes32 indexed marketId, address indexed oracle, uint256 feeBips, bytes data)
func (_NegRisk *NegRiskFilterer) ParseMarketPrepared(log types.Log) (*NegRiskMarketPrepared, error) {
	event := new(NegRiskMarketPrepared)
	if err := _NegRisk.contract.UnpackLog(event, "MarketPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NegRisk contract.
type NegRiskNewAdminIterator struct {
	Event *NegRiskNewAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskNewAdmin)
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
		it.Event = new(NegRiskNewAdmin)
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
func (it *NegRiskNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskNewAdmin represents a NewAdmin event raised by the NegRisk contract.
type NegRiskNewAdmin struct {
	Admin           common.Address
	NewAdminAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_NegRisk *NegRiskFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address, newAdminAddress []common.Address) (*NegRiskNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskNewAdminIterator{contract: _NegRisk.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed admin, address indexed newAdminAddress)
func (_NegRisk *NegRiskFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskNewAdmin, admin []common.Address, newAdminAddress []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "NewAdmin", adminRule, newAdminAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskNewAdmin)
				if err := _NegRisk.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_NegRisk *NegRiskFilterer) ParseNewAdmin(log types.Log) (*NegRiskNewAdmin, error) {
	event := new(NegRiskNewAdmin)
	if err := _NegRisk.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskOutcomeReportedIterator is returned from FilterOutcomeReported and is used to iterate over the raw logs and unpacked data for OutcomeReported events raised by the NegRisk contract.
type NegRiskOutcomeReportedIterator struct {
	Event *NegRiskOutcomeReported // Event containing the contract specifics and raw log

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
func (it *NegRiskOutcomeReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskOutcomeReported)
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
		it.Event = new(NegRiskOutcomeReported)
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
func (it *NegRiskOutcomeReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskOutcomeReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskOutcomeReported represents a OutcomeReported event raised by the NegRisk contract.
type NegRiskOutcomeReported struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Outcome    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOutcomeReported is a free log retrieval operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRisk *NegRiskFilterer) FilterOutcomeReported(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*NegRiskOutcomeReportedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskOutcomeReportedIterator{contract: _NegRisk.contract, event: "OutcomeReported", logs: logs, sub: sub}, nil
}

// WatchOutcomeReported is a free log subscription operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRisk *NegRiskFilterer) WatchOutcomeReported(opts *bind.WatchOpts, sink chan<- *NegRiskOutcomeReported, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "OutcomeReported", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskOutcomeReported)
				if err := _NegRisk.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
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

// ParseOutcomeReported is a log parse operation binding the contract event 0x9e9fa7fd355160bd4cd3f22d4333519354beff1f5689bde87f2c5e63d8d484b2.
//
// Solidity: event OutcomeReported(bytes32 indexed marketId, bytes32 indexed questionId, bool outcome)
func (_NegRisk *NegRiskFilterer) ParseOutcomeReported(log types.Log) (*NegRiskOutcomeReported, error) {
	event := new(NegRiskOutcomeReported)
	if err := _NegRisk.contract.UnpackLog(event, "OutcomeReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the NegRisk contract.
type NegRiskPayoutRedemptionIterator struct {
	Event *NegRiskPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *NegRiskPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskPayoutRedemption)
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
		it.Event = new(NegRiskPayoutRedemption)
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
func (it *NegRiskPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskPayoutRedemption represents a PayoutRedemption event raised by the NegRisk contract.
type NegRiskPayoutRedemption struct {
	Redeemer    common.Address
	ConditionId [32]byte
	Amounts     []*big.Int
	Payout      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRisk *NegRiskFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, conditionId [][32]byte) (*NegRiskPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskPayoutRedemptionIterator{contract: _NegRisk.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRisk *NegRiskFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *NegRiskPayoutRedemption, redeemer []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskPayoutRedemption)
				if err := _NegRisk.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x9140a6a270ef945260c03894b3c6b3b2695e9d5101feef0ff24fec960cfd3224.
//
// Solidity: event PayoutRedemption(address indexed redeemer, bytes32 indexed conditionId, uint256[] amounts, uint256 payout)
func (_NegRisk *NegRiskFilterer) ParsePayoutRedemption(log types.Log) (*NegRiskPayoutRedemption, error) {
	event := new(NegRiskPayoutRedemption)
	if err := _NegRisk.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the NegRisk contract.
type NegRiskPositionSplitIterator struct {
	Event *NegRiskPositionSplit // Event containing the contract specifics and raw log

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
func (it *NegRiskPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskPositionSplit)
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
		it.Event = new(NegRiskPositionSplit)
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
func (it *NegRiskPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskPositionSplit represents a PositionSplit event raised by the NegRisk contract.
type NegRiskPositionSplit struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*NegRiskPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskPositionSplitIterator{contract: _NegRisk.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *NegRiskPositionSplit, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskPositionSplit)
				if err := _NegRisk.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0xbbed930dbfb7907ae2d60ddf78345610214f26419a0128df39b6cc3d9e5df9b0.
//
// Solidity: event PositionSplit(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) ParsePositionSplit(log types.Log) (*NegRiskPositionSplit, error) {
	event := new(NegRiskPositionSplit)
	if err := _NegRisk.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskPositionsConvertedIterator is returned from FilterPositionsConverted and is used to iterate over the raw logs and unpacked data for PositionsConverted events raised by the NegRisk contract.
type NegRiskPositionsConvertedIterator struct {
	Event *NegRiskPositionsConverted // Event containing the contract specifics and raw log

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
func (it *NegRiskPositionsConvertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskPositionsConverted)
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
		it.Event = new(NegRiskPositionsConverted)
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
func (it *NegRiskPositionsConvertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskPositionsConvertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskPositionsConverted represents a PositionsConverted event raised by the NegRisk contract.
type NegRiskPositionsConverted struct {
	Stakeholder common.Address
	MarketId    [32]byte
	IndexSet    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsConverted is a free log retrieval operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRisk *NegRiskFilterer) FilterPositionsConverted(opts *bind.FilterOpts, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (*NegRiskPositionsConvertedIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskPositionsConvertedIterator{contract: _NegRisk.contract, event: "PositionsConverted", logs: logs, sub: sub}, nil
}

// WatchPositionsConverted is a free log subscription operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRisk *NegRiskFilterer) WatchPositionsConverted(opts *bind.WatchOpts, sink chan<- *NegRiskPositionsConverted, stakeholder []common.Address, marketId [][32]byte, indexSet []*big.Int) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var indexSetRule []interface{}
	for _, indexSetItem := range indexSet {
		indexSetRule = append(indexSetRule, indexSetItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "PositionsConverted", stakeholderRule, marketIdRule, indexSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskPositionsConverted)
				if err := _NegRisk.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
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

// ParsePositionsConverted is a log parse operation binding the contract event 0xb03d19dddbc72a87e735ff0ea3b57bef133ebe44e1894284916a84044deb367e.
//
// Solidity: event PositionsConverted(address indexed stakeholder, bytes32 indexed marketId, uint256 indexed indexSet, uint256 amount)
func (_NegRisk *NegRiskFilterer) ParsePositionsConverted(log types.Log) (*NegRiskPositionsConverted, error) {
	event := new(NegRiskPositionsConverted)
	if err := _NegRisk.contract.UnpackLog(event, "PositionsConverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the NegRisk contract.
type NegRiskPositionsMergeIterator struct {
	Event *NegRiskPositionsMerge // Event containing the contract specifics and raw log

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
func (it *NegRiskPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskPositionsMerge)
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
		it.Event = new(NegRiskPositionsMerge)
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
func (it *NegRiskPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskPositionsMerge represents a PositionsMerge event raised by the NegRisk contract.
type NegRiskPositionsMerge struct {
	Stakeholder common.Address
	ConditionId [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, conditionId [][32]byte) (*NegRiskPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskPositionsMergeIterator{contract: _NegRisk.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *NegRiskPositionsMerge, stakeholder []common.Address, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskPositionsMerge)
				if err := _NegRisk.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0xba33ac50d8894676597e6e35dc09cff59854708b642cd069d21eb9c7ca072a04.
//
// Solidity: event PositionsMerge(address indexed stakeholder, bytes32 indexed conditionId, uint256 amount)
func (_NegRisk *NegRiskFilterer) ParsePositionsMerge(log types.Log) (*NegRiskPositionsMerge, error) {
	event := new(NegRiskPositionsMerge)
	if err := _NegRisk.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskQuestionPreparedIterator is returned from FilterQuestionPrepared and is used to iterate over the raw logs and unpacked data for QuestionPrepared events raised by the NegRisk contract.
type NegRiskQuestionPreparedIterator struct {
	Event *NegRiskQuestionPrepared // Event containing the contract specifics and raw log

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
func (it *NegRiskQuestionPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskQuestionPrepared)
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
		it.Event = new(NegRiskQuestionPrepared)
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
func (it *NegRiskQuestionPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskQuestionPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskQuestionPrepared represents a QuestionPrepared event raised by the NegRisk contract.
type NegRiskQuestionPrepared struct {
	MarketId   [32]byte
	QuestionId [32]byte
	Index      *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuestionPrepared is a free log retrieval operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRisk *NegRiskFilterer) FilterQuestionPrepared(opts *bind.FilterOpts, marketId [][32]byte, questionId [][32]byte) (*NegRiskQuestionPreparedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskQuestionPreparedIterator{contract: _NegRisk.contract, event: "QuestionPrepared", logs: logs, sub: sub}, nil
}

// WatchQuestionPrepared is a free log subscription operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRisk *NegRiskFilterer) WatchQuestionPrepared(opts *bind.WatchOpts, sink chan<- *NegRiskQuestionPrepared, marketId [][32]byte, questionId [][32]byte) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "QuestionPrepared", marketIdRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskQuestionPrepared)
				if err := _NegRisk.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
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

// ParseQuestionPrepared is a log parse operation binding the contract event 0xaac410f87d423a922a7b226ac68f0c2eaf5bf6d15e644ac0758c7f96e2c253f7.
//
// Solidity: event QuestionPrepared(bytes32 indexed marketId, bytes32 indexed questionId, uint256 index, bytes data)
func (_NegRisk *NegRiskFilterer) ParseQuestionPrepared(log types.Log) (*NegRiskQuestionPrepared, error) {
	event := new(NegRiskQuestionPrepared)
	if err := _NegRisk.contract.UnpackLog(event, "QuestionPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the NegRisk contract.
type NegRiskRemovedAdminIterator struct {
	Event *NegRiskRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskRemovedAdmin)
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
		it.Event = new(NegRiskRemovedAdmin)
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
func (it *NegRiskRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskRemovedAdmin represents a RemovedAdmin event raised by the NegRisk contract.
type NegRiskRemovedAdmin struct {
	Admin        common.Address
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_NegRisk *NegRiskFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, admin []common.Address, removedAdmin []common.Address) (*NegRiskRemovedAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _NegRisk.contract.FilterLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskRemovedAdminIterator{contract: _NegRisk.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed admin, address indexed removedAdmin)
func (_NegRisk *NegRiskFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskRemovedAdmin, admin []common.Address, removedAdmin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}
	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _NegRisk.contract.WatchLogs(opts, "RemovedAdmin", adminRule, removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskRemovedAdmin)
				if err := _NegRisk.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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
func (_NegRisk *NegRiskFilterer) ParseRemovedAdmin(log types.Log) (*NegRiskRemovedAdmin, error) {
	event := new(NegRiskRemovedAdmin)
	if err := _NegRisk.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
