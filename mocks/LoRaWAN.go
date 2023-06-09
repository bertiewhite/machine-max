// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LoRaWAN is an autogenerated mock type for the LoRaWAN type
type LoRaWAN struct {
	mock.Mock
}

type LoRaWAN_Expecter struct {
	mock *mock.Mock
}

func (_m *LoRaWAN) EXPECT() *LoRaWAN_Expecter {
	return &LoRaWAN_Expecter{mock: &_m.Mock}
}

// RegisterDevEUI provides a mock function with given fields: _a0
func (_m *LoRaWAN) RegisterDevEUI(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoRaWAN_RegisterDevEUI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterDevEUI'
type LoRaWAN_RegisterDevEUI_Call struct {
	*mock.Call
}

// RegisterDevEUI is a helper method to define mock.On call
//   - _a0 string
func (_e *LoRaWAN_Expecter) RegisterDevEUI(_a0 interface{}) *LoRaWAN_RegisterDevEUI_Call {
	return &LoRaWAN_RegisterDevEUI_Call{Call: _e.mock.On("RegisterDevEUI", _a0)}
}

func (_c *LoRaWAN_RegisterDevEUI_Call) Run(run func(_a0 string)) *LoRaWAN_RegisterDevEUI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LoRaWAN_RegisterDevEUI_Call) Return(_a0 bool, _a1 error) *LoRaWAN_RegisterDevEUI_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewLoRaWAN interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoRaWAN creates a new instance of LoRaWAN. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoRaWAN(t mockConstructorTestingTNewLoRaWAN) *LoRaWAN {
	mock := &LoRaWAN{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
