package ports

import (
	"market/internal/core/domain"
)

type CategoryRepository interface {
	NewCategory(category *domain.Category) (*domain.Category, error)
	DeleteCategoryById(categoryId int) error
	CategoryById(id int) (*domain.Category, error)
	AddToCategory(category *domain.Category, productId int) (*domain.Category, error)
	GetAll() []domain.Category
}
