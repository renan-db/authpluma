package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleRequest é o método que será chamado pelo Echo
func HandleRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
