package builders

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

const (
	contractAddress = "0xE7819d9745e64c14541732ca07CC3898670b7650"
	chainId         = 42
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

func TestBuildMarketOrderAndSignature(t *testing.T) {
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

	orderHash, err := marketOrderBuilder.BuildMarketOrderHash(marketOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := marketOrderBuilder.BuildSignature(privateKey, orderHash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	marketOrderAndSignature := marketOrderBuilder.BuildMarketOrderAndSignature(marketOrder, signature)
	assert.NotNil(t, marketOrderAndSignature)

	assert.Equal(t, marketOrderAndSignature.OrderType, "market")

	recoveredSignature, err := hex.DecodeString(marketOrderAndSignature.Signature[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredSignature, signature)

	assert.Positive(t, marketOrderAndSignature.Order.Salt)

	recoveredSigner, err := hex.DecodeString(marketOrderAndSignature.Order.Signer[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredSigner, marketOrder.Signer.Bytes())

	recoveredMaker, err := hex.DecodeString(marketOrderAndSignature.Order.Maker[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredMaker, marketOrder.Maker.Bytes())

	recoveredMakerAsset, err := hex.DecodeString(marketOrderAndSignature.Order.MakerAsset[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredMakerAsset, marketOrder.MakerAsset.Bytes())

	assert.Equal(t, marketOrderAndSignature.Order.MakerAmount, marketOrder.MakerAmount.String())
	assert.Equal(t, marketOrderAndSignature.Order.MakerAssetID, int(marketOrder.MakerAssetID.Int64()))

	recoveredTakerAsset, err := hex.DecodeString(marketOrderAndSignature.Order.TakerAsset[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredTakerAsset, marketOrder.TakerAsset.Bytes())

	assert.Equal(t, marketOrderAndSignature.Order.TakerAssetID, int(marketOrder.TakerAssetID.Int64()))
	assert.Equal(t, marketOrderAndSignature.Order.SigType, int(marketOrder.SigType.Int64()))

}

func TestBuildMarketOrderHash(t *testing.T) {
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

	orderHash, err := marketOrderBuilder.BuildMarketOrderHash(marketOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)
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

	orderHash, err := marketOrderBuilder.BuildMarketOrderHash(marketOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := marketOrderBuilder.BuildSignature(privateKey, orderHash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	signature[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow paper
	match, err := marketOrderBuilder.ValidateSignature(&privateKey.PublicKey, orderHash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
