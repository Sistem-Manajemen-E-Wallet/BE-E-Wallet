package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/transaction"
	"e-wallet/utils/responses"
	"net/http"
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
			return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error insert data: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error insert data: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create transaction", nil))
}

func (th *transactionHandler) GetTransactionByMerchantId(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	result, err := th.ts.GetTransactionByMerchantId(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	data := toCoreList(result)

	if len(result) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all transactions", nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all transactions", data))
}
