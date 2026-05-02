package orm

import (
	"market/internal/core/domain"

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

func (r *ProductRepository) GetProducts() []domain.Product {
	var products []domain.Product
	r.db.Find(&products)
	return products
}

func (r *ProductRepository) GetProduct(id int64) (*domain.Product, error) {
	var product domain.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id int64) error {
	var product domain.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&product).Error
}

func (r *ProductRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	res := r.db.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	res := r.db.Save(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}
