package repositories

import (
	"dot/app/models/domain"
	"errors"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *domain.Book) error
	GetAll() ([]domain.Book, error)
	Show(id uint) (*domain.Book, error)
	Update(book *domain.Book) error
	Delete(id uint) error
	GetByCategory(categoryID uint) ([]domain.Book, error)
	CountByCategory(categoryID uint) (int64, error)
}
type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}
func (r *bookRepository) CountByCategory(categoryID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Book{}).Where("category_id").Count(&count).Error
	return count, err
}
func (r *bookRepository) GetByCategory(categoryID uint) ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Where("category_id = ?", categoryID).Preload("Category").Find(&books).Error
	return books, err
}
func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Book{}, id).Error
}
func (r *bookRepository) Update(book *domain.Book) error {
	return r.db.Save(book).Error
}
func (r *bookRepository) Show(id uint) (*domain.Book, error) {
	var book domain.Book
	err := r.db.Preload("Category").First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &book, err
}
func (r *bookRepository) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Preload("Category").Find(&books).Error
	return books, err
}
