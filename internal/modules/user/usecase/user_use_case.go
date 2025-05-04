// Package usecase contém todos os casos de uso do módulo de usuário.
package usecase

import (
	"context"
	"fmt"

	"project/internal/modules/user/entity"
	interfaces "project/internal/modules/user/interface"
)

// userUseCase implementa a interface UserUseCase.
var _ interfaces.UserUseCase = (*userUseCase)(nil)

// userUseCase representa o caso de uso com a dependência externa injetada.
type userUseCase struct {
	repo interfaces.UserRepository
}

// NewUserUseCase cria uma nova instância de userUseCase.
func NewUserUseCase(repo interfaces.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

// Create executa a lógica principal com a dependência fornecida.
func (u *userUseCase) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	userCreated, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("createUserUseCase: erro ao criar usuário no repositório: %w", err)
	}
	
	return userCreated, nil
}
