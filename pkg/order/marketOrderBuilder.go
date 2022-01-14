package order

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/fatih/structs"
	"github.com/polymarket/order-utils/pkg/eip712"
	"github.com/polymarket/order-utils/pkg/signature"
)

type MarketOrderBuilder interface {
	BuildMarketOrder(
		makerAssetAddress,
		takerAssetAddress,
		makerAddress,
		makerAmount,
		signer string,
		takerAssetID,
		makerAssetID int,
		sigType signature.SignatureType,
	) *MarketOrder
	BuildOrderSignature(walletAddress string, typedData signer.TypedData) (common.Hash, error)
}

type MarketOrderBuilderImpl struct {
	contractAddress string
	chainId         *math.HexOrDecimal256
	rpcClient       *ethclient.Client
	saltGenerator   func() int
}

func NewMarketOrderBuilderImpl(contractAddress string, chainId int, rpcClient *ethclient.Client, saltGenerator func() int) *MarketOrderBuilderImpl {
	if saltGenerator == nil {
		// TODO(REC): Implement a real salt func
		saltGenerator = func() int {
			return 0
		}
	}

	return &MarketOrderBuilderImpl{
		contractAddress: contractAddress,
		chainId:         math.NewHexOrDecimal256(int64(chainId)),
		rpcClient:       rpcClient,
		saltGenerator:   saltGenerator,
	}
}

func (m *MarketOrderBuilderImpl) BuildMarketOrder(
	makerAssetAddress,
	takerAssetAddress,
	makerAddress,
	makerAmount,
	signer string,
	takerAssetID,
	makerAssetID int,
	sigType signature.SignatureType,
) *MarketOrder {
	if signer == "" {
		signer = makerAddress
	}
	if sigType == 0 {
		sigType = signature.EOA
	}

	return &MarketOrder{
		Salt:         m.saltGenerator(),
		Signer:       signer,
		Maker:        makerAddress,
		MakerAsset:   makerAssetAddress,
		MakerAmount:  makerAmount,
		MakerAssetID: makerAssetID,
		TakerAsset:   takerAssetAddress,
		TakerAssetID: takerAssetID,
		SigType:      sigType,
	}
}

func (m *MarketOrderBuilderImpl) BuildOrderTypedData(order *MarketOrder) *signer.TypedData {
	return &signer.TypedData{
		PrimaryType: "MarketOrder",
		Types: signer.Types{
			"MarketOrder":  eip712.MARKET_ORDER_STRUCTURE,
			"EIP712Domain": eip712.EIP712_DOMAIN,
		},
		Domain: signer.TypedDataDomain{
			Name:              eip712.PROTOCOL_NAME,
			Version:           eip712.PROTOCOL_VERSION,
			ChainId:           m.chainId,
			VerifyingContract: m.contractAddress,
		},
		Message: structs.Map(order),
	}
}

func (m *MarketOrderBuilderImpl) BuildOrderSignature(typedData *signer.TypedData) (common.Hash, error) {
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return [common.HashLength]byte{}, err
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return [common.HashLength]byte{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)

	return hash, nil
}
