package interfaces

import (
	"context"
	"project/internal/modules/user/entity"
	"time"
)

type UserCRUD interface {
	CreateUser(ctx context.Context, user *entity.UserEntity) error
	DeleteUser(ctx context.Context, id int32) error
	GetUserById(ctx context.Context, id int32) (*entity.UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, user *entity.UserEntity) error
	ListUsers(ctx context.Context) ([]entity.UserEntity, error)
}

type UserValidationError interface {
	ValidateRequiredFields(user *entity.UserEntity) error
	ValidateNameLength(name string) error
	ValidateNameFormat(name string) error
	ValidateEmailLength(email string) error
	ValidateEmailFormat(email string) error
	ValidateDatabaseConnection() error
	ValidateEmailExistsInDB(email string) error
}

type UserSuccess interface {
	ValidateUserCreation(user *entity.UserEntity) (*entity.UserEntity, error)
}

type UserHooks interface {
	BeforeCreate() time.Time
	BeforeUpdate() time.Time
}

type UserInterface interface {
	UserCRUD
	UserHooks
	UserValidationError
	UserSuccess
	UserValidateAll(user *entity.UserEntity) error
}
