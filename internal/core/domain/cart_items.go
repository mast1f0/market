package domain

import "time"

type CartItem struct {
	ID        int64 `json:"id"`
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`

	PriceSnapshot float64 `json:"price_snapshot"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
