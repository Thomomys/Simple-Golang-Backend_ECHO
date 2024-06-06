package handlers

import (
	"app/responses"
	"app/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerHealth struct {
	server *server.Server
}

func NewHandlerHealth(server *server.Server) *HandlerHealth {
	return &HandlerHealth{server: server}
}

// Refresh godoc
// @Summary Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Router /api/v1/health [get]
func (handler *HandlerHealth) HealthCheck(c echo.Context) error {
	return responses.MessageResponse(c, http.StatusOK, "Server is running!")
}
