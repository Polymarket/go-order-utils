package eip712

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_EIP712_DOMAIN = []abi.Type{
		Bytes32, // typehash
		Bytes32, // name
		Bytes32, // version
		Uint256, // chainId
		Address, // verifyingContract
	}
)

var (
	_EIP712_DOMAIN_HASH = crypto.Keccak256Hash(
		[]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
	)
)
