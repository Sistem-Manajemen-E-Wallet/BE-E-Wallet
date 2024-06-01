package product

import "time"

type Core struct {
	ID            int
	UserID        int
	ProductName   string
	Description   string
	Price         int
	ProductImages string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAllProduct() ([]Core, error)
	SelectProductById(id int) (*Core, error)
	SelectProductByUserId(id int) ([]Core, error)
	Update(id int, input Core) error
	Delete(input int) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetProductById(id int) (*Core, error)
	GetAllProduct() ([]Core, error)
	GetProductByUserId(id int) ([]Core, error)
	Update(id int, input Core) error
	Delete(input, userID int) error
}
