package routes

import (
	"chatApp/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var E = echo.New()

func Init() {
	E.Use(middleware.Logger())
	E.Use(middleware.Recover())
	Web()
	Api()
	port := ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(port))
}
