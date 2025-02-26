package http

import (
	"net/http"
	"project/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *usecase.ExampleService
}

func NewHandler(service *usecase.ExampleService) *Handler {
	return &Handler{service: service}
}

// HandleRequest é o método que será chamado pelo Echo
func (h *Handler) HandleRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
} 