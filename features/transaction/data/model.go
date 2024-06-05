package data

import (
	userData "e-wallet/features/user/data"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint
	OrderID        int
	ProductID      uint
	MerchantID     uint
	Quantity       int
	TotalCost      int
	StatusProgress string
	Additional     string
	StatusPayment  string
	User           userData.User `gorm:"foreignKey:UserID"`
}
