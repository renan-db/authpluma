package route // Ou o nome do seu pacote de rotas

import (
	"github.com/labstack/echo/v4"

	userHandler "project/internal/modules/user/handler"
)

// Agrupa todas as rotas do módulo
type UserUseRoutes struct {
	handles *userHandler.UserHandlers
}

// Cria uma nova instância de todas as rotas do módulo
func NewUserRoutes(handlers *userHandler.UserHandlers) *UserUseRoutes {
	return&UserUseRoutes{
		handles: handlers,
	}
}

// Agrupa todas as rotas do módulo
func (r *UserUseRoutes) RegisterRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("", r.handles.Create.Execute)
}