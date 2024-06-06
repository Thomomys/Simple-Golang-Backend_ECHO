package handlers

import (
	"app/models"
	"app/responses"
	"app/server"
	"app/services/pricing"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlerPricing struct {
	Server *server.Server
}

func NewHandlerPricing(server *server.Server) *HandlerPricing {
	return &HandlerPricing{
		Server: server,
	}
}

// Refresh godoc
// @Summary Read pricing table
// @Tags Pricing
// @Accept json
// @Produce json
// @Param date path string true "Date"
// @Success 200 {object} responses.ResponsePricingTable
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/v1/pricing/table/{date} [get]
func (handler *HandlerPricing) ReadPricingTable(context echo.Context) error {
	// Get date from param
	date, err := time.Parse(time.DateOnly, context.Param("date"))
	if err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid date format. Please use this format: (yyyy-mm-dd).")
	}

	// Get pricing informations from DB
	modelPricings := []models.Pricing{}
	servicePricing := pricing.NewService(handler.Server.DB)
	err = servicePricing.ReadByDate(&modelPricings, date)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to load data from the database.")
	}

	return responses.NewResponsePricingTable(context, http.StatusOK, modelPricings)
}
