package data

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserID  uint
	Balance float64
}