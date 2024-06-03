package service

import (
	"e-wallet/features/transaction"
	"errors"
)

type TransactionService struct {
	transactionData transaction.DataInterface
}

func New(td transaction.DataInterface) transaction.ServiceInterface {
	return &TransactionService{
		transactionData: td,
	}
}

// Create implements transaction.ServiceInterface.
func (t *TransactionService) Create(input transaction.Core) error {
	if input.UserID == 0 || input.OrderID == 0 || input.ProductID == 0 || input.Quantity == 0 {
		return errors.New("[validation] nama/email/pin/phone tidak boleh kosong")
	}

	err := t.transactionData.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

// GetTransactionById implements transaction.ServiceInterface.
func (t *TransactionService) GetTransactionById(id uint) (*transaction.Core, error) {
	panic("unimplemented")
}

// GetTransactionByMerchantId implements transaction.ServiceInterface.
func (t *TransactionService) GetTransactionByMerchantId(id uint) ([]transaction.Core, error) {
	result, err := t.transactionData.SelectTransactionByMerchantId(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	return result, nil
}

// UpdateStatusProgress implements transaction.ServiceInterface.
func (t *TransactionService) UpdateStatusProgress(id uint, input transaction.Core) error {
	panic("unimplemented")
}
