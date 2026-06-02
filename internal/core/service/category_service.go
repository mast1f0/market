package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CategoryService struct {
	repository ports.CategoryRepository
}

var (
	ErrInvalidCategoryID      = errors.New("invalid category id")
	ErrInvalidCategoryName    = errors.New("invalid category name")
	ErrCategoryExists         = errors.New("category already exists")
	ErrCategoryNotFound       = errors.New("category not found")
	ErrFailedToCreateCategory = errors.New("failed to create category")
	ErrFailedToLoadCategory   = errors.New("failed to load category")
	ErrFailedToUpdateCategory = errors.New("failed to update category")
	ErrFailedToDeleteCategory = errors.New("failed to delete category")
)

func NewCategoryService(repository ports.CategoryRepository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) GetCategory(id int64) (*domain.Category, error) {
	category, err := s.repository.GetCategory(id)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return category, nil
}

func (s *CategoryService) CreateCategory(categoryName string) (*domain.Category, error) {
	if len(categoryName) == 0 {
		return nil, ErrInvalidCategoryName
	}
	category, err := s.repository.CreateCategory(categoryName)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryAlreadyExists) {
			return nil, ErrCategoryExists
		}
		return nil, ErrFailedToCreateCategory
	}
	return category, nil
}

func (s *CategoryService) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	newCategory, err := s.repository.UpdateCategory(category)
	if err != nil {
		return nil, ErrFailedToUpdateCategory
	}
	return newCategory, nil
}
func (s *CategoryService) DeleteCategory(id int64) error {
	if id < 1 {
		return ErrInvalidCategoryID
	}
	err := s.repository.DeleteCategory(id)
	if err != nil {
		return ErrFailedToDeleteCategory
	}
	return nil
}
func (s *CategoryService) GetCategoryByName(name string) (*domain.Category, error) {
	if len(name) == 0 {
		return nil, ErrInvalidCategoryName
	}
	category, err := s.repository.GetCategoryByName(name)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return category, nil
}

func (s *CategoryService) GetCategories() ([]domain.Category, error) {
	return s.repository.GetCategories()
}

func (s *CategoryService) GetCategoriesByCategoryID(id int64) ([]domain.Product, error) {
	if id < 1 {
		return nil, ErrInvalidCategoryID
	}
	products, err := s.repository.ProductsByCategory(id)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return products, nil
}
