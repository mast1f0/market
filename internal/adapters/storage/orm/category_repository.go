package orm

import (
	"errors"
	"market/internal/core/domain"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) CreateCategory(categoryName string) (*domain.Category, error) {
	var category = &domain.Category{
		Name: categoryName,
	}
	res := r.db.Create(category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	res := r.db.Save(&category)
	if res.Error != nil {
		return nil, errors.New("не удалось обновить")
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategory(id int64) error {
	var category domain.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}

func (r *CategoryRepository) GetCategory(id int64) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetCategories() []domain.Category {
	var categories = make([]domain.Category, 0)
	r.db.Find(&categories)
	return categories
}

func (r *CategoryRepository) GetCategoryByName(name string) *domain.Category {
	var category domain.Category
	r.db.Where("name = ?", name).First(&category)
	return &category
}

func (r *CategoryRepository) ProductsByCategory(id int) ([]domain.Product, error) {
	var products []domain.Product
	res := r.db.Where("category_id = ?", id).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}
