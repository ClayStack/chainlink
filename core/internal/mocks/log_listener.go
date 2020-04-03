// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LogListener is an autogenerated mock type for the LogListener type
type LogListener struct {
	mock.Mock
}

// HandleLog provides a mock function with given fields: log, err
func (_m *LogListener) HandleLog(log interface{}, err error) {
	_m.Called(log, err)
}

// OnConnect provides a mock function with given fields:
func (_m *LogListener) OnConnect() {
	_m.Called()
}

// OnDisconnect provides a mock function with given fields:
func (_m *LogListener) OnDisconnect() {
	_m.Called()
}
