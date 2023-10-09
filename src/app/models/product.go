package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name             string
	SupplierID       int
	CommissionRuleID int
	Supplier         User `gorm:"foreignKey:SupplierID"`
	CommissionRule   CommissionRule
}
