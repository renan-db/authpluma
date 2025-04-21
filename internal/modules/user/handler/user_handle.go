package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	userDto "project/internal/modules/user/dto"
	userEntity "project/internal/modules/user/entity"
	bindHelper "project/internal/modules/user/helper/bind"
	validateHelper "project/internal/modules/user/helper/validate"
	userInterfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ userInterfaces.UserHandle = (*userHandler)(nil)

// Representa o caso de uso com a dependência externa injetada
type userHandler struct {
	useCase userInterfaces.UserUseCase
}

// Cria uma nova instância com a dependência necessária
func NewUserHandler(uc userInterfaces.UserUseCase) *userHandler {
	return &userHandler{
		useCase: uc,
	}
}

// Implementa o caso de uso para criar um usuário
func (h *userHandler) Create(c echo.Context) error {
	var req userDto.CreateUserRequestDTO

	// Bind request
	binder := bindHelper.NewBinder(c)
	if err := binder.Binder(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Validate Struct request
	validateStruct := validateHelper.NewValidateStruct()
	if err := validateStruct.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Mapear DTO para Entidade
	userEntityDto := &userEntity.UserEntity{
		Name:  req.Name,
		Email: req.Email,
	}

	// Chamar a useCase para criar o usuário
	createdUserEntity, err := h.useCase.Create(c.Request().Context(), userEntityDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := userDto.CreateUserResponseDTO{
		UserEntity: userEntity.UserEntity{
			ID:        createdUserEntity.ID,	
			Name:      createdUserEntity.Name,
			Email:     createdUserEntity.Email,
			CreatedAt: createdUserEntity.CreatedAt,
		},
	}

	return c.JSON(http.StatusCreated, response)
}
