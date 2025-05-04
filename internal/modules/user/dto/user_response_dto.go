// Package dto define os DTOs de resposta do módulo de usuário.
package dto

import (
	"time"

	"project/internal/modules/user/entity"
)

// CreateUserResponse representa a estrutura da resposta ao criar um novo usuário.
type CreateUserResponse struct {
	entity.User
	UpdatedAt time.Time `json:"-"`
	InativedAt time.Time `json:"-"`
}

// ListUserResponse representa a estrutura da resposta ao listar usuários.
type ListUserResponse struct {
	Users []user `json:"users"`
}

// UpdateUserResponse representa a estrutura da resposta ao atualizar um usuário.
type UpdateUserResponse struct {
	entity.User
}
