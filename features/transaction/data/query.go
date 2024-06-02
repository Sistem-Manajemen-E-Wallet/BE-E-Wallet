package data

import (
	"e-wallet/features/product"
	"e-wallet/features/transaction"

	"gorm.io/gorm"
)

type TransactionQuery struct {
	db *gorm.DB
	pd product.DataInterface
}

func New(db *gorm.DB, pd product.DataInterface) transaction.DataInterface {
	return &TransactionQuery{
		db: db,
		pd: pd,
	}
}

// Insert implements transaction.DataInterface.
func (t *TransactionQuery) Insert(input transaction.Core) error {
	result, err := t.pd.SelectProductById(input.ProductID)
	if err != nil {
		return err
	}

	transactionGorm := Transaction{
		Model:          gorm.Model{},
		UserID:         input.UserID,
		ProductID:      input.ProductID,
		Quantity:       input.Quantity,
		TotalCost:      result.Price * input.Quantity,
		StatusProgress: "sedang dimasak",
		Additional:     input.Additional,
		StatusPayment:  "success",
		MerchantID:     result.UserID,
	}

	tx := t.db.Create(&transactionGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAllTransaction implements transaction.DataInterface.
func (t *TransactionQuery) SelectAllTransaction() ([]transaction.Core, error) {
	panic("unimplemented")
}

// SelectTransactionById implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionById(id uint) (*transaction.Core, error) {
	panic("unimplemented")
}

// SelectTransactionByMerchantId implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionByMerchantId(id uint) ([]transaction.Core, error) {
	panic("unimplemented")
}

// UpdateStatusProgress implements transaction.DataInterface.
func (t *TransactionQuery) UpdateStatusProgress(input transaction.Core) error {
	panic("unimplemented")
}
