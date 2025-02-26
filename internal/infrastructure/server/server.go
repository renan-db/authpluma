package server

import (
	"log"
	deliveryHandler "project/internal/delivery/http/handler"
	"project/internal/usecase"

	"github.com/labstack/echo/v4"
)

func StartHTTPServer(e *echo.Echo) error {
	service := &usecase.ExampleService{}
	handler := deliveryHandler.NewHandler(service)

	// Configurar rota com Echo
	e.GET("/", handler.HandleRequest) // Ajuste o nome do método conforme seu handler

	log.Println("Starting HTTP server on :8080")
	return e.Start(":8080")
} 