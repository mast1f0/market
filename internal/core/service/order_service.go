package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type OrderService struct {
	orderRepo ports.OrderRepository
	cartRepo  ports.CartRepository
}

func NewOrderService(
	orderRepo ports.OrderRepository,
	cartRepo ports.CartRepository,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (s *OrderService) GetOrderById(orderId int64) (*domain.Order, error) {
	return s.orderRepo.GetOrderById(orderId)
}
func (s *OrderService) GetByUserId(userId int64) ([]domain.Order, error) {
	return s.orderRepo.GetOrderByUserId(userId)
}

func (s *OrderService) UpdateStatus(orderId int64, status string) error {
	return s.orderRepo.UpdateOrderStatus(orderId, status)
}

func (s *OrderService) CreateFromCart(userId int64) (*domain.Order, error) {
	cart, err := s.cartRepo.GetCartWithItems(userId)
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("no items in cart")
	}

	var totalPrice float64
	var orderedItems []domain.OrderItem

	for _, item := range cart.Items {
		price := item.PriceSnapshot

		if price == 0 && item.Product != nil {
			price = item.Product.Price
		}

		totalPrice += price * float64(item.Quantity)

		orderedItems = append(orderedItems, domain.OrderItem{
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			PriceSnapshot: price,
		})
	}

	order := &domain.Order{
		UserId:     userId,
		Status:     "pending",
		TotalPrice: totalPrice,
	}

	order, err = s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	err = s.orderRepo.AddOrderItems(order.Id, orderedItems)
	if err != nil {
		return nil, err
	}

	err = s.cartRepo.ClearCart(cart.ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
