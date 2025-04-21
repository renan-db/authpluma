package user

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	userHandler "project/internal/modules/user/handler"
	userRepo "project/internal/modules/user/repository"
	userRoute "project/internal/modules/user/route"
	userUsecase "project/internal/modules/user/usecase"
)

func Register(e *echo.Echo, db *sql.DB) {
	repoType := viper.GetString("USER_REPOSITORY_TYPE")

	dbConnections := userRepo.DBConnections{
		Postgres: db,
	}

	userRepositories, err:= userRepo.NewUserRepository(repoType, dbConnections)
	if err != nil {
		fmt.Printf("Failed to create user repository: %v", err)
	}

	userUseCases := userUsecase.NewUserUseCase(userRepositories)
	userHandles := userHandler.NewUserHandler(userUseCases)
	userRoutes := userRoute.NewUserRoute(userHandles)
	userRoutes.Create(*e.Group("/users"))
}
