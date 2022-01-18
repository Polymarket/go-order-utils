package sign

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	signer "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signer interface {
	BuildHash(typedData *signer.TypedData) (common.Hash, error)
	BuildSignature(privateKey *ecdsa.PrivateKey, hash common.Hash) ([]byte, error)
	ValidateSignature(publicKey *ecdsa.PublicKey, hash common.Hash, signature []byte) (bool, error)
}
type SignerImpl struct{}

func NewSignerImpl() *SignerImpl {
	return &SignerImpl{}
}

func (o *SignerImpl) BuildHash(typedData *signer.TypedData) (common.Hash, error) {
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return [common.HashLength]byte{}, err
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return [common.HashLength]byte{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)

	return hash, nil
}

func (o *SignerImpl) BuildSignature(privateKey *ecdsa.PrivateKey, hash common.Hash) ([]byte, error) {
	return crypto.Sign(hash.Bytes(), privateKey)
}

func (o *SignerImpl) ValidateSignature(publicKey *ecdsa.PublicKey, hash common.Hash, signature []byte) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false, err
	}

	matches := bytes.Equal(sigPublicKey, crypto.FromECDSAPub(publicKey))
	return matches, nil
}
