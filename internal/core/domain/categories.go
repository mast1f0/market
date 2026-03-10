package domain

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"null"`
}
