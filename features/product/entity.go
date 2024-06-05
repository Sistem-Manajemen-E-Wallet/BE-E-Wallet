package product

import "time"

type Core struct {
	ID            uint
	UserID        uint
	MerchantName  string
	ProductName   string `validate:"required"`
	Description   string `validate:"required"`
	Price         int    `validate:"required"`
	ProductImages string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAllProduct(offset, limit int) ([]Core, error)
	SelectProductById(id uint) (*Core, error)
	SelectProductByUserId(id uint, offset, limit int) ([]Core, error)
	Update(id uint, input Core) error
	Delete(input uint) error
	CountProductByUserId(id uint) (int, error)
	CountProduct() (int, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetProductById(id uint) (*Core, error)
	GetAllProduct(offset, limit int) ([]Core, int, error)
	GetProductByUserId(id uint, offset, limit int) ([]Core, int, error)
	Update(id uint, input Core) error
	Delete(input uint, userID uint) error
}
