package ports

import (
	"market/internal/core/domain"
)

type CartItemsRepository interface {
	AddCartItem(cartItems *domain.CartItems) (*domain.CartItems, error)
	DeleteCartItem(id int) error
	GetCartItems(id int) (*domain.CartItems, error)
	UpdateCartItem(cartItems *domain.CartItems) (*domain.CartItems, error)
}
