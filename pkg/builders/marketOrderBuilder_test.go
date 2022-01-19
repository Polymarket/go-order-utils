package builders

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

const (
	contractAddress = "0xE7819d9745e64c14541732ca07CC3898670b7650"
	privateKey      = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	chainId         = 5
)

func getMarketOrderBuilderImpl(t *testing.T) MarketOrderBuilder {
	marketOrderBuilder := NewMarketOrderBuilderImpl(
		common.HexToAddress(contractAddress),
		chainId,
		nil,
	)
	assert.NotNil(t, marketOrderBuilder)
	assert.Equal(t, marketOrderBuilder.chainId, math.NewHexOrDecimal256(int64(chainId)))
	assert.Equal(t, marketOrderBuilder.contractAddress.String(), contractAddress)
	assert.NotNil(t, marketOrderBuilder.saltGenerator)
	return marketOrderBuilder
}

func TestBuildMarketOrder(t *testing.T) {
	marketOrderBuilder := getMarketOrderBuilderImpl(t)

	marketOrder := marketOrderBuilder.BuildMarketOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		big.NewInt(int64(100)),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		model.EOA,
	)
	assert.NotNil(t, marketOrder)
	assert.NotNil(t, marketOrder.Salt)
	assert.Equal(t, marketOrder.Signer.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654").String())
	assert.Equal(t, marketOrder.Maker.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653").String())
	assert.Equal(t, marketOrder.MakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651").String())
	assert.Equal(t, marketOrder.MakerAmount.Int64(), int64(100))
	assert.Equal(t, marketOrder.MakerAssetID.Int64(), int64(2))
	assert.Equal(t, marketOrder.TakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652").String())
	assert.Equal(t, marketOrder.TakerAssetID.Int64(), int64(1))
	assert.Equal(t, marketOrder.SigType.Int64(), int64(0))
}

func TestBuildMarketOrderTypedData(t *testing.T) {
	marketOrderBuilder := getMarketOrderBuilderImpl(t)

	marketOrder := marketOrderBuilder.BuildMarketOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		big.NewInt(int64(100)),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		model.EOA,
	)
	assert.NotNil(t, marketOrder)

	orderTypedData := marketOrderBuilder.BuildMarketOrderTypedData(marketOrder)
	assert.NotNil(t, orderTypedData)
	assert.Equal(t, orderTypedData.PrimaryType, "MarketOrder")
	assert.Equal(t, orderTypedData.Types["MarketOrder"], MARKET_ORDER_STRUCTURE)
	assert.Equal(t, orderTypedData.Types["EIP712Domain"], EIP712_DOMAIN)
	assert.Equal(t, orderTypedData.Domain.Name, PROTOCOL_NAME)
	assert.Equal(t, orderTypedData.Domain.Version, PROTOCOL_VERSION)
	assert.Equal(t, orderTypedData.Domain.ChainId, math.NewHexOrDecimal256(int64(chainId)))
	assert.Equal(t, orderTypedData.Domain.VerifyingContract, contractAddress)
	assert.NotNil(t, orderTypedData.Domain.Salt)

	assert.NotNil(t, orderTypedData.Message)
	assert.Equal(t, orderTypedData.Message, signer.TypedDataMessage{
		"salt":         marketOrder.Salt.String(),
		"signer":       marketOrder.Signer.String(),
		"maker":        marketOrder.Maker.String(),
		"makerAsset":   marketOrder.MakerAsset.String(),
		"makerAmount":  marketOrder.MakerAmount.String(),
		"makerAssetID": marketOrder.MakerAssetID.String(),
		"takerAsset":   marketOrder.TakerAsset.String(),
		"takerAssetID": marketOrder.TakerAssetID.String(),
		"sigType":      marketOrder.SigType.String(),
	})
}

func TestMarketOrderBuilderAndSign(t *testing.T) {
	marketOrderBuilder := getMarketOrderBuilderImpl(t)

	marketOrder := marketOrderBuilder.BuildMarketOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		big.NewInt(int64(100)),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		model.EOA,
	)
	assert.NotNil(t, marketOrder)

	orderTypedData := marketOrderBuilder.BuildMarketOrderTypedData(marketOrder)
	assert.NotNil(t, orderTypedData)

	hash, err := marketOrderBuilder.BuildHash(orderTypedData)
	assert.NotNil(t, hash)
	assert.Nil(t, err)

	privateKey, err := crypto.HexToECDSA(privateKey)
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := marketOrderBuilder.BuildSignature(privateKey, hash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	match, err := marketOrderBuilder.ValidateSignature(&privateKey.PublicKey, hash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
