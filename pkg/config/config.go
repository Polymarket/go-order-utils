package config

import "fmt"

type Exchange struct {
	Address string
	Name    string
	Version string
}
type ClobContracts struct {
	Exchange    *Exchange
	Executor    string
	Collateral  string
	Conditional string
}

var (
	MUMBAI_CONTRACTS = &ClobContracts{
		Exchange: &Exchange{
			Address: "0x6D486b31b5c0f724828Aff07c88606b213B0D196",
			Name:    "Polymarket Limit Order Protocol",
			Version: "1",
		},
		Executor:    "0x6b0ab7A1E65ea6AE9072f6c45B4261ACDfB30827",
		Collateral:  "0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961",
		Conditional: "0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43",
	}

	MATIC_CONTRACTS = &ClobContracts{
		Exchange: &Exchange{
			Address: "0x6D486b31b5c0f724828Aff07c88606b213B0D196",
			Name:    "Polymarket Limit Order Protocol",
			Version: "1",
		},
		Executor:    "0x6b0ab7A1E65ea6AE9072f6c45B4261ACDfB30827",
		Collateral:  "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174",
		Conditional: "0x4D97DCd97eC945f40cF65F87097ACe5EA0476045",
	}
)

func GetContracts(chainId int) (*ClobContracts, error) {
	switch chainId {
	case 137:
		return MATIC_CONTRACTS, nil
	case 80001:
		return MUMBAI_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
