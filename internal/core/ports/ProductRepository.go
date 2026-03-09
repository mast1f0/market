package ports

import (
	"market/internal/core/domain"
)

type ProductRepository interface {
	GetProducts() []domain.Product
	GetProduct(id int) *domain.Product
	DeleteProduct(id int) error
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
}
