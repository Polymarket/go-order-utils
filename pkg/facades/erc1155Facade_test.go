package facades

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/contracts/erc1155"
	"github.com/polymarket/go-order-utils/pkg/contracts/limitOrder"

	"github.com/stretchr/testify/assert"
)

func getNewERC1155FacadeImpl(t *testing.T) ERC1155Facade {
	limitOrderProtocolAbi, err := abi.JSON(strings.NewReader(limitOrder.LimitOrderABI))
	assert.NotNil(t, limitOrderProtocolAbi)
	assert.Nil(t, err)

	erc1155Abi, err := abi.JSON(strings.NewReader(erc1155.Erc1155ABI))
	assert.NotNil(t, erc1155Abi)
	assert.Nil(t, err)

	erc1155Facade, err := NewERC1155FacadeImpl()
	assert.NotNil(t, erc1155Facade)
	assert.Nil(t, err)
	assert.Equal(t, erc1155Facade.methodsLimitOrder, limitOrderProtocolAbi.Methods)
	assert.Equal(t, erc1155Facade.methodsERC1155, erc1155Abi.Methods)

	return erc1155Facade
}

func TestERC1155TransferFrom(t *testing.T) {
	erc1155Facade := getNewERC1155FacadeImpl(t)
	assert.NotNil(t, erc1155Facade)

	r, err := erc1155Facade.TransferFrom(
		common.Address{},
		common.Address{},
		common.Address{},
		&big.Int{},
		&big.Int{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestERC1155BalanceOf(t *testing.T) {
	erc1155Facade := getNewERC1155FacadeImpl(t)
	assert.NotNil(t, erc1155Facade)

	r, err := erc1155Facade.BalanceOf(common.Address{}, &big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}
