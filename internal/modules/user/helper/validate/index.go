<<<<<<< HEAD
package helper

import (
	validateImplementation "project/internal/modules/user/helper/validate/implementation"
	validateInterfaces "project/internal/modules/user/helper/validate/interfaces"
)

func NewValidateStruct() validateInterfaces.ValidateStruct {
	return validateImplementation.NewRequestValidateStruct()
}
=======
// Package helper define o submódulo de validate do módulo de usuário.
package helper

import (
	"project/internal/modules/user/helper/validate/implementation"
	"project/internal/modules/user/helper/validate/interfaces"
)

// NewValidater cria uma nova instância do validador.
func NewValidater() interfaces.Validater {
	return implementation.NewRequestValidater()
}
>>>>>>> develop
