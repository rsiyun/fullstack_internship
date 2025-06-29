package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
)

type CategoryIncomeService struct {
	categoryIncomeRepo *repositories.CategoryIncomeRepo
}

func NewCategoryIncomeService(categoryIncomeRepo *repositories.CategoryIncomeRepo) *CategoryIncomeService {
	return &CategoryIncomeService{categoryIncomeRepo: categoryIncomeRepo}
}

func (s *CategoryIncomeService) GetCategoryIncomeByUserId(userId int) (*dtos.CategoryIncomeListResponse, *dtos.ErrorResponse) {
	result, err := s.categoryIncomeRepo.FindCategoryIncomeByUserId(userId)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryIncomeListResponse(result), nil
}
func (s *CategoryIncomeService) GetCategoryIncomeById(categoryId int) (*dtos.CategoryIncomeResponse, *dtos.ErrorResponse) {
	result, err := s.categoryIncomeRepo.FindCategoryIncomeById(categoryId)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryIncomeResponse(result), nil

}
func (s *CategoryIncomeService) CreateCategoryIncome(request *models.IncomeCategory) (*dtos.CategoryIncomeResponse, *dtos.ErrorResponse) {
	result, err := s.categoryIncomeRepo.CreateIncomeCategory(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryIncomeResponse(result), nil
}
func (s *CategoryIncomeService) UpdateCategoryIncome(request *models.IncomeCategory) (*dtos.CategoryIncomeResponse, *dtos.ErrorResponse) {
	existingCategory, err := s.categoryIncomeRepo.FindCategoryIncomeById(int(request.ID))
	if err != nil {
		return nil, err
	}

	existingCategory.Name = request.Name
	existingCategory.Color = request.Color
	existingCategory.Icon = request.Icon

	result, err := s.categoryIncomeRepo.UpdateIncomeCategory(existingCategory)
	if err != nil {
		return nil, err
	}
	return dtos.ToCategoryIncomeResponse(result), nil
}
