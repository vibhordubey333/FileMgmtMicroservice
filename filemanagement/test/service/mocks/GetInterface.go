// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	file_mgmt "FileMgmtMicroservice/restapi/operations/file_mgmt"

	middleware "github.com/go-openapi/runtime/middleware"

	mock "github.com/stretchr/testify/mock"
)

// GetInterface is an autogenerated mock type for the GetInterface type
type GetInterface struct {
	mock.Mock
}

// GetFile provides a mock function with given fields: params
func (_m *GetInterface) GetFile(params file_mgmt.GetFileParams) middleware.Responder {
	ret := _m.Called(params)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(file_mgmt.GetFileParams) middleware.Responder); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
