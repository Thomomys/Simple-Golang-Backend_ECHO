package responses

import (
	"app/models"
	"app/pkgs/functions"

	"github.com/labstack/echo/v4"
)

type ResponsePricingTable struct {
	Table [][]interface{} `json:"table"`
}

func NewResponsePricingTable(context echo.Context, statusCode int, modelPricings []models.Pricing) error {
	return Response(context, statusCode, ResponsePricingTable{
		Table: functions.TransformToPricingTable(modelPricings),
	})
}
