package service

import (
	"e-wallet/features/history"
	"errors"
)

type HistoryService struct {
	hd history.DataInterface
}

func New(hd history.DataInterface) history.ServiceInterface {
	return &HistoryService{
		hd: hd,
	}
}

// GetAllHistory implements history.ServiceInterface.
func (h *HistoryService) GetAllHistory(idUser uint, offset int, limit int) ([]history.Core, int, error) {

	result, err := h.hd.SelectAllHistory(idUser, offset, limit)
	if err != nil {
		return nil, 0, errors.New("history not found")
	}

	totalHistory, err := h.hd.CountHistory(idUser)
	if err != nil {
		return nil, 0, errors.New("history not found")
	}

	return result, totalHistory, nil
}
