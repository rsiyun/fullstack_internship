package dto

type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
