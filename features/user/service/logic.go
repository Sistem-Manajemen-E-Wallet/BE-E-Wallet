package service

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/user"
	encrypts "e-wallet/utils"
	"errors"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}
}

func (u *userService) UpdateProfilePicture(id uint, input user.Core) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}

	if result.DeleteAt.IsZero() {
		return u.userData.UpdateProfilePicture(id, input)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Create implements user.ServiceInterface.
func (u *userService) Create(input user.Core) error {
	if input.Name == "" || input.Email == "" || input.Pin == "" || input.Phone == "" {
		return errors.New("name/email/pin/phone cannot be empty")
	}
	if input.Pin != input.PinConfirm {
		return errors.New("pin does not match")
	}
	result, errHash := u.hashService.HashPassword(input.Pin)
	if errHash != nil {
		return errHash
	}
	input.Pin = result

	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements user.ServiceInterface.
func (u *userService) GetProfileUser(id uint) (*user.Core, error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}
	return u.userData.SelectProfileById(id)
}

// Delete implements user.ServiceInterface.
func (u *userService) Delete(id uint) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}
	if result.DeleteAt.IsZero() {
		return u.userData.Delete(id)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Update implements user.ServiceInterface.
func (u *userService) Update(id uint, input user.Core) error {
	result, err := u.userData.SelectProfileById(id)
	if err != nil {
		return errors.New("user not found. you must login first")
	}

	result2, errHash := u.hashService.HashPassword(input.Pin)
	if errHash != nil {
		return errHash
	}
	if input.Pin != "" {
		input.Pin = result2
	}

	if result.DeleteAt.IsZero() {
		return u.userData.Update(id, input)
	} else {
		return errors.New("user not found. you must login first")
	}
}

// Login implements user.ServiceInterface.
func (u *userService) Login(phone string, Pin string) (data *user.Core, token string, err error) {
	data, err = u.userData.Login(phone)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Pin, Pin)
	if !isLoginValid {
		return nil, "", errors.New("wrong pin")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}
