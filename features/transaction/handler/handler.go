package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/transaction"
	"e-wallet/features/transaction/service"
	"e-wallet/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	ts service.TransactionService
}

func New(ts service.TransactionService) *transactionHandler {
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
