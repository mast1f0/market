package orm

import (
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"

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
	if err := r.db.Create(category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ports.ErrCategoryExists
		}
		return nil, ports.ErrFailedToCreateCategory
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	if err := r.db.Save(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrCategoryNotFound
		}
		return nil, ports.ErrFailedToUpdateCategory
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategory(id int64) error {
	var category domain.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ports.ErrCategoryNotFound
		}
		return ports.ErrFailedToDeleteCategory
	}
	if err := r.db.Delete(&category).Error; err != nil {
		return ports.ErrFailedToDeleteCategory
	}
	return nil
}

func (r *CategoryRepository) GetCategory(id int64) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, ports.ErrCategoryNotFound
	}
	return &category, nil
}

func (r *CategoryRepository) GetCategories() []domain.Category {
	var categories = make([]domain.Category, 0)
	r.db.Find(&categories)
	return categories
}

func (r *CategoryRepository) GetCategoryByName(name string) (*domain.Category, error) {
	var category domain.Category
	if err := r.db.First(&category, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ports.ErrCategoryNotFound
		}
	}
	return &category, nil
}

func (r *CategoryRepository) ProductsByCategory(id int64) ([]domain.Product, error) {
	var products []domain.Product
	res := r.db.Where("category_id = ?", id).Find(&products)
	if res.Error != nil {
		return nil, ports.ErrCategoryNotFound
	}
	return products, nil
}
