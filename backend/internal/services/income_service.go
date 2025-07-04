package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
)

type IncomeService struct {
	incomeRepository         *repositories.IncomeRepository
	walletRepository         *repositories.WalletRepository
	incomeCategoryRepository *repositories.CategoryIncomeRepo
}

func NewIncomeService(incomeRepo *repositories.IncomeRepository, walletRepo *repositories.WalletRepository, incomeCategoryRepo *repositories.CategoryIncomeRepo) *IncomeService {
	return &IncomeService{
		incomeRepository:         incomeRepo,
		walletRepository:         walletRepo,
		incomeCategoryRepository: incomeCategoryRepo,
	}
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
	wallet, err := s.walletRepository.FindWalletByID(int(request.WalletId))
	if err != nil {
		return nil, err
	}
	_, err = s.incomeCategoryRepository.FindCategoryIncomeById(int(request.IncomeCategoryId))
	if err != nil {
		return nil, err
	}
	wallet.Balance = wallet.Balance + request.Amount
	result, err := s.incomeRepository.CreateIncome(request, wallet)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeResponse(result), nil
}
func (s *IncomeService) UpdateIncome(request *models.Income) (*dtos.IncomeResponse, *dtos.ErrorResponse) {
	existingIncome, err := s.incomeRepository.FindIncomeById(int(request.ID))
	if err != nil {
		return nil, err
	}
	_, err = s.incomeCategoryRepository.FindCategoryIncomeById(int(request.IncomeCategoryId))
	if err != nil {
		return nil, err
	}
	var oldWallet, newWallet *models.Wallet
	if existingIncome.WalletId != request.WalletId {
		oldWallet, err = s.walletRepository.FindWalletByID(int(existingIncome.WalletId))
		if err != nil {
			return nil, err
		}
		oldWallet.Balance -= existingIncome.Amount
		newWallet, err = s.walletRepository.FindWalletByID(int(request.WalletId))
		if err != nil {
			return nil, err
		}
		newWallet.Balance += request.Amount
	} else if existingIncome.Amount != request.Amount {
		oldWallet, err = s.walletRepository.FindWalletByID(int(existingIncome.WalletId))
		if err != nil {
			return nil, err
		}
		oldWallet.Balance = oldWallet.Balance - existingIncome.Amount + request.Amount
		newWallet = oldWallet
	}

	existingIncome.WalletId = request.WalletId
	existingIncome.IncomeCategoryId = request.IncomeCategoryId
	existingIncome.Amount = request.Amount
	existingIncome.Description = request.Description

	result, err := s.incomeRepository.UpdateIncome(existingIncome, oldWallet, newWallet)
	if err != nil {
		return nil, err
	}
	return dtos.ToIncomeResponse(result), nil

}
