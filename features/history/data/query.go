package data

import (
	"e-wallet/features/history"

	"gorm.io/gorm"
)

type HistoryQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) history.DataInterface {
	return &HistoryQuery{
		db: db,
	}
}

// InsertHistory implements history.DataInterface.
func (h *HistoryQuery) InsertHistory(input history.Core) error {
	historyGorm := History{
		Model:         gorm.Model{},
		UserID:        input.UserID,
		TransactionID: input.TransactionID,
		TopUpID:       input.TopUpID,
		TrxName:       input.TrxName,
		Amount:        input.Amount,
		Type:          input.Type,
		Status:        input.Status,
	}
	tx := h.db.Create(&historyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAllHistory implements history.DataInterface.
func (h *HistoryQuery) SelectAllHistory(idUser uint) ([]history.Core, error) {
	var historyGorm []History
	tx := h.db.Model(&History{}).Where("user_id = ?", idUser).Find(&historyGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var historyCore []history.Core
	for _, v := range historyGorm {
		historyCore = append(historyCore, history.Core{
			ID:            v.ID,
			UserID:        idUser,
			TransactionID: v.TransactionID,
			TopUpID:       v.TopUpID,
			TrxName:       v.TrxName,
			Amount:        v.Amount,
			Type:          v.Type,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return historyCore, nil
}

func (h *HistoryQuery) UpdateHistoryTopUp(input history.Core) error {
	historyGorm := History{
		Status: input.Status,
	}
	tx := h.db.Model(&History{}).Where("top_up_id = ?", input.TopUpID).Updates(&historyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
