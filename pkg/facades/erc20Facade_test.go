package facades

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/contracts/erc20"
	"github.com/stretchr/testify/assert"
)

func getNewERC20FacadeImpl(t *testing.T) ERC20Facade {
	erc20Abi, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	assert.NotNil(t, erc20Abi)
	assert.Nil(t, err)

	erc20Facade, err := NewERC20FacadeImpl()
	assert.NotNil(t, erc20Facade)
	assert.Nil(t, err)
	assert.Equal(t, erc20Facade.methods, erc20Abi.Methods)

	return erc20Facade
}

func TestERC20TransferFrom(t *testing.T) {
	erc20Facade := getNewERC20FacadeImpl(t)
	assert.NotNil(t, erc20Facade)

	r, err := erc20Facade.TransferFrom(
		common.Address{},
		common.Address{},
		&big.Int{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestERC20BalanceOf(t *testing.T) {
	erc20Facade := getNewERC20FacadeImpl(t)
	assert.NotNil(t, erc20Facade)

	r, err := erc20Facade.BalanceOf(common.Address{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}
