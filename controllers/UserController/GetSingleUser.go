package usercontrollers

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Profile struct {
	User      models.User `json:"user" `
	IsContact bool        `json:"is_contact" gorm:"default:false"`
}

func GetSingleUser(c echo.Context) error {
	var user models.User
	var contact models.Contact
	var profile Profile

	db := database.DBManager()
	input := c.QueryParam("to_id")

	result := db.First(&user, input)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	profile.User = user
	result = db.Where("`to_id` =" + input).Where("`user_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Find(&contact)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	if result.RowsAffected > 0 {
		profile.IsContact = true
	}

	return c.JSON(http.StatusOK, map[string]Profile{
		"singleUser": profile,
	})
}
