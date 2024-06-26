package service

import (
	"e-wallet/features/history"
	"e-wallet/features/topups"
	"e-wallet/features/user"
	"e-wallet/features/wallet"
	"e-wallet/utils/midtranspay"
	"e-wallet/utils/randomstring"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type topupsService struct {
	topupData   topups.DataInterface
	userData    user.DataInterface
	walletData  wallet.DataInterface
	historyData history.DataInterface
	validate    *validator.Validate
	midtrans    midtranspay.Service
}

func New(data topups.DataInterface, walletData wallet.DataInterface, userData user.DataInterface, historyData history.DataInterface) topups.ServiceInterface {
	return &topupsService{
		topupData:   data,
		walletData:  walletData,
		historyData: historyData,
		userData:    userData,
		validate:    validator.New(),
		midtrans:    midtranspay.New(),
	}
}

func (t *topupsService) Create(input topups.Core) (topups.Core, error) {
	err := t.validate.Struct(input)

	if err != nil {
		return topups.Core{}, errors.New("[validation error] " + err.Error())
	}

	channelBank := strings.ToLower(input.ChannelBank)
	if channelBank != "bca" && channelBank != "cimb" && channelBank != "bni" && channelBank != "bri" {
		return topups.Core{}, errors.New("invalid channel bank")
	}

	userData, err := t.userData.SelectProfileById(uint(int(input.UserID)))
	if err != nil {
		return topups.Core{}, errors.New("user not found")
	}

	if userData.Role == "Merchant" {
		return topups.Core{}, errors.New("user not authorized")
	}

	orderID := randomstring.GenerateRandomString()

	payload := midtranspay.Topup{
		OrderID: orderID,
		Amount:  int(input.Amount),
		Bank:    input.ChannelBank,
	}

	vaNumbers, err := t.midtrans.GetVaNumbers(payload)
	if err != nil {
		return topups.Core{}, errors.New("error getting va numbers")
	}

	topup := topups.Core{
		OrderID:     orderID,
		UserID:      input.UserID,
		Amount:      input.Amount,
		Type:        "Bank Transfer",
		ChannelBank: input.ChannelBank,
		Status:      "Pending",
		VaNumbers:   vaNumbers,
	}

	result, err := t.topupData.Insert(topup)
	if err != nil {
		return topups.Core{}, err
	}

	history := history.Core{
		UserID:  uint(input.UserID),
		TrxName: "Bank " + strings.ToUpper(input.ChannelBank),
		Amount:  int(input.Amount),
		TopUpID: uint(result.ID),
		Type:    "Top-Up",
		Status:  "Pending",
	}

	err = t.historyData.InsertHistory(history)
	if err != nil {
		return topups.Core{}, err
	}

	return result, nil
}

func (t *topupsService) GetByUserID(id int) ([]topups.Core, error) {
	if id <= 0 {
		return []topups.Core{}, errors.New("invalid user id")
	}

	result, err := t.topupData.SelectByUserID(id)
	if err != nil {
		return []topups.Core{}, errors.New("topup not found")
	}

	return result, nil
}

func (t *topupsService) Update(input topups.Core) error {
	topup, err := t.topupData.SelectByOrderID(input.OrderID)
	if err != nil {
		return errors.New("topup not found")
	}

	if topup.Status == "Success" {
		return errors.New("topup status cannot be updated")
	}

	if input.Status != "settlement" {
		return errors.New("invalid topup status")
	}

	topup.Status = "Success"
	tx := t.topupData.Update(int(topup.ID), topup)
	if tx != nil {
		return errors.New("error updating topup")
	}

	wallet, err := t.walletData.GetWalletByUserId(uint(topup.UserID))
	if err != nil {
		return errors.New("wallet not found")
	}

	wallet.Balance += int(topup.Amount)
	tx = t.walletData.UpdateBalanceByTopup(wallet)
	if tx != nil {
		return errors.New("error updating wallet")
	}

	history := history.Core{
		TopUpID: uint(topup.ID),
		Status:  "Success",
	}

	tx = t.historyData.UpdateHistoryTopUp(history)
	if tx != nil {
		return errors.New("error updating history")
	}

	return nil
}

func (t *topupsService) GetByID(id int, userID int) (topups.Core, error) {
	if id <= 0 {
		return topups.Core{}, errors.New("invalid topup id")
	}

	result, err := t.topupData.SelectById(id)
	if err != nil {
		return topups.Core{}, errors.New("topup not found")
	}

	if result.UserID != userID {
		return topups.Core{}, errors.New("user not authorized")
	}

	return result, nil
}
