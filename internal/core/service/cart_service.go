package service

import (
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CartService struct {
	repo ports.CartRepository
}

func NewCartService(cartRepo ports.CartRepository) *CartService {
	return &CartService{repo: cartRepo}
}

func (s *CartService) GetCartWithItems(userID int64) (*domain.Cart, error) {
	return s.repo.GetCartWithItems(userID)
}
func (s *CartService) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	return s.repo.FindCartItem(cartID, productID)
}

func (s *CartService) DeleteCartItem(userId int64, productId int64) error {
	return s.repo.DeleteCartItem(userId, productId)
}

func (s *CartService) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	return s.repo.UpdateCartItem(itemId, quantity)
}
func (s *CartService) CreateCart(cartID int64) (*domain.Cart, error) {
	return s.repo.CreateCart(cartID)
}

func (s *CartService) AddCartItem(userID int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	return s.repo.AddCartItem(userID, cartItem)
}
