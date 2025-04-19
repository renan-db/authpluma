package repository

import (
	"database/sql"

	userinterfaces "project/internal/modules/user/interface"
)

// Agrupa todos os repositórios do usuário
type UserRepositories struct{
	Create userinterfaces.CreateUserRepository
}

// Cria uma nova instância de todos os repositórios do usuário
func NewUserRepositories(db *sql.DB) *UserRepositories {
	return &UserRepositories{
		Create: NewUserRepository(db),
	}
}
