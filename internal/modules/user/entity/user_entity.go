package entity

import (
	"time"
)

type UserEntity struct {
	ID         int32     `json:"id,"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	InativedAt time.Time `json:"inative_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
