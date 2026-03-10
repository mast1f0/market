package domain

import "time"

type Product struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"not null;index"`
	Description   string
	Price         float64 `gorm:"not null;check:price > 0"`
	CategoryID    uint    `gorm:"not null;index"`
	Category      Category
	StockQuantity int `gorm:"default:0;check:stock_quantity >= 0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
