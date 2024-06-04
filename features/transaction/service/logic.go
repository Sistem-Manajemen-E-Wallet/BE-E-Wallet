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
