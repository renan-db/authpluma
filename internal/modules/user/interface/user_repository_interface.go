package interfaces

import (
	"context"

	"project/internal/modules/user/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
}
