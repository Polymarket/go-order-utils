package eip712

import (
	"encoding/hex"
	"math/big"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var _ = Describe("EIP712 Tests", func() {
	It("should build the domain separator correctly", func() {
		//Domain separator of Limit Order protocol deployed to kovan here: 0xC563F26Bdf0cFC8bEa1aa41A000E7F54B8569c8d
		expected := common.Hex2Bytes("other")

		nameHash := crypto.Keccak256Hash([]byte("Polymarket CTF Exchange"))
		versionHash := crypto.Keccak256Hash([]byte("1"))
		chainId := big.NewInt(42)
		address := common.HexToAddress("0x0000000000000000000000000000000000000000")
		actual, _ := BuildEIP712DomainSeparator(nameHash, versionHash, chainId, address)
		Expect(actual).Should(Equal(expected))
	})

	It("should correctly hash typed data", func() {
		mockObjTypeHash := crypto.Keccak256Hash([]byte("MockObj(string name, uint256 id)"))
		mockDomainSeparator := []byte("abcdef123")
		var mockDomainSep32Bytes [32]byte
		copy(mockDomainSeparator, mockDomainSep32Bytes[:])

		types := []abi.Type{
			Bytes32,
			String,
			Uint256,
		}
		values := []interface{}{
			mockObjTypeHash,
			"test",
			big.NewInt(1),
		}

		expectedTypedDataHash := "0xc782787af573f8144babf20881d4f0188d1a22bbeccd2f42365f42a141d0a7f9"
		dataHashBytes, err := HashTypedDataV4(mockDomainSep32Bytes, types, values)
		Expect(err).To(BeNil())
		typedDataHash := "0x" + strings.ToLower(hex.EncodeToString(dataHashBytes[:]))

		Expect(typedDataHash).Should(Equal(expectedTypedDataHash))
	})
})
