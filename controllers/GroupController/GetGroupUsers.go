package groupcontroller

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGroupUsers(c echo.Context) error {
	var groupUsers []models.GroupUser
	db := database.DBManager()
	input := c.QueryParam("group_id")

	result := db.Joins("From").Where("`group_id` = " + input).Find(&groupUsers)
	if result.Error != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, map[string][]models.GroupUser{
		"groupUsers": groupUsers,
	})
}
