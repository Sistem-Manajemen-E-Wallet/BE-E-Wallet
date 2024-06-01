package handler

type ProductRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}

type ProductUpdateRequest struct {
	ProductName string `json:"product_name" form:"product_name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}
