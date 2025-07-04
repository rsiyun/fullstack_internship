package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
)

type CategoryExpenseService struct {
	categoryExpenseRepo *repositories.CategoryExpenseRepo
}

func NewCategoryExpenseService(categoryExpenseRepo *repositories.CategoryExpenseRepo) *CategoryExpenseService {
	return &CategoryExpenseService{categoryExpenseRepo: categoryExpenseRepo}
}

func (s *CategoryExpenseService) GetCategoryExpenseByUserId(userId int) (*dtos.CategoryListResponse, *dtos.ErrorResponse) {
	result, err := s.categoryExpenseRepo.FindCategoryExpenseByUserId(userId)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryExpenseListResponse(result), nil
}

func (s *CategoryExpenseService) GetCategoryExpenseById(categoryId int) (*dtos.CategoryResponse, *dtos.ErrorResponse) {
	result, err := s.categoryExpenseRepo.FindCategoryExpenseById(categoryId)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryExpenseResponse(result), nil
}

func (s *CategoryExpenseService) CreateCategoryExpense(request *models.ExpenseCategory) (*dtos.CategoryResponse, *dtos.ErrorResponse) {
	result, err := s.categoryExpenseRepo.CreateExpenseCategory(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryExpenseResponse(result), nil
}

func (s *CategoryExpenseService) UpdateCategoryExpense(request *models.ExpenseCategory) (*dtos.CategoryResponse, *dtos.ErrorResponse) {
	existingCategory, err := s.categoryExpenseRepo.FindCategoryExpenseById(int(request.ID))
	if err != nil {
		return nil, err
	}

	existingCategory.Name = request.Name
	existingCategory.Color = request.Color
	existingCategory.Icon = request.Icon

	result, err := s.categoryExpenseRepo.UpdateExpenseCategory(existingCategory)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryExpenseResponse(result), nil
}
