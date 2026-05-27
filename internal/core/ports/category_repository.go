package ports

import (
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
	CreateCategory(category string) (*domain.Category, error)
	UpdateCategory(category *domain.Category) (*domain.Category, error)
	DeleteCategory(id int64) error
	GetCategory(id int64) (*domain.Category, error)
	GetCategoryByName(name string) (*domain.Category, error)
	GetCategories() []domain.Category
	ProductsByCategory(id int64) ([]domain.Product, error)
}
