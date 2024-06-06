package data

import (
	userData "e-wallet/features/user/data"
	"time"

	"gorm.io/gorm"
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
	DeletedAt     gorm.DeletedAt
	User          userData.User `gorm:"foreignKey:UserID"`
}
