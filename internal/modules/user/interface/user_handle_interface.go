// Package interfaces define as interfaces do módulo de usuário.
package interfaces

import (
	"github.com/labstack/echo/v4"
)

// UserHandle define a interface para os handlers do módulo de usuário.
type UserHandle interface {
	Create(c echo.Context) error
}
