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

func (s *CategoryService) Create(category *domain.Category) (*domain.Category, error) {
	return s.repository.NewCategory(category)
}
func (s *CategoryService) Get(id int) (*domain.Category, error) {
	return s.repository.CategoryById(id)
}

func (s *CategoryService) GetAll() []domain.Category {
	return s.repository.GetAll()
}

func (s *CategoryService) Delete(id int) error {
	return s.Delete(id)
}
