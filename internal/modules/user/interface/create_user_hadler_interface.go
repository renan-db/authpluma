package interfaces

import (
	"github.com/labstack/echo/v4"
)

type CreateUserHandlerInterface interface {
	Execute(c echo.Context) error
}