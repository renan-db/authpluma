package main

import (
	"log"
	"project/internal/infrastructure/config"
	"project/internal/infrastructure/server"
)

func main() {
	// Carrega as variáveis de ambiente diretamente
	if err := config.LoadConfig(""); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Inicializa o servidor HTTP
	if err := server.StartHTTPServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 