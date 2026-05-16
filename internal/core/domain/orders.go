package domain

import "time"

type Order struct {
	ID         int64       `gorm:"primaryKey" json:"id"`
	UserID     int64       `json:"user_id"`
	Status     string      `json:"status"`
	TotalPrice float64     `json:"total_price"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
