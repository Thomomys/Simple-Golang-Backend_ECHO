package main

import (
	"app/server"
	"app/server/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	// Create server
	server, err := server.NewServer()
	if err != nil {
		log.Println("Failed to create server:", err)
		return
	}

	// Configure routes
	routes.ConfigureRoutes(server)

	// Start server
	err = server.Start()
	if err != nil {
		log.Println("Failed to start server:", err)
	}
}
