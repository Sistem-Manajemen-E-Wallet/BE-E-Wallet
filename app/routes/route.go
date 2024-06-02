package routes

import (
	"e-wallet/app/middlewares"
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
	userData := userData.New(db, walletDataService)
	userService := userService.New(userData, hashService)
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
}
