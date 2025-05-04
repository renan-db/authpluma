// Package route contém todas as rotas do módulo de usuário.
package route

import (
	"github.com/labstack/echo/v4"

	userInterfaces "project/internal/modules/user/interface"
)

// userRoute implementa a interface UserRoute.
var _ userInterfaces.UserRoute = (*userRoute)(nil)

// userRoute representa a estrutura que implementa a interface UserRoute.
type userRoute struct {
	userHandler userInterfaces.UserHandle
}

<<<<<<< HEAD
// Cria uma nova instância com a dependência necessária
func NewUserRoute(ch userInterfaces.UserHandle) userInterfaces.UserRoute {
=======
// NewUserRoute cria uma nova instância de userRoute.
func NewUserRoute(ch userInterfaces.UserHandle) userInterfaces.UserRoute{
>>>>>>> develop
	return &userRoute{
		userHandler: ch,
	}
}

// Create adiciona a rota POST para criar um novo usuário.
func (ch *userRoute) Create(group echo.Group) {
	group.POST("/users", ch.userHandler.Create)
}
