package dto

import (
	"time"

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
}
