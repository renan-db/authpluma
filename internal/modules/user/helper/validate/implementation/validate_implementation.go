<<<<<<< HEAD
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
=======
// Package implementation define a implementação do submódulo de validate do módulo de usuário.
package implementation

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	interfaces "project/internal/modules/user/helper/validate/interfaces"
)

// Garante que RequestValidater implementa a interface Validater.
var _ interfaces.Validater = (*RequestValidater)(nil)	

// RequestValidater representa o validador com o validator.
type RequestValidater struct {
	validate *validator.Validate
}

// NewRequestValidater cria uma nova instância com o validador.
func NewRequestValidater() *RequestValidater {
	return &RequestValidater{
		validate: validator.New(),
	}
}

// Validate valida a estrutura da requisição.
func (v *RequestValidater) Validate(data interface{}) error {	
	if err := v.validate.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)	
	}

	return nil
}
>>>>>>> develop
