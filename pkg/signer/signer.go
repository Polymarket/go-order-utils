package signer

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Sign(privateKey *ecdsa.PrivateKey, typedData common.Hash) ([]byte, error) {
	sign, err := crypto.Sign(typedData.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] += 27
	return sign, err
}

func ValidateSignature(signer common.Address, typedData common.Hash, signature []byte) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(typedData.Bytes(), signature)
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
