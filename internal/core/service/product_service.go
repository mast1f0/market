package service

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrInvalidProductID = errors.New("invalid product id")
	ErrForbidden        = errors.New("forbidden")
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
	if id < 0 {
		return nil, ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}

func (s *ProductService) CreateProduct(product *domain.Product) (*domain.Product, error) {
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
func (s *ProductService) UpdateProduct(newProduct *domain.Product, role string) (*domain.Product, error) {
	product, err := s.productRepository.GetProductById(newProduct.ID)
	if err != nil {
		return nil, err
	}
	if product.OwnerID != newProduct.OwnerID && role != "admin" {
		return nil, ErrForbidden
	}
	_, err = s.categoryRepository.GetCategory(newProduct.CategoryID)
	if err != nil {
		return nil, errors.New("category Not Found")
	}
	updatedProduct, err := s.productRepository.UpdateProduct(newProduct)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(id int64, userId int64, role string) error {
	if id < 0 {
		return ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		return err
	}
	if product.OwnerID != userId && role != "admin" {
		return ErrForbidden
	}

	err = s.productRepository.DeleteProduct(id)
	if errors.Is(err, ports.ErrNotFound) {
		return ErrProductNotFound
	}
	return err
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.productRepository.GetProducts()
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	if id < 0 {
		return nil, ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, errors.New("failed to load product")
	}
	return product, nil
}
