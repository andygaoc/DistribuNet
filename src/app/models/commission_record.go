package models

import "gorm.io/gorm"

type CommissionRecord struct {
	gorm.Model
	OrderID          int
	UserID           int
	CommissionAmount float64
	Order            Order
	User             User
}
