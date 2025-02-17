package main

import (
	"log"
	"project/internal/infrastructure/server"
)

func main() {
	// Inicializa o servidor HTTP
	err := server.StartHTTPServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 