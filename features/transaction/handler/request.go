package handler

type CustomerRequest struct {
	OrderID    int    `json:"order_id" form:"order_id"`
	ProductID  uint   `json:"product_id" form:"product_id"`
	Quantity   int    `json:"quantity" form:"quantity"`
	Additional string `json:"additional" form:"additional"`
}

type StatusProgressRequest struct {
	StatusProgress string `json:"status_progress" form:"status_progress"`
}
