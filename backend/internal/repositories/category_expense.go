package repositories

import (
	"api-money-management/internal/dtos"
	"api-money-management/internal/models"

	"gorm.io/gorm"
)

type CategoryExpenseRepo struct {
	db *gorm.DB
}

func NewCategoryExpenseRepository(db *gorm.DB) *CategoryExpenseRepo {
	return &CategoryExpenseRepo{db: db}
}

func (r *CategoryExpenseRepo) FindCategoryExpenseByUserId(userID int) ([]models.ExpenseCategory, *dtos.ErrorResponse) {
	var data []models.ExpenseCategory
	result := r.db.Where("user_id = ?", userID).Find(&data)
	if result.Error != nil {
		return nil, dtos.NewErrorResponse("Failed to retrieve Category", 500, "database error")
	}
	return data, nil
}
func (r *CategoryExpenseRepo) FindCategoryExpenseById(categoryID int) (*models.ExpenseCategory, *dtos.ErrorResponse) {
	var data models.ExpenseCategory
	err := r.db.Where("id =?", categoryID).First(&data).Error
	if err != nil {
		return nil, dtos.NewErrorResponse("Category not found", 404, "category not found")
	}
	return &data, nil
}
func (r *CategoryExpenseRepo) CreateExpenseCategory(ExpenseCategory *models.ExpenseCategory) (*models.ExpenseCategory, *dtos.ErrorResponse) {
	result := r.db.Create(ExpenseCategory)
	if result.RowsAffected > 0 {
		return ExpenseCategory, nil
	}
	return nil, dtos.NewErrorResponse("Failed to create Category", 500, "database error")
}

func (r *CategoryExpenseRepo) UpdateExpenseCategory(expenseCategory *models.ExpenseCategory) (*models.ExpenseCategory, *dtos.ErrorResponse) {
	existingCategory, err := r.FindCategoryExpenseById(int(expenseCategory.ID))
	if err != nil {
		return nil, err
	}
	existingCategory.Name = expenseCategory.Name
	existingCategory.Color = expenseCategory.Color
	existingCategory.Icon = expenseCategory.Icon
	result := r.db.Save(&existingCategory)
	if result.RowsAffected > 0 {
		return existingCategory, nil
	}
	return nil, dtos.NewErrorResponse("Failed to update category", 500, "database error")

}
