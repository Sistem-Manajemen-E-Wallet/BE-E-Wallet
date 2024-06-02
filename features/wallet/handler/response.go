package handler

import "time"

type WalletResponse struct {
	Balance   int
	UpdatedAt time.Time
}
