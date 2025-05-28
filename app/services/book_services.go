package services

import (
	"dot/app/models/domain"
	"dot/app/repositories"
	"errors"
)

type BookServices interface {
	Create(title, writer, description string, categoryID uint) (*domain.Book, error)
	GetAllBooks() ([]domain.Book, error)
	Show(id uint) (*domain.Book, error)
	UpdateBook(id uint, title, writer, description string, categoryID uint) (*domain.Book, error)
	DeleteBook(id uint) (*domain.Book, error)
	GetBooksBycategory(categoryID uint) ([]domain.Book, error)
	CountBooksByCategory(categoryID uint) (int64, error)
}

type bookServices struct {
	bookRepo     repositories.BookRepository
	categoryRepo repositories.CategoryRepository
}

func NewBookService(bookRepo repositories.BookRepository, categoryRepo repositories.CategoryRepository) BookServices {
	return &bookServices{bookRepo: bookRepo, categoryRepo: categoryRepo}
}

func (s *bookServices) Create(title, writer, description string, categoryID uint) (*domain.Book, error) {
	category, err := s.categoryRepo.Show(categoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	book := &domain.Book{
		Title:       title,
		Writer:      writer,
		Description: description,
		CategoryID:  categoryID,
	}

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *bookServices) GetAllBooks() ([]domain.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (s *bookServices) Show(id uint) (*domain.Book, error) {
	book, err := s.bookRepo.Show(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	return book, nil
}
func (s *bookServices) UpdateBook(id uint, title, writer, description string, categoryID uint) (*domain.Book, error) {
	category, err := s.categoryRepo.Show(categoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	book, err := s.bookRepo.Show(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}

	book.Title = title
	book.Writer = writer
	book.Description = description
	book.CategoryID = categoryID

	err = s.bookRepo.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (s *bookServices) DeleteBook(id uint) (*domain.Book, error) {
	book, err := s.bookRepo.Show(id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	if err := s.bookRepo.Delete(id); err != nil {
		return nil, err
	}
	return book, nil
}
func (s *bookServices) GetBooksBycategory(categoryID uint) ([]domain.Book, error) {
	category, err := s.categoryRepo.Show(categoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	// Ambil buku berdasarkan kategori
	books, err := s.bookRepo.GetByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (s *bookServices) CountBooksByCategory(categoryID uint) (int64, error) {
	category, err := s.categoryRepo.Show(categoryID)
	if err != nil {
		return 0, err
	}
	if category == nil {
		return 0, errors.New("category not found")
	}

	// Hitung jumlah buku
	count, err := s.bookRepo.CountByCategory(categoryID)
	if err != nil {
		return 0, err
	}

	return count, nil
}
