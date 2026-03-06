package ports

import "market/internal/domain"

type StorageRepository interface {
	GetAllProducts() []domain.Product
	GetProductById(productId int) (domain.Product, error)
	AddProduct(product domain.Product) (domain.Product, error)
	DeleteProductById(productId int) error
	UpdateProduct(product domain.Product) (domain.Product, error)
}
