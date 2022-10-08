package groupcontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserGroups(c echo.Context) error {
	var usergroups []models.GroupUser
	db := database.DBManager()
	result := db.Joins("Group").Where("`from_id` = " + strconv.Itoa(int(helpers.User(c).ID))).Order("Group.name").Find(&usergroups)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, map[string][]models.GroupUser{
		"groups": usergroups,
	})
}
