// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/room_repo.go
//
// Generated by this command:
//
//	mockgen -source=internal/domain/repository/room_repo.go -destination=internal/mock/domain/repository/mock_room_repository.go -package=repository /internal/domain/repository/room_repo.go RoomRepository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	model "github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomRepository is a mock of RoomRepository interface.
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository.
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance.
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockRoomRepository) CreateRoom(ctx context.Context, room *model.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomRepositoryMockRecorder) CreateRoom(ctx, room any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomRepository)(nil).CreateRoom), ctx, room)
}

// GetRoom mocks base method.
func (m *MockRoomRepository) GetRoom(ctx context.Context, id model.RoomID) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, id)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomRepositoryMockRecorder) GetRoom(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomRepository)(nil).GetRoom), ctx, id)
}

// ListRoom mocks base method.
func (m *MockRoomRepository) ListRoom(ctx context.Context) ([]*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoom", ctx)
	ret0, _ := ret[0].([]*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoom indicates an expected call of ListRoom.
func (mr *MockRoomRepositoryMockRecorder) ListRoom(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoom", reflect.TypeOf((*MockRoomRepository)(nil).ListRoom), ctx)
}

// UpdateRoom mocks base method.
func (m *MockRoomRepository) UpdateRoom(ctx context.Context, room *model.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoom indicates an expected call of UpdateRoom.
func (mr *MockRoomRepositoryMockRecorder) UpdateRoom(ctx, room any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoom", reflect.TypeOf((*MockRoomRepository)(nil).UpdateRoom), ctx, room)
}
