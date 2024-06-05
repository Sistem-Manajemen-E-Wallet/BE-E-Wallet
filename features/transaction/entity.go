package transaction

import "time"

type Core struct {
	ID             uint
	UserID         uint
	Pin            string
	OrderID        int
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
	SelectTransactionById(id uint) (*Core, error)
	SelectTransactionByMerchantId(id uint, offset int, limit int) ([]Core, error)
	UpdateStatusProgress(id uint, input Core) error
	CountByMerchantId(merchantId uint) (int, error)
	VerifyPin(pin string, idUser uint) error
	// SelectTransactionByUserId(id uint) ([]Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetTransactionById(userId uint, id uint) (*Core, error)
	GetTransactionByMerchantId(id uint, offset int, limit int) ([]Core, int, error)
	UpdateStatusProgress(idUser uint, id uint, input Core) error
	VerifyPin(pin string, idUser uint) error
}
