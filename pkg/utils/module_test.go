package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGetVerifyingContractAddress(t *testing.T) {
	// amoy
	contract, err := GetVerifyingContractAddress(big.NewInt(80002), model.CTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40").Hex(), contract.Hex())

	contract, err = GetVerifyingContractAddress(big.NewInt(80002), model.NegRiskCTFExchange)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a").Hex(), contract.Hex())

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
