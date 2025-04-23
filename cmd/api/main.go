package main

import (
	"log"
	_ "project/docs/swagger"
	"project/internal/infrastructure/config"
	"project/internal/infrastructure/database"

	userModule "project/internal/modules/user"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Carrega configurações
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Configura banco de dados
	dbConfig := &database.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	}

	// Conecta ao banco de dados
	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Cria uma nova instância do Echo
	e := echo.New()

	// Adiciona rota do Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Gegistra os módulos
	userModule.Register(e, db)

	// Inicializa o servidor HTTP com Echo
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
