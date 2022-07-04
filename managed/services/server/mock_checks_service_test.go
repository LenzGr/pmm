// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockChecksService is an autogenerated mock type for the checksService type
type mockChecksService struct {
	mock.Mock
}

// CleanupAlerts provides a mock function with given fields:
func (_m *mockChecksService) CleanupAlerts() {
	_m.Called()
}

// CollectChecks provides a mock function with given fields: ctx
func (_m *mockChecksService) CollectChecks(ctx context.Context) {
	_m.Called(ctx)
}

// StartChecks provides a mock function with given fields: checkNames
func (_m *mockChecksService) StartChecks(checkNames []string) error {
	ret := _m.Called(checkNames)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(checkNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateIntervals provides a mock function with given fields: rare, standard, frequent
func (_m *mockChecksService) UpdateIntervals(rare time.Duration, standard time.Duration, frequent time.Duration) {
	_m.Called(rare, standard, frequent)
}