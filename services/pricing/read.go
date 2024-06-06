package pricing

import (
	"app/models"
	"time"
)

func (service *Service) ReadByDate(modelPricings *[]models.Pricing, date time.Time) error {
	return service.DB.
		Where("date = ?", date).
		Order("transfer_amount").
		Find(modelPricings).
		Error
}
