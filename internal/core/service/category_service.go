package service

import (
	"context"
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

func (s *CategoryService) GetCategory(ctx context.Context, id int64) (*domain.Category, error) {
	if id < 0 {
		return nil, ErrInvalidCategoryID
	}
	category, err := s.repository.GetCategory(ctx, id)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return category, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, categoryName string) (*domain.Category, error) {
	if len(categoryName) == 0 {
		return nil, ErrInvalidCategoryName
	}
	category, err := s.repository.CreateCategory(ctx, categoryName)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryAlreadyExists) || errors.Is(err, ports.ErrCategoryExists) {
			return nil, ErrCategoryExists
		}
		return nil, ErrFailedToCreateCategory
	}
	return category, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if category.Name == "" {
		return nil, ErrInvalidCategoryName
	}
	newCategory, err := s.repository.UpdateCategory(ctx, category)
	if err != nil {
		return nil, ErrFailedToUpdateCategory
	}
	return newCategory, nil
}
func (s *CategoryService) DeleteCategory(ctx context.Context, id int64) error {
	if id < 0 {
		return ErrInvalidCategoryID
	}
	err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return ErrFailedToDeleteCategory
	}
	return nil
}
func (s *CategoryService) GetCategoryByName(ctx context.Context, name string) (*domain.Category, error) {
	if len(name) == 0 {
		return nil, ErrInvalidCategoryName
	}
	category, err := s.repository.GetCategoryByName(ctx, name)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return category, nil
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repository.GetCategories(ctx)
}

func (s *CategoryService) GetCategoriesByCategoryID(ctx context.Context, id int64) ([]domain.Product, error) {
	if id < 0 {
		return nil, ErrInvalidCategoryID
	}
	products, err := s.repository.ProductsByCategory(ctx, id)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrFailedToLoadCategory
	}
	return products, nil
}
