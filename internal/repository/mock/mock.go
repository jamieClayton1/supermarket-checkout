// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/repository.go -destination=./internal/repository/repository_test.go
//

// Package mock_repository is a generated GoMock package.
package mock

import (
	reflect "reflect"
	entity "supermarket-checkout/internal/entity"

	gomock "go.uber.org/mock/gomock"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// FetchItem mocks base method.
func (m *MockItemRepository) FetchItem(sku string) (*entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchItem", sku)
	ret0, _ := ret[0].(*entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchItem indicates an expected call of FetchItem.
func (mr *MockItemRepositoryMockRecorder) FetchItem(sku any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchItem", reflect.TypeOf((*MockItemRepository)(nil).FetchItem), sku)
}

// MockBasketRepository is a mock of BasketRepository interface.
type MockBasketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBasketRepositoryMockRecorder
}

// MockBasketRepositoryMockRecorder is the mock recorder for MockBasketRepository.
type MockBasketRepositoryMockRecorder struct {
	mock *MockBasketRepository
}

// NewMockBasketRepository creates a new mock instance.
func NewMockBasketRepository(ctrl *gomock.Controller) *MockBasketRepository {
	mock := &MockBasketRepository{ctrl: ctrl}
	mock.recorder = &MockBasketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBasketRepository) EXPECT() *MockBasketRepositoryMockRecorder {
	return m.recorder
}

// FetchBasket mocks base method.
func (m *MockBasketRepository) FetchBasket(basketId string) (*entity.Basket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBasket", basketId)
	ret0, _ := ret[0].(*entity.Basket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBasket indicates an expected call of FetchBasket.
func (mr *MockBasketRepositoryMockRecorder) FetchBasket(basketId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBasket", reflect.TypeOf((*MockBasketRepository)(nil).FetchBasket), basketId)
}

// PutBasketItem mocks base method.
func (m *MockBasketRepository) PutBasketItem(item *entity.Item, basketId *string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBasketItem", item, basketId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBasketItem indicates an expected call of PutBasketItem.
func (mr *MockBasketRepositoryMockRecorder) PutBasketItem(item, basketId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBasketItem", reflect.TypeOf((*MockBasketRepository)(nil).PutBasketItem), item, basketId)
}
