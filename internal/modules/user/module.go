// Package user contém todos os módulos relacionados ao usuário.
package user

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"project/internal/modules/user/handler"
	"project/internal/modules/user/repository"
	"project/internal/modules/user/route"
	"project/internal/modules/user/usecase"
)

// Register registra os módulos relacionados ao usuário no aplicativo.
func Register(e *echo.Echo, db *sql.DB) {
	repoType := viper.GetString("USER_REPOSITORY_TYPE")

	// Cria uma conexão com o banco de dados
	dbConnections := repository.DBConnections{
		Postgres: db,
	}

	// Cria um repositório de usuário
	userRepositories, err:= repository.NewUserRepository(repoType, dbConnections)
	if err != nil {
		fmt.Printf("Failed to create user repository: %v", err)
	}

	// Cria um caso de uso de usuário
	userUseCases := usecase.NewUserUseCase(userRepositories)

	// Cria um manipulador de usuário
	userHandles := handler.NewUserHandler(userUseCases)

	// Cria uma rota de usuário
	userRoutes := route.NewUserRoute(userHandles)
	userRoutes.Create(*e.Group("/users"))
}
