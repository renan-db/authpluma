package handler

import (
	usecases "project/internal/modules/user/usecase"
)

// Agrupa todos os handlers do módulo
type UserHandlers struct{
	Create *createUserHandler
}

// Cria uma nova instância de todos os handlers do módulo
func NewUserHandlers(useCases *usecases.UserUseCases) *UserHandlers {
	return &UserHandlers{
		Create: NewCreateUserHandler(useCases.Create),
	}
}
