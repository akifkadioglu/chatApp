package usercontrollers

import (
	"chatApp/database"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputGetUser struct {
	ID int `json:"id" validate:"required"`
}

func GetSingleUser(c echo.Context) error {
	var user models.User
	var input inputGetUser
	db := database.DBManager()
	c.Bind(&input)
	err := getUserValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}
	result := db.Where("`id` =" + strconv.Itoa(input.ID)).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, map[string]models.User{
		"user": user,
	})
}

func getUserValidate(c echo.Context, input inputGetUser) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}
