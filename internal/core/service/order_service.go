package service

import (
	"context"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrFailedToLoadOrder = errors.New("failed to load order")
	ErrInvalidOrderId    = errors.New("invalid order id")
	ErrFailedToSaveOrder = errors.New("failed to save order")
	ErrEmptyCart         = errors.New("cart is empty")
)

func mapOrderRepoError(err error) error {
	switch {
	case errors.Is(err, ports.ErrOrderNotFound):
		return ErrOrderNotFound
	case errors.Is(err, ports.ErrFailedToLoadOrder):
		return ErrFailedToLoadOrder
	case errors.Is(err, ports.ErrFailedToSaveOrder):
		return ErrFailedToSaveOrder
	default:
		return ErrFailedToSaveOrder
	}
}

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

func (s *OrderService) GetOrderById(ctx context.Context, orderId int64) (*domain.Order, error) {
	if orderId < 0 {
		return nil, ErrInvalidOrderId
	}
	order, err := s.orderRepo.GetOrderById(ctx, orderId)
	if err != nil {
		if errors.Is(err, ports.ErrOrderNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, ErrFailedToLoadOrder
	}
	return order, nil
}
func (s *OrderService) GetByUserId(ctx context.Context, userId int64) ([]domain.Order, error) {
	if userId < 0 {
		return nil, ErrInvalidUserID
	}
	orders, err := s.orderRepo.GetOrderByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, ports.ErrOrderNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, ErrFailedToLoadOrder
	}
	return orders, nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, orderId int64, status string, userID int64, role string) error {
	if userID < 0 {
		return ErrInvalidUserID
	}
	if orderId < 0 {
		return ErrInvalidOrderId
	}
	order, err := s.GetOrderById(ctx, orderId)
	if err != nil {
		if errors.Is(err, ports.ErrOrderNotFound) {
			return ErrOrderNotFound
		}
		return ErrFailedToLoadOrder
	}
	if order.UserID != userID && role != "admin" || role == "buyer" {
		return ErrForbidden
	}
	err = s.orderRepo.UpdateOrderStatus(ctx, orderId, status)
	if err != nil {
		if errors.Is(err, ports.ErrOrderNotFound) {
			return ErrOrderNotFound
		}
		return ErrFailedToSaveOrder
	}
	return nil
}

func (s *OrderService) CreateFromCart(ctx context.Context, userId int64) (*domain.Order, error) {
	if userId < 0 {
		return nil, ErrInvalidUserID
	}
	cart, err := s.cartRepo.GetCartWithItems(ctx, userId)
	if err != nil {
		return nil, mapCartRepoError(err)
	}

	if len(cart.Items) == 0 {
		return nil, ErrEmptyCart
	}

	var totalPrice float64
	var orderedItems []domain.OrderItem

	for _, item := range cart.Items {
		price := item.PriceSnapshot
		name := "Товар"
		image := ""

		if item.Product != nil {
			if price == 0 {
				price = item.Product.Price
			}
			if item.Product.Name != "" {
				name = item.Product.Name
			}
			image = item.Product.ImageURL
		}

		totalPrice += price * float64(item.Quantity)

		orderedItems = append(orderedItems, domain.OrderItem{
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			PriceSnapshot: price,
			NameSnapshot:  name,
			ImageSnapshot: image,
		})
	}

	order := &domain.Order{
		UserID:     userId,
		Status:     "pending",
		TotalPrice: totalPrice,
	}

	order, err = s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, mapOrderRepoError(err)
	}

	err = s.orderRepo.AddOrderItems(ctx, order.ID, orderedItems)
	if err != nil {
		return nil, mapOrderRepoError(err)
	}

	err = s.cartRepo.ClearCart(ctx, userId)
	if err != nil {
		return nil, mapCartRepoError(err)
	}

	order, err = s.orderRepo.GetOrderById(ctx, order.ID)
	if err != nil {
		return nil, mapOrderRepoError(err)
	}
	return order, nil
}
