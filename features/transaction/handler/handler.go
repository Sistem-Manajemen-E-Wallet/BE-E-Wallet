package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/transaction"
	"e-wallet/utils/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	ts transaction.ServiceInterface
}

func New(ts transaction.ServiceInterface) *transactionHandler {
	return &transactionHandler{
		ts: ts,
	}
}

func (th *transactionHandler) CreateTransaction(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newTransaction := CustomerRequest{}
	errBind := c.Bind(&newTransaction)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := transaction.Core{
		UserID:     uint(idToken),
		OrderID:    newTransaction.OrderID,
		ProductID:  newTransaction.ProductID,
		Quantity:   newTransaction.Quantity,
		Additional: newTransaction.Additional,
	}

	err := th.ts.Create(inputCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error create transaction: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create transaction: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create transaction", nil))
}

func (th *transactionHandler) GetTransactionByMerchantId(c echo.Context) error {
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

	result, totalTransactions, err := th.ts.GetTransactionByMerchantId(uint(idToken), offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}
	if len(result) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("you don't have any transactions", nil))
	}

	data := toCoreList(result)

	response := map[string]interface{}{
		"page":       page,
		"limit":      limit,
		"totalItems": totalTransactions,
		"totalPages": (totalTransactions + limit - 1) / limit,
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponseMeta("success get all transactions", response, data))
}
func (th *transactionHandler) GetTransactionById(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	result, err := th.ts.GetTransactionById(uint(idToken), uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get data: "+err.Error(), nil))
	}

	response := toResponse(*result)
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get transactions", response))
}

func (th *transactionHandler) UpdateStatusProgress(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	updateStatusProgress := StatusProgressRequest{}
	errBind := c.Bind(&updateStatusProgress)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	statusCore := transaction.Core{
		StatusProgress: updateStatusProgress.StatusProgress,
	}

	err2 := th.ts.UpdateStatusProgress(uint(idToken), uint(idConv), statusCore)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update data: "+err2.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success update status progress", nil))
}
