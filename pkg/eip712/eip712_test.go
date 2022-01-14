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
		expected := common.Hex2Bytes("9063f5eedf29e049017dad8bf71ee9328905cb3cf6f0439ffe6a53bf5e8fcfcf")

		nameHash := crypto.Keccak256Hash([]byte("1inch Limit Order Protocol"))
		versionHash := crypto.Keccak256Hash([]byte("1"))
		chainID := big.NewInt(42)
		address := common.HexToAddress("0xC563F26Bdf0cFC8bEa1aa41A000E7F54B8569c8d")
		actual := BuildEIP712DomainSeparator(nameHash, versionHash, chainID, address)
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
		encoded, err := Encode(types, values)
		Expect(err).To(BeNil())

		expectedTypedDataHash := "0xc782787af573f8144babf20881d4f0188d1a22bbeccd2f42365f42a141d0a7f9"
		dataHashBytes := HashTypedDataV4(mockDomainSep32Bytes, crypto.Keccak256Hash(encoded))
		typedDataHash := "0x" + strings.ToLower(hex.EncodeToString(dataHashBytes[:]))

		Expect(typedDataHash).Should(Equal(expectedTypedDataHash))
	})

	It("should correctly verfiy a signature given a hash", func() {
		// Signed message generated the following values

		address := common.HexToAddress("0x798F0fC66080bD1dF04728617cFD3756927D4495")
		structHash := common.Hex2Bytes("ec8822f36e01c9707e826eebddb80c06f277f73f22bd1e8aa6635eb086a5efc1")
		signature := common.Hex2Bytes("f4ad869399a8d786b3da2fe5554f671b9ed04d71043c7168c7b794bc97b8eadd12578cf6778de425f24abb864cfc3fb0849b024081428aa2ae5539a80684b9241c")

		var structHash32Bytes [32]byte
		copy(structHash32Bytes[:], structHash)

		ok := VerifySignature(address, structHash32Bytes, signature)
		Expect(ok).To(BeTrue())
	})
})
