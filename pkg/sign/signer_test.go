package sign

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSigner(t *testing.T) {
	signerImpl := NewSignerImpl()
	assert.NotNil(t, signerImpl)

	hash := common.BytesToHash(crypto.Keccak256([]byte("hashed string")))
	assert.NotNil(t, hash)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := signerImpl.BuildSignature(privateKey, hash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	signer, err := GetPublicAddress(privateKey)
	assert.NotNil(t, signer)
	assert.Nil(t, err)

	signature[64] -= 27 // Transform V from 27/28 to 0/1 according to the yellow paper
	match, err := signerImpl.ValidateSignature(signer, hash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
