/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/common/v1/gen/go"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// InventorySvcClient is an autogenerated mock type for the InventorySvcClient type
type InventorySvcClient struct {
	mock.Mock
}

type InventorySvcClient_Expecter struct {
	mock *mock.Mock
}

func (_m *InventorySvcClient) EXPECT() *InventorySvcClient_Expecter {
	return &InventorySvcClient_Expecter{mock: &_m.Mock}
}

// GetInventory provides a mock function with given fields: ctx, in, opts
func (_m *InventorySvcClient) GetInventory(ctx context.Context, in *_go.GetInventoryRequest, opts ...grpc.CallOption) (*_go.Inventory, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Inventory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetInventoryRequest, ...grpc.CallOption) (*_go.Inventory, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetInventoryRequest, ...grpc.CallOption) *_go.Inventory); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Inventory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.GetInventoryRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InventorySvcClient_GetInventory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInventory'
type InventorySvcClient_GetInventory_Call struct {
	*mock.Call
}

// GetInventory is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.GetInventoryRequest
//   - opts ...grpc.CallOption
func (_e *InventorySvcClient_Expecter) GetInventory(ctx interface{}, in interface{}, opts ...interface{}) *InventorySvcClient_GetInventory_Call {
	return &InventorySvcClient_GetInventory_Call{Call: _e.mock.On("GetInventory",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *InventorySvcClient_GetInventory_Call) Run(run func(ctx context.Context, in *_go.GetInventoryRequest, opts ...grpc.CallOption)) *InventorySvcClient_GetInventory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.GetInventoryRequest), variadicArgs...)
	})
	return _c
}

func (_c *InventorySvcClient_GetInventory_Call) Return(_a0 *_go.Inventory, _a1 error) *InventorySvcClient_GetInventory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InventorySvcClient_GetInventory_Call) RunAndReturn(run func(context.Context, *_go.GetInventoryRequest, ...grpc.CallOption) (*_go.Inventory, error)) *InventorySvcClient_GetInventory_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewInventorySvcClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewInventorySvcClient creates a new instance of InventorySvcClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInventorySvcClient(t mockConstructorTestingTNewInventorySvcClient) *InventorySvcClient {
	mock := &InventorySvcClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}