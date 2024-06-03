package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/topups"
	"e-wallet/utils/responses"
	"github.com/labstack/echo/v4"
	"net/http"
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
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error topup notification: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success topup notification", nil))
}
