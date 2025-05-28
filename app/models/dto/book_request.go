package dto

import "dot/app/models/domain"

type BookRequest struct {
	Title       string `validate:"required,min=3,max=100"`
	Writer      string `validate:"required,min=3,max=100"`
	Description string `validate:"omitempty,max=1000"`
	CategoryID  uint   `validate:"required"`
}

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Writer      string `json:"writer"`
	Description string `json:"description"`
	CategoryId  uint   `json:"categoryid"`
	Category    CategoryResponse
}

type MultipleBookResponse struct {
	Message string `json:"message"`
	Data    []domain.Book
}

type SingleBookResponse struct {
	Message string `json:"message"`
	Data    BookResponse
}

type BookErrorResponse struct {
	Message string `json:"message"`
}
type CountBooksByCategory struct {
	Message string `json:"message"`
	Count   int64  `json:"count"`
}
