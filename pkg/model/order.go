package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OrderSignature = []byte

type OrderHash = common.Hash

type OrderData struct {
	// Maker of the order, i.e the source of funds for the order
	Maker string

	// Address of the order taker. The zero address is used to indicate a public order
	Taker string

	// Token Id of the CTF ERC1155 asset to be bought or sold.
	// If BUY, this is the tokenId of the asset to be bought, i.e the makerAssetId
	// If SELL, this is the tokenId of the asset to be sold, i.e the  takerAssetId
	TokenId string

	// Maker amount, i.e the max amount of tokens to be sold
	MakerAmount string

	// Taker amount, i.e the minimum amount of tokens to be received
	TakerAmount string

	// Fee rate, in basis points, charged to the order maker, charged on proceeds
	FeeRateBps string

	// Nonce used for onchain cancellations
	Nonce string

	// Signer of the order. Optional, if it is not present the signer is the maker of the order.
	Signer string

	// Timestamp after which the order is expired.
	// Optional, if it is not present the value is '0' (no expiration)
	Expiration string

	// The side of the order, BUY or SELL
	Side Side

	// Signature type used by the Order. Default value 'EOA'
	SignatureType SignatureType
}

type Order struct {
	//  Unique salt to ensure entropy
	Salt *big.Int

	// Token Id of the CTF ERC1155 asset to be bought or sold.
	// If BUY, this is the tokenId of the asset to be bought, i.e the makerAssetId
	// If SELL, this is the tokenId of the asset to be sold, i.e the  takerAssetId
	TokenId *big.Int

	// Maker amount, i.e the max amount of tokens to be sold
	MakerAmount *big.Int

	// Taker amount, i.e the minimum amount of tokens to be received
	TakerAmount *big.Int

	// The side of the order, BUY or SELL
	Side *big.Int

	// Timestamp after which the order is expired
	Expiration *big.Int

	// Nonce used for onchain cancellations
	Nonce *big.Int

	// Fee rate, in basis points, charged to the order maker, charged on proceeds
	FeeRateBps *big.Int

	// Signature type used by the Order
	SignatureType *big.Int

	// Maker of the order, i.e the source of funds for the order
	Maker common.Address

	// Address of the order taker. The zero address is used to indicate a public order
	Taker common.Address

	// Signer of the order
	Signer common.Address
}

type SignedOrder struct {
	Order

	// The order signature
	Signature OrderSignature
}
