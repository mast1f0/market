package service

import (
	"context"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"
)

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrInvalidProductID = errors.New("invalid product id")
	ErrForbidden        = errors.New("forbidden")
)

func mapProductRepoError(err error) error {
	switch {
	case errors.Is(err, ports.ErrNotFound):
		return ErrProductNotFound
	case errors.Is(err, ports.ErrAlreadyExists):
		return ErrAlreadyExists
	case errors.Is(err, ports.ErrConflict):
		return ErrConflict
	case errors.Is(err, ports.ErrInvalidData):
		return ErrInvalidData
	case errors.Is(err, ports.ErrFailedToLoadProduct):
		return ErrFailedToLoadProduct
	default:
		return ErrFailedToLoadProduct
	}
}

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

func (s *ProductService) GetProductById(ctx context.Context, id int64) (*domain.Product, error) {
	if id < 0 {
		return nil, ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(ctx, id)
	if err != nil {
		return nil, mapProductRepoError(err)
	}
	return product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	_, err := s.categoryRepository.GetCategory(ctx, product.CategoryID)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	createdProduct, err := s.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, mapProductRepoError(err)
	}
	return createdProduct, nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, newProduct *domain.Product, role string) (*domain.Product, error) {
	product, err := s.productRepository.GetProductById(ctx, newProduct.ID)
	if err != nil {
		return nil, mapProductRepoError(err)
	}
	if product.OwnerID != newProduct.OwnerID && role != "admin" {
		return nil, ErrForbidden
	}
	_, err = s.categoryRepository.GetCategory(ctx, newProduct.CategoryID)
	if err != nil {
		if errors.Is(err, ports.ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	updatedProduct, err := s.productRepository.UpdateProduct(ctx, newProduct)
	if err != nil {
		return nil, mapProductRepoError(err)
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64, userId int64, role string) error {
	if id < 0 {
		return ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(ctx, id)
	if err != nil {
		return mapProductRepoError(err)
	}
	if product.OwnerID != userId && role != "admin" {
		return ErrForbidden
	}

	err = s.productRepository.DeleteProduct(ctx, id)
	if err != nil {
		return mapProductRepoError(err)
	}
	return nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.productRepository.GetProducts(ctx)
}

func (s *ProductService) GetProduct(ctx context.Context, id int64) (*domain.Product, error) {
	if id < 0 {
		return nil, ErrInvalidProductID
	}
	product, err := s.productRepository.GetProductById(ctx, id)
	if err != nil {
		return nil, mapProductRepoError(err)
	}
	return product, nil
}
