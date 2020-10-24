// Code generated by mockery v1.0.0. DO NOT EDIT.

package slack

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// MockHTTPClient is an autogenerated mock type for the HTTPClient type
type MockHTTPClient struct {
	mock.Mock
}

// Do provides a mock function with given fields: req
func (_m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
