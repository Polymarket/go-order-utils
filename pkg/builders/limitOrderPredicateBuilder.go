package builders

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/order-utils/pkg/facades"
)

type LimitOrderPredicateCallData = []byte

type LimitOrderPredicateBuilder interface {
	And(predicates ...LimitOrderPredicateCallData) (LimitOrderPredicateCallData, error)
	Or(predicates ...LimitOrderPredicateCallData) (LimitOrderPredicateCallData, error)
	Eq(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error)
	Lt(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error)
	Gt(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error)
	NonceEquals(makerAddress common.Address, makerNonce *big.Int) (LimitOrderPredicateCallData, error)
	TimestampBelow(timestamp *big.Int) (LimitOrderPredicateCallData, error)
}

type LimitOrderPredicateBuilderImpl struct {
	limitOrderProtocolFacade facades.LimitOrderProtocolFacade
}

func NewLimitOrderPredicateBuilderImpl(limitOrderProtocolFacade facades.LimitOrderProtocolFacade) *LimitOrderPredicateBuilderImpl {
	return &LimitOrderPredicateBuilderImpl{
		limitOrderProtocolFacade: limitOrderProtocolFacade,
	}
}

func (l *LimitOrderPredicateBuilderImpl) And(contractAddress common.Address, predicates ...LimitOrderPredicateCallData) (LimitOrderPredicateCallData, error) {
	addresses := []common.Address{}
	for range predicates {
		addresses = append(addresses, contractAddress)
	}
	return l.limitOrderProtocolFacade.And(addresses, predicates)
}

func (l *LimitOrderPredicateBuilderImpl) Or(contractAddress common.Address, predicates ...LimitOrderPredicateCallData) (LimitOrderPredicateCallData, error) {
	addresses := []common.Address{}
	for range predicates {
		addresses = append(addresses, contractAddress)
	}
	return l.limitOrderProtocolFacade.Or(addresses, predicates)
}

func (l *LimitOrderPredicateBuilderImpl) Eq(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error) {
	return l.limitOrderProtocolFacade.Eq(value, address, callData)
}

func (l *LimitOrderPredicateBuilderImpl) Lt(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error) {
	return l.limitOrderProtocolFacade.Lt(value, address, callData)
}

func (l *LimitOrderPredicateBuilderImpl) Gt(value *big.Int, address common.Address, callData []byte) (LimitOrderPredicateCallData, error) {
	return l.limitOrderProtocolFacade.Gt(value, address, callData)
}

func (l *LimitOrderPredicateBuilderImpl) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (LimitOrderPredicateCallData, error) {
	return l.limitOrderProtocolFacade.NonceEquals(makerAddress, makerNonce)
}

func (l *LimitOrderPredicateBuilderImpl) TimestampBelow(timestamp *big.Int) (LimitOrderPredicateCallData, error) {
	return l.limitOrderProtocolFacade.TimestampBelow(timestamp)
}
