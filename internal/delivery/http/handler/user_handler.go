package http

import (
	"net/http"
	database "project/internal/infrastructure/database/sqlc"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserRequest representa o payload para criar/atualizar usuário
// @Description Estrutura de dados do usuário para requisições
type UserRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
}

// UserResponse representa a resposta com dados do usuário
// @Description Estrutura de dados do usuário para respostas
type UserResponse struct {
	ID        int32  `json:"id" example:"1"`
	Name      string `json:"name" example:"John Doe"`
	Email     string `json:"email" example:"john@example.com"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserHandler struct {
	querier database.Querier
}

func NewUserHandler(querier database.Querier) *UserHandler {
	return &UserHandler{querier: querier}
}

// @Summary Criar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequest true "Dados do usuário"
// @Success 201 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	params := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	createdUser, err := h.querier.CreateUser(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// @Summary Buscar usuário por ID
// @Description Retorna um usuário específico pelo ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} database.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user, err := h.querier.GetUser(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// @Summary Listar todos os usuários
// @Description Retorna uma lista de todos os usuários
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} database.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) ListUsers(c echo.Context) error {
	users, err := h.querier.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// @Summary Atualizar usuário
// @Description Atualiza os dados de um usuário existente
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param user body UserRequest true "Dados do usuário"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var user UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	params := database.UpdateUserParams{
		ID:    int32(id),
		Name:  user.Name,
		Email: user.Email,
	}

	updatedUser, err := h.querier.UpdateUser(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// @Summary Deletar usuário
// @Description Remove um usuário do sistema
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	err = h.querier.DeleteUser(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
} 