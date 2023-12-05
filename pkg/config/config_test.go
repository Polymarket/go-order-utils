package config

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetContracts(t *testing.T) {
	var (
		mumbai = &Contracts{
			Exchange:         common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:        common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:       common.HexToAddress("0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961"),
			Conditional:      common.HexToAddress("0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43"),
			NegRiskFeeModule: common.HexToAddress("0x311FD94978961EeF4B903519B33B47B506aB6823"),
			NegRiskExchange:  common.HexToAddress("0x87d1A0DdB4C63a6301916F02090A51a7241571e4"),
			NegRiskAdapter:   common.HexToAddress("0x9A6930BB811fe3dFE1c35e4502134B38EC54399C"),
		}

		matic = &Contracts{
			Exchange:         common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:        common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:       common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
			Conditional:      common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
			NegRiskFeeModule: common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"),
			NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
			NegRiskAdapter:   common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		}
	)

	c, err := GetContracts(80001)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], mumbai.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], mumbai.FeeModule[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], mumbai.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], mumbai.Conditional[:]))
	assert.True(t, bytes.Equal(c.NegRiskFeeModule[:], mumbai.NegRiskFeeModule[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], mumbai.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], mumbai.NegRiskAdapter[:]))

	c, err = GetContracts(137)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], matic.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], matic.FeeModule[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], matic.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], matic.Conditional[:]))
	assert.True(t, bytes.Equal(c.NegRiskFeeModule[:], matic.NegRiskFeeModule[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], matic.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], matic.NegRiskAdapter[:]))

	c, err = GetContracts(100000)
	assert.Nil(t, c)
	assert.NotNil(t, err)
}
