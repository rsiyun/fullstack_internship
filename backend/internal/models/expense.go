package models

import "time"

type Expense struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	UserId            uint            `json:"user_id" gorm:"not null"`
	User              User            `json:"user" gorm:"foreignKey:UserId;references:ID"`
	WalletId          uint            `json:"wallet_id" gorm:"not null"`
	Wallet            Wallet          `json:"wallet" gorm:"foreignKey:WalletId;references:ID"`
	ExpenseCategoryId uint            `json:"expense_category_id" gorm:"not null"`
	ExpenseCategory   ExpenseCategory `json:"expense_category" gorm:"foreignKey:ExpenseCategoryId;references:ID"`
	Amount            float64         `json:"amount" gorm:"not null"`
	Description       string          `json:"description" gorm:"not null"`
	CreatedAt         time.Time
}
