package builder

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	chainId = new(big.Int).SetInt64(80001)
	// publicly known private key
	privateKey, _ = crypto.ToECDSA(common.Hex2Bytes("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	// private key address
	signerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	salt = int64(479249096354)
)

func TestBuildOrder(t *testing.T) {
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.True(t, order.Salt.Int64() > 0)
	assert.Equal(t, order.Maker, signerAddress)
	assert.Equal(t, order.Signer, signerAddress)
	assert.Equal(t, order.TokenId.String(), "1234")
	assert.Equal(t, order.MakerAmount.String(), "100000000")
	assert.Equal(t, order.TakerAmount.String(), "50000000")
	assert.Equal(t, order.Side.String(), "0")
	assert.Equal(t, order.Expiration.String(), "0")
	assert.Equal(t, order.Nonce.String(), "0")
	assert.Equal(t, order.FeeRateBps.String(), "100")
	assert.Equal(t, order.SignatureType.String(), "0")

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.Equal(t, order.Salt.Int64(), int64(salt))
	assert.Equal(t, order.Maker, signerAddress)
	assert.Equal(t, order.Signer, signerAddress)
	assert.Equal(t, order.TokenId.String(), "1234")
	assert.Equal(t, order.MakerAmount.String(), "100000000")
	assert.Equal(t, order.TakerAmount.String(), "50000000")
	assert.Equal(t, order.Side.String(), "0")
	assert.Equal(t, order.Expiration.String(), "0")
	assert.Equal(t, order.Nonce.String(), "0")
	assert.Equal(t, order.FeeRateBps.String(), "100")
	assert.Equal(t, order.SignatureType.String(), "0")
}

func TestBuildOrderTypedData(t *testing.T) {
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderTypedData, err := builder.BuildOrderTypedData(order)
	assert.NoError(t, err)
	assert.NotNil(t, orderTypedData)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderTypedData, err = builder.BuildOrderTypedData(order)
	assert.NoError(t, err)
	assert.NotNil(t, orderTypedData)

	expectedTypedData := common.Hex2Bytes("c13f7f9d048468a07de3f203b5079d21cd46127b1a0e35750e250889e61ddc05")
	assert.Equal(t, expectedTypedData, orderTypedData[:])
}

func TestBuildOrderSignature(t *testing.T) {
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderTypedData, err := builder.BuildOrderTypedData(order)
	assert.NoError(t, err)
	assert.NotNil(t, orderTypedData)

	orderSignature, err := builder.BuildOrderSignature(privateKey, orderTypedData)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderTypedData, err = builder.BuildOrderTypedData(order)
	assert.NoError(t, err)
	assert.NotNil(t, orderTypedData)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderTypedData)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	expectedSignature := common.Hex2Bytes("369a9e418825caf981c5640f29452a9af0f3063cc345e9163f7395e617bb995965b97ce89a8b1ad229276dc67d0c26caaa76732fbd9ec03d018d8b4bbc81aaaf1b")
	assert.Equal(t, expectedSignature, orderSignature)
}

func TestBuildSignedOrder(t *testing.T) {
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err := builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.True(t, signedOrder.Salt.Int64() > 0)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, signedOrder.Signature)
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	signedOrder, err = builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, signedOrder)

	assert.Equal(t, signedOrder.Salt.Int64(), salt)
	assert.Equal(t, signedOrder.Maker, signerAddress)
	assert.Equal(t, signedOrder.Signer, signerAddress)
	assert.Equal(t, signedOrder.TokenId.String(), "1234")
	assert.Equal(t, signedOrder.MakerAmount.String(), "100000000")
	assert.Equal(t, signedOrder.TakerAmount.String(), "50000000")
	assert.Equal(t, signedOrder.Side.String(), "0")
	assert.Equal(t, signedOrder.Expiration.String(), "0")
	assert.Equal(t, signedOrder.Nonce.String(), "0")
	assert.Equal(t, signedOrder.FeeRateBps.String(), "100")
	assert.Equal(t, signedOrder.SignatureType.String(), "0")
	assert.NotEmpty(t, hex.EncodeToString(signedOrder.Signature))

	expectedSignature := common.Hex2Bytes("369a9e418825caf981c5640f29452a9af0f3063cc345e9163f7395e617bb995965b97ce89a8b1ad229276dc67d0c26caaa76732fbd9ec03d018d8b4bbc81aaaf1b")
	assert.Equal(t, expectedSignature, signedOrder.Signature)
}
