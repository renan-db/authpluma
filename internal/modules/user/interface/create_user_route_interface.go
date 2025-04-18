package interfaces

import "github.com/labstack/echo/v4"

type CreateUserRouteInterface interface{
	Execute(group echo.Group)
}