package messagecontroller

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGroupMessages(c echo.Context) error {
	var message []models.GroupMessages
	db := database.DBManager()

	group_id := c.QueryParam("group_id")

	result := db.Joins("User").Where("`group_id` = " + group_id).Order("created_at desc").Find(&message)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string][]models.GroupMessages{
		"groupMessages": message,
	})
}
