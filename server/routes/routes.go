package routes

import (
	"app/server"
	"app/server/handlers"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(server *server.Server) {
	apiV1 := server.Echo.Group("/api/v1")

	groupHealth := apiV1.Group("/health")
	GroupHealth(server, groupHealth)
}

func GroupHealth(server *server.Server, group *echo.Group) {
	handler := handlers.NewHandlerHealth(server)

	group.GET("", handler.HealthCheck)
}
