package interfaces

import (
	"context"

	"project/internal/modules/user/entity"
)

type UserUseCase interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
