package handler

type TopUpRequest struct {
	Amount      int    `json:"amount" form:"amount"`
	ChannelBank string `json:"channel_bank" form:"channel_bank"`
}

type TopUpNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
}
