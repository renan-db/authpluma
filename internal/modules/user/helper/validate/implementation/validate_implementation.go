package implementation

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	validateStructInterfaces "project/internal/modules/user/helper/validate/interfaces"
)

// Garante que a struct implemente a interface esperada
var _ validateStructInterfaces.ValidateStruct = (*RequestValidateStruct)(nil)

// Representa o validador com o validator
type RequestValidateStruct struct {
	validate *validator.Validate
}

// Cria uma nova instância com o validador
func NewRequestValidateStruct() *RequestValidateStruct {
	return &RequestValidateStruct{
		validate: validator.New(),
	}
}

// Implementa o validador com o validator
func (v *RequestValidateStruct) Validate(req interface{}) error {
	if err := v.validate.Struct(req); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
