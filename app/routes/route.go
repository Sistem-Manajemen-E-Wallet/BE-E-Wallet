package routes

import (
	"e-wallet/app/middlewares"
	productData "e-wallet/features/product/data"
	productHandler "e-wallet/features/product/handler"
	productService "e-wallet/features/product/service"
	userData "e-wallet/features/user/data"
	userHandler "e-wallet/features/user/handler"
	userService "e-wallet/features/user/service"
	encrypts "e-wallet/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	dataService := userData.New(db)
	userService := userService.New(dataService, hashService)
	userHandlerAPI := userHandler.New(userService)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users/customer", userHandlerAPI.RegisterCustomer)
	e.POST("/users/merchant", userHandlerAPI.RegisterMerchant)
	e.GET("/users", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.POST("/users/changeprofilepicture", userHandlerAPI.UpdateProfilePicture, middlewares.JWTMiddleware())

	dataProduct := productData.New(db)
	productService := productService.New(dataProduct, dataService)
	productHandler := productHandler.New(productService)

	e.GET("/products", productHandler.GetAllProduct, middlewares.JWTMiddleware())
	e.POST("/products", productHandler.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", productHandler.GetProductByID, middlewares.JWTMiddleware())
	e.PUT("/products/:id", productHandler.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productHandler.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/users/:id/products", productHandler.GetProductByUserID, middlewares.JWTMiddleware())
	e.POST("/products/:id/images", productHandler.UpdateProductImages, middlewares.JWTMiddleware())
}
