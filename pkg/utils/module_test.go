package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGetVerifyingContractAddress(t *testing.T) {
	// mumbai
	contract, err := GetVerifyingContractAddress(big.NewInt(80001), model.CTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E").Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(80001), model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x87d1A0DdB4C63a6301916F02090A51a7241571e4").Hex(), contract.Hex())

	// polygon mainnet
	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.CTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E").Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a").Hex(), contract.Hex())

	// wrong network
	_, err = GetVerifyingContractAddress(big.NewInt(1), model.CTFExchange)
	assert.Error(t, err)

	_, err = GetVerifyingContractAddress(big.NewInt(1), model.NegRiskCTFExchange)
	assert.Error(t, err)
}
