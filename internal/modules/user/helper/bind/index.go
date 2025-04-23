package helper

import (
	bindImplementation "project/internal/modules/user/helper/bind/implementation"
	bindInterfaces "project/internal/modules/user/helper/bind/interface"

	"github.com/labstack/echo/v4"
)

// Cria uma nova instância do binder com o contexto do echo
func NewBinder(ctx echo.Context) bindInterfaces.Bind {
	return bindImplementation.NewEchoBinder(ctx)
}
