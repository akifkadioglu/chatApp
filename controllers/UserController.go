package controllers

import (
	"chatApp/database"
	helpers "chatApp/helpers"
	models "chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// All users
func GetUsers(c echo.Context) error {
	var users []models.User
	db := database.DBManager()
	helpers.SendEmail("akif.kadioglu.28@gmail.com", "naber", "naber knk")
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
