package usercontrollers

import (
	"chatApp/database"
	models "chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRandomUsers(c echo.Context) error {
	var users []models.User
	db := database.DBManager()
	db.Order("rand()").Limit(3).Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"randomUser": users,
	})
}
