package usecase

import (
	"context"
	"fmt"

	userEntity "project/internal/modules/user/entity"
	userInterfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ userInterfaces.CreateUserUseCaseInterface = (*createUserUseCase)(nil)

// Representa o caso de uso com a dependência externa injetada
type createUserUseCase struct {
	repo userInterfaces.CreateUserRepository
}

// Cria uma nova instância com a dependência necessária
func NewCreateUserUseCase(repo userInterfaces.CreateUserRepository) userInterfaces.CreateUserUseCaseInterface {
	return &createUserUseCase{
		repo: repo,
	}
}

// Executa a lógica principal com a dependência fornecida
func (u *createUserUseCase) Execute(ctx context.Context, user *userEntity.UserEntity) (*userEntity.UserEntity, error) {
	createdUser, err := u.repo.Execute(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("createUserUseCase: erro ao criar usuário no repositório: %w", err)
	}
	
	return createdUser, nil
}
