package contactcontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// the user's all contacts
func GetContacts(c echo.Context) error {
	var contact []models.Contact
	db := database.DBManager()
	db.Joins("To").Where("`from_id` ='" + strconv.Itoa(int(helpers.User(c).ID)) + "'").Order("to.name").Find(&contact)
	return c.JSON(http.StatusOK, map[string][]models.Contact{
		"contacts": contact,
	})
}
