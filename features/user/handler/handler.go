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

func (uh *UserHandler) Register(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := user.Core{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Phone:    newUser.PhoneNumber,
	}
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error insert data: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error insert data: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success add user", nil))

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
	updateUser := UserRequest{}
	errBind := c.Bind(&updateUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateCore := user.Core{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
		Phone:    updateUser.PhoneNumber,
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

	login, token, errLogin := uh.userService.Login(loginUser.Email, loginUser.Password)
	if errLogin != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error login: "+errLogin.Error(), nil))
	}

	var resultResponse = map[string]any{
		"id":    login.ID,
		"name":  login.Name,
		"token": token,
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success login", resultResponse))
}
