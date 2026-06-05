package ports

import (
	"context"
	"errors"
	"market/internal/core/domain"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrFailedToLoadOrder = errors.New("failed to load order")
	ErrFailedToSaveOrder = errors.New("failed to save order")
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderById(ctx context.Context, id int64) (*domain.Order, error)
	GetOrderByUserId(uctx context.Context, serId int64) ([]domain.Order, error)
	AddOrderItems(ctx context.Context, orderId int64, items []domain.OrderItem) error
	UpdateOrderStatus(ctx context.Context, orderId int64, status string) error
}
