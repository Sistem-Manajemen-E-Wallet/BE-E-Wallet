package configs

import (
	"os"
	"strconv"
)

var (
	CLOUDINARY_CLOUD_NAME    string
	CLOUDINARY_API_KEY       string
	CLOUDINARY_API_SECRET    string
	CLOUDINARY_UPLOAD_FOLDER string
	JWT_SECRET               string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.DB_USERNAME = os.Getenv("DBUSER")
	app.DB_PASSWORD = os.Getenv("DBPASS")
	app.DB_HOSTNAME = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error conver dbport")
	}
	app.DB_PORT = portConv
	app.DB_NAME = os.Getenv("DBNAME")
	JWT_SECRET = os.Getenv("JWTSECRET")
	CLOUDINARY_CLOUD_NAME = os.Getenv("CLOUDINARY_CLOUD_NAME")
	CLOUDINARY_API_KEY = os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET = os.Getenv("CLOUDINARY_API_SECRET")
	CLOUDINARY_UPLOAD_FOLDER = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
