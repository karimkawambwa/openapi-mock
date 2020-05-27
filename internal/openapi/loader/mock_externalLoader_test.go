// Code generated by mockery v1.0.0. DO NOT EDIT.

package loader

import (
	openapi3 "github.com/getkin/kin-openapi/openapi3"
	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// mockExternalLoader is an autogenerated mock type for the externalLoader type
type mockExternalLoader struct {
	mock.Mock
}

// LoadSwaggerFromFile provides a mock function with given fields: path
func (_m *mockExternalLoader) LoadSwaggerFromFile(path string) (*openapi3.Swagger, error) {
	ret := _m.Called(path)

	var r0 *openapi3.Swagger
	if rf, ok := ret.Get(0).(func(string) *openapi3.Swagger); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi3.Swagger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSwaggerFromURI provides a mock function with given fields: location
func (_m *mockExternalLoader) LoadSwaggerFromURI(location *url.URL) (*openapi3.Swagger, error) {
	ret := _m.Called(location)

	var r0 *openapi3.Swagger
	if rf, ok := ret.Get(0).(func(*url.URL) *openapi3.Swagger); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi3.Swagger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*url.URL) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
