package data

import (
	"e-wallet/features/topups"
	"gorm.io/gorm"
)

type topupQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) topups.DataInterface {
	return &topupQuery{
		db: db,
	}
}

func (t *topupQuery) Insert(input topups.Core) (topups.Core, error) {
	topupGorm := TopUp{
		OrderID:     input.OrderID,
		UserID:      input.UserID,
		Amount:      input.Amount,
		Type:        input.Type,
		ChannelBank: input.ChannelBank,
		VaNumbers:   input.VaNumbers,
		Status:      input.Status,
	}

	tx := t.db.Create(&topupGorm)
	if tx.Error != nil {
		return topups.Core{}, tx.Error
	}
	return topups.Core{
		ID:          int(topupGorm.ID),
		OrderID:     topupGorm.OrderID,
		UserID:      topupGorm.UserID,
		Amount:      topupGorm.Amount,
		Type:        topupGorm.Type,
		ChannelBank: topupGorm.ChannelBank,
		VaNumbers:   topupGorm.VaNumbers,
		Status:      topupGorm.Status,
		CreatedAt:   topupGorm.CreatedAt,
		UpdatedAt:   topupGorm.UpdatedAt,
	}, nil

	//tx := t.db.Create(&topupGorm)
	//if tx.Error != nil {
	//	return tx.Error
	//}
	//return nil
}

func (t *topupQuery) SelectByUserID(id int) ([]topups.Core, error) {
	var topupGorm []TopUp
	tx := t.db.Where("user_id = ?", id).Find(&topupGorm)
	if tx.Error != nil {
		return []topups.Core{}, tx.Error
	}

	var topupCore []topups.Core
	for _, v := range topupGorm {
		topupCore = append(topupCore, topups.Core{
			ID:          int(v.ID),
			OrderID:     v.OrderID,
			UserID:      v.UserID,
			Amount:      v.Amount,
			Type:        v.Type,
			ChannelBank: v.ChannelBank,
			VaNumbers:   v.VaNumbers,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return topupCore, nil
}

func (t *topupQuery) Update(id int, input topups.Core) error {
	topupGorm := TopUp{
		UserID:      input.UserID,
		OrderID:     input.OrderID,
		Amount:      input.Amount,
		Type:        input.Type,
		ChannelBank: input.ChannelBank,
		VaNumbers:   input.VaNumbers,
		Status:      input.Status,
	}
	tx := t.db.Model(&topupGorm).Where("id = ?", id).Updates(topupGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (t *topupQuery) SelectById(id int) (topups.Core, error) {
	var topupGorm TopUp
	tx := t.db.Where("id = ?", id).First(&topupGorm)
	if tx.Error != nil {
		return topups.Core{}, tx.Error
	}
	return topups.Core{
		ID:          int(topupGorm.ID),
		UserID:      topupGorm.UserID,
		OrderID:     topupGorm.OrderID,
		Amount:      topupGorm.Amount,
		Type:        topupGorm.Type,
		ChannelBank: topupGorm.ChannelBank,
		VaNumbers:   topupGorm.VaNumbers,
		Status:      topupGorm.Status,
		CreatedAt:   topupGorm.CreatedAt,
		UpdatedAt:   topupGorm.UpdatedAt,
	}, nil
}

func (t *topupQuery) SelectByOrderID(id string) (topups.Core, error) {
	var topupGorm TopUp
	tx := t.db.Where("order_id = ?", id).First(&topupGorm)
	if tx.Error != nil {
		return topups.Core{}, tx.Error
	}
	return topups.Core{
		ID:          int(topupGorm.ID),
		UserID:      topupGorm.UserID,
		OrderID:     topupGorm.OrderID,
		Amount:      topupGorm.Amount,
		Type:        topupGorm.Type,
		ChannelBank: topupGorm.ChannelBank,
		VaNumbers:   topupGorm.VaNumbers,
		Status:      topupGorm.Status,
		CreatedAt:   topupGorm.CreatedAt,
		UpdatedAt:   topupGorm.UpdatedAt,
	}, nil
}
