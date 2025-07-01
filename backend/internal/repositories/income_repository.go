package repositories

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"

	"gorm.io/gorm"
)

type IncomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) *IncomeRepository {
	return &IncomeRepository{
		db: db,
	}
}

func (r *IncomeRepository) FindIncomeByUserId(userID int) ([]models.Income, *dtos.ErrorResponse) {
	var data []models.Income
	result := r.db.Where("user_id = ?", userID).Find(&data)
	if result.Error != nil {
		return nil, dtos.NewErrorResponse("Failed to retrieve Category", 500, "database error")
	}
	return data, nil
}
func (r *IncomeRepository) FindIncomeById(incomeID int) (*models.Income, *dtos.ErrorResponse) {
	var data models.Income
	err := r.db.Where("id = ?", incomeID).First(&data).Error
	if err != nil {
		return nil, dtos.NewErrorResponse("Category not found", 404, "category not found")
	}
	return &data, nil
}

func (r *IncomeRepository) CreateIncome(income *models.Income, wallet *models.Wallet) (*models.Income, *dtos.ErrorResponse) {
	tx := r.db.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Save(wallet).Error; err != nil {
		tx.Rollback()
		return nil, dtos.NewErrorResponse("failed to update wallet", 500, err.Error())
	}
	if err := tx.Create(income).Error; err != nil {
		return nil, dtos.NewErrorResponse("failed to create income", 500, err.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return nil, dtos.NewErrorResponse("Failed to commit transaction", 500, err.Error())
	}
	return income, nil
}
func (r *IncomeRepository) UpdateIncome(income *models.Income, oldwallet, newwallet *models.Wallet) (*models.Income, *dtos.ErrorResponse) {
	tx := r.db.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Save(oldwallet).Error; err != nil {
		tx.Rollback()
		return nil, dtos.NewErrorResponse("failed to update old wallet", 500, err.Error())
	}
	if oldwallet.ID != newwallet.ID {
		if err := tx.Save(newwallet).Error; err != nil {
			tx.Rollback()
			return nil, dtos.NewErrorResponse("failed to update new wallet", 500, err.Error())
		}
	}
	if err := tx.Save(income).Error; err != nil {
		tx.Rollback()
		return nil, dtos.NewErrorResponse("failed to update income", 500, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, dtos.NewErrorResponse("Failed to commit transaction", 500, err.Error())
	}
	return income, nil
}
