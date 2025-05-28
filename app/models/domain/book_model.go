package domain

import "time"

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Writer      string `gorm:"not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CategoryID  uint
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
