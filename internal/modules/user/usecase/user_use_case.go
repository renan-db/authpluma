package usecase

import (
	"context"
	"fmt"

	userEntity "project/internal/modules/user/entity"
	userInterfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ userInterfaces.UserUseCase = (*userUseCase)(nil)

// Representa o caso de uso com a dependência externa injetada
type userUseCase struct {
	repo userInterfaces.UserRepository
}

// Cria uma nova instância com a dependência necessária
func NewUserUseCase(repo userInterfaces.UserRepository) userInterfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

// Executa a lógica principal com a dependência fornecida
func (u *userUseCase) Create(ctx context.Context, user *userEntity.UserEntity) (*userEntity.UserEntity, error) {
	createdUser, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("createUserUseCase: erro ao criar usuário no repositório: %w", err)
	}

	return createdUser, nil
}
