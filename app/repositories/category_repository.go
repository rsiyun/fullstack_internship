package repositories

import (
	"dot/app/models/domain"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetAll() ([]domain.Category, error)
	Show(id uint) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetAll() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Find(&categories).Error
	return categories, err
}
func (r *categoryRepository) Show(id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &category, err
}
func (r *categoryRepository) Update(category *domain.Category) error {
	var existingCategory domain.Category
	err := r.db.First(&existingCategory, category.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("category not found")
		}
		return err
	}
	return r.db.Model(&existingCategory).Updates(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Category{}, id).Error
}
