package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/service/mock"
	"testing"
)

func TestDeleteFailedInvalidUserID(t *testing.T) {
	executed := false
	cartRepo := &mock.CartRepoMock{
		DeleteCartItemById: func(userId int64, productId int64) error {
			executed = true
			return nil
		},
	}
	productRepo := &mock.ProductRepoMock{}

	service := NewCartService(cartRepo, productRepo)

	err := service.DeleteCartItem(-1, 2)
	if !errors.Is(err, ErrInvalidUserID) {
		t.Error(err)
	}
	if executed {
		t.Error("unexpected call DeleteCartItemById function")
	}
}

func TestDeleteFailedInvalidProductID(t *testing.T) {
	executed := false
	cartRepo := &mock.CartRepoMock{
		DeleteCartItemById: func(userId int64, productId int64) error {
			executed = true
			return nil
		},
	}
	productRepo := &mock.ProductRepoMock{}
	service := NewCartService(cartRepo, productRepo)
	err := service.DeleteCartItem(1, -2)
	if !errors.Is(err, ErrInvalidProduct) {
		t.Error(err)
	}
	if executed {
		t.Error("unexpected call DeleteCartItemById function")
	}
}

func TestDeleteSuccess(t *testing.T) {
	executed := false
	cartRepo := &mock.CartRepoMock{
		DeleteCartItemById: func(userId int64, productId int64) error {
			executed = true
			return nil
		},
	}
	productRepo := &mock.ProductRepoMock{}
	service := NewCartService(cartRepo, productRepo)
	err := service.DeleteCartItem(1, 1)
	if err != nil {
		t.Error(err)
	}
	if !executed {
		t.Error("expected to call DeleteCartItemById function")
	}
}

func TestAddCartItemSuccess(t *testing.T) {
	executed := false
	cartRepo := &mock.CartRepoMock{
		AddCartItemA: func(userId int64, item *domain.CartItem) (*domain.CartItem, error) {
			executed = true
			return &domain.CartItem{}, nil
		},
	}
	productRepo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(productId int64) (*domain.Product, error) {
			return &domain.Product{
				ID:    123,
				Price: 100,
			}, nil
		},
	}
	service := NewCartService(cartRepo, productRepo)
	_, err := service.AddCartItem(1, &domain.CartItem{
		CartID:    1,
		Quantity:  2,
		ProductID: 123,
	})
	if err != nil {
		t.Error(err)
	}

	if !executed {
		t.Error("expected to call AddCartItem function")
	}
}

func TestAddCartItemFailInvalidId(t *testing.T) {
	cartRepo := &mock.CartRepoMock{
		AddCartItemA: func(userId int64, item *domain.CartItem) (*domain.CartItem, error) {
			return nil, errors.New("error")
		},
	}
	productRepo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(productId int64) (*domain.Product, error) {
			return &domain.Product{}, nil
		},
	}
	service := NewCartService(cartRepo, productRepo)
	_, err := service.AddCartItem(1, &domain.CartItem{
		ID:        -121,
		Quantity:  2,
		ProductID: 123,
	})
	if !errors.Is(err, ErrInvalidItem) {
		t.Error(err)
	}
}

func TestAddCartItemFailInvalidQuantity(t *testing.T) {
	cartRepo := &mock.CartRepoMock{
		AddCartItemA: func(userId int64, item *domain.CartItem) (*domain.CartItem, error) {
			return nil, errors.New("error")
		},
	}
	productRepo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(productId int64) (*domain.Product, error) {
			return &domain.Product{}, nil
		},
	}
	service := NewCartService(cartRepo, productRepo)
	_, err := service.AddCartItem(1, &domain.CartItem{
		ID:        121,
		Quantity:  -2,
		ProductID: 123,
	})
	if !errors.Is(err, ErrInvalidQuantity) {
		t.Error(err)
	}
}

func TestCreateCartSuccess(t *testing.T) {
	executed := false
	var userID int64 = 1
	cartRepo := &mock.CartRepoMock{
		CreateCartByUserId: func(userId int64) (*domain.Cart, error) {
			executed = true
			return &domain.Cart{
				UserID: &userID,
				ID:     1,
			}, nil
		},
	}
	productRepo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(productId int64) (*domain.Product, error) {
			return &domain.Product{
				ID: 1,
			}, nil
		},
	}
	service := NewCartService(cartRepo, productRepo)
	cart, err := service.CreateCart(1)
	if err != nil {
		t.Error(err)
	}
	if !executed {
		t.Error("expected to call CreateCartByUserId function")
	}
	if *cart.UserID != 1 {
		t.Error("strange id")
	}
}

func TestCreateCartFailedInvalidId(t *testing.T) {
	executed := false
	var userID int64 = -1
	cartRepo := &mock.CartRepoMock{
		CreateCartByUserId: func(userId int64) (*domain.Cart, error) {
			executed = true
			return &domain.Cart{
				UserID: &userID,
				ID:     1,
			}, nil
		},
	}
	productRepo := &mock.ProductRepoMock{
		GetProductByIdFunc: func(productId int64) (*domain.Product, error) {
			return &domain.Product{
				ID: 1,
			}, nil
		},
	}
	service := NewCartService(cartRepo, productRepo)
	_, err := service.CreateCart(-11)
	if err == nil {
		t.Error(err)
	}
	if executed {
		t.Error("unexpected to call CreateCartByUserId function")
	}
}
