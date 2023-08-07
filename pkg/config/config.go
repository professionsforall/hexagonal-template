package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Databases    Databases
	App          App
	ShutdownTime time.Duration
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
	Timeout  time.Duration
}

func Apply() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	shutdownTime := os.Getenv("SHUTDOWN_TIME")
	shutdownTimeDuration, err := time.ParseDuration(shutdownTime)
	if err != nil {
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
	mysqlTimeout := os.Getenv("DATABASES_MYSQL_TIMEOUT")
	mysqlTimeoutDuration, err := time.ParseDuration(mysqlTimeout)
	if err != nil {
		return err
	}
	mysql := Mysql{
		UserName: os.Getenv("DATABASES_MYSQL_USER"),
		Password: os.Getenv("DATABASES_MYSQL_PASSWORD"),
		Database: os.Getenv("DATABASES_MYSQL_DATABASE"),
		Port:     os.Getenv("DATABASES_MYSQL_PORT"),
		Sslmode:  os.Getenv("DATABASE_MYSQL_SSLMODE"),
		Timeout:  mysqlTimeoutDuration,
	}
	database := Databases{
		Mysql: mysql,
	}

	AppConfig.App = app
	AppConfig.Databases = database
	AppConfig.ShutdownTime = shutdownTimeDuration

	return nil
}
