package service

import (
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

func (s *CartService) GetCartWithItems(userID int64) (*domain.Cart, error) {
	item, err := s.repo.GetCartWithItems(userID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return item, nil
}
func (s *CartService) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	if cartID < 0 {
		return nil, ErrInvalidCartID
	}
	if productID < 0 {
		return nil, ErrInvalidProduct
	}
	item, err := s.repo.FindCartItem(cartID, productID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return item, nil
}

func (s *CartService) DeleteCartItem(userId int64, productId int64) error {
	err := s.repo.DeleteCartItem(userId, productId)
	return mapCartRepoError(err)
}

func (s *CartService) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	if itemId < 0 {
		return nil, ErrInvalidItem
	}
	if quantity < 0 {
		return nil, ErrInvalidQuantity
	}
	updated, err := s.repo.UpdateCartItem(itemId, quantity)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return updated, nil
}
func (s *CartService) CreateCart(cartID int64) (*domain.Cart, error) {
	if cartID < 0 {
		return nil, ErrInvalidCartID
	}
	cart, err := s.repo.CreateCart(cartID)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return cart, nil
}

func (s *CartService) AddCartItem(userID int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	if userID < 0 {
		return nil, ErrInvalidUserID
	}
	if cartItem.Quantity < 0 {
		return nil, ErrInvalidQuantity
	}
	if cartItem.ID < 0 {
		return nil, ErrInvalidItem
	}
	product, err := s.products.GetProductById(cartItem.ProductID)
	if err != nil {
		return nil, err
	}
	if cartItem.PriceSnapshot <= 0 {
		cartItem.PriceSnapshot = product.Price
	}
	newCartItem, err := s.repo.AddCartItem(userID, cartItem)
	if err != nil {
		return nil, mapCartRepoError(err)
	}
	return newCartItem, nil
}

func (s *CartService) ClearCart(userID int64) error {
	err := s.repo.ClearCart(userID)
	if err != nil {
		return mapCartRepoError(err)
	}
	return nil
}
