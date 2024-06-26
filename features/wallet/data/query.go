package data

import (
	"e-wallet/features/wallet"
	"errors"
	"time"

	"gorm.io/gorm"
)

type walletQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) wallet.DataInterface {
	return &walletQuery{
		db: db,
	}
}

// CreateWallet implements wallet.DataInterface.
func (w *walletQuery) CreateWallet(id uint) error {
	walletGorm := Wallet{
		Model:   gorm.Model{},
		UserID:  id,
		Balance: 0,
	}

	tx := w.db.Create(&walletGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetWalletById implements wallet.DataInterface.
func (w *walletQuery) GetWalletById(id uint) (wallet.Core, error) {
	var walletId Wallet
	tx := w.db.First(&walletId, id)
	if tx.Error != nil {
		return wallet.Core{}, tx.Error
	}

	walletCore := wallet.Core{
		ID:        walletId.ID,
		UserID:    id,
		Balance:   walletId.Balance,
		CreatedAt: walletId.CreatedAt,
		UpdatedAt: walletId.UpdatedAt,
	}

	return walletCore, nil
}

// UpdateBalance implements wallet.DataInterface.
func (w *walletQuery) UpdateBalanceMinus(id uint, amount int) error {
	result, err := w.GetWalletByUserId(id)
	if err != nil {
		return err
	}

	if result.Balance < amount {
		return errors.New("your balance is not enough")
	}

	substraction := result.Balance - amount

	tx := w.db.Model(&Wallet{}).Where("user_id = ?", id).Updates(map[string]interface{}{
		"balance":    substraction,
		"updated_at": time.Now(),
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// UpdateBalance implements wallet.DataInterface.
func (w *walletQuery) UpdateBalancePlus(id uint, amount int) error {
	result, err := w.GetWalletByUserId(id)
	if err != nil {
		return err
	}

	summation := result.Balance + amount

	tx := w.db.Model(&Wallet{}).Where("user_id = ?", id).Updates(map[string]interface{}{
		"balance":    summation,
		"updated_at": time.Now(),
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (w *walletQuery) UpdateBalanceByTopup(input wallet.Core) error {
	walletGorm := Wallet{
		UserID:  input.UserID,
		Balance: input.Balance,
	}

	tx := w.db.Model(&Wallet{}).Where("user_id = ?", input.UserID).Updates(walletGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (w *walletQuery) GetWalletByUserId(id uint) (wallet.Core, error) {
	var walletGorm Wallet
	tx := w.db.Where("user_id = ?", id).First(&walletGorm)
	if tx.Error != nil {
		return wallet.Core{}, tx.Error
	}
	return wallet.Core{
		ID:        walletGorm.ID,
		UserID:    walletGorm.UserID,
		Balance:   walletGorm.Balance,
		CreatedAt: walletGorm.CreatedAt,
		UpdatedAt: walletGorm.UpdatedAt,
	}, nil
}
