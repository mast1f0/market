package ports

import (
	"errors"
	"market/internal/core/domain"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrFailedToLoadOrder = errors.New("failed to load order")
	ErrFailedToSaveOrder = errors.New("failed to save order")
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	GetOrderById(id int64) (*domain.Order, error)
	GetOrderByUserId(userId int64) ([]domain.Order, error)
	AddOrderItems(orderId int64, items []domain.OrderItem) error
	UpdateOrderStatus(orderId int64, status string) error
}
