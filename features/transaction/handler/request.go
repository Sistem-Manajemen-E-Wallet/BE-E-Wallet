package handler

type CustomerRequest struct {
	ProductID  uint   `json:"product_id" form:"product_id"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Additional string `json:"additional" form:"additional"`
}