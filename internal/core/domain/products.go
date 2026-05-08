package domain

import "time"

type Product struct {
	ID          int64     `gorm:"primarykey" json:"id,omitempty"`
	OwnerID     int64     `gorm:"not null" json:"owner_id,omitempty"`
	Name        string    `gorm:"not null;index" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null;check:price > 0" json:"price"`
	CategoryID  int64     `gorm:"not null;index" json:"category_id"`
	ImageURL    string    `json:"image_url"`
	Stock       int       `gorm:"default:0;check:stock >= 0" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
