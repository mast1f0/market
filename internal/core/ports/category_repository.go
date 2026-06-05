package ports

import (
	"context"
	"errors"
	"market/internal/core/domain"
)

var (
	ErrCategoryNotFound       = errors.New("Category not found")
	ErrCategoryExists         = errors.New("Category already exists")
	ErrFailedToCreateCategory = errors.New("Failed to create category")
	ErrFailedToUpdateCategory = errors.New("Failed to update category")
	ErrFailedToDeleteCategory = errors.New("Failed to delete category")
	ErrCategoryAlreadyExists  = errors.New("Category already exists")
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category string) (*domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	GetCategory(ctx context.Context, id int64) (*domain.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*domain.Category, error)
	GetCategories(ctx context.Context) ([]domain.Category, error)
	ProductsByCategory(ctx context.Context, id int64) ([]domain.Product, error)
}
