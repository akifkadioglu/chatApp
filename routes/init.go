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
	addr := env.GoDotEnvVariable("APP_HOST") + ":" + env.GoDotEnvVariable("APP_PORT")
	E.Logger.Fatal(E.Start(addr))
}
