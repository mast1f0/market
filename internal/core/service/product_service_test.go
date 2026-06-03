package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/service/mock"
	"testing"
)

func TestProductByIDSuccess(t *testing.T) {
	repo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				ID:   id,
				Name: "test",
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	product, err := service.GetProductById(1)
	if err != nil {
		t.Error(err)
	}

	if product.ID != 1 {
		t.Error("product id should be 1")
	}

	if product.Name != "test" {
		t.Error("product name should be test")
	}
}

func TestProductByIDInvalid(t *testing.T) {
	repo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				ID:   id,
				Name: "test",
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	product, err := service.GetProductById(-1)
	if err == nil {
		t.Error(err)
	}

	if errors.Is(err, ErrProductNotFound) {
		t.Error("product not found")
	}
	if product != nil {
		t.Error("product should be nil")
	}
}

func TestDeleteProductSuccess(t *testing.T) {
	called := false
	repo := &mock.ProductRepoMock{
		DeleteProductFunc: func(id int64) error {
			called = true
			return nil
		},
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 1,
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	err := service.DeleteProduct(1, 1, "admin")
	if err != nil {
		t.Error(err)
	}

	if !called {
		t.Error("product should be called")
	}
}

func TestDeleteProductForeignProductSuccessAdmin(t *testing.T) {
	called := false
	repo := &mock.ProductRepoMock{
		DeleteProductFunc: func(id int64) error {
			called = true
			return nil
		},
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 2,
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	err := service.DeleteProduct(1, 1, "admin")
	if err != nil {
		t.Error(err)
	}

	if !called {
		t.Error("product should be called")
	}
}

func TestDeleteProductForeignProductErrorSeller(t *testing.T) {
	repo := &mock.ProductRepoMock{
		DeleteProductFunc: func(id int64) error {
			return nil
		},
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 2,
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	err := service.DeleteProduct(1, 1, "seller")
	if !errors.Is(err, ErrForbidden) {
		t.Error(err)
	}
}

func TestUpdateProductSuccess(t *testing.T) {
	called := false
	repo := &mock.ProductRepoMock{
		UpdateProductFunc: func(product *domain.Product) (*domain.Product, error) {
			called = true
			return product, nil
		},
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 1,
			}, nil
		},
	}

	service := NewProductService(repo, &mock.CategoryRepoMock{})

	service.UpdateProduct(&domain.Product{
		ID:      1,
		OwnerID: 1,
	}, "admin")

	if !called {
		t.Error("product should be called")
	}
}

func TestUpdateProductForeignProductErrorSeller(t *testing.T) {
	repo := &mock.ProductRepoMock{
		UpdateProductFunc: func(product *domain.Product) (*domain.Product, error) {
			return product, nil
		},

		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 1,
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	_, err := service.UpdateProduct(&domain.Product{
		ID: 1,
	}, "seller")
	if !errors.Is(err, ErrForbidden) {
		t.Error(err)
	}
}

func TestUpdateProductOwnSellerSuccessn(t *testing.T) {
	called := false
	repo := &mock.ProductRepoMock{
		UpdateProductFunc: func(product *domain.Product) (*domain.Product, error) {
			called = true
			return product, nil
		},
		GetProductByIdFunc: func(id int64) (*domain.Product, error) {
			return &domain.Product{
				OwnerID: 1,
			}, nil
		},
	}
	service := NewProductService(repo, &mock.CategoryRepoMock{})
	_, err := service.UpdateProduct(&domain.Product{
		OwnerID: 1,
	}, "seller")

	if err != nil {
		t.Error(err)
	}

	if !called {
		t.Error("product should be called")
	}
}
