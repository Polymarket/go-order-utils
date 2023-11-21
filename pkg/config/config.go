package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange        common.Address
	FeeModule       common.Address
	Collateral      common.Address
	Conditional     common.Address
	NegRiskAdapter  common.Address
	NegRiskExchange common.Address
}

func GetContracts(chainId int64) (*Contracts, error) {
	var (
		MUMBAI_CONTRACTS = &Contracts{
			Exchange:        common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:       common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:      common.HexToAddress("0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961"),
			Conditional:     common.HexToAddress("0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43"),
			NegRiskAdapter:  common.HexToAddress("0xf16a3BdFFB7B882E3236243E901f6c5953E2EE0d"), // TODO: This is not the final address
			NegRiskExchange: common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3"), // TODO: This is not the final address
		}

		MATIC_CONTRACTS = &Contracts{
			Exchange:        common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
			FeeModule:       common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113"),
			Collateral:      common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
			Conditional:     common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
			NegRiskAdapter:  common.HexToAddress("0xf16a3BdFFB7B882E3236243E901f6c5953E2EE0d"), // TODO: This is not the final address
			NegRiskExchange: common.HexToAddress("0x8599536bECcD7e24b13060d78A54e2237Ed57eF3"), // TODO: This is not the final address
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
