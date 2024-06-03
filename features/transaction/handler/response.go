package handler

import (
	"e-wallet/features/transaction"
)

type GetAllTransactionResponse struct {
	ID             uint
	UserID         uint
	OrderID        int
	ProductID      uint
	Quantity       int
	TotalCost      int
	StatusProgress string
	Additional     string
	StatusPayment  string
}

func toResponse(transaction transaction.Core) GetAllTransactionResponse {
	return GetAllTransactionResponse{
		ID:             transaction.ID,
		UserID:         transaction.UserID,
		OrderID:        transaction.OrderID,
		ProductID:      transaction.ProductID,
		Quantity:       transaction.Quantity,
		TotalCost:      transaction.TotalCost,
		StatusProgress: transaction.StatusProgress,
		Additional:     transaction.Additional,
		StatusPayment:  transaction.StatusPayment,
	}
}

func toCoreList(transaction []transaction.Core) []GetAllTransactionResponse {
	result := []GetAllTransactionResponse{}
	for key := range transaction {
		result = append(result, toResponse(transaction[key]))
	}
	return result
}
