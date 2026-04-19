package dto

import "market/internal/core/domain"

type UpdateCartRequest struct {
	UserId int64 `json:"user_id"`
}

func (r *UpdateCartRequest) ToDomain(cartID int64) *domain.Cart {
	return &domain.Cart{
		Id:     cartID,
		UserId: r.UserId,
	}
}
