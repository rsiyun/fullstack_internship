package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	UserId  uint    `json:"user_id" gorm:"not null"`
	User    User    `json:"user" gorm:"foreignKey:UserId;references:ID"`
	Name    string  `json:"name" gorm:"not null"`
	Balance float64 `json:"balance" gorm:"not null;default:0"`
	gorm.Model
}
