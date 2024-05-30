package main

import (
	"e-wallet/app/configs"
	"e-wallet/app/databases"
	"e-wallet/app/migrations"
	"e-wallet/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMySql := databases.InitDBMysql(cfg)
	// dbMysql := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	migrations.InitialMigration()
	routes.InitRouter(e, dbMySql)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())

	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}

// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
// }))
