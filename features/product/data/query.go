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

func (p *productQuery) SelectAllProduct(offset, limit int) ([]product.Core, error) {
	var productGorm []Product
	tx := p.db.Offset(offset).Limit(limit).Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var productCore []product.Core
	for _, v := range productGorm {
		productCore = append(productCore, product.Core{
			ID:            v.ID,
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

func (p *productQuery) SelectProductById(id uint) (*product.Core, error) {
	var productGorm Product
	tx := p.db.Preload("User").First(&productGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &product.Core{
		ID:            productGorm.ID,
		UserID:        productGorm.UserID,
		MerchantName:  productGorm.User.Name,
		ProductName:   productGorm.ProductName,
		Description:   productGorm.Description,
		Price:         productGorm.Price,
		ProductImages: productGorm.ProductImages,
		CreatedAt:     productGorm.CreatedAt,
		UpdatedAt:     productGorm.UpdatedAt,
	}, nil
}

func (p *productQuery) SelectProductByUserId(id uint, offset, limit int) ([]product.Core, error) {
	var productGorm []Product
	tx := p.db.Where("user_id = ?", id).Offset(offset).Limit(limit).Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productCore []product.Core
	for _, v := range productGorm {
		productCore = append(productCore, product.Core{
			ID:            v.ID,
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

func (p *productQuery) Update(id uint, input product.Core) error {
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

func (p *productQuery) Delete(id uint) error {
	tx := p.db.Delete(&Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (p *productQuery) CountProductByUserId(id uint) (int, error) {
	var count int64
	tx := p.db.Model(&Product{}).Where("user_id = ?", id).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

func (p *productQuery) CountProduct() (int, error) {
	// count product
	var count int64
	tx := p.db.Model(&Product{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

func (p *productQuery) SearchProducts(offset, limit int, search string) ([]product.Core, error) {
	var productGorm []Product
	tx := p.db.Offset(offset).Limit(limit).Where("product_name LIKE ?", "%"+search+"%").Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var productCore []product.Core
	for _, v := range productGorm {
		productCore = append(productCore, product.Core{
			ID:            v.ID,
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

func (p *productQuery) CountProductBySearch(search string) (int, error) {
	var count int64
	tx := p.db.Model(&Product{}).Where("product_name LIKE ?", "%"+search+"%").Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}
