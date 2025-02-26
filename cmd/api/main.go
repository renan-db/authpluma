package main

import (
	"log"
	"project/internal/infrastructure/config"
	"project/internal/infrastructure/database"
	"project/internal/infrastructure/server"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbConfig := &database.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	}

	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Cria uma nova instância do Echo
	e := echo.New()

	// Inicializa o servidor HTTP com Echo
	if err := server.StartHTTPServer(e); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 