package dtos

import (
	"api-money-management/internal/models"
	"time"
)

type RequestCategoryIncome struct {
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Icon   string `json:"icon" validate:"required"`
	Color  string `json:"color" validate:"required"`
}

type CategoryIncomeResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryIncomeListResponse struct {
	Categories []CategoryIncomeResponse `json:"categories"`
}

func ToCategoryIncomeResponse(incomeCategory *models.IncomeCategory) *CategoryIncomeResponse {
	return &CategoryIncomeResponse{
		ID:        incomeCategory.ID,
		UserId:    incomeCategory.UserId,
		Name:      incomeCategory.Name,
		Icon:      incomeCategory.Icon,
		Color:     incomeCategory.Color,
		CreatedAt: incomeCategory.CreatedAt,
		UpdatedAt: incomeCategory.UpdatedAt,
	}
}
func ToCategoryIncomeListResponse(categories []models.IncomeCategory) *CategoryIncomeListResponse {
	var categoryResponse []CategoryIncomeResponse
	for _, wallet := range categories {
		categoryResponse = append(categoryResponse, *ToCategoryIncomeResponse(&wallet))
	}
	return &CategoryIncomeListResponse{
		Categories: categoryResponse,
	}
}
