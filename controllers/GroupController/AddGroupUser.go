package groupcontroller

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputAddGroupUser struct {
	FromID  int `json:"from_id" validate:"required"`
	GroupID int `json:"group_id" validate:"required"`
}

func AddGroupUser(c echo.Context) error {
	var groupUser models.GroupUser
	var input inputAddGroupUser
	db := database.DBManager()
	c.Bind(input)
	err := addGroupUserValidate(c, input)
	if err != nil {
		return err
	}
	result := db.Where("`from_id` = " + strconv.Itoa(input.FromID)).Where("`group_id` = " + strconv.Itoa(input.GroupID)).Find(&groupUser)
	if result.Error != nil && result.RowsAffected > 0 {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	groupUser.FromId = input.FromID
	groupUser.GroupId = input.GroupID
	result = db.Create(&groupUser)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func addGroupUserValidate(c echo.Context, input inputAddGroupUser) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
