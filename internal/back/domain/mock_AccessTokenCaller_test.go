// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAccessTokenCaller is an autogenerated mock type for the AccessTokenCaller type
type MockAccessTokenCaller struct {
	mock.Mock
}

type MockAccessTokenCaller_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccessTokenCaller) EXPECT() *MockAccessTokenCaller_Expecter {
	return &MockAccessTokenCaller_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, token
func (_m *MockAccessTokenCaller) Get(ctx context.Context, token string) (*AccessToken, error) {
	ret := _m.Called(ctx, token)

	var r0 *AccessToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*AccessToken, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *AccessToken); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessTokenCaller_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccessTokenCaller_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *MockAccessTokenCaller_Expecter) Get(ctx interface{}, token interface{}) *MockAccessTokenCaller_Get_Call {
	return &MockAccessTokenCaller_Get_Call{Call: _e.mock.On("Get", ctx, token)}
}

func (_c *MockAccessTokenCaller_Get_Call) Run(run func(ctx context.Context, token string)) *MockAccessTokenCaller_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccessTokenCaller_Get_Call) Return(_a0 *AccessToken, _a1 error) *MockAccessTokenCaller_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessTokenCaller_Get_Call) RunAndReturn(run func(context.Context, string) (*AccessToken, error)) *MockAccessTokenCaller_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAccessTokenCaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAccessTokenCaller creates a new instance of MockAccessTokenCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAccessTokenCaller(t mockConstructorTestingTNewMockAccessTokenCaller) *MockAccessTokenCaller {
	mock := &MockAccessTokenCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
