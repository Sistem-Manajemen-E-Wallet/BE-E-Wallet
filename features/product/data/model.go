package data

import "time"

type Product struct {
	ID            uint
	UserID        int
	ProductName   string
	Description   string
	Price         int
	ProductImages string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
