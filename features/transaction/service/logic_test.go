package service

import (
	"e-wallet/features/product"
	"e-wallet/features/transaction"
	"e-wallet/features/wallet"
	"e-wallet/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repoTransactionMock := new(mocks.TransactionData)
	repoWalletMock := new(mocks.WalletData)
	repoProductMock := new(mocks.ProductData)
	transactionService := New(repoTransactionMock, repoWalletMock, repoProductMock)

	t.Run("success create transaction", func(t *testing.T) {
		input := transaction.Core{
			UserID:     1,
			OrderID:    11,
			ProductID:  1,
			Quantity:   1,
			Additional: "pedes cabe 1",
		}
		input2 := wallet.Core{
			Balance: 10000,
		}
		input3 := &product.Core{
			Price: 5000,
		}

		repoWalletMock.On("GetWalletByUserId", input.UserID).Return(input2, nil).Once()
		repoProductMock.On("SelectProductById", input.ProductID).Return(input3, nil).Once()
		repoTransactionMock.On("Insert", input).Return(nil).Once()

		err := transactionService.Create(input)

		assert.NoError(t, err)
		repoWalletMock.AssertExpectations(t)
		repoProductMock.AssertExpectations(t)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("validation error: UserID is zero", func(t *testing.T) {
		input := transaction.Core{
			UserID: 0,
		}

		err := transactionService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "[validation] you must login first", err.Error())
	})

	t.Run("validation error: OrderID/ProductID/Quantity is zero", func(t *testing.T) {
		input := transaction.Core{
			UserID:    1,
			OrderID:   0,
			ProductID: 0,
			Quantity:  0,
		}

		err := transactionService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "[validation] nomor meja/produk/quantity tidak boleh kosong", err.Error())
	})

	t.Run("error getting wallet", func(t *testing.T) {
		input := transaction.Core{
			UserID:    1,
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
		}

		repoWalletMock.On("GetWalletByUserId", input.UserID).Return(wallet.Core{}, errors.New("wallet not found")).Once()

		err := transactionService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "wallet not found", err.Error())
		repoWalletMock.AssertExpectations(t)
	})

	t.Run("error getting product", func(t *testing.T) {
		input := transaction.Core{
			UserID:    1,
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
		}
		input2 := wallet.Core{
			Balance: 10000,
		}

		repoWalletMock.On("GetWalletByUserId", input.UserID).Return(input2, nil).Once()
		repoProductMock.On("SelectProductById", input.ProductID).Return(&product.Core{}, errors.New("product not found")).Once()

		err := transactionService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())
		repoWalletMock.AssertExpectations(t)
		repoProductMock.AssertExpectations(t)
	})

	t.Run("not enough balance", func(t *testing.T) {
		input := transaction.Core{
			UserID:    1,
			OrderID:   1,
			ProductID: 1,
			Quantity:  1,
		}
		input2 := wallet.Core{
			Balance: 1000,
		}
		input3 := &product.Core{
			Price: 5000,
		}

		repoWalletMock.On("GetWalletByUserId", input.UserID).Return(input2, nil).Once()
		repoProductMock.On("SelectProductById", input.ProductID).Return(input3, nil).Once()
		repoTransactionMock.On("Insert", input).Return(nil).Once()

		err := transactionService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "you don't have enough balance", err.Error())
		repoWalletMock.AssertExpectations(t)
		repoProductMock.AssertExpectations(t)
	})
}

func TestGetTransactionById(t *testing.T) {
	repoTransactionMock := new(mocks.TransactionData)
	repoWalletMock := new(mocks.WalletData)
	repoProductMock := new(mocks.ProductData)
	transactionService := New(repoTransactionMock, repoWalletMock, repoProductMock)

	t.Run("success get transaction", func(t *testing.T) {
		input := &transaction.Core{
			ID:     1,
			UserID: 1,
		}

		repoTransactionMock.On("SelectTransactionById", input.ID).Return(input, nil).Once()

		result, err := transactionService.GetTransactionById(input.UserID, input.ID)

		assert.NoError(t, err)
		assert.Equal(t, input, result)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("transaction not found", func(t *testing.T) {
		input := &transaction.Core{
			ID:     2,
			UserID: 1,
		}

		repoTransactionMock.On("SelectTransactionById", input.ID).Return(nil, errors.New("transaction not found")).Once()

		result, err := transactionService.GetTransactionById(input.UserID, input.ID)

		assert.Error(t, err)
		assert.Nil(t, result)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("transaction does not belong to user", func(t *testing.T) {
		currentUserId := 2
		input := &transaction.Core{
			ID:     1,
			UserID: 1,
		}

		repoTransactionMock.On("SelectTransactionById", input.ID).Return(input, nil).Once()

		result, err := transactionService.GetTransactionById(uint(currentUserId), input.ID)

		assert.Error(t, err)
		assert.EqualError(t, err, "this is not your transaction")
		assert.Nil(t, result)
		repoTransactionMock.AssertExpectations(t)
	})
}

func TestGetTransactionByMerchantId(t *testing.T) {
	repoTransactionMock := new(mocks.TransactionData)
	repoWalletMock := new(mocks.WalletData)
	repoProductMock := new(mocks.ProductData)
	transactionService := New(repoTransactionMock, repoWalletMock, repoProductMock)

	t.Run("success get transactions by merchant id", func(t *testing.T) {
		id := uint(1)
		offset := 0
		limit := 10
		transactions := []transaction.Core{
			{ID: 1, MerchantID: id, StatusProgress: "sedang dimasak"},
			{ID: 2, MerchantID: id, StatusProgress: "sudah diantar"},
		}
		count := 2

		repoTransactionMock.On("SelectTransactionByMerchantId", id, offset, limit).Return(transactions, nil).Once()
		repoTransactionMock.On("CountByMerchantId", id).Return(count, nil).Once()

		result, resultCount, err := transactionService.GetTransactionByMerchantId(id, offset, limit)

		assert.NoError(t, err)
		assert.Equal(t, transactions, result)
		assert.Equal(t, count, resultCount)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("transaction not found", func(t *testing.T) {
		id := uint(1)
		offset := 0
		limit := 10

		repoTransactionMock.On("SelectTransactionByMerchantId", id, offset, limit).Return(nil, errors.New("transaction not found")).Once()

		result, resultCount, err := transactionService.GetTransactionByMerchantId(id, offset, limit)

		assert.Error(t, err)
		assert.EqualError(t, err, "transaction not found")
		assert.Nil(t, result)
		assert.Equal(t, 0, resultCount)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("count error", func(t *testing.T) {
		id := uint(1)
		offset := 0
		limit := 10
		transactions := []transaction.Core{
			{ID: 1, MerchantID: id, StatusProgress: "sedang dimasak"},
			{ID: 2, MerchantID: id, StatusProgress: "sudah diantar"},
		}

		repoTransactionMock.On("SelectTransactionByMerchantId", id, offset, limit).Return(transactions, nil).Once()
		repoTransactionMock.On("CountByMerchantId", id).Return(0, errors.New("count error")).Once()

		result, resultCount, err := transactionService.GetTransactionByMerchantId(id, offset, limit)

		assert.Error(t, err)
		assert.EqualError(t, err, "count error")
		assert.Nil(t, result)
		assert.Equal(t, 0, resultCount)
		repoTransactionMock.AssertExpectations(t)
	})
}

func TestUpdateStatusProgress(t *testing.T) {
	repoTransactionMock := new(mocks.TransactionData)
	repoWalletMock := new(mocks.WalletData)
	repoProductMock := new(mocks.ProductData)
	transactionService := New(repoTransactionMock, repoWalletMock, repoProductMock)

	t.Run("success update status progress", func(t *testing.T) {
		idUser := uint(1)
		idTransaction := uint(1)
		input := transaction.Core{
			ID:             idTransaction,
			StatusProgress: "sudah diantar",
		}
		input2 := &transaction.Core{
			ID:         idTransaction,
			MerchantID: idUser,
		}

		repoTransactionMock.On("SelectTransactionById", idTransaction).Return(input2, nil).Once()
		repoTransactionMock.On("UpdateStatusProgress", idTransaction, input).Return(nil).Once()

		err := transactionService.UpdateStatusProgress(idUser, idTransaction, input)

		assert.NoError(t, err)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("transaction not found", func(t *testing.T) {
		idUser := uint(1)
		idTransaction := uint(2)
		input := transaction.Core{
			ID:             idTransaction,
			StatusProgress: "sudah diantar",
		}

		repoTransactionMock.On("SelectTransactionById", idTransaction).Return(nil, errors.New("transaction not found")).Once()

		err := transactionService.UpdateStatusProgress(idUser, idTransaction, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "transaction not found")
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("transaction does not belong to user", func(t *testing.T) {
		idUser := uint(1)
		idTransaction := uint(2)
		CurrentUserId := uint(3)
		input := transaction.Core{
			ID:             idTransaction,
			StatusProgress: "sudah diantar",
		}
		input2 := &transaction.Core{
			ID:         idTransaction,
			MerchantID: CurrentUserId,
		}

		repoTransactionMock.On("SelectTransactionById", idTransaction).Return(input2, nil).Once()

		err := transactionService.UpdateStatusProgress(idUser, idTransaction, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "it's not your transaction")
		repoTransactionMock.AssertExpectations(t)
	})
}

func TestVerifyPin(t *testing.T) {
	repoTransactionMock := new(mocks.TransactionData)
	repoWalletMock := new(mocks.WalletData)
	repoProductMock := new(mocks.ProductData)
	transactionService := New(repoTransactionMock, repoWalletMock, repoProductMock)

	t.Run("verify successful", func(t *testing.T) {
		idUser := uint(1)
		pin := "123456"

		repoTransactionMock.On("VerifyPin", pin, idUser).Return(nil).Once()

		err := transactionService.VerifyPin(pin, idUser)

		assert.NoError(t, err)
		repoTransactionMock.AssertExpectations(t)
	})

	t.Run("error: UserID is zero", func(t *testing.T) {
		idUser := uint(0)
		pin := "123456"

		err := transactionService.VerifyPin(pin, idUser)

		assert.Error(t, err)
		assert.EqualError(t, err, "you must login first")
		repoTransactionMock.AssertExpectations(t)
	})
}
