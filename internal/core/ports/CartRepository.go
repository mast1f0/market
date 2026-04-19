package ports

import (
	"market/internal/core/domain"
)

type CartRepository interface {
	CreateCart(cart *domain.Cart) (*domain.Cart, error)
	// GetCart возвращает корзину текущего пользователя по user_id.
	GetCart(userID int64) (*domain.Cart, error)
	GetCartByID(id int64) (*domain.Cart, error)
	UpdateCart(cart *domain.Cart) (*domain.Cart, error)
	DeleteCart(userID int64) error
}
