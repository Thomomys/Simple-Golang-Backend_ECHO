package server

import (
	"os"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() (*Server, error) {

	// Create server
	return &Server{
		Echo: echo.New(),
	}, nil
}

func (server *Server) Start() error {
	// Get the values of the environment variables
	port := os.Getenv("PORT")

	// Start searver
	return server.Echo.Start(":" + port)
}
