package utils

import (
	"github.com/go-playground/validator/v10"
)

// NewValidator creates a new validator instance
func NewValidator() *validator.Validate {
	return validator.New()
}

// ValidateStruct validates a struct using the provided validator instance
func ValidateStruct(v *validator.Validate, s interface{}) error {
	return v.Struct(s)
}
