// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	product "e-wallet/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductData is an autogenerated mock type for the DataInterface type
type ProductData struct {
	mock.Mock
}

// CountProduct provides a mock function with given fields:
func (_m *ProductData) CountProduct() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CountProduct")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountProductByUserId provides a mock function with given fields: id
func (_m *ProductData) CountProductByUserId(id uint) (int, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for CountProductByUserId")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (int, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: input
func (_m *ProductData) Delete(input uint) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *ProductData) Insert(input product.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAllProduct provides a mock function with given fields: offset, limit
func (_m *ProductData) SelectAllProduct(offset int, limit int) ([]product.Core, error) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for SelectAllProduct")
	}

	var r0 []product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]product.Core, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []product.Core); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectProductById provides a mock function with given fields: id
func (_m *ProductData) SelectProductById(id uint) (*product.Core, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for SelectProductById")
	}

	var r0 *product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*product.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *product.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectProductByUserId provides a mock function with given fields: id, offset, limit
func (_m *ProductData) SelectProductByUserId(id uint, offset int, limit int) ([]product.Core, error) {
	ret := _m.Called(id, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for SelectProductByUserId")
	}

	var r0 []product.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int) ([]product.Core, error)); ok {
		return rf(id, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int) []product.Core); ok {
		r0 = rf(id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(id, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, input
func (_m *ProductData) Update(id uint, input product.Core) error {
	ret := _m.Called(id, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, product.Core) error); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductData creates a new instance of ProductData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductData(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductData {
	mock := &ProductData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
