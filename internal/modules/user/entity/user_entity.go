// package entity, contém a entidade de usuário.
package entity

import (
	"time"
)

// UserEntity representa a estrutura da entidade de usuário.
type User struct {
	ID         int32     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	InativedAt time.Time `json:"inatived_at"`
}
