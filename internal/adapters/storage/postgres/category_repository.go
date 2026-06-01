package postgres

import (
	"database/sql"
	"errors"

	"market/internal/core/domain"
	"market/internal/core/ports"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(categoryName string) (*domain.Category, error) {
	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`

	var category domain.Category

	err := r.db.QueryRow(query, categoryName).
		Scan(&category.ID, &category.Name, &category.CreatedAt)

	if err != nil {
		var pqErr interface{ SQLState() string }
		if errors.As(err, &pqErr) && pqErr.SQLState() == "23505" {
			return nil, ports.ErrCategoryAlreadyExists
		}
		return nil, ports.ErrFailedToCreateCategory
	}

	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(category *domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories
		SET name = $1
		WHERE id = $2
		RETURNING id, name, created_at
	`

	var updated domain.Category

	err := r.db.QueryRow(query, category.Name, category.ID).
		Scan(&updated.ID, &updated.Name, &updated.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrCategoryNotFound
		}
		return nil, ports.ErrFailedToUpdateCategory
	}

	return &updated, nil
}

func (r *CategoryRepository) DeleteCategory(id int64) error {
	query := `DELETE FROM categories WHERE id = $1`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return ports.ErrFailedToDeleteCategory
	}

	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return ports.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoryRepository) GetCategory(id int64) (*domain.Category, error) {
	query := `
		SELECT id, name, created_at
		FROM categories
		WHERE id = $1
	`

	var category domain.Category

	err := r.db.QueryRow(query, id).
		Scan(&category.ID, &category.Name, &category.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetCategories() ([]domain.Category, error) {
	query := `
		SELECT id, name, created_at
		FROM categories
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *CategoryRepository) GetCategoryByName(name string) (*domain.Category, error) {
	query := `
		SELECT id, name, created_at
		FROM categories
		WHERE name = $1
	`

	var category domain.Category

	err := r.db.QueryRow(query, name).
		Scan(&category.ID, &category.Name, &category.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) ProductsByCategory(id int64) ([]domain.Product, error) {
	query := `
		SELECT id, name, price, category_id
		FROM products
		WHERE category_id = $1
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
