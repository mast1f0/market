//for tests, fixed later

package memory

import (
	"errors"
	"market/internal/core/domain"
	"sync"
)

type Storage struct {
	mu         sync.RWMutex
	products   []domain.Product
	categories []domain.Category
	cart       []domain.Cart
	cartItems  []domain.CartItem
}

func NewMemory() *Storage {
	return &Storage{
		products:   make([]domain.Product, 0),
		categories: make([]domain.Category, 0),
	}
}
func (s *Storage) GetProducts() []domain.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.products
}

func (s *Storage) GetProductById(id int64) (*domain.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (s *Storage) DeleteProduct(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products = append(s.products[:id], s.products[id+1:]...)
	return nil
}

func (s *Storage) CreateProduct(product *domain.Product) (*domain.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products = append(s.products, *product)
	return product, nil
}

func (s *Storage) UpdateProduct(newProduct *domain.Product) (*domain.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, product := range s.products {
		if product.ID == newProduct.ID {
			product.Name = newProduct.Name
			product.Description = newProduct.Description
			product.Stock = newProduct.Stock
			product.CategoryID = newProduct.CategoryID
			return &product, nil
		}
	}
	return nil, nil
}

func (s *Storage) ProductsByCategory(id int) ([]domain.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var products []domain.Product
	for _, product := range s.products {
		for _, category := range s.categories {
			if product.ID == category.ID {
				products = append(products, product)
			}
		}
	}
	return products, nil
}

func (s *Storage) CreateCategory(categoryName string) (*domain.Category, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var category = domain.Category{
		Name: categoryName,
	}
	s.categories = append(s.categories, category)
	return &category, nil

}

func (s *Storage) UpdateCategory(newCategory *domain.Category) (*domain.Category, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, category := range s.categories {
		if category.ID == category.ID {
			category.Name = newCategory.Name
			return &category, nil
		}
	}
	return nil, errors.New("Не удалось обновить")
}

func (s *Storage) DeleteCategory(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for index, category := range s.categories {
		if category.ID == id {
			s.categories = append(s.categories[:index], s.categories[index+1:]...)
			return nil
		}
	}
	return errors.New("Cant find category")

}

func (s *Storage) GetCategory(id int64) (*domain.Category, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, category := range s.categories {
		if category.ID == id {
			return &category, nil
		}
	}
	return nil, errors.New("Cant find category")
}

func (s *Storage) GetCategories() []domain.Category {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.categories
}

func (s *Storage) GetCategoryByName(name string) (*domain.Category, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, category := range s.categories {
		if category.Name == name {
			return &category, nil
		}
	}
	return nil, errors.New("Cant find category")
}

func (s *Storage) CreateCart(id int64) (*domain.Cart, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var cart domain.Cart
	s.cartItems = append(s.cartItems, domain.CartItem{
		ProductID: id,
		Quantity:  1,
	})
	s.cart = append(s.cart, cart)
	return &cart, nil
}

// TODO fix
func (s *Storage) GetCartWithItems(userID int64) (*domain.Cart, error) {
	//s.mu.RLock()
	//defer s.mu.RUnlock()
	//var cart domain.Cart
	//for _, item := range s.cartItems {
	//	if item.ProductID == userID {
	//		re
	//	}
	//}
	return nil, nil
}

func (s *Storage) FindCartItem(cartID, productID int64) (*domain.CartItem, error) {
	return nil, nil
}

func (s *Storage) DeleteCartItem(userId int64, itemId int64) error {
	return nil
}

func (s *Storage) UpdateCartItem(itemId int64, quantity int) (*domain.CartItem, error) {
	return nil, nil
}

func (s *Storage) AddCartItem(userId int64, cartItem *domain.CartItem) (*domain.CartItem, error) {
	return nil, nil
}
