package handler

import (
	"e-wallet/features/product"
	"time"
)

type GetAllProductResponse struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	MerchantName  string    `json:"merchant_name,omitempty"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	Price         int       `json:"price"`
	ProductImages string    `json:"product_images"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func toResponse(product product.Core) GetAllProductResponse {
	return GetAllProductResponse{
		ID:            uint(product.ID),
		UserID:        uint(product.UserID),
		MerchantName:  product.MerchantName,
		ProductName:   product.ProductName,
		Description:   product.Description,
		Price:         product.Price,
		ProductImages: product.ProductImages,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

func toCoreList(product []product.Core) []GetAllProductResponse {
	result := []GetAllProductResponse{}
	for key := range product {
		result = append(result, toResponse(product[key]))
	}
	return result
}
