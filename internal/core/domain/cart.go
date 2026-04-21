package domain

import "time"

type Cart struct {
	ID        int64      `json:"id"`
	UserID    *int64     `json:"user_id"` //приколы с NULL
	Status    string     `json:"status"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
