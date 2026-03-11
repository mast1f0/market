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

func (s *CartService) GetCart(id int) (*domain.Cart, error) {
	return s.repo.GetCart(id)
}

func (s *CartService) UpdateCart(cart *domain.Cart) (*domain.Cart, error) {
	return s.repo.UpdateCart(cart)
}
func (s *CartService) CreateCart(cart *domain.Cart) (*domain.Cart, error) {
	return s.repo.CreateCart(cart)
}

func (s *CartService) DeleteCart(id int) error {
	return s.repo.DeleteCart(id)
}
