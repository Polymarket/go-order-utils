package eip712

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func BuildEIP712DomainSeparator(name, version common.Hash, chainId *big.Int, address common.Address) (common.Hash, error) {
	values := []interface{}{
		_EIP712_DOMAIN_HASH,
		name,
		version,
		chainId,
		address,
	}

	encodedDomainSeparator, err := Encode(_EIP712_DOMAIN, values)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(encodedDomainSeparator), nil
}

func HashTypedDataV4(domainSeparator common.Hash, args []abi.Type, values []interface{}) (common.Hash, error) {
	encoded, err := Encode(args, values)
	if err != nil {
		return common.Hash{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator[:]), string(encoded[:])))
	return crypto.Keccak256Hash(rawData), nil
}
