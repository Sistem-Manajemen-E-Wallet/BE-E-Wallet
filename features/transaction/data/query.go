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
		OrderID:        input.OrderID,
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

// SelectTransactionById implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionById(id uint) (*transaction.Core, error) {
	panic("unimplemented")
}

// SelectTransactionByMerchantId implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionByMerchantId(id uint) ([]transaction.Core, error) {
	var transactionGorm []Transaction
	tx := t.db.Where("merchant_id = ?", id).Find(&transactionGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var transactionCore []transaction.Core
	for _, v := range transactionGorm {
		transactionCore = append(transactionCore, transaction.Core{
			ID:             v.ID,
			UserID:         v.UserID,
			OrderID:        v.OrderID,
			ProductID:      v.ProductID,
			Quantity:       v.Quantity,
			TotalCost:      v.TotalCost,
			StatusProgress: v.StatusProgress,
			Additional:     v.Additional,
			StatusPayment:  v.StatusPayment,
			MerchantID:     id,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		})
	}
	return transactionCore, nil
}

// UpdateStatusProgress implements transaction.DataInterface.
func (t *TransactionQuery) UpdateStatusProgress(input transaction.Core) error {
	panic("unimplemented")
}
