package utils

import "github.com/go-playground/validator/v10"

type GlobalValidator struct {
	validator *validator.Validate
}

func NewGlobalValidator() *GlobalValidator {
	return &GlobalValidator{
		validator: validator.New(),
	}
}
func (gv *GlobalValidator) Validate(i interface{}) error {
	return gv.validator.Struct(i)
}
