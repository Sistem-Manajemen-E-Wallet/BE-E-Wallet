package wallet

import "time"

type Core struct {
	ID        uint
	UserID    uint
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	CreateWallet(id uint) error
	GetWalletById(id uint) (Core, error)
	UpdateBalanceMinus(id uint, amount float64) error
}

type ServiceInterface interface {
	GetWalletById(id uint) (Core, error)
}
