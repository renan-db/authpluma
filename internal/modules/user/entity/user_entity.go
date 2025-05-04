<<<<<<< HEAD
=======
// package entity, contém a entidade de usuário.
>>>>>>> develop
package entity

import (
	"time"
)

<<<<<<< HEAD
type UserEntity struct {
	ID         int32     `json:"id,"`
=======
// UserEntity representa a estrutura da entidade de usuário.
type User struct {
	ID         int32     `json:"id"`
>>>>>>> develop
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
<<<<<<< HEAD
	InativedAt time.Time `json:"inative_at"`
	DeletedAt  time.Time `json:"deleted_at"`
=======
	InativedAt time.Time `json:"inatived_at"`
>>>>>>> develop
}
