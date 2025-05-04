<<<<<<< HEAD
package dto

import (
	userEntity "project/internal/modules/user/entity"
)

type CreateUserRequestDTO struct {
	userEntity.UserEntity `json:"-"`
	Name                  string `json:"name" validate:"required,min=3" example:"John Doe"`
	Email                 string `json:"email" validate:"required,email" example:"john@example.com"`
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
=======
// Package dto define os DTOs de requisição do módulo de usuário.
package dto

// CreateUserRequest representa a estrutura da requisição para criar um novo usuário..
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3" example:"John Doe"`
	Email string `json:"email" validate:"required,email" example:"john@example.com"`
}

// ListUserRequest representa a estrutura da requisição para listar usuários.
type ListUserRequest struct {
>>>>>>> develop
	Limit  int `json:"limit" validate:"min=1"`
	Offset int `json:"offset" validate:"min:0"`
}

<<<<<<< HEAD
type UpdateUserRequestDTO struct {
	userEntity.UserEntity `json:"-"`
	ID                    int    `json:"id" validate:"required"`
	Name                  string `json:"name" validate:"required, min=3"`
	Email                 string `json:"email" validate:"required, email"`
=======
// UpdateUserRequest representa a estrutura da requisição para atualizar um usuário.
type UpdateUserRequest struct {
	ID    int `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required, min=3"`
	Email string `json:"email" validate:"required, email"`
>>>>>>> develop
}
