package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
)

type IncomeService struct {
	incomeRepository *repositories.IncomeRepository
}

func NewIncomeService(incomeRepo *repositories.IncomeRepository) *IncomeService {
	return &IncomeService{incomeRepository: incomeRepo}
}

func (s *IncomeService) GetIncomeByUserId(userID int) (*dtos.IncomeListResponse, *dtos.ErrorResponse) {
	result, err := s.incomeRepository.FindIncomeByUserId(userID)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeListResponse(result), nil
}
func (s *IncomeService) GetIncomeIncomeById(id int) (*dtos.IncomeResponse, *dtos.ErrorResponse) {
	result, err := s.incomeRepository.FindIncomeById(id)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeResponse(result), nil
}

func (s *IncomeService) CreateIncome(request *models.Income) (*dtos.IncomeResponse, *dtos.ErrorResponse) {
	result, err := s.incomeRepository.CreateIncome(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeResponse(result), nil
}
func (s *IncomeService) UpdateIncome(request *models.Income) (*dtos.IncomeResponse, *dtos.ErrorResponse) {
	existingCategory, err := s.incomeRepository.FindIncomeById(int(request.ID))
	if err != nil {
		return nil, err
	}
	existingCategory.WalletId = request.WalletId
	existingCategory.IncomeCategoryId = request.IncomeCategoryId
	existingCategory.Amount = request.Amount
	existingCategory.Description = request.Description

	result, err := s.incomeRepository.UpdateIncome(existingCategory)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeResponse(result), nil

}
