package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type ProductService struct {
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
}

func NewProductService(productRepo ports.ProductRepository, categoryRepo ports.CategoryRepository) *ProductService {
	return &ProductService{
		productRepository:  productRepo,
		categoryRepository: categoryRepo,
	}
}

func (s *ProductService) GetProductById(id int64) (*domain.Product, error) {
	product, err := s.productRepository.GetProduct(id)
	if err != nil {
		return nil, err
	}
	if id < 1 {
		return nil, errors.New("invalid product id")
	}
	return product, nil
}

func (s *ProductService) AddToProduct(product *domain.Product) (*domain.Product, error) {
	_, err := s.categoryRepository.GetCategory(product.CategoryID)
	if err != nil {
		return nil, errors.New("category Not Found")
	}
	createdProduct, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}
func (s *ProductService) UpdateProduct(id int64, userId int64) (*domain.Product, error) {
	product, err := s.productRepository.GetProduct(id)
	if err != nil {
		return nil, err
	}
	if product.OwnerID != userId {
		return nil, errors.New("you are not the owner of this product")
	}
	_, err = s.categoryRepository.GetCategory(product.CategoryID)
	if err != nil {
		return nil, errors.New("category Not Found")
	}
	updatedProduct, err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.productRepository.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() []domain.Product {
	return s.productRepository.GetProducts()
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	return s.productRepository.GetProduct(id)
}
