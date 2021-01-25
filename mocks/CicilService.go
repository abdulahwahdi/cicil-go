// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cicil "github.com/abdulahwahdi/cicil-go"
	mock "github.com/stretchr/testify/mock"
)

// CicilService is an autogenerated mock type for the CicilService type
type CicilService struct {
	mock.Mock
}

// GetCheckoutURL provides a mock function with given fields: param
func (_m *CicilService) GetCheckoutURL(param cicil.CheckoutRequest) (cicil.CheckoutResponse, error) {
	ret := _m.Called(param)

	var r0 cicil.CheckoutResponse
	if rf, ok := ret.Get(0).(func(cicil.CheckoutRequest) cicil.CheckoutResponse); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(cicil.CheckoutResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cicil.CheckoutRequest) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields:
func (_m *CicilService) GetToken() (string, string) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// SetCancelOrder provides a mock function with given fields: param
func (_m *CicilService) SetCancelOrder(param cicil.CancelOrderRequest) (cicil.CancelOrderResponse, error) {
	ret := _m.Called(param)

	var r0 cicil.CancelOrderResponse
	if rf, ok := ret.Get(0).(func(cicil.CancelOrderRequest) cicil.CancelOrderResponse); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(cicil.CancelOrderResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cicil.CancelOrderRequest) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: param
func (_m *CicilService) UpdateStatus(param cicil.UpdateStatusRequest) (cicil.UpdateStatusResponse, error) {
	ret := _m.Called(param)

	var r0 cicil.UpdateStatusResponse
	if rf, ok := ret.Get(0).(func(cicil.UpdateStatusRequest) cicil.UpdateStatusResponse); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(cicil.UpdateStatusResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cicil.UpdateStatusRequest) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}