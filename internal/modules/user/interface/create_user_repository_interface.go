package interfaces

import (
	"context"
	"project/internal/modules/user/entity"
)

type CreateUserRepositoryInterface interface {
	Execute(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
}
