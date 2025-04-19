package helper

import (
	validateImplementation "project/internal/modules/user/helper/validate/implementation"
	validateInterfaces "project/internal/modules/user/helper/validate/interfaces"
)

func NewValidateStruct() validateInterfaces.ValidateStruct {
	return validateImplementation.NewRequestValidateStruct()
}
