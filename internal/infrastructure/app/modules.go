package app

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	userModule "project/internal/modules/user"
)

func RegisterModules(e *echo.Echo, db *sql.DB) {
	userModule.Register(e, db)
}
