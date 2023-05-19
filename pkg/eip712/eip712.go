package eip712

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator[:]), string(crypto.Keccak256Hash(encoded).Bytes())))
	return crypto.Keccak256Hash(rawData), nil
}

// Verifies an EIP712 signature
func ValidateSignature(signer common.Address, hashedData common.Hash, signature []byte) (bool, error) {
	sigCopy := make([]byte, len(signature))
	copy(sigCopy, signature)

	if len(sigCopy) != 65 {
		return false, secp256k1.ErrInvalidSignatureLen
	}

	if sigCopy[64] != 0 && sigCopy[64] != 1 { // in case of ledger signing v might already be 0 or 1
		sigCopy[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow paper
	}

	sigPublicKey, err := crypto.Ecrecover(hashedData.Bytes(), sigCopy)
	if err != nil {
		return false, err
	}

	recoveredPublicKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return false, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*recoveredPublicKey)
	return bytes.Equal(signer.Bytes(), recoveredAddress.Bytes()), nil
}
