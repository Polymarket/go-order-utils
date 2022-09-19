package builder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/eip712"
)

var (
	_PROTOCOL_NAME    = crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	_PROTOCOL_VERSION = crypto.Keccak256Hash([]byte("1"))
)

var (
	_ORDER_STRUCTURE = []abi.Type{
		eip712.Bytes32, // typehash
		eip712.Uint256, // salt
		eip712.Address, // maker
		eip712.Address, // signer
		eip712.Uint256, // tokenId
		eip712.Uint256, // makerAmount
		eip712.Uint256, // takerAmount
		eip712.Uint256, // side
		eip712.Uint256, // expiration
		eip712.Uint256, // nonce
		eip712.Uint256, // feeRateBps
		eip712.Uint256, // signatureType
	}
)

var (
	_ORDER_STRUCTURE_HASH = crypto.Keccak256Hash(
		[]byte("Order(uint256 salt,address maker,address signer,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint256 side,uint256 expiration,uint256 nonce,uint256 feeRateBps,uint256 signatureType)"),
	)
)
