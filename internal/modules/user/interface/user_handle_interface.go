package interfaces

import (
	"github.com/labstack/echo/v4"
)

type UserHandle interface {
	Create(c echo.Context) error
}
