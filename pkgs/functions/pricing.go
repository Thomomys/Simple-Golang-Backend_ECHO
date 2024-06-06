package functions

import "app/models"

func TransformToPricingTable(modelPricings []models.Pricing) [][]interface{} {
	table := [][]interface{}{{"Transfer Amount"}}
	organizationFlags := map[string]bool{}
	transferFlags := map[int]bool{}
	fxRates := map[string]map[int]float64{}
	for _, modelPricing := range modelPricings {
		organizationName := modelPricing.OrganizationName
		if !organizationFlags[organizationName] {
			organizationFlags[organizationName] = true
			table[0] = append(table[0], organizationName)
		}

		transferAmount := modelPricing.TransferAmount
		if !transferFlags[transferAmount] {
			transferFlags[transferAmount] = true
			table = append(table, []interface{}{transferAmount})
		}

		_, isExited := fxRates[organizationName]
		if !isExited {
			fxRates[organizationName] = map[int]float64{}
		}
		fxRates[organizationName][transferAmount] = modelPricing.FxRate
	}

	for index := range table {
		if index != 0 {
			for subIndex, organizationName := range table[0] {
				if subIndex != 0 {
					fxRate, isExited := fxRates[organizationName.(string)][table[index][0].(int)]
					if isExited {
						table[index] = append(table[index], fxRate)
					} else {
						table[index] = append(table[index], nil)
					}
				}
			}
		}
	}
	return table
}
