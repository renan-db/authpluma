package route // Ou o nome do seu pacote de rotas

import (
	"project/internal/modules/user/handler" // Importa o pacote handler

	"github.com/labstack/echo/v4"
)

// Agrupa todas as rotas do módulo
type UserUseRoutes struct {
	handles *handler.UserHandlers
}

// Cria uma nova instância de todas as rotas do módulo
func NewUserRoutes(handlers *handler.UserHandlers) *UserUseRoutes {
	return&UserUseRoutes{
		handles: handlers,
	}
}

// Agrupa todas as rotas do módulo
func (r *UserUseRoutes) RegisterRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("", r.handles.Create.Execute)
}