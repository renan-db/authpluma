package http

import (
	"net/http"
	"project/internal/usecase"
)

type Handler struct {
	service *usecase.ExampleService
}

func NewHandler(service *usecase.ExampleService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Manipulador HTTP
	w.Write([]byte("Hello, World!"))
} 