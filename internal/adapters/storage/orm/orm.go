package orm

import (
	"errors"
	"fmt"
	"market/internal/config"
	"market/internal/core/domain"

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
		panic("failed to connect database")
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

func (s *Storage) GetProduct(id int) *domain.Product {
	var product domain.Product
	s.db.First(&product, id)
	return &product
}

func (s *Storage) DeleteProduct(id int) error {
	var product domain.Product
	s.db.First(&product, id)
	res := s.db.Delete(&product)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
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
		return nil, res.Error
	}
	return category, nil
}

func (s *Storage) DeleteCategory(id int) error {
	var category domain.Category
	s.db.First(&category, id)
	res := s.db.Delete(&category)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Storage) GetCategory(id int) *domain.Category {
	var category domain.Category
	s.db.First(&category, id)
	return &category
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

func (s *Storage) DeleteCart(id int) error {
	var cart domain.Cart
	s.db.First(&cart, id)
	res := s.db.Delete(&cart)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Storage) GetCart(id int) (*domain.Cart, error) {
	var cart domain.Cart
	res := s.db.First(&cart, id)
	if res.Error != nil {
		return nil, errors.New("This cart is not exist")
	}
	return &cart, nil
}
