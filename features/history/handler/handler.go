package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/history"
	"e-wallet/utils/responses"
	"net/http"
	"strconv"

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

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	result, totalHistory, err := hh.HistoryService.GetAllHistory(uint(idToken), offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	data := toCoreList(result)

	if len(result) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("empty", nil))
	}

	response := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total_items": totalHistory,
		"total_pages": (totalHistory + limit - 1) / limit,
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponseMeta("success get all history", response, data))
}
