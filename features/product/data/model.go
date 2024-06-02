package data

import (
	userData "e-wallet/features/user/data"
	"time"
)

type Product struct {
	ID            uint
	UserID        uint
	ProductName   string
	Description   string
	Price         int
	ProductImages string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          userData.User `gorm:"foreignKey:UserID"`
}
