package handler

import (
	"e-wallet/app/middlewares"
	"e-wallet/features/product"
	"e-wallet/utils/responses"
	"e-wallet/utils/upload"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type productHandler struct {
	productService product.ServiceInterface
}

func New(us product.ServiceInterface) *productHandler {
	return &productHandler{
		productService: us,
	}
}

func (ph *productHandler) GetAllProduct(c echo.Context) error {
	results, err := ph.productService.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	data := toCoreList(results)

	if len(results) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all products", nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all products", data))
}

func (ph *productHandler) CreateProduct(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := product.Core{
		UserID:      idToken,
		ProductName: newProduct.ProductName,
		Description: newProduct.Description,
		Price:       newProduct.Price,
	}

	err := ph.productService.Create(inputCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("product created succesfully", nil))
}

func (ph *productHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	// convert id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	result, err := ph.productService.GetProductById(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	data := toResponse(*result)

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get product", data))
}

func (ph *productHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)

	updateProduct := ProductUpdateRequest{}
	errBind := c.Bind(&updateProduct)
	if errBind != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	inputCore := product.Core{
		UserID:      idToken,
		ProductName: updateProduct.ProductName,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
	}
	err = ph.productService.Update(idInt, inputCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("product update succesfully", nil))
}

func (ph *productHandler) UpdateProductImages(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)

	formHeader, err := c.FormFile("product_image")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+err.Error(), nil))
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error bind data: "+err.Error(), nil))
	}

	defer formFile.Close()
	uploadUrl, err := upload.ImageUploadHelper(formFile)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error upload: "+err.Error(), nil))
	}

	inputCore := product.Core{
		UserID:        idToken,
		ProductImages: uploadUrl,
	}

	err = ph.productService.Update(idInt, inputCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("product update succesfully", nil))
}

func (ph *productHandler) DeleteProduct(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	err = ph.productService.Delete(idInt, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error delete data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.WebJSONResponse("product deleted succesfully", nil))
}

func (ph *productHandler) GetProductByUserID(c echo.Context) error {
	id := c.Param("id")

	idUser, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.WebJSONResponse("error convert data: "+err.Error(), nil))
	}

	results, err := ph.productService.GetProductByUserId(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error read data: "+err.Error(), nil))
	}

	if len(results) == 0 {
		return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all products", nil))
	}

	data := toCoreList(results)
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all products", data))
}
