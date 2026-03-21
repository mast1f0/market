package ports

import (
	"market/internal/core/domain"
)

type CartItemsRepository interface {
	AddCartItem(cartItems *domain.CartItems) (*domain.CartItems, error)
	DeleteCartItem(id int64) error
	GetCartItems(id int64) (*domain.CartItems, error)
	UpdateCartItem(cartItems *domain.CartItems) (*domain.CartItems, error)
}
