package handler

import (
	"net/http"

	"project/internal/modules/user/dto"
	"project/internal/modules/user/entity"
	"project/internal/modules/user/helper"
	interfaces "project/internal/modules/user/interface"

	"github.com/labstack/echo/v4"
)

// Garante que a struct implemente a interface esperada
var _ interfaces.CreateUserHandlerInterface = (*createUserHandler)(nil)

// Representa o caso de uso com a dependência externa injetada
type createUserHandler struct {
	useCase interfaces.CreateUserUseCaseInterface
}

// Cria uma nova instância com a dependência necessária
func NewCreateUserHandler(uc interfaces.CreateUserUseCaseInterface) *createUserHandler {
	return &createUserHandler{
		useCase: uc,
	}
}

// Executa a lógica principal com a dependência fornecida
func (h *createUserHandler) Execute(c echo.Context) error {
	var req dto.CreateUserRequestDTO
	
	if err := helper.BindRequest(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := helper.ValidateStruct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()}) 
	}

	// Mapear DTO para Entidade
	userEntity := &entity.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}

	// Chamar a useCase para criar o usuário
	createdUserEntity, err := h.useCase.Execute(c.Request().Context(), userEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := dto.CreateUserResponseDTO{
		UserEntity: entity.UserEntity{
			ID:        createdUserEntity.ID,
			Name:      createdUserEntity.Name,
			Email:     createdUserEntity.Email,
			CreatedAt: createdUserEntity.CreatedAt,
		},
	}

	return c.JSON(http.StatusCreated, response)
}
