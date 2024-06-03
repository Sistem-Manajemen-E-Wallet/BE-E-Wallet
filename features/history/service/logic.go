package service

import "e-wallet/features/history"

type HistoryService struct {
	hd history.DataInterface
}

func New(hd history.DataInterface) history.ServiceInterface {
	return &HistoryService{
		hd: hd,
	}
}

// GetAllHistory implements history.ServiceInterface.
func (h *HistoryService) GetAllHistory(idUser uint) ([]history.Core, error) {
	return h.hd.SelectAllHistory(idUser)
}
