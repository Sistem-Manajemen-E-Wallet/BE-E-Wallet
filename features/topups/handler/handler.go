package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/topups"
	"e-wallet/utils/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type topupHandler struct {
	topupService topups.ServiceInterface
}

func New(data topups.ServiceInterface) *topupHandler {
	return &topupHandler{
		topupService: data,
	}
}

func (th *topupHandler) CreateTopup(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	newTopup := TopUpRequest{}
	errBind := c.Bind(&newTopup)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := topups.Core{
		UserID:      idToken,
		Amount:      float64(newTopup.Amount),
		ChannelBank: newTopup.ChannelBank,
	}

	data, err := th.topupService.Create(inputCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create data: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create topup", toResponse(data)))
}

func (th *topupHandler) TopUpNotification(c echo.Context) error {

	notification := TopUpNotificationRequest{}
	errBind := c.Bind(&notification)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := topups.Core{
		OrderID: notification.OrderID,
		Status:  notification.TransactionStatus,
	}

	err := th.topupService.Update(inputCore)
	if err != nil {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("error topup notification: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success topup notification", nil))
}

func (th *topupHandler) GetByUserID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	results, err := th.topupService.GetByUserID(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	if len(results) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all topups", nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all topups", toResponses(results)))
}

func (th *topupHandler) GetByID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	results, err := th.topupService.GetByID(idInt, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get topup", toResponse(results)))

}
