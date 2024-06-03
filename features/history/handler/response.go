package handler

import (
	"e-wallet/features/history"
	"time"
)

type HistoryResponse struct {
	CreatedAt time.Time `json:"created_at"`
	TrxName   string    `json:"trx_name"`
	Type      string    `json:"type"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
}

func toResponse(history history.Core) HistoryResponse {
	return HistoryResponse{
		CreatedAt: history.CreatedAt,
		TrxName:   history.TrxName,
		Type:      history.Type,
		Amount:    history.Amount,
		Status:    history.Status,
	}
}

func toCoreList(history []history.Core) []HistoryResponse {
	result := []HistoryResponse{}
	for key := range history {
		result = append(result, toResponse(history[key]))
	}
	return result
}
