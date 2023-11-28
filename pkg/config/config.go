package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange         common.Address
	FeeModule        common.Address
	NegRiskExchange  common.Address
	NegRiskFeeModule common.Address
	Collateral       common.Address
	Conditional      common.Address
}

func GetContracts(chainId int64) (*Contracts, error) {
	var (
		MUMBAI_CONTRACTS = &Contracts{
			Exchange:         common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:        common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
			NegRiskFeeModule: common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"),
			Collateral:       common.HexToAddress("0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961"),
			Conditional:      common.HexToAddress("0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43"),
		}

		MATIC_CONTRACTS = &Contracts{
			Exchange:         common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:        common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			NegRiskExchange:  common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
			NegRiskFeeModule: common.HexToAddress("0x78769D50Be1763ed1CA0D5E878D93f05aabff29e"),
			Collateral:       common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
			Conditional:      common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
		}
	)

	switch chainId {
	case 137:
		return MATIC_CONTRACTS, nil
	case 80001:
		return MUMBAI_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
