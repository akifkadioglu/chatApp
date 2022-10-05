package controllers

import (
	"chatApp/database"
	"chatApp/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// the user's all contacts
func GetContacts(c echo.Context) error {
	var contact []models.Contact
	db := database.DBManager()

	userId := c.FormValue("userId")
	if userId == "" {
		userId = "0"
	}

	db.Joins("UserTo").Where("`from` =" + userId).Order("UserTo.name").Find(&contact)
	return c.JSON(http.StatusOK, map[string][]models.Contact{
		"contacts": contact,
	})
}

//Connect a user
func ConnectAUser(c echo.Context) error {
	var contact models.Contact
	var user models.User

	db := database.DBManager()

	contact.FromId = 2
	contact.ToId = 10

	db.Create(&contact)
	db.Find(&user, 10)

	return c.JSON(http.StatusOK, map[string]models.User{
		"contacts": user,
	})
}
