package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	gorm.Model
}
