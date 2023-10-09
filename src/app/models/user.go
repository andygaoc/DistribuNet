package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string
	Role             string
	ParentUserID     int
	CommissionRuleID int
	ParentUser       *User `gorm:"foreignKey:ParentUserID"`
	CommissionRule   CommissionRule
}

type CommissionRule struct {
	gorm.Model
	Name string
	Type string
}
