// Code generated by mockery v2.45.0. DO NOT EDIT.

package entities

import (
	context "context"

	entities "github.com/ariefsn/intrans/entities"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepositoryMock is an autogenerated mock type for the AccountRepository type
type AccountRepositoryMock struct {
	mock.Mock
}

type AccountRepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepositoryMock) EXPECT() *AccountRepositoryMock_Expecter {
	return &AccountRepositoryMock_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, input
func (_m *AccountRepositoryMock) Create(ctx context.Context, input entities.AccountCreatePayload) (*entities.AccountModel, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entities.AccountModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.AccountCreatePayload) (*entities.AccountModel, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.AccountCreatePayload) *entities.AccountModel); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.AccountModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.AccountCreatePayload) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepositoryMock_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountRepositoryMock_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - input entities.AccountCreatePayload
func (_e *AccountRepositoryMock_Expecter) Create(ctx interface{}, input interface{}) *AccountRepositoryMock_Create_Call {
	return &AccountRepositoryMock_Create_Call{Call: _e.mock.On("Create", ctx, input)}
}

func (_c *AccountRepositoryMock_Create_Call) Run(run func(ctx context.Context, input entities.AccountCreatePayload)) *AccountRepositoryMock_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.AccountCreatePayload))
	})
	return _c
}

func (_c *AccountRepositoryMock_Create_Call) Return(_a0 *entities.AccountModel, _a1 error) *AccountRepositoryMock_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepositoryMock_Create_Call) RunAndReturn(run func(context.Context, entities.AccountCreatePayload) (*entities.AccountModel, error)) *AccountRepositoryMock_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AccountRepositoryMock) GetByID(ctx context.Context, id int64) (*entities.AccountModel, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entities.AccountModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entities.AccountModel, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entities.AccountModel); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.AccountModel)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepositoryMock_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type AccountRepositoryMock_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *AccountRepositoryMock_Expecter) GetByID(ctx interface{}, id interface{}) *AccountRepositoryMock_GetByID_Call {
	return &AccountRepositoryMock_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id)}
}

func (_c *AccountRepositoryMock_GetByID_Call) Run(run func(ctx context.Context, id int64)) *AccountRepositoryMock_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *AccountRepositoryMock_GetByID_Call) Return(_a0 *entities.AccountModel, _a1 error) *AccountRepositoryMock_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepositoryMock_GetByID_Call) RunAndReturn(run func(context.Context, int64) (*entities.AccountModel, error)) *AccountRepositoryMock_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountRepositoryMock creates a new instance of AccountRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepositoryMock {
	mock := &AccountRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
