package service

import (
	"e-wallet/features/product"
	"e-wallet/features/transaction"
	"e-wallet/features/wallet"
	"errors"
)

type TransactionService struct {
	transactionData transaction.DataInterface
	walletData      wallet.DataInterface
	productData     product.DataInterface
}

func New(td transaction.DataInterface, wd wallet.DataInterface, pd product.DataInterface) transaction.ServiceInterface {
	return &TransactionService{
		transactionData: td,
		walletData:      wd,
		productData:     pd,
	}
}

// Create implements transaction.ServiceInterface.
func (t *TransactionService) Create(input transaction.Core) error {
	if input.UserID == 0 {
		return errors.New("[validation] you must login first")
	}
	if input.OrderID == 0 || input.ProductID == 0 || input.Quantity == 0 {
		return errors.New("[validation] nomor meja/produk/quantity tidak boleh kosong")
	}

	result, err := t.walletData.GetWalletByUserId(input.UserID)
	if err != nil {
		return err
	}
	result2, err2 := t.productData.SelectProductById(input.ProductID)
	if err2 != nil {
		return err2
	}
	if result.Balance < result2.Price {
		return errors.New("you don't have enough balance")
	}

	err3 := t.transactionData.Insert(input)
	if err3 != nil {
		return err3
	}
	return nil
}

// GetTransactionById implements transaction.ServiceInterface.
func (t *TransactionService) GetTransactionById(userId uint, id uint) (*transaction.Core, error) {
	result, err := t.transactionData.SelectTransactionById(id)
	if err != nil {
		return nil, err
	}
	if result.UserID != userId {
		return nil, errors.New("this is not your transaction")
	}

	return t.transactionData.SelectTransactionById(id)
}

// GetTransactionByMerchantId implements transaction.ServiceInterface.
func (t *TransactionService) GetTransactionByMerchantId(id uint, offset int, limit int) ([]transaction.Core, int, error) {
	result, err := t.transactionData.SelectTransactionByMerchantId(id, offset, limit)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}
	result2, err2 := t.transactionData.CountByMerchantId(id)
	if err2 != nil {
		return nil, 0, err2
	}

	return result, result2, nil
}

// UpdateStatusProgress implements transaction.ServiceInterface.
func (t *TransactionService) UpdateStatusProgress(idUser uint, id uint, input transaction.Core) error {
	result, err := t.transactionData.SelectTransactionById(id)
	if err != nil {
		return errors.New("transaction not found")
	}
	if result.MerchantID != idUser {
		return errors.New("it's not your transaction")
	}

	return t.transactionData.UpdateStatusProgress(id, input)
}

// VerifyPin implements transaction.ServiceInterface.
func (t *TransactionService) VerifyPin(pin string, idUser uint) error {
	if idUser == 0 {
		return errors.New("you must login first")
	}
	return t.transactionData.VerifyPin(pin, idUser)
}
