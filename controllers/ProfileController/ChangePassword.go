package profilecontroller

import (
	"chatApp/database"
	"chatApp/helpers"
	"chatApp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type inputChangePassword struct {
	CurrentPassword         string `json:"current_password" validate:"required; min:8"`
	NewPassword             string `json:"new_password" validate:"required; min:8"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required; min:8"`
}

func ChangePassword(c echo.Context) error {
	var input inputChangePassword
	var user models.User

	c.Bind(&input)
	err := changePasswordValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}

	db := database.DBManager()
	db.Where("`id` = " + strconv.Itoa(int(helpers.User(c).ID))).First(&user)

	err = helpers.CompareHash(user.Password, input.CurrentPassword)
	if err != nil {
		return echo.ErrBadRequest
	}

	if input.NewPassword != input.NewPasswordConfirmation {
		return echo.ErrBadRequest
	}

	user.Password = helpers.Hash(input.NewPassword)
	result := db.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isChanged": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"isChanged": true,
	})
}

func changePasswordValidate(c echo.Context, input inputChangePassword) error {
	v := validator.New()
	err := v.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
