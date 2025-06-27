package services

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"
	"api-money-management/internal/repositories"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
}

func NewWalletService(walletRepo *repositories.WalletRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo}
}

func (s *WalletService) GetWalletByUserID(userID int) (*dtos.WalletListResponse, *dtos.ErrorResponse) {
	wallets, err := s.walletRepo.FindWalletByUserId(userID)
	if err != nil {
		return nil, err
	}
	return dtos.ToWalletListResponse(wallets), nil
}

// func GetWalletByID(walletID int) (models.Wallet, *dtos.ErrorResponse) {

// }

func (s *WalletService) CreateWallet(request *models.Wallet) (*dtos.WalletResponse, *dtos.ErrorResponse) {
	data, err := s.walletRepo.CreateWallet(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToWalletResponse(data), nil
}
