package orm

import (
	"market/internal/core/domain"

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

func (r *OrderRepository) GetOrderById(id int64) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *OrderRepository) GetOrderByUserId(userId int64) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (r *OrderRepository) AddOrderItems(orderId int64, items []domain.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var order domain.Order
		if err := tx.First(&order, orderId).Error; err != nil {
			return err
		}
		for i := range items {
			items[i].OrderID = orderId
		}
		if err := tx.Create(&items).Error; err != nil {
			return err
		}

		return nil
	})
}
func (r *OrderRepository) UpdateOrderStatus(orderId int64, status string) error {
	return r.db.Model(&domain.Order{}).
		Where("id = ?", orderId).
		Update("status", status).Error
}
