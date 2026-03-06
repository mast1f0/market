package memory

import (
	"errors"
	"market/internal/domain"
)

type Memory struct {
	products []domain.Product
}

func NewMemory() *Memory {
	return &Memory{
		products: make([]domain.Product, 0),
	}
}

func (m *Memory) GetAllProducts() []domain.Product {
	return m.products
}

func (m *Memory) GetProductById(productId int) (domain.Product, error) {
	for _, product := range m.products {
		if productId == product.ID {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (m *Memory) AddProduct(product domain.Product) (domain.Product, error) {
	m.products = append(m.products, product)
	return product, nil
}

func (m *Memory) DeleteProductById(productId int) error {
	for i, product := range m.products {
		if product.ID == productId {
			m.products = append(m.products[:i], m.products[i+1:]...)
		}
	}
	return nil
}

func (m *Memory) UpdateProduct(product domain.Product) (domain.Product, error) {
	for i, product := range m.products {
		if product.ID == product.ID {
			m.products[i] = product
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}
