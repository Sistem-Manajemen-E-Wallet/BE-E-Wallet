package data

import (
	"e-wallet/features/history"
	"e-wallet/features/product"
	"e-wallet/features/transaction"
	"e-wallet/features/user"
	"e-wallet/features/wallet"
	encrypts "e-wallet/utils"
	"errors"

	"gorm.io/gorm"
)

type TransactionQuery struct {
	db *gorm.DB
	pd product.DataInterface
	hd history.DataInterface
	wd wallet.DataInterface
	ud user.DataInterface
	eh encrypts.HashInterface
}

func New(db *gorm.DB, pd product.DataInterface, hd history.DataInterface, wd wallet.DataInterface, ud user.DataInterface, eh encrypts.HashInterface) transaction.DataInterface {
	return &TransactionQuery{
		db: db,
		pd: pd,
		hd: hd,
		wd: wd,
		ud: ud,
		eh: eh,
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

	err3 := t.wd.UpdateBalanceMinus(input.UserID, transactionGorm.TotalCost)
	if err3 != nil {
		return err3
	}

	err4 := t.wd.UpdateBalancePlus(result.UserID, transactionGorm.TotalCost)
	if err4 != nil {
		return err4
	}

	tx := t.db.Create(&transactionGorm)
	if tx.Error != nil {
		return tx.Error
	}

	historyGorm := history.Core{
		UserID:        transactionGorm.UserID,
		TransactionID: transactionGorm.ID,
		TrxName:       result.ProductName,
		Amount:        transactionGorm.TotalCost,
		Type:          "payment",
		Status:        transactionGorm.StatusPayment,
		CreatedAt:     transactionGorm.CreatedAt,
	}

	err2 := t.hd.InsertHistory(historyGorm)
	if err2 != nil {
		return err2
	}
	return nil
}

// SelectTransactionById implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionById(id uint) (*transaction.Core, error) {
	var currentTransaction Transaction
	tx := t.db.First(&currentTransaction, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	currentCore := transaction.Core{
		ID:             id,
		UserID:         currentTransaction.UserID,
		OrderID:        currentTransaction.OrderID,
		ProductID:      currentTransaction.ProductID,
		Quantity:       currentTransaction.Quantity,
		TotalCost:      currentTransaction.TotalCost,
		StatusProgress: currentTransaction.StatusProgress,
		Additional:     currentTransaction.Additional,
		StatusPayment:  currentTransaction.StatusPayment,
		MerchantID:     currentTransaction.MerchantID,
		CreatedAt:      currentTransaction.CreatedAt,
		UpdatedAt:      currentTransaction.UpdatedAt,
	}

	return &currentCore, nil
}

// SelectTransactionByMerchantId implements transaction.DataInterface.
func (t *TransactionQuery) SelectTransactionByMerchantId(id uint, offset int, limit int) ([]transaction.Core, error) {
	var transactionGorm []Transaction
	tx := t.db.Where("merchant_id = ?", id).Offset(offset).Limit(limit).Find(&transactionGorm)
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
func (t *TransactionQuery) UpdateStatusProgress(id uint, input transaction.Core) error {
	updateStatus := Transaction{
		StatusProgress: input.StatusProgress,
	}
	tx := t.db.Model(&Transaction{}).Where("id = ?", id).Updates(updateStatus)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CountByMerchantId implements transaction.DataInterface.
func (t *TransactionQuery) CountByMerchantId(merchantId uint) (int, error) {
	var count int64
	tx := t.db.Model(&Transaction{}).Where("merchant_id = ?", merchantId).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

// VerifyPin implements transaction.DataInterface.
func (t *TransactionQuery) VerifyPin(pin string, idUser uint) error {
	result, err := t.ud.SelectProfileById(idUser)
	if err != nil {
		return err
	}

	isVerifyValid := t.eh.CheckPasswordHash(result.Pin, pin)
	if !isVerifyValid {
		return errors.New("pin tidak sesuai")
	}

	return nil
}
