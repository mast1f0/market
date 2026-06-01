package orm

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *domain.Order) (*domain.Order, error) {
	err := r.db.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func ensureOrderItems(order *domain.Order) {
	if order.Items == nil {
		order.Items = []domain.OrderItem{}
	}
	for i, item := range order.Items {
		if item.ImageSnapshot == "" && item.Product != nil {
			order.Items[i].ImageSnapshot = item.Product.ImageURL
		}
		if item.NameSnapshot == "" && item.Product != nil && item.Product.Name != "" {
			order.Items[i].NameSnapshot = item.Product.Name
		}
	}
}

func (r *OrderRepository) GetOrderById(id int64) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items.Product").Where("id = ?", id).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrOrderNotFound
		}
		return nil, ports.ErrFailedToLoadOrder
	}
	ensureOrderItems(&order)
	return &order, nil
}
func (r *OrderRepository) GetOrderByUserId(userId int64) ([]domain.Order, error) {
	orders := make([]domain.Order, 0)
	err := r.db.Preload("Items.Product").Where("user_id = ?", userId).Order("created_at DESC").Find(&orders).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrOrderNotFound
		}
		return nil, ports.ErrFailedToLoadOrder
	}
	for i := range orders {
		ensureOrderItems(&orders[i])
	}
	return orders, nil
}
func (r *OrderRepository) AddOrderItems(orderId int64, items []domain.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var order domain.Order
		if err := tx.First(&order, orderId).Error; err != nil {
			return ports.ErrOrderNotFound
		}
		for i := range items {
			items[i].OrderID = orderId
		}
		if err := tx.Create(&items).Error; err != nil {
			return ports.ErrFailedToSaveOrder
		}

		return nil
	})
}
func (r *OrderRepository) UpdateOrderStatus(orderId int64, status string) error {
	err := r.db.Model(&domain.Order{}).
		Where("id = ?", orderId).
		Update("status", status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrOrderNotFound
		}
		return ports.ErrFailedToSaveOrder
	}
	return nil
}
