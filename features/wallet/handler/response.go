package handler

import "time"

type WalletResponse struct {
	Balance   float64
	UpdatedAt time.Time
}
