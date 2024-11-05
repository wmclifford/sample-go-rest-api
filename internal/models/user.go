package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:100;not null"` // No json tag to exclude this field
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
