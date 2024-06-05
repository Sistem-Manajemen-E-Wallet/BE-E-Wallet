package service

import (
	"e-wallet/features/wallet"
	"e-wallet/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetWalletByUserId(t *testing.T) {
	repoWalletMock := mocks.WalletData{}
	service := New(&repoWalletMock)
	t.Run("failed due to invalid id", func(t *testing.T) {
		_, err := service.GetWalletByUserId(0)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid user id")
		repoWalletMock.AssertNotCalled(t, "GetWalletByUserId", mock.Anything)
	})
	t.Run("success get wallet by user id", func(t *testing.T) {
		input := wallet.Core{
			UserID:  1,
			Balance: 10000,
		}

		repoWalletMock.On("GetWalletByUserId", input.UserID).Return(input, nil).Once()
		srv := New(&repoWalletMock)

		result, err := srv.GetWalletByUserId(input.UserID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.UserID, result.UserID)
		repoWalletMock.AssertExpectations(t)
	})

	t.Run("failed to get wallet data", func(t *testing.T) {
		repoWalletMock.On("GetWalletByUserId", uint(1)).Return(wallet.Core{}, errors.New("data not found")).Once()

		_, err := repoWalletMock.GetWalletByUserId(1)
		assert.Error(t, err)
		assert.EqualError(t, err, "data not found")
		repoWalletMock.AssertExpectations(t)
	})
}
