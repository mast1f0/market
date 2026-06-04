package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/service/mock"
	"testing"
)

func TestUpdateOrderSuccess(t *testing.T) {
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			return &domain.Order{
				ID: 1,
			}, nil
		},
		UpdateOrder: func(orderId int64, status string) error {
			return nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	err := service.UpdateStatus(1, "pending", 1, "admin")

	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOrderFailedInvalidRole(t *testing.T) {
	executed := false
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			return &domain.Order{
				ID: 1,
			}, nil
		},
		UpdateOrder: func(orderId int64, status string) error {
			executed = true
			return nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	err := service.UpdateStatus(1, "pending", 1, "buyer")

	if !errors.Is(err, ErrForbidden) {
		t.Error(err)
	}

	if executed {
		t.Error("It shouldnt execute")
	}
}

func TestUpdateOrderFailedForeignProductSeller(t *testing.T) {
	executed := false
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			return &domain.Order{
				ID: 1,
			}, nil
		},
		UpdateOrder: func(orderId int64, status string) error {
			executed = true
			return nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	err := service.UpdateStatus(1, "pending", 1, "seller")

	if !errors.Is(err, ErrForbidden) {
		t.Error(err)
	}

	if executed {
		t.Error("It shouldnt execute")
	}
}

func TestCreateOrderSuccess(t *testing.T) {
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			return &domain.Order{}, nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	_, err := service.orderRepo.CreateOrder(&domain.Order{
		ID: 1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateOrderFailedInvalidID(t *testing.T) {
	executed := false
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			executed = true
			return &domain.Order{}, nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	_, err := service.orderRepo.CreateOrder(&domain.Order{
		ID: -1,
	})
	if !errors.Is(err, ErrInvalidOrderId) {
		t.Error(err)
	}
	if executed {
		t.Error("It shouldnt execute")
	}
}

func TestGetOrderByUserIdSuccess(t *testing.T) {
	orderRepo := &mock.OrderRepoMock{
		GetOrder: func(id int64) (*domain.Order, error) {
			return &domain.Order{}, nil
		},
	}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	_, err := service.GetOrderById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetOrderByUserIdFailed(t *testing.T) {
	orderRepo := &mock.OrderRepoMock{}
	cartRepo := &mock.CartRepoMock{}
	service := NewOrderService(orderRepo, cartRepo)
	_, err := service.GetOrderById(-1)
	if err == nil {
		t.Error("It should error")
	}
}
