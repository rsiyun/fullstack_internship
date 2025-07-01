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

func (r *IncomeRepository) CreateIncome(income *models.Income) (*models.Income, *dtos.ErrorResponse) {
	result := r.db.Create(income)
	if result.RowsAffected > 0 {
		return income, nil
	}
	return nil, dtos.NewErrorResponse("Failed to create income", 500, "database error")
}
func (r *IncomeRepository) UpdateIncome(income *models.Income) (*models.Income, *dtos.ErrorResponse) {
	result := r.db.Save(income)
	if result.RowsAffected > 0 {
		return income, nil
	}
	return nil, dtos.NewErrorResponse("Failed to update income", 500, "database error")
}
