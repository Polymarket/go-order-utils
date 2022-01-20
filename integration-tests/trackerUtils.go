//nolint:deadcode
package main

import (
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/eip712"
	"github.com/polymarket/go-order-utils/pkg/sign"
)

type ApiKeyCreds struct {
	key        string
	secret     string
	passphrase string
}

type L2HeaderArgs struct {
	method      string
	requestPath string
	body        string
}

const (
	CLOB_STRUCT_NAME = "ClobAuth"
	CLOB_DOMAIN_NAME = "ClobAuthDomain"
	CLOB_VERSION     = "1"
	CLOB_CHAIN_ID    = 1
	MESSAGE_TO_SIGN  = "This message attests that I control the given wallet"
)

func createL1Headers(privateKey *ecdsa.PrivateKey) (map[string]string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	sig, err := buildClobEip712Signature(privateKey, timestamp)
	if err != nil {
		return nil, err
	}

	address, err := getPublicAddress(privateKey)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"POLY_ADDRESS":   address,
		"POLY_SIGNATURE": "0x" + hex.EncodeToString(sig),
		"POLY_TIMESTAMP": timestamp,
	}, nil
}

func createL2Headers(privateKey *ecdsa.PrivateKey, creds *ApiKeyCreds, l2HeadersArgs *L2HeaderArgs) (map[string]string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	address, err := getPublicAddress(privateKey)
	if err != nil {
		return nil, err
	}

	sig, err := buildPolyHmacSignature(
		creds.secret,
		timestamp,
		l2HeadersArgs.method,
		l2HeadersArgs.requestPath,
		l2HeadersArgs.body,
	)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"POLY_ADDRESS":    address,
		"POLY_SIGNATURE":  sig,
		"POLY_TIMESTAMP":  timestamp,
		"POLY_API_KEY":    creds.key,
		"POLY_PASSPHRASE": creds.passphrase,
	}, nil
}

func buildPolyHmacSignature(base64Secret, timestamp, method, requestPath, body string) (string, error) {
	message := timestamp + method + requestPath
	if body != "" {
		message += body
	}

	secret, err := base64.URLEncoding.DecodeString(base64Secret)
	if err != nil {
		return "", err
	}

	sig, err := createHmac(secret, []byte(message))
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(sig), "+", "-"), "/", "_"), nil
}

// Create a SHA256 HMAC signature
func createHmac(key, data []byte) ([]byte, error) {
	// Create a SHA256 HMAC, with the key
	h := hmac.New(sha256.New, key)
	_, err := h.Write(data)

	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func buildClobEip712Signature(privateKey *ecdsa.PrivateKey, timestamp string) ([]byte, error) {
	addressS, err := getPublicAddress(privateKey)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(addressS)

	// Construct the domain hash
	// Note that arbitrary length value types(array, string, byte) MUST be hashed
	types := []abi.Type{
		eip712.Bytes32,
		eip712.Bytes32,
		eip712.Bytes32,
		eip712.Uint256,
	}
	values := []interface{}{
		crypto.Keccak256Hash(
			[]byte("EIP712Domain(string name,string version,uint256 chainId)"),
		),
		crypto.Keccak256Hash([]byte(CLOB_DOMAIN_NAME)),
		crypto.Keccak256Hash([]byte(CLOB_VERSION)),
		big.NewInt(CLOB_CHAIN_ID),
	}
	encodedDomainSeparator, _ := eip712.Encode(types, values)
	domSeparator := crypto.Keccak256(encodedDomainSeparator)
	var domainSep32Bytes [32]byte
	copy(domainSep32Bytes[:], domSeparator[:])

	// Construct the struct type hash
	clobAuthTypeHash := crypto.Keccak256Hash(
		[]byte("ClobAuth(address address,string timestamp,string message)"),
	)

	t := []abi.Type{
		eip712.Bytes32,
		eip712.Address,
		eip712.Bytes32,
		eip712.Bytes32,
	}
	v := []interface{}{
		clobAuthTypeHash,
		address,
		crypto.Keccak256Hash(
			[]byte(timestamp),
		),
		crypto.Keccak256Hash(
			[]byte(MESSAGE_TO_SIGN),
		),
	}

	var structHash32Bytes [32]byte
	encoded, err := eip712.Encode(t, v)
	if err != nil {
		copy(structHash32Bytes[:], []byte(""))
		return structHash32Bytes[:], err
	}

	structHash := eip712.HashTypedDataV4(domainSep32Bytes, crypto.Keccak256Hash(encoded))
	copy(structHash32Bytes[:], structHash)

	signerImpl := sign.NewSignerImpl()
	signature, err := signerImpl.BuildSignature(privateKey, structHash32Bytes)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func getPublicAddress(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error parsing public key")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}
