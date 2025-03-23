package handle

import (
	"net/http"

	database "project/internal/infrastructure/database/sqlc"
	"project/internal/modules/user/dto"
	"project/internal/modules/user/entity"

	"github.com/labstack/echo/v4"
)

// UserHandler representa o manipulador de usuários
type UserHandler struct {
	querier database.Querier
}

// NewUserHandler cria um novo UserHandler com o querier fornecido
func NewUserHandler(querier database.Querier) *UserHandler {
	return &UserHandler{querier: querier}
}

// @Summary Criar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequestDTO true "Dados do usuário"
// @Success 201 {object} dto.CreateUserResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user dto.CreateUserRequestDTO

	// Tenta vincular os dados da requisição ao DTO
	if err := c.Bind(&user); err != nil {
		// Retorna um erro 400 se a requisição for inválida
		// Se não receber o body, com o formato dto.CreateUserRequestDTO, retorna um erro 400
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Prepara os parâmetros para a criação do usuário
	params := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	// Tenta criar o usuário no banco de dados
	createdUser, err := h.querier.CreateUser(c.Request().Context(), params)
	if err != nil {
		// Retorna um erro 500 se houver uma falha ao criar o usuário
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Prepara a resposta com os dados do usuário criado
	response := dto.CreateUserResponseDTO{
		UserEntity: entity.UserEntity{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
		},
	}

	// Retorna a resposta com status 201 (Criado)
	return c.JSON(http.StatusCreated, response)
}
