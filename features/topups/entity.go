package topups

import "time"

type Core struct {
	ID          int
	OrderID     string
	UserID      int
	Amount      float64 `validate:"required,gt=0"`
	Type        string
	ChannelBank string `validate:"required"`
	VaNumbers   string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DataInterface interface {
	Insert(input Core) (Core, error)
	SelectById(id int) (Core, error)
	SelectByUserID(id int) ([]Core, error)
	Update(id int, input Core) error
}

type ServiceInterface interface {
	Create(input Core) (Core, error)
	GetByID(id int) (Core, error)
	GetByUserID(id int) ([]Core, error)
	Update(id int, input Core) error
}
