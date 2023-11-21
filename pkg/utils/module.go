package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/config"
	"github.com/polymarket/go-order-utils/pkg/model"
)

func GetVerifyingContractAddress(chainId *big.Int, module model.CTFModule) (common.Address, error) {
	contracts, err := config.GetContracts(chainId.Int64())
	if err != nil {
		return common.Address{}, err
	}

	switch module {
	case model.FeeModule:
		return contracts.Exchange, nil
	case model.NegRiskModule:
		return contracts.NegRiskExchange, nil
	}

	return common.Address{}, fmt.Errorf("invalid module")
}
