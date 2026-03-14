package service

import (
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type ProductService struct {
	repository ports.ProductRepository
}

func NewProductService(repository ports.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) GetProductById(id int) (*domain.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *ProductService) AddToProduct(product *domain.Product) (*domain.Product, error) {
	return s.repository.CreateProduct(product)
}
func (s *ProductService) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	return s.repository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repository.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() []domain.Product {
	return s.repository.GetProducts()
}
