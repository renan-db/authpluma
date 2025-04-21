package route

import (
	"github.com/labstack/echo/v4"

	userInterfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ userInterfaces.UserRoute = (*userRoute)(nil)

// Representa o caso de uso com a dependência externa injetada
type userRoute struct {
	userHandler userInterfaces.UserHandle
}

// Cria uma nova instância com a dependência necessária
func NewUserRoute(ch userInterfaces.UserHandle) userInterfaces.UserRoute{
	return &userRoute{
		userHandler: ch,
	}
}

// Executa a lógica principal com a dependência fornecida
func (ch *userRoute) Create(group echo.Group) {
	group.POST("/users", ch.userHandler.Create)
}
