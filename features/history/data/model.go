package data

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserID        uint
	TransactionID uint
	TopUpID       uint
	TrxName       string
	Amount        int
	Type          string
	Status        string
}
