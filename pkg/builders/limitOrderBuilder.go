package builders

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/config"
	"github.com/polymarket/go-order-utils/pkg/eip712"
	"github.com/polymarket/go-order-utils/pkg/facades"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/sign"
	"github.com/polymarket/go-order-utils/pkg/utils"
)

type LimitOrderBuilder interface {
	sign.Signer
	BuildLimitOrder(
		exchangeAddress,
		makerAssetAddress,
		makerAddress,
		takerAssetAddress,
		takerAddress,
		signer common.Address,
		predicate []byte,
		makerAmount,
		takerAmount,
		makerAssetID,
		takerAssetID,
		expiry,
		nonce *big.Int,
		sigType model.SignatureType,
	) (*model.LimitOrder, error)
	BuildLimitOrderHash(order *model.LimitOrder) (common.Hash, error)
	BuildLimitOrderAndSignature(order *model.LimitOrder, orderHash common.Hash, signature []byte) (*model.LimitOrderAndSignature, error)
}

type LimitOrderBuilderImpl struct {
	sign.Signer
	contractAddress            common.Address
	chainId                    *math.HexOrDecimal256
	saltGenerator              func() int64
	limitOrderPredicateBuilder LimitOrderPredicateBuilder
	erc20Facade                facades.ERC20Facade
	erc1155Facade              facades.ERC1155Facade
	limitOrderProtocolFacade   facades.LimitOrderProtocolFacade
}

func NewLimitOrderBuilderImpl(contractAddress common.Address, chainId int, saltGenerator func() int64) (*LimitOrderBuilderImpl, error) {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}

	limitOrderProtocolFacade, err := facades.NewLimitOrderProtocolFacadeImpl()
	if err != nil {
		return nil, err
	}

	limitOrderPredicateBuilder := NewLimitOrderPredicateBuilderImpl(limitOrderProtocolFacade)

	erc20Facade, err := facades.NewERC20FacadeImpl()

	if err != nil {
		return nil, err
	}

	erc1155Facade, err := facades.NewERC1155FacadeImpl()

	if err != nil {
		return nil, err
	}

	return &LimitOrderBuilderImpl{
		Signer:                     sign.NewSignerImpl(),
		contractAddress:            contractAddress,
		chainId:                    math.NewHexOrDecimal256(int64(chainId)),
		saltGenerator:              saltGenerator,
		limitOrderPredicateBuilder: limitOrderPredicateBuilder,
		erc20Facade:                erc20Facade,
		erc1155Facade:              erc1155Facade,
		limitOrderProtocolFacade:   limitOrderProtocolFacade,
	}, nil
}

func (l *LimitOrderBuilderImpl) BuildLimitOrder(
	exchangeAddress,
	makerAssetAddress,
	makerAddress,
	takerAssetAddress,
	takerAddress,
	signer common.Address,
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
	if makerAssetID.Int64() >= 0 {
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
	if takerAssetID.Int64() >= 0 {
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

	getMakerAmount, err := l.limitOrderProtocolFacade.GetMakerAmount(makerAmount, takerAmount, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	getTakerAmount, err := l.limitOrderProtocolFacade.GetTakerAmount(makerAmount, takerAmount, big.NewInt(0))
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
			l.contractAddress,
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

	if sigType < model.EOA || sigType > model.POLY_PROXY {
		sigType = model.EOA
	}

	return &model.LimitOrder{
		Salt:           big.NewInt(l.saltGenerator()),
		MakerAsset:     makerAsset,
		TakerAsset:     takerAsset,
		MakerAssetData: makerAssetData,
		TakerAssetData: takerAssetData,
		GetMakerAmount: getMakerAmount[0:68],
		GetTakerAmount: getTakerAmount[0:68],
		Predicate:      predicate,
		Signer:         signer,
		SigType:        big.NewInt(int64(sigType)),
	}, nil
}

func (l *LimitOrderBuilderImpl) BuildLimitOrderHash(order *model.LimitOrder) (common.Hash, error) {
	c, err := config.GetContracts(int((*big.Int)(l.chainId).Int64()))
	if err != nil {
		return [32]byte{}, err
	}
	exchangeAddress := common.HexToAddress(c.Exchange.Address)
	nameHash := crypto.Keccak256Hash([]byte(c.Exchange.Name))
	versionHash := crypto.Keccak256Hash([]byte(c.Exchange.Version))

	domainSeparator, err := eip712.BuildEIP712DomainSeparator(nameHash, versionHash, (*big.Int)(l.chainId), exchangeAddress)
	if err != nil {
		return [32]byte{}, err
	}
	var domainSep32Bytes [32]byte
	copy(domainSep32Bytes[:], domainSeparator)

	types := []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Uint256, // salt
		eip712.Address, // makerAsset
		eip712.Address, // takerAsset
		eip712.Bytes32, // makerAssetData hash
		eip712.Bytes32, // takerAssetData hash
		eip712.Bytes32, // getMakerAmount hash
		eip712.Bytes32, // getTakerAmount hash
		eip712.Bytes32, // predicate hash
		eip712.Address, // signer
		eip712.Uint256, // sig type
	}

	values := []interface{}{
		eip712.LIMIT_ORDER_PROTOCOL_TYPE_HASH,
		order.Salt,
		order.MakerAsset,
		order.TakerAsset,
		crypto.Keccak256Hash(order.MakerAssetData),
		crypto.Keccak256Hash(order.TakerAssetData),
		crypto.Keccak256Hash(order.GetMakerAmount),
		crypto.Keccak256Hash(order.GetTakerAmount),
		crypto.Keccak256Hash(order.Predicate),
		order.Signer,
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

func (l *LimitOrderBuilderImpl) BuildLimitOrderAndSignature(order *model.LimitOrder, orderHash common.Hash,
	signature []byte) (*model.LimitOrderAndSignature, error) {
	sigCopy := make([]byte, len(signature))
	copy(sigCopy, signature)
	sigCopy[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow paper
	match, err := l.ValidateSignature(order.Signer, orderHash, sigCopy)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, fmt.Errorf("order signer does not match with the generated signature")
	}

	return &model.LimitOrderAndSignature{
		Order: &model.CannonicalLimitOrder{
			Salt:           int(order.Salt.Int64()),
			MakerAsset:     "0x" + hex.EncodeToString(order.MakerAsset.Bytes()),
			TakerAsset:     "0x" + hex.EncodeToString(order.TakerAsset.Bytes()),
			MakerAssetData: "0x" + hex.EncodeToString(order.MakerAssetData),
			TakerAssetData: "0x" + hex.EncodeToString(order.TakerAssetData),
			GetMakerAmount: "0x" + hex.EncodeToString(order.GetMakerAmount),
			GetTakerAmount: "0x" + hex.EncodeToString(order.GetTakerAmount),
			Predicate:      "0x" + hex.EncodeToString(order.Predicate),
			Signer:         "0x" + hex.EncodeToString(order.Signer.Bytes()),
			SigType:        int(order.SigType.Int64()),
		},
		Signature: "0x" + hex.EncodeToString(signature),
		OrderType: "limit",
	}, nil
}
