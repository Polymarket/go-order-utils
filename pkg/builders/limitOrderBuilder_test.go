package builders

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/polymarket/go-order-utils/pkg/facades"
	"github.com/polymarket/go-order-utils/pkg/model"
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

func TestBuildLimitOrderTypedData(t *testing.T) {
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

	orderTypedData := limitOrderBuilder.BuildLimitOrderTypedData(limitOrder)
	assert.NotNil(t, orderTypedData)
	assert.Equal(t, orderTypedData.PrimaryType, "LimitOrder")
	assert.Equal(t, orderTypedData.Types["LimitOrder"], LIMIT_ORDER_STRUCTURE)
	assert.Equal(t, orderTypedData.Types["EIP712Domain"], EIP712_DOMAIN)
	assert.Equal(t, orderTypedData.Domain.Name, PROTOCOL_NAME)
	assert.Equal(t, orderTypedData.Domain.Version, PROTOCOL_VERSION)
	assert.Equal(t, orderTypedData.Domain.ChainId, math.NewHexOrDecimal256(int64(chainId)))
	assert.Equal(t, orderTypedData.Domain.VerifyingContract, contractAddress)
	assert.NotNil(t, orderTypedData.Domain.Salt)

	assert.NotNil(t, orderTypedData.Message)
	assert.Equal(t, orderTypedData.Message, signer.TypedDataMessage{
		"salt":           limitOrder.Salt.String(),
		"makerAsset":     limitOrder.MakerAsset.String(),
		"takerAsset":     limitOrder.TakerAsset.String(),
		"makerAssetData": limitOrder.MakerAssetData,
		"takerAssetData": limitOrder.TakerAssetData,
		"getMakerAmount": limitOrder.GetMakerAmount,
		"getTakerAmount": limitOrder.GetTakerAmount,
		"predicate":      limitOrder.Predicate,
		"permit":         limitOrder.Permit,
		"interaction":    limitOrder.Interaction,
		"signer":         limitOrder.Signer.String(),
		"sigType":        limitOrder.SigType.String(),
	})
}

func TestLimitOrderBuilderAndSign(t *testing.T) {
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

	orderTypedData := limitOrderBuilder.BuildLimitOrderTypedData(limitOrder)
	assert.NotNil(t, orderTypedData)

	hash, err := limitOrderBuilder.BuildHash(orderTypedData)
	assert.NotNil(t, hash)
	assert.Nil(t, err)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := limitOrderBuilder.BuildSignature(privateKey, hash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	match, err := limitOrderBuilder.ValidateSignature(&privateKey.PublicKey, hash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
