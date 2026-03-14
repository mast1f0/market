package domain

import "time"

type Product struct {
	ID          uint      `gorm:"primarykey" json:"id,omitempty"`
	Name        string    `gorm:"not null;index" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null;check:price > 0" json:"price"`
	CategoryID  uint      `gorm:"not null;index" json:"category_id"`
	Stock       int       `gorm:"default:0;check:stock >= 0" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
