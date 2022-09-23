package eip712

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestBuildEIP712DomainSeparator(t *testing.T) {
	expectedMumbai := common.Hex2Bytes("66b157d2cc69ba84f05ab5f6ae3a7438d209df5c167c89848968e301da841e1c")
	name := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80001)
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")

	actual, err := BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expectedMumbai, actual[:])

	expectedPolygon := common.Hex2Bytes("0ba60c2a4504ef46ef2d139f27cd131dbc8b9643d0565a54bb14de0e31cd5600")
	chainId = big.NewInt(137)

	actual, err = BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NotEqual(t, expectedMumbai, actual[:])
	assert.Equal(t, expectedPolygon, actual[:])
}

func TestHashTypedDataV4(t *testing.T) {
	name := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80001)
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")

	domainSeparator, err := BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, domainSeparator)

	types := []abi.Type{
		Bytes32,
		String,
		Uint256,
	}
	values := []interface{}{
		crypto.Keccak256Hash([]byte("MockObj(string name, uint256 id)")),
		"test",
		big.NewInt(1),
	}

	dataHashBytes, err := HashTypedDataV4(domainSeparator, types, values)
	assert.NoError(t, err)
	assert.NotEmpty(t, dataHashBytes)

	expectedTypedDataHash := "caaea30a004febfda6fd75f0ccf024a806839d6a6f9e1fd79424787ab078a7e3"
	expectedTypedDataHashBytes := common.Hex2Bytes(expectedTypedDataHash)

	assert.Equal(t, expectedTypedDataHashBytes, dataHashBytes[:])

	typedDataHash := hex.EncodeToString(dataHashBytes[:])
	assert.Equal(t, expectedTypedDataHash, typedDataHash)
}
