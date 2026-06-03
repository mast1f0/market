package mock

import "market/internal/core/domain"

type CategoryRepoMock struct {
	CreateCategoryByName func(category string) (*domain.Category, error)
	UpdateCategoryName   func(category *domain.Category) (*domain.Category, error)
	DeleteCategoryById   func(id int64) error
}

func (r *CategoryRepoMock) CreateCategory(category string) (*domain.Category, error) {
	return r.CreateCategoryByName(category)
}
func (r *CategoryRepoMock) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	return r.UpdateCategoryName(category)
}
func (r *CategoryRepoMock) DeleteCategory(id int64) error {
	return r.DeleteCategoryById(id)
}
func (r *CategoryRepoMock) GetCategory(id int64) (*domain.Category, error) {
	return &domain.Category{}, nil
}
func (r *CategoryRepoMock) GetCategoryByName(name string) (*domain.Category, error) {
	return &domain.Category{}, nil
}
func (r *CategoryRepoMock) GetCategories() ([]domain.Category, error) {
	return []domain.Category{}, nil
}
func (r *CategoryRepoMock) ProductsByCategory(id int64) ([]domain.Product, error) {
	return []domain.Product{}, nil
}
