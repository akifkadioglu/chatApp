package messagecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetContactsMessages(c echo.Context) error {
	var messages []models.Message
	db := database.DBManager()
	result := db.Joins("User").Where("`user_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Or("`to_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Order("created_at desc").Find(&messages)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, map[string][]models.Message{
		"messages": messages,
	})
}
