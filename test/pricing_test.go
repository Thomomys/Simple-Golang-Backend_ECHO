package test

import (
	"app/models"
	"app/pkgs/functions"
	"app/pkgs/types"
	"app/server"
	"app/services/pricing"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	samplePricings = []models.Pricing{
		{ID: 26, Date: types.NewDate(2024, time.January, 10), OrganizationName: "RapidRemit", TransferAmount: 100, FxRate: 4.94},
		{ID: 311, Date: types.NewDate(2024, time.January, 10), OrganizationName: "GlobalSettle", TransferAmount: 100, FxRate: 4.21},
		{ID: 312, Date: types.NewDate(2024, time.January, 10), OrganizationName: "SecureShift", TransferAmount: 300, FxRate: 0.65},
		{ID: 231, Date: types.NewDate(2024, time.January, 10), OrganizationName: "TransactFX", TransferAmount: 300, FxRate: 4.47},
		{ID: 757, Date: types.NewDate(2024, time.January, 10), OrganizationName: "CoinGate", TransferAmount: 300, FxRate: 3.48},
		{ID: 286, Date: types.NewDate(2024, time.January, 10), OrganizationName: "GlobalSettle", TransferAmount: 400, FxRate: 4.2},
		{ID: 567, Date: types.NewDate(2024, time.January, 10), OrganizationName: "SecureShift", TransferAmount: 400, FxRate: 0.71},
		{ID: 869, Date: types.NewDate(2024, time.January, 10), OrganizationName: "SecureShift", TransferAmount: 400, FxRate: 1.25},
		{ID: 970, Date: types.NewDate(2024, time.January, 10), OrganizationName: "GlobalSettle", TransferAmount: 500, FxRate: 1.3},
	}
	samplePricingTable = [][]interface{}{
		{
			"Transfer Amount",
			"RapidRemit",
			"GlobalSettle",
			"SecureShift",
			"TransactFX",
			"CoinGate",
		},
		{
			int(100),
			4.94,
			4.21,
			nil,
			nil,
			nil,
		},
		{
			int(300),
			nil,
			nil,
			0.65,
			4.47,
			3.48,
		},
		{
			int(400),
			nil,
			4.2,
			1.25,
			nil,
			nil,
		},
		{
			int(500),
			nil,
			1.3,
			nil,
			nil,
			nil,
		},
	}
)

func TestReadPricingTable(t *testing.T) {
	// Load environment
	err := godotenv.Load("../.env")
	if assert.Equal(t, nil, err) {

		// Create server
		server, err := server.NewServer()
		if assert.Equal(t, nil, err) {

			// Load data from database
			modelPricings := []models.Pricing{}
			repositoryPricing := pricing.NewService(server.DB)
			err = repositoryPricing.ReadByDate(&modelPricings, time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC))
			if assert.Equal(t, nil, err) {
				assert.Equal(t, samplePricings, modelPricings)

				// Transform data to table structure
				pricingTable := functions.TransformToPricingTable(modelPricings)
				assert.Equal(t, samplePricingTable, pricingTable)
			}
		}
	}
}
