package service

import (
	"e-wallet/features/topups"
	"e-wallet/features/user"
	"e-wallet/features/wallet"
	"e-wallet/mocks"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreate(t *testing.T) {
	topupDataMock := new(mocks.TopUpsData)
	userDataMock := new(mocks.UserData)
	walletDataMock := new(mocks.WalletData)
	historyDataMock := new(mocks.HistoryData)
	midtransMock := new(mocks.Service)
	validate := validator.New()

	//svc := New(topupDataMock, walletDataMock, userDataMock, historyDataMock)
	svc := &topupsService{
		topupData:   topupDataMock,
		userData:    userDataMock,
		walletData:  walletDataMock,
		historyData: historyDataMock,
		midtrans:    midtransMock,
		validate:    validate,
	}

	t.Run("success", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "Customer"}, nil).Once()
		midtransMock.On("GetVaNumbers", mock.Anything).Return("va123", nil).Once()
		topupDataMock.On("Insert", mock.Anything).Return(input, nil).Once()
		historyDataMock.On("InsertHistory", mock.Anything).Return(nil).Once()

		result, err := svc.Create(input)
		assert.NoError(t, err)
		assert.Equal(t, input, result)

		mock.AssertExpectationsForObjects(t, topupDataMock, userDataMock, midtransMock, historyDataMock)
	})

	t.Run("validation error", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      0,
			ChannelBank: "bca",
		}

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "[validation error]")
	})

	t.Run("invalid channel bank", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "invalid_bank",
		}

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "invalid channel bank", err.Error())
	})

	t.Run("user not found", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found")).Once()

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())

		userDataMock.AssertExpectations(t)
	})

	t.Run("user not authorized", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "Merchant"}, nil).Once()

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "user not authorized", err.Error())

		userDataMock.AssertExpectations(t)
	})

	t.Run("error getting va numbers", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "Customer"}, nil).Once()
		midtransMock.On("GetVaNumbers", mock.Anything).Return("", errors.New("error getting va numbers")).Once()

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "error getting va numbers", err.Error())

		mock.AssertExpectationsForObjects(t, userDataMock, midtransMock)
	})

	t.Run("error inserting topup", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "Customer"}, nil).Once()
		midtransMock.On("GetVaNumbers", mock.Anything).Return("va123", nil).Once()
		topupDataMock.On("Insert", mock.Anything).Return(topups.Core{}, errors.New("error inserting topup")).Once()

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "error inserting topup", err.Error())

		mock.AssertExpectationsForObjects(t, userDataMock, midtransMock, topupDataMock)
	})

	t.Run("error inserting history", func(t *testing.T) {
		input := topups.Core{
			UserID:      1,
			Amount:      10000,
			ChannelBank: "bca",
		}

		userDataMock.On("SelectProfileById", uint(1)).Return(&user.Core{Role: "Customer"}, nil).Once()
		midtransMock.On("GetVaNumbers", mock.Anything).Return("va123", nil).Once()
		topupDataMock.On("Insert", mock.Anything).Return(input, nil).Once()
		historyDataMock.On("InsertHistory", mock.Anything).Return(errors.New("error inserting history")).Once()

		_, err := svc.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "error inserting history", err.Error())

		mock.AssertExpectationsForObjects(t, userDataMock, midtransMock, topupDataMock, historyDataMock)
	})
}

func TestGetByUserID(t *testing.T) {
	topupDataMock := new(mocks.TopUpsData)
	userDataMock := new(mocks.UserData)
	walletDataMock := new(mocks.WalletData)
	historyDataMock := new(mocks.HistoryData)
	midtransMock := new(mocks.Service)
	validate := validator.New()

	//svc := New(topupDataMock, walletDataMock, userDataMock, historyDataMock)
	svc := &topupsService{
		topupData:   topupDataMock,
		userData:    userDataMock,
		walletData:  walletDataMock,
		historyData: historyDataMock,
		midtrans:    midtransMock,
		validate:    validate,
	}

	t.Run("success", func(t *testing.T) {
		userID := 1
		topups := []topups.Core{
			{ID: 1, UserID: userID, Amount: 10000, ChannelBank: "bca"},
			{ID: 2, UserID: userID, Amount: 5000, ChannelBank: "bni"},
		}

		topupDataMock.On("SelectByUserID", userID).Return(topups, nil).Once()

		result, err := svc.GetByUserID(userID)
		assert.NoError(t, err)
		assert.Equal(t, topups, result)

		topupDataMock.AssertExpectations(t)
	})

	t.Run("topup not found", func(t *testing.T) {
		userID := 1

		topupDataMock.On("SelectByUserID", userID).Return(nil, errors.New("topup not found")).Once()

		result, err := svc.GetByUserID(userID)
		assert.Error(t, err)
		assert.Equal(t, []topups.Core{}, result)
		assert.Equal(t, "topup not found", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("no topups found", func(t *testing.T) {
		userID := 1

		topupDataMock.On("SelectByUserID", userID).Return([]topups.Core{}, nil).Once()

		result, err := svc.GetByUserID(userID)
		assert.NoError(t, err)
		assert.Empty(t, result)

		topupDataMock.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		userID := 1

		topupDataMock.On("SelectByUserID", userID).Return(nil, errors.New("database error")).Once()

		result, err := svc.GetByUserID(userID)
		assert.Error(t, err)
		assert.Equal(t, []topups.Core{}, result)
		assert.Equal(t, "topup not found", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("invalid user id", func(t *testing.T) {
		invalidUserID := -5

		result, err := svc.GetByUserID(invalidUserID)
		assert.Error(t, err)
		assert.Equal(t, []topups.Core{}, result)
		assert.Equal(t, "invalid user id", err.Error())
	})
}

func TestGetByID(t *testing.T) {
	topupDataMock := new(mocks.TopUpsData)
	userDataMock := new(mocks.UserData)
	walletDataMock := new(mocks.WalletData)
	historyDataMock := new(mocks.HistoryData)
	midtransMock := new(mocks.Service)
	validate := validator.New()

	svc := &topupsService{
		topupData:   topupDataMock,
		userData:    userDataMock,
		walletData:  walletDataMock,
		historyData: historyDataMock,
		midtrans:    midtransMock,
		validate:    validate,
	}

	t.Run("success", func(t *testing.T) {
		topupID := 1
		userID := 1
		topup := topups.Core{
			ID:          topupID,
			UserID:      userID,
			Amount:      10000,
			ChannelBank: "bca",
		}

		topupDataMock.On("SelectById", topupID).Return(topup, nil).Once()

		result, err := svc.GetByID(topupID, userID)
		assert.NoError(t, err)
		assert.Equal(t, topup, result)

		topupDataMock.AssertExpectations(t)
	})

	t.Run("invalid topup id", func(t *testing.T) {
		invalidTopupID := 0
		userID := 1

		result, err := svc.GetByID(invalidTopupID, userID)
		assert.Error(t, err)
		assert.Equal(t, "invalid topup id", err.Error())
		assert.Equal(t, topups.Core{}, result)
	})

	t.Run("topup not found", func(t *testing.T) {
		topupID := 1
		userID := 1

		topupDataMock.On("SelectById", topupID).Return(topups.Core{}, errors.New("topup not found")).Once()

		result, err := svc.GetByID(topupID, userID)
		assert.Error(t, err)
		assert.Equal(t, "topup not found", err.Error())
		assert.Equal(t, topups.Core{}, result)

		topupDataMock.AssertExpectations(t)
	})

	t.Run("user not authorized", func(t *testing.T) {
		topupID := 1
		userID := 1
		unauthorizedUserID := 2

		topup := topups.Core{
			ID:          topupID,
			UserID:      unauthorizedUserID,
			Amount:      10000,
			ChannelBank: "bca",
		}

		topupDataMock.On("SelectById", topupID).Return(topup, nil).Once()

		result, err := svc.GetByID(topupID, userID)
		assert.Error(t, err)
		assert.Equal(t, "user not authorized", err.Error())
		assert.Equal(t, topups.Core{}, result)

		topupDataMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	topupDataMock := new(mocks.TopUpsData)
	userDataMock := new(mocks.UserData)
	walletDataMock := new(mocks.WalletData)
	historyDataMock := new(mocks.HistoryData)
	midtransMock := new(mocks.Service)
	validate := validator.New()

	svc := &topupsService{
		topupData:   topupDataMock,
		userData:    userDataMock,
		walletData:  walletDataMock,
		historyData: historyDataMock,
		midtrans:    midtransMock,
		validate:    validate,
	}

	t.Run("success", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
			UserID:  1,
			Amount:  10000,
		}
		wallet := wallet.Core{
			UserID:  1,
			Balance: 5000,
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()
		topupDataMock.On("Update", int(topup.ID), mock.Anything).Return(nil).Once()
		walletDataMock.On("GetWalletByUserId", uint(topup.UserID)).Return(wallet, nil).Once()
		walletDataMock.On("UpdateBalanceByTopup", mock.Anything).Return(nil).Once()
		historyDataMock.On("UpdateHistoryTopUp", mock.Anything).Return(nil).Once()

		err := svc.Update(input)
		assert.NoError(t, err)

		topupDataMock.AssertExpectations(t)
		walletDataMock.AssertExpectations(t)
		historyDataMock.AssertExpectations(t)
	})

	t.Run("topup not found", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topups.Core{}, errors.New("topup not found")).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "topup not found", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("topup status cannot be updated", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "paid",
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "topup status cannot be updated", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("invalid topup status", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "pending",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "invalid topup status", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("error updating topup", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()
		topupDataMock.On("Update", int(topup.ID), mock.Anything).Return(errors.New("error updating topup")).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "error updating topup", err.Error())

		topupDataMock.AssertExpectations(t)
	})

	t.Run("wallet not found", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
			UserID:  1,
			Amount:  10000,
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()
		topupDataMock.On("Update", int(topup.ID), mock.Anything).Return(nil).Once()
		walletDataMock.On("GetWalletByUserId", uint(topup.UserID)).Return(wallet.Core{}, errors.New("wallet not found")).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "wallet not found", err.Error())

		topupDataMock.AssertExpectations(t)
		walletDataMock.AssertExpectations(t)
	})

	t.Run("error updating wallet", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
			UserID:  1,
			Amount:  10000,
		}
		wallet := wallet.Core{
			UserID:  1,
			Balance: 5000,
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()
		topupDataMock.On("Update", int(topup.ID), mock.Anything).Return(nil).Once()
		walletDataMock.On("GetWalletByUserId", uint(topup.UserID)).Return(wallet, nil).Once()
		walletDataMock.On("UpdateBalanceByTopup", mock.Anything).Return(errors.New("error updating wallet")).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "error updating wallet", err.Error())

		topupDataMock.AssertExpectations(t)
		walletDataMock.AssertExpectations(t)
	})

	t.Run("error updating history", func(t *testing.T) {
		input := topups.Core{
			OrderID: "order123",
			Status:  "settlement",
		}
		topup := topups.Core{
			ID:      1,
			OrderID: "order123",
			Status:  "pending",
			UserID:  1,
			Amount:  10000,
		}
		wallet := wallet.Core{
			UserID:  1,
			Balance: 5000,
		}

		topupDataMock.On("SelectByOrderID", "order123").Return(topup, nil).Once()
		topupDataMock.On("Update", int(topup.ID), mock.Anything).Return(nil).Once()
		walletDataMock.On("GetWalletByUserId", uint(topup.UserID)).Return(wallet, nil).Once()
		walletDataMock.On("UpdateBalanceByTopup", mock.Anything).Return(nil).Once()
		historyDataMock.On("UpdateHistoryTopUp", mock.Anything).Return(errors.New("error updating history")).Once()

		err := svc.Update(input)
		assert.Error(t, err)
		assert.Equal(t, "error updating history", err.Error())

		topupDataMock.AssertExpectations(t)
		walletDataMock.AssertExpectations(t)
		historyDataMock.AssertExpectations(t)
	})
}
