package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description,omitempty"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
