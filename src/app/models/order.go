package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID        int
	SalesAmount      float64
	CommissionAmount float64
	Product          Product
}
