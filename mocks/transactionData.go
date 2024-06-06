// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	transaction "e-wallet/features/transaction"

	mock "github.com/stretchr/testify/mock"
)

// TransactionData is an autogenerated mock type for the DataInterface type
type TransactionData struct {
	mock.Mock
}

// CountByMerchantId provides a mock function with given fields: merchantId
func (_m *TransactionData) CountByMerchantId(merchantId uint) (int, error) {
	ret := _m.Called(merchantId)

	if len(ret) == 0 {
		panic("no return value specified for CountByMerchantId")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int, error)); ok {
		return rf(merchantId)
	}
	if rf, ok := ret.Get(0).(func(uint) int); ok {
		r0 = rf(merchantId)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(merchantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: input
func (_m *TransactionData) Insert(input transaction.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(transaction.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectTransactionById provides a mock function with given fields: id
func (_m *TransactionData) SelectTransactionById(id uint) (*transaction.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for SelectTransactionById")
	}

	var r0 *transaction.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*transaction.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *transaction.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectTransactionByMerchantId provides a mock function with given fields: id, offset, limit
func (_m *TransactionData) SelectTransactionByMerchantId(id uint, offset int, limit int) ([]transaction.Core, error) {
	ret := _m.Called(id, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for SelectTransactionByMerchantId")
	}

	var r0 []transaction.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int) ([]transaction.Core, error)); ok {
		return rf(id, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int) []transaction.Core); ok {
		r0 = rf(id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(id, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusProgress provides a mock function with given fields: id, input
func (_m *TransactionData) UpdateStatusProgress(id uint, input transaction.Core) error {
	ret := _m.Called(id, input)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusProgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, transaction.Core) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyPin provides a mock function with given fields: pin, idUser
func (_m *TransactionData) VerifyPin(pin string, idUser uint) error {
	ret := _m.Called(pin, idUser)

	if len(ret) == 0 {
		panic("no return value specified for VerifyPin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(pin, idUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTransactionData creates a new instance of TransactionData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionData(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionData {
	mock := &TransactionData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
