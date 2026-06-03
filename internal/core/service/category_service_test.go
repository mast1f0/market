package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/service/mock"
	"testing"
)

func TestCreateProductSuccess(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		CreateCategoryByName: func(name string) (*domain.Category, error) {
			return &domain.Category{
				Name: name,
			}, nil
		},
	}
	service := NewCategoryService(repo)
	category, err := service.CreateCategory("name")
	if err != nil {
		t.Error(err)
	}
	if category.Name != "name" {
		t.Error("Category Name should be 'name'")
	}
}

func TestCreateProductInvalidName(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		CreateCategoryByName: func(name string) (*domain.Category, error) {
			return &domain.Category{
				Name: name,
			}, nil
		},
	}
	service := NewCategoryService(repo)
	_, err := service.CreateCategory("")
	if err == nil {
		t.Error(err)
	}
}

func TestCreateProductFail(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		CreateCategoryByName: func(name string) (*domain.Category, error) {
			return nil, ErrFailedToCreateCategory
		},
	}
	service := NewCategoryService(repo)
	_, err := service.CreateCategory("name")
	if err == nil {
		t.Error("Category should not be created")
	}

	if !errors.Is(err, ErrFailedToCreateCategory) {
		t.Error(err)
	}
}

func TestUpdateCategorySuccess(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		UpdateCategoryName: func(category *domain.Category) (*domain.Category, error) {
			return category, nil
		},
	}
	service := NewCategoryService(repo)
	newCategory, err := service.UpdateCategory(&domain.Category{
		Name: "name",
	})
	if err != nil {
		t.Error(err)
	}
	if newCategory.Name != "name" {
		t.Error("Category Name should be 'name'")
	}
}

func TestUpdateCategoryFailed(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		UpdateCategoryName: func(category *domain.Category) (*domain.Category, error) {
			return category, nil
		},
	}
	service := NewCategoryService(repo)
	_, err := service.UpdateCategory(&domain.Category{
		Name: "",
	})
	if err == nil {
		t.Error(err)
	}
}

func TestDeleteCategorySuccess(t *testing.T) {
	repo := &mock.CategoryRepoMock{
		DeleteCategoryById: func(id int64) error {
			return nil
		},
	}
	service := NewCategoryService(repo)
	err := service.DeleteCategory(1)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCategoryInvalidId(t *testing.T) {
	executed := false
	repo := &mock.CategoryRepoMock{
		DeleteCategoryById: func(id int64) error {
			executed = true
			return nil
		},
	}
	service := NewCategoryService(repo)
	err := service.DeleteCategory(-1)
	if !errors.Is(err, ErrInvalidCategoryID) {
		t.Error(err)
	}
	if executed {
		t.Error("Category should be executed")
	}
}

func TestDeleteCategoryFailed(t *testing.T) {
	executed := false
	repo := &mock.CategoryRepoMock{
		DeleteCategoryById: func(id int64) error {
			executed = true
			return ErrFailedToDeleteCategory
		},
	}
	service := NewCategoryService(repo)
	err := service.DeleteCategory(1)
	if !errors.Is(err, ErrFailedToDeleteCategory) {
		t.Error(err)
	}
	if !executed {
		t.Error("Category should be executed")
	}
}
