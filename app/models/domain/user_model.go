package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Active    bool   `gorm:"default:true;not null" json:"active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
