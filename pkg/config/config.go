package config

import "fmt"

type Exchange struct {
	Address string
	Name    string
	Version string
}
type ClobContracts struct {
	Exchange   *Exchange
	Executor   string
	Collateral string
}

var (
	KOVAN_CONTRACTS = &ClobContracts{
		Exchange: &Exchange{
			Address: "0xE7819d9745e64c14541732ca07CC3898670b7650",
			Name:    "1inch Limit Order Protocol",
			Version: "1",
		},
		Executor:   "0x382E8f6a8404eB11aaFd9A5a0B11aa5A24e0830B",
		Collateral: "0xe22da380ee6B445bb8273C81944ADEB6E8450422",
	}

	MUMBAI_CONTRACTS = &ClobContracts{
		Exchange: &Exchange{
			Address: "",
			Name:    "1inch Limit Order Protocol",
			Version: "1",
		},
		Executor:   "",
		Collateral: "",
	}

	MATIC_CONTRACTS = &ClobContracts{
		Exchange: &Exchange{
			Address: "",
			Name:    "1inch Limit Order Protocol",
			Version: "1",
		},
		Executor:   "",
		Collateral: "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174",
	}
)

func GetContracts(chainId int) (*ClobContracts, error) {
	switch chainId {
	case 42:
		return KOVAN_CONTRACTS, nil
	case 137:
		return MATIC_CONTRACTS, nil
	case 80001:
		return MUMBAI_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
