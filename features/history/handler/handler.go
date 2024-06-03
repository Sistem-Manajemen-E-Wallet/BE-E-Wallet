package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/history"
	"e-wallet/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HistoryHandler struct {
	HistoryService history.ServiceInterface
}

func New(hs history.ServiceInterface) *HistoryHandler {
	return &HistoryHandler{
		HistoryService: hs,
	}
}

func (hh *HistoryHandler) GetAllHistory(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := hh.HistoryService.GetAllHistory(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	data := toCoreList(result)

	if len(result) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("empty", nil))
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all history", data))
}
