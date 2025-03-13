// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, key
func (_m *Client) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Client_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Client_Expecter) Delete(ctx interface{}, key interface{}) *Client_Delete_Call {
	return &Client_Delete_Call{Call: _e.mock.On("Delete", ctx, key)}
}

func (_c *Client_Delete_Call) Run(run func(ctx context.Context, key string)) *Client_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_Delete_Call) Return(_a0 error) *Client_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Delete_Call) RunAndReturn(run func(context.Context, string) error) *Client_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePattern provides a mock function with given fields: ctx, prefix
func (_m *Client) DeletePattern(ctx context.Context, prefix string) error {
	ret := _m.Called(ctx, prefix)

	if len(ret) == 0 {
		panic("no return value specified for DeletePattern")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_DeletePattern_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePattern'
type Client_DeletePattern_Call struct {
	*mock.Call
}

// DeletePattern is a helper method to define mock.On call
//   - ctx context.Context
//   - prefix string
func (_e *Client_Expecter) DeletePattern(ctx interface{}, prefix interface{}) *Client_DeletePattern_Call {
	return &Client_DeletePattern_Call{Call: _e.mock.On("DeletePattern", ctx, prefix)}
}

func (_c *Client_DeletePattern_Call) Run(run func(ctx context.Context, prefix string)) *Client_DeletePattern_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_DeletePattern_Call) Return(_a0 error) *Client_DeletePattern_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_DeletePattern_Call) RunAndReturn(run func(context.Context, string) error) *Client_DeletePattern_Call {
	_c.Call.Return(run)
	return _c
}

// FlushAll provides a mock function with given fields: ctx
func (_m *Client) FlushAll(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FlushAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_FlushAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FlushAll'
type Client_FlushAll_Call struct {
	*mock.Call
}

// FlushAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) FlushAll(ctx interface{}) *Client_FlushAll_Call {
	return &Client_FlushAll_Call{Call: _e.mock.On("FlushAll", ctx)}
}

func (_c *Client_FlushAll_Call) Run(run func(ctx context.Context)) *Client_FlushAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_FlushAll_Call) Return(_a0 error) *Client_FlushAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_FlushAll_Call) RunAndReturn(run func(context.Context) error) *Client_FlushAll_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *Client) Get(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Client_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Client_Expecter) Get(ctx interface{}, key interface{}) *Client_Get_Call {
	return &Client_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *Client_Get_Call) Run(run func(ctx context.Context, key string)) *Client_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_Get_Call) Return(_a0 string, _a1 error) *Client_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Get_Call) RunAndReturn(run func(context.Context, string) (string, error)) *Client_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value
func (_m *Client) Set(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Client_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
func (_e *Client_Expecter) Set(ctx interface{}, key interface{}, value interface{}) *Client_Set_Call {
	return &Client_Set_Call{Call: _e.mock.On("Set", ctx, key, value)}
}

func (_c *Client_Set_Call) Run(run func(ctx context.Context, key string, value interface{})) *Client_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *Client_Set_Call) Return(_a0 error) *Client_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Set_Call) RunAndReturn(run func(context.Context, string, interface{}) error) *Client_Set_Call {
	_c.Call.Return(run)
	return _c
}

// SetWithTTL provides a mock function with given fields: ctx, key, value, ttl
func (_m *Client) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	if len(ret) == 0 {
		panic("no return value specified for SetWithTTL")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_SetWithTTL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetWithTTL'
type Client_SetWithTTL_Call struct {
	*mock.Call
}

// SetWithTTL is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - ttl time.Duration
func (_e *Client_Expecter) SetWithTTL(ctx interface{}, key interface{}, value interface{}, ttl interface{}) *Client_SetWithTTL_Call {
	return &Client_SetWithTTL_Call{Call: _e.mock.On("SetWithTTL", ctx, key, value, ttl)}
}

func (_c *Client_SetWithTTL_Call) Run(run func(ctx context.Context, key string, value interface{}, ttl time.Duration)) *Client_SetWithTTL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *Client_SetWithTTL_Call) Return(_a0 error) *Client_SetWithTTL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_SetWithTTL_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) error) *Client_SetWithTTL_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
