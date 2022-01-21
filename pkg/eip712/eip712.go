package eip712

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func toTypedDataHash(domainSeparator [32]byte, structHash [32]byte) []byte {
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator[:]), string(structHash[:])))
	return crypto.Keccak256(rawData)
}

func BuildEIP712DomainSeparator(nameHash [32]byte, versionHash [32]byte, chainID *big.Int, address common.Address) ([]byte, error) {
	types := []abi.Type{
		Bytes32,
		Bytes32,
		Bytes32,
		Uint256,
		Address,
	}

	values := []interface{}{
		EIP_712_TYPEHASH,
		nameHash,
		versionHash,
		chainID,
		address,
	}
	encodedDomainSeparator, err := Encode(types, values)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(encodedDomainSeparator), nil
}

func HashTypedDataV4(domainSeparator [32]byte, structHash [32]byte) []byte {
	return toTypedDataHash(domainSeparator, structHash)
}
