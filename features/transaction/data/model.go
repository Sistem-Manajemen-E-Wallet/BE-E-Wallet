package data

import (
	productData "e-wallet/features/product/data"
	userData "e-wallet/features/user/data"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint
	OrderID        int
	ProductID      uint
	Quantity       int
	TotalCost      int
	StatusProgress string
	Additional     string
	StatusPayment  string
	MerchantID     uint
	User           userData.User       `gorm:"foreignKey:UserID"`
	Product        productData.Product `gorm:"foreignKey:ProductID"`
}
