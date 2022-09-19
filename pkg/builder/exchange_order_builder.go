package builder

import (
	"crypto/ecdsa"

	"github.com/polymarket/go-order-utils/pkg/model"
)

type ExchangeOrderBuilder interface {
	// build an order object including the signature.
	//
	// @param private key
	//
	// @param orderData
	//
	// @returns a SignedOrder object (order + signature)
	BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData) (*model.SignedOrder, error)

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
	// @returns a OrderTypedData that is a 'common.Hash'
	BuildOrderTypedData(order *model.Order) (model.OrderTypedData, error)

	// signs an order
	//
	// @param private key
	//
	// @param orderData
	//
	// @returns a OrderSignature that is a 'common.Hash'
	BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderTypedData model.OrderTypedData) (model.OrderSignature, error)
}