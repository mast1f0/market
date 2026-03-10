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

func (s *CategoryService) GetCategory(id int) *domain.Category {
	return s.repository.GetCategory(id)
}

func (s *CategoryService) CreateCategory(category *domain.Category) (*domain.Category, error) {
	return s.repository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	return s.repository.UpdateCategory(category)
}
func (s *CategoryService) DeleteCategory(id int) error {
	return s.repository.DeleteCategory(id)
}
func (s *CategoryService) GetCategoryByName(name string) *domain.Category {
	return s.repository.GetCategoryByName(name)
}
