package data

import (
	"e-wallet/features/product"
	"e-wallet/features/user"

	"gorm.io/gorm"
)

type productQuery struct {
	db       *gorm.DB
	userData user.DataInterface
}

func New(db *gorm.DB, ud user.DataInterface) product.DataInterface {
	return &productQuery{
		db:       db,
		userData: ud,
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
	tx := p.db.First(&productGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result, err := p.userData.SelectProfileById(productGorm.UserID)
	if err != nil {
		return nil, err
	}

	return &product.Core{
		ID:            productGorm.ID,
		UserID:        productGorm.UserID,
		MerchantName:  result.Name,
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
