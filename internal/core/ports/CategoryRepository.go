package ports

import (
	"market/internal/core/domain"
)

type CategoryRepository interface {
	CreateCategory(category *domain.Category) (*domain.Category, error)
	UpdateCategory(category *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
	GetCategory(id int) (*domain.Category, error)
	GetCategoryByName(name string) *domain.Category
}
