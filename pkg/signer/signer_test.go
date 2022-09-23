package signer

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func getPublicAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error parsing public key")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func TestSigner(t *testing.T) {
	hash := common.BytesToHash(crypto.Keccak256([]byte("hashed string")))
	assert.NotNil(t, hash)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := Sign(privateKey, hash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	signer, err := getPublicAddress(privateKey)
	assert.NotNil(t, signer)
	assert.Nil(t, err)

	match, err := ValidateSignature(signer, hash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
