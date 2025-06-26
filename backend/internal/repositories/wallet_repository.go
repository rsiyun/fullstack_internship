package repositories

import (
	"api-money-management/internal/models"
	"errors"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) FindWalletByUserId(userID int) ([]models.Wallet, error) {
	var data []models.Wallet
	result := r.db.Where("user_id = ?", userID).Find(&data)
	if result.Error != nil {
		return nil, errors.New("Failed to find wallet")
	}
	return data, nil
}

func (r *WalletRepository) FindWalletByID(walletID int) (models.Wallet, error) {
	var data models.Wallet
	err := r.db.Where("id = ?", walletID).First(&data).Error
	if err != nil {
		return models.Wallet{}, errors.New("failed to find wallet")
	}
	return data, err
}

func (r *WalletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	result := r.db.Create(wallet)
	if result.RowsAffected > 0 {
		return wallet, nil
	}
	return nil, errors.New("")

}

func (r *WalletRepository) UpdateWallet(wallet *models.Wallet) (*models.Wallet, error) {

	// mencari data wallet terlebih dahulu
	var existingwallet models.Wallet
	err := r.db.Where("id = ?", wallet.ID).First(&existingwallet).Error
	if err != nil {
		return &models.Wallet{}, errors.New("failed to find wallet")
	}
	// update data wallet
	existingwallet.Balance = wallet.Balance
	existingwallet.Name = wallet.Name

	// masukkan data
	result := r.db.Save(&existingwallet)
	if result.RowsAffected > 0 {
		return wallet, nil
	}
	return nil, errors.New("")

}
