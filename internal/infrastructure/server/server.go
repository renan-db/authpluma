package server

import (
	"log"

	"github.com/labstack/echo/v4"
)

func StartHTTPServer(e *echo.Echo) error {

	log.Println("Starting HTTP server on :8080")
	return e.Start(":8080")
}
