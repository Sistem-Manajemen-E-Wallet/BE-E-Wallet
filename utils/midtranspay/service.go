package midtranspay

import (
	"e-wallet/app/configs"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type service struct {
}

type Service interface {
	GetVaNumbers(transaction Topup) (string, error)
}

func New() *service {
	return &service{}
}

func (s *service) GetVaNumbers(topup Topup) (string, error) {
	c := coreapi.Client{}
	c.New(configs.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  topup.OrderID,
			GrossAmt: int64(topup.Amount),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.Bank(topup.Bank),
		},
	}

	coreApiRes, _ := c.ChargeTransaction(chargeReq)

	return coreApiRes.VaNumbers[0].VANumber, nil

}
