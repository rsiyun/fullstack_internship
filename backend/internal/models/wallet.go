package models

type Wallet struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id" gorm:"not null"`
	User   User `json:"user" gorm:"foreignKey:UserId;references:ID"`
}
