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

type inputUpdateProfile struct {
	Name      string `json:"name"`
	Username  string `json:"username" validate:"min=3"`
	Biography string `json:"biography"`
}

func UpdateProfile(c echo.Context) error {
	var input inputUpdateProfile
	var user models.User
	c.Bind(&input)
	err := updateProfileValidate(c, input)
	if err != nil {
		return echo.ErrBadRequest
	}
	db := database.DBManager()
	db.Where("`id` = " + strconv.Itoa(int(helpers.User(c).ID))).First(&user)

	user.Name = input.Name
	user.Username = input.Username
	user.Biography = input.Biography
	
	result := db.Save(&user)

	if result.Error != nil {
		return c.JSON(http.StatusOK, map[string]bool{
			"isUpdated": false,
		})
	}
	return c.JSON(http.StatusOK, map[string]bool{
		"isUpdated": true,
	})
}
func updateProfileValidate(c echo.Context, input inputUpdateProfile) error {
	v := validator.New()
	err := v.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
