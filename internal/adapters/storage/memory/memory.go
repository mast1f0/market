//for tests, fixed later

package memory

//import (
//	"errors"
//	domain2 "market/internal/core/domain"
//)
//
//type Memory struct {
//	products   []domain2.Product
//	categories []domain2.Category
//}
//
//func NewMemory() *Memory {
//	return &Memory{
//		products:   make([]domain2.Product, 0),
//		categories: make([]domain2.Category, 0),
//	}
//}
//
//func (m *Memory) GetAllProducts() []domain2.Product {
//	return m.products
//}
//
//func (m *Memory) GetProductById(productId int) (*domain2.Product, error) {
//	for _, product := range m.products {
//		if productId == product.ID {
//			return &product, nil
//		}
//	}
//	return &domain2.Product{}, errors.New("product not found")
//}
//
//func (m *Memory) AddProduct(product *domain2.Product) (*domain2.Product, error) {
//	m.products = append(m.products, *product)
//	return product, nil
//}
//
//func (m *Memory) DeleteProductById(productId int) error {
//	for i, product := range m.products {
//		if product.ID == productId {
//			m.products = append(m.products[:i], m.products[i+1:]...)
//		}
//	}
//	return nil
//}
//
//func (m *Memory) UpdateProduct(newProduct *domain2.Product) (*domain2.Product, error) {
//	for i, product := range m.products {
//		if newProduct.ID == product.ID {
//			m.products[i] = *newProduct
//			return newProduct, nil
//		}
//	}
//	return &domain2.Product{}, errors.New("product not found")
//}
//
//func (m *Memory) NewCategory(category *domain2.Category) (*domain2.Category, error) {
//	m.categories = append(m.categories, *category)
//	return category, nil
//}
//
//func (m *Memory) DeleteCategoryById(categoryId int) error {
//	for i, category := range m.categories {
//		if category.ID == categoryId {
//			m.categories = append(m.categories[:i], m.categories[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("category not found")
//}
//
//func (m *Memory) CategoryById(id int) (*domain2.Category, error) {
//	for _, category := range m.categories {
//		if category.ID == id {
//			return &category, nil
//		}
//	}
//	return &domain2.Category{}, errors.New("category not found")
//}
//
//func (m *Memory) AddToCategory(category *domain2.Category, productId int) (*domain2.Category, error) {
//	category.ProductId = append(category.ProductId, productId)
//	return category, nil
//}
//
//func (m *Memory) GetAll() []domain2.Category {
//	return m.categories
//}
