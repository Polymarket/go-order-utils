package eip712

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestBuildEIP712DomainSeparator(t *testing.T) {
	expectedAmoy := common.HexToHash("0x029ca49b31e01a0230787b3673a6977ee0dbb02f78e68d2ee41d330893e17594")
	name := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80002)
	address := common.HexToAddress("0x0000000000000000000000000000000000000000")

	actual, err := BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expectedAmoy.String(), actual.String())

	expectedPolygon := common.HexToHash("0ba60c2a4504ef46ef2d139f27cd131dbc8b9643d0565a54bb14de0e31cd5600")
	chainId = big.NewInt(137)

	actual, err = BuildEIP712DomainSeparator(name, version, chainId, address)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NotEqual(t, expectedAmoy.String(), actual.String())
	assert.Equal(t, expectedPolygon.String(), actual.String())
}

func TestBuildEIP712DomainSeparatorNoContract(t *testing.T) {
	// Calculated in foundry
	expectedAmoy := common.HexToHash("0xf231a704e0942ea4f4f4a68a19fa01e198ddae4595d821697230159c60e562b0")
	chainId := big.NewInt(80002)

	name := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))

	actual, err := BuildEIP712DomainSeparatorNoContract(name, version, chainId)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expectedAmoy.String(), actual.String())

	// Calculated in foundry
	expectedPolygon := common.HexToHash("aee1d7dd93bb10f6c6a59417017905bc5dbec7ddbd71475cd19d8a95845e632d")
	chainId = big.NewInt(137)

	actual, err = BuildEIP712DomainSeparatorNoContract(name, version, chainId)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual)
	assert.NotEqual(t, expectedAmoy.String(), actual.String())
	assert.Equal(t, expectedPolygon.String(), actual.String())
}

func TestHashTypedDataV4(t *testing.T) {
	name := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
	version := crypto.Keccak256Hash([]byte("1"))
	chainId := big.NewInt(80002)
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

	expectedTypedDataHash := "0x5b9fd6432c2890e6293b3880add5e3d40ead280b31f7cd70f06b43f8d37dcc13"
	assert.Equal(t, expectedTypedDataHash, dataHashBytes.String())
}
