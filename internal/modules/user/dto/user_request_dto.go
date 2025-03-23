package dto

import (
	"project/internal/modules/user/entity"
)

type CreateUserRequestDTO struct {
	entity.UserEntity `json:"-"`
	Name       string `json:"user" validate:"required,min=3" example:"John Doe"`
	Email      string `json:"email" validate:"required,email" example:"john@example.com"`
}

type DeleteUserRequestDTO struct {
	// Este arquivo não contém struct, o ID é extraído da URL.
	// Exemplo: DELETE/user/{id}
}

type GetByIdUserRequestDTO struct {
	// Este arquivo não contém struct, o ID é extraído da URL.
	// Exemplo: GET /user/{id}
}

type ListUserRequestDTO struct {
	Limit  int `json:"limit" validate:"min=1"`
	Offset int `json:"offset" validate:"min:0"`
}

type UpdateUserRequestDTO struct {
	entity.UserEntity `json:"-"`
	ID         int    `json:"id" validate:"required"`
	Name       string `json:"user" validate:"required, min=3"`
	Email      string `json:"email" validate:"required, email"`
}
