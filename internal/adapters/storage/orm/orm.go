package orm

import (
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
