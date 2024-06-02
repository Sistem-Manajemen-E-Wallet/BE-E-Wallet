package transaction

import "time"

type Core struct {
	ID             uint
	UserID         uint
	ProductID      uint
	Quantity       int
	TotalCost      int
	StatusProgress string
	Additional     string
	StatusPayment  string
	MerchantID     uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAllTransaction() ([]Core, error)
	SelectTransactionById(id uint) (*Core, error)
	SelectTransactionByMerchantId(id uint) ([]Core, error)
	UpdateStatusProgress(input Core) error
	// SelectTransactionByUserId(id uint) ([]Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetAllTransaction() ([]Core, error)
	GetTransactionById(id uint) (*Core, error)
	GetTransactionByMerchantId(id uint) ([]Core, error)
	UpdateStatusProgress(id uint, input Core) error
}
