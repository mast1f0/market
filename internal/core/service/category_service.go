package service

import (
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CategoryService struct {
	repository ports.CategoryRepository
}

func NewCategoryService(repository ports.CategoryRepository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) GetCategory(id int64) (*domain.Category, error) {
	return s.repository.GetCategory(id)
}

func (s *CategoryService) CreateCategory(category string) (*domain.Category, error) {
	return s.repository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	return s.repository.UpdateCategory(category)
}
func (s *CategoryService) DeleteCategory(id int64) error {
	return s.repository.DeleteCategory(id)
}
func (s *CategoryService) GetCategoryByName(name string) *domain.Category {
	return s.repository.GetCategoryByName(name)
}

func (s *CategoryService) GetCategories() []domain.Category {
	return s.repository.GetCategories()
}

func (s *CategoryService) GetCategoriesByCategoryID(id int) ([]domain.Product, error) {
	return s.repository.ProductsByCategory(id)
}
