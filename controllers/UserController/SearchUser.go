package usercontrollers

import (
	"chatApp/database"
	models "chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)



func SearchUser(c echo.Context) error {
	var users []models.User
	input :=c.QueryParam("username") 

	if input == "" {
		return c.JSON(http.StatusOK, map[string][]string{
			"searchUser": make([]string, 0),
		})
	}

	db := database.DBManager()
	db.Where(`username LIKE "%` + input + `%"`).Order("username").Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"searchUser": users,
	})
}
