package data

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserID        uint
	TransactionID uint `gorm:"default:null"`
	TopUpID       uint `gorm:"default:null"`
	TrxName       string
	Amount        int
	Type          string
	Status        string
}
