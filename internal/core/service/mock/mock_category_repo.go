package mock

import "market/internal/core/domain"

//заглушечка

type CategoryRepoMock struct {
}

func (r *CategoryRepoMock) CreateCategory(category string) (*domain.Category, error) {
	return &domain.Category{}, nil
}
func (r *CategoryRepoMock) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	return &domain.Category{}, nil
}
func (r *CategoryRepoMock) DeleteCategory(id int64) error {
	return nil
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
