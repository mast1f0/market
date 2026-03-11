package ports

import (
	"market/internal/core/domain"
)

type CartRepository interface {
	CreateCart(cart *domain.Cart) (*domain.Cart, error)
	GetCart(id int) (*domain.Cart, error)
	UpdateCart(cart *domain.Cart) (*domain.Cart, error)
	DeleteCart(id int) error
}
