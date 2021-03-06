// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import errors "github.com/onedaycat/errors"
import mock "github.com/stretchr/testify/mock"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// Confirm provides a mock function with given fields: name, val, confirmName, confirmValue, msg
func (_m *Validator) Confirm(name string, val interface{}, confirmName string, confirmValue interface{}, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, confirmName, confirmValue)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// Email provides a mock function with given fields: name, val, msg
func (_m *Validator) Email(name string, val string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// EmptyOrEmail provides a mock function with given fields: name, val, msg
func (_m *Validator) EmptyOrEmail(name string, val string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// EqualInt provides a mock function with given fields: name, val, equal, msg
func (_m *Validator) EqualInt(name string, val int, equal int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, equal)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// EqualString provides a mock function with given fields: name, val, size, msg
func (_m *Validator) EqualString(name string, val string, size int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, size)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// GetError provides a mock function with given fields:
func (_m *Validator) GetError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasError provides a mock function with given fields:
func (_m *Validator) HasError() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ISO8601DataTime provides a mock function with given fields: name, val, msg
func (_m *Validator) ISO8601DataTime(name string, val string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// InFloat64 provides a mock function with given fields: name, val, list, msg
func (_m *Validator) InFloat64(name string, val float64, list []float64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, list)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// InInt provides a mock function with given fields: name, val, list, msg
func (_m *Validator) InInt(name string, val int, list []int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, list)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// InInt64 provides a mock function with given fields: name, val, list, msg
func (_m *Validator) InInt64(name string, val int64, list []int64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, list)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// InString provides a mock function with given fields: name, val, list, msg
func (_m *Validator) InString(name string, val string, list []string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, list)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// IsValid provides a mock function with given fields:
func (_m *Validator) IsValid() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MaxFloat64 provides a mock function with given fields: name, val, max, msg
func (_m *Validator) MaxFloat64(name string, val float64, max float64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MaxInt provides a mock function with given fields: name, val, max, msg
func (_m *Validator) MaxInt(name string, val int, max int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MaxInt64 provides a mock function with given fields: name, val, max, msg
func (_m *Validator) MaxInt64(name string, val int64, max int64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MaxString provides a mock function with given fields: name, val, max, msg
func (_m *Validator) MaxString(name string, val string, max int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MinFloat64 provides a mock function with given fields: name, val, min, msg
func (_m *Validator) MinFloat64(name string, val float64, min float64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MinInt provides a mock function with given fields: name, val, min, msg
func (_m *Validator) MinInt(name string, val int, min int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MinInt64 provides a mock function with given fields: name, val, min, msg
func (_m *Validator) MinInt64(name string, val int64, min int64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// MinString provides a mock function with given fields: name, val, min, msg
func (_m *Validator) MinString(name string, val string, min int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// NotEmptyBool provides a mock function with given fields: name, val, msg
func (_m *Validator) NotEmptyBool(name string, val bool, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// NotEmptyFloat64 provides a mock function with given fields: name, val, msg
func (_m *Validator) NotEmptyFloat64(name string, val float64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// NotEmptyInt provides a mock function with given fields: name, val, msg
func (_m *Validator) NotEmptyInt(name string, val int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// NotEmptyInt64 provides a mock function with given fields: name, val, msg
func (_m *Validator) NotEmptyInt64(name string, val int64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// NotEmptyString provides a mock function with given fields: name, val, msg
func (_m *Validator) NotEmptyString(name string, val string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// RangeFloat64 provides a mock function with given fields: name, val, min, max, msg
func (_m *Validator) RangeFloat64(name string, val float64, min float64, max float64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// RangeInt provides a mock function with given fields: name, val, min, max, msg
func (_m *Validator) RangeInt(name string, val int, min int, max int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// RangeInt64 provides a mock function with given fields: name, val, min, max, msg
func (_m *Validator) RangeInt64(name string, val int64, min int64, max int64, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// RangeString provides a mock function with given fields: name, val, min, max, msg
func (_m *Validator) RangeString(name string, val string, min int, max int, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val, min, max)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// Required provides a mock function with given fields: name, val, msg
func (_m *Validator) Required(name string, val bool, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, name, val)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// SetError provides a mock function with given fields: msg
func (_m *Validator) SetError(msg string) {
	_m.Called(msg)
}

// Wrap provides a mock function with given fields: _a0
func (_m *Validator) Wrap(_a0 errors.Error) errors.Error {
	ret := _m.Called(_a0)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(errors.Error) errors.Error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}
