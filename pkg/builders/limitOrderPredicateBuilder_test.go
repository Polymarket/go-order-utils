package builders

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/facades"
	"github.com/stretchr/testify/assert"
)

func getNewLimitOrderPredicateBuilderImpl(t *testing.T) LimitOrderPredicateBuilder {
	limitOrderProtocolFacade, err := facades.NewLimitOrderProtocolFacadeImpl()
	assert.NotNil(t, limitOrderProtocolFacade)
	assert.Nil(t, err)

	limitOrderPredicateBuilder := NewLimitOrderPredicateBuilderImpl(limitOrderProtocolFacade)
	assert.NotNil(t, limitOrderPredicateBuilder)
	assert.Nil(t, err)
	assert.Equal(t, limitOrderPredicateBuilder.limitOrderProtocolFacade, limitOrderProtocolFacade)

	return limitOrderPredicateBuilder
}

func TestAnd(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.And(
		common.Address{},
		[]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestOr(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.Or(
		common.Address{},
		[]byte{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestEq(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.Eq(
		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestLt(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.Lt(
		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestGt(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.Gt(

		&big.Int{},
		common.Address{},
		[]byte{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestTimestampBelow(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.TimestampBelow(&big.Int{})
	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestNonceEquals(t *testing.T) {
	limitOrderPredicateBuilder := getNewLimitOrderPredicateBuilderImpl(t)
	assert.NotNil(t, limitOrderPredicateBuilder)

	r, err := limitOrderPredicateBuilder.NonceEquals(
		common.Address{},
		&big.Int{},
	)
	assert.NotNil(t, r)
	assert.Nil(t, err)
}
