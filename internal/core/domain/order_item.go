package domain

import "time"

type OrderItem struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	OrderID       int64     `json:"order_id"`
	ProductID     int64     `json:"product_id"`
	Quantity      int       `json:"quantity"`
	PriceSnapshot float64   `json:"price_snapshot"`
	NameSnapshot  string    `json:"name_snapshot"`
	ImageSnapshot string    `json:"image_snapshot"`
	Product       *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt     time.Time `json:"created_at"`
}
