package models

import "gorm.io/gorm"

type Settlement struct {
	gorm.Model
	SettlementPeriod string
	SettlementAmount float64
	UserID           int
	User             User
}
