// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Blank is an autogenerated mock type for the Blank type
type Blank struct {
	mock.Mock
}

// Create provides a mock function with given fields: x
func (_m *Blank) Create(x interface{}) error {
	ret := _m.Called(x)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(x)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlank creates a new instance of Blank. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlank(t testing.TB) *Blank {
	mock := &Blank{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
