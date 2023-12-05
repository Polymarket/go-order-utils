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
		Taker:       "0x0",
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
	assert.Equal(t, order.Taker, common.HexToAddress("0x0"))
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
		Taker:       "0x1",
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
	assert.Equal(t, order.Taker, common.HexToAddress("0x1"))
	assert.Equal(t, order.TokenId.String(), "1234")
	assert.Equal(t, order.MakerAmount.String(), "100000000")
	assert.Equal(t, order.TakerAmount.String(), "50000000")
	assert.Equal(t, order.Side.String(), "0")
	assert.Equal(t, order.Expiration.String(), "0")
	assert.Equal(t, order.Nonce.String(), "0")
	assert.Equal(t, order.FeeRateBps.String(), "100")
	assert.Equal(t, order.SignatureType.String(), "0")
}

func TestBuildOrderHash(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err := builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	expectedOrderHash := common.Hex2Bytes("bf58957703791db2ab057528d03d1cff5375d9a475b14a9073bb7d892398dc96")
	assert.Equal(t, expectedOrderHash, orderHash[:])

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	expectedOrderHash = common.Hex2Bytes("053c3169ec6c9e99e3640cb12b9c9245917daf36bd4fd39ea09d201a07b53952")
	assert.Equal(t, expectedOrderHash, orderHash[:])
}

func TestBuildOrderSignature(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	order, err := builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err := builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err := builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.CTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	expectedSignature := common.Hex2Bytes("3874d2cfe79c0e82577f4f76ec58d950522ee5923a6f08441927d382c8178a5a2190fd4d5c49705e94d75999a58ec843f94a432e87ebc15cdb2186d315b3cc201b")
	assert.Equal(t, expectedSignature, orderSignature)

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	// specific salt
	builder = NewExchangeOrderBuilderImpl(chainId, func() int64 { return salt })

	order, err = builder.BuildOrder(&model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	orderHash, err = builder.BuildOrderHash(order, model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.NotNil(t, orderHash)

	orderSignature, err = builder.BuildOrderSignature(privateKey, orderHash)
	assert.NoError(t, err)
	assert.NotNil(t, orderSignature)

	expectedSignature = common.Hex2Bytes("b47e588cfb8630ffa255d1a04a4bb2b996967c2143fc107ab443282ed7dcc123288842968fa29e6f2e789e39ea02f13654d4dd55002b8672e9a91e2ba9045aa21b")
	assert.Equal(t, expectedSignature, orderSignature)
}

func TestBuildSignedOrder(t *testing.T) {
	// FEE
	// random salt
	builder := NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err := builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.CTFExchange)
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
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.CTFExchange)
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

	expectedSignature := common.Hex2Bytes("3874d2cfe79c0e82577f4f76ec58d950522ee5923a6f08441927d382c8178a5a2190fd4d5c49705e94d75999a58ec843f94a432e87ebc15cdb2186d315b3cc201b")
	assert.Equal(t, expectedSignature, signedOrder.Signature)

	// NegRisk
	// random salt
	builder = NewExchangeOrderBuilderImpl(chainId, nil)

	signedOrder, err = builder.BuildSignedOrder(privateKey, &model.OrderData{
		Maker:       signerAddress.Hex(),
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.NegRiskCTFExchange)
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
		Taker:       common.HexToAddress("0x0").Hex(),
		TokenId:     "1234",
		MakerAmount: "100000000",
		TakerAmount: "50000000",
		Side:        model.BUY,
		FeeRateBps:  "100",
		Nonce:       "0",
	}, model.NegRiskCTFExchange)
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

	expectedSignature = common.Hex2Bytes("b47e588cfb8630ffa255d1a04a4bb2b996967c2143fc107ab443282ed7dcc123288842968fa29e6f2e789e39ea02f13654d4dd55002b8672e9a91e2ba9045aa21b")
	assert.Equal(t, expectedSignature, signedOrder.Signature)
}
