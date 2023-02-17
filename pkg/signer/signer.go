package signer

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func Sign(privateKey *ecdsa.PrivateKey, hashedData common.Hash) ([]byte, error) {
	sign, err := crypto.Sign(hashedData.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] += 27
	return sign, err
}

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
