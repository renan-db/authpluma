package route

import (
	"project/internal/modules/user/handle"

	"github.com/labstack/echo/v4"
)

// RegisterUserRoutes registra as rotas do usuário
func RegisterUserRoutes(e *echo.Echo, userHandler *handle.UserHandler) {
	// Grupo de rotas de usuário
	userGroup := e.Group("/users")
	
	// Registra rotas
	userGroup.POST("", userHandler.CreateUser)
	}