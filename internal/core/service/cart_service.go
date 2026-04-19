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

func (s *CartService) GetCart(userID int64) (*domain.Cart, error) {
	return s.repo.GetCart(userID)
}

func (s *CartService) GetCartByID(id int64) (*domain.Cart, error) {
	return s.repo.GetCartByID(id)
}

func (s *CartService) UpdateCart(cart *domain.Cart) (*domain.Cart, error) {
	return s.repo.UpdateCart(cart)
}
func (s *CartService) CreateCart(cart *domain.Cart) (*domain.Cart, error) {
	return s.repo.CreateCart(cart)
}

func (s *CartService) DeleteCart(userID int64) error {
	return s.repo.DeleteCart(userID)
}
