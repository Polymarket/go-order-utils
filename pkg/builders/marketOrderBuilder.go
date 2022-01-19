package builders

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/sign"
	"github.com/polymarket/go-order-utils/pkg/utils"
)

type MarketOrderBuilder interface {
	sign.Signer
	BuildMarketOrder(
		makerAssetAddress,
		takerAssetAddress,
		makerAddress,
		signer common.Address,
		makerAmount,
		takerAssetID,
		makerAssetID *big.Int,
		sigType model.SignatureType,
	) *model.MarketOrder
	BuildMarketOrderTypedData(order *model.MarketOrder) *signer.TypedData
}

type MarketOrderBuilderImpl struct {
	sign.Signer
	contractAddress common.Address
	chainId         *math.HexOrDecimal256
	saltGenerator   func() int64
}

func NewMarketOrderBuilderImpl(contractAddress common.Address, chainId int, saltGenerator func() int64) *MarketOrderBuilderImpl {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}

	return &MarketOrderBuilderImpl{
		Signer:          sign.NewSignerImpl(),
		contractAddress: contractAddress,
		chainId:         math.NewHexOrDecimal256(int64(chainId)),
		saltGenerator:   saltGenerator,
	}
}

func (m *MarketOrderBuilderImpl) BuildMarketOrder(
	makerAssetAddress,
	takerAssetAddress,
	makerAddress,
	signer common.Address,
	makerAmount,
	takerAssetID,
	makerAssetID *big.Int,
	sigType model.SignatureType,
) *model.MarketOrder {
	if signer.String() == "" {
		signer = makerAddress
	}

	if sigType < model.EOA || sigType > model.POLY_PROXY {
		sigType = model.EOA
	}

	return &model.MarketOrder{
		Salt:         big.NewInt(m.saltGenerator()),
		Signer:       signer,
		Maker:        makerAddress,
		MakerAsset:   makerAssetAddress,
		MakerAmount:  makerAmount,
		MakerAssetID: makerAssetID,
		TakerAsset:   takerAssetAddress,
		TakerAssetID: takerAssetID,
		SigType:      big.NewInt(int64(sigType)),
	}
}

func (m *MarketOrderBuilderImpl) BuildMarketOrderTypedData(order *model.MarketOrder) *signer.TypedData {
	return &signer.TypedData{
		PrimaryType: "MarketOrder",
		Types: signer.Types{
			"MarketOrder":  MARKET_ORDER_STRUCTURE,
			"EIP712Domain": EIP712_DOMAIN,
		},
		Domain: signer.TypedDataDomain{
			Name:              PROTOCOL_NAME,
			Version:           PROTOCOL_VERSION,
			ChainId:           m.chainId,
			VerifyingContract: m.contractAddress.String(),
			Salt:              fmt.Sprintf("%d", order.Salt.Int64()),
		},
		Message: signer.TypedDataMessage{
			"salt":         order.Salt.String(),
			"signer":       order.Signer.String(),
			"maker":        order.Maker.String(),
			"makerAsset":   order.MakerAsset.String(),
			"makerAmount":  order.MakerAmount.String(),
			"makerAssetID": order.MakerAssetID.String(),
			"takerAsset":   order.TakerAsset.String(),
			"takerAssetID": order.TakerAssetID.String(),
			"sigType":      order.SigType.String(),
		},
	}
}
