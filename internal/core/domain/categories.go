package domain

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id,omitempty"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `gorm:"null" json:"created_at,omitempty"`
}
