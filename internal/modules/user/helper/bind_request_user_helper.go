package helper

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Monta os dados dentro da DTO atribuída a req
// Não valida os dados somente a estrutura
func BindRequest(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}
	return nil
}

