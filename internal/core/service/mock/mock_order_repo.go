package mock

import "market/internal/core/domain"

type OrderRepoMock struct {
	GetOrder    func(id int64) (*domain.Order, error)
	UpdateOrder func(orderId int64, status string) error
}

func (o OrderRepoMock) CreateOrder(order *domain.Order) (*domain.Order, error) {
	return nil, nil
}

func (o OrderRepoMock) GetOrderById(id int64) (*domain.Order, error) {
	return o.GetOrder(id)
}

func (o OrderRepoMock) GetOrderByUserId(userId int64) ([]domain.Order, error) {
	return nil, nil
}

func (o OrderRepoMock) AddOrderItems(orderId int64, items []domain.OrderItem) error {
	return nil
}

func (o OrderRepoMock) UpdateOrderStatus(orderId int64, status string) error {
	return o.UpdateOrder(orderId, status)
}
