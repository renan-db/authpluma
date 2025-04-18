package usecase

import (
	"context"
	"fmt"

	"project/internal/modules/user/entity"
	interfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ interfaces.CreateUserUseCaseInterface = (*createUserUseCase)(nil)

// Representa o caso de uso com a dependência externa injetada
type createUserUseCase struct {
	repo interfaces.CreateUserRepositoryInterface
}

// Cria uma nova instância com a dependência necessária
func NewCreateUserUseCase(repo interfaces.CreateUserRepositoryInterface) interfaces.CreateUserUseCaseInterface {
	return &createUserUseCase{
		repo: repo,
	}
}

// Executa a lógica principal com a dependência fornecida
func (u *createUserUseCase) Execute(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	createdUser, err := u.repo.Execute(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("createUserUseCase: erro ao criar usuário no repositório: %w", err)
	}
	
	return createdUser, nil
}
