package routes

import (
	"chatApp/adapter"
	"chatApp/env"
	"github.com/labstack/echo/v4"
)

var E = echo.New()
var Network = E.Group("", adapter.ConsoleAdapter)

func Init() {
	Web()
	Api()
	port := ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(port))
}
