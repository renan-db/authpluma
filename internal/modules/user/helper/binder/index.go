// Package helper define o submódulo de bind do módulo de usuário.
package helper

import (
	binderImplementation "project/internal/modules/user/helper/binder/implementation"
	binderInterfaces "project/internal/modules/user/helper/binder/interface"

	"github.com/labstack/echo/v4"
)

// NewBinder cria uma nova instância do binder com o contexto do echo.
func NewBinder(ctx echo.Context) binderInterfaces.Binder {
	return binderImplementation.NewEchoBinder(ctx)
}
