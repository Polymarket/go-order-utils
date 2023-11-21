// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	model "github.com/polymarket/go-order-utils/pkg/model"
)

// ExchangeOrderBuilder is an autogenerated mock type for the ExchangeOrderBuilder type
type ExchangeOrderBuilder struct {
	mock.Mock
}

// BuildOrder provides a mock function with given fields: orderData
func (_m *ExchangeOrderBuilder) BuildOrder(orderData *model.OrderData) (*model.Order, error) {
	ret := _m.Called(orderData)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.OrderData) (*model.Order, error)); ok {
		return rf(orderData)
	}
	if rf, ok := ret.Get(0).(func(*model.OrderData) *model.Order); ok {
		r0 = rf(orderData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.OrderData) error); ok {
		r1 = rf(orderData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildOrderHash provides a mock function with given fields: order, module
func (_m *ExchangeOrderBuilder) BuildOrderHash(order *model.Order, module int) (common.Hash, error) {
	ret := _m.Called(order, module)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Order, int) (common.Hash, error)); ok {
		return rf(order, module)
	}
	if rf, ok := ret.Get(0).(func(*model.Order, int) common.Hash); ok {
		r0 = rf(order, module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Order, int) error); ok {
		r1 = rf(order, module)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildOrderSignature provides a mock function with given fields: privateKey, orderHash
func (_m *ExchangeOrderBuilder) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash common.Hash) ([]byte, error) {
	ret := _m.Called(privateKey, orderHash)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, common.Hash) ([]byte, error)); ok {
		return rf(privateKey, orderHash)
	}
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, common.Hash) []byte); ok {
		r0 = rf(privateKey, orderHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*ecdsa.PrivateKey, common.Hash) error); ok {
		r1 = rf(privateKey, orderHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildSignedOrder provides a mock function with given fields: privateKey, orderData, module
func (_m *ExchangeOrderBuilder) BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderData, module int) (*model.SignedOrder, error) {
	ret := _m.Called(privateKey, orderData, module)

	var r0 *model.SignedOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, *model.OrderData, int) (*model.SignedOrder, error)); ok {
		return rf(privateKey, orderData, module)
	}
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, *model.OrderData, int) *model.SignedOrder); ok {
		r0 = rf(privateKey, orderData, module)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SignedOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(*ecdsa.PrivateKey, *model.OrderData, int) error); ok {
		r1 = rf(privateKey, orderData, module)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExchangeOrderBuilder creates a new instance of ExchangeOrderBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExchangeOrderBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExchangeOrderBuilder {
	mock := &ExchangeOrderBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
