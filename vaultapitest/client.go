package vaultapitest

import mock "github.com/stretchr/testify/mock"
import time "time"
import vaultapi "github.com/shoenig/vaultapi"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AccessorCapabilities provides a mock function with given fields: path, accessor
func (_m *Client) AccessorCapabilities(path string, accessor string) ([]string, error) {
	ret := _m.Called(path, accessor)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(path, accessor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, accessor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateToken provides a mock function with given fields: opts
func (_m *Client) CreateToken(opts vaultapi.TokenOptions) (vaultapi.CreatedToken, error) {
	ret := _m.Called(opts)

	var r0 vaultapi.CreatedToken
	if rf, ok := ret.Get(0).(func(vaultapi.TokenOptions) vaultapi.CreatedToken); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(vaultapi.CreatedToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(vaultapi.TokenOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: path
func (_m *Client) Delete(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePolicy provides a mock function with given fields: name
func (_m *Client) DeletePolicy(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: path
func (_m *Client) Get(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPolicy provides a mock function with given fields: name
func (_m *Client) GetPolicy(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields:
func (_m *Client) Health() (vaultapi.Health, error) {
	ret := _m.Called()

	var r0 vaultapi.Health
	if rf, ok := ret.Get(0).(func() vaultapi.Health); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vaultapi.Health)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Keys provides a mock function with given fields: path
func (_m *Client) Keys(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// Leader provides a mock function with given fields:
func (_m *Client) Leader() (vaultapi.Leader, error) {
	ret := _m.Called()

	var r0 vaultapi.Leader
	if rf, ok := ret.Get(0).(func() vaultapi.Leader); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vaultapi.Leader)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMounts provides a mock function with given fields:
func (_m *Client) ListMounts() (vaultapi.Mounts, error) {
	ret := _m.Called()

	var r0 vaultapi.Mounts
	if rf, ok := ret.Get(0).(func() vaultapi.Mounts); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vaultapi.Mounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicies provides a mock function with given fields:
func (_m *Client) ListPolicies() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupLease provides a mock function with given fields: id
func (_m *Client) LookupLease(id string) (vaultapi.Lease, error) {
	ret := _m.Called(id)

	var r0 vaultapi.Lease
	if rf, ok := ret.Get(0).(func(string) vaultapi.Lease); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vaultapi.Lease)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupSelfToken provides a mock function with given fields:
func (_m *Client) LookupSelfToken() (vaultapi.LookedUpToken, error) {
	ret := _m.Called()

	var r0 vaultapi.LookedUpToken
	if rf, ok := ret.Get(0).(func() vaultapi.LookedUpToken); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vaultapi.LookedUpToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupToken provides a mock function with given fields: id
func (_m *Client) LookupToken(id string) (vaultapi.LookedUpToken, error) {
	ret := _m.Called(id)

	var r0 vaultapi.LookedUpToken
	if rf, ok := ret.Get(0).(func(string) vaultapi.LookedUpToken); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vaultapi.LookedUpToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: path, value
func (_m *Client) Put(path string, value string) error {
	ret := _m.Called(path, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(path, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenewToken provides a mock function with given fields: id, increment
func (_m *Client) RenewToken(id string, increment time.Duration) (vaultapi.RenewedToken, error) {
	ret := _m.Called(id, increment)

	var r0 vaultapi.RenewedToken
	if rf, ok := ret.Get(0).(func(string, time.Duration) vaultapi.RenewedToken); ok {
		r0 = rf(id, increment)
	} else {
		r0 = ret.Get(0).(vaultapi.RenewedToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(id, increment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SealStatus provides a mock function with given fields:
func (_m *Client) SealStatus() (vaultapi.SealStatus, error) {
	ret := _m.Called()

	var r0 vaultapi.SealStatus
	if rf, ok := ret.Get(0).(func() vaultapi.SealStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vaultapi.SealStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelfCapabilities provides a mock function with given fields: path
func (_m *Client) SelfCapabilities(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// SetPolicy provides a mock function with given fields: name, content
func (_m *Client) SetPolicy(name string, content string) error {
	ret := _m.Called(name, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StepDown provides a mock function with given fields:
func (_m *Client) StepDown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokenCapabilities provides a mock function with given fields: path, token
func (_m *Client) TokenCapabilities(path string, token string) ([]string, error) {
	ret := _m.Called(path, token)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(path, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
