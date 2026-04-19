package dto

import (
	"errors"
	"strings"

	"market/internal/core/domain"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (r *CreateCategoryRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func (r *CreateCategoryRequest) ToDomain() *domain.Category {
	return &domain.Category{Name: strings.TrimSpace(r.Name)}
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

func (r *UpdateCategoryRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}
