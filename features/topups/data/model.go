package data

import (
	userData "e-wallet/features/user/data"
	"time"
)

type TopUp struct {
	ID          int
	OrderID     string
	UserID      int
	Amount      float64
	Type        string
	ChannelBank string
	VaNumbers   string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        userData.User `gorm:"foreignKey:UserID"`
}
