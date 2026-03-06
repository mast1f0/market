package ports

import "market/internal/domain"

type CategoryRepository interface {
	NewCategory(category *domain.Category) (*domain.Category, error)
	DeleteCategoryById(categoryId int) error
	CategoryByName(name string) (*domain.Category, error)
	AddToCategory(category *domain.Category, productId int) (*domain.Category, error)
}
