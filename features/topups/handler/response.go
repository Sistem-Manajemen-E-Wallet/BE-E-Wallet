package handler

import (
	"e-wallet/features/topups"
	"time"
)

type TopupResponse struct {
	ID          uint      `json:"id"`
	OrderID     string    `json:"order_id"`
	UserID      uint      `json:"user_id"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	ChannelBank string    `json:"channel_bank"`
	VaNumbers   string    `json:"va_numbers"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toResponse(topup topups.Core) TopupResponse {
	return TopupResponse{
		ID:          uint(topup.ID),
		OrderID:     topup.OrderID,
		UserID:      uint(topup.UserID),
		Amount:      topup.Amount,
		Type:        topup.Type,
		ChannelBank: topup.ChannelBank,
		VaNumbers:   topup.VaNumbers,
		Status:      topup.Status,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

func toResponses(topups []topups.Core) []TopupResponse {
	var topupResponses []TopupResponse
	for _, topup := range topups {
		topupResponses = append(topupResponses, toResponse(topup))
	}
	return topupResponses
}
