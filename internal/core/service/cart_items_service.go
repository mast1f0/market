package service

import (
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CartItemsService struct {
	repo ports.CartItemsRepository
}

func NewCartItemsService(repo ports.CartItemsRepository) *CartItemsService {
	return &CartItemsService{repo: repo}
}

func (s *CartItemsService) AddCartItem(cartItems *domain.CartItems) (*domain.CartItems, error) {
	return s.repo.AddCartItem(cartItems)
}

func (s *CartItemsService) DeleteCartItem(id int64) error {
	return s.repo.DeleteCartItem(id)
}

func (s *CartItemsService) GetCartItems(id int64) (*domain.CartItems, error) {
	return s.repo.GetCartItems(id)
}

func (s *CartItemsService) UpdateCartItem(cartItems *domain.CartItems) (*domain.CartItems, error) {
	return s.repo.UpdateCartItem(cartItems)
}
