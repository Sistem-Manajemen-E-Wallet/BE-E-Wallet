package history

import "time"

type Core struct {
	ID            uint `gorm:"primarykey"`
	UserID        uint
	TransactionID uint
	TopUpID       uint
	TrxName       string
	Amount        int
	Type          string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DataInterface interface {
	InsertHistory(input Core) error
	SelectAllHistory(idUser uint, offset int, limit int) ([]Core, error)
	UpdateHistoryTopUp(input Core) error
	CountHistory(idUser uint) (int, error)
}

type ServiceInterface interface {
	GetAllHistory(idUser uint, offset int, limit int) ([]Core, int, error)
}
