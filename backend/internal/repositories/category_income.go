package repositories

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"

	"gorm.io/gorm"
)

type CategoryIncomeRepo struct {
	db *gorm.DB
}

func NewCategoryIncomeRepository(db *gorm.DB) *CategoryIncomeRepo {
	return &CategoryIncomeRepo{db: db}
}

func (r *CategoryIncomeRepo) FindCategoryIncomeByUserId(userID int) ([]models.IncomeCategory, *dtos.ErrorResponse) {
	var data []models.IncomeCategory
	result := r.db.Where("user_id = ?", userID).Find(&data)
	if result.Error != nil {
		return nil, dtos.NewErrorResponse("Failed to retrieve Category", 500, "database error")
	}
	return data, nil
}
func (r *CategoryIncomeRepo) FindCategoryIncomeById(categoryID int) (*models.IncomeCategory, *dtos.ErrorResponse) {
	var data models.IncomeCategory
	err := r.db.Where("id =?", categoryID).First(&data).Error
	if err != nil {
		return nil, dtos.NewErrorResponse("Category not found", 404, "category not found")
	}
	return &data, nil
}
func (r *CategoryIncomeRepo) CreateIncomeCategory(incomeCategory *models.IncomeCategory) (*models.IncomeCategory, *dtos.ErrorResponse) {
	result := r.db.Create(incomeCategory)
	if result.RowsAffected > 0 {
		return incomeCategory, nil
	}
	return nil, dtos.NewErrorResponse("Failed to create Category", 500, "database error")
}

func (r *CategoryIncomeRepo) UpdateIncomeCategory(incomeCategory *models.IncomeCategory) (*models.IncomeCategory, *dtos.ErrorResponse) {
	result := r.db.Save(incomeCategory)
	if result.RowsAffected > 0 {
		return incomeCategory, nil
	}
	return nil, dtos.NewErrorResponse("Failed to update category", 500, "database error")

}
