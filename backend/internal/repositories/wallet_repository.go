package repositories

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) FindWalletByUserId(userID int) ([]models.Wallet, *dtos.ErrorResponse) {
	var data []models.Wallet
	result := r.db.Where("user_id = ?", userID).Find(&data)
	if result.Error != nil {
		return nil, dtos.NewErrorResponse("Failed to retrieve wallets", 500, "database error")
	}
	return data, nil
}

func (r *WalletRepository) FindWalletByID(walletID int) (models.Wallet, *dtos.ErrorResponse) {
	var data models.Wallet
	err := r.db.Where("id = ?", walletID).First(&data).Error
	if err != nil {
		return models.Wallet{}, dtos.NewErrorResponse("Wallet not found", 404, "wallet not found")
	}
	return data, nil
}

func (r *WalletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, *dtos.ErrorResponse) {
	result := r.db.Create(wallet)
	if result.RowsAffected > 0 {
		return wallet, nil
	}
	return nil, dtos.NewErrorResponse("Failed to create wallet", 500, "database error")

}

func (r *WalletRepository) UpdateWallet(wallet *models.Wallet) (*models.Wallet, *dtos.ErrorResponse) {

	// mencari data wallet terlebih dahulu
	var existingwallet models.Wallet
	err := r.db.Where("id = ?", wallet.ID).First(&existingwallet).Error
	if err != nil {
		return nil, dtos.NewErrorResponse("Wallet not found", 404, "wallet not found")
	}
	// update data wallet
	existingwallet.Balance = wallet.Balance
	existingwallet.Name = wallet.Name

	// masukkan data
	result := r.db.Save(&existingwallet)
	if result.RowsAffected > 0 {
		return wallet, nil
	}
	return nil, dtos.NewErrorResponse("Failed to update wallet", 500, "database error")

}
