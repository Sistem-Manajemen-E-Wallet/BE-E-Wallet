package data

import (
	userData "e-wallet/features/user/data"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserID  uint
	Balance int
	User    userData.User `gorm:"foreignKey:UserID"`
}
