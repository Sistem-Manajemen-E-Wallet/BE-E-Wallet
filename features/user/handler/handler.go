package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/user"
	"e-wallet/utils/responses"
	"e-wallet/utils/upload"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceInterface
}

func New(us user.ServiceInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) UpdateProfilePicture(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	_, err := uh.userService.GetProfileUser(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	formHeader, err := c.FormFile("profile_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+err.Error(), nil))
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error formfile: "+err.Error(), nil))
	}

	uploadUrl, err := upload.ImageUploadHelper(formFile)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error upload: "+err.Error(), nil))
	}

	inputCore := user.Core{
		ProfilePicture: uploadUrl,
	}

	errUpdate := uh.userService.UpdateProfilePicture(uint(idToken), inputCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update data: "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success upload", uploadUrl))
}

func (uh *UserHandler) RegisterCustomer(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := user.Core{
		Name:       newUser.Name,
		Email:      newUser.Email,
		Phone:      newUser.PhoneNumber,
		Pin:        newUser.Pin,
		PinConfirm: newUser.PinConfirm,
		Role:       "Customer",
	}
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "pin") {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse(errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error insert data: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success add customer", nil))

}

func (uh *UserHandler) RegisterMerchant(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := user.Core{
		Name:       newUser.Name,
		Email:      newUser.Email,
		Phone:      newUser.PhoneNumber,
		Pin:        newUser.Pin,
		PinConfirm: newUser.PinConfirm,
		Role:       "Merchant",
	}
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "pin") {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse(errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error insert data: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success add merchant", nil))

}

func (uh *UserHandler) GetProfileUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := uh.userService.GetProfileUser(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	userResponse := UserResponse{
		ID:             result.ID,
		Name:           result.Name,
		Email:          result.Email,
		Address:        result.Address,
		Role:           result.Role,
		ProfilePicture: result.ProfilePicture,
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success read data", userResponse))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	tx := uh.userService.Delete(uint(idToken))
	if tx != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error delete data: "+tx.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete user", nil))
}

func (uh *UserHandler) Update(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	updateUser := UserUpdateRequest{}
	errBind := c.Bind(&updateUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateCore := user.Core{
		Name:    updateUser.Name,
		Email:   updateUser.Email,
		Phone:   updateUser.PhoneNumber,
		Address: updateUser.Address,
	}

	errUpdate := uh.userService.Update(uint(idToken), updateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update data: "+errUpdate.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success update user", nil))
}

func (uh *UserHandler) Login(c echo.Context) error {
	loginUser := LoginRequest{}
	errBind := c.Bind(&loginUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	login, token, errLogin := uh.userService.Login(loginUser.PhoneNumber, loginUser.Pin)
	if errLogin != nil {
		if strings.Contains(errLogin.Error(), "record not found") {
			return c.JSON(http.StatusNotFound, responses.WebJSONResponse(errLogin.Error(), nil))
		}
		if strings.Contains(errLogin.Error(), "wrong pin") {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse(errLogin.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse(errLogin.Error(), nil))
	}

	var resultResponse = map[string]any{
		"id":    login.ID,
		"name":  login.Name,
		"role":  login.Role,
		"token": token,
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success login", resultResponse))
}
