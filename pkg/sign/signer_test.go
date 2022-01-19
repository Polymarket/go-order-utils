package sign

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
)

func TestSigner(t *testing.T) {
	signerImpl := NewSignerImpl()
	assert.NotNil(t, signerImpl)

	signerData := &signer.TypedData{
		Types: signer.Types{
			"Test": []signer.Type{
				{Name: "timestamp", Type: "string"},
			},
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "version", Type: "string"},
			},
		},
		PrimaryType: "Test",
		Domain: signer.TypedDataDomain{
			Name:    "Test",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(int64(1)),
		},
		Message: signer.TypedDataMessage{
			"timestamp": "1",
		},
	}

	hash, err := signerImpl.BuildHash(signerData)
	assert.NotNil(t, hash)
	assert.Nil(t, err)

	privateKey, err := crypto.GenerateKey()
	assert.NotNil(t, privateKey)
	assert.Nil(t, err)

	signature, err := signerImpl.BuildSignature(privateKey, hash)
	assert.NotNil(t, signature)
	assert.Nil(t, err)

	match, err := signerImpl.ValidateSignature(&privateKey.PublicKey, hash, signature)
	assert.NotNil(t, match)
	assert.True(t, match)
	assert.Nil(t, err)
}
