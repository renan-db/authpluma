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
