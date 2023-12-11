package builder

import (
	"crypto/ecdsa"

	"github.com/polymarket/go-order-utils/pkg/model"
)

//go:generate mockery --name ExchangeOrderBuilder
type ExchangeOrderBuilder interface {
	// build an order object including the signature.
	//
	// @param private key
	//
	// @param orderData
	//
	// @returns a SignedOrder object (order + signature)
	BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, contract model.VerifyingContract) (*model.SignedOrder, error)

	// Creates an Order object from order data.
	//
	// @param orderData
	//
	// @returns a Order object (not signed)
	BuildOrder(orderData *model.OrderData) (*model.Order, error)

	// Generates the hash of the order from a EIP712TypedData object.
	//
	// @param Order
	//
	// @returns a OrderHash that is a 'common.Hash'
	BuildOrderHash(order *model.Order, contract model.VerifyingContract) (model.OrderHash, error)

	// signs an order
	//
	// @param private key
	//
	// @param order hash
	//
	// @returns a OrderSignature that is []byte
	BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHash) (model.OrderSignature, error)
}
