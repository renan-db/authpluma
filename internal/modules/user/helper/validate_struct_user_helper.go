package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Valida os dados da DTO marcados com 'validate'
// Validação acontece de acordo com as regras declaradas no DTO
func ValidateStruct(req interface{}) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return fmt.Errorf("validation failed: %w", err)
		
	}

	return nil
}
