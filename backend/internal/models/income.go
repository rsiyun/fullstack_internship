package models

import "time"

type Income struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	UserId           uint           `json:"user_id" gorm:"not null"`
	User             User           `json:"user" gorm:"foreignKey:UserId;references:ID"`
	WalletId         uint           `json:"wallet_id" gorm:"not null"`
	Wallet           Wallet         `json:"wallet" gorm:"foreignKey:WalletId;references:ID"`
	IncomeCategoryId uint           `json:"income_category_id" gorm:"not null"`
	IncomeCategory   IncomeCategory `json:"income_category" gorm:"foreignKey:IncomeCategoryId;references:ID"`
	Amount           float64        `json:"amount" gorm:"not null"`
	Description      string         `json:"description" gorm:"not null"`
	CreatedAt        time.Time
}
