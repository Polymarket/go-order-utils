package sign

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
	BuildSignature(privateKey *ecdsa.PrivateKey, hash common.Hash) ([]byte, error)
	ValidateSignature(publicKey *ecdsa.PublicKey, hash common.Hash, signature []byte) (bool, error)
}
type SignerImpl struct{}

func NewSignerImpl() *SignerImpl {
	return &SignerImpl{}
}

func (o *SignerImpl) BuildSignature(privateKey *ecdsa.PrivateKey, hash common.Hash) ([]byte, error) {
	sign, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] += 27
	return sign, err
}

func (o *SignerImpl) ValidateSignature(publicKey *ecdsa.PublicKey, hash common.Hash, signature []byte) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false, err
	}

	matches := bytes.Equal(sigPublicKey, crypto.FromECDSAPub(publicKey))
	return matches, nil
}
