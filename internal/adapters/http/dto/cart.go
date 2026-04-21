package dto

import "market/internal/core/domain"

type AddCartItemRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity"`
}

type RemoveCartItemRequest struct {
	ProductID int64 `json:"product_id"`
}

func (r *AddCartItemRequest) ToDomain() *domain.CartItem {
	return &domain.CartItem{
		ProductID: r.ProductID,
		Quantity:  r.Quantity,
	}
}
