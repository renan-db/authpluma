package interfaces

import (
	"context"
	"project/internal/modules/user/entity"
)

type CreateUserRepository interface {
	Execute(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
}
