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

func (s *WalletService) GetWalletByID(walletID int) (*dtos.WalletResponse, *dtos.ErrorResponse) {
	wallet, err := s.walletRepo.FindWalletByID(walletID)
	if err != nil {
		return nil, err
	}
	return dtos.ToWalletResponse(wallet), nil
}

func (s *WalletService) CreateWallet(request *models.Wallet) (*dtos.WalletResponse, *dtos.ErrorResponse) {
	data, err := s.walletRepo.CreateWallet(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToWalletResponse(data), nil
}
func (s *WalletService) UpdateWallet(request *models.Wallet) (*dtos.WalletResponse, *dtos.ErrorResponse) {
	existingwallet, err := s.walletRepo.FindWalletByID(int(request.ID))
	if err != nil {
		return nil, err
	}

	existingwallet.Balance = request.Balance
	existingwallet.Name = request.Name

	data, err := s.walletRepo.UpdateWallet(request)
	if err != nil {
		return nil, err
	}
	return dtos.ToWalletResponse(data), nil
}
