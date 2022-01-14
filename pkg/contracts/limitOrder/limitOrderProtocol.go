// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package limitOrder

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

// LimitOrderMetaData contains all meta data concerning the LimitOrder contract.
var LimitOrderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"ProxyFactoryChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIMIT_ORDER_RFQ_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIMIT_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"makingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"takingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"thresholdAmounts\",\"type\":\"uint256[]\"}],\"name\":\"batchFillOrders\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"oracle1\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"oracle2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"doublePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"func_20xtkDI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"func_40aVqeY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"func_50BkM4K\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC1155\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"func_733NCGU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderMakerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapTakerAmount\",\"type\":\"uint256\"}],\"name\":\"getMakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderMakerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapMakerAmount\",\"type\":\"uint256\"}],\"name\":\"getTakerAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.MarketOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"immutableOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"contractIProxyWalletFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"remainingsRaw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"results\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryAddress\",\"type\":\"address\"}],\"name\":\"setProxyWalletFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"simulateCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inverseAndSpread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"singlePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"timestampBelow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getMakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"getTakerAmount\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.LimitOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateLimitOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sigType\",\"type\":\"uint256\"}],\"internalType\":\"structOrderTypes.MarketOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateMarketOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LimitOrderABI is the input ABI used to generate the binding from.
// Deprecated: Use LimitOrderMetaData.ABI instead.
var LimitOrderABI = LimitOrderMetaData.ABI

// LimitOrder is an auto generated Go binding around an Ethereum contract.
type LimitOrder struct {
	LimitOrderCaller     // Read-only binding to the contract
	LimitOrderTransactor // Write-only binding to the contract
	LimitOrderFilterer   // Log filterer for contract events
}

// LimitOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LimitOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LimitOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LimitOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LimitOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LimitOrderSession struct {
	Contract     *LimitOrder       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LimitOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LimitOrderCallerSession struct {
	Contract *LimitOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LimitOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LimitOrderTransactorSession struct {
	Contract     *LimitOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LimitOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LimitOrderRaw struct {
	Contract *LimitOrder // Generic contract binding to access the raw methods on
}

// LimitOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LimitOrderCallerRaw struct {
	Contract *LimitOrderCaller // Generic read-only contract binding to access the raw methods on
}

// LimitOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LimitOrderTransactorRaw struct {
	Contract *LimitOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLimitOrder creates a new instance of LimitOrder, bound to a specific deployed contract.
func NewLimitOrder(address common.Address, backend bind.ContractBackend) (*LimitOrder, error) {
	contract, err := bindLimitOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LimitOrder{LimitOrderCaller: LimitOrderCaller{contract: contract}, LimitOrderTransactor: LimitOrderTransactor{contract: contract}, LimitOrderFilterer: LimitOrderFilterer{contract: contract}}, nil
}

// NewLimitOrderCaller creates a new read-only instance of LimitOrder, bound to a specific deployed contract.
func NewLimitOrderCaller(address common.Address, caller bind.ContractCaller) (*LimitOrderCaller, error) {
	contract, err := bindLimitOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderCaller{contract: contract}, nil
}

// NewLimitOrderTransactor creates a new write-only instance of LimitOrder, bound to a specific deployed contract.
func NewLimitOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*LimitOrderTransactor, error) {
	contract, err := bindLimitOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LimitOrderTransactor{contract: contract}, nil
}

// NewLimitOrderFilterer creates a new log filterer instance of LimitOrder, bound to a specific deployed contract.
func NewLimitOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*LimitOrderFilterer, error) {
	contract, err := bindLimitOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LimitOrderFilterer{contract: contract}, nil
}

// bindLimitOrder binds a generic wrapper to an already deployed contract.
func bindLimitOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LimitOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrder *LimitOrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrder.Contract.LimitOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrder *LimitOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrder.Contract.LimitOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrder *LimitOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrder.Contract.LimitOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LimitOrder *LimitOrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LimitOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LimitOrder *LimitOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LimitOrder *LimitOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LimitOrder.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LimitOrder *LimitOrderSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LimitOrder.Contract.DOMAINSEPARATOR(&_LimitOrder.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LimitOrder.Contract.DOMAINSEPARATOR(&_LimitOrder.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) LIMITORDERRFQTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "LIMIT_ORDER_RFQ_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderSession) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.LIMITORDERRFQTYPEHASH(&_LimitOrder.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.LIMITORDERRFQTYPEHASH(&_LimitOrder.CallOpts)
}

// LIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x54dd5f74.
//
// Solidity: function LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) LIMITORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "LIMIT_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x54dd5f74.
//
// Solidity: function LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderSession) LIMITORDERTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.LIMITORDERTYPEHASH(&_LimitOrder.CallOpts)
}

// LIMITORDERTYPEHASH is a free data retrieval call binding the contract method 0x54dd5f74.
//
// Solidity: function LIMIT_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) LIMITORDERTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.LIMITORDERTYPEHASH(&_LimitOrder.CallOpts)
}

// MARKETORDERTYPEHASH is a free data retrieval call binding the contract method 0x1c1cad66.
//
// Solidity: function MARKET_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) MARKETORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "MARKET_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETORDERTYPEHASH is a free data retrieval call binding the contract method 0x1c1cad66.
//
// Solidity: function MARKET_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderSession) MARKETORDERTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.MARKETORDERTYPEHASH(&_LimitOrder.CallOpts)
}

// MARKETORDERTYPEHASH is a free data retrieval call binding the contract method 0x1c1cad66.
//
// Solidity: function MARKET_ORDER_TYPEHASH() view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) MARKETORDERTYPEHASH() ([32]byte, error) {
	return _LimitOrder.Contract.MARKETORDERTYPEHASH(&_LimitOrder.CallOpts)
}

// And is a free data retrieval call binding the contract method 0x961d5b1e.
//
// Solidity: function and(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderCaller) And(opts *bind.CallOpts, targets []common.Address, data [][]byte) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "and", targets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0x961d5b1e.
//
// Solidity: function and(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderSession) And(targets []common.Address, data [][]byte) (bool, error) {
	return _LimitOrder.Contract.And(&_LimitOrder.CallOpts, targets, data)
}

// And is a free data retrieval call binding the contract method 0x961d5b1e.
//
// Solidity: function and(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) And(targets []common.Address, data [][]byte) (bool, error) {
	return _LimitOrder.Contract.And(&_LimitOrder.CallOpts, targets, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) ArbitraryStaticCall(opts *bind.CallOpts, target common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "arbitraryStaticCall", target, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_LimitOrder *LimitOrderSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _LimitOrder.Contract.ArbitraryStaticCall(&_LimitOrder.CallOpts, target, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _LimitOrder.Contract.ArbitraryStaticCall(&_LimitOrder.CallOpts, target, data)
}

// CheckPredicate is a free data retrieval call binding the contract method 0xc2cc178f.
//
// Solidity: function checkPredicate((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bool)
func (_LimitOrder *LimitOrderCaller) CheckPredicate(opts *bind.CallOpts, order OrderTypesLimitOrder) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "checkPredicate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPredicate is a free data retrieval call binding the contract method 0xc2cc178f.
//
// Solidity: function checkPredicate((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bool)
func (_LimitOrder *LimitOrderSession) CheckPredicate(order OrderTypesLimitOrder) (bool, error) {
	return _LimitOrder.Contract.CheckPredicate(&_LimitOrder.CallOpts, order)
}

// CheckPredicate is a free data retrieval call binding the contract method 0xc2cc178f.
//
// Solidity: function checkPredicate((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) CheckPredicate(order OrderTypesLimitOrder) (bool, error) {
	return _LimitOrder.Contract.CheckPredicate(&_LimitOrder.CallOpts, order)
}

// DoublePrice is a free data retrieval call binding the contract method 0x36006bf3.
//
// Solidity: function doublePrice(address oracle1, address oracle2, uint256 spread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) DoublePrice(opts *bind.CallOpts, oracle1 common.Address, oracle2 common.Address, spread *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "doublePrice", oracle1, oracle2, spread, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoublePrice is a free data retrieval call binding the contract method 0x36006bf3.
//
// Solidity: function doublePrice(address oracle1, address oracle2, uint256 spread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderSession) DoublePrice(oracle1 common.Address, oracle2 common.Address, spread *big.Int, amount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.DoublePrice(&_LimitOrder.CallOpts, oracle1, oracle2, spread, amount)
}

// DoublePrice is a free data retrieval call binding the contract method 0x36006bf3.
//
// Solidity: function doublePrice(address oracle1, address oracle2, uint256 spread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) DoublePrice(oracle1 common.Address, oracle2 common.Address, spread *big.Int, amount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.DoublePrice(&_LimitOrder.CallOpts, oracle1, oracle2, spread, amount)
}

// Eq is a free data retrieval call binding the contract method 0x32565d61.
//
// Solidity: function eq(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCaller) Eq(opts *bind.CallOpts, value *big.Int, target common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "eq", value, target, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0x32565d61.
//
// Solidity: function eq(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderSession) Eq(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Eq(&_LimitOrder.CallOpts, value, target, data)
}

// Eq is a free data retrieval call binding the contract method 0x32565d61.
//
// Solidity: function eq(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) Eq(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Eq(&_LimitOrder.CallOpts, value, target, data)
}

// GetMakerAmount is a free data retrieval call binding the contract method 0xf4a215c3.
//
// Solidity: function getMakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapTakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderCaller) GetMakerAmount(opts *bind.CallOpts, orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapTakerAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "getMakerAmount", orderMakerAmount, orderTakerAmount, swapTakerAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMakerAmount is a free data retrieval call binding the contract method 0xf4a215c3.
//
// Solidity: function getMakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapTakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderSession) GetMakerAmount(orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapTakerAmount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.GetMakerAmount(&_LimitOrder.CallOpts, orderMakerAmount, orderTakerAmount, swapTakerAmount)
}

// GetMakerAmount is a free data retrieval call binding the contract method 0xf4a215c3.
//
// Solidity: function getMakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapTakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) GetMakerAmount(orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapTakerAmount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.GetMakerAmount(&_LimitOrder.CallOpts, orderMakerAmount, orderTakerAmount, swapTakerAmount)
}

// GetTakerAmount is a free data retrieval call binding the contract method 0x296637bf.
//
// Solidity: function getTakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapMakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderCaller) GetTakerAmount(opts *bind.CallOpts, orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapMakerAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "getTakerAmount", orderMakerAmount, orderTakerAmount, swapMakerAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTakerAmount is a free data retrieval call binding the contract method 0x296637bf.
//
// Solidity: function getTakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapMakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderSession) GetTakerAmount(orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapMakerAmount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.GetTakerAmount(&_LimitOrder.CallOpts, orderMakerAmount, orderTakerAmount, swapMakerAmount)
}

// GetTakerAmount is a free data retrieval call binding the contract method 0x296637bf.
//
// Solidity: function getTakerAmount(uint256 orderMakerAmount, uint256 orderTakerAmount, uint256 swapMakerAmount) pure returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) GetTakerAmount(orderMakerAmount *big.Int, orderTakerAmount *big.Int, swapMakerAmount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.GetTakerAmount(&_LimitOrder.CallOpts, orderMakerAmount, orderTakerAmount, swapMakerAmount)
}

// Gt is a free data retrieval call binding the contract method 0x057702e9.
//
// Solidity: function gt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCaller) Gt(opts *bind.CallOpts, value *big.Int, target common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "gt", value, target, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x057702e9.
//
// Solidity: function gt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderSession) Gt(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Gt(&_LimitOrder.CallOpts, value, target, data)
}

// Gt is a free data retrieval call binding the contract method 0x057702e9.
//
// Solidity: function gt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) Gt(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Gt(&_LimitOrder.CallOpts, value, target, data)
}

// Hash is a free data retrieval call binding the contract method 0x532cd199.
//
// Solidity: function hash((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) Hash(opts *bind.CallOpts, order OrderTypesLimitOrder) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "hash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x532cd199.
//
// Solidity: function hash((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderSession) Hash(order OrderTypesLimitOrder) ([32]byte, error) {
	return _LimitOrder.Contract.Hash(&_LimitOrder.CallOpts, order)
}

// Hash is a free data retrieval call binding the contract method 0x532cd199.
//
// Solidity: function hash((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) Hash(order OrderTypesLimitOrder) ([32]byte, error) {
	return _LimitOrder.Contract.Hash(&_LimitOrder.CallOpts, order)
}

// Hash0 is a free data retrieval call binding the contract method 0xc25fd216.
//
// Solidity: function hash((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderCaller) Hash0(opts *bind.CallOpts, order OrderTypesMarketOrder) ([32]byte, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "hash0", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xc25fd216.
//
// Solidity: function hash((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderSession) Hash0(order OrderTypesMarketOrder) ([32]byte, error) {
	return _LimitOrder.Contract.Hash0(&_LimitOrder.CallOpts, order)
}

// Hash0 is a free data retrieval call binding the contract method 0xc25fd216.
//
// Solidity: function hash((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order) view returns(bytes32)
func (_LimitOrder *LimitOrderCallerSession) Hash0(order OrderTypesMarketOrder) ([32]byte, error) {
	return _LimitOrder.Contract.Hash0(&_LimitOrder.CallOpts, order)
}

// ImmutableOwner is a free data retrieval call binding the contract method 0x8ec73568.
//
// Solidity: function immutableOwner() view returns(address)
func (_LimitOrder *LimitOrderCaller) ImmutableOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "immutableOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImmutableOwner is a free data retrieval call binding the contract method 0x8ec73568.
//
// Solidity: function immutableOwner() view returns(address)
func (_LimitOrder *LimitOrderSession) ImmutableOwner() (common.Address, error) {
	return _LimitOrder.Contract.ImmutableOwner(&_LimitOrder.CallOpts)
}

// ImmutableOwner is a free data retrieval call binding the contract method 0x8ec73568.
//
// Solidity: function immutableOwner() view returns(address)
func (_LimitOrder *LimitOrderCallerSession) ImmutableOwner() (common.Address, error) {
	return _LimitOrder.Contract.ImmutableOwner(&_LimitOrder.CallOpts)
}

// Lt is a free data retrieval call binding the contract method 0x871919d5.
//
// Solidity: function lt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCaller) Lt(opts *bind.CallOpts, value *big.Int, target common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "lt", value, target, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0x871919d5.
//
// Solidity: function lt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderSession) Lt(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Lt(&_LimitOrder.CallOpts, value, target, data)
}

// Lt is a free data retrieval call binding the contract method 0x871919d5.
//
// Solidity: function lt(uint256 value, address target, bytes data) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) Lt(value *big.Int, target common.Address, data []byte) (bool, error) {
	return _LimitOrder.Contract.Lt(&_LimitOrder.CallOpts, value, target, data)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_LimitOrder *LimitOrderSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _LimitOrder.Contract.Nonce(&_LimitOrder.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _LimitOrder.Contract.Nonce(&_LimitOrder.CallOpts, arg0)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_LimitOrder *LimitOrderCaller) NonceEquals(opts *bind.CallOpts, makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "nonceEquals", makerAddress, makerNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_LimitOrder *LimitOrderSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _LimitOrder.Contract.NonceEquals(&_LimitOrder.CallOpts, makerAddress, makerNonce)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _LimitOrder.Contract.NonceEquals(&_LimitOrder.CallOpts, makerAddress, makerNonce)
}

// Or is a free data retrieval call binding the contract method 0xe6133301.
//
// Solidity: function or(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderCaller) Or(opts *bind.CallOpts, targets []common.Address, data [][]byte) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "or", targets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0xe6133301.
//
// Solidity: function or(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderSession) Or(targets []common.Address, data [][]byte) (bool, error) {
	return _LimitOrder.Contract.Or(&_LimitOrder.CallOpts, targets, data)
}

// Or is a free data retrieval call binding the contract method 0xe6133301.
//
// Solidity: function or(address[] targets, bytes[] data) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) Or(targets []common.Address, data [][]byte) (bool, error) {
	return _LimitOrder.Contract.Or(&_LimitOrder.CallOpts, targets, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrder *LimitOrderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrder *LimitOrderSession) Owner() (common.Address, error) {
	return _LimitOrder.Contract.Owner(&_LimitOrder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LimitOrder *LimitOrderCallerSession) Owner() (common.Address, error) {
	return _LimitOrder.Contract.Owner(&_LimitOrder.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_LimitOrder *LimitOrderCaller) ProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "proxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_LimitOrder *LimitOrderSession) ProxyFactory() (common.Address, error) {
	return _LimitOrder.Contract.ProxyFactory(&_LimitOrder.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_LimitOrder *LimitOrderCallerSession) ProxyFactory() (common.Address, error) {
	return _LimitOrder.Contract.ProxyFactory(&_LimitOrder.CallOpts)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) Remaining(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "remaining", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _LimitOrder.Contract.Remaining(&_LimitOrder.CallOpts, orderHash)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _LimitOrder.Contract.Remaining(&_LimitOrder.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) RemainingRaw(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "remainingRaw", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _LimitOrder.Contract.RemainingRaw(&_LimitOrder.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _LimitOrder.Contract.RemainingRaw(&_LimitOrder.CallOpts, orderHash)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[] results)
func (_LimitOrder *LimitOrderCaller) RemainingsRaw(opts *bind.CallOpts, orderHashes [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "remainingsRaw", orderHashes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[] results)
func (_LimitOrder *LimitOrderSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _LimitOrder.Contract.RemainingsRaw(&_LimitOrder.CallOpts, orderHashes)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[] results)
func (_LimitOrder *LimitOrderCallerSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _LimitOrder.Contract.RemainingsRaw(&_LimitOrder.CallOpts, orderHashes)
}

// SinglePrice is a free data retrieval call binding the contract method 0xc05435f1.
//
// Solidity: function singlePrice(address oracle, uint256 inverseAndSpread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderCaller) SinglePrice(opts *bind.CallOpts, oracle common.Address, inverseAndSpread *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "singlePrice", oracle, inverseAndSpread, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SinglePrice is a free data retrieval call binding the contract method 0xc05435f1.
//
// Solidity: function singlePrice(address oracle, uint256 inverseAndSpread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderSession) SinglePrice(oracle common.Address, inverseAndSpread *big.Int, amount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.SinglePrice(&_LimitOrder.CallOpts, oracle, inverseAndSpread, amount)
}

// SinglePrice is a free data retrieval call binding the contract method 0xc05435f1.
//
// Solidity: function singlePrice(address oracle, uint256 inverseAndSpread, uint256 amount) view returns(uint256)
func (_LimitOrder *LimitOrderCallerSession) SinglePrice(oracle common.Address, inverseAndSpread *big.Int, amount *big.Int) (*big.Int, error) {
	return _LimitOrder.Contract.SinglePrice(&_LimitOrder.CallOpts, oracle, inverseAndSpread, amount)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_LimitOrder *LimitOrderCaller) TimestampBelow(opts *bind.CallOpts, time *big.Int) (bool, error) {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "timestampBelow", time)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_LimitOrder *LimitOrderSession) TimestampBelow(time *big.Int) (bool, error) {
	return _LimitOrder.Contract.TimestampBelow(&_LimitOrder.CallOpts, time)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_LimitOrder *LimitOrderCallerSession) TimestampBelow(time *big.Int) (bool, error) {
	return _LimitOrder.Contract.TimestampBelow(&_LimitOrder.CallOpts, time)
}

// ValidateLimitOrder is a free data retrieval call binding the contract method 0x8c7a8aac.
//
// Solidity: function validateLimitOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderCaller) ValidateLimitOrder(opts *bind.CallOpts, order OrderTypesLimitOrder, signature []byte) error {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "validateLimitOrder", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateLimitOrder is a free data retrieval call binding the contract method 0x8c7a8aac.
//
// Solidity: function validateLimitOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderSession) ValidateLimitOrder(order OrderTypesLimitOrder, signature []byte) error {
	return _LimitOrder.Contract.ValidateLimitOrder(&_LimitOrder.CallOpts, order, signature)
}

// ValidateLimitOrder is a free data retrieval call binding the contract method 0x8c7a8aac.
//
// Solidity: function validateLimitOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderCallerSession) ValidateLimitOrder(order OrderTypesLimitOrder, signature []byte) error {
	return _LimitOrder.Contract.ValidateLimitOrder(&_LimitOrder.CallOpts, order, signature)
}

// ValidateMarketOrder is a free data retrieval call binding the contract method 0xba329a25.
//
// Solidity: function validateMarketOrder((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderCaller) ValidateMarketOrder(opts *bind.CallOpts, order OrderTypesMarketOrder, signature []byte) error {
	var out []interface{}
	err := _LimitOrder.contract.Call(opts, &out, "validateMarketOrder", order, signature)

	if err != nil {
		return err
	}

	return err

}

// ValidateMarketOrder is a free data retrieval call binding the contract method 0xba329a25.
//
// Solidity: function validateMarketOrder((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderSession) ValidateMarketOrder(order OrderTypesMarketOrder, signature []byte) error {
	return _LimitOrder.Contract.ValidateMarketOrder(&_LimitOrder.CallOpts, order, signature)
}

// ValidateMarketOrder is a free data retrieval call binding the contract method 0xba329a25.
//
// Solidity: function validateMarketOrder((uint256,address,address,address,uint256,uint256,address,uint256,uint256) order, bytes signature) view returns()
func (_LimitOrder *LimitOrderCallerSession) ValidateMarketOrder(order OrderTypesMarketOrder, signature []byte) error {
	return _LimitOrder.Contract.ValidateMarketOrder(&_LimitOrder.CallOpts, order, signature)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_LimitOrder *LimitOrderTransactor) AdvanceNonce(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "advanceNonce", amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_LimitOrder *LimitOrderSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _LimitOrder.Contract.AdvanceNonce(&_LimitOrder.TransactOpts, amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_LimitOrder *LimitOrderTransactorSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _LimitOrder.Contract.AdvanceNonce(&_LimitOrder.TransactOpts, amount)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x8bbd66d4.
//
// Solidity: function batchFillOrders((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[] orders, bytes[] signatures, uint256[] makingAmounts, uint256[] takingAmounts, uint256[] thresholdAmounts) returns(uint256[], uint256[])
func (_LimitOrder *LimitOrderTransactor) BatchFillOrders(opts *bind.TransactOpts, orders []OrderTypesLimitOrder, signatures [][]byte, makingAmounts []*big.Int, takingAmounts []*big.Int, thresholdAmounts []*big.Int) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "batchFillOrders", orders, signatures, makingAmounts, takingAmounts, thresholdAmounts)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x8bbd66d4.
//
// Solidity: function batchFillOrders((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[] orders, bytes[] signatures, uint256[] makingAmounts, uint256[] takingAmounts, uint256[] thresholdAmounts) returns(uint256[], uint256[])
func (_LimitOrder *LimitOrderSession) BatchFillOrders(orders []OrderTypesLimitOrder, signatures [][]byte, makingAmounts []*big.Int, takingAmounts []*big.Int, thresholdAmounts []*big.Int) (*types.Transaction, error) {
	return _LimitOrder.Contract.BatchFillOrders(&_LimitOrder.TransactOpts, orders, signatures, makingAmounts, takingAmounts, thresholdAmounts)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0x8bbd66d4.
//
// Solidity: function batchFillOrders((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256)[] orders, bytes[] signatures, uint256[] makingAmounts, uint256[] takingAmounts, uint256[] thresholdAmounts) returns(uint256[], uint256[])
func (_LimitOrder *LimitOrderTransactorSession) BatchFillOrders(orders []OrderTypesLimitOrder, signatures [][]byte, makingAmounts []*big.Int, takingAmounts []*big.Int, thresholdAmounts []*big.Int) (*types.Transaction, error) {
	return _LimitOrder.Contract.BatchFillOrders(&_LimitOrder.TransactOpts, orders, signatures, makingAmounts, takingAmounts, thresholdAmounts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x05d0d562.
//
// Solidity: function cancelOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) returns()
func (_LimitOrder *LimitOrderTransactor) CancelOrder(opts *bind.TransactOpts, order OrderTypesLimitOrder) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x05d0d562.
//
// Solidity: function cancelOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) returns()
func (_LimitOrder *LimitOrderSession) CancelOrder(order OrderTypesLimitOrder) (*types.Transaction, error) {
	return _LimitOrder.Contract.CancelOrder(&_LimitOrder.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x05d0d562.
//
// Solidity: function cancelOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order) returns()
func (_LimitOrder *LimitOrderTransactorSession) CancelOrder(order OrderTypesLimitOrder) (*types.Transaction, error) {
	return _LimitOrder.Contract.CancelOrder(&_LimitOrder.TransactOpts, order)
}

// FillOrder is a paid mutator transaction binding the contract method 0x342b3879.
//
// Solidity: function fillOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactor) FillOrder(opts *bind.TransactOpts, order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "fillOrder", order, signature, makingAmount, takingAmount, thresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x342b3879.
//
// Solidity: function fillOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount) returns(uint256, uint256)
func (_LimitOrder *LimitOrderSession) FillOrder(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrder(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x342b3879.
//
// Solidity: function fillOrder((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactorSession) FillOrder(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrder(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0x8aa76226.
//
// Solidity: function fillOrderTo((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactor) FillOrderTo(opts *bind.TransactOpts, order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "fillOrderTo", order, signature, makingAmount, takingAmount, thresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0x8aa76226.
//
// Solidity: function fillOrderTo((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target) returns(uint256, uint256)
func (_LimitOrder *LimitOrderSession) FillOrderTo(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrderTo(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0x8aa76226.
//
// Solidity: function fillOrderTo((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactorSession) FillOrderTo(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrderTo(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount, target)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xfe05c0e8.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target, bytes permit) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactor) FillOrderToWithPermit(opts *bind.TransactOpts, order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "fillOrderToWithPermit", order, signature, makingAmount, takingAmount, thresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xfe05c0e8.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target, bytes permit) returns(uint256, uint256)
func (_LimitOrder *LimitOrderSession) FillOrderToWithPermit(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrderToWithPermit(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xfe05c0e8.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,bytes,bytes,bytes,bytes,bytes,bytes,bytes,address,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, uint256 thresholdAmount, address target, bytes permit) returns(uint256, uint256)
func (_LimitOrder *LimitOrderTransactorSession) FillOrderToWithPermit(order OrderTypesLimitOrder, signature []byte, makingAmount *big.Int, takingAmount *big.Int, thresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.FillOrderToWithPermit(&_LimitOrder.TransactOpts, order, signature, makingAmount, takingAmount, thresholdAmount, target, permit)
}

// Func20xtkDI is a paid mutator transaction binding the contract method 0x23b872e0.
//
// Solidity: function func_20xtkDI(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderTransactor) Func20xtkDI(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "func_20xtkDI", from, to, tokenId, token)
}

// Func20xtkDI is a paid mutator transaction binding the contract method 0x23b872e0.
//
// Solidity: function func_20xtkDI(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderSession) Func20xtkDI(from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func20xtkDI(&_LimitOrder.TransactOpts, from, to, tokenId, token)
}

// Func20xtkDI is a paid mutator transaction binding the contract method 0x23b872e0.
//
// Solidity: function func_20xtkDI(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderTransactorSession) Func20xtkDI(from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func20xtkDI(&_LimitOrder.TransactOpts, from, to, tokenId, token)
}

// Func40aVqeY is a paid mutator transaction binding the contract method 0x23b872df.
//
// Solidity: function func_40aVqeY(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderTransactor) Func40aVqeY(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "func_40aVqeY", from, to, tokenId, token)
}

// Func40aVqeY is a paid mutator transaction binding the contract method 0x23b872df.
//
// Solidity: function func_40aVqeY(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderSession) Func40aVqeY(from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func40aVqeY(&_LimitOrder.TransactOpts, from, to, tokenId, token)
}

// Func40aVqeY is a paid mutator transaction binding the contract method 0x23b872df.
//
// Solidity: function func_40aVqeY(address from, address to, uint256 tokenId, address token) returns()
func (_LimitOrder *LimitOrderTransactorSession) Func40aVqeY(from common.Address, to common.Address, tokenId *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func40aVqeY(&_LimitOrder.TransactOpts, from, to, tokenId, token)
}

// Func50BkM4K is a paid mutator transaction binding the contract method 0x23b872de.
//
// Solidity: function func_50BkM4K(address from, address to, uint256 amount, address token) returns()
func (_LimitOrder *LimitOrderTransactor) Func50BkM4K(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "func_50BkM4K", from, to, amount, token)
}

// Func50BkM4K is a paid mutator transaction binding the contract method 0x23b872de.
//
// Solidity: function func_50BkM4K(address from, address to, uint256 amount, address token) returns()
func (_LimitOrder *LimitOrderSession) Func50BkM4K(from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func50BkM4K(&_LimitOrder.TransactOpts, from, to, amount, token)
}

// Func50BkM4K is a paid mutator transaction binding the contract method 0x23b872de.
//
// Solidity: function func_50BkM4K(address from, address to, uint256 amount, address token) returns()
func (_LimitOrder *LimitOrderTransactorSession) Func50BkM4K(from common.Address, to common.Address, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func50BkM4K(&_LimitOrder.TransactOpts, from, to, amount, token)
}

// Func733NCGU is a paid mutator transaction binding the contract method 0x23b872e1.
//
// Solidity: function func_733NCGU(address from, address to, uint256 amount, address token, uint256 tokenId, bytes data) returns()
func (_LimitOrder *LimitOrderTransactor) Func733NCGU(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, token common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "func_733NCGU", from, to, amount, token, tokenId, data)
}

// Func733NCGU is a paid mutator transaction binding the contract method 0x23b872e1.
//
// Solidity: function func_733NCGU(address from, address to, uint256 amount, address token, uint256 tokenId, bytes data) returns()
func (_LimitOrder *LimitOrderSession) Func733NCGU(from common.Address, to common.Address, amount *big.Int, token common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func733NCGU(&_LimitOrder.TransactOpts, from, to, amount, token, tokenId, data)
}

// Func733NCGU is a paid mutator transaction binding the contract method 0x23b872e1.
//
// Solidity: function func_733NCGU(address from, address to, uint256 amount, address token, uint256 tokenId, bytes data) returns()
func (_LimitOrder *LimitOrderTransactorSession) Func733NCGU(from common.Address, to common.Address, amount *big.Int, token common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.Func733NCGU(&_LimitOrder.TransactOpts, from, to, amount, token, tokenId, data)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_LimitOrder *LimitOrderTransactor) IncreaseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "increaseNonce")
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_LimitOrder *LimitOrderSession) IncreaseNonce() (*types.Transaction, error) {
	return _LimitOrder.Contract.IncreaseNonce(&_LimitOrder.TransactOpts)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_LimitOrder *LimitOrderTransactorSession) IncreaseNonce() (*types.Transaction, error) {
	return _LimitOrder.Contract.IncreaseNonce(&_LimitOrder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrder *LimitOrderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrder *LimitOrderSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrder.Contract.RenounceOwnership(&_LimitOrder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LimitOrder *LimitOrderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LimitOrder.Contract.RenounceOwnership(&_LimitOrder.TransactOpts)
}

// SetProxyWalletFactory is a paid mutator transaction binding the contract method 0xa50453bf.
//
// Solidity: function setProxyWalletFactory(address factoryAddress) returns()
func (_LimitOrder *LimitOrderTransactor) SetProxyWalletFactory(opts *bind.TransactOpts, factoryAddress common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "setProxyWalletFactory", factoryAddress)
}

// SetProxyWalletFactory is a paid mutator transaction binding the contract method 0xa50453bf.
//
// Solidity: function setProxyWalletFactory(address factoryAddress) returns()
func (_LimitOrder *LimitOrderSession) SetProxyWalletFactory(factoryAddress common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.SetProxyWalletFactory(&_LimitOrder.TransactOpts, factoryAddress)
}

// SetProxyWalletFactory is a paid mutator transaction binding the contract method 0xa50453bf.
//
// Solidity: function setProxyWalletFactory(address factoryAddress) returns()
func (_LimitOrder *LimitOrderTransactorSession) SetProxyWalletFactory(factoryAddress common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.SetProxyWalletFactory(&_LimitOrder.TransactOpts, factoryAddress)
}

// SimulateCalls is a paid mutator transaction binding the contract method 0x7f29a59d.
//
// Solidity: function simulateCalls(address[] targets, bytes[] data) returns()
func (_LimitOrder *LimitOrderTransactor) SimulateCalls(opts *bind.TransactOpts, targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "simulateCalls", targets, data)
}

// SimulateCalls is a paid mutator transaction binding the contract method 0x7f29a59d.
//
// Solidity: function simulateCalls(address[] targets, bytes[] data) returns()
func (_LimitOrder *LimitOrderSession) SimulateCalls(targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.SimulateCalls(&_LimitOrder.TransactOpts, targets, data)
}

// SimulateCalls is a paid mutator transaction binding the contract method 0x7f29a59d.
//
// Solidity: function simulateCalls(address[] targets, bytes[] data) returns()
func (_LimitOrder *LimitOrderTransactorSession) SimulateCalls(targets []common.Address, data [][]byte) (*types.Transaction, error) {
	return _LimitOrder.Contract.SimulateCalls(&_LimitOrder.TransactOpts, targets, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrder *LimitOrderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrder *LimitOrderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.TransferOwnership(&_LimitOrder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LimitOrder *LimitOrderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LimitOrder.Contract.TransferOwnership(&_LimitOrder.TransactOpts, newOwner)
}

// LimitOrderNonceIncreasedIterator is returned from FilterNonceIncreased and is used to iterate over the raw logs and unpacked data for NonceIncreased events raised by the LimitOrder contract.
type LimitOrderNonceIncreasedIterator struct {
	Event *LimitOrderNonceIncreased // Event containing the contract specifics and raw log

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
func (it *LimitOrderNonceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderNonceIncreased)
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
		it.Event = new(LimitOrderNonceIncreased)
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
func (it *LimitOrderNonceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderNonceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderNonceIncreased represents a NonceIncreased event raised by the LimitOrder contract.
type LimitOrderNonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncreased is a free log retrieval operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_LimitOrder *LimitOrderFilterer) FilterNonceIncreased(opts *bind.FilterOpts, maker []common.Address) (*LimitOrderNonceIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _LimitOrder.contract.FilterLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderNonceIncreasedIterator{contract: _LimitOrder.contract, event: "NonceIncreased", logs: logs, sub: sub}, nil
}

// WatchNonceIncreased is a free log subscription operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_LimitOrder *LimitOrderFilterer) WatchNonceIncreased(opts *bind.WatchOpts, sink chan<- *LimitOrderNonceIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _LimitOrder.contract.WatchLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderNonceIncreased)
				if err := _LimitOrder.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
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

// ParseNonceIncreased is a log parse operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_LimitOrder *LimitOrderFilterer) ParseNonceIncreased(log types.Log) (*LimitOrderNonceIncreased, error) {
	event := new(LimitOrderNonceIncreased)
	if err := _LimitOrder.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the LimitOrder contract.
type LimitOrderOrderFilledIterator struct {
	Event *LimitOrderOrderFilled // Event containing the contract specifics and raw log

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
func (it *LimitOrderOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderOrderFilled)
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
		it.Event = new(LimitOrderOrderFilled)
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
func (it *LimitOrderOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderOrderFilled represents a OrderFilled event raised by the LimitOrder contract.
type LimitOrderOrderFilled struct {
	Maker     common.Address
	OrderHash [32]byte
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_LimitOrder *LimitOrderFilterer) FilterOrderFilled(opts *bind.FilterOpts, maker []common.Address) (*LimitOrderOrderFilledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _LimitOrder.contract.FilterLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderOrderFilledIterator{contract: _LimitOrder.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_LimitOrder *LimitOrderFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *LimitOrderOrderFilled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _LimitOrder.contract.WatchLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderOrderFilled)
				if err := _LimitOrder.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_LimitOrder *LimitOrderFilterer) ParseOrderFilled(log types.Log) (*LimitOrderOrderFilled, error) {
	event := new(LimitOrderOrderFilled)
	if err := _LimitOrder.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LimitOrder contract.
type LimitOrderOwnershipTransferredIterator struct {
	Event *LimitOrderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LimitOrderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderOwnershipTransferred)
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
		it.Event = new(LimitOrderOwnershipTransferred)
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
func (it *LimitOrderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderOwnershipTransferred represents a OwnershipTransferred event raised by the LimitOrder contract.
type LimitOrderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrder *LimitOrderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LimitOrderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LimitOrderOwnershipTransferredIterator{contract: _LimitOrder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LimitOrder *LimitOrderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LimitOrderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LimitOrder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderOwnershipTransferred)
				if err := _LimitOrder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LimitOrder *LimitOrderFilterer) ParseOwnershipTransferred(log types.Log) (*LimitOrderOwnershipTransferred, error) {
	event := new(LimitOrderOwnershipTransferred)
	if err := _LimitOrder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LimitOrderProxyFactoryChangedIterator is returned from FilterProxyFactoryChanged and is used to iterate over the raw logs and unpacked data for ProxyFactoryChanged events raised by the LimitOrder contract.
type LimitOrderProxyFactoryChangedIterator struct {
	Event *LimitOrderProxyFactoryChanged // Event containing the contract specifics and raw log

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
func (it *LimitOrderProxyFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LimitOrderProxyFactoryChanged)
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
		it.Event = new(LimitOrderProxyFactoryChanged)
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
func (it *LimitOrderProxyFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LimitOrderProxyFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LimitOrderProxyFactoryChanged represents a ProxyFactoryChanged event raised by the LimitOrder contract.
type LimitOrderProxyFactoryChanged struct {
	OldFactory common.Address
	NewFactory common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyFactoryChanged is a free log retrieval operation binding the contract event 0x71afcbb97d2433336627c4de6ebc02c1bc2b1cfbe837a75ba8644a04a344b9a9.
//
// Solidity: event ProxyFactoryChanged(address oldFactory, address newFactory)
func (_LimitOrder *LimitOrderFilterer) FilterProxyFactoryChanged(opts *bind.FilterOpts) (*LimitOrderProxyFactoryChangedIterator, error) {

	logs, sub, err := _LimitOrder.contract.FilterLogs(opts, "ProxyFactoryChanged")
	if err != nil {
		return nil, err
	}
	return &LimitOrderProxyFactoryChangedIterator{contract: _LimitOrder.contract, event: "ProxyFactoryChanged", logs: logs, sub: sub}, nil
}

// WatchProxyFactoryChanged is a free log subscription operation binding the contract event 0x71afcbb97d2433336627c4de6ebc02c1bc2b1cfbe837a75ba8644a04a344b9a9.
//
// Solidity: event ProxyFactoryChanged(address oldFactory, address newFactory)
func (_LimitOrder *LimitOrderFilterer) WatchProxyFactoryChanged(opts *bind.WatchOpts, sink chan<- *LimitOrderProxyFactoryChanged) (event.Subscription, error) {

	logs, sub, err := _LimitOrder.contract.WatchLogs(opts, "ProxyFactoryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LimitOrderProxyFactoryChanged)
				if err := _LimitOrder.contract.UnpackLog(event, "ProxyFactoryChanged", log); err != nil {
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

// ParseProxyFactoryChanged is a log parse operation binding the contract event 0x71afcbb97d2433336627c4de6ebc02c1bc2b1cfbe837a75ba8644a04a344b9a9.
//
// Solidity: event ProxyFactoryChanged(address oldFactory, address newFactory)
func (_LimitOrder *LimitOrderFilterer) ParseProxyFactoryChanged(log types.Log) (*LimitOrderProxyFactoryChanged, error) {
	event := new(LimitOrderProxyFactoryChanged)
	if err := _LimitOrder.contract.UnpackLog(event, "ProxyFactoryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}