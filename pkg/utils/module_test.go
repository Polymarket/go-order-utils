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
	contract, err := GetVerifyingContractAddress(big.NewInt(80001), model.FeeModule)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E").Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(80001), model.NegRiskModule)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3").Hex(), contract.Hex())

	// polygon mainnet
	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.FeeModule)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E").Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(137), model.NegRiskModule)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3").Hex(), contract.Hex())

	// wrong network
	_, err = GetVerifyingContractAddress(big.NewInt(1), model.FeeModule)
	assert.Error(t, err)

	_, err = GetVerifyingContractAddress(big.NewInt(1), model.NegRiskModule)
	assert.Error(t, err)
}
