package builders

import (
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const (
	PROTOCOL_NAME    = "1inch Limit Order Protocol"
	PROTOCOL_VERSION = "1"
)

var EIP712_DOMAIN = []signer.Type{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
	{Name: "salt", Type: "string"},
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

var LIMIT_ORDER_STRUCTURE = []signer.Type{
	{Name: "salt", Type: "uint256"},
	{Name: "makerAsset", Type: "address"},
	{Name: "takerAsset", Type: "address"},
	{Name: "makerAssetData", Type: "bytes"},
	{Name: "takerAssetData", Type: "bytes"},
	{Name: "getMakerAmount", Type: "bytes"},
	{Name: "getTakerAmount", Type: "bytes"},
	{Name: "predicate", Type: "bytes"},
	{Name: "permit", Type: "bytes"},
	{Name: "interaction", Type: "bytes"},
	{Name: "signer", Type: "address"},
	{Name: "sigType", Type: "uint256"},
}
