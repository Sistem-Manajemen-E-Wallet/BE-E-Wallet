package data

import (
	"e-wallet/features/product/data"
	userData "e-wallet/features/user/data"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint
	CustName       string
	OrderID        int
	ProductID      uint
	ProductName    string
	MerchantID     uint
	Quantity       int
	TotalCost      int
	StatusProgress string
	Additional     string
	StatusPayment  string
	User           userData.User `gorm:"foreignKey:UserID"`
	Product        data.Product  `gorm:"foreignKey:ProductID"`
}
