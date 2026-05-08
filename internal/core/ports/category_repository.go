package ports

import (
	"market/internal/core/domain"
)

type CategoryRepository interface {
	CreateCategory(category string) (*domain.Category, error)
	UpdateCategory(category *domain.Category) (*domain.Category, error)
	DeleteCategory(id int64) error
	GetCategory(id int64) (*domain.Category, error)
	GetCategoryByName(name string) *domain.Category
	GetCategories() []domain.Category
	ProductsByCategory(id int64) ([]domain.Product, error)
}
