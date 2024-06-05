package service

import (
	"e-wallet/features/history"
	"e-wallet/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllHistory(t *testing.T) {
	mockHistoryData := new(mocks.HistoryData)
	historyService := New(mockHistoryData)

	userID := uint(1)
	offset := 0
	limit := 10

	mockHistories := []history.Core{
		{
			ID:            1,
			UserID:        userID,
			TransactionID: 1,
			TopUpID:       1,
			TrxName:       "topup",
			Amount:        10000,
			Type:          "topup",
			Status:        "paid",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            2,
			UserID:        userID,
			TransactionID: 1,
			TopUpID:       1,
			TrxName:       "topup",
			Amount:        10000,
			Type:          "topup",
			Status:        "paid",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	mockCount := 2

	t.Run("Success", func(t *testing.T) {
		mockHistoryData.On("SelectAllHistory", userID, offset, limit).Return(mockHistories, nil).Once()
		mockHistoryData.On("CountHistory", userID).Return(mockCount, nil).Once()

		histories, count, err := historyService.GetAllHistory(userID, offset, limit)

		assert.NoError(t, err)
		assert.Equal(t, mockHistories, histories)
		assert.Equal(t, mockCount, count)

		mockHistoryData.AssertExpectations(t)
	})

	t.Run("Error in SelectAllHistory", func(t *testing.T) {
		mockHistoryData.On("SelectAllHistory", userID, offset, limit).Return(nil, errors.New("database error")).Once()

		histories, count, err := historyService.GetAllHistory(userID, offset, limit)

		assert.Error(t, err)
		assert.Nil(t, histories)
		assert.Equal(t, 0, count)
		assert.Equal(t, "history not found", err.Error())

		mockHistoryData.AssertExpectations(t)
	})

	t.Run("Error in CountHistory", func(t *testing.T) {
		mockHistoryData.On("SelectAllHistory", userID, offset, limit).Return(mockHistories, nil).Once()
		mockHistoryData.On("CountHistory", userID).Return(0, errors.New("database error")).Once()

		histories, count, err := historyService.GetAllHistory(userID, offset, limit)

		assert.Error(t, err)
		assert.Nil(t, histories)
		assert.Equal(t, 0, count)
		assert.Equal(t, "history not found", err.Error())

		mockHistoryData.AssertExpectations(t)
	})
}
