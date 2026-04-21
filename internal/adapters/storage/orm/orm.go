package orm

import (
	"errors"
	"fmt"
	"market/internal/core/domain"
	"market/internal/engine/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage() *Storage {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB_USER, cfg.DB_PASSWORD,
		cfg.DB_HOST, cfg.DB_PORT,
		cfg.DB_NAME,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < 10; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}
	}

	return &Storage{
		db: db,
	}
}

func (s *Storage) GetProducts() []domain.Product {
	var products []domain.Product
	s.db.Find(&products)
	return products
}

func (s *Storage) GetProduct(id int64) (*domain.Product, error) {
	var product domain.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Storage) DeleteProduct(id int64) error {
	var product domain.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return err
	}
	return s.db.Delete(&product).Error
}

func (s *Storage) CreateProduct(product *domain.Product) (*domain.Product, error) {
	res := s.db.Create(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (s *Storage) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	res := s.db.Save(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (s *Storage) ProductsByCategory(id int) ([]domain.Product, error) {
	var products []domain.Product
	res := s.db.Where("category_id = ?", id).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

func (s *Storage) CreateCategory(category *domain.Category) (*domain.Category, error) {
	res := s.db.Create(&category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}

func (s *Storage) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	res := s.db.Save(&category)
	if res.Error != nil {
		return nil, errors.New("Не удалось обновить")
	}
	return category, nil
}

func (s *Storage) DeleteCategory(id int64) error {
	var category domain.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return err
	}
	return s.db.Delete(&category).Error
}

func (s *Storage) GetCategory(id int64) (*domain.Category, error) {
	var category domain.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *Storage) GetCategories() []domain.Category {
	var categories = make([]domain.Category, 0)
	s.db.Find(&categories)
	return categories
}

func (s *Storage) GetCategoryByName(name string) *domain.Category {
	var category domain.Category
	s.db.Where("name = ?", name).First(&category)
	return &category
}

func (s *Storage) CreateCart(cart *domain.Cart) (*domain.Cart, error) {
	res := s.db.Create(&cart)
	if res.Error != nil {
		return nil, errors.New("Не удалось создать корзину")
	}
	return cart, nil
}

func (s *Storage) UpdateCart(cart *domain.Cart) (*domain.Cart, error) {
	res := s.db.Save(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return cart, nil
}

func (s *Storage) DeleteCart(userID int64) error {
	var cart domain.Cart
	if err := s.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}
	return s.db.Delete(&cart).Error
}

func (s *Storage) GetCart(userID int64) (*domain.Cart, error) {
	var cart domain.Cart
	if err := s.db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (s *Storage) GetCartByID(id int64) (*domain.Cart, error) {
	var cart domain.Cart
	if err := s.db.First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (s *Storage) AddCartItem(cartItems *domain.CartItems) (*domain.CartItems, error) {
	res := s.db.Create(&cartItems)
	if res.Error != nil {
		return nil, res.Error
	}
	return cartItems, nil
}

func (s *Storage) DeleteCartItem(id int64) error {
	var item domain.CartItems
	if err := s.db.First(&item, id).Error; err != nil {
		return err
	}
	return s.db.Delete(&item).Error
}

func (s *Storage) GetCartItems(id int64) (*domain.CartItems, error) {
	var item domain.CartItems
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *Storage) UpdateCartItem(cartItems *domain.CartItems) (*domain.CartItems, error) {
	res := s.db.Save(&cartItems)
	if res.Error != nil {
		return nil, res.Error
	}
	return cartItems, nil
}
