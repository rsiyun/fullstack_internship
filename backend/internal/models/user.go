package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
