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
			Exchange:        common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:       common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:      common.HexToAddress("0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961"),
			Conditional:     common.HexToAddress("0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43"),
			NegRiskAdapter:  common.HexToAddress("0xf16a3BdFFB7B882E3236243E901f6c5953E2EE0d"),
			NegRiskExchange: common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3"),
		}

		matic = &Contracts{
			Exchange:        common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:       common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:      common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
			Conditional:     common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
			NegRiskAdapter:  common.HexToAddress("0xf16a3BdFFB7B882E3236243E901f6c5953E2EE0d"),
			NegRiskExchange: common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3"),
		}
	)

	c, err := GetContracts(80001)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], mumbai.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], mumbai.FeeModule[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], mumbai.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], mumbai.Conditional[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], mumbai.NegRiskAdapter[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], mumbai.NegRiskExchange[:]))

	c, err = GetContracts(137)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], matic.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], matic.FeeModule[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], matic.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], matic.Conditional[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], mumbai.NegRiskAdapter[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], mumbai.NegRiskExchange[:]))

	c, err = GetContracts(100000)
	assert.Nil(t, c)
	assert.NotNil(t, err)
}
