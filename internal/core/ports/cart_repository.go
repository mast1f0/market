package ports

import (
	"context"
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
	CreateCart(ctx context.Context, id int64) (*domain.Cart, error)
	GetCartWithItems(ctx context.Context, userID int64) (*domain.Cart, error)
	ClearCart(ctx context.Context, userId int64) error

	FindCartItem(ctx context.Context, cartID, productID int64) (*domain.CartItem, error)
	DeleteCartItem(ctx context.Context, userId int64, productId int64) error
	UpdateCartItem(ctx context.Context, itemId int64, quantity int) (*domain.CartItem, error)
	AddCartItem(ctx context.Context, userId int64, cartItem *domain.CartItem) (*domain.CartItem, error)
}
