// Package postgre_sql contém o repositório de usuário para PostgreSQL.
package postgre_sql

import (
	"context"
	"database/sql"
	"project/internal/modules/user/entity"

	database "project/internal/infrastructure/database/sqlc"
	userinterfaces "project/internal/modules/user/interface"
)

// Garante que postgresUserRepository implementa UserRepository
var _ userinterfaces.UserRepository = (*userRepositoryPostgre)(nil)

// Representa a estrutura do repositório de usuário para PostgreSQL
type userRepositoryPostgre struct {
	userQueries *database.Queries
	sqlDB       *sql.DB
}

// Crie e retorne uma nova instância do repositório de usuário para PostgreSQL
func NewUserRepositoryPostgre(db *sql.DB) userinterfaces.UserRepository {
	return &userRepositoryPostgre{
		userQueries: database.New(db),
		sqlDB:       db,
	}
}

// Create implementa o método Create da interface UserRepository.
<<<<<<< HEAD
func (r *userRepositoryPostgre) Create(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {

	params := database.CreateUserParams{
=======
func (r *userRepositoryPostgre) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	
	paramsUserToCreate := database.CreateUserParams{
>>>>>>> develop
		Name:  user.Name,
		Email: user.Email,
	}

	// Executa a querie gerada pelo sqlc
	userCreatedDB, err := r.userQueries.CreateUser(ctx, paramsUserToCreate)
	if err != nil {
		return nil, err
	}

	// Cria uma nova entidade de usuário com os dados do banco de dados
	userCreated := &entity.User{
		ID:        userCreatedDB.ID,
		Name:      userCreatedDB.Name,
		Email:     userCreatedDB.Email,
		CreatedAt: userCreatedDB.CreatedAt,
		UpdatedAt: userCreatedDB.UpdatedAt,
	}

	return userCreated, nil
}
