<<<<<<< HEAD
=======
// Package dto define os DTOs de resposta do módulo de usuário.
>>>>>>> develop
package dto

import (
	"time"

<<<<<<< HEAD
	userEntity "project/internal/modules/user/entity"
)

type CreateUserResponseDTO struct {
	userEntity.UserEntity
	UpdatedAt  time.Time `json:"-"`
	InativedAt time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}

type DeleteUserResponseDTO struct {
	userEntity.UserEntity
	ID      int32  `json:"id"`
	Message string `json:"message"`
}

type GetByIDUserResponseDTO struct {
	userEntity.UserEntity
	DeletedAt time.Time `json:"-"`
}

type user struct {
	userEntity.UserEntity
	DeletedAt time.Time `json:"-"`
}

type ListUserResponseDTO struct {
	Users []user `json:"users"`
}

type UpdateUserResponseDTO struct {
	userEntity.UserEntity
	DeletedAt time.Time `json:"-"`
=======
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
>>>>>>> develop
}
