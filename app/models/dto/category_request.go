package dto

import "dot/app/models/domain"

type CategoryCreate struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SingleCategoryResponse struct {
	Message string           `json:"message"`
	Data    CategoryResponse `json:"data"`
}

type MultipleCategoryResponse struct {
	Message string `json:"message"`
	Data    []domain.Category
}

type ErrorResponse struct {
	Message string `json:"message"`
}
