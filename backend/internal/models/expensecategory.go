package models

import "time"

type ExpenseCategory struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserId    uint   `json:"user_id" gorm:"not null"`
	User      User   `json:"user" gorm:"foreignKey:UserId;references:ID"`
	Name      string `json:"name" gorm:"not null"`
	Icon      string `json:"icon" gorm:"not null"`
	Color     string `json:"color" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
