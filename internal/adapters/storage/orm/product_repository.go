package orm

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(id int64) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id int64) error {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrNotFound
		}
		return err
	}
	err = r.db.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	err := r.db.Create(product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ports.ErrAlreadyExists
		}
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	err := r.db.Save(product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ports.ErrAlreadyExists
		}
		return nil, err
	}
	return product, nil
}
