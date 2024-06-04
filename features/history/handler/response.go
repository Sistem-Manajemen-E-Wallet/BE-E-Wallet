package handler

import (
	"e-wallet/features/history"
	"time"
)

type HistoryResponse struct {
	ID            uint      `json:"id"`
	TransactionID uint      `json:"transaction_id,omitempty"`
	TopUpID       uint      `json:"topup_id,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	TrxName       string    `json:"trx_name"`
	Type          string    `json:"type"`
	Amount        int       `json:"amount"`
	Status        string    `json:"status"`
}

func toResponse(history history.Core) HistoryResponse {
	return HistoryResponse{
		ID:            history.ID,
		TransactionID: history.TransactionID,
		TopUpID:       history.TopUpID,
		CreatedAt:     history.CreatedAt,
		TrxName:       history.TrxName,
		Type:          history.Type,
		Amount:        history.Amount,
		Status:        history.Status,
	}
}

func toCoreList(history []history.Core) []HistoryResponse {
	result := []HistoryResponse{}
	for key := range history {
		result = append(result, toResponse(history[key]))
	}
	return result
}
