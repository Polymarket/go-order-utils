package builders

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/config"
	"github.com/polymarket/go-order-utils/pkg/eip712"
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
	BuildMarketOrderHash(order *model.MarketOrder) (common.Hash, error)
	BuildMarketOrderAndSignature(order *model.MarketOrder, signature []byte) *model.MarketOrderAndSignature
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

func (m *MarketOrderBuilderImpl) BuildMarketOrderHash(order *model.MarketOrder) (common.Hash, error) {
	c, err := config.GetContracts(int((*big.Int)(m.chainId).Int64()))
	if err != nil {
		return [32]byte{}, err
	}

	exchangeAddress := common.HexToAddress(c.Exchange.Address)
	nameHash := crypto.Keccak256Hash([]byte(c.Exchange.Name))
	versionHash := crypto.Keccak256Hash([]byte(c.Exchange.Version))

	domainSeparator, err := eip712.BuildEIP712DomainSeparator(nameHash, versionHash, (*big.Int)(m.chainId), exchangeAddress)
	if err != nil {
		return [32]byte{}, err
	}
	var domainSep32Bytes [32]byte
	copy(domainSep32Bytes[:], domainSeparator)

	types := []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Uint256, // salt
		eip712.Address, // signer
		eip712.Address, // maker
		eip712.Address, // makerAsset
		eip712.Uint256, // makerAmount
		eip712.Uint256, // makerAssetID
		eip712.Address, // takerAsset
		eip712.Uint256, // takerAssetID
		eip712.Uint256, // sig type
	}

	values := []interface{}{
		eip712.MARKET_ORDER_TYPE_HASH,
		order.Salt,
		order.Signer,
		order.Maker,
		order.MakerAsset,
		order.MakerAmount,
		order.MakerAssetID,
		order.TakerAsset,
		order.TakerAssetID,
		order.SigType,
	}

	encoded, err := eip712.Encode(
		types, values,
	)
	if err != nil {
		return [32]byte{}, err
	}

	orderHash := eip712.HashTypedDataV4(domainSep32Bytes,
		crypto.Keccak256Hash(
			encoded,
		),
	)
	var orderHash32Bytes [32]byte
	copy(orderHash32Bytes[:], orderHash)
	return orderHash32Bytes, nil
}

func (m *MarketOrderBuilderImpl) BuildMarketOrderAndSignature(order *model.MarketOrder, signature []byte) *model.MarketOrderAndSignature {
	return &model.MarketOrderAndSignature{
		Order: &model.CannonicalMarketOrder{
			Salt:         int(order.Salt.Int64()),
			Signer:       "0x" + hex.EncodeToString(order.Signer.Bytes()),
			Maker:        "0x" + hex.EncodeToString(order.Maker.Bytes()),
			MakerAsset:   "0x" + hex.EncodeToString(order.MakerAsset.Bytes()),
			MakerAmount:  order.MakerAmount.String(),
			MakerAssetID: int(order.MakerAssetID.Int64()),
			TakerAsset:   "0x" + hex.EncodeToString(order.TakerAsset.Bytes()),
			TakerAssetID: int(order.TakerAssetID.Int64()),
			SigType:      int(order.SigType.Int64()),
		},
		Signature: "0x" + hex.EncodeToString(signature),
		OrderType: "market",
	}
}
