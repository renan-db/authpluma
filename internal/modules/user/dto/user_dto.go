// Package dto define os DTOs do módulo de usuário.
package dto

import (
	"project/internal/modules/user/entity"
)

// user representa a estrutura completa do DTO de usuário.
type user struct {
	entity.User
}