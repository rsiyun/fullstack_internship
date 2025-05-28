package dto

type Register struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}
