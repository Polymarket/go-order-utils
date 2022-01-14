package builders

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/fatih/structs"
	"github.com/polymarket/order-utils/pkg/facades"
	"github.com/polymarket/order-utils/pkg/model"
	"github.com/polymarket/order-utils/pkg/utils"
)

type LimitOrderBuilder interface {
	BuildLimitOrder(
		exchangeAddress,
		makerAssetAddress,
		makerAddress,
		takerAssetAddress,
		takerAddress,
		signer common.Address,
		permit,
		interaction,
		predicate []byte,
		makerAmount,
		takerAmount,
		makerAssetID,
		takerAssetID,
		expiry,
		nonce *big.Int,
		sigType model.SignatureType,
	) (*model.LimitOrder, error)
	BuildLimitOrderTypedData(order *model.LimitOrder) *signer.TypedData
	BuildLimitOrderSignature(typedData *signer.TypedData) (common.Hash, error)
}

type LimitOrderBuilderImpl struct {
	contractAddress            common.Address
	chainId                    *math.HexOrDecimal256
	saltGenerator              func() int64
	limitOrderPredicateBuilder LimitOrderPredicateBuilder
	erc20Facade                facades.ERC20Facade
	erc1155Facade              facades.ERC1155Facade
	limitOrderProtocolFacade   facades.LimitOrderProtocolFacade
}

func NewLimitOrderBuilderImpl(contractAddress common.Address, chainId int, saltGenerator func() int64,
	limitOrderPredicateBuilder LimitOrderPredicateBuilder, erc20Facade facades.ERC20Facade,
	erc1155Facade facades.ERC1155Facade, limitOrderProtocolFacade facades.LimitOrderProtocolFacade) *LimitOrderBuilderImpl {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}

	return &LimitOrderBuilderImpl{
		contractAddress:            contractAddress,
		chainId:                    math.NewHexOrDecimal256(int64(chainId)),
		saltGenerator:              saltGenerator,
		limitOrderPredicateBuilder: limitOrderPredicateBuilder,
		erc20Facade:                erc20Facade,
		erc1155Facade:              erc1155Facade,
		limitOrderProtocolFacade:   limitOrderProtocolFacade,
	}
}

func (l *LimitOrderBuilderImpl) BuildLimitOrder(
	exchangeAddress,
	makerAssetAddress,
	makerAddress,
	takerAssetAddress,
	takerAddress,
	signer common.Address,
	permit,
	interaction,
	predicate []byte,
	makerAmount,
	takerAmount,
	makerAssetID,
	takerAssetID,
	expiry,
	nonce *big.Int,
	sigType model.SignatureType,
) (*model.LimitOrder, error) {
	var makerAssetData []byte
	var makerAsset common.Address
	if makerAssetID.Int64() > 0 {
		makerAsset = exchangeAddress

		var err error
		makerAssetData, err = l.erc1155Facade.TransferFrom(makerAddress, takerAddress, makerAssetAddress, makerAmount, makerAssetID)
		if err != nil {
			return nil, err
		}
	} else {
		makerAsset = makerAssetAddress

		var err error
		makerAssetData, err = l.erc20Facade.TransferFrom(makerAddress, takerAddress, makerAmount)
		if err != nil {
			return nil, err
		}
	}

	var takerAssetData []byte
	var takerAsset common.Address
	if takerAssetID.Int64() > 0 {
		takerAsset = exchangeAddress

		var err error
		takerAssetData, err = l.erc1155Facade.TransferFrom(takerAddress, makerAddress, takerAssetAddress, takerAmount, takerAssetID)
		if err != nil {
			return nil, err
		}
	} else {
		takerAsset = takerAssetAddress

		var err error
		takerAssetData, err = l.erc20Facade.TransferFrom(takerAddress, makerAddress, takerAmount)
		if err != nil {
			return nil, err
		}
	}

	getMakerAmount, err := l.limitOrderProtocolFacade.GetMakerAmount(makerAmount, takerAmount, &big.Int{})
	if err != nil {
		return nil, err
	}

	getTakerAmount, err := l.limitOrderProtocolFacade.GetTakerAmount(makerAmount, takerAmount, &big.Int{})
	if err != nil {
		return nil, err
	}

	if expiry.Int64() > 0 && nonce.Int64() > 0 {
		timestampBelow, err := l.limitOrderPredicateBuilder.TimestampBelow(expiry)
		if err != nil {
			return nil, err
		}

		nonceEquals, err := l.limitOrderPredicateBuilder.NonceEquals(makerAddress, nonce)
		if err != nil {
			return nil, err
		}

		predicate, err = l.limitOrderPredicateBuilder.And(
			timestampBelow,
			nonceEquals,
		)
		if err != nil {
			return nil, err
		}
	}

	if signer.String() == "" {
		signer = makerAddress
	}

	if sigType == 0 {
		sigType = model.EOA
	}

	return &model.LimitOrder{
		Salt:           big.NewInt(l.saltGenerator()),
		MakerAsset:     makerAsset,
		TakerAsset:     takerAsset,
		MakerAssetData: makerAssetData,
		TakerAssetData: takerAssetData,
		GetMakerAmount: getMakerAmount,
		GetTakerAmount: getTakerAmount,
		Predicate:      predicate,
		Permit:         permit,
		Interaction:    interaction,
		Signer:         signer,
		SigType:        big.NewInt(int64(sigType)),
	}, nil
}

func (l *LimitOrderBuilderImpl) BuildLimitOrderTypedData(order *model.LimitOrder) *signer.TypedData {
	return &signer.TypedData{
		PrimaryType: "LimitOrder",
		Types: signer.Types{
			"LimitOrder":   LIMIT_ORDER_STRUCTURE,
			"EIP712Domain": EIP712_DOMAIN,
		},
		Domain: signer.TypedDataDomain{
			Name:              PROTOCOL_NAME,
			Version:           PROTOCOL_VERSION,
			ChainId:           l.chainId,
			VerifyingContract: l.contractAddress.String(),
			Salt:              fmt.Sprintf("%d", order.Salt.Int64()),
		},
		Message: structs.Map(order),
	}
}

func (l *LimitOrderBuilderImpl) BuildLimitOrderSignature(typedData *signer.TypedData) (common.Hash, error) {
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
