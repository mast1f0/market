package ports

import "market/internal/core/domain"

type CartRepository interface {
	CreateCart(id int64) (*domain.Cart, error)
	GetCartWithItems(userID int64) (*domain.Cart, error)

	FindCartItem(cartID, productID int64) (*domain.CartItem, error)
	DeleteCartItem(userId int64, productId int64) error
	UpdateCartItem(cartItems *domain.CartItem) (*domain.CartItem, error)
	AddCartItem(userId int64, cartItem *domain.CartItem) (*domain.CartItem, error)
}
