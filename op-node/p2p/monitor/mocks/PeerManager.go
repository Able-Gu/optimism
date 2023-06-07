// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	time "time"
)

// PeerManager is an autogenerated mock type for the PeerManager type
type PeerManager struct {
	mock.Mock
}

type PeerManager_Expecter struct {
	mock *mock.Mock
}

func (_m *PeerManager) EXPECT() *PeerManager_Expecter {
	return &PeerManager_Expecter{mock: &_m.Mock}
}

// BanPeer provides a mock function with given fields: _a0, _a1
func (_m *PeerManager) BanPeer(_a0 peer.ID, _a1 time.Time) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID, time.Time) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PeerManager_BanPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BanPeer'
type PeerManager_BanPeer_Call struct {
	*mock.Call
}

// BanPeer is a helper method to define mock.On call
//   - _a0 peer.ID
//   - _a1 time.Time
func (_e *PeerManager_Expecter) BanPeer(_a0 interface{}, _a1 interface{}) *PeerManager_BanPeer_Call {
	return &PeerManager_BanPeer_Call{Call: _e.mock.On("BanPeer", _a0, _a1)}
}

func (_c *PeerManager_BanPeer_Call) Run(run func(_a0 peer.ID, _a1 time.Time)) *PeerManager_BanPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID), args[1].(time.Time))
	})
	return _c
}

func (_c *PeerManager_BanPeer_Call) Return(_a0 error) *PeerManager_BanPeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PeerManager_BanPeer_Call) RunAndReturn(run func(peer.ID, time.Time) error) *PeerManager_BanPeer_Call {
	_c.Call.Return(run)
	return _c
}

// GetPeerScore provides a mock function with given fields: id
func (_m *PeerManager) GetPeerScore(id peer.ID) (float64, error) {
	ret := _m.Called(id)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID) (float64, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) float64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PeerManager_GetPeerScore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPeerScore'
type PeerManager_GetPeerScore_Call struct {
	*mock.Call
}

// GetPeerScore is a helper method to define mock.On call
//   - id peer.ID
func (_e *PeerManager_Expecter) GetPeerScore(id interface{}) *PeerManager_GetPeerScore_Call {
	return &PeerManager_GetPeerScore_Call{Call: _e.mock.On("GetPeerScore", id)}
}

func (_c *PeerManager_GetPeerScore_Call) Run(run func(id peer.ID)) *PeerManager_GetPeerScore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID))
	})
	return _c
}

func (_c *PeerManager_GetPeerScore_Call) Return(_a0 float64, _a1 error) *PeerManager_GetPeerScore_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PeerManager_GetPeerScore_Call) RunAndReturn(run func(peer.ID) (float64, error)) *PeerManager_GetPeerScore_Call {
	_c.Call.Return(run)
	return _c
}

// IsStatic provides a mock function with given fields: _a0
func (_m *PeerManager) IsStatic(_a0 peer.ID) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(peer.ID) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PeerManager_IsStatic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsStatic'
type PeerManager_IsStatic_Call struct {
	*mock.Call
}

// IsStatic is a helper method to define mock.On call
//   - _a0 peer.ID
func (_e *PeerManager_Expecter) IsStatic(_a0 interface{}) *PeerManager_IsStatic_Call {
	return &PeerManager_IsStatic_Call{Call: _e.mock.On("IsStatic", _a0)}
}

func (_c *PeerManager_IsStatic_Call) Run(run func(_a0 peer.ID)) *PeerManager_IsStatic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID))
	})
	return _c
}

func (_c *PeerManager_IsStatic_Call) Return(_a0 bool) *PeerManager_IsStatic_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PeerManager_IsStatic_Call) RunAndReturn(run func(peer.ID) bool) *PeerManager_IsStatic_Call {
	_c.Call.Return(run)
	return _c
}

// Peers provides a mock function with given fields:
func (_m *PeerManager) Peers() []peer.ID {
	ret := _m.Called()

	var r0 []peer.ID
	if rf, ok := ret.Get(0).(func() []peer.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]peer.ID)
		}
	}

	return r0
}

// PeerManager_Peers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Peers'
type PeerManager_Peers_Call struct {
	*mock.Call
}

// Peers is a helper method to define mock.On call
func (_e *PeerManager_Expecter) Peers() *PeerManager_Peers_Call {
	return &PeerManager_Peers_Call{Call: _e.mock.On("Peers")}
}

func (_c *PeerManager_Peers_Call) Run(run func()) *PeerManager_Peers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PeerManager_Peers_Call) Return(_a0 []peer.ID) *PeerManager_Peers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PeerManager_Peers_Call) RunAndReturn(run func() []peer.ID) *PeerManager_Peers_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPeerManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewPeerManager creates a new instance of PeerManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPeerManager(t mockConstructorTestingTNewPeerManager) *PeerManager {
	mock := &PeerManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}