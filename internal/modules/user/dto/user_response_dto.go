package dto

import (
	"project/internal/modules/user/entity"
	"time"
)

type CreateUserResponseDTO struct {
	entity.UserEntity
	UpdatedAt time.Time `json:"-"`
	InativedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type DeleteUserResponseDTO struct {
	entity.UserEntity
	ID         int32  `json:"id"`
	Message    string `json:"message"`
}

type GetByIDUserResponseDTO struct {
	entity.UserEntity
	DeletedAt time.Time `json:"-"`
}

type user struct {
	entity.UserEntity
	DeletedAt time.Time `json:"-"`
}

type ListUserResponseDTO struct {
	Users []user `json:"users"`
}

type UpdateUserResponseDTO struct {
	entity.UserEntity
	DeletedAt time.Time `json:"-"`
}
