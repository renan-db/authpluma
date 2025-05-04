// Package dto define os DTOs de requisição do módulo de usuário.
package dto

// CreateUserRequest representa a estrutura da requisição para criar um novo usuário..
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3" example:"John Doe"`
	Email string `json:"email" validate:"required,email" example:"john@example.com"`
}

// ListUserRequest representa a estrutura da requisição para listar usuários.
type ListUserRequest struct {
	Limit  int `json:"limit" validate:"min=1"`
	Offset int `json:"offset" validate:"min:0"`
}

// UpdateUserRequest representa a estrutura da requisição para atualizar um usuário.
type UpdateUserRequest struct {
	ID    int `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required, min=3"`
	Email string `json:"email" validate:"required, email"`
}
