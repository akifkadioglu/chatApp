package usercontrollers

import (
	"chatApp/database"
	models "chatApp/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type inputsSearchUsers struct {
	Username string `json:"username" validate:"required"`
}

func SearchUser(c echo.Context) error {
	var users []models.User
	var input inputsSearchUsers

	c.Bind(&input)
	err := searchUserValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}

	db := database.DBManager()
	db.Where(`name LIKE "%` + input.Username + `%"`).Find(&users)
	return c.JSON(http.StatusOK, map[string][]models.User{
		"users": users,
	})
}

func searchUserValidate(c echo.Context, input inputsSearchUsers) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
