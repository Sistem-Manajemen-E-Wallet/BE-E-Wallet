package handler

type TopUpRequest struct {
	Amount      int    `json:"amount" form:"amount"`
	ChannelBank string `json:"channel_bank" form:"channel_bank"`
}
