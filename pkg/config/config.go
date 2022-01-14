package config

type ClobContracts struct {
	Exchange   string
	Executor   string
	Collateral string
}

var (
	KOVAN_CONTRACTS = ClobContracts{
		Exchange:   "0xE7819d9745e64c14541732ca07CC3898670b7650",
		Executor:   "0x382E8f6a8404eB11aaFd9A5a0B11aa5A24e0830B",
		Collateral: "0xe22da380ee6B445bb8273C81944ADEB6E8450422",
	}

	MUMBAI_CONTRACTS = ClobContracts{
		Exchange:   "", // TODO
		Executor:   "",
		Collateral: "",
	}

	MATIC_CONTRACTS = ClobContracts{
		Exchange:   "", // TODO
		Executor:   "",
		Collateral: "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174",
	}
)

func GetContracts(chainId int) ClobContracts {
	switch chainId {
	case 42:
		return KOVAN_CONTRACTS
	case 137:
		return MATIC_CONTRACTS
	case 80001:
		return MUMBAI_CONTRACTS
	default:
		panic("Invalid chain id")
	}
}
