package controllers

import (
	"chatApp/database"
	models "chatApp/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// All users
func GetUsers(c echo.Context) error {
	var users []models.User
	db := database.DBManager()

	db.Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"users": users,
	})
}

//Search users
func SearchUser(c echo.Context) error {
	var users []models.User
	db := database.DBManager()

	username := c.FormValue("name")

	db.Where(`name LIKE "%` + username + `%"`).Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"users": users,
	})
}
