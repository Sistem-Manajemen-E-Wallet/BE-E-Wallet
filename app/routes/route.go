package routes

import (
	"e-wallet/app/middlewares"
	productData "e-wallet/features/product/data"
	productHandler "e-wallet/features/product/handler"
	productService "e-wallet/features/product/service"
	userData "e-wallet/features/user/data"
	userHandler "e-wallet/features/user/handler"
	userService "e-wallet/features/user/service"
	walletData "e-wallet/features/wallet/data"
	"e-wallet/features/wallet/handler"
	walletService "e-wallet/features/wallet/service"
	encrypts "e-wallet/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	walletDataService := walletData.New(db)
	walletService := walletService.New(walletDataService)
	WalletHandler := handler.New(walletService)

	hashService := encrypts.NewHashService()
	userDataService := userData.New(db, walletDataService)
	userService := userService.New(userDataService, hashService)
	userHandler := userHandler.New(userService)

	e.POST("/login", userHandler.Login)

	e.POST("/users/customer", userHandler.RegisterCustomer)
	e.POST("/users/merchant", userHandler.RegisterMerchant)
	e.GET("/users", userHandler.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandler.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandler.Update, middlewares.JWTMiddleware())
	e.POST("/users/changeprofilepicture", userHandler.UpdateProfilePicture, middlewares.JWTMiddleware())

	e.GET("/wallets", WalletHandler.GetWalletById, middlewares.JWTMiddleware())

	// e.POST("/topup")

	// e.GET("/histories")

	// e.POST("/transactions")
	// e.GET("/transactions")
	// e.PUT("/transactions/:id")

	dataProduct := productData.New(db)
	productService := productService.New(dataProduct, userDataService)
	productHandler := productHandler.New(productService)

	e.GET("/products", productHandler.GetAllProduct, middlewares.JWTMiddleware())
	e.POST("/products", productHandler.CreateProduct, middlewares.JWTMiddleware())
	e.GET("/products/:id", productHandler.GetProductByID, middlewares.JWTMiddleware())
	e.PUT("/products/:id", productHandler.UpdateProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productHandler.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/users/:id/products", productHandler.GetProductByUserID, middlewares.JWTMiddleware())
	e.POST("/products/:id/images", productHandler.UpdateProductImages, middlewares.JWTMiddleware())
}
