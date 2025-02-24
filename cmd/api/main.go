package main

import (
	"log"
	"os"
	"project/internal/infrastructure/config"
	"project/internal/infrastructure/server"
)

func main() {
	// Pega o ambiente dos argumentos da linha de comando
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev" // Default para desenvolvimento
	}

	// Carrega o arquivo .env específico do ambiente
	if err := config.LoadConfig(env); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Inicializa o servidor HTTP
	if err := server.StartHTTPServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 