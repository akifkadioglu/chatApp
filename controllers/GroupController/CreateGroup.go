package groupcontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type inputCreateGroup struct {
	Name string `json:"name" validate:"required"`
}

func CreateGroup(c echo.Context) error {
	var input inputCreateGroup
	var group models.Group

	c.Bind(&input)
	err := createGroupValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}
	group.Name = input.Name
	db := database.DBManager()
	result := groupTransaction(c, group, db)
	if result != nil {
		return err
	}
	return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}

func groupTransaction(c echo.Context, group models.Group, db gorm.DB) error {
	var groupUser models.GroupUser

	return db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&group)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}

		groupUser.GroupId = int(group.ID)
		groupUser.FromId = int(helpers.User(c).ID)
		groupUser.IsAdmin = true
		result = tx.Create(&groupUser)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return nil
	})
}

func createGroupValidate(c echo.Context, input inputCreateGroup) error {
	v := validator.New()

	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
