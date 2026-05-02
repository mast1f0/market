package dto

type OrderItemResponse struct {
	ID        int64   `json:"id"`
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	ImageURL  string  `json:"image_url"`
}

type OrderResponse struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"user_id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"total_price"`

	Items []OrderItemResponse `json:"items,omitempty"`

	CreatedAt string `json:"created_at"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}
