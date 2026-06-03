package mock

import "market/internal/core/domain"

type ProductRepoMock struct {
	GetProductByIdFunc func(id int64) (*domain.Product, error)
	DeleteProductFunc  func(id int64) error
	UpdateProductFunc  func(p *domain.Product) (*domain.Product, error)
}

func (r *ProductRepoMock) GetProducts() ([]domain.Product, error) {
	return nil, nil
}

func (r *ProductRepoMock) GetProductById(id int64) (*domain.Product, error) {
	return r.GetProductByIdFunc(id)
}

func (r *ProductRepoMock) DeleteProduct(id int64) error {
	return r.DeleteProductFunc(id)
}
func (r *ProductRepoMock) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return product, nil
}
func (r *ProductRepoMock) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	return r.UpdateProductFunc(product)
}
