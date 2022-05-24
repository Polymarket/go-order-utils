package facades

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/contracts/limitOrder"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/stretchr/testify/assert"
)

func getNewLimitOrderProtocolFacadeImpl(t *testing.T) LimitOrderProtocolFacade {
	limitOrderProtocolAbi, err := abi.JSON(strings.NewReader(limitOrder.LimitOrderABI))
	assert.NotNil(t, limitOrderProtocolAbi)
	assert.Nil(t, err)

	limitOrderProtocolFacade, err := NewLimitOrderProtocolFacadeImpl()
	assert.NotNil(t, limitOrderProtocolFacade)
	assert.Nil(t, err)
	assert.Equal(t, limitOrderProtocolFacade.methods, limitOrderProtocolAbi.Methods)

	return limitOrderProtocolFacade
}

func getLimitOrder() model.LimitOrder {
	return model.LimitOrder{
		Salt:           &big.Int{},
		MakerAsset:     common.Address{},
		TakerAsset:     common.Address{},
		MakerAssetData: []byte{},
		TakerAssetData: []byte{},
		GetMakerAmount: []byte{},
		GetTakerAmount: []byte{},
		Predicate:      []byte{},
		Signer:         common.Address{},
		SigType:        &big.Int{},
	}
}

func TestGetMakerAmount(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.GetMakerAmount(
		&big.Int{}, &big.Int{}, &big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestGetTakerAmount(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.GetTakerAmount(
		&big.Int{}, &big.Int{}, &big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestFillOrder(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.FillOrder(getLimitOrder(), []byte{},
		&big.Int{}, &big.Int{}, &big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestCancelOrder(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.CancelOrder(getLimitOrder())
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestNonce(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Nonce(common.Address{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestAdvanceNonce(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.AdvanceNonce(uint8(1))
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestIncreaseNonce(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.IncreaseNonce()
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestAnd(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.And(
		[]common.Address{{}},
		[][]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestOr(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Or(
		[]common.Address{{}},
		[][]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestEq(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Eq(
		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestLt(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Lt(
		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestGt(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Gt(

		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestTimestampBelow(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.TimestampBelow(&big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestNonceEquals(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.NonceEquals(
		common.Address{},
		&big.Int{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestRemaining(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.Remaining([32]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestCheckPredicate(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.CheckPredicate(getLimitOrder())
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestRemainingsRaw(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.RemainingsRaw([][32]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestSimulateCalls(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.SimulateCalls(
		[]common.Address{{}},
		[][]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestDomainSeparator(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.DomainSeparator()
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestBatchFillOrders(t *testing.T) {
	limitOrderProtocolFacade := getNewLimitOrderProtocolFacadeImpl(t)
	assert.NotNil(t, limitOrderProtocolFacade)

	r, err := limitOrderProtocolFacade.BatchFillOrders(
		[]model.LimitOrder{getLimitOrder()}, [][]byte{},
		[]*big.Int{{}},
		[]*big.Int{{}},
		[]*big.Int{{}},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}
