// package server é responsável por iniciar o servidor HTTP.
package server

import (
	"log"

	"github.com/labstack/echo/v4"
)

// StartHTTPServer inicia o servidor HTTP.
func StartHTTPServer(e *echo.Echo) error {
	// Inicia o servidor HTTP
	log.Println("Starting HTTP server on :8080")
	return e.Start(":8080")
}
