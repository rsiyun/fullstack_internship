package services

import (
	"dot/app/models/domain"
	"dot/app/repositories"
	"errors"
)

type CategoryServices interface {
	GetAll() ([]domain.Category, error)
	Create(name string) (*domain.Category, error)
	Show(id uint) (*domain.Category, error)
	Update(id uint, name string) (*domain.Category, error)
	Delete(id uint) (*domain.Category, error)
}

type categoryServices struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryServices {
	return &categoryServices{categoryRepo: categoryRepo}
}

func (s *categoryServices) GetAll() ([]domain.Category, error) {
	resp, err := s.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *categoryServices) Create(name string) (*domain.Category, error) {
	req := &domain.Category{
		Name: name,
	}
	err := s.categoryRepo.Create(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *categoryServices) Show(id uint) (*domain.Category, error) {
	resp, err := s.categoryRepo.Show(id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (s *categoryServices) Update(id uint, name string) (*domain.Category, error) {
	category := &domain.Category{
		ID:   id,
		Name: name,
	}
	err := s.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryServices) Delete(id uint) (*domain.Category, error) {
	category, err := s.categoryRepo.Show(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	if err := s.categoryRepo.Delete(id); err != nil {
		return nil, err
	}
	return category, nil
}
