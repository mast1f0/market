package domain

import "time"

type CartItem struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`

	PriceSnapshot float64  `json:"price_snapshot"`
	Product       *Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
