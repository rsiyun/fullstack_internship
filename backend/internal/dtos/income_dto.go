package dtos

import (
	"api-money-management/internal/models"
	"time"
)

type RequestIncome struct {
	WalletId         int     `json:"wallet_id" validate:"required"`
	UserId           int     `json:"user_id" validate:"required"`
	IncomeCategoryId int     `json:"income_category_id" validate:"required"`
	Amount           float64 `json:"amount" validate:"required"`
	Description      string  `json:"description" validate:"required"`
}

type IncomeResponse struct {
	ID               uint      `json:"id"`
	UserId           uint      `json:"user_id"`
	WalletId         uint      `json:"wallet_id"`
	IncomeCategoryId uint      `json:"income_category_id"`
	Amount           float64   `json:"amount"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type IncomeListResponse struct {
	Incomes []IncomeResponse
}

func ToIncomeResponse(income *models.Income) *IncomeResponse {
	return &IncomeResponse{
		ID:               income.ID,
		UserId:           income.UserId,
		WalletId:         income.WalletId,
		IncomeCategoryId: income.IncomeCategoryId,
		Amount:           income.Amount,
		Description:      income.Description,
		CreatedAt:        income.CreatedAt,
		UpdatedAt:        income.UpdatedAt,
	}
}
func ToIncomeListResponse(incomes []models.Income) *IncomeListResponse {
	var incomeResponse []IncomeResponse
	for _, income := range incomes {
		incomeResponse = append(incomeResponse, *ToIncomeResponse(&income))
	}
	return &IncomeListResponse{
		Incomes: incomeResponse,
	}
}
