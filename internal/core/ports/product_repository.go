package ports

import (
	"market/internal/core/domain"
)

type ProductRepository interface {
	GetProducts() []domain.Product
	GetProduct(id int64) (*domain.Product, error)
	DeleteProduct(id int64) error
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
}
