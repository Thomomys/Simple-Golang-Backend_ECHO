package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() (*Server, error) {
	// Get the values of the environment variables
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	// Open database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Create server
	return &Server{
		Echo: echo.New(),
		DB:   db,
	}, nil
}

func (server *Server) Start() error {
	// Get the values of the environment variables
	port := os.Getenv("SERVER_PORT")

	// Start searver
	return server.Echo.Start(":" + port)
}
