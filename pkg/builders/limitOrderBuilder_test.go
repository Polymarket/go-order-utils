package builders

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/facades"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/sign"
	"github.com/stretchr/testify/assert"
)

func getLimitOrderBuilderImpl(t *testing.T) LimitOrderBuilder {
	limitOrderProtocolFacade, err := facades.NewLimitOrderProtocolFacadeImpl()
	assert.NotNil(t, limitOrderProtocolFacade)
	assert.Nil(t, err)

	limitOrderPredicateBuilder := NewLimitOrderPredicateBuilderImpl(limitOrderProtocolFacade)
	assert.NotNil(t, limitOrderPredicateBuilder)

	erc20Facade, err := facades.NewERC20FacadeImpl()
	assert.NotNil(t, erc20Facade)
	assert.Nil(t, err)

	erc1155Facade, err := facades.NewERC1155FacadeImpl()
	assert.NotNil(t, erc1155Facade)
	assert.Nil(t, err)

	limitOrderBuilder, err := NewLimitOrderBuilderImpl(
		common.HexToAddress(contractAddress),
		chainId,
		nil,
	)
	assert.NotNil(t, limitOrderBuilder)
	assert.Nil(t, err)
	assert.Equal(t, limitOrderBuilder.chainId, math.NewHexOrDecimal256(int64(chainId)))
	assert.Equal(t, limitOrderBuilder.contractAddress.String(), contractAddress)
	assert.NotNil(t, limitOrderBuilder.saltGenerator)
	assert.Equal(t, limitOrderBuilder.limitOrderProtocolFacade, limitOrderProtocolFacade)
	assert.Equal(t, limitOrderBuilder.limitOrderPredicateBuilder, limitOrderPredicateBuilder)
	assert.Equal(t, limitOrderBuilder.erc20Facade, erc20Facade)
	assert.Equal(t, limitOrderBuilder.erc1155Facade, erc1155Facade)

	return limitOrderBuilder
}

func TestBuildLimitOrder(t *testing.T) {
	limitOrderBuilder := getLimitOrderBuilderImpl(t)

	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7655"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7656"),
		[]byte("1"),
		[]byte("1"),
		[]byte("3"),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		big.NewInt(int64(3)),
		big.NewInt(int64(4)),
		big.NewInt(int64(5)),
		big.NewInt(int64(6)),
		model.POLY_PROXY,
	)
	assert.NotNil(t, limitOrder)
	assert.Nil(t, err)

	assert.NotNil(t, limitOrder.Salt)
	assert.Equal(t, limitOrder.MakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651").String())
	assert.Equal(t, limitOrder.TakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651").String())
	assert.NotNil(t, limitOrder.MakerAssetData)
	assert.NotNil(t, limitOrder.TakerAssetData)
	assert.NotNil(t, limitOrder.GetMakerAmount)
	assert.NotNil(t, limitOrder.GetTakerAmount)
	assert.NotNil(t, limitOrder.Predicate)
	assert.NotNil(t, limitOrder.Permit)
	assert.NotNil(t, limitOrder.Interaction)
	assert.Equal(t, limitOrder.Signer.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7656").String())
	assert.Equal(t, limitOrder.SigType.Int64(), int64(model.POLY_PROXY))

	limitOrder, err = limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7655"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7656"),
		[]byte("1"),
		[]byte("1"),
		[]byte("3"),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		big.NewInt(int64(0)),
		big.NewInt(int64(0)),
		big.NewInt(int64(5)),
		big.NewInt(int64(6)),
		model.POLY_PROXY,
	)
	assert.NotNil(t, limitOrder)
	assert.Nil(t, err)

	assert.NotNil(t, limitOrder.Salt)
	assert.Equal(t, limitOrder.MakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652").String())
	assert.Equal(t, limitOrder.TakerAsset.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654").String())
	assert.NotNil(t, limitOrder.MakerAssetData)
	assert.NotNil(t, limitOrder.TakerAssetData)
	assert.NotNil(t, limitOrder.GetMakerAmount)
	assert.NotNil(t, limitOrder.GetTakerAmount)
	assert.NotNil(t, limitOrder.Predicate)
	assert.NotNil(t, limitOrder.Permit)
	assert.NotNil(t, limitOrder.Interaction)
	assert.Equal(t, limitOrder.Signer.String(), common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7656").String())
	assert.Equal(t, limitOrder.SigType.Int64(), int64(model.POLY_PROXY))
}

func TestBuildLimitOrderHash(t *testing.T) {
	limitOrderBuilder := getLimitOrderBuilderImpl(t)

	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7655"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7656"),
		[]byte("1"),
		[]byte("1"),
		[]byte("3"),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		big.NewInt(int64(3)),
		big.NewInt(int64(4)),
		big.NewInt(int64(5)),
		big.NewInt(int64(6)),
		model.POLY_PROXY,
	)
	assert.NotNil(t, limitOrder)
	assert.Nil(t, err)

	orderHash, err := limitOrderBuilder.BuildLimitOrderHash(limitOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)
}

func TestBuildLimitOrderAndSignature(t *testing.T) {
	limitOrderBuilder := getLimitOrderBuilderImpl(t)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signer, err := sign.GetPublicAddress(privateKey)
	assert.NotNil(t, signer)
	assert.Nil(t, err)

	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7655"),
		signer,
		[]byte("1"),
		[]byte("1"),
		[]byte("3"),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		big.NewInt(int64(3)),
		big.NewInt(int64(4)),
		big.NewInt(int64(5)),
		big.NewInt(int64(6)),
		model.POLY_PROXY,
	)
	assert.NotNil(t, limitOrder)
	assert.Nil(t, err)

	orderHash, err := limitOrderBuilder.BuildLimitOrderHash(limitOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)

	signature, err := limitOrderBuilder.BuildSignature(privateKey, orderHash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	limitOrderAndSignature, err := limitOrderBuilder.BuildLimitOrderAndSignature(limitOrder, orderHash, signature)
	assert.NotNil(t, limitOrderAndSignature)
	assert.Nil(t, err)

	assert.Equal(t, limitOrderAndSignature.OrderType, "limit")

	recoveredSignature, err := hex.DecodeString(limitOrderAndSignature.Signature[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredSignature, signature)

	assert.Positive(t, limitOrderAndSignature.Order.Salt)

	recoveredMakerAsset, err := hex.DecodeString(limitOrderAndSignature.Order.MakerAsset[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredMakerAsset, limitOrder.MakerAsset.Bytes())

	recoveredTakerAsset, err := hex.DecodeString(limitOrderAndSignature.Order.TakerAsset[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredTakerAsset, limitOrder.TakerAsset.Bytes())

	recoveredMakerAssetData, err := hex.DecodeString(limitOrderAndSignature.Order.MakerAssetData[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredMakerAssetData, limitOrder.MakerAssetData)

	recoveredTakerAssetData, err := hex.DecodeString(limitOrderAndSignature.Order.TakerAssetData[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredTakerAssetData, limitOrder.TakerAssetData)

	recoveredGetMakerAmount, err := hex.DecodeString(limitOrderAndSignature.Order.GetMakerAmount[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredGetMakerAmount, limitOrder.GetMakerAmount)

	recoveredGetTakerAmount, err := hex.DecodeString(limitOrderAndSignature.Order.GetTakerAmount[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredGetTakerAmount, limitOrder.GetTakerAmount)

	recoveredPredicate, err := hex.DecodeString(limitOrderAndSignature.Order.Predicate[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredPredicate, limitOrder.Predicate)

	recoveredPermit, err := hex.DecodeString(limitOrderAndSignature.Order.Permit[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredPermit, limitOrder.Permit)

	recoveredInteraction, err := hex.DecodeString(limitOrderAndSignature.Order.Interaction[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredInteraction, limitOrder.Interaction)

	recoveredSigner, err := hex.DecodeString(limitOrderAndSignature.Order.Signer[2:])
	assert.Nil(t, err)
	assert.Equal(t, recoveredSigner, limitOrder.Signer.Bytes())

	assert.Equal(t, limitOrderAndSignature.Order.SigType, int(limitOrder.SigType.Int64()))
}

func TestLimitOrderBuilderAndSign(t *testing.T) {
	limitOrderBuilder := getLimitOrderBuilderImpl(t)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signer, err := sign.GetPublicAddress(privateKey)
	assert.NotNil(t, signer)
	assert.Nil(t, err)

	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7651"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7652"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7653"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7654"),
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7655"),
		signer,
		[]byte("1"),
		[]byte("1"),
		[]byte("3"),
		big.NewInt(int64(1)),
		big.NewInt(int64(2)),
		big.NewInt(int64(3)),
		big.NewInt(int64(4)),
		big.NewInt(int64(5)),
		big.NewInt(int64(6)),
		model.POLY_PROXY,
	)
	assert.NotNil(t, limitOrder)
	assert.Nil(t, err)

	orderHash, err := limitOrderBuilder.BuildLimitOrderHash(limitOrder)
	assert.NotNil(t, orderHash)
	assert.Nil(t, err)

	signature, err := limitOrderBuilder.BuildSignature(privateKey, orderHash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	signature[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow papers
	match, err := limitOrderBuilder.ValidateSignature(limitOrder.Signer, orderHash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
