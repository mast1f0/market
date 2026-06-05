package ports

import (
	"context"
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
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetProductById(ctx context.Context, id int64) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductsByName(ctx context.Context, name string) ([]domain.Product, error)
}
