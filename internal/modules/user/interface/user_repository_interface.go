// Package interfaces define as interfaces do módulo de usuário.
package interfaces

import (
	"context"

	"project/internal/modules/user/entity"
)

// UserRepository define a interface para o repositório de usuário.
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
