package dto

type BookRequest struct {
	Title       string `validate:"required,min=3,max=100"`
	Writer      string `validate:"required,min=3,max=100"`
	Description string `validate:"omitempty,max=1000"`
	CategoryID  uint   `validate:"required"`
}
