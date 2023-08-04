package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Databases Databases
	App       App
}

type App struct {
	AppName  string
	AppPort  string
	AppDebug bool
}

type Databases struct {
	Postgres Postgres
	Mysql    Mysql
}

type Postgres struct {
	Host     string
	UserName string
	Password string
	Database string
	Port     string
	Sslmode  string
}

type Mysql struct {
	Host     string
	UserName string
	Database string
	Password string
	Port     string
	Sslmode  string
}

func Apply() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	appDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return err
	}

	app := App{
		AppName:  os.Getenv("APP_NAME"),
		AppPort:  os.Getenv("APP_PORT"),
		AppDebug: appDebug,
	}

	mysql := Mysql{
		UserName: os.Getenv("DATABASES_MYSQL_USER"),
		Password: os.Getenv("DATABASES_MYSQL_PASSWORD"),
		Database: os.Getenv("DATABASES_MYSQL_DATABASE"),
		Port:     os.Getenv("DATABASES_MYSQL_PORT"),
		Sslmode:  os.Getenv("DATABASE_MYSQL_SSLMODE"),
	}
	database := Databases{
		Mysql: mysql,
	}

	AppConfig.App = app
	AppConfig.Databases = database

	return nil
}
