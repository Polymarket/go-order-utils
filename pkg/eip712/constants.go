package eip712

import "github.com/ethereum/go-ethereum/crypto"

var EIP_712_TYPEHASH [32]byte = crypto.Keccak256Hash(
	[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
)
var LIMIT_ORDER_PROTOCOL_TYPE_HASH [32]byte = crypto.Keccak256Hash(
	[]byte("LimitOrder(uint256 salt,address makerAsset,address takerAsset,bytes makerAssetData,bytes takerAssetData,bytes getMakerAmount,bytes getTakerAmount,bytes predicate,bytes permit,bytes interaction,address signer,uint256 sigType)"),
)
var MARKET_ORDER_TYPE_HASH [32]byte = crypto.Keccak256Hash(
	[]byte("MarketOrder(uint256 salt,address signer,address maker,address makerAsset,uint256 makerAmount,uint256 makerAssetID,address takerAsset,uint256 takerAssetID,uint256 sigType)"),
)
