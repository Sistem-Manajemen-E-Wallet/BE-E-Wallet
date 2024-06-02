package data

import (
	"e-wallet/features/product"
	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &productQuery{
		db: db,
	}
}

func (p *productQuery) SelectAllProduct() ([]product.Core, error) {
	var productGorm []Product
	tx := p.db.Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var productCore []product.Core
	for _, v := range productGorm {
		productCore = append(productCore, product.Core{
			ID:            int(v.ID),
			UserID:        v.UserID,
			ProductName:   v.ProductName,
			Description:   v.Description,
			Price:         v.Price,
			ProductImages: v.ProductImages,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return productCore, nil
}

func (p *productQuery) Insert(input product.Core) error {
	productGorm := Product{
		UserID:        input.UserID,
		ProductName:   input.ProductName,
		Description:   input.Description,
		Price:         input.Price,
		ProductImages: input.ProductImages,
	}
	tx := p.db.Create(&productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *productQuery) SelectProductById(id int) (*product.Core, error) {
	var productGorm Product
	tx := p.db.First(&productGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product.Core{
		ID:            int(productGorm.ID),
		UserID:        productGorm.UserID,
		ProductName:   productGorm.ProductName,
		Description:   productGorm.Description,
		Price:         productGorm.Price,
		ProductImages: productGorm.ProductImages,
		CreatedAt:     productGorm.CreatedAt,
		UpdatedAt:     productGorm.UpdatedAt,
	}, nil
}

func (p *productQuery) SelectProductByUserId(id int) ([]product.Core, error) {
	var productGorm []Product
	tx := p.db.Where("user_id = ?", id).Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productCore []product.Core
	for _, v := range productGorm {
		productCore = append(productCore, product.Core{
			ID:            int(v.ID),
			UserID:        v.UserID,
			ProductName:   v.ProductName,
			Description:   v.Description,
			Price:         v.Price,
			ProductImages: v.ProductImages,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}
	return productCore, nil
}

func (p *productQuery) Update(id int, input product.Core) error {
	productGorm := Product{
		UserID:        input.UserID,
		ProductName:   input.ProductName,
		Description:   input.Description,
		Price:         input.Price,
		ProductImages: input.ProductImages,
	}
	tx := p.db.Model(&productGorm).Where("id = ?", id).Updates(productGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *productQuery) Delete(id int) error {
	tx := p.db.Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
