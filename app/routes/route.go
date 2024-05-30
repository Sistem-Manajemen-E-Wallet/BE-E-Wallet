package routes

import (
	"e-wallet/app/middlewares"
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
}
