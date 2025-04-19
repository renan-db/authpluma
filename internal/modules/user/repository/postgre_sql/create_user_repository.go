package repository

import (
	"context"
	"database/sql"
	"fmt"

	database "project/internal/infrastructure/database/sqlc"
	userEntity "project/internal/modules/user/entity"
	userInterfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ userInterfaces.CreateUserRepository = (*userRepository)(nil)

// Representa o caso de uso com a dependência externa injetada
type userRepository struct {
	querier database.Querier
	db      *sql.DB
}

// Cria uma nova instância com a dependência necessária
func NewUserRepository(db *sql.DB) userInterfaces.CreateUserRepository {
	return &userRepository{
		querier: database.New(db),
		db:      db,
	}
}

// Executa a lógica principal com a dependência fornecida
func (r *userRepository) Execute(ctx context.Context, user *userEntity.UserEntity) (*userEntity.UserEntity, error) {
	params := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	// Usa a querie do SQLC para criar o usuário no banco de dados
	createdDbUser, err := r.querier.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("userRepository: erro ao executar query para criar usuário: %w", err)
	}

	return &userEntity.UserEntity{
		ID:        createdDbUser.ID,
		Name:      createdDbUser.Name,
		Email:     createdDbUser.Email,
		CreatedAt: createdDbUser.CreatedAt,
		UpdatedAt: createdDbUser.UpdatedAt,
	}, nil
}
