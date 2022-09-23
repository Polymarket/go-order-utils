package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange    common.Address
	Collateral  common.Address
	Conditional common.Address
}

func GetContracts(chainId int64) (*Contracts, error) {
	var (
		MUMBAI_CONTRACTS = &Contracts{
			Exchange:    common.HexToAddress("0xfffd6f0dB1ec30A58884B23546B4F1bB333f818f"),
			Collateral:  common.HexToAddress("0x2E8DCfE708D44ae2e406a1c02DFE2Fa13012f961"),
			Conditional: common.HexToAddress("0x7D8610E9567d2a6C9FBf66a5A13E9Ba8bb120d43"),
		}

		MATIC_CONTRACTS = &Contracts{
			Exchange:    common.HexToAddress("0xfffd6f0dB1ec30A58884B23546B4F1bB333f818f"),
			Collateral:  common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
			Conditional: common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
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
