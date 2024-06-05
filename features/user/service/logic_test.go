package service

import (
	"e-wallet/features/user"
	"e-wallet/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProfilePicture(t *testing.T) {
	t.Run("Success update profile picture", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		input := user.Core{ProfilePicture: "new_picture.jpg"}

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{}, nil)
		repoUserMock.On("UpdateProfilePicture", uint(1), input).Return(nil)

		srv := New(repoUserMock, hashMock)

		err := srv.UpdateProfilePicture(uint(1), input)
		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("Error update profile picture due to user not found", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		repoUserMock.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found"))

		srv := New(repoUserMock, hashMock)
		input := user.Core{ProfilePicture: "new_picture.jpg"}
		err := srv.UpdateProfilePicture(uint(1), input)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found. you must login first")
		repoUserMock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repoUserMock := new(mocks.UserData)
	hashMock := new(mocks.Hash)

	t.Run("success create user", func(t *testing.T) {
		input := user.Core{
			Name:       "alta",
			Email:      "alta@mail.com",
			Pin:        "123456",
			PinConfirm: "123456",
			Phone:      "11111",
		}

		hashedPin := "hashed_password"
		hashMock.On("HashPassword", input.Pin).Return(hashedPin, nil)
		repoUserMock.On("Insert", mock.Anything).Return(nil)

		srv := New(repoUserMock, hashMock)

		err := srv.Create(input)
		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
		hashMock.AssertExpectations(t)
	})

	t.Run("failed create user due to invalid input", func(t *testing.T) {
		invalidInput := user.Core{}

		srv := New(repoUserMock, hashMock)
		err := srv.Create(invalidInput)

		assert.Error(t, err)
		assert.EqualError(t, err, "[validation] nama/email/pin/phone tidak boleh kosong")
	})

	t.Run("failed create user due to pin mismatch", func(t *testing.T) {
		input := user.Core{
			Name:       "alta",
			Email:      "alta@mail.com",
			Pin:        "123456",
			PinConfirm: "123457",
			Phone:      "11111",
		}

		srv := New(repoUserMock, hashMock)
		err := srv.Create(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "[validation] pin tidak sama")
	})

	t.Run("failed create user due to hash error", func(t *testing.T) {
		input := user.Core{
			Name:       "alta",
			Email:      "alta@mail.com",
			Pin:        "123456",
			PinConfirm: "123456",
			Phone:      "11111",
		}

		hashMock.On("HashPassword", input.Pin).Return("", errors.New("hash error"))

		srv := New(repoUserMock, hashMock)
		err := srv.Create(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "hash error")
		hashMock.AssertExpectations(t)
	})

	t.Run("failed create user due to insert error", func(t *testing.T) {
		input := user.Core{
			Name:       "alta",
			Email:      "alta@mail.com",
			Pin:        "123456",
			PinConfirm: "123456",
			Phone:      "11111",
		}

		hashedPin := "hashed_password"
		hashMock.On("HashPassword", input.Pin).Return(hashedPin, nil)
		repoUserMock.On("Insert", mock.Anything).Return(errors.New("insert error"))

		srv := New(repoUserMock, hashMock)
		err := srv.Create(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "insert error")
		repoUserMock.AssertExpectations(t)
		hashMock.AssertExpectations(t)
	})
}

func TestGetProfileUser(t *testing.T) {
	repoUserMock := new(mocks.UserData)
	hashMock := new(mocks.Hash)
	returnData := &user.Core{
		ID:             1,
		Name:           "alta",
		Email:          "alta@mail.com",
		Role:           "user",
		ProfilePicture: "http://cloudinary.co.id/new_picture.jpg",
	}

	t.Run("success get profile user", func(t *testing.T) {
		repoUserMock.On("SelectProfileById", uint(1)).Return(returnData, nil)

		srv := New(repoUserMock, hashMock)
		result, err := srv.GetProfileUser(uint(1))

		assert.NoError(t, err)
		assert.Equal(t, returnData, result)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("failed get profile user due to invalid id", func(t *testing.T) {
		srv := New(repoUserMock, hashMock)
		result, err := srv.GetProfileUser(uint(0))

		assert.Error(t, err)
		assert.EqualError(t, err, "id not valid")
		assert.Nil(t, result)
		repoUserMock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success delete user", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{}, nil)
		repoUserMock.On("Delete", uint(1)).Return(nil)

		srv := New(repoUserMock, hashMock)
		err := srv.Delete(uint(1))

		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("Error deleting user due to user not found", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		repoUserMock.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found"))

		srv := New(repoUserMock, hashMock)
		err := srv.Delete(uint(1))

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found. you must login first")
		repoUserMock.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Success update", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		input := user.Core{
			Name:  "alta",
			Email: "alta@mail.com",
			Pin:   "123456",
			Phone: "11111",
		}

		hashedPin := "hashed_pin"
		hashMock.On("HashPassword", input.Pin).Return(hashedPin, nil)

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{}, nil)
		repoUserMock.On("Update", uint(1), mock.Anything).Return(nil)

		srv := New(repoUserMock, hashMock)
		err := srv.Update(uint(1), input)

		assert.NoError(t, err)
		repoUserMock.AssertExpectations(t)
	})

	t.Run("Error due to user not found", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		repoUserMock.On("SelectProfileById", uint(1)).Return(nil, errors.New("user not found"))

		srv := New(repoUserMock, hashMock)
		input := user.Core{
			Name:  "alta",
			Email: "alta@mail.com",
			Pin:   "123456",
			Phone: "11111",
		}
		err := srv.Update(uint(1), input)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found. you must login first")
		repoUserMock.AssertExpectations(t)
	})

	t.Run("Error hashing pin", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		input := user.Core{
			Name:  "alta",
			Email: "alta@mail.com",
			Pin:   "alta123",
			Phone: "11111",
		}

		repoUserMock.On("SelectProfileById", uint(1)).Return(&user.Core{}, nil)
		hashMock.On("HashPassword", input.Pin).Return("", errors.New("hashing error"))

		srv := New(repoUserMock, hashMock)
		err := srv.Update(uint(1), input)

		assert.Error(t, err)
		assert.EqualError(t, err, "hashing error")
		repoUserMock.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success login", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		phone := "11111"
		pin := "123456"
		hashedPin := "hashed_123456"

		userData := &user.Core{
			ID:  1,
			Pin: hashedPin,
		}

		repoUserMock.On("Login", phone).Return(userData, nil)
		hashMock.On("CheckPasswordHash", hashedPin, pin).Return(true)

		srv := New(repoUserMock, hashMock)
		data, generatedToken, err := srv.Login(phone, pin)

		assert.NoError(t, err)
		assert.Equal(t, userData, data)
		assert.NotEmpty(t, generatedToken)
		repoUserMock.AssertExpectations(t)
		hashMock.AssertExpectations(t)
	})

	t.Run("login failed due to invalid phone number", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)

		phone := "11111"
		pin := "123456"

		repoUserMock.On("Login", phone).Return(nil, errors.New("invalid phone number"))

		srv := New(repoUserMock, nil)
		_, _, err := srv.Login(phone, pin)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid phone number")
		repoUserMock.AssertExpectations(t)
	})

	t.Run("login failed due to invalid pin", func(t *testing.T) {
		repoUserMock := new(mocks.UserData)
		hashMock := new(mocks.Hash)

		phone := "11111"
		pin := "123456"
		hashedPin := "hashed_123456"

		userData := &user.Core{
			ID:  1,
			Pin: hashedPin,
		}

		repoUserMock.On("Login", phone).Return(userData, nil)
		hashMock.On("CheckPasswordHash", hashedPin, pin).Return(false)

		srv := New(repoUserMock, hashMock)
		data, _, err := srv.Login(phone, pin)

		assert.Error(t, err)
		assert.EqualError(t, err, "[validation] pin tidak sesuai")
		assert.Nil(t, data)
		repoUserMock.AssertExpectations(t)
		hashMock.AssertExpectations(t)
	})
}
