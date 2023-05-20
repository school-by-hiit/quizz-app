// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthRepository is an autogenerated mock type for the AuthRepository type
type MockAuthRepository struct {
	mock.Mock
}

type MockAuthRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthRepository) EXPECT() *MockAuthRepository_Expecter {
	return &MockAuthRepository_Expecter{mock: &_m.Mock}
}

// AddRoleToUser provides a mock function with given fields: ctx, userId, role
func (_m *MockAuthRepository) AddRoleToUser(ctx context.Context, userId string, role Role) error {
	ret := _m.Called(ctx, userId, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, Role) error); ok {
		r0 = rf(ctx, userId, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_AddRoleToUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRoleToUser'
type MockAuthRepository_AddRoleToUser_Call struct {
	*mock.Call
}

// AddRoleToUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - role Role
func (_e *MockAuthRepository_Expecter) AddRoleToUser(ctx interface{}, userId interface{}, role interface{}) *MockAuthRepository_AddRoleToUser_Call {
	return &MockAuthRepository_AddRoleToUser_Call{Call: _e.mock.On("AddRoleToUser", ctx, userId, role)}
}

func (_c *MockAuthRepository_AddRoleToUser_Call) Run(run func(ctx context.Context, userId string, role Role)) *MockAuthRepository_AddRoleToUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(Role))
	})
	return _c
}

func (_c *MockAuthRepository_AddRoleToUser_Call) Return(_a0 error) *MockAuthRepository_AddRoleToUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_AddRoleToUser_Call) RunAndReturn(run func(context.Context, string, Role) error) *MockAuthRepository_AddRoleToUser_Call {
	_c.Call.Return(run)
	return _c
}

// CreateToken provides a mock function with given fields: ctx, token
func (_m *MockAuthRepository) CreateToken(ctx context.Context, token *AccessToken) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *AccessToken) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_CreateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateToken'
type MockAuthRepository_CreateToken_Call struct {
	*mock.Call
}

// CreateToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token *AccessToken
func (_e *MockAuthRepository_Expecter) CreateToken(ctx interface{}, token interface{}) *MockAuthRepository_CreateToken_Call {
	return &MockAuthRepository_CreateToken_Call{Call: _e.mock.On("CreateToken", ctx, token)}
}

func (_c *MockAuthRepository_CreateToken_Call) Run(run func(ctx context.Context, token *AccessToken)) *MockAuthRepository_CreateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*AccessToken))
	})
	return _c
}

func (_c *MockAuthRepository_CreateToken_Call) Return(_a0 error) *MockAuthRepository_CreateToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_CreateToken_Call) RunAndReturn(run func(context.Context, *AccessToken) error) *MockAuthRepository_CreateToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *MockAuthRepository) CreateUser(ctx context.Context, user *User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockAuthRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *User
func (_e *MockAuthRepository_Expecter) CreateUser(ctx interface{}, user interface{}) *MockAuthRepository_CreateUser_Call {
	return &MockAuthRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *MockAuthRepository_CreateUser_Call) Run(run func(ctx context.Context, user *User)) *MockAuthRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*User))
	})
	return _c
}

func (_c *MockAuthRepository_CreateUser_Call) Return(_a0 error) *MockAuthRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_CreateUser_Call) RunAndReturn(run func(context.Context, *User) error) *MockAuthRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// FindTokenByTokenStr provides a mock function with given fields: ctx, tokenStr
func (_m *MockAuthRepository) FindTokenByTokenStr(ctx context.Context, tokenStr string) (*AccessToken, error) {
	ret := _m.Called(ctx, tokenStr)

	var r0 *AccessToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*AccessToken, error)); ok {
		return rf(ctx, tokenStr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *AccessToken); ok {
		r0 = rf(ctx, tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthRepository_FindTokenByTokenStr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTokenByTokenStr'
type MockAuthRepository_FindTokenByTokenStr_Call struct {
	*mock.Call
}

// FindTokenByTokenStr is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenStr string
func (_e *MockAuthRepository_Expecter) FindTokenByTokenStr(ctx interface{}, tokenStr interface{}) *MockAuthRepository_FindTokenByTokenStr_Call {
	return &MockAuthRepository_FindTokenByTokenStr_Call{Call: _e.mock.On("FindTokenByTokenStr", ctx, tokenStr)}
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) Run(run func(ctx context.Context, tokenStr string)) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) Return(_a0 *AccessToken, _a1 error) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) RunAndReturn(run func(context.Context, string) (*AccessToken, error)) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserById provides a mock function with given fields: ctx, id
func (_m *MockAuthRepository) FindUserById(ctx context.Context, id string) (*User, error) {
	ret := _m.Called(ctx, id)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthRepository_FindUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserById'
type MockAuthRepository_FindUserById_Call struct {
	*mock.Call
}

// FindUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAuthRepository_Expecter) FindUserById(ctx interface{}, id interface{}) *MockAuthRepository_FindUserById_Call {
	return &MockAuthRepository_FindUserById_Call{Call: _e.mock.On("FindUserById", ctx, id)}
}

func (_c *MockAuthRepository_FindUserById_Call) Run(run func(ctx context.Context, id string)) *MockAuthRepository_FindUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthRepository_FindUserById_Call) Return(_a0 *User, _a1 error) *MockAuthRepository_FindUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthRepository_FindUserById_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *MockAuthRepository_FindUserById_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveAllRoleFromUser provides a mock function with given fields: ctx, userId
func (_m *MockAuthRepository) RemoveAllRoleFromUser(ctx context.Context, userId string) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_RemoveAllRoleFromUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveAllRoleFromUser'
type MockAuthRepository_RemoveAllRoleFromUser_Call struct {
	*mock.Call
}

// RemoveAllRoleFromUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
func (_e *MockAuthRepository_Expecter) RemoveAllRoleFromUser(ctx interface{}, userId interface{}) *MockAuthRepository_RemoveAllRoleFromUser_Call {
	return &MockAuthRepository_RemoveAllRoleFromUser_Call{Call: _e.mock.On("RemoveAllRoleFromUser", ctx, userId)}
}

func (_c *MockAuthRepository_RemoveAllRoleFromUser_Call) Run(run func(ctx context.Context, userId string)) *MockAuthRepository_RemoveAllRoleFromUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAuthRepository_RemoveAllRoleFromUser_Call) Return(_a0 error) *MockAuthRepository_RemoveAllRoleFromUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_RemoveAllRoleFromUser_Call) RunAndReturn(run func(context.Context, string) error) *MockAuthRepository_RemoveAllRoleFromUser_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAuthRepository creates a new instance of MockAuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAuthRepository(t mockConstructorTestingTNewMockAuthRepository) *MockAuthRepository {
	mock := &MockAuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}