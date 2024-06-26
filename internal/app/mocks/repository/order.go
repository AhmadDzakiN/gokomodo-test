// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/order.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/repository/order.go -destination=internal/app/mocks/repository/order.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "gokomodo-assignment/internal/app/entity"
	payloads "gokomodo-assignment/internal/app/payloads"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockIOrderRepository is a mock of IOrderRepository interface.
type MockIOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderRepositoryMockRecorder
}

// MockIOrderRepositoryMockRecorder is the mock recorder for MockIOrderRepository.
type MockIOrderRepositoryMockRecorder struct {
	mock *MockIOrderRepository
}

// NewMockIOrderRepository creates a new mock instance.
func NewMockIOrderRepository(ctrl *gomock.Controller) *MockIOrderRepository {
	mock := &MockIOrderRepository{ctrl: ctrl}
	mock.recorder = &MockIOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrderRepository) EXPECT() *MockIOrderRepositoryMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockIOrderRepository) Accept(ctx *gin.Context, id uint64, sellerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", ctx, id, sellerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockIOrderRepositoryMockRecorder) Accept(ctx, id, sellerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockIOrderRepository)(nil).Accept), ctx, id, sellerID)
}

// Create mocks base method.
func (m *MockIOrderRepository) Create(ctx *gin.Context, order *entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIOrderRepositoryMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIOrderRepository)(nil).Create), ctx, order)
}

// GetByID mocks base method.
func (m *MockIOrderRepository) GetByID(ctx *gin.Context, id uint64) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIOrderRepositoryMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIOrderRepository)(nil).GetByID), ctx, id)
}

// GetList mocks base method.
func (m *MockIOrderRepository) GetList(ctx *gin.Context, params payloads.GetOrderListParams) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, params)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockIOrderRepositoryMockRecorder) GetList(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockIOrderRepository)(nil).GetList), ctx, params)
}
