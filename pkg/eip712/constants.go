package eip712

import (
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const (
	PROTOCOL_NAME    = "1inch Limit Order Protocol"
	PROTOCOL_VERSION = "1"
)

var EIP_712_TypeHASH [32]byte = crypto.Keccak256Hash(
	[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
)
var LIMIT_ORDER_PROTOCOL_Type_HASH [32]byte = crypto.Keccak256Hash(
	[]byte("Order(uint256 salt,address makerAsset,address takerAsset,bytes makerAssetData,bytes takerAssetData,bytes getMakerAmount,bytes getTakerAmount,bytes predicate,bytes permit,bytes interaction)"),
)

var EIP712_DOMAIN = []signer.Type{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}

var MARKET_ORDER_STRUCTURE = []signer.Type{
	{Name: "salt", Type: "uint256"},
	{Name: "signer", Type: "address"},
	{Name: "maker", Type: "address"},
	{Name: "makerAsset", Type: "address"},
	{Name: "makerAmount", Type: "uint256"},
	{Name: "makerAssetID", Type: "uint256"},
	{Name: "takerAsset", Type: "address"},
	{Name: "takerAssetID", Type: "uint256"},
	{Name: "sigType", Type: "uint256"},
}
