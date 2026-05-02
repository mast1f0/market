package ports

import "market/internal/core/domain"

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	GetOrderById(id int64) (*domain.Order, error)
	GetOrderByUserId(userId int64) ([]domain.Order, error)
	AddOrderItems(orderId int64, items []domain.OrderItem) error
	UpdateOrderStatus(orderId int64, status string) error
}
