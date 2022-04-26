// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// RequesterVariadic is an autogenerated mock type for the RequesterVariadic type
type RequesterVariadic struct {
	mock.Mock
}

// Get provides a mock function with given fields: values
func (_m *RequesterVariadic) Get(values ...string) bool {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LogMethodToFile provides a mock function with given fields: level
func (_m *RequesterVariadic) LogMethodToFile(level string) func(string, ...io.Writer) {
	ret := _m.Called(level)

	var r0 func(string, ...io.Writer)
	if rf, ok := ret.Get(0).(func(string) func(string, ...io.Writer)); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string, ...io.Writer))
		}
	}

	return r0
}

// LogMethodf provides a mock function with given fields: level
func (_m *RequesterVariadic) LogMethodf(level string) func(string, ...interface{}) {
	ret := _m.Called(level)

	var r0 func(string, ...interface{})
	if rf, ok := ret.Get(0).(func(string) func(string, ...interface{})); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string, ...interface{}))
		}
	}

	return r0
}

// MultiWriteToFile provides a mock function with given fields: filename, w
func (_m *RequesterVariadic) MultiWriteToFile(filename string, w ...io.Writer) string {
	_va := make([]interface{}, len(w))
	for _i := range w {
		_va[_i] = w[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, filename)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...io.Writer) string); ok {
		r0 = rf(filename, w...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OneInterface provides a mock function with given fields: a
func (_m *RequesterVariadic) OneInterface(a ...interface{}) bool {
	var _ca []interface{}
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(a...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Sprintf provides a mock function with given fields: format, a
func (_m *RequesterVariadic) Sprintf(format string, a ...interface{}) string {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewRequesterVariadic creates a new instance of RequesterVariadic. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterVariadic(t testing.TB) *RequesterVariadic {
	mock := &RequesterVariadic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
