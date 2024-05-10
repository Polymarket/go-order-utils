package builder

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/eip712"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/signer"
	"github.com/polymarket/go-order-utils/pkg/utils"
)

type ExchangeOrderBuilderImpl struct {
	chainId       *big.Int
	saltGenerator func() int64
}

var _ ExchangeOrderBuilder = (*ExchangeOrderBuilderImpl)(nil)

func NewExchangeOrderBuilderImpl(chainId *big.Int, saltGenerator func() int64) *ExchangeOrderBuilderImpl {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}
	return &ExchangeOrderBuilderImpl{
		chainId:       chainId,
		saltGenerator: saltGenerator,
	}
}

// build an order object including the signature.
//
// @param private key
//
// @param orderData
//
// @returns a SignedOrder object (order + signature)
func (e *ExchangeOrderBuilderImpl) BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, contract model.VerifyingContract) (*model.SignedOrder, error) {
	order, err := e.BuildOrder(orderData)
	if err != nil {
		return nil, err
	}

	orderHash, err := e.BuildOrderHash(order, contract)
	if err != nil {
		return nil, err
	}

	signature, err := e.BuildOrderSignature(privateKey, orderHash)
	if err != nil {
		return nil, err
	}

	ok, err := signer.ValidateSignature(order.Signer, orderHash, signature)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("signature error")
	}

	return &model.SignedOrder{
		Order:     *order,
		Signature: signature,
	}, nil
}

// Creates an Order object from order data.
//
// @param orderData
//
// @returns a Order object (not signed)
func (e *ExchangeOrderBuilderImpl) BuildOrder(orderData *model.OrderData) (*model.Order, error) {
	var signer common.Address
	if orderData.Signer == "" {
		signer = common.HexToAddress(orderData.Maker)
	} else {
		signer = common.HexToAddress(orderData.Signer)
	}

	var tokenId *big.Int
	var ok bool
	if tokenId, ok = new(big.Int).SetString(orderData.TokenId, 10); !ok {
		return nil, fmt.Errorf("can't parse TokenId: %s as valid *big.Int", orderData.TokenId)
	}

	var makerAmount *big.Int
	if makerAmount, ok = new(big.Int).SetString(orderData.MakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse MakerAmount: %s as valid *big.Int", orderData.MakerAmount)
	}

	var takerAmount *big.Int
	if takerAmount, ok = new(big.Int).SetString(orderData.TakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse TakerAmount: %s as valid *big.Int", orderData.TakerAmount)
	}

	var expiration *big.Int
	if orderData.Expiration == "" {
		orderData.Expiration = "0"
	}
	if expiration, ok = new(big.Int).SetString(orderData.Expiration, 10); !ok {
		return nil, fmt.Errorf("can't parse Expiration: %s as valid *big.Int", orderData.Expiration)
	}

	var nonce *big.Int
	if nonce, ok = new(big.Int).SetString(orderData.Nonce, 10); !ok {
		return nil, fmt.Errorf("can't parse Nonce: %s as valid *big.Int", orderData.Nonce)
	}

	var feeRateBps *big.Int
	if feeRateBps, ok = new(big.Int).SetString(orderData.FeeRateBps, 10); !ok {
		return nil, fmt.Errorf("can't parse FeeRateBps: %s as valid *big.Int", orderData.FeeRateBps)
	}

	return &model.Order{
		Salt:          new(big.Int).SetInt64(e.saltGenerator()),
		Maker:         common.HexToAddress(orderData.Maker),
		Taker:         common.HexToAddress(orderData.Taker),
		Signer:        signer,
		TokenId:       tokenId,
		MakerAmount:   makerAmount,
		TakerAmount:   takerAmount,
		Side:          new(big.Int).SetInt64(int64(orderData.Side)),
		Expiration:    expiration,
		Nonce:         nonce,
		FeeRateBps:    feeRateBps,
		SignatureType: new(big.Int).SetInt64(int64(orderData.SignatureType)),
	}, nil
}

// Generates the hash of the order from a EIP712TypedData object.
//
// @param Order
//
// @returns a OrderHash that is a 'common.Hash'
func (e *ExchangeOrderBuilderImpl) BuildOrderHash(order *model.Order, contract model.VerifyingContract) (model.OrderHash, error) {
	verifyingContract, err := utils.GetVerifyingContractAddress(e.chainId, contract)
	if err != nil {
		return model.OrderHash{}, err
	}

	domainSeparator, err := eip712.BuildEIP712DomainSeparator(_PROTOCOL_NAME, _PROTOCOL_VERSION, e.chainId, verifyingContract)
	if err != nil {
		return model.OrderHash{}, err
	}

	values := []interface{}{
		_ORDER_STRUCTURE_HASH,
		order.Salt,
		order.Maker,
		order.Signer,
		order.Taker,
		order.TokenId,
		order.MakerAmount,
		order.TakerAmount,
		order.Expiration,
		order.Nonce,
		order.FeeRateBps,
		uint8(order.Side.Uint64()),
		uint8(order.SignatureType.Uint64()),
	}
	orderHash, err := eip712.HashTypedDataV4(domainSeparator, _ORDER_STRUCTURE, values)
	if err != nil {
		return model.OrderHash{}, err
	}

	return orderHash, nil
}

// signs an order
//
// @param private key
//
// @param order hash
//
// @returns a OrderSignature that is []byte
func (e *ExchangeOrderBuilderImpl) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error) {
	return signer.Sign(privateKey, orderHash)
}
