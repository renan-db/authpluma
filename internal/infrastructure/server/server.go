package server

import (
	"log"
	"net/http"
	deliveryHandler "project/internal/delivery/http/handler"
	"project/internal/usecase"
)

func StartHTTPServer() error {
	service := &usecase.ExampleService{}
	handler := deliveryHandler.NewHandler(service)

	http.Handle("/", handler)

	log.Println("Starting HTTP server on :8080")
	return http.ListenAndServe(":8080", nil)
} 