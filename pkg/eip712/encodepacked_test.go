package eip712

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var _ = Describe("EncodePacked Tests", func() {
	It("Encode packed test", func() {
		// Test example from https://docs.soliditylang.org/en/v0.8.7/abi-spec.html#non-standard-packed-mode
		// int16(-1), bytes1(0x42), uint16(0x03), string("Hello, world!")
		want := common.Hex2Bytes("ffff42000348656c6c6f2c20776f726c6421")

		types := []abi.Type{
			Int16T,
			Bytes1T,
			Uint16,
			String,
		}
		values := []interface{}{
			int16(-1),
			[]byte{0x42},
			uint16(0x03),
			"Hello, world!",
		}
		encoded, err := EncodePacked(types, values)
		Expect(err).To(BeNil())
		Expect(encoded).Should(Equal(want))
	})

})

func TestTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EncodePacked tests")
}
