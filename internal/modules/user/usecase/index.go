package usecase

import (
	userInterfaces "project/internal/modules/user/interface"
	repository "project/internal/modules/user/repository/postgre_sql"
)

// Agrupa todos os caso de uso do módulo
type UserUseCases struct {
	Create userInterfaces.CreateUserUseCaseInterface
}

// Cria uma nova instância de todos os casos de uso do módulo
func NewUserUseCases(repos *repository.UserRepositories) *UserUseCases {
	return &UserUseCases{
		Create: NewCreateUserUseCase(repos.Create),
	}
}