package service

import (
	"market/internal/core/domain"
	"market/internal/core/ports"
)

type ProductService struct {
	repository ports.ProductRepository
	S3storage  ports.S3Repository
}

func NewProductService(repository ports.ProductRepository, s3 ports.S3Repository) *ProductService {
	return &ProductService{repository: repository, S3storage: s3}
}

func (s *ProductService) GetProductById(id int64) (*domain.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *ProductService) AddToProduct(product *domain.Product) (*domain.Product, error) {
	return s.repository.CreateProduct(product)
}
func (s *ProductService) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	return s.repository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.repository.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() []domain.Product {
	return s.repository.GetProducts()
}

func (s *ProductService) GetProduct(id int64) (*domain.Product, error) {
	return s.repository.GetProduct(id)
}

func (s *ProductService) UploadFile(location string, filename string) (string, error) {
	return s.S3storage.UploadFile(location, filename)
}
