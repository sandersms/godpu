/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/common/v1/gen/go"

	mock "github.com/stretchr/testify/mock"
)

// InvClient is an autogenerated mock type for the InvClient type
type InvClient struct {
	mock.Mock
}

type InvClient_Expecter struct {
	mock *mock.Mock
}

func (_m *InvClient) EXPECT() *InvClient_Expecter {
	return &InvClient_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx
func (_m *InvClient) Get(ctx context.Context) (*_go.Inventory, error) {
	ret := _m.Called(ctx)

	var r0 *_go.Inventory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*_go.Inventory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *_go.Inventory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Inventory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvClient_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type InvClient_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
func (_e *InvClient_Expecter) Get(ctx interface{}) *InvClient_Get_Call {
	return &InvClient_Get_Call{Call: _e.mock.On("Get", ctx)}
}

func (_c *InvClient_Get_Call) Run(run func(ctx context.Context)) *InvClient_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *InvClient_Get_Call) Return(_a0 *_go.Inventory, _a1 error) *InvClient_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InvClient_Get_Call) RunAndReturn(run func(context.Context) (*_go.Inventory, error)) *InvClient_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewInvClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvClient creates a new instance of InvClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvClient(t mockConstructorTestingTNewInvClient) *InvClient {
	mock := &InvClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}