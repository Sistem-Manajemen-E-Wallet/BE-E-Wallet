package service

import (
	"e-wallet/features/product"
	"e-wallet/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type productService struct {
	productData product.DataInterface
	userData    user.DataInterface
	validate    *validator.Validate
}

func New(data product.DataInterface, userData user.DataInterface) product.ServiceInterface {
	return &productService{
		productData: data,
		userData:    userData,
		validate:    validator.New(),
	}
}

func (p *productService) GetAllProduct(offset, limit int) ([]product.Core, int, error) {

	result, err := p.productData.SelectAllProduct(offset, limit)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	totalProduct, err := p.productData.CountProduct()
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	return result, totalProduct, nil
}

func (p *productService) Create(input product.Core) error {
	errValidate := p.validate.Struct(input)

	if errValidate != nil {
		return errors.New("[validation error] " + errValidate.Error())
	}

	product := product.Core{
		UserID:        input.UserID,
		ProductName:   input.ProductName,
		Description:   input.Description,
		Price:         input.Price,
		ProductImages: "https://res.cloudinary.com/dikdesfng/image/upload/v1717222202/NCI_Visuals_Food_Hamburger_ek9rxe.jpg",
	}

	userData, err := p.userData.SelectProfileById(uint(input.UserID))
	if err != nil {
		return errors.New("user not found")
	}

	if userData.Role != "Merchant" {
		return errors.New("user not authorized")
	}

	return p.productData.Insert(product)
}

func (p *productService) GetProductById(id uint) (*product.Core, error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}

	result, err := p.productData.SelectProductById(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return result, nil
}

func (p *productService) GetProductByUserId(id uint, offset, limit int) ([]product.Core, int, error) {
	if id <= 0 {
		return nil, 0, errors.New("id not valid")
	}

	userData, err := p.userData.SelectProfileById(uint(id))
	if err != nil {
		return nil, 0, errors.New("user not found")
	}
	if userData.Role != "Merchant" {
		return nil, 0, errors.New("this not a merchant")
	}

	result, err := p.productData.SelectProductByUserId(id, offset, limit)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	totalProduct, err := p.productData.CountProductByUserId(id)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	return result, totalProduct, nil
}

func (p *productService) Update(id uint, input product.Core) error {

	product, err := p.productData.SelectProductById(id)
	if err != nil {
		return errors.New("product not found")
	}

	if product.UserID != input.UserID {
		return errors.New("user not authorized")
	}

	err = p.productData.Update(id, input)
	if err != nil {
		return errors.New("product not found")
	}
	return nil
}

func (p *productService) Delete(id uint, userID uint) error {

	product, err := p.productData.SelectProductById(id)
	if err != nil {
		return errors.New("product not found")
	}

	if product.UserID != userID {
		return errors.New("user not authorized")
	}

	err = p.productData.Delete(id)
	if err != nil {
		return errors.New("product not found")
	}
	return nil
}

func (p *productService) SearchProducts(offset, limit int, search string) ([]product.Core, int, error) {
	result, err := p.productData.SearchProducts(offset, limit, search)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	totalProduct, err := p.productData.CountProductBySearch(search)
	if err != nil {
		return nil, 0, errors.New("product not found")
	}

	return result, totalProduct, nil
}
