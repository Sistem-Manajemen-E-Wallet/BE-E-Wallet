package service

import (
	"e-wallet/features/product"
	"e-wallet/features/user"
	"e-wallet/mocks"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetAllProduct(t *testing.T) {
	mockProductData := new(mocks.ProductData)
	mockUserData := new(mocks.UserData)
	validate := validator.New()

	service := productService{
		productData: mockProductData,
		userData:    mockUserData,
		validate:    validate,
	}

	t.Run("Success", func(t *testing.T) {
		expectedProducts := []product.Core{
			{ProductName: "Product1"},
			{ProductName: "Product2"},
		}
		expectedTotal := 2

		mockProductData.On("SelectAllProduct", 0, 2).Return(expectedProducts, nil).Once()
		mockProductData.On("CountProduct").Return(expectedTotal, nil).Once()

		products, total, err := service.GetAllProduct(0, 2)

		assert.Nil(t, err)
		assert.Equal(t, expectedProducts, products)
		assert.Equal(t, expectedTotal, total)

		mockProductData.AssertExpectations(t)
	})

	t.Run("Error when fetching products", func(t *testing.T) {
		mockProductData.On("SelectAllProduct", 0, 2).Return(nil, errors.New("product not found")).Once()

		products, total, err := service.GetAllProduct(0, 2)

		assert.NotNil(t, err)
		assert.Nil(t, products)
		assert.Equal(t, 0, total)
		assert.EqualError(t, err, "product not found")

		mockProductData.AssertExpectations(t)
	})

	t.Run("Error when counting products", func(t *testing.T) {
		expectedProducts := []product.Core{
			{ProductName: "Product1"},
			{ProductName: "Product2"},
		}

		mockProductData.On("SelectAllProduct", 0, 2).Return(expectedProducts, nil).Once()
		mockProductData.On("CountProduct").Return(0, errors.New("product not found")).Once()

		products, total, err := service.GetAllProduct(0, 2)

		assert.NotNil(t, err)
		assert.Nil(t, products)
		assert.Equal(t, 0, total)
		assert.EqualError(t, err, "product not found")

		mockProductData.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockProductData := new(mocks.ProductData)
	mockUserData := new(mocks.UserData)
	validate := validator.New()

	service := productService{
		productData: mockProductData,
		userData:    mockUserData,
		validate:    validate,
	}

	t.Run("Success", func(t *testing.T) {
		input := product.Core{
			UserID:      1,
			ProductName: "Product1",
			Description: "Description1",
			Price:       1000,
		}
		mockUser := user.Core{ID: 1, Role: "Merchant"}

		mockUserData.On("SelectProfileById", uint(1)).Return(&mockUser, nil).Once()
		mockProductData.On("Insert", mock.Anything).Return(nil).Once()

		err := service.Create(input)

		assert.Nil(t, err)
		mockUserData.AssertExpectations(t)
		mockProductData.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		input := product.Core{
			UserID: 1,
		}

		err := service.Create(input)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "[validation error]")
	})

	t.Run("User Not Found", func(t *testing.T) {
		input := product.Core{
			UserID:      1,
			ProductName: "Product1",
			Description: "Description1",
			Price:       1000,
		}

		mockUserData.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found")).Once()

		err := service.Create(input)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "user not found")
		mockUserData.AssertExpectations(t)
	})

	t.Run("User Not Authorized", func(t *testing.T) {
		input := product.Core{
			UserID:      1,
			ProductName: "Product1",
			Description: "Description1",
			Price:       1000,
		}
		mockUser := user.Core{ID: 1, Role: "Customer"}

		mockUserData.On("SelectProfileById", uint(1)).Return(&mockUser, nil).Once()

		err := service.Create(input)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "user not authorized")
		mockUserData.AssertExpectations(t)
	})
}

func TestGetProductById(t *testing.T) {
	mockProductData := new(mocks.ProductData)
	mockUserData := new(mocks.UserData)

	productService := productService{
		productData: mockProductData,
		userData:    mockUserData,
	}

	t.Run("success", func(t *testing.T) {
		mockProduct := product.Core{
			ID:          1,
			ProductName: "Product 1",
			Description: "Description 1",
			Price:       1000,
			UserID:      1,
		}
		mockProductData.On("SelectProductById", uint(1)).Return(&mockProduct, nil)

		result, err := productService.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockProduct.ProductName, result.ProductName)
		mockProductData.AssertExpectations(t)
	})

	t.Run("invalid id", func(t *testing.T) {
		result, err := productService.GetProductById(0)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "id not valid", err.Error())
	})

	t.Run("product not found", func(t *testing.T) {
		mockProductData.On("SelectProductById", uint(2)).Return(nil, errors.New("product not found"))

		result, err := productService.GetProductById(2)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "product not found", err.Error())
		mockProductData.AssertExpectations(t)
	})
}

func TestGetProductByUserId(t *testing.T) {
	mockProductData := new(mocks.ProductData)
	mockUserData := new(mocks.UserData)
	productService := productService{
		productData: mockProductData,
		userData:    mockUserData,
	}

	t.Run("success", func(t *testing.T) {
		mockUser := user.Core{
			ID:   1,
			Role: "Merchant",
		}
		mockProducts := []product.Core{
			{ID: 1, ProductName: "Product 1", UserID: 1},
			{ID: 2, ProductName: "Product 2", UserID: 1},
		}
		mockTotalProducts := 2

		mockUserData.On("SelectProfileById", uint(1)).Return(&mockUser, nil)
		mockProductData.On("SelectProductByUserId", uint(1), 0, 10).Return(mockProducts, nil)
		mockProductData.On("CountProductByUserId", uint(1)).Return(mockTotalProducts, nil)

		result, total, err := productService.GetProductByUserId(1, 0, 10)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockTotalProducts, total)
		mockUserData.AssertExpectations(t)
		mockProductData.AssertExpectations(t)
	})

	t.Run("invalid id", func(t *testing.T) {
		result, total, err := productService.GetProductByUserId(0, 0, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, total)
		assert.Equal(t, "id not valid", err.Error())
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserData.On("SelectProfileById", uint(2)).Return(nil, errors.New("user not found"))

		result, total, err := productService.GetProductByUserId(2, 0, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, total)
		assert.Equal(t, "user not found", err.Error())
		mockUserData.AssertExpectations(t)
	})

	t.Run("user not a merchant", func(t *testing.T) {
		mockUser := user.Core{
			ID:   3,
			Role: "Customer",
		}

		mockUserData.On("SelectProfileById", uint(3)).Return(&mockUser, nil)

		result, total, err := productService.GetProductByUserId(3, 0, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, total)
		assert.Equal(t, "this not a merchant", err.Error())
		mockUserData.AssertExpectations(t)
	})

	t.Run("product not found", func(t *testing.T) {
		mockUser := user.Core{
			ID:   4,
			Role: "Merchant",
		}

		mockUserData.On("SelectProfileById", uint(4)).Return(&mockUser, nil)
		mockProductData.On("SelectProductByUserId", uint(4), 0, 10).Return(nil, errors.New("product not found"))

		result, total, err := productService.GetProductByUserId(4, 0, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, total)
		assert.Equal(t, "product not found", err.Error())
		mockUserData.AssertExpectations(t)
		mockProductData.AssertExpectations(t)
	})

	t.Run("count product error", func(t *testing.T) {
		mockUser := user.Core{
			ID:   5,
			Role: "Merchant",
		}
		mockProducts := []product.Core{
			{ID: 1, ProductName: "Product 1", UserID: 5},
			{ID: 2, ProductName: "Product 2", UserID: 5},
		}

		mockUserData.On("SelectProfileById", uint(5)).Return(&mockUser, nil)
		mockProductData.On("SelectProductByUserId", uint(5), 0, 10).Return(mockProducts, nil)
		mockProductData.On("CountProductByUserId", uint(5)).Return(0, errors.New("count error"))

		result, total, err := productService.GetProductByUserId(5, 0, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, 0, total)
		assert.Equal(t, "product not found", err.Error())
		mockUserData.AssertExpectations(t)
		mockProductData.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockProductData := new(mocks.ProductData)
	mockUserData := new(mocks.UserData)
	validate := validator.New()

	service := &productService{
		productData: mockProductData,
		userData:    mockUserData,
		validate:    validate,
	}

	t.Run("success update product", func(t *testing.T) {
		productID := uint(1)
		userID := uint(1)
		input := product.Core{
			UserID:      userID,
			ProductName: "Updated Product",
			Description: "Updated Description",
			Price:       200,
		}
		productFromDB := &product.Core{
			UserID:      userID,
			ProductName: "Old Product",
			Description: "Old Description",
			Price:       100,
		}

		mockProductData.On("SelectProductById", productID).Return(productFromDB, nil)
		mockProductData.On("Update", productID, input).Return(nil)

		err := service.Update(productID, input)
		assert.Nil(t, err)
		mockProductData.AssertExpectations(t)
	})

	t.Run("user not authorized", func(t *testing.T) {
		productID := uint(1)
		userID := uint(2)
		input := product.Core{
			UserID:      userID,
			ProductName: "Updated Product",
			Description: "Updated Description",
			Price:       200,
		}
		productFromDB := &product.Core{
			UserID:      uint(1),
			ProductName: "Old Product",
			Description: "Old Description",
			Price:       100,
		}

		mockProductData.On("SelectProductById", productID).Return(productFromDB, nil)

		err := service.Update(productID, input)
		assert.EqualError(t, err, "user not authorized")
		mockProductData.AssertExpectations(t)
	})

	t.Run("failed to update product", func(t *testing.T) {
		productID := uint(1)
		userID := uint(1)
		input := product.Core{
			UserID:      userID,
			ProductName: "Updated Product",
			Description: "Updated Description",
			Price:       200,
		}
		productFromDB := &product.Core{
			UserID:      userID,
			ProductName: "Old Product",
			Description: "Old Description",
			Price:       100,
		}

		mockProductData.On("SelectProductById", productID).Return(productFromDB, nil)
		mockProductData.On("Update", productID, input).Return(nil)

		err := service.Update(productID, input)
		assert.NoError(t, err)
		mockProductData.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	productDataMock := &mocks.ProductData{}
	userDataMock := &mocks.UserData{}

	productService := &productService{
		productData: productDataMock,
		userData:    userDataMock,
	}

	productID := uint(1)
	userID := uint(123)

	t.Run("Successful deletion", func(t *testing.T) {
		productDataMock.On("SelectProductById", productID).Return(&product.Core{UserID: userID}, nil)
		productDataMock.On("Delete", productID).Return(nil)

		err := productService.Delete(productID, userID)

		assert.NoError(t, err)
		productDataMock.AssertCalled(t, "SelectProductById", productID)
		productDataMock.AssertCalled(t, "Delete", productID)
	})
}
