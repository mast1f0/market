package dto

import (
	"errors"

	"market/internal/core/domain"
)

type AddCartItemRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func (r *AddCartItemRequest) Validate() error {
	if r.ProductID <= 0 {
		return errors.New("product_id is required")
	}
	if r.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}
	return nil
}

func (r *AddCartItemRequest) ToDomain(cartID int64) *domain.CartItems {
	return &domain.CartItems{
		CartId:    cartID,
		ProductId: r.ProductID,
		Quantity:  r.Quantity,
	}
}

type UpdateCartItemRequest struct {
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func (r *UpdateCartItemRequest) Validate() error {
	if r.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}
	return nil
}

func (r *UpdateCartItemRequest) ApplyTo(item *domain.CartItems) {
	if r.CartId > 0 {
		item.CartId = r.CartId
	}
	if r.ProductId > 0 {
		item.ProductId = r.ProductId
	}
	item.Quantity = r.Quantity
}
