package validator

import (
	"github.com/go-playground/validator"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}
