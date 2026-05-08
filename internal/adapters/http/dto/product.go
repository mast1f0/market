package dto

import (
	"errors"
	"strings"

	"market/internal/core/domain"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int64   `json:"category_id"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock"`
}

func (r *CreateProductRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	if r.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if r.CategoryID == 0 {
		return errors.New("category_id is required")
	}
	if r.Stock < 0 {
		return errors.New("stock must be non-negative")
	}
	return nil
}

func (r *CreateProductRequest) ToDomain(ownerID int64) *domain.Product {
	return &domain.Product{
		OwnerID:     ownerID,
		Name:        strings.TrimSpace(r.Name),
		Description: r.Description,
		Price:       r.Price,
		CategoryID:  r.CategoryID,
		ImageURL:    strings.TrimSpace(r.ImageURL),
		Stock:       r.Stock,
	}
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int64   `json:"category_id"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock"`
}

func (r *UpdateProductRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	if r.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if r.CategoryID == 0 {
		return errors.New("category_id is required")
	}
	if r.Stock < 0 {
		return errors.New("stock must be non-negative")
	}
	return nil
}

func (r *UpdateProductRequest) ApplyTo(p *domain.Product) {
	p.Name = strings.TrimSpace(r.Name)
	p.Description = r.Description
	p.Price = r.Price
	p.CategoryID = r.CategoryID
	p.ImageURL = strings.TrimSpace(r.ImageURL)
	p.Stock = r.Stock
}
