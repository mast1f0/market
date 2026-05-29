package postgres

import (
	"database/sql"
	"errors"
	"market/internal/core/domain"
	"market/internal/core/ports"

	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, owner_id, name, description, price, category_id, stock, image_url, created_at FROM products")
	if err != nil {
		return nil, ports.ErrNotFound
	}
	defer rows.Close()
	var products []domain.Product

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.ID, &product.OwnerID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Stock,
			&product.ImageURL, &product.CreatedAt)
		if err != nil {
			return nil, ports.ErrNotFound
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, ports.ErrNotFound
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(id int64) (*domain.Product, error) {
	query := `SELECT * FROM products WHERE id = $1;`
	row := r.db.QueryRow(query, id)
	var product domain.Product
	err := row.Scan(&product.ID, &product.OwnerID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Stock,
		&product.ImageURL, &product.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrNotFound
		}
		return nil, ports.ErrFailedToLoadProduct
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id int64) error {
	query := `DELETE FROM products WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return ports.ErrNotFound
	}
	return nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := `INSERT INTO products(id, owner_id, name, description, price, category_id, stock, image_url, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	err := r.db.QueryRow(query, product.ID, product.OwnerID, product.Name, product.Description, product.Price, product.CategoryID, product.Stock, product.ImageURL, product.CreatedAt).Scan(&product.ID)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, ports.ErrAlreadyExists
			}
		}
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	res, err := r.db.Exec(`
		UPDATE products
		SET name = $1,
			description = $2,
			price = $3
		WHERE id = $4
	`,
		product.Name,
		product.Description,
		product.Price,
		product.ID,
	)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, ports.ErrAlreadyExists
			}
		}

		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, ports.ErrNotFound
	}

	return product, nil
}
