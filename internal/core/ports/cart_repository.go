package ports

import (
	"errors"
	"market/internal/core/domain"
)

var (
	ErrCartNotFound           = errors.New("cart not found")
	ErrCartItemNotFound       = errors.New("cart item not found")
	ErrFailedToLoadCart       = errors.New("failed to load cart")
	ErrFailedToSaveCart       = errors.New("failed to save cart")
	ErrFailedToDeleteCartItem = errors.New("failed to delete cart item")
	ErrFailedToUpdateCartItem = errors.New("failed to update cart item")
	ErrFailedToClearCart      = errors.New("failed to clear cart")
	ErrFailedToLoadCartItem   = errors.New("failed to load cart item")
)

type CartRepository interface {
	CreateCart(id int64) (*domain.Cart, error)
	GetCartWithItems(userID int64) (*domain.Cart, error)
	ClearCart(userId int64) error

	FindCartItem(cartID, productID int64) (*domain.CartItem, error)
	DeleteCartItem(userId int64, productId int64) error
	UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error)
	AddCartItem(userId int64, cartItem *domain.CartItem) (*domain.CartItem, error)
}
