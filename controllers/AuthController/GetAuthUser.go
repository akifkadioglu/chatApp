package authcontroller

import (
	"chatApp/helpers"
	"chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAuthUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]models.JwtCustomClaims{
		"getAuthUser": *helpers.User(c),
	})
}
