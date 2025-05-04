// Package handler define os handlers para o módulo de usuário.
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"project/internal/modules/user/dto"
	"project/internal/modules/user/entity"
	helperBinder "project/internal/modules/user/helper/binder"
	helperValidate "project/internal/modules/user/helper/validate"
	interfaces "project/internal/modules/user/interface"
)

// Garante que userHandler implementa a interface UserHandle.
var _ interfaces.UserHandle = (*userHandler)(nil)

// userHandler implementa os handlers para o módulo de usuário.
type userHandler struct {
	useCase interfaces.UserUseCase
}

// NewUserHandler cria uma nova instância de userHandler.
func NewUserHandler(uc interfaces.UserUseCase) *userHandler {
	return &userHandler{
		useCase: uc,
	}
}

// Create processa a requisição para criar um novo usuário.
func (h *userHandler) Create(c echo.Context) error {
	var req dto.CreateUserRequest

	// Preenche o DTO de request com os dados recebidos.
	binder := helperBinder.NewBinder(c)
	if err := binder.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})	
	}

	// Valida os dados da requisição.
	validateStruct := helperValidate.NewValidater()
	if err := validateStruct.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Preenche a entidade com os dados do DTO.
	userToCreate := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// Executa o caso de uso para criação do usuário.
	userCreated, err := h.useCase.Create(c.Request().Context(), userToCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

<<<<<<< HEAD
	response := userDto.CreateUserResponseDTO{
		UserEntity: userEntity.UserEntity{
			ID:        createdUserEntity.ID,
			Name:      createdUserEntity.Name,
			Email:     createdUserEntity.Email,
			CreatedAt: createdUserEntity.CreatedAt,
=======
	// Preenche a entidade com os dados do DTO.
	response := dto.CreateUserResponse{
		User: entity.User{
			ID:        userCreated.ID,
			Name:      userCreated.Name,
			Email:     userCreated.Email,
			CreatedAt: userCreated.CreatedAt,
>>>>>>> develop
		},
	}

	return c.JSON(http.StatusCreated, response)
}
