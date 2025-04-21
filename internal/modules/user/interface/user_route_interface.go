package interfaces

import (
	"github.com/labstack/echo/v4"
)

type UserRoute interface {
	Create(group echo.Group)
}
