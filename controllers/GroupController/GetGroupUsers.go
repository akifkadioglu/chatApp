package groupcontroller

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputGroupUsers struct {
	GroupId int `json:"group_id" validate:"required"`
}

func GetGroupUsers(c echo.Context) error {
	var groupUsers []models.GroupUser
	var input inputGroupUsers
	db := database.DBManager()

	c.Bind(&input)
	err := groupUsersValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}
	db.Joins("From").Where("`group_id` = " + strconv.Itoa(input.GroupId)).Find(&groupUsers)
	return c.JSON(http.StatusOK, map[string][]models.GroupUser{
		"groupUsers": groupUsers,
	})
}

func groupUsersValidate(c echo.Context, input inputGroupUsers) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
