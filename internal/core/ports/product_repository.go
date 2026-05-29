package ports

import (
	"errors"
	"market/internal/core/domain"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrConflict            = errors.New("conflict")
	ErrInvalidData         = errors.New("invalid data")
	ErrFailedToLoadProduct = errors.New("failed to load product")
)

type ProductRepository interface {
	GetProducts() ([]domain.Product, error)
	GetProductById(id int64) (*domain.Product, error)
	DeleteProduct(id int64) error
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
}
