// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

type Engine_Expecter struct {
	mock *mock.Mock
}

func (_m *Engine) EXPECT() *Engine_Expecter {
	return &Engine_Expecter{mock: &_m.Mock}
}

// ChargeLevel provides a mock function with given fields:
func (_m *Engine) ChargeLevel() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Engine_ChargeLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChargeLevel'
type Engine_ChargeLevel_Call struct {
	*mock.Call
}

// ChargeLevel is a helper method to define mock.On call
func (_e *Engine_Expecter) ChargeLevel() *Engine_ChargeLevel_Call {
	return &Engine_ChargeLevel_Call{Call: _e.mock.On("ChargeLevel")}
}

func (_c *Engine_ChargeLevel_Call) Run(run func()) *Engine_ChargeLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Engine_ChargeLevel_Call) Return(_a0 int) *Engine_ChargeLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEngine interface {
	mock.TestingT
	Cleanup(func())
}

// NewEngine creates a new instance of Engine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEngine(t mockConstructorTestingTNewEngine) *Engine {
	mock := &Engine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
