package user

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	userHandler "project/internal/modules/user/handler"
	userRepo "project/internal/modules/user/repository/postgre_sql"
	userRoute "project/internal/modules/user/route"
	userUsecase "project/internal/modules/user/usecase"
)

func Register(e *echo.Echo, db *sql.DB) {
	userRepositories := userRepo.NewUserRepositories(db)
	userUseCases := userUsecase.NewUserUseCases(userRepositories)
	userHandles := userHandler.NewUserHandlers(userUseCases)
	userRoutes := userRoute.NewUserRoutes(userHandles)
	userRoutes.RegisterRoutes(e)
}
