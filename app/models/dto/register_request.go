package dto

type Register struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"object"`
}
