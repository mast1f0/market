package ports

import (
	"market/internal/core/domain"
)

type CartRepository interface {
	CreateCart(cart *domain.Cart) (*domain.Cart, error)
	GetCart(id int64) (*domain.Cart, error)
	UpdateCart(cart *domain.Cart) (*domain.Cart, error)
	DeleteCart(id int64) error
}
