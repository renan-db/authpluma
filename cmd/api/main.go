package main

import (
	"log"
	"project/internal/infrastructure/config"
	"project/internal/infrastructure/server"

	"github.com/labstack/echo/v4"
)

func main() {
	// Carrega as variáveis de ambiente diretamente
	if err := config.LoadConfig(""); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Cria uma nova instância do Echo
	e := echo.New()

	// Inicializa o servidor HTTP com Echo
	if err := server.StartHTTPServer(e); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 