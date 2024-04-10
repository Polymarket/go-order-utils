package config

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetContracts(t *testing.T) {
	var (
		amoy = &Contracts{
			Exchange:         common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40"),
			FeeModule:        common.HexToAddress("0x9A9faEf45C671cc57B3e117c5B3053075416490f"),
			NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
			NegRiskFeeModule: common.HexToAddress("0x78769d50be1763ed1ca0d5e878d93f05aabff29e"),
			NegRiskAdapter:   common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
			Collateral:       common.HexToAddress("0x9c4e1703476e875070ee25b56a58b008cfb8fa78"),
			Conditional:      common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
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

	c, err := GetContracts(80002)
	assert.NotNil(t, c)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(c.Exchange[:], amoy.Exchange[:]))
	assert.True(t, bytes.Equal(c.FeeModule[:], amoy.FeeModule[:]))
	assert.True(t, bytes.Equal(c.Collateral[:], amoy.Collateral[:]))
	assert.True(t, bytes.Equal(c.Conditional[:], amoy.Conditional[:]))
	assert.True(t, bytes.Equal(c.NegRiskFeeModule[:], amoy.NegRiskFeeModule[:]))
	assert.True(t, bytes.Equal(c.NegRiskExchange[:], amoy.NegRiskExchange[:]))
	assert.True(t, bytes.Equal(c.NegRiskAdapter[:], amoy.NegRiskAdapter[:]))

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
