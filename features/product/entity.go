package product

import "time"

type Core struct {
	ID            uint
	UserID        uint
	ProductName   string `validate:"required"`
	Description   string `validate:"required"`
	Price         int    `validate:"required"`
	ProductImages string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAllProduct() ([]Core, error)
	SelectProductById(id uint) (*Core, error)
	SelectProductByUserId(id uint) ([]Core, error)
	Update(id uint, input Core) error
	Delete(input uint) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetProductById(id uint) (*Core, error)
	GetAllProduct() ([]Core, error)
	GetProductByUserId(id uint) ([]Core, error)
	Update(id uint, input Core) error
	Delete(input uint, userID uint) error
}
