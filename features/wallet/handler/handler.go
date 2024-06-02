package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/wallet"
	"e-wallet/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	walletService wallet.ServiceInterface
}

func New(ws wallet.ServiceInterface) *WalletHandler {
	return &WalletHandler{
		walletService: ws,
	}
}

func (wh *WalletHandler) GetWalletById(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, err := wh.walletService.GetWalletById(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	walletResponse := WalletResponse{
		Balance:   result.Balance,
		UpdatedAt: result.UpdatedAt,
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success read data", walletResponse))
}
