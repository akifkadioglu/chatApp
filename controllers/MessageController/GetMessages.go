package messagecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMessages(c echo.Context) error {
	var message []models.Message
	db := database.DBManager()

	to_id := c.QueryParam("to_id")

	result := db.Where("`to_id` = " + to_id).Where("`from_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Or("`to_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Where("`from_id` = " + to_id).Order("created_at desc").Find(&message)

	if result.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string][]models.Message{
		"messages": message,
	})
}
