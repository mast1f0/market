package service

import (
	"context"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CartService struct {
	repo     ports.CartRepository
	products ports.ProductRepository
}

var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrInvalidProduct  = errors.New("invalid product id")
	ErrInvalidItem     = errors.New("invalid item id")
	ErrInvalidQuantity = errors.New("invalid quantity")

	ErrCartNotFound     = errors.New("cart not found")
	ErrCartItemNotFound = errors.New("cart item not found")
	ErrInvalidCartID    = errors.New("invalid cart id")

	ErrInternal = errors.New("internal service error")
)

func mapCartRepoError(err error) error {
	switch err {

	case ports.ErrCartNotFound:
		return ErrCartNotFound

	case ports.ErrCartItemNotFound:
		return ErrCartItemNotFound

	case ports.ErrFailedToLoadCart,
		ports.ErrFailedToSaveCart,
		ports.ErrFailedToDeleteCartItem,
		ports.ErrFailedToUpdateCartItem,
		ports.ErrFailedToClearCart,
		ports.ErrFailedToLoadCartItem:
		return ErrInternal

	default:
		return ErrInternal
	}
}

func NewCartService(cartRepo ports.CartRepository, productRepo ports.ProductRepository) *CartService {
	return &CartService{repo: cartRepo, products: productRepo}
}

func (s *CartService) GetCartWithItems(ctx context.Context, userID int64) (*domain.Cart, error) {
	if userID < 0 {
		return nil, ErrInvalidUserID
	}
	item, err := s.repo.GetCartWithItems(ctx, userID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return item, nil
}
func (s *CartService) FindCartItem(ctx context.Context, cartID, productID int64) (*domain.CartItem, error) {
	if cartID < 0 {
		return nil, ErrInvalidCartID
	}
	if productID < 0 {
		return nil, ErrInvalidProduct
	}
	item, err := s.repo.FindCartItem(ctx, cartID, productID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return item, nil
}

func (s *CartService) DeleteCartItem(ctx context.Context, userId int64, productId int64) error {
	if userId < 0 {
		return ErrInvalidUserID
	}
	if productId < 0 {
		return ErrInvalidProduct
	}
	err := s.repo.DeleteCartItem(ctx, userId, productId)
	if err != nil {
		mapCartRepoError(err)
	}
	return nil
}

func (s *CartService) UpdateCartItem(ctx context.Context, itemId int64, quantity int) (*domain.CartItem, error) {
	if itemId < 0 {
		return nil, ErrInvalidItem
	}
	if quantity < 0 {
		return nil, ErrInvalidQuantity
	}
	updated, err := s.repo.UpdateCartItem(ctx, itemId, quantity)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return updated, nil
}
func (s *CartService) CreateCart(ctx context.Context, userID int64) (*domain.Cart, error) {
	if userID < 0 {
		return nil, ErrInvalidCartID
	}
	cart, err := s.repo.CreateCart(ctx, userID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return cart, nil
}

func (s *CartService) AddCartItem(ctx context.Context, userID int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	if userID < 0 {
		return nil, ErrInvalidUserID
	}
	if cartItem.Quantity < 0 {
		return nil, ErrInvalidQuantity
	}
	if cartItem.ID < 0 {
		return nil, ErrInvalidItem
	}
	product, err := s.products.GetProductById(ctx, cartItem.ProductID)
	if err != nil {
		return nil, err
	}
	if cartItem.PriceSnapshot <= 0 {
		cartItem.PriceSnapshot = product.Price
	}
	newCartItem, err := s.repo.AddCartItem(ctx, userID, cartItem)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return newCartItem, nil
}

func (s *CartService) ClearCart(ctx context.Context, userID int64) error {
	if userID < 0 {
		return ErrInvalidUserID
	}
	err := s.repo.ClearCart(ctx, userID)
	if err != nil {
		return mapCartRepoError(err)
	}
	return nil
}
