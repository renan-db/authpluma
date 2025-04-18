package route

import (
	"github.com/labstack/echo/v4"

	interfaces "project/internal/modules/user/interface"
)

// Garante que a struct implemente a interface esperada
var _ interfaces.CreateUserRouteInterface = (*createUserRoute)(nil)

// Representa o caso de uso com a dependência externa injetada
type createUserRoute struct {
	createUserHandler interfaces.CreateUserHandlerInterface
}

// Cria uma nova instância com a dependência necessária
func NewCreateUserRoute(ch interfaces.CreateUserHandlerInterface) interfaces.CreateUserRouteInterface{
	return &createUserRoute{
		createUserHandler: ch,
	}
}

// Executa a lógica principal com a dependência fornecida
func (ch *createUserRoute) Execute(group echo.Group) {
	group.POST("/users", ch.createUserHandler.Execute)
}
