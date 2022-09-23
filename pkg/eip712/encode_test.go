package eip712

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	// Test example from https://docs.soliditylang.org/en/v0.8.7/abi-spec.html#non-standard-packed-mode
	// Solidity abi.encode(int16(-1), bytes1(0x42), uint16(0x03), string("Hello, world!"))
	want := common.Hex2Bytes("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff420000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000d48656c6c6f2c20776f726c642100000000000000000000000000000000000000")

	types := []abi.Type{
		Int16T,
		Bytes1T,
		Uint16,
		String,
	}
	values := []interface{}{
		int16(-1),
		[1]byte{0x42},
		uint16(0x03),
		"Hello, world!",
	}

	encoded, err := Encode(types, values)
	assert.NoError(t, err)
	assert.NotEmpty(t, encoded)
	assert.Equal(t, want, encoded)
}
