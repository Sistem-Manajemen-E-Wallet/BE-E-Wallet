package service

import (
	"e-wallet/features/product"
	"e-wallet/features/user"
	"errors"
)

type productService struct {
	productData product.DataInterface
	userData    user.DataInterface
}

func New(data product.DataInterface, userData user.DataInterface) product.ServiceInterface {
	return &productService{
		productData: data,
		userData:    userData,
	}
}

func (p productService) GetAllProduct() ([]product.Core, error) {

	result, err := p.productData.SelectAllProduct()
	if err != nil {
		return nil, errors.New("product not found")
	}

	return result, nil
}

func (p productService) Create(input product.Core) error {
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

func (p productService) GetProductById(id int) (*product.Core, error) {
	if id < 0 {
		return nil, errors.New("id not valid")
	}

	result, err := p.productData.SelectProductById(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return result, nil
}

func (p productService) GetProductByUserId(id int) ([]product.Core, error) {
	if id < 0 {
		return nil, errors.New("id not valid")
	}

	userData, err := p.userData.SelectProfileById(uint(id))
	if err != nil {
		return nil, errors.New("user not found")
	}
	if userData.Role != "Merchant" {
		return nil, errors.New("this not a merchant")
	}

	result, err := p.productData.SelectProductByUserId(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return result, nil
}

func (p productService) Update(id int, input product.Core) error {

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

func (p productService) Delete(id int, userID int) error {

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
