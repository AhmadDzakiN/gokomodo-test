// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/product.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/repository/product.go -destination=internal/app/mocks/repository/product.go
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

// MockIProductRepository is a mock of IProductRepository interface.
type MockIProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProductRepositoryMockRecorder
}

// MockIProductRepositoryMockRecorder is the mock recorder for MockIProductRepository.
type MockIProductRepositoryMockRecorder struct {
	mock *MockIProductRepository
}

// NewMockIProductRepository creates a new mock instance.
func NewMockIProductRepository(ctrl *gomock.Controller) *MockIProductRepository {
	mock := &MockIProductRepository{ctrl: ctrl}
	mock.recorder = &MockIProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductRepository) EXPECT() *MockIProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProductRepository) Create(ctx *gin.Context, product *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIProductRepositoryMockRecorder) Create(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProductRepository)(nil).Create), ctx, product)
}

// GetByID mocks base method.
func (m *MockIProductRepository) GetByID(ctx *gin.Context, id uint64) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIProductRepositoryMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIProductRepository)(nil).GetByID), ctx, id)
}

// GetList mocks base method.
func (m *MockIProductRepository) GetList(ctx *gin.Context, params payloads.GetProductListParams) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, params)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockIProductRepositoryMockRecorder) GetList(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockIProductRepository)(nil).GetList), ctx, params)
}