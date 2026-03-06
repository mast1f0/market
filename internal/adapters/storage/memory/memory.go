package memory

import (
	"errors"
	"market/internal/domain"
)

type Memory struct {
	products   []domain.Product
	categories []domain.Category
}

func NewMemory() *Memory {
	return &Memory{
		products:   make([]domain.Product, 0),
		categories: make([]domain.Category, 0),
	}
}

func (m *Memory) GetAllProducts() []domain.Product {
	return m.products
}

func (m *Memory) GetProductById(productId int) (*domain.Product, error) {
	for _, product := range m.products {
		if productId == product.ID {
			return &product, nil
		}
	}
	return &domain.Product{}, errors.New("product not found")
}

func (m *Memory) AddProduct(product *domain.Product) (*domain.Product, error) {
	m.products = append(m.products, *product)
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

func (m *Memory) UpdateProduct(newProduct *domain.Product) (*domain.Product, error) {
	for i, product := range m.products {
		if newProduct.ID == product.ID {
			m.products[i] = *newProduct
			return newProduct, nil
		}
	}
	return &domain.Product{}, errors.New("product not found")
}

func (m *Memory) NewCategory(category *domain.Category) (*domain.Category, error) {
	m.categories = append(m.categories, *category)
	return category, nil
}

func (m *Memory) DeleteCategoryById(categoryId int) error {
	for i, category := range m.categories {
		if category.ID == categoryId {
			m.categories = append(m.categories[:i], m.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}

func (m *Memory) CategoryByName(name string) (*domain.Category, error) {
	for _, category := range m.categories {
		if category.Name == name {
			return &category, nil
		}
	}
	return &domain.Category{}, errors.New("category not found")
}

func (m *Memory) AddToCategory(category *domain.Category, productId int) (*domain.Category, error) {
	category.ProductId = append(category.ProductId, productId)
	return category, nil
}
