package models

import "app/pkgs/types"

type Pricing struct {
	ID               int        `gorm:"column:id"`
	Date             types.Date `gorm:"column:date"`
	OrganizationName string     `gorm:"column:organization_name"`
	TransferAmount   int        `gorm:"column:transfer_amount"`
	FxRate           float64    `gorm:"column:fx_rate"`
}

func (Pricing) TableName() string {
	return "pricing"
}
