package service

import (
	"e-wallet/features/wallet"
	"errors"
)

type walletService struct {
	walletData wallet.DataInterface
}

func New(wd wallet.DataInterface) wallet.ServiceInterface {
	return &walletService{
		walletData: wd,
	}
}

// GetWalletById implements wallet.ServiceInterface.
func (w *walletService) GetWalletById(id uint) (wallet.Core, error) {
	if id == 0 {
		return wallet.Core{}, errors.New("invalid user id")
	}
	return w.walletData.GetWalletById(id)
}
