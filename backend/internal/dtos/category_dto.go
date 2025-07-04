package dtos

import (
	"api-money-management/internal/models"
	"time"
)

type RequestCategory struct {
	Name  string `json:"name" validate:"required"`
	Icon  string `json:"icon" validate:"required"`
	Color string `json:"color" validate:"required"`
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

func ToCategoryIncomeResponse(incomeCategory *models.IncomeCategory) *CategoryResponse {
	return &CategoryResponse{
		ID:        incomeCategory.ID,
		UserId:    incomeCategory.UserId,
		Name:      incomeCategory.Name,
		Icon:      incomeCategory.Icon,
		Color:     incomeCategory.Color,
		CreatedAt: incomeCategory.CreatedAt,
		UpdatedAt: incomeCategory.UpdatedAt,
	}
}
func ToCategoryIncomeListResponse(categories []models.IncomeCategory) *CategoryListResponse {
	var categoryResponse []CategoryResponse
	for _, wallet := range categories {
		categoryResponse = append(categoryResponse, *ToCategoryIncomeResponse(&wallet))
	}
	return &CategoryListResponse{
		Categories: categoryResponse,
	}
}

func ToCategoryExpenseResponse(incomeCategory *models.ExpenseCategory) *CategoryResponse {
	return &CategoryResponse{
		ID:        incomeCategory.ID,
		UserId:    incomeCategory.UserId,
		Name:      incomeCategory.Name,
		Icon:      incomeCategory.Icon,
		Color:     incomeCategory.Color,
		CreatedAt: incomeCategory.CreatedAt,
		UpdatedAt: incomeCategory.UpdatedAt,
	}
}
func ToCategoryExpenseListResponse(categories []models.ExpenseCategory) *CategoryListResponse {
	var categoryResponse []CategoryResponse
	for _, wallet := range categories {
		categoryResponse = append(categoryResponse, *ToCategoryExpenseResponse(&wallet))
	}
	return &CategoryListResponse{
		Categories: categoryResponse,
	}
}
